package http_rest_api

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/services/url_shortner/objects/models"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	"strings"
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
		var (
			id    = uuid.New().String()
			route fiber.Route
		)

		router.All("/use/*", func(ctx fiber.Ctx) (err error) {
			var (
				reduce string
				url    *models.ShortUrlInfo
			)

			// Получение сокращения
			{
				reduce = strings.Replace(string(ctx.Request().URI().Path()), strings.Replace(route.Path, "/*", "/", 1), "", 1)
			}

			// Получение url
			{
				var cErr c_errors.RestAPI

				if url, cErr = srv.controllers.Urls.GetByReduceFromRedisDB(ctx.Context(), reduce); cErr != nil {
					srv.components.Logger.Warn().
						Format("Failed to get information on a short route: '%s'. ", cErr).Write()

					if errors.Is(cErr, error_list.ShortUrlNotFound()) {
						return ctx.Redirect().To("/errors/403")
					}

					return ctx.Redirect().To("/errors/50x")
				}
			}

			// Выполнение инструкций
			{
				switch url.Properties.Type {
				case "redirect":
					{
						if err = ctx.Redirect().To(url.Source); err != nil {
							srv.components.Logger.Warn().
								Format("Failed to redirect a remote resource: '%s'. ", err).
								Field("url", url).Write()

							return ctx.Redirect().To("/errors/50x")
						}

						return
					}
				case "proxy":
					{
						var client = new(fasthttp.Client)

						ctx.Request().URI().Update(url.Source)

						fmt.Println(ctx.Request().URI().String())

						if err = client.Do(ctx.Request(), ctx.Response()); err != nil {
							srv.components.Logger.Warn().
								Format("Failed to proxy a remote resource: '%s'. ", err).
								Field("url", url).Write()

							return ctx.Redirect().To("/errors/50x")
						}

						return
					}
				}
			}

			return
		}).Name(id)

		route = srv.app.GetRoute(id)
	}

	srv.components.Logger.Info().
		Text("Http rest api server routes are initialized. ").Write()

	return nil
}
