package languages_adapter

import (
	"context"
	languages_controller "sm-box/internal/services/i18n/infrastructure/controllers/languages"
	"sm-box/internal/services/i18n/objects/models"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator_HttpRestAPI = "infrastructure-[adapters]=languages-(HttpRestAPI)"
)

// Adapter_HttpRestAPI - адаптер контроллера для http rest api.
type Adapter_HttpRestAPI struct {
	components *components

	controller interface {
		GetList(ctx context.Context) (list []*models.Language, cErr c_errors.Error)
		Remove(ctx context.Context, code string) (cErr c_errors.Error)
		Update(ctx context.Context, code, name string) (cErr c_errors.Error)
		Create(ctx context.Context, code string, name string) (cErr c_errors.Error)

		Activate(ctx context.Context, code string) (cErr c_errors.Error)
		Deactivate(ctx context.Context, code string) (cErr c_errors.Error)
	}

	ctx context.Context
}

// components - компоненты адаптера.
type components struct {
	Logger logger.Logger
}

// New_HttpRestAPI - создание адаптера контроллера для http rest api.
func New_HttpRestAPI(ctx context.Context) (adapter *Adapter_HttpRestAPI, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelMain, tracer.LevelController)

		trace.FunctionCall(ctx)

		defer func() { trace.Error(err).FunctionCallFinished(adapter) }()
	}

	adapter = new(Adapter_HttpRestAPI)
	adapter.ctx = ctx

	// Компоненты
	{
		adapter.components = new(components)

		// Logger
		{
			if adapter.components.Logger, err = logger.New(loggerInitiator_HttpRestAPI); err != nil {
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

// GetList - получение списка языков.
func (adapter *Adapter_HttpRestAPI) GetList(ctx context.Context) (list []*models.Language, cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(cErr).FunctionCallFinished(list) }()
	}

	var proxyErr c_errors.Error

	if list, proxyErr = adapter.controller.GetList(ctx); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// Remove - удаление языка.
// Текста и ресурсы локализации так же удаляются.
func (adapter *Adapter_HttpRestAPI) Remove(ctx context.Context, code string) (cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, code)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if proxyErr = adapter.controller.Remove(ctx, code); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// Update - обновление данных языка.
func (adapter *Adapter_HttpRestAPI) Update(ctx context.Context, code, name string) (cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, code, name)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if proxyErr = adapter.controller.Update(ctx, code, name); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// Create - создание языка.
func (adapter *Adapter_HttpRestAPI) Create(ctx context.Context, code string, name string) (cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, code, name)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if proxyErr = adapter.controller.Create(ctx, code, name); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// Activate - активировать язык.
func (adapter *Adapter_HttpRestAPI) Activate(ctx context.Context, code string) (cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, code)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if proxyErr = adapter.controller.Activate(ctx, code); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// Deactivate - деактивировать язык.
func (adapter *Adapter_HttpRestAPI) Deactivate(ctx context.Context, code string) (cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, code)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if proxyErr = adapter.controller.Deactivate(ctx, code); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}
