package authentication_usecase

import (
	"context"
	"database/sql"
	"errors"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/common/objects/entities"
	"sm-box/internal/common/types"
	authentication_repository "sm-box/internal/services/authentication/infrastructure/repositories/authentication"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator = "infrastructure-[usecases]=authentication"
)

// UseCase - логика аутентификации пользователей.
type UseCase struct {
	components   *components
	repositories *repositories

	conf *Config
	ctx  context.Context
}

// repositories - репозитории логики.
type repositories struct {
	Authentication interface {
		GetToken(ctx context.Context, data string) (tok *entities.JwtToken, err error)
		SetTokenOwner(ctx context.Context, tokenID, ownerID types.ID) (err error)
		BasicAuth(ctx context.Context, username, password string) (us *entities.User, err error)
	}
}

// components - компоненты логики.
type components struct {
	Logger logger.Logger
}

// New - создание логики.
func New(ctx context.Context) (usecase *UseCase, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelMain, tracer.LevelUseCase)

		trace.FunctionCall(ctx)

		defer func() { trace.Error(err).FunctionCallFinished(usecase) }()
	}

	usecase = new(UseCase)
	usecase.ctx = ctx

	// Конфигурация
	{
		usecase.conf = new(Config).Default()

		if err = usecase.conf.Read(); err != nil {
			return
		}
	}

	// Компоненты
	{
		usecase.components = new(components)

		// Logger
		{
			if usecase.components.Logger, err = logger.New(loggerInitiator); err != nil {
				return
			}
		}
	}

	// Репозитории
	{
		usecase.repositories = new(repositories)

		// Authentication
		{
			if usecase.repositories.Authentication, err = authentication_repository.New(ctx); err != nil {
				return
			}
		}
	}

	usecase.components.Logger.Info().
		Format("A '%s' usecase has been created. ", "authentication").
		Field("config", usecase.conf).Write()

	return
}

// BasicAuth - базовая авторизация пользователя в системе.
// Для авторизации используется имя пользователя и пароль.
func (usecase *UseCase) BasicAuth(ctx context.Context, tokenData, username, password string) (tok *entities.JwtToken, us *entities.User, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, tokenData, username, password)
		defer func() { trc.Error(cErr).FunctionCallFinished(tok, us) }()
	}

	// Проверка пароля
	{
		//var (
		//	err        error
		//	passwordDB []byte
		//)
		//
		//if passwordDB, err = rsa.EncryptOAEP(
		//	sha256.New(),
		//	rand.Reader,
		//	env.Vars.EncryptionKeys.Public,
		//	[]byte(password),
		//	[]byte("password")); err != nil {
		//	usecase.components.Logger.Error().
		//		Format("The decryption of the user's password failed: '%s'. ", err).
		//		Field("username", username).
		//		Field("password", password).Write()
		//
		//	cErr = error_list.InternalServerError()
		//	cErr.SetError(err)
		//	return
		//}
		//
		//password = string(passwordDB)
	}

	// Получение данных пользователя
	{
		var err error

		if us, err = usecase.repositories.Authentication.BasicAuth(ctx, username, password); err != nil {
			us = nil

			if errors.Is(err, sql.ErrNoRows) {
				usecase.components.Logger.Warn().
					Format("User authorization error: '%s'. ", err).Write()

				cErr = error_list.UserNotFound()
				cErr.SetError(err)

				return
			}

			usecase.components.Logger.Error().
				Format("User authorization error: '%s'. ", err).
				Field("username", username).
				Field("password", password).Write()

			cErr = error_list.InternalServerError()
			cErr.SetError(err)

			return
		}
	}

	// Получение токена
	{
		var err error

		if tok, err = usecase.repositories.Authentication.GetToken(ctx, tokenData); err != nil {
			usecase.components.Logger.Error().
				Format("Failed to get token data: '%s'. ", err).
				Field("data", tokenData).Write()

			cErr = error_list.TokenNotFound()
			cErr.SetError(err)

			return
		}
	}

	// Проверка что уже авторизован
	{
		if tok.UserID != 0 {
			usecase.components.Logger.Warn().
				Text("The user is already logged in. ").
				Field("user_id", us.ID).
				Field("token", tok).Write()

			cErr = error_list.AlreadyAuthorized()

			return
		}
	}

	// Обновление данных токена
	{
		var err error

		if err = usecase.repositories.Authentication.SetTokenOwner(ctx, tok.ID, us.ID); err != nil {
			usecase.components.Logger.Error().
				Format("The token owner could not be identified: '%s'. ", err).
				Field("owner_id", us.ID).
				Field("token", tok).Write()

			cErr = error_list.InternalServerError()
			cErr.SetError(err)

			return
		}
	}

	tok.UserID = us.ID

	return
}
