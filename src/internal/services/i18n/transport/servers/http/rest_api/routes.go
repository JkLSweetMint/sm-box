package http_rest_api

import (
	"github.com/gofiber/fiber/v3"
	error_list "sm-box/internal/common/errors"
	authentication_entities "sm-box/internal/services/authentication/objects/entities"
	"sm-box/internal/services/i18n/objects/models"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/http/rest_api/io"
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

	// /languages
	{
		var router = router.Group("/languages")

		// GET /list
		{
			router.Get("/list", func(ctx fiber.Ctx) (err error) {
				type Response struct {
					List []*models.Language `json:"list" xml:"List"`
				}

				var (
					response = new(Response)
				)

				// Обработка
				{
					var cErr c_errors.RestAPI

					if response.List, cErr = srv.controllers.Languages.GetList(ctx.Context()); cErr != nil {
						srv.components.Logger.Error().
							Format("Couldn't get a list of localization languages: '%s'. ", cErr).Write()

						return http_rest_api_io.WriteError(ctx, cErr)
					}
				}

				// Отправка ответа
				{
					if err = http_rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
						srv.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
					}
					return
				}
			})
		}

		// DELETE /
		{
			router.Delete("/", func(ctx fiber.Ctx) (err error) {
				type Request struct {
					Code string `json:"code"`
				}
				type Response struct{}

				var (
					request  = new(Request)
					response = new(Response)
				)

				// Чтение данных
				{
					if err = ctx.Bind().Body(request); err != nil {
						srv.components.Logger.Error().
							Format("The request body data could not be read: '%s'. ", err).Write()

						if err = http_rest_api_io.WriteError(ctx, error_list.RequestBodyDataCouldNotBeRead_RestAPI()); err != nil {
							srv.components.Logger.Error().
								Format("The response could not be recorded: '%s'. ", err).Write()

							var cErr = error_list.ResponseCouldNotBeRecorded_RestAPI()
							cErr.SetError(err)

							return http_rest_api_io.WriteError(ctx, cErr)
						}

						return
					}
				}

				// Обработка
				{
					var cErr c_errors.RestAPI

					if cErr = srv.controllers.Languages.Remove(ctx.Context(), request.Code); cErr != nil {
						srv.components.Logger.Error().
							Format("The localization language could not be deleted: '%s'. ", cErr).Write()

						return http_rest_api_io.WriteError(ctx, cErr)
					}
				}

				// Отправка ответа
				{
					if err = http_rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
						srv.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
					}
					return
				}
			})
		}

		// PUT /
		{
			router.Put("/", func(ctx fiber.Ctx) (err error) {
				type Request struct {
					Code string `json:"code"`
					Name string `json:"name"`
				}
				type Response struct{}

				var (
					request  = new(Request)
					response = new(Response)
				)

				// Чтение данных
				{
					if err = ctx.Bind().Body(request); err != nil {
						srv.components.Logger.Error().
							Format("The request body data could not be read: '%s'. ", err).Write()

						if err = http_rest_api_io.WriteError(ctx, error_list.RequestBodyDataCouldNotBeRead_RestAPI()); err != nil {
							srv.components.Logger.Error().
								Format("The response could not be recorded: '%s'. ", err).Write()

							var cErr = error_list.ResponseCouldNotBeRecorded_RestAPI()
							cErr.SetError(err)

							return http_rest_api_io.WriteError(ctx, cErr)
						}

						return
					}
				}

				// Обработка
				{
					var cErr c_errors.RestAPI

					if cErr = srv.controllers.Languages.Update(ctx.Context(), request.Code, request.Name); cErr != nil {
						srv.components.Logger.Error().
							Format("The localization language data could not be updated: '%s'. ", cErr).Write()

						return http_rest_api_io.WriteError(ctx, cErr)
					}
				}

				// Отправка ответа
				{
					if err = http_rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
						srv.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
					}
					return
				}
			})
		}

		// POST /
		{
			router.Put("/", func(ctx fiber.Ctx) (err error) {
				type Request struct {
					Code string `json:"code"`
					Name string `json:"name"`
				}
				type Response struct{}

				var (
					request  = new(Request)
					response = new(Response)
				)

				// Чтение данных
				{
					if err = ctx.Bind().Body(request); err != nil {
						srv.components.Logger.Error().
							Format("The request body data could not be read: '%s'. ", err).Write()

						if err = http_rest_api_io.WriteError(ctx, error_list.RequestBodyDataCouldNotBeRead_RestAPI()); err != nil {
							srv.components.Logger.Error().
								Format("The response could not be recorded: '%s'. ", err).Write()

							var cErr = error_list.ResponseCouldNotBeRecorded_RestAPI()
							cErr.SetError(err)

							return http_rest_api_io.WriteError(ctx, cErr)
						}

						return
					}
				}

				// Обработка
				{
					var cErr c_errors.RestAPI

					if cErr = srv.controllers.Languages.Create(ctx.Context(), request.Code, request.Name); cErr != nil {
						srv.components.Logger.Error().
							Format("The localization language could not be created: '%s'. ", cErr).Write()

						return http_rest_api_io.WriteError(ctx, cErr)
					}
				}

				// Отправка ответа
				{
					if err = http_rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
						srv.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
					}
					return
				}
			})
		}

		// POST /activate
		{
			router.Post("/activate", func(ctx fiber.Ctx) (err error) {
				type Request struct {
					Code string `json:"code"`
				}
				type Response struct{}

				var (
					request  = new(Request)
					response = new(Response)
				)

				// Чтение данных
				{
					if err = ctx.Bind().Body(request); err != nil {
						srv.components.Logger.Error().
							Format("The request body data could not be read: '%s'. ", err).Write()

						if err = http_rest_api_io.WriteError(ctx, error_list.RequestBodyDataCouldNotBeRead_RestAPI()); err != nil {
							srv.components.Logger.Error().
								Format("The response could not be recorded: '%s'. ", err).Write()

							var cErr = error_list.ResponseCouldNotBeRecorded_RestAPI()
							cErr.SetError(err)

							return http_rest_api_io.WriteError(ctx, cErr)
						}

						return
					}
				}

				// Обработка
				{
					var cErr c_errors.RestAPI

					if cErr = srv.controllers.Languages.Activate(ctx.Context(), request.Code); cErr != nil {
						srv.components.Logger.Error().
							Format("The localization language could not be activated: '%s'. ", cErr).Write()

						return http_rest_api_io.WriteError(ctx, cErr)
					}
				}

				// Отправка ответа
				{
					if err = http_rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
						srv.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
					}
					return
				}
			})
		}

		// POST /deactivate
		{
			router.Post("/deactivate", func(ctx fiber.Ctx) (err error) {
				type Request struct {
					Code string `json:"code"`
				}
				type Response struct{}

				var (
					request  = new(Request)
					response = new(Response)
				)

				// Чтение данных
				{
					if err = ctx.Bind().Body(request); err != nil {
						srv.components.Logger.Error().
							Format("The request body data could not be read: '%s'. ", err).Write()

						if err = http_rest_api_io.WriteError(ctx, error_list.RequestBodyDataCouldNotBeRead_RestAPI()); err != nil {
							srv.components.Logger.Error().
								Format("The response could not be recorded: '%s'. ", err).Write()

							var cErr = error_list.ResponseCouldNotBeRecorded_RestAPI()
							cErr.SetError(err)

							return http_rest_api_io.WriteError(ctx, cErr)
						}

						return
					}
				}

				// Обработка
				{
					var cErr c_errors.RestAPI

					if cErr = srv.controllers.Languages.Deactivate(ctx.Context(), request.Code); cErr != nil {
						srv.components.Logger.Error().
							Format("The localization language could not be deactivated: '%s'. ", cErr).Write()

						return http_rest_api_io.WriteError(ctx, cErr)
					}
				}

				// Отправка ответа
				{
					if err = http_rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
						srv.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
					}
					return
				}
			})
		}

		// POST /set
		{
			router.Post("/set", func(ctx fiber.Ctx) (err error) {
				type Request struct {
					Code string `json:"code"`
				}
				type Response struct{}

				var (
					request  = new(Request)
					response = new(Response)
				)

				// Чтение данных
				{
					if err = ctx.Bind().Body(request); err != nil {
						srv.components.Logger.Error().
							Format("The request body data could not be read: '%s'. ", err).Write()

						if err = http_rest_api_io.WriteError(ctx, error_list.RequestBodyDataCouldNotBeRead_RestAPI()); err != nil {
							srv.components.Logger.Error().
								Format("The response could not be recorded: '%s'. ", err).Write()

							var cErr = error_list.ResponseCouldNotBeRecorded_RestAPI()
							cErr.SetError(err)

							return http_rest_api_io.WriteError(ctx, cErr)
						}

						return
					}
				}

				// Обработка
				{

				}

				// Отправка ответа
				{
					if err = http_rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
						srv.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
					}
					return
				}
			})
		}
	}

	// /texts
	{
		var router = router.Group("/texts")

		// GET /dictionary
		{
			router.Get("/dictionary", func(ctx fiber.Ctx) (err error) {
				type Response struct {
					Dictionary models.Dictionary `json:"dictionary" xml:"Dictionary"`
				}
				type QueryArgs struct {
					Paths []string `uri:"paths"`
				}

				var (
					response  = new(Response)
					queryArgs = new(QueryArgs)
				)

				// Чтение данных
				{
					if err = ctx.Bind().Query(queryArgs); err != nil {
						srv.components.Logger.Error().
							Format("The request body data could not be read: '%s'. ", err).Write()

						if err = http_rest_api_io.WriteError(ctx, error_list.RequestBodyDataCouldNotBeRead_RestAPI()); err != nil {
							srv.components.Logger.Error().
								Format("The response could not be recorded: '%s'. ", err).Write()

							var cErr = error_list.ResponseCouldNotBeRecorded_RestAPI()
							cErr.SetError(err)

							return http_rest_api_io.WriteError(ctx, cErr)
						}

						return
					}
				}

				// Обработка
				{
					var (
						rawSessionToken = ctx.Cookies(srv.conf.Components.AccessSystem.CookieKeyForSessionToken)
						sessionToken    = new(authentication_entities.JwtSessionToken)
					)

					// Получение токена сессии
					{
						if err = sessionToken.Parse(rawSessionToken); err != nil {
							srv.components.Logger.Error().
								Format("Failed to get session token data: '%s'. ", err).
								Field("raw", rawSessionToken).Write()

							if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(error_list.InvalidToken())); err != nil {
								srv.components.Logger.Error().
									Format("The response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
							}
							return
						}
					}

					// Получение текстов
					{
						var cErr c_errors.RestAPI

						if response.Dictionary, cErr = srv.controllers.Texts.AssembleDictionary(ctx.Context(), sessionToken.Claims.Language, queryArgs.Paths); cErr != nil {
							srv.components.Logger.Error().
								Format("The localization dictionary could not be assembled: '%s'. ", cErr).Write()

							return http_rest_api_io.WriteError(ctx, cErr)
						}
					}
				}

				// Отправка ответа
				{
					if err = http_rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
						srv.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
					}
					return
				}
			})
		}
	}

	srv.components.Logger.Info().
		Text("Http rest api server routes are initialized. ").Write()

	return nil
}
