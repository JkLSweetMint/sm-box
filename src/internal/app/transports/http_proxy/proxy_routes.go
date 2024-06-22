package http_proxy

import (
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

	eng.components.Logger.Info().
		Text("The proxy http routes are initialized. ").Write()
}
