package rest_api

import (
	"github.com/gofiber/fiber/v3"
	"sm-box/pkg/core/components/tracer"
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

		router.Post("/basic-auth", eng.components.AccessSystem.BasicUserAuth).Name(route.Name)

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
