package projects_controller

import (
	"context"
	projects_usecase "sm-box/internal/app/infrastructure/usecases/projects"
	"sm-box/internal/app/objects/entities"
	"sm-box/internal/app/objects/models"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator = "infrastructure-[controllers]=projects"
)

// Controller - контроллер проектов системы.
type Controller struct {
	components *components
	usecases   *usecases

	conf *Config
	ctx  context.Context
}

// usecases - логика контроллера.
type usecases struct {
	Projects interface {
		Get(ctx context.Context, ids ...types.ID) (list entities.ProjectList, cErr c_errors.Error)
		GetOne(ctx context.Context, id types.ID) (project *entities.Project, cErr c_errors.Error)
	}
}

// components - компоненты контроллера.
type components struct {
	Logger logger.Logger
}

// New - создание контроллера.
func New(ctx context.Context) (controller *Controller, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelMain, tracer.LevelController)

		trace.FunctionCall(ctx)

		defer func() { trace.Error(err).FunctionCallFinished(controller) }()
	}

	controller = new(Controller)
	controller.ctx = ctx

	// Конфигурация
	{
		controller.conf = new(Config).Default()

		if err = controller.conf.Read(); err != nil {
			return
		}
	}

	// Компоненты
	{
		controller.components = new(components)

		// Logger
		{
			if controller.components.Logger, err = logger.New(loggerInitiator); err != nil {
				return
			}
		}
	}

	// Логика
	{
		controller.usecases = new(usecases)

		// Projects
		{
			if controller.usecases.Projects, err = projects_usecase.New(ctx); err != nil {
				return
			}
		}
	}

	controller.components.Logger.Info().
		Format("A '%s' controller has been created. ", "projects").
		Field("config", controller.conf).Write()

	return
}

// Get - получение проектов по ID.
func (controller *Controller) Get(ctx context.Context, ids ...types.ID) (list models.ProjectList, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, ids)
		defer func() { trc.Error(cErr).FunctionCallFinished(list) }()
	}

	var projects []*entities.Project

	if projects, cErr = controller.usecases.Projects.Get(ctx, ids...); cErr != nil {
		controller.components.Logger.Error().
			Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

		return
	}

	// Преобразование в модели
	{
		list = make(models.ProjectList, 0, len(projects))

		if projects != nil {
			for _, project := range projects {
				list = append(list, project.ToModel())
			}
		}
	}

	return
}

// GetOne - получение проекта по ID.
func (controller *Controller) GetOne(ctx context.Context, id types.ID) (project *models.ProjectInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(cErr).FunctionCallFinished(project) }()
	}

	var proj *entities.Project

	if proj, cErr = controller.usecases.Projects.GetOne(ctx, id); cErr != nil {
		controller.components.Logger.Error().
			Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

		return
	}

	// Преобразование в модель
	{
		if proj != nil {
			project = proj.ToModel()
		}
	}

	return
}
