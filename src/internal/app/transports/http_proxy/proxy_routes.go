package http_proxy

import (
	"github.com/gofiber/fiber/v3/middleware/proxy"
	"sm-box/pkg/core/components/tracer"
)

// initProxyRoutes - регистрация маршрутов для проксирования.
func (eng *engine) initProxyRoutes() {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal, tracer.LevelTransport)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished() }()
	}

	eng.components.Logger.Info().
		Text("Starting initialization of proxy http routes... ").Write()

	// Сервис аутентификации
	{
		for _, src := range eng.conf.Proxy.Sources {
			eng.components.Logger.Info().
				Text("Registration of the proxying source... ").
				Field("path", src.Path).
				Field("remote_addr", src.RemoteAddr).Write()

			eng.app.All(src.Path, proxy.DomainForward(eng.conf.Engine.Domain, src.RemoteAddr))
		}
	}

	eng.components.Logger.Info().
		Text("The proxy http routes are initialized. ").Write()
}
