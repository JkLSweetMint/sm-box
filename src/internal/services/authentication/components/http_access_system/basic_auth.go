package http_access_system

import (
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	app_errors "sm-box/internal/app/objects/errors"
	app_models "sm-box/internal/app/objects/models"
	common_errors "sm-box/internal/common/errors"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/authentication/objects/entities"
	srv_errors "sm-box/internal/services/authentication/objects/errors"
	"sm-box/internal/services/users/objects"
	users_models "sm-box/internal/services/users/objects/models"
	c_errors "sm-box/pkg/errors"
	http_rest_api_io "sm-box/pkg/http/rest_api/io"
	"time"
)

// BasicAuthentication - промежуточное программное обеспечение для аутентификации пользователя по http маршрутам.
func (acc *accessSystem) BasicAuthentication(ctx fiber.Ctx) (err error) {
	var (
		sessionToken *entities.JwtSessionToken
		accessToken  *entities.JwtAccessToken
		refreshToken *entities.JwtRefreshToken
		route        *entities.HttpRoute
	)

	var (
		remoteAddr = string(ctx.Request().Header.Peek("X-Real-IP"))
		userAgent  = string(ctx.Request().Header.UserAgent())
	)

	// Работа с токенами
	{
		// Доступа
		{
			var cErr c_errors.Error

			if accessToken, cErr = acc.basicAuthenticationProcessingAccessToken(ctx); cErr != nil {
				acc.components.Logger.Warn().
					Format("Access token processing failed: '%s'. ", cErr).Write()
			}

			// Проверка что токен не спиздили
			{
				if (accessToken != nil && accessToken.Params != nil) && (remoteAddr != accessToken.Params.RemoteAddr ||
					userAgent != accessToken.Params.UserAgent) {

					if raw := ctx.Cookies(acc.conf.CookieKeyForSessionToken); raw != "" {
						ctx.Cookie(&fiber.Cookie{
							Name:        acc.conf.CookieKeyForSessionToken,
							Value:       "",
							Path:        "/",
							Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
							MaxAge:      0,
							Expires:     time.Unix(0, 0),
							Secure:      false,
							HTTPOnly:    false,
							SameSite:    fiber.CookieSameSiteNoneMode,
							SessionOnly: false,
						})
					}

					if raw := ctx.Cookies(acc.conf.CookieKeyForAccessToken); raw != "" {
						ctx.Cookie(&fiber.Cookie{
							Name:        acc.conf.CookieKeyForAccessToken,
							Value:       "",
							Path:        "/",
							Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
							MaxAge:      0,
							Expires:     time.Unix(0, 0),
							Secure:      false,
							HTTPOnly:    false,
							SameSite:    fiber.CookieSameSiteNoneMode,
							SessionOnly: false,
						})
					}

					if raw := ctx.Cookies(acc.conf.CookieKeyForRefreshToken); raw != "" {
						ctx.Cookie(&fiber.Cookie{
							Name:        acc.conf.CookieKeyForRefreshToken,
							Value:       "",
							Path:        "/",
							Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
							MaxAge:      0,
							Expires:     time.Unix(0, 0),
							Secure:      false,
							HTTPOnly:    false,
							SameSite:    fiber.CookieSameSiteNoneMode,
							SessionOnly: false,
						})
					}
				}
			}
		}

		// Сессия
		{
			var cErr c_errors.Error

			if sessionToken, cErr = acc.basicAuthenticationProcessingSessionToken(ctx, accessToken); cErr != nil {
				if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(cErr)); err != nil {
					acc.components.Logger.Error().
						Format("The response could not be recorded: '%s'. ", err).Write()

					var cErr = common_errors.ResponseCouldNotBeRecorded_RestAPI()
					cErr.SetError(err)

					if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
						acc.components.Logger.Error().
							Format("The error response could not be recorded: '%s'. ", err).Write()

						return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
					}

					return
				}
				return
			}

			// Проверка что токен не спиздили
			{
				if (sessionToken != nil && sessionToken.Params != nil) && (remoteAddr != sessionToken.Params.RemoteAddr ||
					userAgent != sessionToken.Params.UserAgent) {

					if raw := ctx.Cookies(acc.conf.CookieKeyForSessionToken); raw != "" {
						ctx.Cookie(&fiber.Cookie{
							Name:        acc.conf.CookieKeyForSessionToken,
							Value:       "",
							Path:        "/",
							Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
							MaxAge:      0,
							Expires:     time.Unix(0, 0),
							Secure:      false,
							HTTPOnly:    false,
							SameSite:    fiber.CookieSameSiteNoneMode,
							SessionOnly: false,
						})
					}

					if raw := ctx.Cookies(acc.conf.CookieKeyForAccessToken); raw != "" {
						ctx.Cookie(&fiber.Cookie{
							Name:        acc.conf.CookieKeyForAccessToken,
							Value:       "",
							Path:        "/",
							Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
							MaxAge:      0,
							Expires:     time.Unix(0, 0),
							Secure:      false,
							HTTPOnly:    false,
							SameSite:    fiber.CookieSameSiteNoneMode,
							SessionOnly: false,
						})
					}

					if raw := ctx.Cookies(acc.conf.CookieKeyForRefreshToken); raw != "" {
						ctx.Cookie(&fiber.Cookie{
							Name:        acc.conf.CookieKeyForRefreshToken,
							Value:       "",
							Path:        "/",
							Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
							MaxAge:      0,
							Expires:     time.Unix(0, 0),
							Secure:      false,
							HTTPOnly:    false,
							SameSite:    fiber.CookieSameSiteNoneMode,
							SessionOnly: false,
						})
					}
				}
			}
		}

		// Обновления (отрабатывает если нет токена доступа)
		{
			if accessToken == nil {
				// Получение
				{
					var cErr c_errors.Error

					if refreshToken, cErr = acc.basicAuthenticationProcessingRefreshToken(ctx); cErr != nil {
						acc.components.Logger.Warn().
							Format("Refresh token processing failed: '%s'. ", cErr).Write()
					}

					// Проверка что токен не спиздили
					{
						if (refreshToken != nil && refreshToken.Params != nil) && (remoteAddr != refreshToken.Params.RemoteAddr ||
							userAgent != refreshToken.Params.UserAgent) {

							if raw := ctx.Cookies(acc.conf.CookieKeyForSessionToken); raw != "" {
								ctx.Cookie(&fiber.Cookie{
									Name:        acc.conf.CookieKeyForSessionToken,
									Value:       "",
									Path:        "/",
									Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
									MaxAge:      0,
									Expires:     time.Unix(0, 0),
									Secure:      false,
									HTTPOnly:    false,
									SameSite:    fiber.CookieSameSiteNoneMode,
									SessionOnly: false,
								})
							}

							if raw := ctx.Cookies(acc.conf.CookieKeyForAccessToken); raw != "" {
								ctx.Cookie(&fiber.Cookie{
									Name:        acc.conf.CookieKeyForAccessToken,
									Value:       "",
									Path:        "/",
									Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
									MaxAge:      0,
									Expires:     time.Unix(0, 0),
									Secure:      false,
									HTTPOnly:    false,
									SameSite:    fiber.CookieSameSiteNoneMode,
									SessionOnly: false,
								})
							}

							if raw := ctx.Cookies(acc.conf.CookieKeyForRefreshToken); raw != "" {
								ctx.Cookie(&fiber.Cookie{
									Name:        acc.conf.CookieKeyForRefreshToken,
									Value:       "",
									Path:        "/",
									Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
									MaxAge:      0,
									Expires:     time.Unix(0, 0),
									Secure:      false,
									HTTPOnly:    false,
									SameSite:    fiber.CookieSameSiteNoneMode,
									SessionOnly: false,
								})
							}
						}
					}
				}

				if refreshToken != nil {
					var (
						user    *users_models.UserInfo
						project *app_models.ProjectInfo
					)

					// Получение данных пользователя
					{
						var cErr c_errors.Error

						if user, cErr = acc.gateways.Users.GetOne(ctx.Context(), sessionToken.UserID); cErr != nil {
							acc.components.Logger.Error().
								Format("Failed to get the user data: '%s'. ", cErr).
								Field("id", sessionToken.UserID).Write()

							if errors.Is(cErr, sql.ErrNoRows) {
								cErr = srv_errors.NotAccess()
							} else {
								cErr = common_errors.InternalServerError()
							}

							if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(cErr)); err != nil {
								acc.components.Logger.Error().
									Format("The error response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
							}

							return
						}

						acc.components.Logger.Info().
							Text("The user's data has been successfully received. ").
							Field("user", user).Write()
					}

					// Получение данных проекта
					{
						var cErr c_errors.Error

						if project, cErr = acc.gateways.Projects.GetOne(ctx.Context(), sessionToken.ProjectID); cErr != nil {
							acc.components.Logger.Error().
								Format("Failed to get the project: '%s'. ", cErr).
								Field("id", sessionToken.ProjectID).Write()

							if errors.Is(cErr, sql.ErrNoRows) {
								cErr = srv_errors.NotAccess()
							} else {
								cErr = common_errors.InternalServerError()
							}

							if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(cErr)); err != nil {
								acc.components.Logger.Error().
									Format("The error response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
							}

							return
						}

						acc.components.Logger.Info().
							Text("The project data has been successfully received. ").
							Field("project", project).Write()
					}

					// Проверка доступа
					{
						var (
							ids  = make(map[common_types.ID]struct{})
							cErr c_errors.Error
						)

						// Список id проектов
						{
							var writeInheritance func(rl *users_models.RoleInfo)

							writeInheritance = func(rl *users_models.RoleInfo) {
								if id := rl.ProjectID; id != 0 {
									ids[id] = struct{}{}
								}

								for _, child := range rl.Inheritances {
									writeInheritance(child.RoleInfo)
								}
							}

							for _, rl := range user.Accesses.Roles {
								writeInheritance(rl)
							}
						}

						if _, ok := ids[project.ID]; !ok {
							acc.components.Logger.Error().
								Format("The user does not have access to the project: '%s'. ", cErr).
								Field("project_id", project.ID).
								Field("user_id", user.ID).Write()

							if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(app_errors.NotAccessToProject())); err != nil {
								acc.components.Logger.Error().
									Format("The error response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
							}

							return
						}
					}

					// Создание новых токенов
					{
						// Создание токена обновления
						{
							refreshToken = &entities.JwtRefreshToken{
								JwtToken: &entities.JwtToken{
									ProjectID: sessionToken.ProjectID,
									ParentID:  refreshToken.ID,
									UserID:    user.ID,

									Params: sessionToken.Params,
								},
							}

							if err = refreshToken.Generate(); err != nil {
								acc.components.Logger.Error().
									Format("User session token generation failed: '%s'. ", err).Write()

								if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(common_errors.InternalServerError())); err != nil {
									acc.components.Logger.Error().
										Format("The error response could not be recorded: '%s'. ", err).Write()

									return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
								}

								return
							}

							// Сохранение в базе
							{
								if err = acc.repositories.JwtTokens.RegisterJwtRefreshToken(ctx.Context(), refreshToken); err != nil {
									acc.components.Logger.Error().
										Format("The client's refresh token could not be registered in the database: '%s'. ", err).Write()

									if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(common_errors.InternalServerError())); err != nil {
										acc.components.Logger.Error().
											Format("The error response could not be recorded: '%s'. ", err).Write()

										return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
									}

									return
								}
							}
						}

						// Создание токена доступа
						{
							accessToken = &entities.JwtAccessToken{
								JwtToken: &entities.JwtToken{
									ParentID:  refreshToken.ID,
									ProjectID: sessionToken.ProjectID,
									UserID:    user.ID,

									Params: sessionToken.Params,
								},
							}

							// Запись доступов пользователя
							{
								accessToken.UserInfo = &entities.JwtAccessTokenUserInfo{
									Accesses: user.Accesses,
								}
							}

							if err = accessToken.Generate(); err != nil {
								acc.components.Logger.Error().
									Format("User session token generation failed: '%s'. ", err).Write()

								if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(common_errors.InternalServerError())); err != nil {
									acc.components.Logger.Error().
										Format("The error response could not be recorded: '%s'. ", err).Write()

									return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
								}

								return
							}

							// Сохранение в базе
							{
								if err = acc.repositories.JwtTokens.RegisterJwtAccessToken(ctx.Context(), accessToken); err != nil {
									acc.components.Logger.Error().
										Format("The client's access token could not be registered in the database: '%s'. ", err).Write()

									if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(common_errors.InternalServerError())); err != nil {
										acc.components.Logger.Error().
											Format("The error response could not be recorded: '%s'. ", err).Write()

										return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
									}

									return
								}
							}
						}
					}

					// Запись печенек
					{
						ctx.Cookie(&fiber.Cookie{
							Name:        acc.conf.CookieKeyForAccessToken,
							Value:       accessToken.Raw,
							Path:        "/",
							Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
							MaxAge:      0,
							Expires:     accessToken.ExpiresAt,
							Secure:      true,
							HTTPOnly:    true,
							SameSite:    fiber.CookieSameSiteLaxMode,
							SessionOnly: false,
						})

						ctx.Cookie(&fiber.Cookie{
							Name:        acc.conf.CookieKeyForRefreshToken,
							Value:       refreshToken.Raw,
							Path:        "/",
							Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
							MaxAge:      0,
							Expires:     refreshToken.ExpiresAt,
							Secure:      true,
							HTTPOnly:    true,
							SameSite:    fiber.CookieSameSiteLaxMode,
							SessionOnly: false,
						})
					}
				}
			}

			// Пересоздание токена сессии чтоб сбросить данные пользователя в нём
			{
				if accessToken == nil {
					refreshToken = nil

					if sessionToken.UserID != 0 && sessionToken.ProjectID != 0 {
						// Создание нового токена сессии
						{
							sessionToken = &entities.JwtSessionToken{
								JwtToken: &entities.JwtToken{
									ParentID: sessionToken.ID,

									Params: sessionToken.Params,
								},
							}

							if err = sessionToken.Generate(); err != nil {
								acc.components.Logger.Error().
									Format("User token generation failed: '%s'. ", err).Write()

								if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(common_errors.InternalServerError())); err != nil {
									acc.components.Logger.Error().
										Format("The response could not be recorded: '%s'. ", err).Write()

									var cErr = common_errors.ResponseCouldNotBeRecorded_RestAPI()
									cErr.SetError(err)

									if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
										acc.components.Logger.Error().
											Format("The error response could not be recorded: '%s'. ", err).Write()

										return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
									}

									return
								}
								return
							}
						}

						// Печеньки
						{
							ctx.Cookie(&fiber.Cookie{
								Name:        acc.conf.CookieKeyForSessionToken,
								Value:       sessionToken.Raw,
								Path:        "/",
								Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
								MaxAge:      0,
								Expires:     sessionToken.ExpiresAt,
								Secure:      false,
								HTTPOnly:    true,
								SameSite:    fiber.CookieSameSiteLaxMode,
								SessionOnly: false,
							})
						}
					}

					if raw := ctx.Cookies(acc.conf.CookieKeyForAccessToken); len(raw) > 0 {
						ctx.Cookie(&fiber.Cookie{
							Name:        acc.conf.CookieKeyForAccessToken,
							Value:       "",
							Path:        "/",
							Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
							MaxAge:      0,
							Expires:     time.Unix(0, 0),
							Secure:      false,
							HTTPOnly:    false,
							SameSite:    fiber.CookieSameSiteNoneMode,
							SessionOnly: false,
						})
					}

					if raw := ctx.Cookies(acc.conf.CookieKeyForRefreshToken); len(raw) > 0 {
						ctx.Cookie(&fiber.Cookie{
							Name:        acc.conf.CookieKeyForRefreshToken,
							Value:       "",
							Path:        "/",
							Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
							MaxAge:      0,
							Expires:     time.Unix(0, 0),
							Secure:      false,
							HTTPOnly:    false,
							SameSite:    fiber.CookieSameSiteNoneMode,
							SessionOnly: false,
						})
					}
				}
			}
		}
	}

	// Получение маршрута
	{
		var (
			protocol = string(ctx.Request().Header.Peek("X-Forwarded-Protocol"))
			method   = string(ctx.Request().Header.Peek("X-Forwarded-Method"))
			path     = string(ctx.Request().Header.Peek("X-Forwarded-Path"))
		)

		if string(ctx.Request().Header.Peek("Upgrade")) == "websocket" {
			switch protocol {
			case "http":
				protocol = "ws"
			case "https":
				protocol = "wss"
			}
		}

		if route, err = acc.repositories.HttpRoutesRedis.Get(ctx.Context(), protocol, method, path); err != nil {
			acc.components.Logger.Error().
				Format("Failed to get the route data: '%s'. ", err).
				Field("protocol", protocol).
				Field("method", method).
				Field("path", path).Write()

			if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(common_errors.InternalServerError())); err != nil {
				acc.components.Logger.Error().
					Format("The error response could not be recorded: '%s'. ", err).Write()

				return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
			}

			return
		}

		if route == nil || !route.Active {
			return http_rest_api_io.WriteError(ctx, common_errors.RouteNotFound_RestAPI())
		}

		// Проверка доступа
		{
			var allowed bool

			// Обработка если токен доступа есть
			{
				if accessToken != nil {
					if accessToken.UserInfo == nil || accessToken.UserInfo.Accesses == nil {
						return http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(srv_errors.NotAccess()))
					}

					// Роли
					{
						if accessToken.UserInfo.Accesses.Roles != nil {
						CheckAccessForRouteByRole:
							for _, userRole := range accessToken.UserInfo.Accesses.Roles {
								for _, routeRoleID := range route.Accesses.Roles {
									if userRole.ID == routeRoleID {
										allowed = true
										break CheckAccessForRouteByRole
									}
								}
							}
						}
					}

					// Права
					{
						if accessToken.UserInfo.Accesses.Permissions != nil && !allowed {
						CheckAccessForRouteByPermission:
							for _, userPermission := range accessToken.UserInfo.Accesses.Permissions {
								for _, routePermissionID := range route.Accesses.Permissions {
									if userPermission.ID == routePermissionID {
										allowed = true
										break CheckAccessForRouteByPermission
									}
								}
							}
						}
					}
				} else {
					for _, routeRoleID := range route.Accesses.Roles {
						if objects.RoleIDGuest == routeRoleID {
							allowed = true
							break
						}
					}
				}
			}

			if !allowed {
				switch {
				case sessionToken == nil:
					return http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(srv_errors.NotAccess()))
				case sessionToken.UserID == 0 || sessionToken.ProjectID == 0:
					return http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(srv_errors.Unauthorized()))
				default:
					return http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(srv_errors.NotAccess()))
				}
			}
		}
	}

	// X-Authorization-State
	{
		if sessionToken != nil {
			var state string

			switch {
			case sessionToken.UserID == 0 && sessionToken.ProjectID == 0:
				state = "auth"
			case sessionToken.UserID != 0 && sessionToken.ProjectID == 0:
				state = "project-select"
			case sessionToken.UserID != 0 && sessionToken.ProjectID != 0:
				state = "done"
			default:
				state = "unknown"
			}

			ctx.Response().Header.Set("X-Authorization-State", state)
		}
	}

	// Отправка ответа
	{
		if err = http_rest_api_io.Write(ctx.Status(fiber.StatusOK), nil); err != nil {
			acc.components.Logger.Error().
				Format("The response could not be recorded: '%s'. ", err).Write()

			var cErr = common_errors.ResponseCouldNotBeRecorded_RestAPI()
			cErr.SetError(err)

			if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
				acc.components.Logger.Error().
					Format("The error response could not be recorded: '%s'. ", err).Write()

				return http_rest_api_io.WriteError(ctx, common_errors.ResponseCouldNotBeRecorded_RestAPI())
			}

			return
		}

		return
	}
}

