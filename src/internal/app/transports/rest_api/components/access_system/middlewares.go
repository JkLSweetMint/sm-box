package access_system

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"regexp"
	"sm-box/internal/app/transports/rest_api/io"
	"sm-box/internal/common/entities"
	errors2 "sm-box/internal/common/errors"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
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

			if err = rest_api_io.WriteError(ctx, errors2.RouteNotFound_RestAPI()); err != nil {
				acc.components.Logger.Error().
					Format("The response could not be recorded: '%s'. ", err).Write()

				return rest_api_io.WriteError(ctx, errors2.ResponseCouldNotBeRecorded_RestAPI())
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

			if err = rest_api_io.WriteError(ctx, errors2.TokenNotFound_RestAPI()); err != nil {
				acc.components.Logger.Error().
					Format("The response could not be recorded: '%s'. ", err).Write()

				return rest_api_io.WriteError(ctx, errors2.ResponseCouldNotBeRecorded_RestAPI())
			}
			return
		}
	}

	// Проверка требуется ли авторизация, если да проводим
	{
		if route.Authorize {
			if token.UserID == 0 {
				return rest_api_io.WriteError(ctx, errors2.Unauthorized_RestAPI())
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

						if err = rest_api_io.WriteError(ctx, errors2.UserNotFound_RestAPI()); err != nil {
							acc.components.Logger.Error().
								Format("The response could not be recorded: '%s'. ", err).Write()

							return rest_api_io.WriteError(ctx, errors2.ResponseCouldNotBeRecorded_RestAPI())
						}
						return
					}

					acc.components.Logger.Error().
						Format("User data could not be retrieved: '%s'. ", err).
						Field("user_id", token.UserID).Write()

					if err = rest_api_io.WriteError(ctx, errors2.InternalServerError_RestAPI()); err != nil {
						acc.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						return rest_api_io.WriteError(ctx, errors2.ResponseCouldNotBeRecorded_RestAPI())
					}
					return
				}
			}

			if !us.CheckHttpRouteAccesses(route.Accesses) {
				acc.components.Logger.Error().
					Text("The user does not have access to visit this route. ").
					Field("user_id", token.UserID).
					Field("route", route).Write()

				if err = rest_api_io.WriteError(ctx, errors2.Forbidden_RestAPI()); err != nil {
					acc.components.Logger.Error().
						Format("The response could not be recorded: '%s'. ", err).Write()

					return rest_api_io.WriteError(ctx, errors2.ResponseCouldNotBeRecorded_RestAPI())
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

				if err = rest_api_io.WriteError(ctx, errors2.TokenNotFound_RestAPI()); err != nil {
					acc.components.Logger.Error().
						Format("The response could not be recorded: '%s'. ", err).Write()

					return rest_api_io.WriteError(ctx, errors2.ResponseCouldNotBeRecorded_RestAPI())
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

				var cErr = errors2.ValidityPeriodOfUserTokenHasNotStarted_RestAPI()
				cErr.Details().Set("not_before", token.NotBefore.UTC().Format(time.RFC3339Nano))

				if err = rest_api_io.WriteError(ctx, cErr); err != nil {
					acc.components.Logger.Error().
						Format("The response could not be recorded: '%s'. ", err).Write()

					return rest_api_io.WriteError(ctx, errors2.ResponseCouldNotBeRecorded_RestAPI())
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

// generateToken - генерация токена.
func (acc *accessSystem) generateToken(ctx fiber.Ctx, token *entities.JwtToken) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelComponentInternal)

		trc.FunctionCall(ctx, token)
		defer func() { trc.Error(err).FunctionCallFinished(token) }()
	}

	// Генерация данных токена
	{
		var (
			claims = &jwt.RegisteredClaims{
				Issuer: env.Vars.SystemName,
				Audience: jwt.ClaimStrings{
					string(ctx.Request().Header.UserAgent()),
				},
				ExpiresAt: &jwt.NumericDate{Time: token.ExpiresAt},
				NotBefore: &jwt.NumericDate{Time: token.NotBefore},
				IssuedAt:  &jwt.NumericDate{Time: token.IssuedAt},
			}
		)

		if err = token.Generate(claims); err != nil {
			acc.components.Logger.Error().
				Format("Failed to generate a token for the client: '%s'. ", err).Write()

			if err = rest_api_io.WriteError(ctx, errors2.InternalServerError_RestAPI()); err != nil {
				acc.components.Logger.Error().
					Format("The response could not be recorded: '%s'. ", err).Write()

				return rest_api_io.WriteError(ctx, errors2.ResponseCouldNotBeRecorded_RestAPI())
			}
			return
		}
	}

	ctx.Cookie(&fiber.Cookie{
		Name:        acc.conf.CookieKeyForToken,
		Value:       token.Data,
		Path:        "/",
		Domain:      string(ctx.Request().Host()),
		MaxAge:      0,
		Expires:     token.ExpiresAt,
		Secure:      false,
		HTTPOnly:    false,
		SameSite:    fiber.CookieSameSiteLaxMode,
		SessionOnly: false,
	})

	// Сохранить в базу
	{
		if err = acc.repository.RegisterToken(ctx.Context(), token); err != nil {
			acc.components.Logger.Error().
				Format("The client's current could not be registered in the database: '%s'. ", err).Write()

			if err = rest_api_io.WriteError(ctx, errors2.InternalServerError_RestAPI()); err != nil {
				acc.components.Logger.Error().
					Format("The response could not be recorded: '%s'. ", err).Write()

				return rest_api_io.WriteError(ctx, errors2.ResponseCouldNotBeRecorded_RestAPI())
			}
			return
		}
	}

	return
}
