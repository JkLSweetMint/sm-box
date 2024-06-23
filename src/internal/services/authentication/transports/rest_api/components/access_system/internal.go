package access_system

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"regexp"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/common/objects/entities"
	src_access_system "sm-box/internal/common/transports/rest_api/components/access_system"
	rest_api_io "sm-box/internal/common/transports/rest_api/io"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	error_details "sm-box/pkg/errors/entities/details"
	error_messages "sm-box/pkg/errors/entities/messages"
)

// accessSystem - компонент системы доступа http rest api.
type accessSystem struct {
	src_access_system.AccessSystem

	conf *src_access_system.Config
	ctx  context.Context

	components *components
	repository interface {
		BasicAuth(ctx context.Context, username, password string) (us *entities.User, err error)

		GetToken(ctx context.Context, data string) (tok *entities.JwtToken, err error)
		SetTokenOwner(ctx context.Context, tokenID, ownerID types.ID) (err error)
	}
}

// components - компоненты компонента системы доступа http rest api.
type components struct {
	Logger logger.Logger
}

// BasicUserAuth - базовый обработчик для авторизации пользователя.
// Для авторизации используется имя пользователя и пароль.
func (acc *accessSystem) BasicUserAuth(ctx fiber.Ctx) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelComponent)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var requestData = new(BasicUserAuthData)

	// Чтение данных
	{
		if err = ctx.Bind().Body(requestData); err != nil {
			acc.components.Logger.Error().
				Format("The request body data could not be read: '%s'. ", err).Write()

			if err = rest_api_io.WriteError(ctx, error_list.RequestBodyDataCouldNotBeRead_RestAPI()); err != nil {
				acc.components.Logger.Error().
					Format("The response could not be recorded: '%s'. ", err).Write()

				var cErr = error_list.ResponseCouldNotBeRecorded_RestAPI()
				cErr.SetError(err)

				return rest_api_io.WriteError(ctx, cErr)
			}

			return
		}
	}

	// Проверка данных
	{
		if requestData.Username == "" || requestData.Password == "" {
			var cErr = error_list.InvalidArgumentsValue_RestAPI()

			if requestData.Username == "" {
				cErr.Details().SetField(new(error_details.FieldKey).Add("username"), new(error_messages.TextMessage).Text("Zero value. "))
			}

			if requestData.Password == "" {
				cErr.Details().SetField(new(error_details.FieldKey).Add("password"), new(error_messages.TextMessage).Text("Zero value. "))
			}

			if err = rest_api_io.WriteError(ctx, cErr); err != nil {
				acc.components.Logger.Error().
					Format("The response could not be recorded: '%s'. ", err).Write()

				var cErr = error_list.ResponseCouldNotBeRecorded_RestAPI()
				cErr.SetError(err)

				return rest_api_io.WriteError(ctx, cErr)
			}

			return
		}
	}

	// Обработка
	{
		var (
			us    *entities.User
			token *entities.JwtToken
		)

		// Получение данных пользователя
		{
			if us, err = acc.repository.BasicAuth(ctx.Context(), requestData.Username, requestData.Password); err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					acc.components.Logger.Warn().
						Format("User authorization error: '%s'. ", err).Write()

					if err = rest_api_io.WriteError(ctx, error_list.UserNotFound_RestAPI()); err != nil {
						acc.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						return rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
					}
					return
				}

				acc.components.Logger.Error().
					Format("User authorization error: '%s'. ", err).
					Field("username", requestData.Username).
					Field("password", requestData.Password).Write()

				if err = rest_api_io.WriteError(ctx, error_list.InternalServerError_RestAPI()); err != nil {
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

				if err = rest_api_io.WriteError(ctx, error_list.TokenNotFound_RestAPI()); err != nil {
					acc.components.Logger.Error().
						Format("The response could not be recorded: '%s'. ", err).Write()

					return rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
				}
				return
			}
		}

		// Проверка что уже авторизован
		{
			if token.UserID != 0 {
				acc.components.Logger.Warn().
					Text("The user is already logged in. ").
					Field("user_id", us.ID).
					Field("token", token).Write()

				if err = rest_api_io.WriteError(ctx, error_list.AlreadyAuthorized_RestAPI()); err != nil {
					acc.components.Logger.Error().
						Format("The response could not be recorded: '%s'. ", err).Write()

					return rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
				}
				return
			}
		}

		// Обновление данных токена
		{
			if err = acc.repository.SetTokenOwner(ctx.Context(), token.ID, us.ID); err != nil {
				acc.components.Logger.Error().
					Format("The token owner could not be identified: '%s'. ", err).
					Field("owner_id", us.ID).
					Field("token", token).Write()

				if err = rest_api_io.WriteError(ctx, error_list.InternalServerError_RestAPI()); err != nil {
					acc.components.Logger.Error().
						Format("The response could not be recorded: '%s'. ", err).Write()

					return rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
				}
				return
			}
		}
	}

	// Отправка ответа
	{
		if err = rest_api_io.Write(ctx.Status(fiber.StatusOK), nil); err != nil {
			acc.components.Logger.Error().
				Format("The response could not be recorded: '%s'. ", err).Write()

			return rest_api_io.WriteError(ctx, error_list.ResponseCouldNotBeRecorded_RestAPI())
		}
		return
	}
}
