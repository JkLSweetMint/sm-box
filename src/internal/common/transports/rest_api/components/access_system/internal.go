package access_system

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/common/objects/entities"
	rest_api_io "sm-box/internal/common/transports/rest_api/io"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	c_errors "sm-box/pkg/errors"
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

	return
}
