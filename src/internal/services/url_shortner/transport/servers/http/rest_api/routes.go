package http_rest_api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	common_errors "sm-box/internal/common/errors"
	common_types "sm-box/internal/common/types"
	authentication_entities "sm-box/internal/services/authentication/objects/entities"
	"sm-box/internal/services/url_shortner/objects"
	"sm-box/internal/services/url_shortner/objects/models"
	"sm-box/internal/services/url_shortner/objects/types"
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

	// ALL /use/
	{
		var (
			id    = uuid.New().String()
			route fiber.Route
		)

		router.All("/use/*", func(ctx fiber.Ctx) (err error) {
			var (
				reduction string
				token     *authentication_entities.JwtSessionToken
			)

			// Получение токена
			{
				if raw := ctx.Cookies(srv.conf.Components.AccessSystem.CookieKeyForSessionToken); len(raw) > 0 {
					token = new(authentication_entities.JwtSessionToken)

					if err = token.Parse(raw); err != nil {
						srv.components.Logger.Error().
							Format("Failed to get session token data: '%s'. ", err).
							Field("raw", raw).Write()

						err = ctx.Redirect().To("/errors/403")
						return
					}
				}
			}

			// Получение сокращения
			{
				reduction = strings.Replace(string(ctx.Request().URI().Path()), strings.Replace(route.Path, "/*", "/", 1), "", 1)
			}

			// Использование
			{
				var (
					url    *models.ShortUrlInfo
					status types.ShortUrlUsageHistoryStatus
				)

				// Обработка
				{
					var cErr c_errors.RestAPI

					if url, status, cErr = srv.controllers.Urls.Use(ctx.Context(), reduction, token); cErr != nil {
						srv.components.Logger.Error().
							Format("The short url could not be used: '%s'. ", cErr).Write()

						switch status {
						case types.ShortUrlUsageHistoryStatusForbidden:
							return ctx.Redirect().To("/errors/403")
						case types.ShortUrlUsageHistoryStatusFailed:
							return ctx.Redirect().To("/errors/50x")
						}

						if cErr.StatusCode() >= 500 {
							return ctx.Redirect().To("/errors/50x")
						} else if cErr.StatusCode() == 403 {
							return ctx.Redirect().To("/errors/403")
						}

						if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
							srv.components.Logger.Error().
								Format("The error response could not be recorded: '%s'. ", err).Write()

							return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
						}

						return
					}

					switch status {
					case types.ShortUrlUsageHistoryStatusForbidden:
						return ctx.Redirect().To("/errors/403")
					case types.ShortUrlUsageHistoryStatusFailed:
						return ctx.Redirect().To("/errors/50x")
					}
				}

				// Выполнение http запроса
				{
					switch url.Properties.Type {
					case types.ShortUrlTypeRedirect:
						{
							if err = ctx.Redirect().To(url.Source); err != nil {
								srv.components.Logger.Warn().
									Format("Failed to redirect a remote resource: '%s'. ", err).
									Field("url", url).Write()

								err = ctx.Redirect().To("/errors/50x")
								return
							}
						}
					case types.ShortUrlTypeProxy:
						{
							var client = new(fasthttp.Client)

							ctx.Request().URI().Update(url.Source)

							if err = client.Do(ctx.Request(), ctx.Response()); err != nil {
								srv.components.Logger.Warn().
									Format("Failed to proxy a remote resource: '%s'. ", err).
									Field("url", url).Write()

								err = ctx.Redirect().To("/errors/50x")
								return
							}
						}
					}
				}
			}

			return
		}).Name(id)

		route = srv.app.GetRoute(id)
	}

	// /management
	{
		var (
			router       = router.Group("/management")
			postmanGroup = srv.postman.AddItemGroup("Управление короткими url. ")
		)

		// GET /list
		{
			var id = uuid.New().String()

			router.Get("/list", func(ctx fiber.Ctx) (err error) {
				type Response struct {
					Count int64                  `json:"count" xml:"count,attr"`
					List  []*models.ShortUrlInfo `json:"list"  xml:"List"`
				}
				type QueryArgs struct {
					Search string `query:"search"`

					Limit  *int64 `query:"limit"`
					Offset *int64 `query:"offset"`

					SortKey  string `query:"sort_key"`
					SortType string `query:"sort_type"`

					FilterActive *string `query:"filter_active"`
					FilterType   *string `query:"filter_type"`

					FilterNumberOfUses     *int64  `query:"filter_number_of_uses"`
					FilterNumberOfUsesType *string `query:"filter_number_of_uses_type"`

					FilterStartActive     *string `query:"filter_start_active"`
					FilterStartActiveType *string `query:"filter_start_active_type"`

					FilterEndActive     *string `query:"filter_end_active"`
					FilterEndActiveType *string `query:"filter_end_active_type"`
				}

				var (
					response  = new(Response)
					queryArgs = new(QueryArgs)
				)

				// Чтение данных
				{
					if err = ctx.Bind().Query(queryArgs); err != nil {
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
						search     *objects.ShortUrlsListSearch
						sort       *objects.ShortUrlsListSort
						pagination *objects.ShortUrlsListPagination
						filters    *objects.ShortUrlsListFilters
					)

					// Обработка входных данных
					{
						// Поиск
						{
							queryArgs.Search = strings.TrimSpace(queryArgs.Search)

							search = &objects.ShortUrlsListSearch{
								Global: queryArgs.Search,
							}
						}

						// Сортировки
						{
							queryArgs.SortKey = strings.TrimSpace(queryArgs.SortKey)
							queryArgs.SortType = strings.TrimSpace(queryArgs.SortType)

							if queryArgs.SortKey != "" && queryArgs.SortType != "" {
								sort = &objects.ShortUrlsListSort{
									Key:  queryArgs.SortKey,
									Type: queryArgs.SortType,
								}
							}
						}

						// Пагинация
						{
							pagination = &objects.ShortUrlsListPagination{
								Offset: queryArgs.Offset,
								Limit:  queryArgs.Limit,
							}
						}

						// Фильтрация
						{
							filters = new(objects.ShortUrlsListFilters)

							if queryArgs.FilterActive != nil {
								switch strings.ToLower(strings.TrimSpace(*queryArgs.FilterActive)) {
								case "true", "on", "enable":
									{
										var v = true
										filters.Active = &v
									}
								case "false", "off", "disable":
									{
										var v = false
										filters.Active = &v
									}
								}
							}

							if queryArgs.FilterType != nil && *queryArgs.FilterType != "" {
								var v = types.ShortUrlType(strings.TrimSpace(*queryArgs.FilterType))
								filters.Type = &v
							}

							if queryArgs.FilterNumberOfUses != nil {
								filters.NumberOfUses = queryArgs.FilterNumberOfUses
								filters.NumberOfUsesType = queryArgs.FilterNumberOfUsesType
							}

							if queryArgs.FilterStartActive != nil && *queryArgs.FilterStartActive != "" {
								var v = strings.TrimSpace(*queryArgs.FilterStartActive)

								if tm, e := time.Parse(time.RFC3339Nano, v); e == nil {
									filters.StartActive = &tm
									filters.StartActiveType = queryArgs.FilterStartActiveType
								}
							}

							if queryArgs.FilterEndActive != nil && *queryArgs.FilterEndActive != "" {
								var v = strings.TrimSpace(*queryArgs.FilterEndActive)

								if tm, e := time.Parse(time.RFC3339Nano, v); e == nil {
									filters.EndActive = &tm
									filters.EndActiveType = queryArgs.FilterEndActiveType
								}
							}
						}
					}

					var cErr c_errors.RestAPI

					if response.Count, response.List, cErr = srv.controllers.UrlsManagement.GetList(ctx.Context(), search, sort, pagination, filters); cErr != nil {
						srv.components.Logger.Error().
							Format("Couldn't get a list of short url: '%s'. ", cErr).Write()

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
				Name: "Получение списка коротких url. ",
				Description: `
Используется для получения списка коротких url.
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
								Key:   "sort_key",
								Value: "",
							},
							{
								Key:   "sort_type",
								Value: "",
							},
							{
								Key:   "filter_active",
								Value: "",
							},
							{
								Key:   "filter_type",
								Value: "",
							},
							{
								Key:   "filter_start_active",
								Value: "",
							},
							{
								Key:   "filter_start_active_type",
								Value: "",
							},
							{
								Key:   "filter_end_active",
								Value: "",
							},
							{
								Key:   "filter_end_active_type",
								Value: "",
							},
							{
								Key:   "filter_number_of_uses",
								Value: "",
							},
							{
								Key:   "filter_number_of_uses_type",
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
        "count": 5,
        "list": [
            {
                "id": 1,
                "source": "https://www.youtube.com/watch?v=71pOiq-E_X4",
                "reduction": "Fu884hxjuhjEDSU0",
                "accesses": {
                    "roles_id": [],
                    "permissions_id": []
                },
                "properties": {
                    "type": "redirect",
                    "active": true,
                    "number_of_uses": 1,
                    "remained_number_of_uses": 0,
                    "start_active": "0001-01-01T04:00:00+04:00",
                    "end_active": "0001-01-01T04:00:00+04:00"
                }
            },
            {
                "id": 2,
                "source": "https://samgk.ru/files/officialdocs/admission/vacant_22.07.2024.pdf",
                "reduction": "9sC6RY73Q8Vyt06c",
                "accesses": {
                    "roles_id": [],
                    "permissions_id": []
                },
                "properties": {
                    "type": "proxy",
                    "active": true,
                    "number_of_uses": 0,
                    "remained_number_of_uses": 0,
                    "start_active": "0001-01-01T04:00:00+04:00",
                    "end_active": "0001-01-01T04:00:00+04:00"
                }
            },
            {
                "id": 7,
                "source": "https://habr.com/ru/companies/selectel/articles/831980/",
                "reduction": "twv3TNknSoXB2R0o",
                "accesses": {
                    "roles_id": [],
                    "permissions_id": []
                },
                "properties": {
                    "type": "redirect",
                    "active": false,
                    "number_of_uses": 1,
                    "remained_number_of_uses": 1,
                    "start_active": "2024-07-30T00:00:00+04:00",
                    "end_active": "2024-07-30T23:59:59+04:00"
                }
            }
        ]
    }
}
`,
					},
				},
			})
		}

		// GET /by_id/:id
		{
			var id = uuid.New().String()

			router.Get("/by_id/:id", func(ctx fiber.Ctx) (err error) {
				type Response struct {
					Url *models.ShortUrlInfo `json:"url" xml:"Url"`
				}
				type UriArgs struct {
					ID common_types.ID `uri:"id"`
				}

				var (
					response = new(Response)
					uriArgs  = new(UriArgs)
				)

				// Чтение данных
				{
					if err = ctx.Bind().URI(uriArgs); err != nil {
						srv.components.Logger.Error().
							Format("The request uri data could not be read: '%s'. ", err).Write()

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
					var cErr c_errors.RestAPI

					if response.Url, cErr = srv.controllers.UrlsManagement.GetOne(ctx.Context(), uriArgs.ID); cErr != nil {
						srv.components.Logger.Error().
							Format("Could not get the short url data by id: '%s'. ", cErr).Write()

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
				Name: "Получение короткого url по ID. ",
				Description: `
Используется для получения короткого url по ID.
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
        "url": {
            "id": 1,
            "source": "https://www.youtube.com/watch?v=71pOiq-E_X4",
            "reduction": "Fu884hxjuhjEDSU0",
            "accesses": {
                "roles_id": [],
                "permissions_id": []
            },
            "properties": {
                "type": "redirect",
                "active": true,
                "number_of_uses": 1,
                "remained_number_of_uses": 0,
                "start_active": "0001-01-01T04:00:00+04:00",
                "end_active": "0001-01-01T04:00:00+04:00"
            }
        }
    }
}
`,
					},
				},
			})
		}

		// GET /by_reduction/:reduction
		{
			var id = uuid.New().String()

			router.Get("/by_reduction/:reduction", func(ctx fiber.Ctx) (err error) {
				type Response struct {
					Url *models.ShortUrlInfo `json:"url" xml:"Url"`
				}
				type UriArgs struct {
					Reduction string `uri:"reduction"`
				}

				var (
					response = new(Response)
					uriArgs  = new(UriArgs)
				)

				// Чтение данных
				{
					if err = ctx.Bind().URI(uriArgs); err != nil {
						srv.components.Logger.Error().
							Format("The request uri data could not be read: '%s'. ", err).Write()

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
					var cErr c_errors.RestAPI

					if response.Url, cErr = srv.controllers.UrlsManagement.GetOneByReduction(ctx.Context(), uriArgs.Reduction); cErr != nil {
						srv.components.Logger.Error().
							Format("Could not get the short url data by reduction: '%s'. ", cErr).Write()

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
				Name: "Получение короткого url по сокращению. ",
				Description: `
Используется для получения короткого url по сокращению.
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
        "url": {
            "id": 1,
            "source": "https://www.youtube.com/watch?v=71pOiq-E_X4",
            "reduction": "Fu884hxjuhjEDSU0",
            "accesses": {
                "roles_id": [],
                "permissions_id": []
            },
            "properties": {
                "type": "redirect",
                "active": true,
                "number_of_uses": 1,
                "remained_number_of_uses": 0,
                "start_active": "0001-01-01T04:00:00+04:00",
                "end_active": "0001-01-01T04:00:00+04:00"
            }
        }
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
					ID        common_types.ID `json:"id"`
					Reduction string          `json:"reduction"`
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
					var cErr c_errors.RestAPI

					if request.Reduction != "" {
						if cErr = srv.controllers.UrlsManagement.RemoveByReduction(ctx.Context(), request.Reduction); cErr != nil {
							srv.components.Logger.Error().
								Format("Could not remove the short url data by reduction: '%s'. ", cErr).Write()

							if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
								srv.components.Logger.Error().
									Format("The error response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
							}

							return
						}
					} else {
						if cErr = srv.controllers.UrlsManagement.Remove(ctx.Context(), request.ID); cErr != nil {
							srv.components.Logger.Error().
								Format("Could not remove the short url data by id: '%s'. ", cErr).Write()

							if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
								srv.components.Logger.Error().
									Format("The error response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
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

						return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
					}

					return
				}
			}).Name(id)

			var route = srv.app.GetRoute(id)

			postmanGroup.AddItem(&postman.Items{
				Name: "Удаление короткого url. ",
				Description: `
Используется для удаления короткого url.
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
	"id": 0,
	"reduction: ""
}`,
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

		// POST /
		{
			var id = uuid.New().String()

			router.Post("/", func(ctx fiber.Ctx) (err error) {
				type Request struct {
					Source string `json:"source"`

					Properties *struct {
						Type types.ShortUrlType `json:"type"`

						NumberOfUses int64 `json:"number_of_uses"`

						StartActive time.Time `json:"start_active"`
						EndActive   time.Time `json:"end_active"`
					}
				}
				type Response struct {
					Url *models.ShortUrlInfo `json:"url" xml:"Url"`
				}

				var (
					request  = new(Request)
					response = new(Response)
				)

				// Чтение данных
				{
					if err = ctx.Bind().Body(request); err != nil {
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

				// Проверка данных
				{
					if request.Properties == nil {
						srv.components.Logger.Error().
							Text("Invalid arguments were received. ").Write()

						var cErr = common_errors.InvalidArguments()
						cErr.Details().Set("properties", "Is empty. ")

						if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(cErr)); err != nil {
							srv.components.Logger.Error().
								Format("The error response could not be recorded: '%s'. ", err).Write()

							return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
						}

						return
					}
				}

				// Обработка
				{
					var cErr c_errors.RestAPI

					if response.Url, cErr = srv.controllers.UrlsManagement.Create(ctx.Context(),
						request.Source,
						request.Properties.Type,
						request.Properties.NumberOfUses,
						request.Properties.StartActive,
						request.Properties.EndActive); cErr != nil {
						srv.components.Logger.Error().
							Format("Could not get the short url data by reduction: '%s'. ", cErr).Write()

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
				Name: "Создание короткого url. ",
				Description: `
Используется для создания короткого url.
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
	"source": "",
	"properties": {
		"type": "",
		"number_of_uses": 0,
		"start_active": "0001-01-01T00:00:00+00:00",
		"end_active": "0001-01-01T00:00:00+00:00"
	}
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
    "data": {
        "url": {
            "id": 10,
            "source": "https://translate.yandex.ru/",
            "reduction": "webAA4TxCJEFlS7s",
            "accesses": {
                "roles_id": [],
                "permissions_id": []
            },
            "properties": {
                "type": "redirect",
                "active": false,
                "number_of_uses": -1,
                "remained_number_of_uses": -1,
                "start_active": "0001-01-01T04:00:00+04:00",
                "end_active": "0001-01-01T04:00:00+04:00"
            }
        }
    }
}
`,
					},
				},
			})
		}

		// PUT /activate
		{
			var id = uuid.New().String()

			router.Put("/activate", func(ctx fiber.Ctx) (err error) {
				type Request struct {
					ID        common_types.ID `json:"id"`
					Reduction string          `json:"reduction"`
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
					var cErr c_errors.RestAPI

					if request.Reduction != "" {
						if cErr = srv.controllers.UrlsManagement.ActivateByReduction(ctx.Context(), request.Reduction); cErr != nil {
							srv.components.Logger.Error().
								Format("Could not activate the short url by reduction: '%s'. ", cErr).Write()

							if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
								srv.components.Logger.Error().
									Format("The error response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
							}

							return
						}
					} else {
						if cErr = srv.controllers.UrlsManagement.Activate(ctx.Context(), request.ID); cErr != nil {
							srv.components.Logger.Error().
								Format("Could not activate the short url by id: '%s'. ", cErr).Write()

							if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
								srv.components.Logger.Error().
									Format("The error response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
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

						return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
					}

					return
				}
			}).Name(id)

			var route = srv.app.GetRoute(id)

			postmanGroup.AddItem(&postman.Items{
				Name: "Активация короткого url. ",
				Description: `
Используется для активации короткого url.
`,
				Request: &postman.Request{
					URL: &postman.URL{
						Protocol: srv.conf.Postman.Protocol,
						Host:     strings.Split(srv.conf.Postman.Host, "."),
						Path:     strings.Split(route.Path, "/"),
					},
					Method: postman.Method(route.Method),
					Description: `
{
	"id": 0,
	"reduction": ""
}
`,
					Body: &postman.Body{
						Mode: "raw",
						Raw: `
{
	"id": 0,
	"reduction": ""
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

		// PUT /deactivate
		{
			var id = uuid.New().String()

			router.Put("/deactivate", func(ctx fiber.Ctx) (err error) {
				type Request struct {
					ID        common_types.ID `json:"id"`
					Reduction string          `json:"reduction"`
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
					var cErr c_errors.RestAPI

					if request.Reduction != "" {
						if cErr = srv.controllers.UrlsManagement.DeactivateByReduction(ctx.Context(), request.Reduction); cErr != nil {
							srv.components.Logger.Error().
								Format("Could not deactivate the short url by reduction: '%s'. ", cErr).Write()

							if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
								srv.components.Logger.Error().
									Format("The error response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
							}

							return
						}
					} else {
						if cErr = srv.controllers.UrlsManagement.Deactivate(ctx.Context(), request.ID); cErr != nil {
							srv.components.Logger.Error().
								Format("Could not deactivate the short url by id: '%s'. ", cErr).Write()

							if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
								srv.components.Logger.Error().
									Format("The error response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
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

						return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
					}

					return
				}
			}).Name(id)

			var route = srv.app.GetRoute(id)

			postmanGroup.AddItem(&postman.Items{
				Name: "Деактивация короткого url. ",
				Description: `
Используется для деактивации короткого url.
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
	"id": 0,
	"reduction": ""
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

		// GET /history
		{
			var (
				router       = router.Group("/history")
				postmanGroup = postmanGroup.AddItemGroup("История использований. ")
			)

			// GET /by_id/:id
			{
				var id = uuid.New().String()

				router.Get("/by_id/:id", func(ctx fiber.Ctx) (err error) {
					type Response struct {
						Count   int64                              `json:"count"   xml:"count,attr"`
						History []*models.ShortUrlUsageHistoryInfo `json:"history" xml:"History"`
					}
					type UriArgs struct {
						ID common_types.ID `uri:"id"`
					}
					type QueryArgs struct {
						Limit  *int64 `query:"limit"`
						Offset *int64 `query:"offset"`

						SortKey  string `query:"sort_key"`
						SortType string `query:"sort_type"`

						FilterStatus *string `query:"filter_status"`

						FilterTimestamp     *string `query:"filter_timestamp"`
						FilterTimestampType *string `query:"filter_timestamp_type"`
					}

					var (
						response  = new(Response)
						uriArgs   = new(UriArgs)
						queryArgs = new(QueryArgs)
					)

					// Чтение данных
					{
						if err = ctx.Bind().URI(uriArgs); err != nil {
							srv.components.Logger.Error().
								Format("The request uri data could not be read: '%s'. ", err).Write()

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

						if err = ctx.Bind().Query(queryArgs); err != nil {
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
							sort       *objects.ShortUrlsUsageHistoryListSort
							pagination *objects.ShortUrlsUsageHistoryListPagination
							filters    *objects.ShortUrlsUsageHistoryListFilters
						)

						// Обработка входных данных
						{
							// Сортировки
							{
								queryArgs.SortKey = strings.TrimSpace(queryArgs.SortKey)
								queryArgs.SortType = strings.TrimSpace(queryArgs.SortType)

								if queryArgs.SortKey != "" && queryArgs.SortType != "" {
									sort = &objects.ShortUrlsUsageHistoryListSort{
										Key:  queryArgs.SortKey,
										Type: queryArgs.SortType,
									}
								}
							}

							// Пагинация
							{
								pagination = &objects.ShortUrlsUsageHistoryListPagination{
									Offset: queryArgs.Offset,
									Limit:  queryArgs.Limit,
								}
							}

							// Фильтрация
							{
								filters = new(objects.ShortUrlsUsageHistoryListFilters)

								if queryArgs.FilterTimestampType != nil && *queryArgs.FilterTimestampType != "" {
									var v = types.ShortUrlUsageHistoryStatus(strings.TrimSpace(*queryArgs.FilterTimestampType))
									filters.Status = &v
								}

								if queryArgs.FilterTimestamp != nil && *queryArgs.FilterTimestamp != "" {
									var v = strings.TrimSpace(*queryArgs.FilterTimestamp)

									if tm, e := time.Parse(time.RFC3339Nano, v); e == nil {
										filters.Timestamp = &tm
										filters.TimestampType = queryArgs.FilterTimestampType
									}
								}
							}
						}

						var cErr c_errors.RestAPI

						if response.Count, response.History, cErr = srv.controllers.UrlsManagement.GetUsageHistory(ctx.Context(), uriArgs.ID, sort, pagination, filters); cErr != nil {
							srv.components.Logger.Error().
								Format("The short url usage history could not be retrieved: '%s'. ", cErr).Write()

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
					Name: "Получение историю использования короткого url по ID. ",
					Description: `
Используется для получения истории использования короткого url по ID.
`,
					Request: &postman.Request{
						URL: &postman.URL{
							Protocol: srv.conf.Postman.Protocol,
							Host:     strings.Split(srv.conf.Postman.Host, "."),
							Path:     strings.Split(route.Path, "/"),
							Query: []*postman.QueryParam{
								{
									Key:   "offset",
									Value: "0",
								},
								{
									Key:   "limit",
									Value: "20",
								},
								{
									Key:   "sort_key",
									Value: "",
								},
								{
									Key:   "sort_type",
									Value: "",
								},
								{
									Key:   "filter_status",
									Value: "",
								},
								{
									Key:   "filter_timestamp",
									Value: "",
								},
								{
									Key:   "filter_timestamp_type",
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
        "count": 14,
        "history": [
            {
                "status": "success",
                "timestamp": "2024-07-30T14:24:22.722266+04:00",
                "token_info": {
                    "ID": "2de0d61b-dafe-49ff-820a-88de302cf86e",
                    "ParentID": "00000000-0000-0000-0000-000000000000",
                    "UserID": 0,
                    "ProjectID": 0,
                    "Type": "session",
                    "Raw": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjI5Mzc2OTEsIm5iZiI6MTcyMjMzMjg5MSwiaWF0IjoxNzIyMzMyODkxLCJUb2tlbiI6eyJJRCI6IjJkZTBkNjFiLWRhZmUtNDlmZi04MjBhLTg4ZGUzMDJjZjg2ZSIsIlBhcmVudElEIjoiMDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAwIiwiVXNlcklEIjowLCJQcm9qZWN0SUQiOjAsIlBhcmFtcyI6eyJSZW1vdGVBZGRyIjoiMTcyLjIyLjAuMSIsIlVzZXJBZ2VudCI6Ik1vemlsbGEvNS4wIChXaW5kb3dzIE5UIDEwLjA7IFdpbjY0OyB4NjQpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS8xMjQuMC4wLjAgWWFCcm93c2VyLzI0LjYuMC4wIFNhZmFyaS81MzcuMzYifX0sIkxhbmd1YWdlIjoiIn0.u-7m54VOe9qyk22YjjoCg9DAjPc6j8waStCiRYE0s5XkWlrmKpFMVxSgXaduJ3WS5hrX_Yj_LOwez3IqgABERrMjgERvndUMoshRiTNoiVHosDVjuPvVQzYXlHom5fskR9lN-LiGL8gKXeJqCQw3GZ5NqrhzM0esYbO9q1XRrkjuO1Qs4SH2mkTYMF5iazIyRLDMOm5T9p9c13hQwDOFys58sHuSXUoD-yr0aA1WEe6AUHuOUNTLtuEQIs-o8xnFiTdQ9F5UsktcicH_5GKadI_USdNCMUqJRmRwXFyLeGGf6HV0xlRKIn8HZe6JlIZkSIpGfiM5wlwRxQ4nP4WZ2KQnZOT6D3Aoiuk8QxFhjXuqmGHFn6tjlbDDer2KJVMdeePEblIOQT-bKbiYRjVunoWZBbwTWc-E2TuGQectWkH2elqXMpT0bF9Kj8qqLIwkhMluFAuoD4gBTm-91ATKncUOH8fMOqGhFH1Y_XzirBNhd4dgr1h-AodPPyHGNUReHX_fPfGxeMEKUuVe58oubHOl7I4bMOZ7A52093BmJyzVFJbSXebXZqj6vkGwdx2buDx2sFCRsOHT5hy54M9oznOZciYepJ-GTOc1LlUhxbkWjRJYmIYFA8Jqgm-CGHPt_vkG0DSYppWsnX6ehvvuEur0kSQmYBKIiFgcQmGebt4",
                    "ExpiresAt": "2024-08-06T09:48:11Z",
                    "NotBefore": "2024-07-30T09:48:11Z",
                    "IssuedAt": "2024-07-30T09:48:11Z",
                    "Params": {
                        "RemoteAddr": "172.22.0.1",
                        "UserAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 YaBrowser/24.6.0.0 Safari/537.36"
                    },
                    "Language": ""
                }
            },
            {
                "status": "success",
                "timestamp": "2024-07-30T14:26:11.222774+04:00",
                "token_info": {
                    "ID": "7d7171ad-2814-449d-8132-c7aeeeb39788",
                    "ParentID": "00000000-0000-0000-0000-000000000000",
                    "UserID": 0,
                    "ProjectID": 0,
                    "Type": "session",
                    "Raw": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjI5Mzk4NjcsIm5iZiI6MTcyMjMzNTA2NywiaWF0IjoxNzIyMzM1MDY3LCJUb2tlbiI6eyJJRCI6IjdkNzE3MWFkLTI4MTQtNDQ5ZC04MTMyLWM3YWVlZWIzOTc4OCIsIlBhcmVudElEIjoiMDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAwIiwiVXNlcklEIjowLCJQcm9qZWN0SUQiOjAsIlBhcmFtcyI6eyJSZW1vdGVBZGRyIjoiMTcyLjIzLjAuMSIsIlVzZXJBZ2VudCI6Ik1vemlsbGEvNS4wIChXaW5kb3dzIE5UIDEwLjA7IFdpbjY0OyB4NjQpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS8xMjQuMC4wLjAgWWFCcm93c2VyLzI0LjYuMC4wIFNhZmFyaS81MzcuMzYifX0sIkxhbmd1YWdlIjoiIn0.HXQmQnLnpruFJaS9WVZzRQyRIpHArdHm6nWugdQUuV-dY_aDhEnycEQS97cNK5N_DnTixmYaLWDIwQToFBPXx8CE0Dlx2W6as3VpDVW3bTVx1UQzyGCEQX1rna88TyNDGB5QHFV0Mc4o9IL8VRdc_vXoonEq-7ZY2WCkPNdXPL1ypRqCjP3omWX53GEXZIec4UGQTadDnD11-T1eJTU_uEL3I9_bJr93oNSnV0iA65pKuE_tjpmW7P1zOul14rBcIV6qaLbzU-k7MEUSSnUDG2JvOaNmYHNlaC9SB6wPz9UjdeBlPQx0Yug1Fo9zg1dp0IAshmF6PpkzyQJI8N5BZ5yfKsTx2DUvyKSKmc56dJNr2z3D6GZDUutP6sQCDXlghmThWGsSm0eNnsDj0mw2TycOUMbjg6eThEyM5pLzXQlL6L06XvIhakSnJEy6qIbRBHGDy4a24tXi3Zxo0s50ZWieL_dqdjYRdPi2D42AClSeGeKMJMkFyKvnkF3R4UDwNsKFE9ecztHfkPghfuyLc4D980jxPxZenPhIzCO5O_SV_XoDeTNiu8I8ByTnuSa4myiitW1bnCihKyV037UtBQercLKtd8MWu_Y1v_9Jl7vyVBII5-HfGL7tmCaY5-pgYJCFzlRhBxBb49bZrw9MfAdBiymTcq09uHRP3rum8wA",
                    "ExpiresAt": "2024-08-06T10:24:27Z",
                    "NotBefore": "2024-07-30T10:24:27Z",
                    "IssuedAt": "2024-07-30T10:24:27Z",
                    "Params": {
                        "RemoteAddr": "172.23.0.1",
                        "UserAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 YaBrowser/24.6.0.0 Safari/537.36"
                    },
                    "Language": ""
                }
            },
            {
                "status": "success",
                "timestamp": "2024-07-30T14:27:12.406202+04:00",
                "token_info": {
                    "ID": "7d7171ad-2814-449d-8132-c7aeeeb39788",
                    "ParentID": "00000000-0000-0000-0000-000000000000",
                    "UserID": 0,
                    "ProjectID": 0,
                    "Type": "session",
                    "Raw": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjI5Mzk4NjcsIm5iZiI6MTcyMjMzNTA2NywiaWF0IjoxNzIyMzM1MDY3LCJUb2tlbiI6eyJJRCI6IjdkNzE3MWFkLTI4MTQtNDQ5ZC04MTMyLWM3YWVlZWIzOTc4OCIsIlBhcmVudElEIjoiMDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAwIiwiVXNlcklEIjowLCJQcm9qZWN0SUQiOjAsIlBhcmFtcyI6eyJSZW1vdGVBZGRyIjoiMTcyLjIzLjAuMSIsIlVzZXJBZ2VudCI6Ik1vemlsbGEvNS4wIChXaW5kb3dzIE5UIDEwLjA7IFdpbjY0OyB4NjQpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS8xMjQuMC4wLjAgWWFCcm93c2VyLzI0LjYuMC4wIFNhZmFyaS81MzcuMzYifX0sIkxhbmd1YWdlIjoiIn0.HXQmQnLnpruFJaS9WVZzRQyRIpHArdHm6nWugdQUuV-dY_aDhEnycEQS97cNK5N_DnTixmYaLWDIwQToFBPXx8CE0Dlx2W6as3VpDVW3bTVx1UQzyGCEQX1rna88TyNDGB5QHFV0Mc4o9IL8VRdc_vXoonEq-7ZY2WCkPNdXPL1ypRqCjP3omWX53GEXZIec4UGQTadDnD11-T1eJTU_uEL3I9_bJr93oNSnV0iA65pKuE_tjpmW7P1zOul14rBcIV6qaLbzU-k7MEUSSnUDG2JvOaNmYHNlaC9SB6wPz9UjdeBlPQx0Yug1Fo9zg1dp0IAshmF6PpkzyQJI8N5BZ5yfKsTx2DUvyKSKmc56dJNr2z3D6GZDUutP6sQCDXlghmThWGsSm0eNnsDj0mw2TycOUMbjg6eThEyM5pLzXQlL6L06XvIhakSnJEy6qIbRBHGDy4a24tXi3Zxo0s50ZWieL_dqdjYRdPi2D42AClSeGeKMJMkFyKvnkF3R4UDwNsKFE9ecztHfkPghfuyLc4D980jxPxZenPhIzCO5O_SV_XoDeTNiu8I8ByTnuSa4myiitW1bnCihKyV037UtBQercLKtd8MWu_Y1v_9Jl7vyVBII5-HfGL7tmCaY5-pgYJCFzlRhBxBb49bZrw9MfAdBiymTcq09uHRP3rum8wA",
                    "ExpiresAt": "2024-08-06T10:24:27Z",
                    "NotBefore": "2024-07-30T10:24:27Z",
                    "IssuedAt": "2024-07-30T10:24:27Z",
                    "Params": {
                        "RemoteAddr": "172.23.0.1",
                        "UserAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 YaBrowser/24.6.0.0 Safari/537.36"
                    },
                    "Language": ""
                }
            }
        ]
    }
}
`,
						},
					},
				})
			}

			// GET /by_reduction/:reduction
			{
				var id = uuid.New().String()

				router.Get("/by_reduction/:reduction", func(ctx fiber.Ctx) (err error) {
					type Response struct {
						Count   int64                              `json:"count"   xml:"count,attr"`
						History []*models.ShortUrlUsageHistoryInfo `json:"history" xml:"History"`
					}
					type UriArgs struct {
						Reduction string `uri:"reduction"`
					}
					type QueryArgs struct {
						Limit  *int64 `query:"limit"`
						Offset *int64 `query:"offset"`

						SortKey  string `query:"sort_key"`
						SortType string `query:"sort_type"`

						FilterStatus *string `query:"filter_status"`

						FilterTimestamp     *string `query:"filter_timestamp"`
						FilterTimestampType *string `query:"filter_timestamp_type"`
					}

					var (
						response  = new(Response)
						uriArgs   = new(UriArgs)
						queryArgs = new(QueryArgs)
					)

					// Чтение данных
					{
						if err = ctx.Bind().URI(uriArgs); err != nil {
							srv.components.Logger.Error().
								Format("The request uri data could not be read: '%s'. ", err).Write()

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

						if err = ctx.Bind().Query(queryArgs); err != nil {
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
							sort       *objects.ShortUrlsUsageHistoryListSort
							pagination *objects.ShortUrlsUsageHistoryListPagination
							filters    *objects.ShortUrlsUsageHistoryListFilters
						)

						// Обработка входных данных
						{
							// Сортировки
							{
								queryArgs.SortKey = strings.TrimSpace(queryArgs.SortKey)
								queryArgs.SortType = strings.TrimSpace(queryArgs.SortType)

								if queryArgs.SortKey != "" && queryArgs.SortType != "" {
									sort = &objects.ShortUrlsUsageHistoryListSort{
										Key:  queryArgs.SortKey,
										Type: queryArgs.SortType,
									}
								}
							}

							// Пагинация
							{
								pagination = &objects.ShortUrlsUsageHistoryListPagination{
									Offset: queryArgs.Offset,
									Limit:  queryArgs.Limit,
								}
							}

							// Фильтрация
							{
								filters = new(objects.ShortUrlsUsageHistoryListFilters)

								if queryArgs.FilterTimestampType != nil && *queryArgs.FilterTimestampType != "" {
									var v = types.ShortUrlUsageHistoryStatus(strings.TrimSpace(*queryArgs.FilterTimestampType))
									filters.Status = &v
								}

								if queryArgs.FilterTimestamp != nil && *queryArgs.FilterTimestamp != "" {
									var v = strings.TrimSpace(*queryArgs.FilterTimestamp)

									if tm, e := time.Parse(time.RFC3339Nano, v); e == nil {
										filters.Timestamp = &tm
										filters.TimestampType = queryArgs.FilterTimestampType
									}
								}
							}
						}

						var cErr c_errors.RestAPI

						if response.Count, response.History, cErr = srv.controllers.UrlsManagement.GetUsageHistoryByReduction(ctx.Context(), uriArgs.Reduction, sort, pagination, filters); cErr != nil {
							srv.components.Logger.Error().
								Format("The short url usage history could not be retrieved: '%s'. ", cErr).Write()

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
					Name: "Получение историю использования короткого url по сокращению. ",
					Description: `
Используется для получения истории использования короткого url по сокращению.
`,
					Request: &postman.Request{
						URL: &postman.URL{
							Protocol: srv.conf.Postman.Protocol,
							Host:     strings.Split(srv.conf.Postman.Host, "."),
							Path:     strings.Split(route.Path, "/"),
							Query: []*postman.QueryParam{
								{
									Key:   "offset",
									Value: "0",
								},
								{
									Key:   "limit",
									Value: "20",
								},
								{
									Key:   "sort_key",
									Value: "",
								},
								{
									Key:   "sort_type",
									Value: "",
								},
								{
									Key:   "filter_status",
									Value: "",
								},
								{
									Key:   "filter_timestamp",
									Value: "",
								},
								{
									Key:   "filter_timestamp_type",
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
        "count": 14,
        "history": [
            {
                "status": "success",
                "timestamp": "2024-07-30T14:24:22.722266+04:00",
                "token_info": {
                    "ID": "2de0d61b-dafe-49ff-820a-88de302cf86e",
                    "ParentID": "00000000-0000-0000-0000-000000000000",
                    "UserID": 0,
                    "ProjectID": 0,
                    "Type": "session",
                    "Raw": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjI5Mzc2OTEsIm5iZiI6MTcyMjMzMjg5MSwiaWF0IjoxNzIyMzMyODkxLCJUb2tlbiI6eyJJRCI6IjJkZTBkNjFiLWRhZmUtNDlmZi04MjBhLTg4ZGUzMDJjZjg2ZSIsIlBhcmVudElEIjoiMDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAwIiwiVXNlcklEIjowLCJQcm9qZWN0SUQiOjAsIlBhcmFtcyI6eyJSZW1vdGVBZGRyIjoiMTcyLjIyLjAuMSIsIlVzZXJBZ2VudCI6Ik1vemlsbGEvNS4wIChXaW5kb3dzIE5UIDEwLjA7IFdpbjY0OyB4NjQpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS8xMjQuMC4wLjAgWWFCcm93c2VyLzI0LjYuMC4wIFNhZmFyaS81MzcuMzYifX0sIkxhbmd1YWdlIjoiIn0.u-7m54VOe9qyk22YjjoCg9DAjPc6j8waStCiRYE0s5XkWlrmKpFMVxSgXaduJ3WS5hrX_Yj_LOwez3IqgABERrMjgERvndUMoshRiTNoiVHosDVjuPvVQzYXlHom5fskR9lN-LiGL8gKXeJqCQw3GZ5NqrhzM0esYbO9q1XRrkjuO1Qs4SH2mkTYMF5iazIyRLDMOm5T9p9c13hQwDOFys58sHuSXUoD-yr0aA1WEe6AUHuOUNTLtuEQIs-o8xnFiTdQ9F5UsktcicH_5GKadI_USdNCMUqJRmRwXFyLeGGf6HV0xlRKIn8HZe6JlIZkSIpGfiM5wlwRxQ4nP4WZ2KQnZOT6D3Aoiuk8QxFhjXuqmGHFn6tjlbDDer2KJVMdeePEblIOQT-bKbiYRjVunoWZBbwTWc-E2TuGQectWkH2elqXMpT0bF9Kj8qqLIwkhMluFAuoD4gBTm-91ATKncUOH8fMOqGhFH1Y_XzirBNhd4dgr1h-AodPPyHGNUReHX_fPfGxeMEKUuVe58oubHOl7I4bMOZ7A52093BmJyzVFJbSXebXZqj6vkGwdx2buDx2sFCRsOHT5hy54M9oznOZciYepJ-GTOc1LlUhxbkWjRJYmIYFA8Jqgm-CGHPt_vkG0DSYppWsnX6ehvvuEur0kSQmYBKIiFgcQmGebt4",
                    "ExpiresAt": "2024-08-06T09:48:11Z",
                    "NotBefore": "2024-07-30T09:48:11Z",
                    "IssuedAt": "2024-07-30T09:48:11Z",
                    "Params": {
                        "RemoteAddr": "172.22.0.1",
                        "UserAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 YaBrowser/24.6.0.0 Safari/537.36"
                    },
                    "Language": ""
                }
            },
            {
                "status": "success",
                "timestamp": "2024-07-30T14:26:11.222774+04:00",
                "token_info": {
                    "ID": "7d7171ad-2814-449d-8132-c7aeeeb39788",
                    "ParentID": "00000000-0000-0000-0000-000000000000",
                    "UserID": 0,
                    "ProjectID": 0,
                    "Type": "session",
                    "Raw": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjI5Mzk4NjcsIm5iZiI6MTcyMjMzNTA2NywiaWF0IjoxNzIyMzM1MDY3LCJUb2tlbiI6eyJJRCI6IjdkNzE3MWFkLTI4MTQtNDQ5ZC04MTMyLWM3YWVlZWIzOTc4OCIsIlBhcmVudElEIjoiMDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAwIiwiVXNlcklEIjowLCJQcm9qZWN0SUQiOjAsIlBhcmFtcyI6eyJSZW1vdGVBZGRyIjoiMTcyLjIzLjAuMSIsIlVzZXJBZ2VudCI6Ik1vemlsbGEvNS4wIChXaW5kb3dzIE5UIDEwLjA7IFdpbjY0OyB4NjQpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS8xMjQuMC4wLjAgWWFCcm93c2VyLzI0LjYuMC4wIFNhZmFyaS81MzcuMzYifX0sIkxhbmd1YWdlIjoiIn0.HXQmQnLnpruFJaS9WVZzRQyRIpHArdHm6nWugdQUuV-dY_aDhEnycEQS97cNK5N_DnTixmYaLWDIwQToFBPXx8CE0Dlx2W6as3VpDVW3bTVx1UQzyGCEQX1rna88TyNDGB5QHFV0Mc4o9IL8VRdc_vXoonEq-7ZY2WCkPNdXPL1ypRqCjP3omWX53GEXZIec4UGQTadDnD11-T1eJTU_uEL3I9_bJr93oNSnV0iA65pKuE_tjpmW7P1zOul14rBcIV6qaLbzU-k7MEUSSnUDG2JvOaNmYHNlaC9SB6wPz9UjdeBlPQx0Yug1Fo9zg1dp0IAshmF6PpkzyQJI8N5BZ5yfKsTx2DUvyKSKmc56dJNr2z3D6GZDUutP6sQCDXlghmThWGsSm0eNnsDj0mw2TycOUMbjg6eThEyM5pLzXQlL6L06XvIhakSnJEy6qIbRBHGDy4a24tXi3Zxo0s50ZWieL_dqdjYRdPi2D42AClSeGeKMJMkFyKvnkF3R4UDwNsKFE9ecztHfkPghfuyLc4D980jxPxZenPhIzCO5O_SV_XoDeTNiu8I8ByTnuSa4myiitW1bnCihKyV037UtBQercLKtd8MWu_Y1v_9Jl7vyVBII5-HfGL7tmCaY5-pgYJCFzlRhBxBb49bZrw9MfAdBiymTcq09uHRP3rum8wA",
                    "ExpiresAt": "2024-08-06T10:24:27Z",
                    "NotBefore": "2024-07-30T10:24:27Z",
                    "IssuedAt": "2024-07-30T10:24:27Z",
                    "Params": {
                        "RemoteAddr": "172.23.0.1",
                        "UserAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 YaBrowser/24.6.0.0 Safari/537.36"
                    },
                    "Language": ""
                }
            },
            {
                "status": "success",
                "timestamp": "2024-07-30T14:27:12.406202+04:00",
                "token_info": {
                    "ID": "7d7171ad-2814-449d-8132-c7aeeeb39788",
                    "ParentID": "00000000-0000-0000-0000-000000000000",
                    "UserID": 0,
                    "ProjectID": 0,
                    "Type": "session",
                    "Raw": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjI5Mzk4NjcsIm5iZiI6MTcyMjMzNTA2NywiaWF0IjoxNzIyMzM1MDY3LCJUb2tlbiI6eyJJRCI6IjdkNzE3MWFkLTI4MTQtNDQ5ZC04MTMyLWM3YWVlZWIzOTc4OCIsIlBhcmVudElEIjoiMDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAwIiwiVXNlcklEIjowLCJQcm9qZWN0SUQiOjAsIlBhcmFtcyI6eyJSZW1vdGVBZGRyIjoiMTcyLjIzLjAuMSIsIlVzZXJBZ2VudCI6Ik1vemlsbGEvNS4wIChXaW5kb3dzIE5UIDEwLjA7IFdpbjY0OyB4NjQpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS8xMjQuMC4wLjAgWWFCcm93c2VyLzI0LjYuMC4wIFNhZmFyaS81MzcuMzYifX0sIkxhbmd1YWdlIjoiIn0.HXQmQnLnpruFJaS9WVZzRQyRIpHArdHm6nWugdQUuV-dY_aDhEnycEQS97cNK5N_DnTixmYaLWDIwQToFBPXx8CE0Dlx2W6as3VpDVW3bTVx1UQzyGCEQX1rna88TyNDGB5QHFV0Mc4o9IL8VRdc_vXoonEq-7ZY2WCkPNdXPL1ypRqCjP3omWX53GEXZIec4UGQTadDnD11-T1eJTU_uEL3I9_bJr93oNSnV0iA65pKuE_tjpmW7P1zOul14rBcIV6qaLbzU-k7MEUSSnUDG2JvOaNmYHNlaC9SB6wPz9UjdeBlPQx0Yug1Fo9zg1dp0IAshmF6PpkzyQJI8N5BZ5yfKsTx2DUvyKSKmc56dJNr2z3D6GZDUutP6sQCDXlghmThWGsSm0eNnsDj0mw2TycOUMbjg6eThEyM5pLzXQlL6L06XvIhakSnJEy6qIbRBHGDy4a24tXi3Zxo0s50ZWieL_dqdjYRdPi2D42AClSeGeKMJMkFyKvnkF3R4UDwNsKFE9ecztHfkPghfuyLc4D980jxPxZenPhIzCO5O_SV_XoDeTNiu8I8ByTnuSa4myiitW1bnCihKyV037UtBQercLKtd8MWu_Y1v_9Jl7vyVBII5-HfGL7tmCaY5-pgYJCFzlRhBxBb49bZrw9MfAdBiymTcq09uHRP3rum8wA",
                    "ExpiresAt": "2024-08-06T10:24:27Z",
                    "NotBefore": "2024-07-30T10:24:27Z",
                    "IssuedAt": "2024-07-30T10:24:27Z",
                    "Params": {
                        "RemoteAddr": "172.23.0.1",
                        "UserAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 YaBrowser/24.6.0.0 Safari/537.36"
                    },
                    "Language": ""
                }
            }
        ]
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