// basicAuthenticationProcessingSessionToken - обработка токена сессия в промежуточном программном обеспечении
// для аутентификации пользователя по http маршрутам.
func (acc *accessSystem) basicAuthenticationProcessingSessionToken(ctx fiber.Ctx, accessToken *entities.JwtAccessToken) (token *entities.JwtSessionToken, cErr c_errors.Error) {
	var expired bool

	// Получение
	{
		if raw := ctx.Cookies(acc.conf.CookieKeyForSessionToken); len(raw) > 0 {
			token = new(entities.JwtSessionToken)

			if err := token.Parse(raw); err != nil {
				acc.components.Logger.Error().
					Format("Failed to get session token data: '%s'. ", err).
					Field("raw", raw).Write()

				cErr = srv_errors.InvalidToken()
				cErr.SetError(err)

				return
			}
		}
	}

	// Проверки
	{
		var tm = time.Now()

		// Срок действия ещё не начался
		{
			if token != nil && tm.Before(token.NotBefore) {
				acc.components.Logger.Warn().
					Text("The validity period of the session token has not started yet. ").
					Field("id", token.ID).
					Field("raw", token.Raw).Write()

				token = nil
			}
		}

		// Срок действия уже закончился
		{
			if token != nil && tm.After(token.ExpiresAt) {
				acc.components.Logger.Warn().
					Text("The validity period of the session token has already expired. ").
					Field("id", token.ID).
					Field("raw", token.Raw).Write()

				token = nil
			}
		}
	}

	// Если нужно создать, создаём
	{
		if token == nil || expired {
			// Создание токена
			{
				if expired && token != nil {
					token.ParentID = token.ID
					token.ID = uuid.UUID{}

					token.ExpiresAt = time.Time{}
					token.NotBefore = time.Time{}
					token.IssuedAt = time.Time{}

				} else {
					token = &entities.JwtSessionToken{
						JwtToken: &entities.JwtToken{
							Params: &entities.JwtTokenParams{
								RemoteAddr: string(ctx.Request().Header.Peek("X-Real-IP")),
								UserAgent:  string(ctx.Request().Header.UserAgent()),
							},
						},
					}

					if accessToken != nil {
						token.UserID = accessToken.UserID
						token.ProjectID = accessToken.ProjectID
					}
				}

				if err := token.Generate(); err != nil {
					acc.components.Logger.Error().
						Format("User token generation failed: '%s'. ", err).Write()

					cErr = common_errors.InternalServerError()
					cErr.SetError(err)

					return
				}
			}

			// Печеньки
			{
				ctx.Cookie(&fiber.Cookie{
					Name:        acc.conf.CookieKeyForSessionToken,
					Value:       token.Raw,
					Path:        "/",
					Domain:      string(ctx.Request().Header.Peek("X-Forwarded-Host")),
					MaxAge:      0,
					Expires:     token.ExpiresAt,
					Secure:      false,
					HTTPOnly:    true,
					SameSite:    fiber.CookieSameSiteLaxMode,
					SessionOnly: false,
				})
			}
		}
	}

	return
}

