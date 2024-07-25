package http_rest_api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	app_models "sm-box/internal/app/objects/models"
	error_list "sm-box/internal/common/errors"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/authentication/objects/models"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/http/postman"
	http_rest_api_io "sm-box/pkg/http/rest_api/io"
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

	// /basic-auth
	{
		var router = router.Group("/basic-auth")

		// /projects
		{
			var router = router.Group("/projects")

			// GET /select
			{
				var id = uuid.New().String()

				router.Get("/select", func(ctx fiber.Ctx) (err error) {
					type Response struct {
						Projects app_models.ProjectList `json:"projects" xml:"Projects"`
					}

					var response = new(Response)

					// Обработка
					{
						var rawSessionToken = ctx.Cookies(srv.conf.Components.AccessSystem.CookieKeyForSessionToken)

						// Получение списка
						{
							var cErr c_errors.RestAPI

							if response.Projects, cErr = srv.controllers.BasicAuthentication.GetUserProjectList(ctx.Context(), rawSessionToken); cErr != nil {
								srv.components.Logger.Error().
									Format("The list of user's projects could not be retrieved: '%s'. ", cErr).Write()

								if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
									srv.components.Logger.Error().
										Format("The error response could not be recorded: '%s'. ", err).Write()

									return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
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

							return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
						}

						return
					}
				}).Name(id)

				var route = srv.app.GetRoute(id)

				srv.postman.AddItem(&postman.Items{
					Name: "Получение списка проектов пользователя для выбора после авторизации. ",
					Request: &postman.Request{
						URL: &postman.URL{
							Protocol: srv.conf.Postman.Protocol,
							Host:     strings.Split(srv.conf.Postman.Host, "."),
							Path:     strings.Split(route.Path, "/"),
						},
						Method: postman.Method(route.Method),
						Description: `
Используется для получение списка проектов пользователя для выбора после авторизации. 
После выбора проекта пользователю будут доступны методы управления проектом.
`,
						Body: &postman.Body{
							Mode: "raw",
							Raw: `
{
   "username": "",
   "password": ""
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
							Name:   "Пример успешного ответа. ",
							Status: string(fiber.StatusOK),
							Code:   fiber.StatusOK,
							Body: `
{
    "code": 200,
    "code_message": "OK",
    "status": "success",
    "data": {
        "projects": [
            {
                "id": 1,
                "name": "System",
                "version": ""
            },
            {
                "id": 2,
                "name": "Test",
                "version": ""
            },
            {
                "id": 3,
                "name": "Test 2",
                "version": "v2.0"
            }
        ]
    }
}
`,
						},
						{
							Name:   "Пример успешного ответа, но с отсутствием проектов. ",
							Status: string(fiber.StatusOK),
							Code:   fiber.StatusOK,
							Body: `
{
    "code": 200,
    "code_message": "OK",
    "status": "success",
    "data": {
        "projects": []
    }
}
`,
						},
						{
							Name:   "Ошибка если проект уже выбран. ",
							Status: string(fiber.StatusBadRequest),
							Code:   fiber.StatusBadRequest,
							Body: `
{
    "code": 400,
    "code_message": "Bad Request",
    "status": "failed",
    "error": {
        "id": "E-000007",
        "type": "system",
        "status": "error",
        "message": "The project has already been selected, it is not possible to re-select it. ",
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
						ProjectID common_types.ID `json:"project_id"`
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

								if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
									srv.components.Logger.Error().
										Format("The error response could not be recorded: '%s'. ", err).Write()

									return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
								}

								return
							}

							return
						}
					}

					// Обработка
					{
						var (
							rawSessionToken                         = ctx.Cookies(srv.conf.Components.AccessSystem.CookieKeyForSessionToken)
							cErr                                    c_errors.RestAPI
							sessionToken, accessToken, refreshToken *models.JwtTokenInfo
						)

						if sessionToken, accessToken, refreshToken, cErr = srv.controllers.BasicAuthentication.SetTokenProject(ctx.Context(), rawSessionToken, request.ProjectID); cErr != nil {
							srv.components.Logger.Error().
								Format("The project value for the user token could not be set: '%s'. ", cErr).Write()

							if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
								srv.components.Logger.Error().
									Format("The error response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
							}

							return
						}

						// Запись печенек
						{
							if sessionToken != nil {
								var cookie = &fiber.Cookie{
									Name:        srv.conf.Components.AccessSystem.CookieKeyForSessionToken,
									Value:       sessionToken.Raw,
									Path:        "/",
									Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
									MaxAge:      0,
									Expires:     sessionToken.ExpiresAt,
									Secure:      false,
									HTTPOnly:    true,
									SameSite:    fiber.CookieSameSiteLaxMode,
									SessionOnly: false,
								}

								ctx.Cookie(cookie)
							}

							if accessToken != nil {
								var cookie = &fiber.Cookie{
									Name:        srv.conf.Components.AccessSystem.CookieKeyForAccessToken,
									Value:       accessToken.Raw,
									Path:        "/",
									Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
									MaxAge:      0,
									Expires:     accessToken.ExpiresAt,
									Secure:      true,
									HTTPOnly:    true,
									SameSite:    fiber.CookieSameSiteLaxMode,
									SessionOnly: false,
								}

								ctx.Cookie(cookie)
							}

							if refreshToken != nil {
								var cookie = &fiber.Cookie{
									Name:        srv.conf.Components.AccessSystem.CookieKeyForRefreshToken,
									Value:       refreshToken.Raw,
									Path:        "/",
									Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
									MaxAge:      0,
									Expires:     refreshToken.ExpiresAt,
									Secure:      true,
									HTTPOnly:    true,
									SameSite:    fiber.CookieSameSiteLaxMode,
									SessionOnly: false,
								}

								ctx.Cookie(cookie)
							}
						}
					}

					// Отправка ответа
					{
						if err = http_rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
							srv.components.Logger.Error().
								Format("The error response could not be recorded: '%s'. ", err).Write()

							return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
						}

						return
					}
				}).Name(id)

				var route = srv.app.GetRoute(id)

				srv.postman.AddItem(&postman.Items{
					Name: "Выбрать проект пользователя для дальнейше работы после авторизации. ",
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
   "project_id": 0
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
							Name:   "Пример успешного ответа. ",
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
						{
							Name:   "Ошибка если проект уже выбран. ",
							Status: string(fiber.StatusBadRequest),
							Code:   fiber.StatusBadRequest,
							Body: `
{
    "code": 400,
    "code_message": "Bad Request",
    "status": "failed",
    "error": {
        "id": "E-000007",
        "type": "system",
        "status": "error",
        "message": "The project has already been selected, it is not possible to re-select it. ",
        "details": {}
    }
}
`,
						},
						{
							Name:   "Ошибка если проект не найден. ",
							Status: string(fiber.StatusNotFound),
							Code:   fiber.StatusNotFound,
							Body: `
{
    "code": 404,
    "code_message": "Not Found",
    "status": "failed",
    "error": {
        "id": "E-000008",
        "type": "system",
        "status": "error",
        "message": "The project was not found. ",
        "details": {}
    }
}
`,
						},
						{
							Name:   "Ошибка если нет доступа к проекту. ",
							Status: string(fiber.StatusBadRequest),
							Code:   fiber.StatusBadRequest,
							Body: `
{
    "code": 400,
    "code_message": "Bad Request",
    "status": "failed",
    "error": {
        "id": "E-000009",
        "type": "system",
        "status": "error",
        "message": "There is no access to the project. ",
        "details": {}
    }
}
`,
						},
					},
				})
			}
		}

		// POST /
		{
			var id = uuid.New().String()

			router.Post("/", func(ctx fiber.Ctx) (err error) {
				type Request struct {
					Username string `json:"username"`
					Password string `json:"password"`
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

							if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
								srv.components.Logger.Error().
									Format("The error response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
							}

							return
						}

						return
					}
				}

				// Обработка
				{
					var (
						rawSessionToken                         string
						cErr                                    c_errors.RestAPI
						sessionToken, accessToken, refreshToken *models.JwtTokenInfo
					)

					// Получение токена сессии
					{
						rawSessionToken = ctx.Cookies(srv.conf.Components.AccessSystem.CookieKeyForSessionToken)
					}

					// Авторизация
					{
						if sessionToken, cErr = srv.controllers.BasicAuthentication.Auth(ctx.Context(),
							rawSessionToken,
							request.Username,
							request.Password); cErr != nil {
							srv.components.Logger.Error().
								Format("User authorization failed: '%s'. ", cErr).Write()

							if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
								srv.components.Logger.Error().
									Format("The error response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
							}

							return
						}
					}

					// Установка проекта если он у пользователя один
					{
						var projects app_models.ProjectList

						if projects, cErr = srv.controllers.BasicAuthentication.GetUserProjectList(ctx.Context(), sessionToken.Raw); cErr != nil {
							srv.components.Logger.Error().
								Format("The list of user's projects could not be retrieved: '%s'. ", cErr).Write()

							cErr = nil
						}

						if len(projects) == 1 {
							var project = projects[0]

							if project != nil {
								var newSessionToken *models.JwtTokenInfo

								if newSessionToken, accessToken, refreshToken, cErr = srv.controllers.BasicAuthentication.SetTokenProject(ctx.Context(), sessionToken.Raw, project.ID); cErr != nil {
									srv.components.Logger.Error().
										Format("The project value for the user token could not be set: '%s'. ", cErr).Write()

									cErr = nil
								} else {
									sessionToken = newSessionToken
									ctx.Response().Header.Set("X-Authorization-State", "done")

									srv.components.Logger.Info().
										Format("The user was automatically logged into the project. ").
										Field("project", project).Write()
								}
							}
						}
					}

					// Запись печенек
					{
						if sessionToken != nil {
							var cookie = &fiber.Cookie{
								Name:        srv.conf.Components.AccessSystem.CookieKeyForSessionToken,
								Value:       sessionToken.Raw,
								Path:        "/",
								Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
								MaxAge:      0,
								Expires:     sessionToken.ExpiresAt,
								Secure:      false,
								HTTPOnly:    true,
								SameSite:    fiber.CookieSameSiteLaxMode,
								SessionOnly: false,
							}

							ctx.Cookie(cookie)
						}

						if accessToken != nil {
							var cookie = &fiber.Cookie{
								Name:        srv.conf.Components.AccessSystem.CookieKeyForAccessToken,
								Value:       accessToken.Raw,
								Path:        "/",
								Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
								MaxAge:      0,
								Expires:     accessToken.ExpiresAt,
								Secure:      true,
								HTTPOnly:    true,
								SameSite:    fiber.CookieSameSiteLaxMode,
								SessionOnly: false,
							}

							ctx.Cookie(cookie)
						}

						if refreshToken != nil {
							var cookie = &fiber.Cookie{
								Name:        srv.conf.Components.AccessSystem.CookieKeyForRefreshToken,
								Value:       refreshToken.Raw,
								Path:        "/",
								Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
								MaxAge:      0,
								Expires:     refreshToken.ExpiresAt,
								Secure:      true,
								HTTPOnly:    true,
								SameSite:    fiber.CookieSameSiteLaxMode,
								SessionOnly: false,
							}

							ctx.Cookie(cookie)
						}
					}
				}

				// Отправка ответа
				{
					if err = http_rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
						srv.components.Logger.Error().
							Format("The error response could not be recorded: '%s'. ", err).Write()

						return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
					}

					return
				}
			}).Name(id)

			var route = srv.app.GetRoute(id)

			srv.postman.AddItem(&postman.Items{
				Name: "Запрос для базовой авторизации пользователя. ",
				Request: &postman.Request{
					URL: &postman.URL{
						Protocol: srv.conf.Postman.Protocol,
						Host:     strings.Split(srv.conf.Postman.Host, "."),
						Path:     strings.Split(route.Path, "/"),
					},
					Method: postman.Method(route.Method),
					Description: `Используется для аутентификации пользователя по логину и паролю. Этот запрос принимает два параметра: username 
(имя пользователя) и password (пароль). Сервер проверяет предоставленные учетные данные и, если они верны, авторизирует
токен хранящийся в Cookie, который затем будет использован для выполнения других операций в системе.`,
					Body: &postman.Body{
						Mode: "raw",
						Raw: `
{
   "username": "",
   "password": ""
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
						Name:   "Пример успешного ответа. ",
						Status: string(fiber.StatusOK),
						Code:   fiber.StatusOK,
						Body: `
{
    "code": 200,
    "code_message": "OK",
    "status": "success"
}
`,
					},
					{
						Name:   "Пример ответа если пользователь не найден. ",
						Status: string(fiber.StatusNotFound),
						Code:   fiber.StatusNotFound,
						Body: `
{
    "code": 404,
    "code_message": "Not Found",
    "status": "failed",
    "error": {
        "id": "E-000004",
        "type": "system",
        "status": "error",
        "message": "The user was not found. ",
        "details": {}
    }
}
`,
					},
					{
						Name:   "Пример ответа при не корректном запросе. ",
						Status: string(fiber.StatusBadRequest),
						Code:   fiber.StatusBadRequest,
						Body: `
{
    "code": 400,
    "code_message": "Bad Request",
    "status": "failed",
    "error": {
        "id": "ERA-000002",
        "type": "system",
        "status": "error",
        "message": "The request body data could not be read. ",
        "details": {}
    }
}
`,
					},
				},
			})
		}

		// POST /logout
		{
			var id = uuid.New().String()

			router.Get("/logout", func(ctx fiber.Ctx) (err error) {
				type Response struct{}

				var response = new(Response)

				// Обработка
				{
					var rawSessionToken, rawAccessToken, rawRefreshToken string

					if cErr := srv.controllers.BasicAuthentication.Logout(ctx.Context(), rawSessionToken, rawAccessToken, rawRefreshToken); cErr != nil {
						srv.components.Logger.Error().
							Format("The user's session could not be terminated: '%s'. ", cErr).Write()

						if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
							srv.components.Logger.Error().
								Format("The error response could not be recorded: '%s'. ", err).Write()

							return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
						}

						return
					}

					// Очистка куков
					{
						ctx.Cookie(&fiber.Cookie{
							Name:        srv.conf.Components.AccessSystem.CookieKeyForSessionToken,
							Value:       "",
							Path:        "/",
							Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
							MaxAge:      0,
							Expires:     time.Unix(0, 0),
							Secure:      false,
							HTTPOnly:    false,
							SameSite:    fiber.CookieSameSiteNoneMode,
							SessionOnly: false,
						})

						ctx.Cookie(&fiber.Cookie{
							Name:        srv.conf.Components.AccessSystem.CookieKeyForAccessToken,
							Value:       "",
							Path:        "/",
							Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
							MaxAge:      0,
							Expires:     time.Unix(0, 0),
							Secure:      false,
							HTTPOnly:    false,
							SameSite:    fiber.CookieSameSiteNoneMode,
							SessionOnly: false,
						})

						ctx.Cookie(&fiber.Cookie{
							Name:        srv.conf.Components.AccessSystem.CookieKeyForRefreshToken,
							Value:       "",
							Path:        "/",
							Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
							MaxAge:      0,
							Expires:     time.Unix(0, 0),
							Secure:      false,
							HTTPOnly:    false,
							SameSite:    fiber.CookieSameSiteNoneMode,
							SessionOnly: false,
						})
					}
				}

				// Отправка ответа
				{
					if err = http_rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
						srv.components.Logger.Error().
							Format("The error response could not be recorded: '%s'. ", err).Write()

						return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
					}

					return
				}
			}).Name(id)

			var route = srv.app.GetRoute(id)

			srv.postman.AddItem(&postman.Items{
				Name: "Запрос для завершения сессии пользователя. ",
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
						Name:   "Пример успешного ответа. ",
						Status: string(fiber.StatusOK),
						Code:   fiber.StatusOK,
						Body: `
{
    "code": 200,
    "code_message": "OK",
    "status": "success"
}
`,
					},
				},
			})
		}
	}

	// /nginx
	{
		var (
			router = router.Group("/nginx")
		)

		// GET /auth
		{
			router.Get("/auth", srv.components.AccessSystem.BasicAuthentication)
		}
	}

	srv.components.Logger.Info().
		Text("Http rest api server routes are initialized. ").Write()

	return nil
}
