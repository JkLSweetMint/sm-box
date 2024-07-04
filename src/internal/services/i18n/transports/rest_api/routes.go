package rest_api

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"regexp"
	error_list "sm-box/internal/common/errors"
	common_entities "sm-box/internal/common/objects/entities"
	common_models "sm-box/internal/common/objects/models"
	rest_api_io "sm-box/internal/common/transports/rest_api/io"
	"sm-box/internal/services/i18n/infrastructure/objects/models"
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

	// /texts
	{
		var (
			router     = router.Group("/texts")
			postmanGrp = eng.postman.AddItemGroup("Текста локализации. ")
		)

		// /dictionary
		{
			var route = &common_entities.HttpRouteConstructor{
				Name: "Получение текстов локализации на секции. ",
				Description: `
Используется для получение текстов локализации на секции, передается путь к секции родителя,
запрос возвращает текста в том числе с дочерних секций.
`,

				Authorize: false,
			}

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

						if tok, cErr = eng.controllers.Identification.GetToken(ctx.Context(), data); cErr != nil {
							eng.components.Logger.Error().
								Format("Failed to get token data: '%s'. ", cErr).Write()

							return rest_api_io.WriteError(ctx, cErr)
						}
					}

					// Получение текстов
					{
						var cErr c_errors.RestAPI

						if response.Dictionary, cErr = eng.controllers.Texts.AssembleDictionary(ctx.Context(), tok.Language, queryArgs.Paths); cErr != nil {
							eng.components.Logger.Error().
								Format("The localization dictionary could not be assembled: '%s'. ", cErr).Write()

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

			postmanGrp.AddItem(&postman.Items{
				Name: route.Name,
				Request: &postman.Request{
					URL: &postman.URL{
						Protocol: eng.conf.Postman.Protocol,
						Host:     strings.Split(eng.conf.Postman.Host, "."),
						Path:     strings.Split(route.Path, "/"),
						Query: []*postman.QueryParam{
							{
								Key:   "paths",
								Value: "",
							},
						},
					},
					Method:      postman.Method(route.Method),
					Description: route.Description,
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

	eng.components.Logger.Info().
		Text("Http rest api routes are initialized. ").Write()
}