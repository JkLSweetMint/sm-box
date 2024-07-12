package projects_usecase

import (
	"context"
	"database/sql"
	"errors"
	projects_repository "sm-box/internal/app/infrastructure/repositories/projects"
	"sm-box/internal/app/objects/entities"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	err_details "sm-box/pkg/errors/entities/details"
	err_messages "sm-box/pkg/errors/entities/messages"
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
		Get(ctx context.Context, id types.ID) (project *entities.Project, err error)
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

	// Валидация
	{
		// ID пользователя
		{
			if userID == 0 {
				if cErr == nil {
					cErr = error_list.InvalidDataWasTransmitted()
				}

				cErr.Details().SetField(
					new(err_details.FieldKey).Add("user_id"),
					new(err_messages.TextMessage).Text("Zero value. "),
				)
			}

			if cErr != nil {
				usecase.components.Logger.Error().
					Text("Invalid data was transmitted to receive projects. ").
					Field("user_id", userID).Write()

				return
			}
		}
	}

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
	}

	usecase.components.Logger.Info().
		Format("Getting the list of the user's projects has been completed successfully. The user has access to '%d' projects. ", len(list)).
		Field("user_id", userID).
		Field("projects", list).Write()

	return
}

// Get - получение проекта.
func (usecase *UseCase) Get(ctx context.Context, id types.ID) (project *entities.Project, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(cErr).FunctionCallFinished(project) }()
	}

	usecase.components.Logger.Info().
		Text("Project receipt started... ").
		Field("id", id).Write()

	// Валидация
	{
		// ID
		{
			if id == 0 {
				if cErr == nil {
					cErr = error_list.InvalidDataWasTransmitted()
				}

				cErr.Details().SetField(
					new(err_details.FieldKey).Add("id"),
					new(err_messages.TextMessage).Text("Zero value. "),
				)
			}

			if cErr != nil {
				usecase.components.Logger.Error().
					Text("Invalid data was transmitted to receive projects. ").
					Field("id", id).Write()

				return
			}
		}
	}

	// Получение
	{
		var err error

		if project, err = usecase.repositories.Projects.Get(ctx, id); err != nil {
			usecase.components.Logger.Error().
				Format("Failed to get the project: '%s'. ", err).
				Field("id", id).Write()

			if errors.Is(err, sql.ErrNoRows) {
				cErr = error_list.ProjectNotFound()
				cErr.SetError(err)
				return
			}

			cErr = error_list.InternalServerError()
			cErr.SetError(err)
			return
		}
	}

	usecase.components.Logger.Info().
		Text("The project was successfully received. ").
		Field("project", project).Write()

	return
}
