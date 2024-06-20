package access_system

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"regexp"
	"sm-box/internal/app/errors"
	entities2 "sm-box/internal/app/infrastructure/objects/entities"
	"sm-box/internal/app/infrastructure/types"
	rest_api_io "sm-box/internal/app/transports/rest_api/io"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"strings"
	"time"
)

// accessSystem - компонент системы доступа http rest api.
type accessSystem struct {
	conf *Config
	ctx  context.Context

	components *components
	repository interface {
		GetUser(ctx context.Context, id types.ID) (us *entities2.User, err error)
		BasicAuth(ctx context.Context, username, password string) (us *entities2.User, err error)

		GetRoute(ctx context.Context, method, path string) (route *entities2.HttpRoute, err error)
		GetActiveRoute(ctx context.Context, method, path string) (route *entities2.HttpRoute, err error)
		RegisterRoute(ctx context.Context, route *entities2.HttpRoute) (err error)
		SetInactiveRoutes(ctx context.Context) (err error)

		GetToken(ctx context.Context, data string) (tok *entities2.JwtToken, err error)
		RegisterToken(ctx context.Context, tok *entities2.JwtToken) (err error)
		SetTokenOwner(ctx context.Context, tokenID, ownerID types.ID) (err error)
	}
}

// components - компоненты компонента системы доступа http rest api.
type components struct {
	Logger logger.Logger
}

// RegisterRoutes - регистрация маршрутов в системе.
func (acc *accessSystem) RegisterRoutes(list ...*fiber.Route) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelComponent)

		trc.FunctionCall(list)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	for _, r := range list {
		var route = new(entities2.HttpRoute).FillEmptyFields()

		route.Active = true
		route.Method = strings.ToUpper(r.Method)
		route.Path = r.Path
		route.RegisterTime = time.Now()

		if err = acc.repository.RegisterRoute(acc.ctx, route); err != nil {
			acc.components.Logger.Error().
				Format("Failed to register http rest api route: '%s'. ", err).Write()
			return
		}
	}

	return
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

			if err = rest_api_io.WriteError(ctx, error_list.ErrRequestBodyDataCouldNotBeRead_RestAPI()); err != nil {
				acc.components.Logger.Error().
					Format("The response could not be recorded: '%s'. ", err).Write()

				var cErr = error_list.ErrResponseCouldNotBeRecorded_RestAPI()
				cErr.SetError(err)

				return rest_api_io.WriteError(ctx, cErr)
			}

			return
		}
	}

	// Обработка
	{
		var (
			us    *entities2.User
			token *entities2.JwtToken
		)

		// Получение данных пользователя
		{
			if us, err = acc.repository.BasicAuth(ctx.Context(), requestData.Username, requestData.Password); err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					acc.components.Logger.Warn().
						Format("User authorization error: '%s'. ", err).Write()

					if err = rest_api_io.WriteError(ctx, error_list.ErrUserNotFound_RestAPI()); err != nil {
						acc.components.Logger.Error().
							Format("The response could not be recorded: '%s'. ", err).Write()

						return rest_api_io.WriteError(ctx, error_list.ErrResponseCouldNotBeRecorded_RestAPI())
					}
					return
				}

				acc.components.Logger.Error().
					Format("User authorization error: '%s'. ", err).
					Field("username", requestData.Username).
					Field("password", requestData.Password).Write()

				if err = rest_api_io.WriteError(ctx, error_list.ErrInternalServerError_RestAPI()); err != nil {
					acc.components.Logger.Error().
						Format("The response could not be recorded: '%s'. ", err).Write()

					return rest_api_io.WriteError(ctx, error_list.ErrResponseCouldNotBeRecorded_RestAPI())
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

				if err = rest_api_io.WriteError(ctx, error_list.ErrTokenNotFound_RestAPI()); err != nil {
					acc.components.Logger.Error().
						Format("The response could not be recorded: '%s'. ", err).Write()

					return rest_api_io.WriteError(ctx, error_list.ErrResponseCouldNotBeRecorded_RestAPI())
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

				if err = rest_api_io.WriteError(ctx, error_list.ErrAlreadyAuthorized_RestAPI()); err != nil {
					acc.components.Logger.Error().
						Format("The response could not be recorded: '%s'. ", err).Write()

					return rest_api_io.WriteError(ctx, error_list.ErrResponseCouldNotBeRecorded_RestAPI())
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

				if err = rest_api_io.WriteError(ctx, error_list.ErrInternalServerError_RestAPI()); err != nil {
					acc.components.Logger.Error().
						Format("The response could not be recorded: '%s'. ", err).Write()

					return rest_api_io.WriteError(ctx, error_list.ErrResponseCouldNotBeRecorded_RestAPI())
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

			return rest_api_io.WriteError(ctx, error_list.ErrResponseCouldNotBeRecorded_RestAPI())
		}
		return
	}
}
