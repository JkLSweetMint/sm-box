package languages_adapter

import (
	"context"
	languages_controller "sm-box/internal/services/i18n/infrastructure/controllers/languages"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
)

const (
	loggerInitiator_RestAPI = "infrastructure-[adapters]=languages-(RestAPI)"
)

// Adapter_RestAPI - адаптер контроллера для rest api.
type Adapter_RestAPI struct {
	components *components

	controller interface {
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
		if adapter.controller, err = languages_controller.New(ctx); err != nil {
			return
		}
	}

	adapter.components.Logger.Info().
		Format("A '%s' adapter for RestAPI has been created. ", "languages").Write()

	return
}
