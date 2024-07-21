package http_rest_api

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
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

	var router = srv.router

	// ALL /use/
	{
		router.All("/use/*", func(ctx fiber.Ctx) (err error) {
			fmt.Println(ctx.Request().String())

			return
		})
	}

	srv.components.Logger.Info().
		Text("Http rest api server routes are initialized. ").Write()

	return nil
}
