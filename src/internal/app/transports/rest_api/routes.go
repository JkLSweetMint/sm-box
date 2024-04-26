package rest_api

import "sm-box/pkg/core/components/tracer"

// initRoutes - регистрация маршрутов системы.
func (eng *engine) initRoutes() {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal, tracer.LevelTransport)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished() }()
	}

	var router = eng.router.Group("/", eng.components.SystemAccess.Middleware)

	// /users
	{
		var router = router.Group("/users")

		router.Post("/auth", eng.components.SystemAccess.BasicAuth)
	}
}