// basicAuthenticationProcessingAccessToken - обработка токена доступа в промежуточном программном обеспечении
// для аутентификации пользователя по http маршрутам.
func (acc *accessSystem) basicAuthenticationProcessingAccessToken(ctx fiber.Ctx) (token *entities.JwtAccessToken, cErr c_errors.Error) {
	// Получение
	{
		var raw string

		if raw = ctx.Cookies(acc.conf.CookieKeyForAccessToken); len(raw) == 0 {
			return
		}

		token = new(entities.JwtAccessToken)

		// Читаем данные из raw
		{
			if err := token.Parse(raw); err != nil {
				acc.components.Logger.Error().
					Format("Failed to get access token data from raw: '%s'. ", err).
					Field("raw", raw).Write()

				cErr = srv_errors.InvalidToken()
				cErr.SetError(err)

				return
			}
		}

		// Получение из redis
		{
			var (
				err     error
				tokenID = token.ID
			)

			if token, err = acc.repositories.JwtTokens.GetJwtAccessToken(ctx.Context(), tokenID); err != nil {
				if !errors.Is(err, redis.Nil) {
					acc.components.Logger.Error().
						Format("Failed to get access token data from redis: '%s'. ", err).
						Field("id", tokenID).Write()

					cErr = srv_errors.InvalidToken()
					cErr.SetError(err)

					return
				}

				acc.components.Logger.Warn().
					Format("Failed to get access token data from redis: '%s'. ", err).
					Field("id", tokenID).Write()
			}
		}
	}

	// Проверки
	{
		var tm = time.Now()

		// Срок действия ещё не начался
		{
			if token != nil && tm.Before(token.NotBefore) {
				acc.components.Logger.Warn().
					Text("The validity period of the access token has not started yet. ").
					Field("id", token.ID).
					Field("raw", token.Raw).Write()

				cErr = srv_errors.ValidityPeriodOfUserTokenHasNotStarted()
				cErr.Details().Set("not_before", token.NotBefore.Format(time.RFC3339Nano))

				return
			}
		}

		// Срок действия уже закончился
		{
			if token != nil && tm.After(token.ExpiresAt) {
				acc.components.Logger.Warn().
					Text("The validity period of the access token has already expired. ").
					Field("id", token.ID).
					Field("raw", token.Raw).Write()

				token = nil
			}
		}
	}

	return
}

