package http_rest_api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	common_errors "sm-box/internal/common/errors"
	"sm-box/internal/services/users/objects/models"
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

	// /access_system
	{
		var (
			router       = router.Group("/access_system")
			postmanGroup = srv.postman.AddItemGroup("Система доступа. ")
		)

		// /roles
		{
			var (
				router       = router.Group("/roles")
				postmanGroup = postmanGroup.AddItemGroup("Роли. ")
			)

			// GET /select
			{
				var id = uuid.New().String()

				router.Get("/select", func(ctx *fiber.Ctx) (err error) {
					type Response struct {
						List []*models.RoleInfo `json:"list" xml:"List"`
					}

					var response = new(Response)

					// Обработка
					{
						var cErr c_errors.RestAPI

						if response.List, cErr = srv.controllers.AccessSystem.GetRolesListForSelect(ctx.Context()); cErr != nil {
							srv.components.Logger.Error().
								Format("Couldn't get a list of access system roles for selects: '%s'. ", cErr).Write()

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
					Name: "Получение списка ролей системы доступа для select'ов. ",
					Description: `
Используется для получения списка ролей системы доступа для select'ов.
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

		// /permissions
		{
			var (
				router       = router.Group("/permissions")
				postmanGroup = postmanGroup.AddItemGroup("Права. ")
			)

			// GET /select
			{
				var id = uuid.New().String()

				router.Get("/select", func(ctx *fiber.Ctx) (err error) {
					type Response struct {
						List []*models.PermissionInfo `json:"list" xml:"List"`
					}

					var response = new(Response)

					// Обработка
					{
						var cErr c_errors.RestAPI

						if response.List, cErr = srv.controllers.AccessSystem.GetPermissionsListForSelect(ctx.Context()); cErr != nil {
							srv.components.Logger.Error().
								Format("Couldn't get a list of access system permissions for selects: '%s'. ", cErr).Write()

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
					Name: "Получение списка прав системы доступа для select'ов. ",
					Description: `
Используется для получения списка прав системы доступа для select'ов.
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
	}

	srv.components.Logger.Info().
		Text("Http rest api server routes are initialized. ").Write()

	return nil
}
