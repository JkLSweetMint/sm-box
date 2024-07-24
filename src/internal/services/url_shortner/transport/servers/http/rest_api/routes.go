package http_rest_api

import (
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	error_list "sm-box/internal/common/errors"
	authentication_entities "sm-box/internal/services/authentication/objects/entities"
	"sm-box/internal/services/url_shortner/objects/models"
	"sm-box/internal/services/url_shortner/objects/types"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	"strings"
	"time"
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
				status types.ShortUrlUsageHistoryStatus
				url    *models.ShortUrlInfo
				token  *authentication_entities.JwtSessionToken
			)

			// Запись в историю
			{
				defer func() {
					go func() {
						if url != nil {
							var cErr c_errors.RestAPI

							if cErr = srv.controllers.Urls.WriteCallToHistory(ctx.Context(), url.ID, status, token); cErr != nil {
								srv.components.Logger.Warn().
									Format("The call data could not be recorded in the history: '%s'. ", cErr).
									Field("url", url).
									Field("status", status).Write()
							}
						}
					}()
				}()
			}

			// Получение токена
			{
				if raw := ctx.Cookies(srv.conf.Components.AccessSystem.CookieKeyForSessionToken); len(raw) > 0 {
					token = new(authentication_entities.JwtSessionToken)

					if err = token.Parse(raw); err != nil {
						srv.components.Logger.Error().
							Format("Failed to get session token data: '%s'. ", err).
							Field("raw", raw).Write()

						status = types.ShortUrlUsageHistoryStatusFailed

						err = ctx.Redirect().To("/errors/403")
						return
					}
				}
			}

			// Получение сокращения
			{
				reduce = strings.Replace(string(ctx.Request().URI().Path()), strings.Replace(route.Path, "/*", "/", 1), "", 1)
			}

			// Получение и обработка короткого url
			{
				var cErr c_errors.RestAPI

				// Получение
				{
					if url, cErr = srv.controllers.Urls.GetByReduceFromRedisDB(ctx.Context(), reduce); cErr != nil {
						srv.components.Logger.Warn().
							Format("Failed to get information on a short url: '%s'. ", cErr).Write()

						if errors.Is(cErr, error_list.ShortUrlNotFound()) {
							status = types.ShortUrlUsageHistoryStatusForbidden

							err = ctx.Redirect().To("/errors/403")
							return
						}

						status = types.ShortUrlUsageHistoryStatusFailed

						err = ctx.Redirect().To("/errors/50x")
						return
					}
				}

				// Проверки
				{
					var (
						tm      = time.Now()
						emptyTm time.Time
					)

					// Ещё не начал действовать
					{
						if tm.Before(url.Properties.StartActive) && !url.Properties.StartActive.Equal(emptyTm) {
							url = nil

							srv.components.Logger.Warn().
								Text("The validity period of the short url has not yet begun. ").Write()

							err = ctx.Redirect().To("/errors/403")
							return
						}
					}

					// Уже закончился
					{
						if tm.After(url.Properties.EndActive) && !url.Properties.EndActive.Equal(emptyTm) {
							srv.components.Logger.Warn().
								Text("The validity period of the short url has already been completed. ").Write()

							// Удаление из базы данных Redis
							{
								if cErr = srv.controllers.Urls.RemoveByReduceFromRedisDB(ctx.Context(), url.Reduction); cErr != nil {
									srv.components.Logger.Warn().
										Format("The short url could not be deleted from the redis database: '%s'. ", cErr).
										Field("url", url).Write()
								}
							}

							url = nil

							err = ctx.Redirect().To("/errors/403")

							return
						}
					}

					// Кол-во использований превышено
					{
						if url.Properties.NumberOfUses == 0 {
							srv.components.Logger.Warn().
								Text("The number of uses of the short url  is overestimated. ").Write()

							// Удаление из базы данных Redis
							{
								if cErr = srv.controllers.Urls.RemoveByReduceFromRedisDB(ctx.Context(), url.Reduction); cErr != nil {
									srv.components.Logger.Warn().
										Format("The short url could not be deleted from the redis database: '%s'. ", cErr).
										Field("url", url).Write()
								}
							}

							url = nil

							err = ctx.Redirect().To("/errors/403")

							return
						}
					}

					// Доступов пользователя
					{

					}
				}
			}

			// Выполнение инструкций
			{
				// Обработка
				{
					switch url.Properties.Type {
					case types.ShortUrlTypeRedirect:
						{
							if err = ctx.Redirect().To(url.Source); err != nil {
								srv.components.Logger.Warn().
									Format("Failed to redirect a remote resource: '%s'. ", err).
									Field("url", url).Write()

								status = types.ShortUrlUsageHistoryStatusFailed

								err = ctx.Redirect().To("/errors/50x")
								return
							}
						}
					case types.ShortUrlTypeProxy:
						{
							var client = new(fasthttp.Client)

							ctx.Request().URI().Update(url.Source)

							if err = client.Do(ctx.Request(), ctx.Response()); err != nil {
								srv.components.Logger.Warn().
									Format("Failed to proxy a remote resource: '%s'. ", err).
									Field("url", url).Write()

								status = types.ShortUrlUsageHistoryStatusFailed

								err = ctx.Redirect().To("/errors/50x")
								return
							}
						}
					default:
						{
							srv.components.Logger.Error().
								Text("Unknown type of shortened url. ").
								Field("url", url).Write()

							status = types.ShortUrlUsageHistoryStatusForbidden

							err = ctx.Redirect().To("/errors/403")
							return
						}
					}

					status = types.ShortUrlUsageHistoryStatusSuccess
				}

				// Обновление данных в базе если кол-во использований не бесконечное
				{
					if url.Properties.NumberOfUses != -1 {
						var cErr c_errors.RestAPI

						url.Properties.NumberOfUses--

						if url.Properties.NumberOfUses == 0 {
							if cErr = srv.controllers.Urls.RemoveByReduceFromRedisDB(ctx.Context(), url.Reduction); cErr != nil {
								srv.components.Logger.Warn().
									Format("The short url could not be deleted from the redis database: '%s'. ", cErr).
									Field("url", url).Write()

								return
							}
						} else {
							if cErr = srv.controllers.Urls.UpdateInRedisDB(ctx.Context(), url); cErr != nil {
								srv.components.Logger.Warn().
									Format("Failed to update the short url data in the redis database: '%s'. ", cErr).
									Field("url", url).Write()

								return
							}
						}
					}
				}

				return
			}
		}).Name(id)

		route = srv.app.GetRoute(id)
	}

	srv.components.Logger.Info().
		Text("Http rest api server routes are initialized. ").Write()

	return nil
}
