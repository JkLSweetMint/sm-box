package authentication_controller

import (
	"context"
	authentication_usecase "sm-box/internal/services/authentication/infrastructure/usecases/authentication"
	"sm-box/internal/services/authentication/objects/entities"
	"sm-box/internal/services/authentication/objects/models"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator = "infrastructure-[controllers]=authentication"
)

// Controller - контроллер аутентификации пользователей.
type Controller struct {
	components *components
	usecases   *usecases

	conf *Config
	ctx  context.Context
}

// usecases - логика контроллера.
type usecases struct {
	Authentication interface {
		BasicAuth(ctx context.Context, rawToken, username, password string) (token *entities.JwtToken, cErr c_errors.Error)
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

		// Authentication
		{
			if controller.usecases.Authentication, err = authentication_usecase.New(ctx); err != nil {
				return
			}
		}
	}

	controller.components.Logger.Info().
		Format("A '%s' controller has been created. ", "authentication").
		Field("config", controller.conf).Write()

	return
}

// BasicAuth - базовая авторизация пользователя в системе.
// Для авторизации используется имя пользователя и пароль.
func (controller *Controller) BasicAuth(ctx context.Context, rawToken, username, password string) (token *models.JwtTokenInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, rawToken, username, password)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var tok *entities.JwtToken

	if tok, cErr = controller.usecases.Authentication.BasicAuth(ctx, rawToken, username, password); cErr != nil {
		controller.components.Logger.Error().
			Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

		return
	}

	// Преобразование в модели
	{
		if tok != nil {
			token = tok.ToModel()
		}
	}

	return
}
