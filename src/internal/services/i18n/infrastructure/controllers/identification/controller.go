package identification_controller

import (
	"context"
	common_entities "sm-box/internal/common/objects/entities"
	common_models "sm-box/internal/common/objects/models"
	identification_usecase "sm-box/internal/services/i18n/infrastructure/usecases/identification"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator = "infrastructure-[controllers]=identification"
)

// Controller - контроллер идентификации пользователей.
type Controller struct {
	components *components
	usecases   *usecases

	conf *Config
	ctx  context.Context
}

// usecases - логика контроллера.
type usecases struct {
	Identification interface {
		GetToken(ctx context.Context, data string) (tok *common_entities.JwtToken, cErr c_errors.Error)
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

		// Identification
		{
			if controller.usecases.Identification, err = identification_usecase.New(ctx); err != nil {
				return
			}
		}
	}

	controller.components.Logger.Info().
		Format("A '%s' controller has been created. ", "identification").
		Field("config", controller.conf).Write()

	return
}

// GetToken - получение jwt токена.
func (controller *Controller) GetToken(ctx context.Context, data string) (token *common_models.JwtTokenInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, data)
		defer func() { trc.Error(cErr).FunctionCallFinished(token) }()
	}

	var tok *common_entities.JwtToken

	if tok, cErr = controller.usecases.Identification.GetToken(ctx, data); cErr != nil {
		controller.components.Logger.Error().
			Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

		return
	}

	// Преобразование в модели
	{
		if tok != nil {
			token = tok.Model()
		}
	}

	return
}
