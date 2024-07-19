package http_rest_api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	error_list "sm-box/internal/common/errors"
	authentication_entities "sm-box/internal/services/authentication/objects/entities"
	"sm-box/internal/services/i18n/objects/models"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/http/postman"
	"sm-box/pkg/http/rest_api/io"
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

	// /languages
	{
		var router = router.Group("/languages")

		// GET /list
		{
			var id = uuid.New().String()

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
			}).Name(id)

			var route = srv.app.GetRoute(id)

			srv.postman.AddItem(&postman.Items{
				Name: "Получение списка языков локализации. ",
				Description: `
Используется для получение языков локализации.
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
					{
						Name:   "Успешный ответ. ",
						Status: string(fiber.StatusOK),
						Code:   fiber.StatusOK,
						Body: `
{
    "code": 200,
    "code_message": "OK",
    "status": "success",
    "data": {
        "list": [
            {
                "code": "en-US",
                "name": "English",
                "active": true
            },
            {
                "code": "ru-RU",
                "name": "Русский",
                "active": true
            },
            {
                "code": "zh-CN",
                "name": "中文",
                "active": true
            }
        ]
    }
}
`,
					},
					{
						Name:   "Языки локализации не найдены. ",
						Status: string(fiber.StatusOK),
						Code:   fiber.StatusOK,
						Body: `
{
    "code": 200,
    "code_message": "OK",
    "status": "success",
    "data": {
        "list": []
    }
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
			}).Name(id)

			var route = srv.app.GetRoute(id)

			srv.postman.AddItem(&postman.Items{
				Name: "Удаление языка локализации. ",
				Description: `
Используется для удаления языка локализации.
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

		// PUT /
		{
			var id = uuid.New().String()

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
			}).Name(id)

			var route = srv.app.GetRoute(id)

			srv.postman.AddItem(&postman.Items{
				Name: "Обновления данных язык локализации. ",
				Description: `
Используется для обновления данных языка локализации.
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

		// POST /
		{
			var id = uuid.New().String()

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
			}).Name(id)

			var route = srv.app.GetRoute(id)

			srv.postman.AddItem(&postman.Items{
				Name: "Создание языка локализации. ",
				Description: `
Используется для создания языка локализации.
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

		// POST /activate
		{
			var id = uuid.New().String()

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
			}).Name(id)

			var route = srv.app.GetRoute(id)

			srv.postman.AddItem(&postman.Items{
				Name: "Активация языка локализации. ",
				Description: `
Используется для активации языка локализации.
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

		// POST /deactivate
		{
			var id = uuid.New().String()

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
			}).Name(id)

			var route = srv.app.GetRoute(id)

			srv.postman.AddItem(&postman.Items{
				Name: "Деактивация языка локализации. ",
				Description: `
Используется для деактивации языка локализации.
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

		// POST /set
		{
			var id = uuid.New().String()

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
			}).Name(id)

			var route = srv.app.GetRoute(id)

			srv.postman.AddItem(&postman.Items{
				Name: "Установить язык локализации пользователю. ",
				Description: `
Используется для установка языка локализации пользователю по токену.
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
	}

	// /texts
	{
		var router = router.Group("/texts")

		// GET /dictionary
		{
			var id = uuid.New().String()

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

						if response.Dictionary, cErr = srv.controllers.Texts.AssembleDictionary(ctx.Context(), sessionToken.Language, queryArgs.Paths); cErr != nil {
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
			}).Name(id)

			var route = srv.app.GetRoute(id)

			srv.postman.AddItem(&postman.Items{
				Name: "Получение текстов локализации на секции. ",
				Description: `
Используется для получение текстов локализации на секции, передается путь к секции родителя,
запрос возвращает текста в том числе с дочерних секций.
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
					{
						Name:   "Успешный ответ. ",
						Status: string(fiber.StatusOK),
						Code:   fiber.StatusOK,
						Body: `
{
    "code": 200,
    "code_message": "OK",
    "status": "success",
    "data": {
        "dictionary": {
            "auth": {
                "form": {
                    "buttons": {
                        "log_in": {
                            "text": "Войти"
                        }
                    },
                    "description": "Пожалуйста, укажите свои учетные данные для авторизации, чтобы продолжить. ",
                    "errors": {
                        "field_is_required": "Это поле обязательное. ",
                        "invalid_value": "Недопустимое значение. "
                    },
                    "inputs": {
                        "password": {
                            "name": "Пароль"
                        },
                        "username": {
                            "name": "Имя пользователя"
                        }
                    },
                    "title": "Добро пожаловать в SM-Box"
                }
            },
            "toasts": {
                "error": {
                    "title": "Произошла ошибка"
                }
            }
        }
    }
}
`,
					},
					{
						Name:   "Ошибка, не переданы пути текстов локализации. ",
						Status: string(fiber.StatusBadRequest),
						Code:   fiber.StatusBadRequest,
						Body: `
{
    "code": 400,
    "code_message": "Bad Request",
    "status": "failed",
    "error": {
        "id": "E-000010",
        "type": "system",
        "status": "error",
        "message": "Invalid value of text localization paths. ",
        "details": {
            "paths": "Invalid value. "
        }
    }
}
`,
					},
					{
						Name:   "Текста локализации не найдены. ",
						Status: string(fiber.StatusOK),
						Code:   fiber.StatusOK,
						Body: `
{
    "code": 200,
    "code_message": "OK",
    "status": "success",
    "data": {
        "dictionary": {}
    }
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
