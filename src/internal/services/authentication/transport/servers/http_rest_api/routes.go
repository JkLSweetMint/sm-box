package http_rest_api

import (
	"sm-box/pkg/core/components/tracer"
)

// registerRoutes - регистрация маршрутов сервера.
func (srv *server) registerRoutes() {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal, tracer.LevelTransport)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished() }()
	}

	srv.components.Logger.Info().
		Text("Starting initialization of http rest api server routes... ").Write()

	var router = srv.router

	// /nginx
	{
		var (
			router = router.Group("/nginx")
		)

		// GET /auth
		{
			router.All("/auth", srv.components.AccessSystem.AuthenticationMiddlewareForRestAPI)
		}
	}

	srv.components.Logger.Info().
		Text("Http rest api server routes are initialized. ").Write()
}