// basicAuthenticationProcessingAccessToken - обработка токена доступа в промежуточном программном обеспечении
// для аутентификации пользователя по http маршрутам.
func (acc *accessSystem) basicAuthenticationProcessingRefreshToken(ctx fiber.Ctx) (token *entities.JwtRefreshToken, cErr c_errors.Error) {
	// Получение
	{
		var raw string

		if raw = ctx.Cookies(acc.conf.CookieKeyForRefreshToken); len(raw) == 0 {
			return
		}

		token = new(entities.JwtRefreshToken)

		// Читаем данные из raw
		{
			if err := token.Parse(raw); err != nil {
				acc.components.Logger.Error().
					Format("Failed to get access token data from raw: '%s'. ", err).
					Field("raw", raw).Write()

				cErr = srv_errors.InvalidToken()
				cErr.SetError(err)

				return
			}
		}

		// Получение из redis
		{
			var (
				err     error
				tokenID = token.ID
			)

			if token, err = acc.repositories.JwtTokens.GetJwtRefreshToken(ctx.Context(), tokenID); err != nil {
				if !errors.Is(err, redis.Nil) {
					acc.components.Logger.Error().
						Format("Failed to get refresh token data from redis: '%s'. ", err).
						Field("id", tokenID).Write()

					cErr = srv_errors.InvalidToken()
					cErr.SetError(err)

					return
				}

				acc.components.Logger.Warn().
					Format("Failed to get refresh token data from redis: '%s'. ", err).
					Field("id", tokenID).Write()
			}
		}
	}

	// Проверки
	{
		var tm = time.Now()

		// Срок действия ещё не начался
		{
			if token != nil && tm.Before(token.NotBefore) {
				acc.components.Logger.Warn().
					Text("The validity period of the refresh token has not started yet. ").
					Field("id", token.ID).
					Field("raw", token.Raw).Write()

				token = nil
			}
		}

		// Срок действия уже закончился
		{
			if token != nil && tm.After(token.ExpiresAt) {
				acc.components.Logger.Warn().
					Text("The validity period of the refresh token has already expired. ").
					Field("id", token.ID).
					Field("raw", token.Raw).Write()

				token = nil
			}
		}
	}

	// Завершение жизни
	{
		if token != nil {
			if err := acc.repositories.JwtTokens.Remove(ctx.Context(), token.ID); err != nil {
				acc.components.Logger.Error().
					Format("The refresh token lifetime could not be completed: '%s'. ", err).
					Field("id", token.ID).
					Field("raw", token.Raw).Write()

				token = nil
			}
		}
	}

	return
}
