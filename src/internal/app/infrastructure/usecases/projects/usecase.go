package projects_usecase

import (
	"context"
	"database/sql"
	"errors"
	projects_repository "sm-box/internal/app/infrastructure/repositories/projects"
	"sm-box/internal/app/objects/entities"
	error_list "sm-box/internal/common/errors"
	common_types "sm-box/internal/common/types"
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
		Get(ctx context.Context, ids []common_types.ID) (list entities.ProjectList, err error)
		GetOne(ctx context.Context, id common_types.ID) (project *entities.Project, err error)
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

// Get - получение проектов по ID.
func (usecase *UseCase) Get(ctx context.Context, ids ...common_types.ID) (list entities.ProjectList, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, ids)
		defer func() { trc.Error(cErr).FunctionCallFinished(list) }()
	}

	usecase.components.Logger.Info().
		Text("The process of obtaining information about projects has been launched... ").
		Field("ids", ids).Write()

	// Получение данных пользователя
	{
		if len(ids) > 0 {
			var err error

			if list, err = usecase.repositories.Projects.Get(ctx, ids); err != nil {
				usecase.components.Logger.Error().
					Format("Projects data could not be retrieved: '%s'. ", err).
					Field("ids", ids).Write()

				cErr = error_list.InternalServerError()
				cErr.SetError(err)
				return
			}

			usecase.components.Logger.Info().
				Text("The projects data has been successfully received. ").
				Field("projects", list).Write()
		}
	}

	usecase.components.Logger.Info().
		Text("Obtaining projects information is completed. ").
		Field("ids", ids).Write()

	return
}

// GetOne - получение проекта по ID.
func (usecase *UseCase) GetOne(ctx context.Context, id common_types.ID) (project *entities.Project, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(cErr).FunctionCallFinished(project) }()
	}

	usecase.components.Logger.Info().
		Text("The process of obtaining information about project has been launched... ").
		Field("id", id).Write()

	// Получение данных пользователя
	{
		var err error

		if project, err = usecase.repositories.Projects.GetOne(ctx, id); err != nil {
			project = nil

			usecase.components.Logger.Error().
				Format("Project data could not be retrieved: '%s'. ", err).
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

		usecase.components.Logger.Info().
			Text("The project data has been successfully received. ").
			Field("project", project).Write()
	}

	usecase.components.Logger.Info().
		Text("Obtaining project information is completed. ").
		Field("id", id).Write()

	return
}
