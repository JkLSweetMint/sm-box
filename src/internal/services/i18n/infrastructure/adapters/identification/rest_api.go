package identification_adapter

import (
	"context"
	common_models "sm-box/internal/common/objects/models"
	identification_controller "sm-box/internal/services/i18n/infrastructure/controllers/identification"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator_RestAPI = "infrastructure-[adapters]=identification-(RestAPI)"
)

// Adapter_RestAPI - адаптер контроллера для rest api.
type Adapter_RestAPI struct {
	components *components

	controller interface {
		GetToken(ctx context.Context, data string) (token *common_models.JwtTokenInfo, cErr c_errors.Error)
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
		if adapter.controller, err = identification_controller.New(ctx); err != nil {
			return
		}
	}

	adapter.components.Logger.Info().
		Format("A '%s' adapter for RestAPI has been created. ", "identification").Write()

	return
}

// GetToken - получение jwt токена.
func (adapter *Adapter_RestAPI) GetToken(ctx context.Context, data string) (token *common_models.JwtTokenInfo, cErr c_errors.RestAPI) {
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
