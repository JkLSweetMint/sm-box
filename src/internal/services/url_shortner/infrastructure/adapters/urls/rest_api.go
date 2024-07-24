package urls_adapter

import (
	"context"
	common_types "sm-box/internal/common/types"
	authentication_entities "sm-box/internal/services/authentication/objects/entities"
	urls_controller "sm-box/internal/services/url_shortner/infrastructure/controllers/urls"
	"sm-box/internal/services/url_shortner/objects/models"
	"sm-box/internal/services/url_shortner/objects/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator_RestAPI = "infrastructure-[adapters]=urls-(RestAPI)"
)

// Adapter_RestAPI - адаптер контроллера для rest api.
type Adapter_RestAPI struct {
	components *components

	controller interface {
		GetByReduceFromRedisDB(ctx context.Context, reduce string) (url *models.ShortUrlInfo, cErr c_errors.Error)
		UpdateInRedisDB(ctx context.Context, url *models.ShortUrlInfo) (cErr c_errors.Error)
		RemoveByReduceFromRedisDB(ctx context.Context, reduce string) (cErr c_errors.Error)

		WriteCallToHistory(ctx context.Context, id common_types.ID, status types.ShortUrlUsageHistoryStatus, token *authentication_entities.JwtSessionToken) (cErr c_errors.Error)
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
		if adapter.controller, err = urls_controller.New(ctx); err != nil {
			return
		}
	}

	adapter.components.Logger.Info().
		Format("A '%s' adapter for RestAPI has been created. ", "urls").Write()

	return
}

// GetByReduceFromRedisDB - получение короткого маршрута по сокращению из базы данных redis.
func (adapter *Adapter_RestAPI) GetByReduceFromRedisDB(ctx context.Context, reduce string) (url *models.ShortUrlInfo, cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, reduce)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if url, proxyErr = adapter.controller.GetByReduceFromRedisDB(ctx, reduce); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// UpdateInRedisDB - обновление короткого маршрута в базу данных redis.
func (adapter *Adapter_RestAPI) UpdateInRedisDB(ctx context.Context, url *models.ShortUrlInfo) (cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, url)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if proxyErr = adapter.controller.UpdateInRedisDB(ctx, url); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// RemoveByReduceFromRedisDB - получение короткого маршрута по сокращению из базы данных redis.
func (adapter *Adapter_RestAPI) RemoveByReduceFromRedisDB(ctx context.Context, reduce string) (cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, reduce)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if proxyErr = adapter.controller.RemoveByReduceFromRedisDB(ctx, reduce); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// WriteCallToHistory - записать обращение по короткой ссылке в историю.
func (adapter *Adapter_RestAPI) WriteCallToHistory(ctx context.Context, id common_types.ID, status types.ShortUrlUsageHistoryStatus, token *authentication_entities.JwtSessionToken) (cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, id, status, token)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if proxyErr = adapter.controller.WriteCallToHistory(ctx, id, status, token); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}
