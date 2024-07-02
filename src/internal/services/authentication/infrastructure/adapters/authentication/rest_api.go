package authentication_adapter

import (
	"context"
	"sm-box/internal/common/objects/models"
	"sm-box/internal/common/types"
	authentication_controller "sm-box/internal/services/authentication/infrastructure/controllers/authentication"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator_RestAPI = "infrastructure-[adapters]=authentication-(RestAPI)"
)

// Adapter_RestAPI - адаптер контроллера для rest api.
type Adapter_RestAPI struct {
	components *components

	controller interface {
		BasicAuth(ctx context.Context, tokenData, username, password string) (token *models.JwtTokenInfo, user *models.UserInfo, cErr c_errors.Error)
		SetTokenProject(ctx context.Context, tokenID, projectID types.ID) (cErr c_errors.Error)
		GetUserProjectsList(ctx context.Context, tokenID, userID types.ID) (list models.ProjectList, cErr c_errors.Error)
		GetToken(ctx context.Context, data string) (token *models.JwtTokenInfo, cErr c_errors.Error)
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
		if adapter.controller, err = authentication_controller.New(ctx); err != nil {
			return
		}
	}

	adapter.components.Logger.Info().
		Format("A '%s' adapter for RestAPI has been created. ", "authentication").Write()

	return
}

// BasicAuth - базовая авторизация пользователя в системе.
// Для авторизации используется имя пользователя и пароль.
func (adapter *Adapter_RestAPI) BasicAuth(ctx context.Context, tokenData, username, password string) (token *models.JwtTokenInfo, user *models.UserInfo, cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, tokenData, username, password)
		defer func() { trc.Error(cErr).FunctionCallFinished(token, user) }()
	}

	var proxyErr c_errors.Error

	if token, user, proxyErr = adapter.controller.BasicAuth(ctx, tokenData, username, password); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// SetTokenProject - установить проект для токена.
func (adapter *Adapter_RestAPI) SetTokenProject(ctx context.Context, tokenID, projectID types.ID) (cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, tokenID, projectID)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if proxyErr = adapter.controller.SetTokenProject(ctx, tokenID, projectID); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// GetUserProjectsList - получение списка проектов пользователя.
func (adapter *Adapter_RestAPI) GetUserProjectsList(ctx context.Context, tokenID, userID types.ID) (list models.ProjectList, cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, tokenID, userID)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if list, proxyErr = adapter.controller.GetUserProjectsList(ctx, tokenID, userID); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// GetToken - получение jwt токена.
func (adapter *Adapter_RestAPI) GetToken(ctx context.Context, data string) (token *models.JwtTokenInfo, cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, data)
		defer func() { trc.Error(cErr).FunctionCallFinished(token) }()
	}

	var proxyErr c_errors.Error

	if token, proxyErr = adapter.controller.GetToken(ctx, data); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}
