package basic_authentication_controller

import (
	"context"
	basic_authentication_usecase "sm-box/internal/services/users/infrastructure/usecases/basic_authentication"
	"sm-box/internal/services/users/objects/entities"
	"sm-box/internal/services/users/objects/models"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator = "infrastructure-[controllers]=basic_authentication"
)

// Controller - контроллер для аутентификации пользователей.
type Controller struct {
	components *components
	usecases   *usecases

	conf *Config
	ctx  context.Context
}

// usecases - логика контроллера.
type usecases struct {
	BasicAuthentication interface {
		Auth(ctx context.Context, username, password string) (us *entities.User, cErr c_errors.Error)
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

		// BasicAuthentication
		{
			if controller.usecases.BasicAuthentication, err = basic_authentication_usecase.New(ctx); err != nil {
				return
			}
		}
	}

	controller.components.Logger.Info().
		Format("A '%s' controller has been created. ", "basic_authentication").
		Field("config", controller.conf).Write()

	return
}

// Auth - получение информации о пользователе
// с использованием механизма базовой авторизации.
func (controller *Controller) Auth(ctx context.Context, username, password string) (user *models.UserInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, username, password)
		defer func() { trc.Error(cErr).FunctionCallFinished(user) }()
	}

	var us *entities.User

	// Выполнения инструкций
	{
		if us, cErr = controller.usecases.BasicAuthentication.Auth(ctx, username, password); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	// Преобразование в модели
	{
		if us != nil {
			user = us.ToModel()
		}
	}

	return
}
