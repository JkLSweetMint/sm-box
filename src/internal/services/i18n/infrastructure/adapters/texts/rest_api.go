package texts_adapter

import (
	"context"
	texts_controller "sm-box/internal/services/i18n/infrastructure/controllers/texts"
	"sm-box/internal/services/i18n/infrastructure/objects/models"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator_RestAPI = "infrastructure-[adapters]=texts-(RestAPI)"
)

// Adapter_RestAPI - адаптер контроллера для rest api.
type Adapter_RestAPI struct {
	components *components

	controller interface {
		AssembleDictionary(ctx context.Context, lang string, paths []string) (dictionary models.Dictionary, cErr c_errors.Error)
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
		if adapter.controller, err = texts_controller.New(ctx); err != nil {
			return
		}
	}

	adapter.components.Logger.Info().
		Format("A '%s' adapter for RestAPI has been created. ", "texts").Write()

	return
}

// AssembleDictionary - собрать словарь локализации.
func (adapter *Adapter_RestAPI) AssembleDictionary(ctx context.Context, lang string, paths []string) (dictionary models.Dictionary, cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, lang, paths)
		defer func() { trc.Error(cErr).FunctionCallFinished(dictionary) }()
	}

	var proxyErr c_errors.Error

	if dictionary, proxyErr = adapter.controller.AssembleDictionary(ctx, lang, paths); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}
