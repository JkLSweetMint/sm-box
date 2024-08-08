package http_rest_api

import (
	"context"
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	common_errors "sm-box/internal/common/errors"
	common_types "sm-box/internal/common/types"
	authentication_entities "sm-box/internal/services/authentication/objects/entities"
	"sm-box/internal/services/notifications/components/notification_notifier"
	"sm-box/internal/services/notifications/objects"
	"sm-box/internal/services/notifications/objects/constructors"
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

	// /ws
	{
		var router = router.Group("/ws")

		router.Use("/", func(ctx *fiber.Ctx) (err error) {
			if cookie := ctx.Cookies(srv.conf.Components.AccessSystem.CookieKeyForSessionToken); websocket.IsWebSocketUpgrade(ctx) && cookie != "" {
				ctx.Locals("allowed", true)
				return ctx.Next()
			}
			return fiber.ErrUpgradeRequired
		})

		router.Get("/", websocket.New(func(conn *websocket.Conn) {
			var (
				ctx, disconnect = context.WithCancel(context.Background())
				recipient       = new(notification_notifier.Recipient)
			)

			// Закрытие соединения при завершении работы
			{
				defer func() {
					if len(recipient.Keys) > 0 {
						srv.components.NotificationNotifier.RemoveRecipient(recipient.JwtToken.ID)
					}

					if err := conn.Close(); err != nil {
						srv.components.Logger.Warn().
							Format("An error occurred when closing the Web socket channel: '%s'. ", err).
							Field("addr", conn.RemoteAddr()).
							Field("recipient", recipient).Write()
					}

					srv.components.Logger.Info().
						Text("Client disconnected. ").
						Field("addr", conn.RemoteAddr()).
						Field("recipient", recipient).Write()
				}()
			}

			// Подготовка
			{
				srv.components.Logger.Info().
					Text("Client connected. ").
					Field("addr", conn.RemoteAddr()).Write()

				// Формирование ключей
				{
					var raw = conn.Cookies(srv.conf.Components.AccessSystem.CookieKeyForSessionToken)

					recipient.JwtToken = new(authentication_entities.JwtSessionToken)

					if err := recipient.JwtToken.Parse(raw); err != nil {
						srv.components.Logger.Error().
							Format("Failed to get session token data: '%s'. ", err).
							Field("raw", raw).Write()

						return
					}

					recipient.Keys = []string{
						fmt.Sprintf("session:%s", recipient.JwtToken.ID.String()),
						fmt.Sprintf("users:%d", recipient.JwtToken.UserID),
					}
				}

				// Получение канала
				{
					srv.components.NotificationNotifier.RegisterRecipient(recipient)
				}
			}

			// Обработка web-сокета
			{
				go func() {
					for {
						if _, _, err := conn.ReadMessage(); err != nil {
							if err.Error() != "websocket: close 1000 (normal)" {
								srv.components.Logger.Warn().
									Format("Error reading the data to the web socket channel: '%s'. ", err).
									Field("addr", conn.RemoteAddr()).
									Field("recipient", recipient).Write()
							}

							disconnect()
							break
						}
					}
				}()
			}

			// Логика
			{
			WsHandler:
				for {
					select {
					case data := <-recipient.Channel():
						{
							srv.components.Logger.Info().
								Text("Sending data to a web socket channel. ").
								Field("addr", conn.RemoteAddr()).
								Field("recipient", recipient).
								Field("data", data).Write()

							if err := conn.WriteJSON(data); err != nil {
								srv.components.Logger.Warn().
									Format("Error writing the response to the web socket channel: '%s'. ", err).
									Field("addr", conn.RemoteAddr()).
									Field("recipient", recipient).Write()

								disconnect()
								break WsHandler
							}
						}
					case <-ctx.Done():
						break WsHandler
					case <-srv.ctx.Done():
						disconnect()
						break WsHandler
					}
				}
			}
		}))
	}

	// /users
	{
		var (
			router       = router.Group("/users")
			postmanGroup = srv.postman.AddItemGroup("Пользовательские уведомления. ")
		)

		// GET /list
		{
			var id = uuid.New().String()

			router.Get("/list", func(ctx *fiber.Ctx) (err error) {
				type Response struct {
					Count        int64 `json:"count"          xml:"count,attr"`
					CountNotRead int64 `json:"count_not_read" xml:"count_not_read,attr"`

					List []*models.UserNotificationInfo `json:"list" xml:"List"`
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
					if err = ctx.QueryParser(queryArgs); err != nil {
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
						search      *objects.UserNotificationSearch
						pagination  *objects.UserNotificationPagination
						filters     *objects.UserNotificationFilters
						recipientID common_types.ID
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

									err = ctx.Redirect("/errors/403")
									return
								}
							}

							recipientID = token.UserID
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

					if response.Count, response.CountNotRead, response.List, cErr = srv.controllers.UserNotifications.GetList(ctx.Context(), recipientID, search, pagination, filters); cErr != nil {
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

			router.Put("/read", func(ctx *fiber.Ctx) (err error) {
				type Request struct {
					IDs []common_types.ID `json:"ids"`
				}
				type Response struct{}

				var (
					request  = new(Request)
					response = new(Response)
				)

				// Чтение данных
				{
					if err = ctx.BodyParser(request); err != nil {
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
					var recipientID common_types.ID

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

									err = ctx.Redirect("/errors/403")
									return
								}
							}

							recipientID = token.UserID
						}
					}

					if cErr := srv.controllers.UserNotifications.Read(ctx.Context(), recipientID, request.IDs...); cErr != nil {
						srv.components.Logger.Error().
							Format("Several user notifications could not be read: '%s'. ", cErr).Write()

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
	"ids": []
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

			router.Delete("/", func(ctx *fiber.Ctx) (err error) {
				type Request struct {
					IDs []common_types.ID `json:"ids"`
				}
				type Response struct{}

				var (
					request  = new(Request)
					response = new(Response)
				)

				// Чтение данных
				{
					if err = ctx.BodyParser(request); err != nil {
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
					var recipientID common_types.ID

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

									err = ctx.Redirect("/errors/403")
									return
								}
							}

							recipientID = token.UserID
						}
					}

					if cErr := srv.controllers.UserNotifications.Remove(ctx.Context(), recipientID, request.IDs...); cErr != nil {
						srv.components.Logger.Error().
							Format("Several user notifications could not be removed: '%s'. ", cErr).Write()

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
	"ids": []
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

		// /create
		{
			var (
				router = router.Group("/create")
			)

			// POST /
			{
				var id = uuid.New().String()

				router.Post("/", func(ctx *fiber.Ctx) (err error) {
					type Request struct {
						Type types.NotificationType `json:"type"`

						SenderID    common_types.ID `json:"sender_id"`
						RecipientID common_types.ID `json:"recipient_id"`

						Title     string    `json:"title"`
						TitleI18n uuid.UUID `json:"title_i18n"`

						Text     string    `json:"text"`
						TextI18n uuid.UUID `json:"text_i18n"`
					}
					type Response struct {
						Notification *models.UserNotificationInfo `json:"notification" xml:"Notification"`
					}

					var (
						request  = new(Request)
						response = new(Response)
					)

					// Чтение данных
					{
						if err = ctx.BodyParser(request); err != nil {
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
						var (
							cErr        c_errors.RestAPI
							constructor *constructors.UserNotification
						)

						// Создание конструктора
						{
							constructor = &constructors.UserNotification{
								Type: request.Type,

								SenderID:    request.SenderID,
								RecipientID: request.RecipientID,

								Title:     request.Title,
								TitleI18n: request.TitleI18n,

								Text:     request.Text,
								TextI18n: request.TextI18n,
							}
						}

						if response.Notification, cErr = srv.controllers.UserNotifications.CreateOne(ctx.Context(), constructor); cErr != nil {
							srv.components.Logger.Error().
								Format("Failed to create a user notification: '%s'. ", cErr).Write()

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
					Name: "Создание пользовательского уведомления. ",
					Description: `
Используется для создания пользовательского уведомления.
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
	"type": "",
	"sender_id": 0,
	"recipient_id": 0,
	"title": "",
	"title_i18n": "",
	"text": "",
	"text_i18n": ""
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

			// POST /multiple
			{
				var id = uuid.New().String()

				router.Post("/multiple", func(ctx *fiber.Ctx) (err error) {
					type Request struct {
						Notifications []*struct {
							Type types.NotificationType `json:"type"`

							SenderID    common_types.ID `json:"sender_id"`
							RecipientID common_types.ID `json:"recipient_id"`

							Title     string    `json:"title"`
							TitleI18n uuid.UUID `json:"title_i18n"`

							Text     string    `json:"text"`
							TextI18n uuid.UUID `json:"text_i18n"`
						} `json:"notifications"`
					}
					type Response struct {
						Notifications []*models.UserNotificationInfo `json:"notifications" xml:"Notifications>Notification"`
					}

					var (
						request  = new(Request)
						response = new(Response)
					)

					// Чтение данных
					{
						if err = ctx.BodyParser(request); err != nil {
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
						var (
							cErr c_errors.RestAPI
							list []*constructors.UserNotification
						)

						// Создание конструкторов
						{
							for _, notification := range request.Notifications {
								list = append(list, &constructors.UserNotification{
									Type: notification.Type,

									SenderID:    notification.SenderID,
									RecipientID: notification.RecipientID,

									Title:     notification.Title,
									TitleI18n: notification.TitleI18n,

									Text:     notification.Text,
									TextI18n: notification.TextI18n,
								})
							}
						}

						if response.Notifications, cErr = srv.controllers.UserNotifications.Create(ctx.Context(), list...); cErr != nil {
							srv.components.Logger.Error().
								Format("Failed to create a user notifications: '%s'. ", cErr).Write()

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
					Name: "Создание нескольких пользовательских уведомлений. ",
					Description: `
Используется для создания нескольких пользовательских уведомлений.
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
	"notifications": [
		{
			"type": "",
			"sender_id": 0,
			"recipient_id": 0,
			"title": "",
			"title_i18n": "",
			"text": "",
			"text_i18n": ""
		}
	]
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
	}

	srv.components.Logger.Info().
		Text("Http rest api server routes are initialized. ").Write()

	return nil
}
