package http_rest_api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	common_errors "sm-box/internal/common/errors"
	common_types "sm-box/internal/common/types"
	authentication_entities "sm-box/internal/services/authentication/objects/entities"
	"sm-box/internal/services/notifications/objects"
	"sm-box/internal/services/notifications/objects/models"
	"sm-box/internal/services/notifications/objects/types"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/http/postman"
	http_rest_api_io "sm-box/pkg/http/rest_api/io"
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

	{
		var (
			router       = router.Group("/users")
			postmanGroup = srv.postman.AddItemGroup("Пользовательские уведомления. ")
		)

		// /ws
		{

		}

		// GET /list
		{
			var id = uuid.New().String()

			router.Get("/list", func(ctx fiber.Ctx) (err error) {
				type Response struct {
					Count int64                          `json:"count" xml:"count,attr"`
					List  []*models.UserNotificationInfo `json:"list"  xml:"List"`
				}
				type QueryArgs struct {
					Search string `query:"search"`

					Limit  *int64 `query:"limit"`
					Offset *int64 `query:"offset"`

					FilterType     *string          `query:"filter_type"`
					FilterNotRead  *string          `query:"filter_not_read"`
					FilterSenderID *common_types.ID `query:"filter_sender_id"`
				}

				var (
					response  = new(Response)
					queryArgs = new(QueryArgs)
				)

				// Чтение данных
				{
					if err = ctx.Bind().Query(queryArgs); err != nil {
						srv.components.Logger.Error().
							Format("The request query data could not be read: '%s'. ", err).Write()

						if err = http_rest_api_io.WriteError(ctx, common_errors.RequestBodyDataCouldNotBeRead_RestAPI()); err != nil {
							srv.components.Logger.Error().
								Format("The response could not be recorded: '%s'. ", err).Write()

							var cErr = common_errors.ResponseCouldNotBeRecorded_RestAPI()
							cErr.SetError(err)

							if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
								srv.components.Logger.Error().
									Format("The error response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
							}

							return
						}

						return
					}
				}

				// Обработка
				{
					var (
						search     *objects.UserNotificationSearch
						pagination *objects.UserNotificationPagination
						filters    *objects.UserNotificationFilters
						userID     common_types.ID
					)

					// Получение ID пользователя
					{
						// Получение токена
						{
							var token *authentication_entities.JwtSessionToken

							if raw := ctx.Cookies(srv.conf.Components.AccessSystem.CookieKeyForSessionToken); len(raw) > 0 {
								token = new(authentication_entities.JwtSessionToken)

								if err = token.Parse(raw); err != nil {
									srv.components.Logger.Error().
										Format("Failed to get session token data: '%s'. ", err).
										Field("raw", raw).Write()

									err = ctx.Redirect().To("/errors/403")
									return
								}
							}

							userID = token.UserID
						}
					}

					// Обработка входных данных
					{
						// Поиск
						{
							queryArgs.Search = strings.TrimSpace(queryArgs.Search)

							search = &objects.UserNotificationSearch{
								Global: queryArgs.Search,
							}
						}

						// Пагинация
						{
							pagination = &objects.UserNotificationPagination{
								Offset: queryArgs.Offset,
								Limit:  queryArgs.Limit,
							}
						}

						// Фильтрация
						{
							filters = new(objects.UserNotificationFilters)

							if queryArgs.FilterNotRead != nil {
								switch strings.ToLower(strings.TrimSpace(*queryArgs.FilterNotRead)) {
								case "true", "on", "enable":
									{
										var v = true
										filters.NotRead = &v
									}
								case "false", "off", "disable":
									{
										var v = false
										filters.NotRead = &v
									}
								}
							}

							if queryArgs.FilterType != nil && *queryArgs.FilterType != "" {
								var v = types.NotificationType(strings.TrimSpace(*queryArgs.FilterType))
								filters.Type = &v
							}

							if queryArgs.FilterSenderID != nil {
								filters.SenderID = queryArgs.FilterSenderID
							}
						}
					}

					var cErr c_errors.RestAPI

					if response.Count, response.List, cErr = srv.controllers.UserNotifications.GetList(ctx.Context(), userID, search, pagination, filters); cErr != nil {
						srv.components.Logger.Error().
							Format("Couldn't get a list of user notifications: '%s'. ", cErr).Write()

						if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
							srv.components.Logger.Error().
								Format("The error response could not be recorded: '%s'. ", err).Write()

							return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
						}

						return
					}
				}

				// Отправка ответа
				{
					if err = http_rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
						srv.components.Logger.Error().
							Format("The error response could not be recorded: '%s'. ", err).Write()

						return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
					}

					return
				}
			}).Name(id)

			var route = srv.app.GetRoute(id)

			postmanGroup.AddItem(&postman.Items{
				Name: "Получение списка пользовательских уведомлений. ",
				Description: `
Используется для получения списка пользовательских уведомлений.
`,
				Request: &postman.Request{
					URL: &postman.URL{
						Protocol: srv.conf.Postman.Protocol,
						Host:     strings.Split(srv.conf.Postman.Host, "."),
						Path:     strings.Split(route.Path, "/"),
						Query: []*postman.QueryParam{
							{
								Key:   "search",
								Value: "",
							},
							{
								Key:   "offset",
								Value: "0",
							},
							{
								Key:   "limit",
								Value: "20",
							},
							{
								Key:   "filter_type",
								Value: "",
							},
							{
								Key:   "filter_not_read",
								Value: "",
							},
							{
								Key:   "filter_sender_id",
								Value: "",
							},
						},
					},
					Method:      postman.Method(route.Method),
					Description: ``,
					Body: &postman.Body{
						Mode: "raw",
						Raw:  ``,
						Options: &postman.BodyOptions{
							Raw: postman.BodyOptionsRaw{
								Language: postman.JSON,
							},
						},
					},
				},
				Responses: []*postman.Response{
					{
						Name:   "Произошла внутренняя ошибка сервера. ",
						Status: string(fiber.StatusInternalServerError),
						Code:   fiber.StatusInternalServerError,
						Body: `
{
    "code": 500,
    "code_message": "Internal Server Error",
    "status": "error",
    "error": {
        "id": "I-000001",
        "type": "system",
        "status": "error",
        "message": "Internal server error. ",
        "details": {}
    }
}
`,
					},
				},
			})
		}

		// PUT /read
		{
			var id = uuid.New().String()

			router.Put("/read", func(ctx fiber.Ctx) (err error) {
				type Request struct {
					ID  common_types.ID   `json:"id"`
					IDs []common_types.ID `json:"ids"`
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

						if err = http_rest_api_io.WriteError(ctx, common_errors.RequestBodyDataCouldNotBeRead_RestAPI()); err != nil {
							srv.components.Logger.Error().
								Format("The response could not be recorded: '%s'. ", err).Write()

							var cErr = common_errors.ResponseCouldNotBeRecorded_RestAPI()
							cErr.SetError(err)

							if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
								srv.components.Logger.Error().
									Format("The error response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
							}

							return
						}

						return
					}
				}

				// Обработка
				{
					var userID common_types.ID

					// Получение ID пользователя
					{
						// Получение токена
						{
							var token *authentication_entities.JwtSessionToken

							if raw := ctx.Cookies(srv.conf.Components.AccessSystem.CookieKeyForSessionToken); len(raw) > 0 {
								token = new(authentication_entities.JwtSessionToken)

								if err = token.Parse(raw); err != nil {
									srv.components.Logger.Error().
										Format("Failed to get session token data: '%s'. ", err).
										Field("raw", raw).Write()

									err = ctx.Redirect().To("/errors/403")
									return
								}
							}

							userID = token.UserID
						}
					}

					var cErr c_errors.RestAPI

					if len(request.IDs) > 0 {
						if cErr = srv.controllers.UserNotifications.Read(ctx.Context(), userID, request.IDs...); cErr != nil {
							srv.components.Logger.Error().
								Format("Several user notifications could not be read: '%s'. ", cErr).Write()

							if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
								srv.components.Logger.Error().
									Format("The error response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
							}

							return
						}
					} else {
						if cErr = srv.controllers.UserNotifications.ReadOne(ctx.Context(), userID, request.ID); cErr != nil {
							srv.components.Logger.Error().
								Format("The user's notification could not be read: '%s'. ", cErr).Write()

							if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
								srv.components.Logger.Error().
									Format("The error response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
							}

							return
						}
					}
				}

				// Отправка ответа
				{
					if err = http_rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
						srv.components.Logger.Error().
							Format("The error response could not be recorded: '%s'. ", err).Write()

						return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
					}

					return
				}
			}).Name(id)

			var route = srv.app.GetRoute(id)

			postmanGroup.AddItem(&postman.Items{
				Name: "Чтение пользовательских уведомлений. ",
				Description: `
Используется для чтения пользовательских уведомлений.
`,
				Request: &postman.Request{
					URL: &postman.URL{
						Protocol: srv.conf.Postman.Protocol,
						Host:     strings.Split(srv.conf.Postman.Host, "."),
						Path:     strings.Split(route.Path, "/"),
					},
					Method:      postman.Method(route.Method),
					Description: ``,
					Body: &postman.Body{
						Mode: "raw",
						Raw: `
{
	"source": "",
	"accesses": {
		"roles_id": [],
		"permissions_id": []
	},
	"properties": {
		"type": "",
		"number_of_uses": 0,
		"active": false,
		"start_active": "0001-01-01T00:00:00+00:00",
		"end_active": "0001-01-01T00:00:00+00:00"
	}
}
`,
						Options: &postman.BodyOptions{
							Raw: postman.BodyOptionsRaw{
								Language: postman.JSON,
							},
						},
					},
				},
				Responses: []*postman.Response{
					{
						Name:   "Произошла внутренняя ошибка сервера. ",
						Status: string(fiber.StatusInternalServerError),
						Code:   fiber.StatusInternalServerError,
						Body: `
{
    "code": 500,
    "code_message": "Internal Server Error",
    "status": "error",
    "error": {
        "id": "I-000001",
        "type": "system",
        "status": "error",
        "message": "Internal server error. ",
        "details": {}
    }
}
`,
					},
					{
						Name:   "Успешный ответ. ",
						Status: string(fiber.StatusOK),
						Code:   fiber.StatusOK,
						Body: `
{
    "code": 200,
    "code_message": "OK",
    "status": "success",
    "data": {}
}
`,
					},
				},
			})
		}

		// DELETE /
		{
			var id = uuid.New().String()

			router.Delete("/", func(ctx fiber.Ctx) (err error) {
				type Request struct {
					ID  common_types.ID   `json:"id"`
					IDs []common_types.ID `json:"ids"`
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

						if err = http_rest_api_io.WriteError(ctx, common_errors.RequestBodyDataCouldNotBeRead_RestAPI()); err != nil {
							srv.components.Logger.Error().
								Format("The response could not be recorded: '%s'. ", err).Write()

							var cErr = common_errors.ResponseCouldNotBeRecorded_RestAPI()
							cErr.SetError(err)

							if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
								srv.components.Logger.Error().
									Format("The error response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
							}

							return
						}

						return
					}
				}

				// Обработка
				{
					var userID common_types.ID

					// Получение ID пользователя
					{
						// Получение токена
						{
							var token *authentication_entities.JwtSessionToken

							if raw := ctx.Cookies(srv.conf.Components.AccessSystem.CookieKeyForSessionToken); len(raw) > 0 {
								token = new(authentication_entities.JwtSessionToken)

								if err = token.Parse(raw); err != nil {
									srv.components.Logger.Error().
										Format("Failed to get session token data: '%s'. ", err).
										Field("raw", raw).Write()

									err = ctx.Redirect().To("/errors/403")
									return
								}
							}

							userID = token.UserID
						}
					}

					var cErr c_errors.RestAPI

					if len(request.IDs) > 0 {
						if cErr = srv.controllers.UserNotifications.Remove(ctx.Context(), userID, request.IDs...); cErr != nil {
							srv.components.Logger.Error().
								Format("Several user notifications could not be removed: '%s'. ", cErr).Write()

							if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
								srv.components.Logger.Error().
									Format("The error response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
							}

							return
						}
					} else {
						if cErr = srv.controllers.UserNotifications.RemoveOne(ctx.Context(), userID, request.ID); cErr != nil {
							srv.components.Logger.Error().
								Format("The user's notification could not be removed: '%s'. ", cErr).Write()

							if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
								srv.components.Logger.Error().
									Format("The error response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
							}

							return
						}
					}
				}

				// Отправка ответа
				{
					if err = http_rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
						srv.components.Logger.Error().
							Format("The error response could not be recorded: '%s'. ", err).Write()

						return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
					}

					return
				}
			}).Name(id)

			var route = srv.app.GetRoute(id)

			postmanGroup.AddItem(&postman.Items{
				Name: "Удаление пользовательских уведомлений. ",
				Description: `
Используется для удаления пользовательских уведомлений.
`,
				Request: &postman.Request{
					URL: &postman.URL{
						Protocol: srv.conf.Postman.Protocol,
						Host:     strings.Split(srv.conf.Postman.Host, "."),
						Path:     strings.Split(route.Path, "/"),
					},
					Method:      postman.Method(route.Method),
					Description: ``,
					Body: &postman.Body{
						Mode: "raw",
						Raw: `
{
	"source": "",
	"accesses": {
		"roles_id": [],
		"permissions_id": []
	},
	"properties": {
		"type": "",
		"number_of_uses": 0,
		"active": false,
		"start_active": "0001-01-01T00:00:00+00:00",
		"end_active": "0001-01-01T00:00:00+00:00"
	}
}
`,
						Options: &postman.BodyOptions{
							Raw: postman.BodyOptionsRaw{
								Language: postman.JSON,
							},
						},
					},
				},
				Responses: []*postman.Response{
					{
						Name:   "Произошла внутренняя ошибка сервера. ",
						Status: string(fiber.StatusInternalServerError),
						Code:   fiber.StatusInternalServerError,
						Body: `
{
    "code": 500,
    "code_message": "Internal Server Error",
    "status": "error",
    "error": {
        "id": "I-000001",
        "type": "system",
        "status": "error",
        "message": "Internal server error. ",
        "details": {}
    }
}
`,
					},
					{
						Name:   "Успешный ответ. ",
						Status: string(fiber.StatusOK),
						Code:   fiber.StatusOK,
						Body: `
{
    "code": 200,
    "code_message": "OK",
    "status": "success",
    "data": {}
}
`,
					},
				},
			})
		}
	}

	srv.components.Logger.Info().
		Text("Http rest api server routes are initialized. ").Write()

	return nil
}
