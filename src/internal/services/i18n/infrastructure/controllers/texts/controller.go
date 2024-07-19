package texts_controller

import (
	"context"
	texts_usecase "sm-box/internal/services/i18n/infrastructure/usecases/texts"
	"sm-box/internal/services/i18n/objects/entities"
	"sm-box/internal/services/i18n/objects/models"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator = "infrastructure-[controllers]=texts"
)

// Controller - контроллер текстов локализации.
type Controller struct {
	components *components
	usecases   *usecases

	conf *Config
	ctx  context.Context
}

// usecases - логика контроллера.
type usecases struct {
	Texts interface {
		AssembleDictionary(ctx context.Context, lang string, paths []string) (dictionary entities.Dictionary, cErr c_errors.Error)
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

		// Texts
		{
			if controller.usecases.Texts, err = texts_usecase.New(ctx); err != nil {
				return
			}
		}
	}

	controller.components.Logger.Info().
		Format("A '%s' controller has been created. ", "texts").
		Field("config", controller.conf).Write()

	return
}

// AssembleDictionary - собрать словарь локализации.
func (controller *Controller) AssembleDictionary(ctx context.Context, lang string, paths []string) (dictionary models.Dictionary, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, lang, paths)
		defer func() { trc.Error(cErr).FunctionCallFinished(dictionary) }()
	}

	var dict entities.Dictionary

	// Выполнения инструкций
	{
		if dict, cErr = controller.usecases.Texts.AssembleDictionary(ctx, lang, paths); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	// Преобразование в модели
	{
		if dict != nil {
			dictionary = dict.ToModel()
		}
	}

	return
}
