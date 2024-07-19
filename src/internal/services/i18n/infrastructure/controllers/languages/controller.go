package languages_controller

import (
	"context"
	languages_usecase "sm-box/internal/services/i18n/infrastructure/usecases/languages"
	"sm-box/internal/services/i18n/objects/entities"
	"sm-box/internal/services/i18n/objects/models"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
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
		GetList(ctx context.Context) (list []*entities.Language, cErr c_errors.Error)
		Remove(ctx context.Context, code string) (cErr c_errors.Error)
		Update(ctx context.Context, code, name string) (cErr c_errors.Error)
		Create(ctx context.Context, code string, name string) (cErr c_errors.Error)

		Activate(ctx context.Context, code string) (cErr c_errors.Error)
		Deactivate(ctx context.Context, code string) (cErr c_errors.Error)
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

// GetList - получение списка языков.
func (controller *Controller) GetList(ctx context.Context) (list []*models.Language, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(cErr).FunctionCallFinished(list) }()
	}

	var languages []*entities.Language

	// Выполнения инструкций
	{
		if languages, cErr = controller.usecases.Languages.GetList(ctx); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	// Преобразование в модели
	{
		list = make([]*models.Language, 0, len(list))

		if languages != nil {
			for _, lang := range languages {
				list = append(list, lang.ToModel())
			}
		}
	}

	return
}

// Remove - удаление языка.
// Текста и ресурсы локализации так же удаляются.
func (controller *Controller) Remove(ctx context.Context, code string) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, code)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	// Выполнения инструкций
	{
		if cErr = controller.usecases.Languages.Remove(ctx, code); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	return
}

// Update - обновление данных языка.
func (controller *Controller) Update(ctx context.Context, code, name string) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, code, name)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	// Выполнения инструкций
	{
		if cErr = controller.usecases.Languages.Update(ctx, code, name); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	return
}

// Create - создание языка.
func (controller *Controller) Create(ctx context.Context, code string, name string) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, code, name)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	// Выполнения инструкций
	{
		if cErr = controller.usecases.Languages.Create(ctx, code, name); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	return
}

// Activate - активировать язык.
func (controller *Controller) Activate(ctx context.Context, code string) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, code)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	// Выполнения инструкций
	{
		if cErr = controller.usecases.Languages.Activate(ctx, code); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	return
}

// Deactivate - деактивировать язык.
func (controller *Controller) Deactivate(ctx context.Context, code string) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, code)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	// Выполнения инструкций
	{
		if cErr = controller.usecases.Languages.Activate(ctx, code); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	return
}
