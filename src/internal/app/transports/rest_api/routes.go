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

	eng.components.Logger.Info().
		Text("Starting initialization of http rest api routes... ").Write()

	var router = eng.router

	// /users
	{
		var router = router.Group("/users")

		router.Post("/auth", eng.components.AccessSystem.BasicUserAuth)
	}

	eng.components.Logger.Info().
		Text("Http rest api routes are initialized. ").Write()
}
