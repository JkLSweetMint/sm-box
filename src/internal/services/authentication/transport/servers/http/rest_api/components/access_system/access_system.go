package access_system

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"net/url"
	app_models "sm-box/internal/app/objects/models"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/common/types"
	"sm-box/internal/services/authentication/objects/entities"
	users_models "sm-box/internal/services/users/objects/models"
	"sm-box/pkg/core/components/logger"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/http/rest_api/io"
	"strings"
	"time"
)

// accessSystem - компонент системы доступа http маршрутов.
type accessSystem struct {
	conf *Config
	ctx  context.Context

	components   *components
	gateways     *gateways
	repositories *repositories
}

type (
	// components - компоненты компонента.
	components struct {
		Logger logger.Logger
	}

	// gateways - шлюзы компонента.
	gateways struct {
		Projects interface {
			Get(ctx context.Context, ids ...types.ID) (list app_models.ProjectList, cErr c_errors.Error)
			GetOne(ctx context.Context, id types.ID) (project *app_models.ProjectInfo, cErr c_errors.Error)
		}
		Users interface {
			Get(ctx context.Context, ids ...types.ID) (list []*users_models.UserInfo, cErr c_errors.Error)
			GetOne(ctx context.Context, id types.ID) (project *users_models.UserInfo, cErr c_errors.Error)
		}
	}

	// repositories - репозитории компонента.
	repositories struct {
		HttpRoutes interface {
			Get(ctx context.Context, protocol, method, path string) (route *entities.HttpRoute, err error)
			GetActive(ctx context.Context, protocol, method, path string) (route *entities.HttpRoute, err error)
		}
		JwtTokens interface {
			Register(ctx context.Context, tok *entities.JwtToken) (err error)
			Disable(ctx context.Context, raw string) (err error)
			GetToken(ctx context.Context, raw string) (tok *entities.JwtToken, err error)
		}
	}
)

