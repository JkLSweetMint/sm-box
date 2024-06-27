package rest_api

import (
	"github.com/gofiber/fiber/v3"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/common/objects/models"
	rest_api_io "sm-box/internal/common/transports/rest_api/io"
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

	// /basic-auth
	{
		var route = fiber.Route{Name: "Запрос для базовой авторизации пользователя. "}

		router.Post("/basic-auth", func(ctx fiber.Ctx) (err error) {
			type Response struct {
				Token *models.JwtTokenInfo `json:"token" xml:"Token"`
				User  *models.UserInfo     `json:"user"  xml:"User"`
			}
			type Request struct {
				Username string `json:"username" xml:"Username"`
				Password string `json:"password" xml:"Password"`
			}

			var (
				response = new(Response)
				request  = new(Request)
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
					tokenData = ctx.Cookies(eng.conf.Components.AccessSystem.CookieKeyForToken)
					cErr      c_errors.RestAPI
				)

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

		route = eng.app.GetRoute(route.Name)

		eng.postman.AddItem(&postman.Items{
			Name: route.Name,
			Request: &postman.Request{
				URL: &postman.URL{
					Protocol: eng.conf.Postman.Protocol,
					Host:     strings.Split(eng.conf.Postman.Host, "."),
					Path:     strings.Split(route.Path, "/"),
				},
				Method: postman.Method(route.Method),
				Description: `
Используется для аутентификации пользователя по логину и паролю. Этот запрос принимает два параметра: username 
(имя пользователя) и password (пароль). Сервер проверяет предоставленные учетные данные и, если они верны, авторизирует
токен хранящийся в Cookie, который затем будет использован для выполнения других операций в системе.
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
        "id": "E-000104",
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
        "id": "I-000003",
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

	eng.components.Logger.Info().
		Text("Http rest api routes are initialized. ").Write()
}
