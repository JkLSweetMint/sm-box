package projects

import (
	"context"
	"errors"
	g_uuid "github.com/google/uuid"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/common/types"
	usecase_projects "sm-box/internal/system/projects_cli/infrastructure/usecases/projects"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	c_errors "sm-box/pkg/errors"
	"strconv"
)

// Controller - контроллер для управления проектами.
type Controller struct {
	usecases   *usecases
	components *components

	conf *Config
	ctx  context.Context
}

// usecases - логики контроллера.
type usecases struct {
	Projects interface {
		Create(ctx context.Context, name, description, version string) (cErr c_errors.Error)
		RemoveByID(ctx context.Context, id types.ID) (cErr c_errors.Error)
		RemoveByUUID(ctx context.Context, uuid g_uuid.UUID) (cErr c_errors.Error)
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
		var trace = tracer.New(tracer.LevelMain)

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
			if controller.components.Logger, err = logger.New(env.Vars.SystemName); err != nil {
				return
			}
		}
	}

	// Логика
	{
		controller.usecases = new(usecases)

		if controller.usecases.Projects, err = usecase_projects.New(controller.ctx); err != nil {
			return
		}
	}

	controller.components.Logger.Info().
		Format("A '%s' controller has been created. ", "projects").
		Field("config", controller.conf).Write()

	return
}

// Create - создание проекта.
func (controller *Controller) Create(ctx context.Context, name, description, version string) (cErr c_errors.Error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelController)

		trace.FunctionCall(ctx, name, description, version)

		defer func() { trace.Error(cErr).FunctionCallFinished() }()
	}

	if cErr = controller.usecases.Projects.Create(ctx, name, description, version); cErr != nil {
		controller.components.Logger.Error().
			Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// Remove - удаление проекта.
func (controller *Controller) Remove(ctx context.Context, id string) (cErr c_errors.Error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelController)

		trace.FunctionCall(ctx, id)

		defer func() { trace.Error(cErr).FunctionCallFinished() }()
	}

	if v, err := g_uuid.Parse(id); err == nil {
		if cErr = controller.usecases.Projects.RemoveByUUID(ctx, v); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()
			return
		}
	} else if v, err := strconv.Atoi(id); err == nil {
		if cErr = controller.usecases.Projects.RemoveByID(ctx, types.ID(v)); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()
			return
		}
	} else {
		cErr = error_list.FailedRemoveProject()
		cErr.SetError(errors.New("Invalid identifier value. "))
		cErr.Details().Set("id", "Invalid value. ")

		controller.components.Logger.Error().
			Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

		return
	}

	return
}
