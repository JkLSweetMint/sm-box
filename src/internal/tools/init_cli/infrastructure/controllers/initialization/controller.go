package initialization

import (
	"context"
	usecase_initialization "sm-box/internal/tools/init_cli/infrastructure/usecases/initialization"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	c_errors "sm-box/pkg/errors"
)

// Controller - контроллер инициализации.
type Controller struct {
	usecases   *usecases
	components *components

	conf *Config
	ctx  context.Context
}

// usecases - логики контроллера.
type usecases struct {
	Initialization interface {
		Initialize(ctx context.Context) (cErr c_errors.Error)
		Clear(ctx context.Context) (cErr c_errors.Error)
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
		controller.conf = new(Config)

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

		if controller.usecases.Initialization, err = usecase_initialization.New(controller.ctx); err != nil {
			return
		}
	}

	controller.components.Logger.Info().
		Format("A '%s' controller has been created. ", "initialization").
		Field("config", controller.conf).Write()

	return
}

// Initialize - инициализировать систему.
func (controller *Controller) Initialize(ctx context.Context) (cErr c_errors.Error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelController)

		trace.FunctionCall(ctx)

		defer func() { trace.Error(cErr).FunctionCallFinished() }()
	}

	if cErr = controller.usecases.Initialization.Initialize(ctx); cErr != nil {
		controller.components.Logger.Error().
			Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// Clear - очистить систему.
func (controller *Controller) Clear(ctx context.Context) (cErr c_errors.Error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelController)

		trace.FunctionCall(ctx)

		defer func() { trace.Error(cErr).FunctionCallFinished() }()
	}

	if cErr = controller.usecases.Initialization.Clear(ctx); cErr != nil {
		controller.components.Logger.Error().
			Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}
