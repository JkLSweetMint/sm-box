package basic_authentication_controller

import (
	"context"
	app_entities "sm-box/internal/app/objects/entities"
	app_models "sm-box/internal/app/objects/models"
	"sm-box/internal/common/types"
	basic_authentication_usecase "sm-box/internal/services/authentication/infrastructure/usecases/basic_authentication"
	"sm-box/internal/services/authentication/objects/entities"
	"sm-box/internal/services/authentication/objects/models"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator = "infrastructure-[controllers]=basic_authentication"
)

// Controller - контроллер базовой аутентификации пользователей.
type Controller struct {
	components *components
	usecases   *usecases

	conf *Config
	ctx  context.Context
}

// usecases - логика контроллера.
type usecases struct {
	BasicAuthentication interface {
		Auth(ctx context.Context, rawSessionToken, username, password string) (sessionToken *entities.JwtSessionToken, cErr c_errors.Error)
		GetUserProjectList(ctx context.Context, rawSessionToken string) (list app_entities.ProjectList, cErr c_errors.Error)
		SetTokenProject(ctx context.Context, rawSessionToken string, projectID types.ID) (
			sessionToken *entities.JwtSessionToken,
			accessToken *entities.JwtAccessToken,
			refreshToken *entities.JwtRefreshToken,
			cErr c_errors.Error)
		Logout(ctx context.Context, rawSessionToken, rawAccessToken, rawRefreshToken string) (cErr c_errors.Error)
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

// Auth - базовая авторизация пользователя в системе.
// Для авторизации используется имя пользователя и пароль.
func (controller *Controller) Auth(ctx context.Context, rawSessionToken, username, password string) (
	sessionToken *models.JwtTokenInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, rawSessionToken, username, password)
		defer func() { trc.Error(cErr).FunctionCallFinished(sessionToken) }()
	}

	var sessionTok *entities.JwtSessionToken

	if sessionTok, cErr = controller.usecases.BasicAuthentication.Auth(ctx, rawSessionToken, username, password); cErr != nil {
		controller.components.Logger.Error().
			Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

		return
	}

	// Преобразование в модели
	{
		if sessionTok != nil {
			sessionToken = sessionTok.ToModel()
		}
	}

	return
}

// GetUserProjectList - получение списка проектов пользователя.
func (controller *Controller) GetUserProjectList(ctx context.Context, rawSessionToken string) (
	list app_models.ProjectList, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, rawSessionToken)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var projects app_entities.ProjectList

	if projects, cErr = controller.usecases.BasicAuthentication.GetUserProjectList(ctx, rawSessionToken); cErr != nil {
		controller.components.Logger.Error().
			Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

		return
	}

	// Преобразование в модели
	{
		list = make(app_models.ProjectList, 0, len(list))

		if projects != nil {
			for _, project := range projects {
				list = append(list, project.ToModel())
			}
		}
	}

	return
}

// SetTokenProject - установить проект для токена.
func (controller *Controller) SetTokenProject(ctx context.Context, rawSessionToken string, projectID types.ID) (
	sessionToken, accessToken, refreshToken *models.JwtTokenInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, rawSessionToken, projectID)
		defer func() { trc.Error(cErr).FunctionCallFinished(sessionToken, accessToken, refreshToken) }()
	}

	var (
		sessionTok *entities.JwtSessionToken
		accessTok  *entities.JwtAccessToken
		refreshTok *entities.JwtRefreshToken
	)

	if sessionTok, accessTok, refreshTok, cErr = controller.usecases.BasicAuthentication.SetTokenProject(ctx, rawSessionToken, projectID); cErr != nil {
		controller.components.Logger.Error().
			Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

		return
	}

	// Преобразование в модели
	{
		if sessionTok != nil {
			sessionToken = sessionTok.ToModel()
		}

		if accessTok != nil {
			accessToken = accessTok.ToModel()
		}

		if refreshTok != nil {
			refreshToken = refreshTok.ToModel()
		}
	}

	return
}

// Logout - завершение действия токена пользователя.
func (controller *Controller) Logout(ctx context.Context, rawSessionToken, rawAccessToken, rawRefreshToken string) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, rawSessionToken, rawAccessToken, rawRefreshToken)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	if cErr = controller.usecases.BasicAuthentication.Logout(ctx, rawSessionToken, rawAccessToken, rawRefreshToken); cErr != nil {
		controller.components.Logger.Error().
			Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

		return
	}

	return
}
