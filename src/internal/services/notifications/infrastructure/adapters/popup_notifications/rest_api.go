package popup_notifications_adapter

import (
	"context"
	popup_notifications_controller "sm-box/internal/services/notifications/infrastructure/controllers/popup_notifications"
	"sm-box/internal/services/notifications/objects/constructors"
	"sm-box/internal/services/notifications/objects/models"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator_HttpRestAPI = "infrastructure-[adapters]=popup_notifications-(HttpRestAPI)"
)

// Adapter_HttpRestAPI - адаптер контроллера для http rest api.
type Adapter_HttpRestAPI struct {
	components *components

	controller interface {
		CreateOne(ctx context.Context, constructor *constructors.PopupNotification) (notification *models.PopupNotificationInfo, cErr c_errors.Error)
		Create(ctx context.Context, constructors ...*constructors.PopupNotification) (notifications []*models.PopupNotificationInfo, cErr c_errors.Error)
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
		if adapter.controller, err = popup_notifications_controller.New(ctx); err != nil {
			return
		}
	}

	adapter.components.Logger.Info().
		Format("A '%s' adapter for RestAPI has been created. ", "popup_notifications").Write()

	return
}

// CreateOne - создание всплывающего уведомления.
func (adapter *Adapter_HttpRestAPI) CreateOne(ctx context.Context, constructor *constructors.PopupNotification) (notification *models.PopupNotificationInfo, cErr c_errors.RestAPI) {
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

// Create - создание всплывающих уведомлений.
func (adapter *Adapter_HttpRestAPI) Create(ctx context.Context, constructors ...*constructors.PopupNotification) (notifications []*models.PopupNotificationInfo, cErr c_errors.RestAPI) {
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
