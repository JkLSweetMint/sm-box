package user_notifications_adapter

import (
	"context"
	common_types "sm-box/internal/common/types"
	user_notifications_controller "sm-box/internal/services/notifications/infrastructure/controllers/user_notifications"
	"sm-box/internal/services/notifications/objects"
	"sm-box/internal/services/notifications/objects/constructors"
	"sm-box/internal/services/notifications/objects/models"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator_HttpRestAPI = "infrastructure-[adapters]=user_notifications-(HttpRestAPI)"
)

// Adapter_HttpRestAPI - адаптер контроллера для http rest api.
type Adapter_HttpRestAPI struct {
	components *components

	controller interface {
		GetList(ctx context.Context,
			userID common_types.ID,
			search *objects.UserNotificationSearch,
			pagination *objects.UserNotificationPagination,
			filters *objects.UserNotificationFilters,
		) (count int64, list []*models.UserNotificationInfo, cErr c_errors.Error)

		CreateOne(ctx context.Context, constructor *constructors.UserNotification) (notification *models.UserNotificationInfo, cErr c_errors.Error)
		Create(ctx context.Context, constructors ...*constructors.UserNotification) (notifications []*models.UserNotificationInfo, cErr c_errors.Error)

		RemoveOne(ctx context.Context, userID common_types.ID, id common_types.ID) (cErr c_errors.Error)
		Remove(ctx context.Context, userID common_types.ID, ids ...common_types.ID) (cErr c_errors.Error)

		ReadOne(ctx context.Context, userID common_types.ID, id common_types.ID) (cErr c_errors.Error)
		Read(ctx context.Context, userID common_types.ID, ids ...common_types.ID) (cErr c_errors.Error)
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
			if adapter.components.Logger, err = logger.New(loggerInitiator_HttpRestAPI); err != nil {
				return
			}
		}
	}

	// Контроллер
	{
		if adapter.controller, err = user_notifications_controller.New(ctx); err != nil {
			return
		}
	}

	adapter.components.Logger.Info().
		Format("A '%s' adapter for RestAPI has been created. ", "user_notifications").Write()

	return
}

// GetList - получение списка пользовательских уведомлений.
func (adapter *Adapter_HttpRestAPI) GetList(ctx context.Context,
	userID common_types.ID,
	search *objects.UserNotificationSearch,
	pagination *objects.UserNotificationPagination,
	filters *objects.UserNotificationFilters,
) (count int64, list []*models.UserNotificationInfo, cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, userID, search, pagination, filters)
		defer func() { trc.Error(cErr).FunctionCallFinished(count, list) }()
	}

	var proxyErr c_errors.Error

	if count, list, proxyErr = adapter.controller.GetList(ctx, userID, search, pagination, filters); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// CreateOne - создание пользовательского уведомления.
func (adapter *Adapter_HttpRestAPI) CreateOne(ctx context.Context, constructor *constructors.UserNotification) (notification *models.UserNotificationInfo, cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, constructor)
		defer func() { trc.Error(cErr).FunctionCallFinished(notification) }()
	}

	var proxyErr c_errors.Error

	if notification, proxyErr = adapter.controller.CreateOne(ctx, constructor); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// Create - создание пользовательских уведомлений.
func (adapter *Adapter_HttpRestAPI) Create(ctx context.Context, constructors ...*constructors.UserNotification) (notifications []*models.UserNotificationInfo, cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, constructors)
		defer func() { trc.Error(cErr).FunctionCallFinished(notifications) }()
	}

	var proxyErr c_errors.Error

	if notifications, proxyErr = adapter.controller.Create(ctx, constructors...); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// RemoveOne - удаление пользовательского уведомления.
func (adapter *Adapter_HttpRestAPI) RemoveOne(ctx context.Context, userID, id common_types.ID) (cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, userID, id)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if proxyErr = adapter.controller.RemoveOne(ctx, userID, id); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// Remove - удаление пользовательских уведомлений.
func (adapter *Adapter_HttpRestAPI) Remove(ctx context.Context, userID common_types.ID, ids ...common_types.ID) (cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, userID, ids)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if proxyErr = adapter.controller.Remove(ctx, userID, ids...); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// ReadOne - чтение пользовательского уведомления.
func (adapter *Adapter_HttpRestAPI) ReadOne(ctx context.Context, userID, id common_types.ID) (cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, userID, id)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if proxyErr = adapter.controller.ReadOne(ctx, userID, id); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}

// Read - чтение пользовательских уведомлений.
func (adapter *Adapter_HttpRestAPI) Read(ctx context.Context, userID common_types.ID, ids ...common_types.ID) (cErr c_errors.RestAPI) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelAdapter)

		trc.FunctionCall(ctx, userID, ids)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var proxyErr c_errors.Error

	if proxyErr = adapter.controller.Read(ctx, userID, ids...); proxyErr != nil {
		cErr = c_errors.ToRestAPI(proxyErr)

		adapter.components.Logger.Error().
			Format("The controller method was executed with an error: '%s'. ", cErr).Write()
		return
	}

	return
}
