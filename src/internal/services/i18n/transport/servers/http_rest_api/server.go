package http_rest_api

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"sm-box/internal/services/i18n/transport/servers/http_rest_api/config"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
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
		Text("The http rest api server is listening... ").Write()

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		if err = srv.app.Listen(srv.conf.Server.Addr, srv.conf.Server.ToFiberListenConfig()); err != nil {
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
