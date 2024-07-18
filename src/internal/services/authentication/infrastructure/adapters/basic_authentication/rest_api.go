package basic_authentication_adapter

import (
	"context"
	app_models "sm-box/internal/app/objects/models"
	"sm-box/internal/common/types"
	basic_authentication_controller "sm-box/internal/services/authentication/infrastructure/controllers/basic_authentication"
	"sm-box/internal/services/authentication/objects/models"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator_RestAPI = "infrastructure-[adapters]=basic_authentication-(RestAPI)"
)

// Adapter_RestAPI - адаптер контроллера для rest api.
type Adapter_RestAPI struct {
	components *components

	controller interface {
		Auth(ctx context.Context, rawSessionToken, username, password string) (sessionToken *models.JwtTokenInfo, cErr c_errors.Error)
		GetUserProjectList(ctx context.Context, rawSessionToken string) (list app_models.ProjectList, cErr c_errors.Error)
		SetTokenProject(ctx context.Context, rawSessionToken string, projectID types.ID) (
			sessionToken, accessToken, refreshToken *models.JwtTokenInfo, cErr c_errors.Error)
		Logout(ctx context.Context, rawSessionToken, rawAccessToken, rawRefreshToken string) (cErr c_errors.Error)
	}

	ctx context.Context
}

// components - компоненты адаптера.
type components struct {
	Logger logger.Logger
}

// New_RestAPI - создание контроллера для rest api.
func New_RestAPI(ctx context.Context) (adapter *Adapter_RestAPI, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelMain, tracer.LevelController)

		trace.FunctionCall(ctx)

		defer func() { trace.Error(err).FunctionCallFinished(adapter) }()
	}

	adapter = new(Adapter_RestAPI)
	adapter.ctx = ctx

	// Компоненты
	{
		adapter.components = new(components)

		// Logger
		{
			if adapter.components.Logger, err = logger.New(loggerInitiator_RestAPI); err != nil {
				return
			}
		}
	}

	// Контроллер
	{
		if adapter.controller, err = basic_authentication_controller.New(ctx); err != nil {
			return
		}
	}

	adapter.components.Logger.Info().
		Format("A '%s' adapter for RestAPI has been created. ", "basic_authentication").Write()

	return
}

// Auth - базовая авторизация пользователя в системе.
// Для авторизации используется имя пользователя и пароль.
func (adapter *Adapter_RestAPI) Auth(ctx context.Context, rawSessionToken, username, password string) (
	sessionToken *models.JwtTokenInfo, cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, rawSessionToken, username, password)
		defer func() { trc.Error(cErr).FunctionCallFinished(sessionToken) }()
	}

	var proxyErr c_errors.Error

	if sessionToken, proxyErr = adapter.controller.Auth(ctx, rawSessionToken, username, password); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// GetUserProjectList - получение списка проектов пользователя.
func (adapter *Adapter_RestAPI) GetUserProjectList(ctx context.Context, rawSessionToken string) (
	list app_models.ProjectList, cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, rawSessionToken)
		defer func() { trc.Error(cErr).FunctionCallFinished(list) }()
	}

	var proxyErr c_errors.Error

	if list, proxyErr = adapter.controller.GetUserProjectList(ctx, rawSessionToken); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// SetTokenProject - установить проект для токена.
func (adapter *Adapter_RestAPI) SetTokenProject(ctx context.Context, rawSessionToken string, projectID types.ID) (
	sessionToken, accessToken, refreshToken *models.JwtTokenInfo, cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, rawSessionToken, projectID)
		defer func() { trc.Error(cErr).FunctionCallFinished(sessionToken, accessToken, refreshToken) }()
	}

	var proxyErr c_errors.Error

	if sessionToken, accessToken, refreshToken, proxyErr = adapter.controller.SetTokenProject(ctx, rawSessionToken, projectID); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// Logout - завершение действия токена пользователя.
func (adapter *Adapter_RestAPI) Logout(ctx context.Context, rawSessionToken, rawAccessToken, rawRefreshToken string) (cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, rawSessionToken, rawAccessToken, rawRefreshToken)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if proxyErr = adapter.controller.Logout(ctx, rawSessionToken, rawAccessToken, rawRefreshToken); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}
