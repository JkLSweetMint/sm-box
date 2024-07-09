package http_rest_api

import (
	"sm-box/pkg/core/components/tracer"
)

// registerRoutes - регистрация маршрутов сервера.
func (srv *server) registerRoutes() error {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal, tracer.LevelTransport)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished() }()
	}

	srv.components.Logger.Info().
		Text("Starting initialization of http rest api server routes... ").Write()

	//var router = srv.router

	srv.components.Logger.Info().
		Text("Http rest api server routes are initialized. ").Write()

	return nil
}