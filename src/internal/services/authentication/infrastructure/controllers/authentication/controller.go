package authentication_controller

import (
	"context"
	app_entities "sm-box/internal/app/objects/entities"
	app_models "sm-box/internal/app/objects/models"
	"sm-box/internal/common/types"
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
		SetTokenProject(ctx context.Context, rawToken string, projectID types.ID) (token *entities.JwtToken, cErr c_errors.Error)
		GetUserProjectList(ctx context.Context, rawToken string) (list app_entities.ProjectList, cErr c_errors.Error)
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

// GetUserProjectList - получение списка проектов пользователя.
func (controller *Controller) GetUserProjectList(ctx context.Context, rawToken string) (list app_models.ProjectList, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, rawToken)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var projects app_entities.ProjectList

	if projects, cErr = controller.usecases.Authentication.GetUserProjectList(ctx, rawToken); cErr != nil {
		controller.components.Logger.Error().
			Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

		return
	}

	list = make(app_models.ProjectList, 0, len(list))

	// Преобразование в модели
	{
		if projects != nil {
			for _, project := range projects {
				list = append(list, project.ToModel())
			}
		}
	}

	return
}

// SetTokenProject - установить проект для токена.
func (controller *Controller) SetTokenProject(ctx context.Context, rawToken string, projectID types.ID) (token *entities.JwtToken, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, rawToken, projectID)
		defer func() { trc.Error(cErr).FunctionCallFinished(token) }()
	}

	if token, cErr = controller.usecases.Authentication.SetTokenProject(ctx, rawToken, projectID); cErr != nil {
		controller.components.Logger.Error().
			Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

		return
	}

	return
}
