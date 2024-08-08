package http_rest_api

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"path"
	"sm-box/internal/app/transport/servers/http/rest_api/config"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	env_mode "sm-box/pkg/core/env/mode"
	"sm-box/pkg/http/postman"
	"sm-box/pkg/tools/file"
	"sync"
	"time"
)

// server - сервер http rest api.
type server struct {
	app    *fiber.App
	router fiber.Router

	conf *config.Config
	ctx  context.Context

	controllers *controllers
	components  *components

	postman *postman.Collection
}

// controllers - контроллеры сервера.
type controllers struct {
}

// components - компоненты сервера.
type components struct {
	Logger logger.Logger
}

// Listen - запуск сервера.
func (srv *server) Listen() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelTransport)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	srv.components.Logger.Info().
		Text("The http rest api server is listening... ").
		Field("config", srv.conf).Write()

	// Postman
	{
		if env.Mode == env_mode.Dev {
			var (
				p    = path.Join(env.Paths.SystemLocation, env.Paths.Src.Path, "/transport/postman", env.Vars.SystemName)
				name = fmt.Sprintf("service.%s@%s.json", srv.conf.Server.Name, srv.conf.Server.Version)
			)

			// Проверка наличия файла
			{
				var exist bool

				if exist, err = file.Exists(path.Join(p, name)); err != nil {
					srv.components.Logger.Error().
						Format("Could not verify the existence of the postman collection file: '%s'. ", err).Write()
					return
				}

				if exist {
					if err = os.Remove(path.Join(p, name)); err != nil {
						srv.components.Logger.Error().
							Format("The postman collection file could not be deleted: '%s'. ", err).Write()
						return
					}
				} else {
					if err = os.MkdirAll(p, 0666); err != nil {
						srv.components.Logger.Error().
							Format("Failed to create a directory for the postman collection file: '%s'. ", err).Write()
						return
					}
				}
			}

			var fl *os.File

			if fl, err = os.Create(path.Join(p, name)); err != nil {
				srv.components.Logger.Error().
					Format("Failed to create a file for the postman collection: '%s'. ", err).Write()
				return
			}

			defer fl.Close()

			if err = srv.postman.Write(fl, postman.V210); err != nil {
				srv.components.Logger.Error().
					Format("Failed to write postman collection data: '%s'. ", err).Write()
				return
			}
		}
	}

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		if err = srv.app.Listen(srv.conf.Server.Addr); err != nil {
			srv.components.Logger.Error().
				Format("An error occurred when starting the http rest api server maintenance: '%s'. ", err).Write()
			return
		}
	}()

	time.Sleep(time.Second)

	if err != nil {
		return
	}

	srv.components.Logger.Info().
		Text("The http rest api server is listened. ").Write()

	wg.Wait()

	return
}

// Shutdown - завершение работы сервера.
func (srv *server) Shutdown() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelTransport)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	srv.components.Logger.Info().
		Text("Shutting down the http rest api server... ").Write()

	if err = srv.app.Shutdown(); err != nil {
		srv.components.Logger.Error().
			Format("An error occurred when completing http rest api server maintenance: '%s'. ", err).Write()
		return
	}

	srv.components.Logger.Info().
		Text("The http rest api server is turned off. ").Write()

	return
}
