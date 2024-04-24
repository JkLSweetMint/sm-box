package rest_api

import (
	"context"
	"github.com/gofiber/fiber/v3"
	cache_middleware "github.com/gofiber/fiber/v3/middleware/cache"
	compress_middleware "github.com/gofiber/fiber/v3/middleware/compress"
	cors_middleware "github.com/gofiber/fiber/v3/middleware/cors"
	"sm-box/src/internal/app/transports/rest_api/config"
	"sm-box/src/pkg/core/components/logger"
	"sm-box/src/pkg/core/components/tracer"
	"sync"
	"time"
)

type engine struct {
	app    *fiber.App
	router fiber.Router

	conf *config.Config
	ctx  context.Context

	components *components
}

type components struct {
	logger logger.Logger
}

func (eng *engine) Listen() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelTransport)

		trc.FunctionCall()
		trc.Error(err).FunctionCallFinished(eng)
	}

	eng.components.logger.Info().
		Text("The http rest engine is listening... ").Write()

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		if err = eng.app.Listen(eng.conf.Engine.Addr, eng.conf.Engine.ToFiberListenConfig()); err != nil {
			eng.components.logger.Error().
				Format("An error occurred when starting the http router maintenance: '%s'. ", err).Write()
			return
		}
	}()

	time.Sleep(time.Second)

	if err != nil {
		return
	}

	eng.components.logger.Info().
		Text("The http rest engine is listened. ").Write()

	wg.Wait()

	return
}

func (eng *engine) Shutdown() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelTransport)

		trc.FunctionCall()
		trc.Error(err).FunctionCallFinished(eng)
	}

	eng.components.logger.Info().
		Text("Shutting down the http rest engine... ").Write()

	if err = eng.app.Shutdown(); err != nil {
		eng.components.logger.Error().
			Format("An error occurred when completing router maintenance: '%s'. ", err).Write()
		return
	}

	eng.components.logger.Info().
		Text("The http rest engine is turned off. ").Write()

	return
}

func (eng *engine) initFiberApp() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelTransport)

		trc.FunctionCall()
		trc.Error(err).FunctionCallFinished(eng)
	}

	eng.app = fiber.New(eng.conf.Engine.ToFiberConfig())

	// Маршрутизатор
	{
		var prefix string

		if eng.conf.Engine.Name != "" {
			prefix += "/" + eng.conf.Engine.Name
		}

		if eng.conf.Engine.Version != "" {
			prefix += "/" + eng.conf.Engine.Version
		}

		eng.router = eng.app.Group(prefix)
	}

	// Промежуточные слои
	{
		if eng.conf.Middlewares != nil &&
			eng.conf.Middlewares.Compress != nil &&
			eng.conf.Middlewares.Compress.Enable {
			eng.app.Use(compress_middleware.New(eng.conf.Middlewares.Compress.ToFiberConfig()))
		}

		if eng.conf.Middlewares != nil &&
			eng.conf.Middlewares.Cache != nil &&
			eng.conf.Middlewares.Cache.Enable {
			eng.app.Use(cache_middleware.New(eng.conf.Middlewares.Cache.ToFiberConfig()))
		}

		if eng.conf.Middlewares != nil &&
			eng.conf.Middlewares.Cors != nil &&
			eng.conf.Middlewares.Cors.Enable {
			eng.app.Use(cors_middleware.New(eng.conf.Middlewares.Cors.ToFiberConfig()))
		}
	}

	eng.initBaseRoutes()

	if err = eng.app.ShutdownWithContext(eng.ctx); err != nil {
		return
	}

	return
}
