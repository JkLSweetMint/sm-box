package rest_api

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"regexp"
	error_list "sm-box/internal/common/errors"
	common_entities "sm-box/internal/common/objects/entities"
	common_models "sm-box/internal/common/objects/models"
	rest_api_io "sm-box/internal/common/transports/rest_api/io"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/http/postman"
	"strings"
)

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

	// POST /basic-auth
	{
		var route = &common_entities.HttpRouteConstructor{
			Name: "Запрос для базовой авторизации пользователя. ",
			Description: `
Используется для аутентификации пользователя по логину и паролю. Этот запрос принимает два параметра: username 
(имя пользователя) и password (пароль). Сервер проверяет предоставленные учетные данные и, если они верны, авторизирует
токен хранящийся в Cookie, который затем будет использован для выполнения других операций в системе.
`,

			Authorize: false,
		}

		router.Post("/basic-auth", func(ctx fiber.Ctx) (err error) {
			type Request struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}
			type Response struct {
				Token *common_models.JwtTokenInfo `json:"token" xml:"Token"`
				User  *common_models.UserInfo     `json:"user"  xml:"User"`
			}

			var (
				request  = new(Request)
				response = new(Response)
			)

			// Чтение данных
			{
				if err = ctx.Bind().Body(request); err != nil {
					eng.components.Logger.Error().
						Format("The request body data could not be read: '%s'. ", err).Write()

					if err = rest_api_io.WriteError(ctx, error_list.RequestBodyDataCouldNotBeRead_RestAPI()); err != nil {
						eng.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						var cErr = error_list.ResponseCouldNotBeRecorded_RestAPI()
						cErr.SetError(err)

						return rest_api_io.WriteError(ctx, cErr)
					}

					return
				}
			}

			// Обработка
			{
				var (
					tokenData string
					cErr      c_errors.RestAPI
				)

				// Получение токена
				{
					if tokenData = ctx.Cookies(eng.conf.Components.AccessSystem.CookieKeyForToken); tokenData == "" {
						var (
							value   = ctx.Response().Header.PeekCookie(eng.conf.Components.AccessSystem.CookieKeyForToken)
							pattern = fmt.Sprintf(`^%s=([\s\S]+);\sexpires=[\s\S]+;\sdomain=[\s\S]+;\spath=[\s\S]+;\sSameSite=[\s\S]+$`, eng.conf.Components.AccessSystem.CookieKeyForToken)
							re      = regexp.MustCompile(pattern)
						)

						tokenData = re.FindStringSubmatch(string(value))[1]
					}
				}

				if response.Token, response.User, cErr = eng.controllers.Authentication.BasicAuth(ctx.Context(),
					tokenData,
					request.Username,
					request.Password); cErr != nil {
					eng.components.Logger.Error().
						Format("User authorization failed: '%s'. ", cErr).Write()

					return rest_api_io.WriteError(ctx, cErr)
				}
			}

			// Отправка ответа
			{
				if err = rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
					eng.components.Logger.Error().
						Format("The response could not be recorded: '%s'. ", err).Write()

					return rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
				}
				return
			}
		}).Name(route.Name)

		route.Fill(eng.app.GetRoute(route.Name))

		if err := eng.components.AccessSystem.RegisterRoutes(route); err != nil {
			eng.components.Logger.Error().
				Format("An error occurred during the registration of http router routes: '%s'. ", err).Write()
			return
		}

		eng.postman.AddItem(&postman.Items{
			Name: route.Name,
			Request: &postman.Request{
				URL: &postman.URL{
					Protocol: eng.conf.Postman.Protocol,
					Host:     strings.Split(eng.conf.Postman.Host, "."),
					Path:     strings.Split(route.Path, "/"),
				},
				Method:      postman.Method(route.Method),
				Description: route.Description,
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

	// /projects
	{
		var router = router.Group("/projects")

		// GET /select
		{
			var route = &common_entities.HttpRouteConstructor{
				Name: "Получение списка проектов пользователя для выбора после авторизации. ",
				Description: `
Используется для получение списка проектов пользователя для выбора после авторизации. 
После выбора проекта пользователю будут доступны методы управления проектом.
`,

				Authorize: true,
			}

			router.Get("/select", func(ctx fiber.Ctx) (err error) {
				type Response struct {
					Projects common_models.ProjectList `json:"projects" xml:"Projects"`
				}

				var response = new(Response)

				// Обработка
				{
					var tok *common_models.JwtTokenInfo

					// Получение токена
					{
						var (
							data string
							cErr c_errors.RestAPI
						)

						if data = ctx.Cookies(eng.conf.Components.AccessSystem.CookieKeyForToken); data == "" {
							var (
								value   = ctx.Response().Header.PeekCookie(eng.conf.Components.AccessSystem.CookieKeyForToken)
								pattern = fmt.Sprintf(`^%s=([\s\S]+);\sexpires=[\s\S]+;\sdomain=[\s\S]+;\spath=[\s\S]+;\sSameSite=[\s\S]+$`, eng.conf.Components.AccessSystem.CookieKeyForToken)
								re      = regexp.MustCompile(pattern)
							)

							data = re.FindStringSubmatch(string(value))[1]
						}

						if tok, cErr = eng.controllers.Authentication.GetToken(ctx.Context(), data); cErr != nil {
							eng.components.Logger.Error().
								Format("Failed to get token data: '%s'. ", cErr).Write()

							return rest_api_io.WriteError(ctx, cErr)
						}
					}

					// Получение списка
					{
						var cErr c_errors.RestAPI

						if response.Projects, cErr = eng.controllers.Authentication.GetUserProjectsList(ctx.Context(), tok.ID, tok.UserID); cErr != nil {
							eng.components.Logger.Error().
								Format("The list of user's projects could not be retrieved: '%s'. ", cErr).Write()

							return rest_api_io.WriteError(ctx, cErr)
						}
					}
				}

				// Отправка ответа
				{
					if err = rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
						eng.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						return rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
					}
					return
				}
			}).Name(route.Name)

			route.Fill(eng.app.GetRoute(route.Name))

			if err := eng.components.AccessSystem.RegisterRoutes(route); err != nil {
				eng.components.Logger.Error().
					Format("An error occurred during the registration of http router routes: '%s'. ", err).Write()
				return
			}

			eng.postman.AddItem(&postman.Items{
				Name: route.Name,
				Request: &postman.Request{
					URL: &postman.URL{
						Protocol: eng.conf.Postman.Protocol,
						Host:     strings.Split(eng.conf.Postman.Host, "."),
						Path:     strings.Split(route.Path, "/"),
					},
					Method:      postman.Method(route.Method),
					Description: route.Description,
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
			var route = &common_entities.HttpRouteConstructor{
				Name:        "Выбрать проект пользователя для дальнейше работы после авторизации. ",
				Description: ``,

				Authorize: true,
			}

			router.Post("/set", func(ctx fiber.Ctx) (err error) {
				type Request struct {
					ProjectID types.ID `json:"project_id"`
				}
				type Response struct{}

				var (
					request  = new(Request)
					response = new(Response)
				)

				// Чтение данных
				{
					if err = ctx.Bind().Body(request); err != nil {
						eng.components.Logger.Error().
							Format("The request body data could not be read: '%s'. ", err).Write()

						if err = rest_api_io.WriteError(ctx, error_list.RequestBodyDataCouldNotBeRead_RestAPI()); err != nil {
							eng.components.Logger.Error().
								Format("The response could not be recorded: '%s'. ", err).Write()

							var cErr = error_list.ResponseCouldNotBeRecorded_RestAPI()
							cErr.SetError(err)

							return rest_api_io.WriteError(ctx, cErr)
						}

						return
					}
				}

				// Обработка
				{
					var tok *common_models.JwtTokenInfo

					// Получение токена
					{
						var (
							data string
							cErr c_errors.RestAPI
						)

						if data = ctx.Cookies(eng.conf.Components.AccessSystem.CookieKeyForToken); data == "" {
							var (
								value   = ctx.Response().Header.PeekCookie(eng.conf.Components.AccessSystem.CookieKeyForToken)
								pattern = fmt.Sprintf(`^%s=([\s\S]+);\sexpires=[\s\S]+;\sdomain=[\s\S]+;\spath=[\s\S]+;\sSameSite=[\s\S]+$`, eng.conf.Components.AccessSystem.CookieKeyForToken)
								re      = regexp.MustCompile(pattern)
							)

							data = re.FindStringSubmatch(string(value))[1]
						}

						if tok, cErr = eng.controllers.Authentication.GetToken(ctx.Context(), data); cErr != nil {
							eng.components.Logger.Error().
								Format("Failed to get token data: '%s'. ", cErr).Write()

							return rest_api_io.WriteError(ctx, cErr)
						}
					}

					// Установка значения
					{
						var cErr c_errors.RestAPI

						if cErr = eng.controllers.Authentication.SetTokenProject(ctx.Context(), tok.ID, request.ProjectID); cErr != nil {
							eng.components.Logger.Error().
								Format("The project value for the user token could not be set: '%s'. ", cErr).Write()

							return rest_api_io.WriteError(ctx, cErr)
						}
					}
				}

				// Отправка ответа
				{
					if err = rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
						eng.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						return rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
					}
					return
				}
			}).Name(route.Name)

			route.Fill(eng.app.GetRoute(route.Name))

			if err := eng.components.AccessSystem.RegisterRoutes(route); err != nil {
				eng.components.Logger.Error().
					Format("An error occurred during the registration of http router routes: '%s'. ", err).Write()
				return
			}

			eng.postman.AddItem(&postman.Items{
				Name: route.Name,
				Request: &postman.Request{
					URL: &postman.URL{
						Protocol: eng.conf.Postman.Protocol,
						Host:     strings.Split(eng.conf.Postman.Host, "."),
						Path:     strings.Split(route.Path, "/"),
					},
					Method:      postman.Method(route.Method),
					Description: route.Description,
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

	eng.components.Logger.Info().
		Text("Http rest api routes are initialized. ").Write()
}
