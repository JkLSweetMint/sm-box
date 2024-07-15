package access_system

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"net/url"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/common/types"
	"sm-box/internal/services/authentication/objects/entities"
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
	}

	// repositories - репозитории компонента.
	repositories struct {
		HttpRoutes interface {
			Get(ctx context.Context, protocol, method, path string) (route *entities.HttpRoute, err error)
			GetActive(ctx context.Context, protocol, method, path string) (route *entities.HttpRoute, err error)
		}
		JwtTokens interface {
			Register(ctx context.Context, tok *entities.JwtToken) (err error)
		}
	}
)

// AuthenticationMiddleware - промежуточное программное обеспечение аутентификации пользователя по http маршрутам.
func (acc *accessSystem) AuthenticationMiddleware(ctx fiber.Ctx) (err error) {
	var (
		sessionToken *entities.JwtSessionToken
		accessToken  *entities.JwtAccessToken
		route        *entities.HttpRoute
	)

	// Работа с токеном
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
