package rest_api

import (
	"github.com/gofiber/fiber/v3"
	cache_middleware "github.com/gofiber/fiber/v3/middleware/cache"
	compress_middleware "github.com/gofiber/fiber/v3/middleware/compress"
	cors_middleware "github.com/gofiber/fiber/v3/middleware/cors"
	"sm-box/pkg/core/components/tracer"
)

// initFiberApp - инициализация http сервера fiber.
func (eng *engine) initFiberApp() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal, tracer.LevelTransport)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
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
	eng.initRoutes()

	if err = eng.app.ShutdownWithContext(eng.ctx); err != nil {
		return
	}

	return
}
