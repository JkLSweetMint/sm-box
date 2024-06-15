package rest_api

import (
	"errors"
	"github.com/gofiber/fiber/v3"
	"sm-box/internal/app/transports/rest_api/io"
	"sm-box/internal/common/errors"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/http/postman"
	"strings"
	"time"
)

// initBaseRoutes - регистрация базовых маршрутов системы.
func (eng *engine) initBaseRoutes() {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal, tracer.LevelTransport)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished() }()
	}

	eng.components.Logger.Info().
		Text("Starting initialization of basic http rest api routes... ").Write()

	// /sys
	{
		var (
			router     = eng.router.Group("/sys")
			postmanGrp = eng.postman.AddItemGroup("Системные запросы. ")
		)

		// GET /ping
		{
			var route = fiber.Route{Name: "Ping-pong запрос. "}

			router.Get("/ping", func(ctx fiber.Ctx) (err error) {
				type Response struct {
					Message string `json:"message" xml:"Message"`
				}

				var (
					response = new(Response)
				)

				// Обработка данных
				{
					response.Message = "pong"
				}

				// Отправка ответа
				{
					if err = rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
						eng.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						var cErr = error_list.ErrResponseCouldNotBeRecorded_RestAPI()
						cErr.SetError(err)

						return rest_api_io.WriteError(ctx, cErr)
					}

					return
				}
			}).Name(route.Name)

			route = eng.app.GetRoute(route.Name)

			postmanGrp.AddItem(&postman.Items{
				Name: route.Name,
				Request: &postman.Request{
					URL: &postman.URL{
						Protocol: eng.conf.Postman.Protocol,
						Host:     strings.Split(eng.conf.Postman.Host, "."),
						Path:     strings.Split(route.Path, "/"),
					},
					Method: postman.Method(route.Method),
					Description: `
Запрос /ping в HTTP используется для проверки работоспособности сервера. Этот запрос очень простой и обычно состоит из 
одной строки, содержащей только слово "ping". Сервер, получив такой запрос, должен ответить сообщением об успешной 
обработке запроса.

Зачем это нужно? Во-первых, такие запросы используются при тестировании сети, чтобы убедиться, что сервер доступен и 
работает корректно. Во-вторых, они могут использоваться в клиентском программном обеспечении для периодической проверки
состояния соединения с сервером. Например, если клиент долгое время не получает ответов от сервера, он может отправить
запрос /ping, чтобы понять, не отключился ли сервер.

Важно отметить, что запрос /ping является информационным и не требует никаких данных от клиента. Это означает, что он 
не занимает много времени и ресурсов сервера, поэтому его можно использовать достаточно часто без риска перегрузки 
системы.`,
					Body: &postman.Body{
						Mode: "raw",
						Options: &postman.BodyOptions{
							Raw: postman.BodyOptionsRaw{
								Language: postman.JSON,
							},
						},
					},
				},
				Responses: []*postman.Response{
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
        "message": "pong"
    }
}
`,
					},
				},
			})
		}

		// GET /health
		{
			var route = fiber.Route{Name: "Используется для проверки работоспособности сервера. "}

			router.Get("/health", func(ctx fiber.Ctx) (err error) {
				type Response struct {
					SystemName string `json:"system_name" xml:"SystemName"`
					Mode       string `json:"mode"        xml:"Mode"`
					Version    string `json:"version"     xml:"Version"`
					OS         string `json:"os"          xml:"OS"`

					LaunchTime string `json:"launch_time" xml:"LaunchTime"`
					UpTime     string `json:"up_time"     xml:"UpTime"`
				}

				var (
					response = new(Response)
				)

				// Обработка данных
				{
					response.SystemName = env.Vars.SystemName
					response.Mode = env.Mode.String()
					response.Version = env.Version
					response.OS = env.OS

					response.LaunchTime = env.Vars.LaunchTime.UTC().Format(time.RFC3339Nano)
					response.UpTime = time.Now().Sub(env.Vars.LaunchTime).String()
				}

				// Отправка ответа
				{
					if err = rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
						eng.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						var cErr = error_list.ErrResponseCouldNotBeRecorded_RestAPI()
						cErr.SetError(err)

						return rest_api_io.WriteError(ctx, cErr)
					}

					return
				}
			}).Name(route.Name)

			route = eng.app.GetRoute(route.Name)

			postmanGrp.AddItem(&postman.Items{
				Name: route.Name,
				Request: &postman.Request{
					URL: &postman.URL{
						Protocol: eng.conf.Postman.Protocol,
						Host:     strings.Split(eng.conf.Postman.Host, "."),
						Path:     strings.Split(route.Path, "/"),
					},
					Method: postman.Method(route.Method),
					Description: `
Запрос /health в HTTP также используется для проверки работоспособности сервера, но в отличие от /ping, он 
предоставляет более подробную информацию о состоянии сервера. 

Этот запрос обычно возвращает JSON-объект, который содержит различные метрики, связанные с работой сервера, такие как
версия ПО, список активных подключений, информация о нагрузке на процессор и память, а также другие важные параметры.

Зачем это нужно? Во-первых, такие запросы используются при развертывании новых версий ПО, чтобы убедиться, что все 
компоненты системы работают корректно. Во-вторых, они могут использоваться в системах мониторинга для автоматического 
обнаружения проблем в работе сервера.

Важно отметить, что запрос /health является информационным и не требует никаких действий со стороны клиента. Это 
означает, что он не занимает много времени и ресурсов сервера, поэтому его можно использовать достаточно часто без 
риска перегрузки системы.`,
					Body: &postman.Body{
						Mode: "raw",
						Options: &postman.BodyOptions{
							Raw: postman.BodyOptionsRaw{
								Language: postman.JSON,
							},
						},
					},
				},
				Responses: []*postman.Response{
					{
						Name:   "Пример ответа. ",
						Status: string(fiber.StatusOK),
						Code:   fiber.StatusOK,
						Body: `
{
    "code": 200,
    "code_message": "OK",
    "status": "success",
    "data": {
        "system_name": "box",
        "mode": "DEV",
        "version": "24.0.16",
        "os": "windows - amd64",
        "launch_time": "2024-06-10T07:58:55.981117Z",
        "up_time": "5.2450825s"
    }
}
`,
					},
				},
			})
		}

		// GET /error
		{
			var route = fiber.Route{Name: "Используется для получения примера ошибки для возможности обработки клиентом. "}

			router.Get("/error", func(ctx fiber.Ctx) (err error) {
				var (
					response c_errors.RestAPI
				)

				// Обработка данных
				{
					response = error_list.ErrUnknown_RestAPI()
					response.SetError(errors.New("Test. "))
				}

				// Отправка ответа
				{
					if err = rest_api_io.WriteError(ctx, response); err != nil {
						eng.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						var cErr = error_list.ErrResponseCouldNotBeRecorded_RestAPI()
						cErr.SetError(err)

						return rest_api_io.WriteError(ctx, cErr)
					}

					return
				}
			}).Name(route.Name)

			route = eng.app.GetRoute(route.Name)

			postmanGrp.AddItem(&postman.Items{
				Name: route.Name,
				Request: &postman.Request{
					URL: &postman.URL{
						Protocol: eng.conf.Postman.Protocol,
						Host:     strings.Split(eng.conf.Postman.Host, "."),
						Path:     strings.Split(route.Path, "/"),
					},
					Method: postman.Method(route.Method),
					Description: `
Запрос /error в HTTP используется для получения примера ошибки для возможности обработки клиентом.

Важно отметить, что запрос /error является информационным и не требует никаких действий со стороны клиента. Это 
означает, что он не занимает много времени и ресурсов сервера, поэтому его можно использовать достаточно часто без 
риска перегрузки системы.
`,
					Body: &postman.Body{
						Mode: "raw",
						Options: &postman.BodyOptions{
							Raw: postman.BodyOptionsRaw{
								Language: postman.JSON,
							},
						},
					},
				},
				Responses: []*postman.Response{
					{
						Name:   "Пример возврата ошибки. ",
						Status: string(fiber.StatusInternalServerError),
						Code:   fiber.StatusInternalServerError,
						Body: `
{
    "code": 500,
    "code_message": "Internal Server Error",
    "status": "error",
    "error": {
        "id": "I-000000",
        "type": "unknown",
        "status": "unknown",
        "message": "Unknown error. ",
        "details": {
            "timestamp": 1718005889
        }
    }
}
`,
					},
				},
			})
		}
	}

	eng.components.Logger.Info().
		Text("The basic http rest api routes are initialized. ").Write()
}
