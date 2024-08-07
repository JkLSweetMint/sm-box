package http_rest_api

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	common_errors "sm-box/internal/common/errors"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/http/postman"
	"sm-box/pkg/http/rest_api/io"
	"strings"
	"time"
)

// registerBaseRoutes - регистрация базовых маршрутов сервера.
func (srv *server) registerBaseRoutes() error {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal, tracer.LevelTransport)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished() }()
	}

	srv.components.Logger.Info().
		Text("Starting initialization of basic http rest api server routes... ").Write()

	var router = srv.router

	// /sys
	{
		var (
			router       = router.Group("/sys")
			postmanGroup = srv.postman.AddItemGroup("Системные запросы. ")
		)

		// GET /ping
		{
			var id = uuid.New().String()

			router.Get("/ping", func(ctx *fiber.Ctx) (err error) {
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
					if err = http_rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
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
			}).Name(id)

			var route = srv.app.GetRoute(id)

			postmanGroup.AddItem(&postman.Items{
				Name: "Ping-pong запрос. ",
				Request: &postman.Request{
					URL: &postman.URL{
						Protocol: srv.conf.Postman.Protocol,
						Host:     strings.Split(srv.conf.Postman.Host, "."),
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
			var id = uuid.New().String()

			router.Get("/health", func(ctx *fiber.Ctx) (err error) {
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
					if err = http_rest_api_io.Write(ctx.Status(fiber.StatusOK), response); err != nil {
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
			}).Name(id)

			var route = srv.app.GetRoute(id)

			postmanGroup.AddItem(&postman.Items{
				Name: "Используется для проверки работоспособности сервера. ",
				Request: &postman.Request{
					URL: &postman.URL{
						Protocol: srv.conf.Postman.Protocol,
						Host:     strings.Split(srv.conf.Postman.Host, "."),
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
			var id = uuid.New().String()

			router.Get("/error", func(ctx *fiber.Ctx) (err error) {
				var (
					response c_errors.RestAPI
				)

				// Обработка данных
				{
					response = c_errors.ToRestAPI(common_errors.Unknown())
					response.SetError(errors.New("Test. "))
				}

				// Отправка ответа
				{
					if err = http_rest_api_io.WriteError(ctx, response); err != nil {
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
			}).Name(id)

			var route = srv.app.GetRoute(id)

			postmanGroup.AddItem(&postman.Items{
				Name: "Используется для получения примера ошибки для возможности обработки клиентом. ",
				Request: &postman.Request{
					URL: &postman.URL{
						Protocol: srv.conf.Postman.Protocol,
						Host:     strings.Split(srv.conf.Postman.Host, "."),
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
            "key": "value"
        }
    }
}
`,
					},
				},
			})
		}
	}

	srv.components.Logger.Info().
		Text("The basic http rest api server routes are initialized. ").Write()

	return nil
}
