package languages_controller

import (
	"context"
	languages_usecase "sm-box/internal/services/i18n/infrastructure/usecases/languages"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
)

const (
	loggerInitiator = "infrastructure-[controllers]=languages"
)

// Controller - контроллер языков локализации.
type Controller struct {
	components *components
	usecases   *usecases

	conf *Config
	ctx  context.Context
}

// usecases - логика контроллера.
type usecases struct {
	Languages interface {
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

		// Languages
		{
			if controller.usecases.Languages, err = languages_usecase.New(ctx); err != nil {
				return
			}
		}
	}

	controller.components.Logger.Info().
		Format("A '%s' controller has been created. ", "languages").
		Field("config", controller.conf).Write()

	return
}
