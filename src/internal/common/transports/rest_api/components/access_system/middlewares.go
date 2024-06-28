package access_system

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"regexp"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/common/objects/entities"
	rest_api_io "sm-box/internal/common/transports/rest_api/io"
	c_errors "sm-box/pkg/errors"
	"time"
)

// AuthenticationMiddleware - промежуточное программное обеспечение для аутентификации пользователя.
func (acc *accessSystem) AuthenticationMiddleware(ctx fiber.Ctx) (err error) {
	var (
		route *entities.HttpRoute
		token *entities.JwtToken
	)

	// Получение маршрута
	{
		var (
			method = string(ctx.Request().Header.Method())
			path   = string(ctx.Request().URI().Path())
		)

		if route, err = acc.repository.GetActiveRoute(ctx.Context(), method, path); err != nil {
			acc.components.Logger.Error().
				Format("Failed to get a route: '%s'. ", err).
				Field("method", method).
				Field("path", path).Write()

			if err = rest_api_io.WriteError(ctx, error_list.RouteNotFound_RestAPI()); err != nil {
				acc.components.Logger.Error().
					Format("The response could not be recorded: '%s'. ", err).Write()

				return rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
			}
			return
		}
	}

	// Получение токена
	{
		var data string

		if data = ctx.Cookies(acc.conf.CookieKeyForToken); data == "" {
			var (
				value   = ctx.Response().Header.PeekCookie(acc.conf.CookieKeyForToken)
				pattern = fmt.Sprintf(`^%s=([\s\S]+);\sexpires=[\s\S]+;\sdomain=[\s\S]+;\spath=[\s\S]+;\sSameSite=[\s\S]+$`, acc.conf.CookieKeyForToken)
				re      = regexp.MustCompile(pattern)
			)

			data = re.FindStringSubmatch(string(value))[1]
		}

		if token, err = acc.repository.GetToken(ctx.Context(), data); err != nil {
			acc.components.Logger.Error().
				Format("Failed to get token data: '%s'. ", err).
				Field("data", data).Write()

			if err = rest_api_io.WriteError(ctx, error_list.TokenHasNotBeenTransferred_RestAPI()); err != nil {
				acc.components.Logger.Error().
					Format("The response could not be recorded: '%s'. ", err).Write()

				return rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
			}
			return
		}
	}

	// Проверка требуется ли авторизация, если да проводим
	{
		if route.Authorize {
			if token.UserID == 0 {
				return rest_api_io.WriteError(ctx, error_list.Unauthorized_RestAPI())
			}
		}
	}

	// Проверка доступа к маршруту, если требуется авторизация
	{
		if route.Authorize {
			var us *entities.User

			// Получение данных пользователя
			{
				if us, err = acc.repository.GetUser(ctx.Context(), token.UserID); err != nil {
					if errors.Is(err, sql.ErrNoRows) {
						acc.components.Logger.Warn().
							Format("User not found: '%s'. ", err).Write()

						if err = rest_api_io.WriteError(ctx, c_errors.ToRestAPI(error_list.UserNotFound())); err != nil {
							acc.components.Logger.Error().
								Format("The response could not be recorded: '%s'. ", err).Write()

							return rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
						}
						return
					}

					acc.components.Logger.Error().
						Format("User data could not be retrieved: '%s'. ", err).
						Field("user_id", token.UserID).Write()

					var cErr = error_list.InternalServerError()
					cErr.SetError(err)

					if err = rest_api_io.WriteError(ctx, c_errors.ToRestAPI(cErr)); err != nil {
						acc.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						return rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
					}
					return
				}
			}

			if !us.CheckHttpRouteAccesses(route.Accesses) {
				acc.components.Logger.Error().
					Text("The user does not have access to visit this route. ").
					Field("user_id", token.UserID).
					Field("route", route).Write()

				if err = rest_api_io.WriteError(ctx, error_list.Forbidden_RestAPI()); err != nil {
					acc.components.Logger.Error().
						Format("The response could not be recorded: '%s'. ", err).Write()

					return rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
				}
				return
			}
		}
	}

	return ctx.Next()
}

// IdentificationMiddleware - промежуточное программное обеспечение для идентификации клиента.
func (acc *accessSystem) IdentificationMiddleware(ctx fiber.Ctx) (err error) {
	var token *entities.JwtToken

	// Получение токена, если нет, то создаём
	{
		if data := ctx.Cookies(acc.conf.CookieKeyForToken); data == "" {
			token = &entities.JwtToken{
				ID:        0,
				UserID:    0,
				Data:      "",
				ExpiresAt: time.Now().Add(time.Hour * 8),
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
			if token, err = acc.repository.GetToken(ctx.Context(), data); err != nil {
				acc.components.Logger.Error().
					Format("Failed to get token data: '%s'. ", err).
					Field("data", data).Write()

				if err = rest_api_io.WriteError(ctx, error_list.TokenHasNotBeenTransferred_RestAPI()); err != nil {
					acc.components.Logger.Error().
						Format("The response could not be recorded: '%s'. ", err).Write()

					return rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
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

				if err = rest_api_io.WriteError(ctx, c_errors.ToRestAPI(cErr)); err != nil {
					acc.components.Logger.Error().
						Format("The response could not be recorded: '%s'. ", err).Write()

					return rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
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
					Data:      "",
					ExpiresAt: time.Now().Add(time.Hour * 8),
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

	return ctx.Next()
}
