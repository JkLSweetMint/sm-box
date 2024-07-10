package authentication_controller

import (
	"context"
	app_models "sm-box/internal/app/objects/models"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/common/types"
	authentication_usecase "sm-box/internal/services/authentication/infrastructure/usecases/authentication"
	"sm-box/internal/services/authentication/objects/entities"
	"sm-box/internal/services/authentication/objects/models"
	projects_service_gateway "sm-box/internal/services/authentication/transport/gateways/grpc/projects_service"
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
	gateways   *gateways
	usecases   *usecases

	conf *Config
	ctx  context.Context
}

// gateways - шлюзы контроллера.
type gateways struct {
	Projects interface {
		GetListByUser(ctx context.Context, userID types.ID) (list app_models.ProjectList, cErr c_errors.Error)
	}
}

// usecases - логика контроллера.
type usecases struct {
	Authentication interface {
		BasicAuth(ctx context.Context, rawToken, username, password string) (token *entities.JwtToken, cErr c_errors.Error)
		SetTokenProject(ctx context.Context, rawToken string, projectID types.ID) (token *entities.JwtToken, cErr c_errors.Error)
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

	// Шлюзы
	{
		controller.gateways = new(gateways)

		// Authentication
		{
			if controller.gateways.Projects, err = projects_service_gateway.New(ctx); err != nil {
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
func (controller *Controller) GetUserProjectList(ctx context.Context, userID types.ID) (list app_models.ProjectList, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, userID)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var err error

	if list, err = controller.gateways.Projects.GetListByUser(ctx, userID); err != nil {
		cErr = error_list.InternalServerError()
		cErr.SetError(err)
		return
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
