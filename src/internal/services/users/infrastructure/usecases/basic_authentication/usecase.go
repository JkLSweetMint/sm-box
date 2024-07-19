package basic_authentication_usecase

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	error_list "sm-box/internal/common/errors"
	basic_authentication_repository "sm-box/internal/services/users/infrastructure/repositories/basic_authentication"
	"sm-box/internal/services/users/objects/entities"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator = "infrastructure-[usecases]=basic_authentication"
)

// UseCase - логика для аутентификации пользователей.
type UseCase struct {
	components   *components
	repositories *repositories

	conf *Config
	ctx  context.Context
}

// repositories - репозитории логики.
type repositories struct {
	BasicAuthentication interface {
		GetByUsername(ctx context.Context, username string) (us *entities.User, err error)
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

		// BasicAuthentication
		{
			if usecase.repositories.BasicAuthentication, err = basic_authentication_repository.New(ctx); err != nil {
				return
			}
		}
	}

	usecase.components.Logger.Info().
		Format("A '%s' usecase has been created. ", "basic_authentication").
		Field("config", usecase.conf).Write()

	return
}

// BasicAuth - получение информации о пользователе
// с использованием механизма базовой авторизации.
func (usecase *UseCase) Auth(ctx context.Context, username, password string) (us *entities.User, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, username, password)
		defer func() { trc.Error(cErr).FunctionCallFinished(us) }()
	}

	usecase.components.Logger.Info().
		Text("The process of obtaining user information has been started using the basic authorization mechanism... ").
		Field("username", username).
		Field("password", password).Write()

	// Получение данных пользователя
	{
		var err error

		if us, err = usecase.repositories.BasicAuthentication.GetByUsername(ctx, username); err != nil {
			us = nil

			usecase.components.Logger.Error().
				Format("User authorization error: '%s'. ", err).
				Field("username", username).
				Field("password", password).Write()

			if errors.Is(err, sql.ErrNoRows) {
				cErr = error_list.UserNotFound()
				cErr.SetError(err)
				return
			}

			cErr = error_list.InternalServerError()
			cErr.SetError(err)
			return
		}

		usecase.components.Logger.Info().
			Text("The user's data has been successfully received. ").
			Field("user", us).Write()
	}

	// Сравнение паролей
	{
		var (
			err                              error
			cryptoPassword1, cryptoPassword2 []byte
		)

		// cryptoPassword1
		{
			if cryptoPassword1, err = base64.StdEncoding.DecodeString(us.Password); err != nil {
				usecase.components.Logger.Error().
					Format("The decryption of the user's password failed: '%s'. ", err).
					Field("username", username).
					Field("password", password).Write()

				cErr = error_list.InternalServerError()
				cErr.SetError(err)
				return
			}

			if cryptoPassword1, err = rsa.DecryptOAEP(
				sha256.New(),
				rand.Reader,
				env.Vars.EncryptionKeys.Private,
				cryptoPassword1,
				[]byte("password")); err != nil {
				usecase.components.Logger.Error().
					Format("The decryption of the user's password failed: '%s'. ", err).
					Field("username", username).
					Field("password", us.Password).Write()

				cErr = error_list.InternalServerError()
				cErr.SetError(err)
				return
			}
		}

		// cryptoPassword2
		{
			if cryptoPassword2, err = base64.StdEncoding.DecodeString(password); err != nil {
				usecase.components.Logger.Error().
					Format("The decryption of the user's password failed: '%s'. ", err).
					Field("username", username).
					Field("password", password).Write()

				cErr = error_list.InternalServerError()
				cErr.SetError(err)
				return
			}

			if cryptoPassword2, err = rsa.DecryptOAEP(
				sha256.New(),
				rand.Reader,
				env.Vars.EncryptionKeys.Private,
				cryptoPassword2,
				[]byte("password")); err != nil {
				usecase.components.Logger.Error().
					Format("The decryption of the user's password failed: '%s'. ", err).
					Field("username", username).
					Field("password", us.Password).Write()

				cErr = error_list.InternalServerError()
				cErr.SetError(err)
				return
			}
		}

		if string(cryptoPassword1) != string(cryptoPassword2) {
			usecase.components.Logger.Error().
				Text("Authorization failed, passwords do not match. ").
				Field("username", username).Write()

			cErr = error_list.UserNotFound()
			cErr.SetError(err)
			return
		}
	}

	usecase.components.Logger.Info().
		Text("Obtaining user information using the basic authorization mechanism is completed. ").
		Field("username", username).Write()

	return
}
