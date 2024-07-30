package urls_management_adapter

import (
	"context"
	common_types "sm-box/internal/common/types"
	urls_management_controller "sm-box/internal/services/url_shortner/infrastructure/controllers/urls_management"
	"sm-box/internal/services/url_shortner/objects"
	"sm-box/internal/services/url_shortner/objects/models"
	"sm-box/internal/services/url_shortner/objects/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	"time"
)

const (
	loggerInitiator_RestAPI = "infrastructure-[adapters]=urls_management-(RestAPI)"
)

// Adapter_HttpRestAPI - адаптер контроллера для http rest api.
type Adapter_HttpRestAPI struct {
	components *components

	controller interface {
		GetList(ctx context.Context,
			search *objects.ShortUrlsListSearch,
			sort *objects.ShortUrlsListSort,
			pagination *objects.ShortUrlsListPagination,
			filters *objects.ShortUrlsListFilters,
		) (list []*models.ShortUrlInfo, cErr c_errors.Error)
		GetOne(ctx context.Context, id common_types.ID) (url *models.ShortUrlInfo, cErr c_errors.Error)
		GetOneByReduction(ctx context.Context, reduction string) (url *models.ShortUrlInfo, cErr c_errors.Error)

		GetUsageHistory(ctx context.Context, id common_types.ID,
			sort *objects.ShortUrlsUsageHistoryListSort,
			pagination *objects.ShortUrlsUsageHistoryListPagination,
			filters *objects.ShortUrlsUsageHistoryListFilters,
		) (history []*models.ShortUrlUsageHistoryInfo, cErr c_errors.Error)
		GetUsageHistoryByReduction(ctx context.Context, reduction string,
			sort *objects.ShortUrlsUsageHistoryListSort,
			pagination *objects.ShortUrlsUsageHistoryListPagination,
			filters *objects.ShortUrlsUsageHistoryListFilters,
		) (history []*models.ShortUrlUsageHistoryInfo, cErr c_errors.Error)

		Create(ctx context.Context,
			source string,
			type_ types.ShortUrlType,
			numberOfUses int64,
			startActive, endActive time.Time,
		) (url *models.ShortUrlInfo, cErr c_errors.Error)

		Remove(ctx context.Context, id common_types.ID) (cErr c_errors.Error)
		RemoveByReduction(ctx context.Context, reduction string) (cErr c_errors.Error)

		Activate(ctx context.Context, id common_types.ID) (cErr c_errors.Error)
		ActivateByReduction(ctx context.Context, reduction string) (cErr c_errors.Error)

		Deactivate(ctx context.Context, id common_types.ID) (cErr c_errors.Error)
		DeactivateByReduction(ctx context.Context, reduction string) (cErr c_errors.Error)
	}

	ctx context.Context
}

// components - компоненты адаптера.
type components struct {
	Logger logger.Logger
}

// New_RestAPI - создание контроллера для rest api.
func New_RestAPI(ctx context.Context) (adapter *Adapter_HttpRestAPI, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelMain, tracer.LevelAdapter)

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
			if adapter.components.Logger, err = logger.New(loggerInitiator_RestAPI); err != nil {
				return
			}
		}
	}

	// Контроллер
	{
		if adapter.controller, err = urls_management_controller.New(ctx); err != nil {
			return
		}
	}

	adapter.components.Logger.Info().
		Format("A '%s' adapter for RestAPI has been created. ", "urls_management").Write()

	return
}

// GetList - получение списка сокращенных url.
func (adapter *Adapter_HttpRestAPI) GetList(ctx context.Context,
	search *objects.ShortUrlsListSearch,
	sort *objects.ShortUrlsListSort,
	pagination *objects.ShortUrlsListPagination,
	filters *objects.ShortUrlsListFilters,
) (list []*models.ShortUrlInfo, cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, search, sort, pagination, filters)
		defer func() { trc.Error(cErr).FunctionCallFinished(list) }()
	}

	var proxyErr c_errors.Error

	if list, proxyErr = adapter.controller.GetList(ctx, search, sort, pagination, filters); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// GetOne - получение сокращенного url.
