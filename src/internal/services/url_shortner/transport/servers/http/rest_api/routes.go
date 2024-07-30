package http_rest_api

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	common_errors "sm-box/internal/common/errors"
	common_types "sm-box/internal/common/types"
	authentication_entities "sm-box/internal/services/authentication/objects/entities"
	"sm-box/internal/services/url_shortner/objects"
	srv_errors "sm-box/internal/services/url_shortner/objects/errors"
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
				status    types.ShortUrlUsageHistoryStatus
				url       *models.ShortUrlInfo
				token     *authentication_entities.JwtSessionToken
			)

			// Запись в историю
			{
				defer func() {
					go func() {
						if url != nil {
							var cErr c_errors.RestAPI

							if cErr = srv.controllers.Urls.WriteCallToHistory(ctx.Context(), url.ID, status, token); cErr != nil {
								srv.components.Logger.Warn().
									Format("The call data could not be recorded in the history: '%s'. ", cErr).
									Field("url", url).
									Field("status", status).Write()
							}
						}
					}()
				}()
			}

			// Получение токена
			{
				if raw := ctx.Cookies(srv.conf.Components.AccessSystem.CookieKeyForSessionToken); len(raw) > 0 {
					token = new(authentication_entities.JwtSessionToken)

					if err = token.Parse(raw); err != nil {
						srv.components.Logger.Error().
							Format("Failed to get session token data: '%s'. ", err).
							Field("raw", raw).Write()

						status = types.ShortUrlUsageHistoryStatusFailed

						err = ctx.Redirect().To("/errors/403")
						return
					}
				}
			}

			// Получение сокращения
			{
				reduction = strings.Replace(string(ctx.Request().URI().Path()), strings.Replace(route.Path, "/*", "/", 1), "", 1)
			}

			// Получение и обработка короткого url
			{
				var cErr c_errors.RestAPI

				// Получение
				{
					if url, cErr = srv.controllers.Urls.GetByReductionFromRedisDB(ctx.Context(), reduction); cErr != nil {
						srv.components.Logger.Warn().
							Format("Failed to get information on a short url: '%s'. ", cErr).Write()

						if errors.Is(cErr, srv_errors.ShortUrlNotFound()) {
							status = types.ShortUrlUsageHistoryStatusForbidden

							err = ctx.Redirect().To("/errors/403")
							return
						}

						status = types.ShortUrlUsageHistoryStatusFailed

						err = ctx.Redirect().To("/errors/50x")
						return
					}
				}

				// Проверки
				{
					var (
						tm      = time.Now()
						emptyTm time.Time
					)

					// Ещё не начал действовать
					{
						if tm.Before(url.Properties.StartActive) && !url.Properties.StartActive.Equal(emptyTm) {
							url = nil

							srv.components.Logger.Warn().
								Text("The validity period of the short url has not yet begun. ").Write()

							err = ctx.Redirect().To("/errors/403")
							return
						}
					}

					// Уже закончился
					{
						if tm.After(url.Properties.EndActive) && !url.Properties.EndActive.Equal(emptyTm) {
							srv.components.Logger.Warn().
								Text("The validity period of the short url has already been completed. ").Write()

							// Удаление из базы данных Redis
							{
								if cErr = srv.controllers.Urls.RemoveByReductionFromRedisDB(ctx.Context(), url.Reduction); cErr != nil {
									srv.components.Logger.Warn().
										Format("The short url could not be deleted from the redis database: '%s'. ", cErr).
										Field("url", url).Write()
								}
							}

							url = nil

							err = ctx.Redirect().To("/errors/403")

							return
						}
					}

					// Кол-во использований превышено
					{
						if url.Properties.NumberOfUses == 0 {
							srv.components.Logger.Warn().
								Text("The number of uses of the short url  is overestimated. ").Write()

							// Удаление из базы данных Redis
							{
								if cErr = srv.controllers.Urls.RemoveByReductionFromRedisDB(ctx.Context(), url.Reduction); cErr != nil {
									srv.components.Logger.Warn().
										Format("The short url could not be deleted from the redis database: '%s'. ", cErr).
										Field("url", url).Write()
								}
							}

							url = nil

							err = ctx.Redirect().To("/errors/403")

							return
						}
					}

					// Доступов пользователя
					{
						fmt.Printf("%+v\n", url)
					}
				}
			}

			// Выполнение инструкций
			{
				// Обработка
				{
					switch url.Properties.Type {
					case types.ShortUrlTypeRedirect:
						{
							if err = ctx.Redirect().To(url.Source); err != nil {
								srv.components.Logger.Warn().
									Format("Failed to redirect a remote resource: '%s'. ", err).
									Field("url", url).Write()

								status = types.ShortUrlUsageHistoryStatusFailed

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

								status = types.ShortUrlUsageHistoryStatusFailed

								err = ctx.Redirect().To("/errors/50x")
								return
							}
						}
					default:
						{
							srv.components.Logger.Error().
								Text("Unknown type of shortened url. ").
								Field("url", url).Write()

							status = types.ShortUrlUsageHistoryStatusForbidden

							err = ctx.Redirect().To("/errors/403")
							return
						}
					}

					status = types.ShortUrlUsageHistoryStatusSuccess
				}

				// Обновление данных в базе если кол-во использований не бесконечное
				{
					if url.Properties.NumberOfUses > 0 {
						var cErr c_errors.RestAPI

						url.Properties.RemainedNumberOfUses--

						if url.Properties.RemainedNumberOfUses == 0 {
							if cErr = srv.controllers.Urls.RemoveByReductionFromRedisDB(ctx.Context(), url.Reduction); cErr != nil {
								srv.components.Logger.Warn().
									Format("The short url could not be deleted from the redis database: '%s'. ", cErr).
									Field("url", url).Write()

								return
							}
						} else {
							if cErr = srv.controllers.Urls.UpdateInRedisDB(ctx.Context(), url); cErr != nil {
								srv.components.Logger.Warn().
									Format("Failed to update the short url data in the redis database: '%s'. ", cErr).
									Field("url", url).Write()

								return
							}
						}
					}
				}

				return
			}
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
					List []*models.ShortUrlInfo `json:"list" xml:"List"`
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

					if response.List, cErr = srv.controllers.UrlsManagement.GetList(ctx.Context(), search, sort, pagination, filters); cErr != nil {
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
				},
			})
		}

		// GET /by_reduce/:reduction
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
				},
			})
		}

		// GET /history
		{
			var id = uuid.New().String()

			router.Get("/history", func(ctx fiber.Ctx) (err error) {
				type Response struct {
					History []*models.ShortUrlUsageHistoryInfo `json:"history" xml:"History"`
				}

				var (
					response = new(Response)
				)

				// Обработка
				{
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
				Name: "Получение историю использования короткого url. ",
				Description: `
Используется для получения истории использования короткого url.
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

	srv.components.Logger.Info().
		Text("Http rest api server routes are initialized. ").Write()

	return nil
}
