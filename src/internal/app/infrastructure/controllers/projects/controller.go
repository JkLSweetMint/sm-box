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
		GetListByUser(ctx context.Context, userID types.ID) (list entities.ProjectList, cErr c_errors.Error)
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

// GetListByUser - получение списка проектов пользователя.
func (controller *Controller) GetListByUser(ctx context.Context, userID types.ID) (list models.ProjectList, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, userID)
		defer func() { trc.Error(cErr).FunctionCallFinished(list) }()
	}

	var list_ entities.ProjectList

	if list_, cErr = controller.usecases.Projects.GetListByUser(ctx, userID); cErr != nil {
		controller.components.Logger.Error().
			Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

		return
	}

	// Преобразование в модели
	{
		if list_ != nil {
			list = list_.ToModel()
		}
	}

	return
}