func (adapter *Adapter_HttpRestAPI) GetOne(ctx context.Context, id common_types.ID) (url *models.ShortUrlInfo, cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(cErr).FunctionCallFinished(url) }()
	}

	var proxyErr c_errors.Error

	if url, proxyErr = adapter.controller.GetOne(ctx, id); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// GetOneByReduction - получение сокращенного url по сокращению.
func (adapter *Adapter_HttpRestAPI) GetOneByReduction(ctx context.Context, reduction string) (url *models.ShortUrlInfo, cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, reduction)
		defer func() { trc.Error(cErr).FunctionCallFinished(url) }()
	}

	var proxyErr c_errors.Error

	if url, proxyErr = adapter.controller.GetOneByReduction(ctx, reduction); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// GetUsageHistory - получение истории использования сокращенного url.
func (adapter *Adapter_HttpRestAPI) GetUsageHistory(ctx context.Context, id common_types.ID,
	sort *objects.ShortUrlsUsageHistoryListSort,
	pagination *objects.ShortUrlsUsageHistoryListPagination,
	filters *objects.ShortUrlsUsageHistoryListFilters,
) (history []*models.ShortUrlUsageHistoryInfo, cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, id, sort, pagination, filters)
		defer func() { trc.Error(cErr).FunctionCallFinished(history) }()
	}

	var proxyErr c_errors.Error

	if history, proxyErr = adapter.controller.GetUsageHistory(ctx, id, sort, pagination, filters); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// GetUsageHistoryByReduction - получение истории использования сокращенного url по сокращению.
func (adapter *Adapter_HttpRestAPI) GetUsageHistoryByReduction(ctx context.Context, reduction string,
	sort *objects.ShortUrlsUsageHistoryListSort,
	pagination *objects.ShortUrlsUsageHistoryListPagination,
	filters *objects.ShortUrlsUsageHistoryListFilters,
) (history []*models.ShortUrlUsageHistoryInfo, cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, reduction, sort, pagination, filters)
		defer func() { trc.Error(cErr).FunctionCallFinished(history) }()
	}

	var proxyErr c_errors.Error

	if history, proxyErr = adapter.controller.GetUsageHistoryByReduction(ctx, reduction, sort, pagination, filters); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// Create - создание сокращенного url.
func (adapter *Adapter_HttpRestAPI) Create(ctx context.Context,
	source string,
	type_ types.ShortUrlType,
	numberOfUses int64,
	startActive, endActive time.Time) (url *models.ShortUrlInfo, cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, source, type_, numberOfUses, startActive, endActive)
		defer func() { trc.Error(cErr).FunctionCallFinished(url) }()
	}

	var proxyErr c_errors.Error

	if url, proxyErr = adapter.controller.Create(ctx, source, type_, numberOfUses, startActive, endActive); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// Remove - удаление сокращенного url.
func (adapter *Adapter_HttpRestAPI) Remove(ctx context.Context, id common_types.ID) (cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if proxyErr = adapter.controller.Remove(ctx, id); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// RemoveByReduction - удаление сокращенного url по сокращению.
func (adapter *Adapter_HttpRestAPI) RemoveByReduction(ctx context.Context, reduction string) (cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, reduction)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if proxyErr = adapter.controller.RemoveByReduction(ctx, reduction); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// Activate - активация сокращенного url.
func (adapter *Adapter_HttpRestAPI) Activate(ctx context.Context, id common_types.ID) (cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if proxyErr = adapter.controller.Activate(ctx, id); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// ActivateByReduction - активация сокращенного url по сокращению.
func (adapter *Adapter_HttpRestAPI) ActivateByReduction(ctx context.Context, reduction string) (cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, reduction)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if proxyErr = adapter.controller.ActivateByReduction(ctx, reduction); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// Deactivate - деактивация сокращенного url.
func (adapter *Adapter_HttpRestAPI) Deactivate(ctx context.Context, id common_types.ID) (cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if proxyErr = adapter.controller.Deactivate(ctx, id); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// DeactivateByReduction - деактивация сокращенного url по сокращению.
func (adapter *Adapter_HttpRestAPI) DeactivateByReduction(ctx context.Context, reduction string) (cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, reduction)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if proxyErr = adapter.controller.DeactivateByReduction(ctx, reduction); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}