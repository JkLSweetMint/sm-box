package rest_api

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"sm-box/internal/app/transports/rest_api/components/access_system"
	"sm-box/internal/app/transports/rest_api/config"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"strings"
	"sync"
	"time"
)

// engine - движок http rest api коробки.
type engine struct {
	app    *fiber.App
	router fiber.Router

	conf *config.Config
	ctx  context.Context

	components *components
}

// components - компоненты движка http rest api коробки.
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

	// Регистрация запросов
	{
		var (
			prefix string
			list   = make([]*fiber.Route, 0, 100)
		)

		// prefix
		{
			if eng.conf.Engine.Name != "" {
				prefix += "/" + eng.conf.Engine.Name
			}

			if eng.conf.Engine.Version != "" {
				prefix += "/" + eng.conf.Engine.Version
			}
		}

		for _, stack := range eng.app.Stack() {
			for _, route := range stack {
				if strings.HasPrefix(route.Path, prefix) {
					eng.components.Logger.Info().
						Format("The route '%s %s' (%d) is registered. ", route.Method, route.Path, len(route.Handlers)).Write()

					list = append(list, route)
				}
			}
		}

		// Система доступа
		{
			if err = eng.components.AccessSystem.RegisterRoutes(list...); err != nil {
				eng.components.Logger.Error().
					Format("An error occurred during the registration of http router routes: '%s'. ", err).Write()
				return
			}
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