// Middleware - промежуточное программное обеспечение аутентификации пользователя по http маршрутам.
func (acc *accessSystem) Middleware(ctx fiber.Ctx) (err error) {
	var (
		sessionToken *entities.JwtSessionToken
		accessToken  *entities.JwtAccessToken
		route        *entities.HttpRoute
	)

	// Работа с токенами
	{
		// Сессия
		{
			var (
				token   = sessionToken
				expired bool
			)

			// Получение
			{
				if raw := ctx.Cookies(acc.conf.CookieKeyForSessionToken); len(raw) > 0 {
					token = new(entities.JwtSessionToken)

					if err = token.Parse(raw); err != nil {
						acc.components.Logger.Error().
							Format("Failed to get session token data: '%s'. ", err).
							Field("raw", raw).Write()

						if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(error_list.InvalidToken())); err != nil {
							acc.components.Logger.Error().
								Format("The response could not be recorded: '%s'. ", err).Write()

							return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
						}
						return
					}
				}
			}

			// Проверки
			{
				if token != nil {
					var tm = time.Now()

					// Срок действия ещё не начался
					{
						if tm.Before(token.NotBefore) {
							acc.components.Logger.Warn().
								Text("The validity period of the user's token has not started yet. ").
								Field("token", token).Write()

							var cErr = error_list.ValidityPeriodOfUserTokenHasNotStarted()
							cErr.Details().Set("not_before", token.NotBefore.Format(time.RFC3339Nano))

							if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(cErr)); err != nil {
								acc.components.Logger.Error().
									Format("The response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
							}
							return
						}
					}

					// Срок действия уже закончился, пересоздаём
					{
						if tm.After(token.ExpiresAt) {
							expired = true
						}
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
										RemoteAddr: fmt.Sprintf("%s:%s", ctx.IP(), ctx.Port()),
										UserAgent:  string(ctx.Request().Header.UserAgent()),
									},
								},
							}
						}

						if err = token.Generate(); err != nil {
							acc.components.Logger.Error().
								Format("User token generation failed: '%s'. ", err).Write()

							var cErr = error_list.InternalServerError()
							cErr.SetError(err)

							if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(cErr)); err != nil {
								acc.components.Logger.Error().
									Format("The response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
							}
							return
						}
					}

					// Регистрация в базе
					{
						if err = acc.repositories.JwtTokens.Register(ctx.Context(), token.JwtToken); err != nil {
							acc.components.Logger.Error().
								Format("The client's session token could not be registered in the database: '%s'. ", err).Write()

							var cErr = error_list.InternalServerError()
							cErr.SetError(err)

							if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(cErr)); err != nil {
								acc.components.Logger.Error().
									Format("The response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
							}
							return
						}
					}

					// Печеньки
					{
						ctx.Cookie(&fiber.Cookie{
							Name:        acc.conf.CookieKeyForSessionToken,
							Value:       token.Raw,
							Path:        "/",
							Domain:      string(ctx.Request().Header.Peek("X-Original-HOST")),
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

			sessionToken = token
		}

		// Обновления
		{

		}

		// Доступа
		{
			var token = accessToken

			// Получение
			{
				var raw = strings.Replace(string(ctx.Request().Header.Peek("Authorization")), "Bearer ", "", 1)

				if len(raw) == 0 {
					raw = ctx.Cookies(acc.conf.CookieKeyForAccessToken)
				}

				if len(raw) > 0 {
					token = new(entities.JwtAccessToken)

					if err = token.Parse(raw); err != nil {
						acc.components.Logger.Error().
							Format("Failed to get access token data: '%s'. ", err).
							Field("raw", raw).Write()

						if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(error_list.InvalidToken())); err != nil {
							acc.components.Logger.Error().
								Format("The response could not be recorded: '%s'. ", err).Write()

							return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
						}
						return
					}
				}
			}

			// Проверки
			{
				if token != nil {
					var tm = time.Now()

					// Срок действия ещё не начался
					{
						if tm.Before(token.NotBefore) {
							acc.components.Logger.Warn().
								Text("The validity period of the user's token has not started yet. ").
								Field("token", token).Write()

							var cErr = error_list.ValidityPeriodOfUserTokenHasNotStarted()
							cErr.Details().Set("not_before", token.NotBefore.Format(time.RFC3339Nano))

							if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(cErr)); err != nil {
								acc.components.Logger.Error().
									Format("The response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
							}
							return
						}
					}

					// Срок действия уже закончился
					{
						if tm.After(token.ExpiresAt) {
							token = nil
						}
					}
				}
			}

			// Если устарел или не передан, создаём новый если есть токен обновления
			{
				if raw := ctx.Cookies(acc.conf.CookieKeyForRefreshToken); token == nil && len(raw) > 0 {
					var (
						refreshToken = new(entities.JwtRefreshToken)
						user         *users_models.UserInfo
						project      *app_models.ProjectInfo
					)

					// Получение токена обновления если ещё живой
					{
						if refreshToken.JwtToken, err = acc.repositories.JwtTokens.GetToken(ctx.Context(), raw); err != nil {
							acc.components.Logger.Error().
								Format("Failed to get refresh token: '%s'. ", err).
								Field("raw", raw).Write()

							var cErr c_errors.RestAPI

							if errors.Is(err, sql.ErrNoRows) {
								cErr = c_errors.ToRestAPI(error_list.AnUnregisteredTokenWasTransferred())
							} else {
								cErr = c_errors.ToRestAPI(error_list.InternalServerError())
							}

							if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
								acc.components.Logger.Error().
									Format("The response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
							}
							return
						}
					}

					// Завершения токена обновления
					{
						if err = acc.repositories.JwtTokens.Disable(ctx.Context(), raw); err != nil {
							acc.components.Logger.Error().
								Format("Failed to complete the validity period of the refresh token: '%s'. ", err).
								Field("raw", raw).Write()

							var cErr = c_errors.ToRestAPI(error_list.InternalServerError())

							if err = http_rest_api_io.WriteError(ctx, cErr); err != nil {
								acc.components.Logger.Error().
									Format("The response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
							}
							return
						}
					}

					// Получение данных пользователя
					{
						var cErr c_errors.Error

						if user, cErr = acc.gateways.Users.GetOne(ctx.Context(), sessionToken.UserID); cErr != nil {
							acc.components.Logger.Error().
								Format("Failed to get the user data: '%s'. ", cErr).
								Field("id", sessionToken.UserID).Write()

							if errors.Is(cErr, sql.ErrNoRows) {
								cErr = error_list.NotAccess()
							} else {
								cErr = error_list.InternalServerError()
							}

							if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(cErr)); err != nil {
								acc.components.Logger.Error().
									Format("The response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
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
								cErr = error_list.NotAccess()
							} else {
								cErr = error_list.InternalServerError()
							}

							if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(cErr)); err != nil {
								acc.components.Logger.Error().
									Format("The response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
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
							ids  = make(map[types.ID]struct{})
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

							for _, rl := range user.Accesses {
								writeInheritance(rl.RoleInfo)
							}
						}

						if _, ok := ids[project.ID]; !ok {
							acc.components.Logger.Error().
								Format("The user does not have access to the project: '%s'. ", cErr).
								Field("project_id", project.ID).
								Field("user_id", user.ID).Write()

							if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(error_list.NotAccessToProject())); err != nil {
								acc.components.Logger.Error().
									Format("The response could not be recorded: '%s'. ", err).Write()

								return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
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
								Claims: nil,
							}

							if err = refreshToken.Generate(); err != nil {
								acc.components.Logger.Error().
									Format("User session token generation failed: '%s'. ", err).Write()

								if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(error_list.InternalServerError())); err != nil {
									acc.components.Logger.Error().
										Format("The response could not be recorded: '%s'. ", err).Write()

									return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
								}
								return
							}

							// Сохранение в базе
							{
								if err = acc.repositories.JwtTokens.Register(ctx.Context(), refreshToken.JwtToken); err != nil {
									acc.components.Logger.Error().
										Format("The client's token could not be registered in the database: '%s'. ", err).Write()

									if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(error_list.InternalServerError())); err != nil {
										acc.components.Logger.Error().
											Format("The response could not be recorded: '%s'. ", err).Write()

										return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
									}
									return
								}
							}
						}

						// Создание токена доступа
						{
							token = &entities.JwtAccessToken{
								JwtToken: &entities.JwtToken{
									ProjectID: sessionToken.ProjectID,
									ParentID:  refreshToken.ID,
									UserID:    user.ID,

									Params: sessionToken.Params,
								},
								Claims: nil,
							}

							// Запись доступов пользователя
							{
								token.Claims = &entities.JwtAccessTokenClaims{
									Accesses: user.Accesses.ListIDs(),
								}
							}

							if err = token.Generate(); err != nil {
								acc.components.Logger.Error().
									Format("User session token generation failed: '%s'. ", err).Write()

								if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(error_list.InternalServerError())); err != nil {
									acc.components.Logger.Error().
										Format("The response could not be recorded: '%s'. ", err).Write()

									return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
								}
								return
							}

							// Сохранение в базе
							{
								if err = acc.repositories.JwtTokens.Register(ctx.Context(), token.JwtToken); err != nil {
									acc.components.Logger.Error().
										Format("The client's token could not be registered in the database: '%s'. ", err).Write()

									if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(error_list.InternalServerError())); err != nil {
										acc.components.Logger.Error().
											Format("The response could not be recorded: '%s'. ", err).Write()

										return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
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
							Value:       token.Raw,
							Path:        "/",
							Domain:      string(ctx.Request().Header.Peek("X-Original-HOST")),
							MaxAge:      0,
							Expires:     token.ExpiresAt,
							Secure:      true,
							HTTPOnly:    true,
							SameSite:    fiber.CookieSameSiteLaxMode,
							SessionOnly: false,
						})

						ctx.Cookie(&fiber.Cookie{
							Name:        acc.conf.CookieKeyForRefreshToken,
							Value:       refreshToken.Raw,
							Path:        "/",
							Domain:      string(ctx.Request().Header.Peek("X-Original-HOST")),
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

			accessToken = token
		}
	}

	// Работа с маршрутом
	{
		// Получение маршрута
		{
			var (
				protocol string
				method   = string(ctx.Request().Header.Peek("X-Original-Method"))
				urlStr   = string(ctx.Request().Header.Peek("X-Original-URL"))
				u        *url.URL
			)

			// Парсинг URL
			{
				if u, err = url.Parse(urlStr); err != nil {
					acc.components.Logger.Error().
						Format("The URL could not be parsed: '%s'. ", err).
						Field("method", method).
						Field("url", urlStr).Write()

					if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(error_list.InternalServerError())); err != nil {
						acc.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
					}
					return
				}
			}

			protocol = u.Scheme

			if string(ctx.Request().Header.Peek("Upgrade")) == "websocket" {
				switch protocol {
				case "http":
					protocol = "ws"
				case "https":
					protocol = "wss"
				}
			}

			if route, err = acc.repositories.HttpRoutes.GetActive(ctx.Context(), protocol, method, u.Path); err != nil {
				acc.components.Logger.Error().
					Format("Failed to get a route: '%s'. ", err).
					Field("method", method).
					Field("path", u.Path).Write()

				if err = http_rest_api_io.WriteError(ctx, error_list.RouteNotFound_RestAPI()); err != nil {
					acc.components.Logger.Error().
						Format("The response could not be recorded: '%s'. ", err).Write()

					return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
				}
				return
			}
		}

		// Проверка требуется ли авторизация, если да проводим
		{
			if route.Authorize {
				if accessToken == nil || accessToken.UserID == 0 || accessToken.ProjectID == 0 {
					return http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(error_list.Unauthorized()))
				}
			}
		}

		// Проверка доступа к маршруту, если требуется авторизация
		{
			if route.Authorize {
				var forbidden = true

				for _, access := range route.Accesses {
					for _, id := range accessToken.Claims.Accesses {
						if types.ID(access) == id {
							forbidden = false
							break
						}
					}
				}

				if forbidden {
					if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(error_list.NotAccess())); err != nil {
						acc.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
					}
					return
				}
			}
		}
	}

	// X-Authorization-State
	{
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

	// Отправка ответа
	{
		if err = http_rest_api_io.Write(ctx.Status(fiber.StatusOK), nil); err != nil {
			acc.components.Logger.Error().
				Format("The response could not be recorded: '%s'. ", err).Write()

			var cErr = error_list.ResponseCouldNotBeRecorded_RestAPI()
			cErr.SetError(err)

			return http_rest_api_io.WriteError(ctx, cErr)
		}

		return
	}
}
