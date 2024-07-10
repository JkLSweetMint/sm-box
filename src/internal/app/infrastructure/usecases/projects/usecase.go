package projects_usecase

import (
	"context"
	projects_repository "sm-box/internal/app/infrastructure/repositories/projects"
	"sm-box/internal/app/objects/entities"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator = "infrastructure-[usecases]=projects"
)

// UseCase - логика проектов системы.
type UseCase struct {
	components   *components
	repositories *repositories

	conf *Config
	ctx  context.Context
}

// repositories - репозитории логики.
type repositories struct {
	Projects interface {
		GetListByUser(ctx context.Context, userID types.ID) (list entities.ProjectList, err error)
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

		// Projects
		{
			if usecase.repositories.Projects, err = projects_repository.New(ctx); err != nil {
				return
			}
		}
	}

	usecase.components.Logger.Info().
		Format("A '%s' usecase has been created. ", "projects").
		Field("config", usecase.conf).Write()

	return
}

// GetListByUser - получение списка проектов пользователя.
func (usecase *UseCase) GetListByUser(ctx context.Context, userID types.ID) (list entities.ProjectList, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, userID)
		defer func() { trc.Error(cErr).FunctionCallFinished(list) }()
	}

	usecase.components.Logger.Info().
		Text("The list of user's projects has been started... ").
		Field("user_id", userID).Write()

	// Получение проектов
	{
		var err error

		if list, err = usecase.repositories.Projects.GetListByUser(ctx, userID); err != nil {
			usecase.components.Logger.Error().
				Format("The list of user's projects could not be retrieved: '%s'. ", err).
				Field("user_id", userID).Write()

			cErr = error_list.ListUserProjectsCouldNotBeRetrieved()
			cErr.SetError(err)
			return
		}

		usecase.components.Logger.Info().
			Format("The user has access to '%d' projects ", len(list)).
			Field("user_id", userID).
			Field("projects", list).Write()

	}

	usecase.components.Logger.Info().
		Text("Getting the list of the user's projects has been completed successfully. ").
		Field("user_id", userID).Write()

	return
}
