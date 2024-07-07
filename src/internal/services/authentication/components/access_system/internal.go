package access_system

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/services/authentication/objects/entities"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	c_errors "sm-box/pkg/errors"
	http_rest_api_io "sm-box/pkg/transport/http_rest_api/io"
)

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
		var claims = &jwt.RegisteredClaims{
			Issuer: env.Vars.SystemName,
			Audience: jwt.ClaimStrings{
				string(ctx.Request().Header.UserAgent()),
			},
			ExpiresAt: &jwt.NumericDate{Time: token.ExpiresAt},
			NotBefore: &jwt.NumericDate{Time: token.NotBefore},
			IssuedAt:  &jwt.NumericDate{Time: token.IssuedAt},
		}

		fmt.Printf("%+v\n", claims)

		if err = token.Generate(claims); err != nil {
			acc.components.Logger.Error().
				Format("Failed to generate a token for the client: '%s'. ", err).Write()

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

	var cookie = &fiber.Cookie{
		Name:        acc.conf.CookieKeyForToken,
		Value:       token.Data,
		Path:        "/",
		Domain:      acc.conf.CookieDomain,
		MaxAge:      0,
		Expires:     token.ExpiresAt,
		Secure:      false,
		HTTPOnly:    false,
		SameSite:    fiber.CookieSameSiteLaxMode,
		SessionOnly: false,
	}

	fmt.Printf("%+v\n", cookie)

	ctx.Cookie(cookie)

	// Сохранить в базу
	{
		if err = acc.repositories.JwtTokens.Register(ctx.Context(), token); err != nil {
			acc.components.Logger.Error().
				Format("The client's current could not be registered in the database: '%s'. ", err).Write()

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

	return
}
