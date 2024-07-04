package rest_api

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"os"
	"path"
	common_models "sm-box/internal/common/objects/models"
	"sm-box/internal/common/transports/rest_api/components/access_system"
	"sm-box/internal/services/i18n/infrastructure/objects/models"
	"sm-box/internal/services/i18n/transports/rest_api/config"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/http/postman"
	"sm-box/pkg/tools/file"
	"sync"
	"time"
)

// engine - движок http rest api сервиса.
type engine struct {
	app    *fiber.App
	router fiber.Router

	conf *config.Config
	ctx  context.Context

	controllers *controllers
	components  *components

	postman *postman.Collection
}

// controllers - контроллеры движка.
type controllers struct {
	Texts interface {
		AssembleDictionary(ctx context.Context, lang string, paths []string) (dictionary models.Dictionary, cErr c_errors.RestAPI)
	}
	Languages interface {
	}
	Identification interface {
		GetToken(ctx context.Context, data string) (token *common_models.JwtTokenInfo, cErr c_errors.RestAPI)
	}
}

// components - компоненты движка http rest api сервиса.
type components struct {
	Logger       logger.Logger
	AccessSystem access_system.AccessSystem
}

// Listen - запуск движка.
func (eng *engine) Listen() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelTransport)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	eng.components.Logger.Info().
		Text("The http rest engine is listening... ").Write()

	// Postman
	{
		var (
			p    = path.Join(env.Paths.SystemLocation, env.Paths.System.Path, "/services", env.Vars.SystemName, "/transports/postman")
			name = fmt.Sprintf("service.%s@%s.json", eng.conf.Engine.Name, eng.conf.Engine.Version)
		)

		// Проверка наличия файла
		{
			var exist bool

			if exist, err = file.Exists(path.Join(p, name)); err != nil {
				eng.components.Logger.Error().
					Format("Could not verify the existence of the postman collection file: '%s'. ", err).Write()
				return
			}

			if exist {
				if err = os.Remove(path.Join(p, name)); err != nil {
					eng.components.Logger.Error().
						Format("The postman collection file could not be deleted: '%s'. ", err).Write()
					return
				}
			} else {
				if err = os.MkdirAll(p, 0666); err != nil {
					eng.components.Logger.Error().
						Format("Failed to create a directory for the postman collection file: '%s'. ", err).Write()
					return
				}
			}
		}

		var fl *os.File

		if fl, err = os.Create(path.Join(p, name)); err != nil {
			eng.components.Logger.Error().
				Format("Failed to create a file for the postman collection: '%s'. ", err).Write()
			return
		}

		defer fl.Close()

		if err = eng.postman.Write(fl, postman.V210); err != nil {
			eng.components.Logger.Error().
				Format("Failed to write postman collection data: '%s'. ", err).Write()
			return
		}
	}

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		if err = eng.app.Listen(eng.conf.Engine.Addr, eng.conf.Engine.ToFiberListenConfig()); err != nil {
			eng.components.Logger.Error().
				Format("An error occurred when starting the http router maintenance: '%s'. ", err).Write()
			return
		}
	}()

	time.Sleep(time.Second)

	if err != nil {
		return
	}

	eng.components.Logger.Info().
		Text("The http rest engine is listened. ").Write()

	wg.Wait()

	return
}

// Shutdown - завершение работы движка.
func (eng *engine) Shutdown() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelTransport)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	eng.components.Logger.Info().
		Text("Shutting down the http rest engine... ").Write()

	if err = eng.app.Shutdown(); err != nil {
		eng.components.Logger.Error().
			Format("An error occurred when completing router maintenance: '%s'. ", err).Write()
		return
	}

	eng.components.Logger.Info().
		Text("The http rest engine is turned off. ").Write()

	return
}