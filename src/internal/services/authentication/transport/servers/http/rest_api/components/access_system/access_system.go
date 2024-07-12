package access_system

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"net/url"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/services/authentication/objects/entities"
	"sm-box/pkg/core/components/logger"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/http/rest_api/io"
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
		token *entities.JwtToken
		route *entities.HttpRoute
	)

	// Работа с токеном
	{
		// Получение токена, если нет, то создаём
		{
			if raw := ctx.Cookies(acc.conf.CookieKeyForToken); raw == "" {
				token = &entities.JwtToken{
					Params: &entities.JwtTokenParams{
						RemoteAddr: fmt.Sprintf("%s:%s", ctx.IP(), ctx.Port()),
						UserAgent:  string(ctx.Request().Header.UserAgent()),
					},
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

				// Запись печеньки
				{
					var cookie = &fiber.Cookie{
						Name:        acc.conf.CookieKeyForToken,
						Value:       token.Raw,
						Path:        "/",
						Domain:      acc.conf.CookieDomain,
						MaxAge:      0,
						Expires:     token.ExpiresAt,
						Secure:      false,
						HTTPOnly:    false,
						SameSite:    fiber.CookieSameSiteLaxMode,
						SessionOnly: false,
					}

					ctx.Cookie(cookie)
				}

				// Сохранение в базе
				{
					if err = acc.repositories.JwtTokens.Register(ctx.Context(), token); err != nil {
						acc.components.Logger.Error().
							Format("The client's token could not be registered in the database: '%s'. ", err).Write()

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
			} else {
				if len(raw) == 0 {
					if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(error_list.TokenWasNotTransferred())); err != nil {
						acc.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
					}
					return
				}

				token = new(entities.JwtToken)
				token.FillEmptyFields()

				if err = token.Parse(raw); err != nil {
					acc.components.Logger.Error().
						Format("Failed to get token data: '%s'. ", err).
						Field("raw", raw).Write()

					if err = http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(error_list.InternalServerError())); err != nil {
						acc.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
					}
					return
				}
			}
		}

		// Проверка времени жизни токена, если закончился, нужно пересоздать
		{
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
					token = &entities.JwtToken{
						ParentID: token.ID,

						UserID:    token.UserID,
						ProjectID: token.ProjectID,

						Params: &entities.JwtTokenParams{
							RemoteAddr: fmt.Sprintf("%s:%s", ctx.IP(), ctx.Port()),
							UserAgent:  string(ctx.Request().Header.UserAgent()),
						},
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

					// Запись печеньки
					{
						var cookie = &fiber.Cookie{
							Name:        acc.conf.CookieKeyForToken,
							Value:       token.Raw,
							Path:        "/",
							Domain:      acc.conf.CookieDomain,
							MaxAge:      0,
							Expires:     token.ExpiresAt,
							Secure:      false,
							HTTPOnly:    false,
							SameSite:    fiber.CookieSameSiteLaxMode,
							SessionOnly: false,
						}

						ctx.Cookie(cookie)
					}

					// Сохранение в базе
					{
						if err = acc.repositories.JwtTokens.Register(ctx.Context(), token); err != nil {
							acc.components.Logger.Error().
								Format("The client's token could not be registered in the database: '%s'. ", err).Write()

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
				}
			}
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
				if token.UserID == 0 || token.ProjectID == 0 {
					return http_rest_api_io.WriteError(ctx, c_errors.ToRestAPI(error_list.Unauthorized()))
				}
			}
		}

		// Проверка доступа к маршруту, если требуется авторизация
		{
			if route.Authorize {

			}
		}
	}

	// Authorization заголовок
	{
		if token.UserID != 0 {
			ctx.Set("Authorization", fmt.Sprintf("Bearer %s", token.Raw))
		}
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
