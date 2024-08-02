package user_notifications_controller

import (
	"context"
	common_types "sm-box/internal/common/types"
	user_notifications_usecase "sm-box/internal/services/notifications/infrastructure/usecases/user_notifications"
	"sm-box/internal/services/notifications/objects"
	"sm-box/internal/services/notifications/objects/constructors"
	"sm-box/internal/services/notifications/objects/entities"
	"sm-box/internal/services/notifications/objects/models"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator = "infrastructure-[controllers]=user_notifications"
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
	UserNotifications interface {
		GetList(ctx context.Context,
			userID common_types.ID,
			search *objects.UserNotificationSearch,
			pagination *objects.UserNotificationPagination,
			filters *objects.UserNotificationFilters,
		) (count int64, list []*entities.UserNotification, cErr c_errors.Error)

		CreateOne(ctx context.Context, constructor *constructors.UserNotification) (notification *entities.UserNotification, cErr c_errors.Error)
		Create(ctx context.Context, constructors ...*constructors.UserNotification) (notifications []*entities.UserNotification, cErr c_errors.Error)

		RemoveOne(ctx context.Context, userID common_types.ID, id common_types.ID) (cErr c_errors.Error)
		Remove(ctx context.Context, userID common_types.ID, ids ...common_types.ID) (cErr c_errors.Error)

		ReadOne(ctx context.Context, userID common_types.ID, id common_types.ID) (cErr c_errors.Error)
		Read(ctx context.Context, userID common_types.ID, ids ...common_types.ID) (cErr c_errors.Error)
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

		// UserNotifications
		{
			if controller.usecases.UserNotifications, err = user_notifications_usecase.New(ctx); err != nil {
				return
			}
		}
	}

	controller.components.Logger.Info().
		Format("A '%s' controller has been created. ", "user_notifications").
		Field("config", controller.conf).Write()

	return
}

// GetList - получение списка пользовательских уведомлений.
func (controller *Controller) GetList(ctx context.Context,
	userID common_types.ID,
	search *objects.UserNotificationSearch,
	pagination *objects.UserNotificationPagination,
	filters *objects.UserNotificationFilters,
) (count int64, list []*models.UserNotificationInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, userID, search, pagination, filters)
		defer func() { trc.Error(cErr).FunctionCallFinished(count, list) }()
	}

	// Выполнения инструкций
	{
		var notifications []*entities.UserNotification

		if count, notifications, cErr = controller.usecases.UserNotifications.GetList(ctx, userID, search, pagination, filters); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}

		// Преобразование в модель
		{
			list = make([]*models.UserNotificationInfo, 0, len(notifications))

			for _, notification := range notifications {
				list = append(list, notification.ToModel())
			}
		}
	}

	return
}

// CreateOne - создание пользовательского уведомления.
func (controller *Controller) CreateOne(ctx context.Context, constructor *constructors.UserNotification) (notification *models.UserNotificationInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, constructor)
		defer func() { trc.Error(cErr).FunctionCallFinished(notification) }()
	}

	// Выполнения инструкций
	{
		var notification_ *entities.UserNotification

		if notification_, cErr = controller.usecases.UserNotifications.CreateOne(ctx, constructor); cErr != nil {
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

// Create - создание пользовательских уведомлений.
func (controller *Controller) Create(ctx context.Context, constructors ...*constructors.UserNotification) (notifications []*models.UserNotificationInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, constructors)
		defer func() { trc.Error(cErr).FunctionCallFinished(notifications) }()
	}

	// Выполнения инструкций
	{
		var notifications_ []*entities.UserNotification

		if notifications_, cErr = controller.usecases.UserNotifications.Create(ctx, constructors...); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}

		// Преобразование в модель
		{
			notifications = make([]*models.UserNotificationInfo, 0, len(notifications))

			for _, notification := range notifications_ {
				notifications = append(notifications, notification.ToModel())
			}
		}
	}

	return
}

// RemoveOne - удаление пользовательского уведомления.
func (controller *Controller) RemoveOne(ctx context.Context, userID, id common_types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, userID, id)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	// Выполнения инструкций
	{
		if cErr = controller.usecases.UserNotifications.RemoveOne(ctx, userID, id); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	return
}

// Remove - удаление пользовательских уведомлений.
func (controller *Controller) Remove(ctx context.Context, userID common_types.ID, ids ...common_types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, userID, ids)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	// Выполнения инструкций
	{
		if cErr = controller.usecases.UserNotifications.Remove(ctx, userID, ids...); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	return
}

// ReadOne - чтение пользовательского уведомления.
func (controller *Controller) ReadOne(ctx context.Context, userID, id common_types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, userID, id)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	// Выполнения инструкций
	{
		if cErr = controller.usecases.UserNotifications.ReadOne(ctx, userID, id); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	return
}

// Read - чтение пользовательских уведомлений.
func (controller *Controller) Read(ctx context.Context, userID common_types.ID, ids ...common_types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, userID, ids)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	// Выполнения инструкций
	{
		if cErr = controller.usecases.UserNotifications.Read(ctx, userID, ids...); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	return
}
