package users_usecase

import (
	"context"
	"database/sql"
	"errors"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/common/types"
	"sm-box/internal/services/users/infrastructure/repositories/users"
	"sm-box/internal/services/users/objects/entities"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator = "infrastructure-[usecases]=users"
)

// UseCase - логика для работы с пользователями.
type UseCase struct {
	components   *components
	repositories *repositories

	conf *Config
	ctx  context.Context
}

// repositories - репозитории логики.
type repositories struct {
	Users interface {
		Get(ctx context.Context, ids []types.ID) (list []*entities.User, err error)
		GetOne(ctx context.Context, id types.ID) (us *entities.User, err error)
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
			if usecase.repositories.Users, err = users_repository.New(ctx); err != nil {
				return
			}
		}
	}

	usecase.components.Logger.Info().
		Format("A '%s' usecase has been created. ", "users").
		Field("config", usecase.conf).Write()

	return
}

// Get - получение пользователей по ID.
func (usecase *UseCase) Get(ctx context.Context, ids ...types.ID) (list []*entities.User, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, ids)
		defer func() { trc.Error(cErr).FunctionCallFinished(list) }()
	}

	usecase.components.Logger.Info().
		Text("The process of obtaining users information has been started.... ").
		Field("ids", ids).Write()

	// Получение данных пользователя
	{
		if len(ids) > 0 {
			var err error

			if list, err = usecase.repositories.Users.Get(ctx, ids); err != nil {
				usecase.components.Logger.Error().
					Format("Users data could not be retrieved: '%s'. ", err).
					Field("ids", ids).Write()

				cErr = error_list.InternalServerError()
				cErr.SetError(err)
				return
			}

			usecase.components.Logger.Info().
				Text("The user's data has been successfully received. ").
				Field("users", list).Write()
		}
	}

	usecase.components.Logger.Info().
		Text("Obtaining users information is completed. ").
		Field("ids", ids).Write()

	return
}

// GetOne - получение пользователя по ID.
func (usecase *UseCase) GetOne(ctx context.Context, id types.ID) (us *entities.User, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(cErr).FunctionCallFinished(us) }()
	}

	usecase.components.Logger.Info().
		Text("The process of obtaining user information has been started .... ").
		Field("id", id).Write()

	// Получение данных пользователя
	{
		var err error

		if us, err = usecase.repositories.Users.GetOne(ctx, id); err != nil {
			us = nil

			usecase.components.Logger.Error().
				Format("User data could not be retrieved: '%s'. ", err).
				Field("id", id).Write()

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

	usecase.components.Logger.Info().
		Text("Obtaining user information is completed. ").
		Field("id", id).Write()

	return
}
