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
	http_rest_api_io "sm-box/pkg/transport/http_rest_api/io"
	"time"
)

// accessSystem - компонент системы доступа http rest api.
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

// AuthenticationMiddlewareForRestAPI - промежуточное программное обеспечение аутентификации пользователя для rest api.
func (acc *accessSystem) AuthenticationMiddlewareForRestAPI(ctx fiber.Ctx) (err error) {
	var (
		token *entities.JwtToken
		route *entities.HttpRoute
	)

	// Работа с токеном
	{
		// Получение токена, если нет, то создаём
		{
			if data := ctx.Cookies(acc.conf.CookieKeyForToken); data == "" {
				token = &entities.JwtToken{
					ID:        0,
					UserID:    0,
					Raw:       "",
					ExpiresAt: time.Now().Add(time.Hour),
					NotBefore: time.Now(),
					IssuedAt:  time.Now(),
					Params: &entities.JwtTokenParams{
						RemoteAddr: fmt.Sprintf("%s:%s", ctx.IP(), ctx.Port()),
						UserAgent:  string(ctx.Request().Header.UserAgent()),
					},
				}

				token.FillEmptyFields()

				if err = acc.generateToken(ctx, token); err != nil {
					return
				}
			} else {
				if len(data) == 0 {
					if err = http_rest_api_io.WriteError(ctx, error_list.TokenHasNotBeenTransferred_RestAPI()); err != nil {
						acc.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						return http_rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
					}
					return
				}

				token = new(entities.JwtToken)
				token.FillEmptyFields()

				if err = token.Parse(data); err != nil {
					acc.components.Logger.Error().
						Format("Failed to get token data: '%s'. ", err).
						Field("data", data).Write()

					if err = http_rest_api_io.WriteError(ctx, error_list.AnUnregisteredTokenWasTransderred_RestAPI()); err != nil {
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
					cErr.Details().Set("not_before", token.NotBefore)

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
						ID:        0,
						UserID:    0,
						Raw:       "",
						ExpiresAt: time.Now().Add(time.Hour),
						NotBefore: time.Now(),
						IssuedAt:  time.Now(),
						Params: &entities.JwtTokenParams{
							RemoteAddr: fmt.Sprintf("%s:%s", ctx.IP(), ctx.Port()),
							UserAgent:  string(ctx.Request().Header.UserAgent()),
						},
					}

					token.FillEmptyFields()

					if err = acc.generateToken(ctx, token); err != nil {
						return
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
				method   = string(ctx.Request().Header.Method())
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
				if token.UserID == 0 {
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
