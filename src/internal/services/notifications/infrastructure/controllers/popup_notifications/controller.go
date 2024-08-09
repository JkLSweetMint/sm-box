package popup_notifications_controller

import (
	"context"
	popup_notifications_usecase "sm-box/internal/services/notifications/infrastructure/usecases/popup_notifications"
	"sm-box/internal/services/notifications/objects/constructors"
	"sm-box/internal/services/notifications/objects/entities"
	"sm-box/internal/services/notifications/objects/models"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator = "infrastructure-[controllers]=popup_notifications"
)

// Controller - контроллер пользовательских уведомлений.
type Controller struct {
	components *components
	usecases   *usecases

	conf *Config
	ctx  context.Context
}

// usecases - логика контроллера.
type usecases struct {
	PopupNotifications interface {
		CreateOne(ctx context.Context, constructor *constructors.PopupNotification) (notification *entities.PopupNotification, cErr c_errors.Error)
		Create(ctx context.Context, constructors ...*constructors.PopupNotification) (notifications []*entities.PopupNotification, cErr c_errors.Error)
	}
}

// components - компоненты контроллера.
type components struct {
	Logger logger.Logger
}

// New - создание контроллера.
func New(ctx context.Context) (controller *Controller, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelMain, tracer.LevelController)

		trace.FunctionCall(ctx)

		defer func() { trace.Error(err).FunctionCallFinished(controller) }()
	}

	controller = new(Controller)
	controller.ctx = ctx

	// Конфигурация
	{
		controller.conf = new(Config).Default()

		if err = controller.conf.Read(); err != nil {
			return
		}
	}

	// Компоненты
	{
		controller.components = new(components)

		// Logger
		{
			if controller.components.Logger, err = logger.New(loggerInitiator); err != nil {
				return
			}
		}
	}

	// Логика
	{
		controller.usecases = new(usecases)

		// PopupNotifications
		{
			if controller.usecases.PopupNotifications, err = popup_notifications_usecase.New(ctx); err != nil {
				return
			}
		}
	}

	controller.components.Logger.Info().
		Format("A '%s' controller has been created. ", "popup_notifications").
		Field("config", controller.conf).Write()

	return
}

// CreateOne - создание всплывающего уведомления.
func (controller *Controller) CreateOne(ctx context.Context, constructor *constructors.PopupNotification) (notification *models.PopupNotificationInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, constructor)
		defer func() { trc.Error(cErr).FunctionCallFinished(notification) }()
	}

	// Выполнения инструкций
	{
		var notification_ *entities.PopupNotification

		if notification_, cErr = controller.usecases.PopupNotifications.CreateOne(ctx, constructor); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}

		// Преобразование в модель
		{
			if notification_ != nil {
				notification = notification_.ToModel()
			}
		}
	}

	return
}

// Create - создание всплывающих уведомлений.
func (controller *Controller) Create(ctx context.Context, constructors ...*constructors.PopupNotification) (notifications []*models.PopupNotificationInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, constructors)
		defer func() { trc.Error(cErr).FunctionCallFinished(notifications) }()
	}

	// Выполнения инструкций
	{
		var notifications_ []*entities.PopupNotification

		if notifications_, cErr = controller.usecases.PopupNotifications.Create(ctx, constructors...); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}

		// Преобразование в модель
		{
			notifications = make([]*models.PopupNotificationInfo, 0, len(notifications))

			for _, notification := range notifications_ {
				notifications = append(notifications, notification.ToModel())
			}
		}
	}

	return
}
