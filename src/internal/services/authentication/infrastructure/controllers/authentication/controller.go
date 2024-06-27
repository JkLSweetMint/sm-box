package authentication_controller

import (
	"context"
	"sm-box/internal/common/objects/entities"
	"sm-box/internal/common/objects/models"
	authentication_usecase "sm-box/internal/services/authentication/infrastructure/usecases/authentication"
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
		BasicAuth(ctx context.Context, tokenData, username, password string) (token *entities.JwtToken, us *entities.User, cErr c_errors.Error)
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
func (controller *Controller) BasicAuth(ctx context.Context, tokenData, username, password string) (token *models.JwtTokenInfo, user *models.UserInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, tokenData, username, password)
		defer func() { trc.Error(cErr).FunctionCallFinished(token, user) }()
	}

	var (
		tok *entities.JwtToken
		us  *entities.User
	)

	if tok, us, cErr = controller.usecases.Authentication.BasicAuth(ctx, tokenData, username, password); cErr != nil {
		controller.components.Logger.Error().
			Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()
		return
	}

	// Преобразование в модели
	{
		if tok != nil {
			token = tok.Model()
		}

		if us != nil {
			user = us.Model()
		}
	}

	return
}
