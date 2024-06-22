package http_proxy

import (
	"github.com/gofiber/fiber/v3"
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

	eng.components.Logger.Info().
		Text("Starting the initialization of the fiber http server... ").
		Field("config", eng.conf.Engine.ToFiberConfig()).Write()

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

	eng.initBaseRoutes()
	eng.initProxyRoutes()

	if err = eng.app.ShutdownWithContext(eng.ctx); err != nil {
		return
	}

	eng.components.Logger.Info().
		Text("The fiber http server has been initialized. ").Write()

	return
}
