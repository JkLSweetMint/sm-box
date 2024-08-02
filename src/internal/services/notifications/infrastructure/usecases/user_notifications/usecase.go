package user_notifications_usecase

import (
	"context"
	common_errors "sm-box/internal/common/errors"
	common_types "sm-box/internal/common/types"
	user_notifications_repository "sm-box/internal/services/notifications/infrastructure/repositories/user_notifications"
	"sm-box/internal/services/notifications/objects"
	"sm-box/internal/services/notifications/objects/constructors"
	"sm-box/internal/services/notifications/objects/entities"
	"sm-box/internal/services/notifications/objects/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	"strings"
)

const (
	loggerInitiator = "infrastructure-[usecases]=user_notifications"
)

// UseCase - логика пользовательских уведомлений.
type UseCase struct {
	components   *components
	repositories *repositories

	conf *Config
	ctx  context.Context
}

// repositories - репозитории логики.
type repositories struct {
	UserNotifications interface {
		GetList(ctx context.Context,
			userID common_types.ID,
			search *objects.UserNotificationSearch,
			pagination *objects.UserNotificationPagination,
			filters *objects.UserNotificationFilters,
		) (count int64, list []*entities.UserNotification, err error)

		CreateOne(ctx context.Context, constructor *constructors.UserNotification) (notification *entities.UserNotification, err error)
		Create(ctx context.Context, constructors ...*constructors.UserNotification) (notifications []*entities.UserNotification, err error)

		RemoveOne(ctx context.Context, userID, id common_types.ID) (err error)
		Remove(ctx context.Context, userID common_types.ID, ids ...common_types.ID) (err error)

		ReadOne(ctx context.Context, userID, id common_types.ID) (err error)
		Read(ctx context.Context, userID common_types.ID, ids ...common_types.ID) (err error)
	}
}

// components - компоненты логики.
type components struct {
	Logger logger.Logger
}

// New - создание логики.
func New(ctx context.Context) (usecase *UseCase, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelMain, tracer.LevelUseCase)

		trace.FunctionCall(ctx)

		defer func() { trace.Error(err).FunctionCallFinished(usecase) }()
	}

	usecase = new(UseCase)
	usecase.ctx = ctx

	// Конфигурация
	{
		usecase.conf = new(Config).Default()

		if err = usecase.conf.Read(); err != nil {
			return
		}
	}

	// Компоненты
	{
		usecase.components = new(components)

		// Logger
		{
			if usecase.components.Logger, err = logger.New(loggerInitiator); err != nil {
				return
			}
		}
	}

	// Репозитории
	{
		usecase.repositories = new(repositories)

		// UserNotifications
		{
			if usecase.repositories.UserNotifications, err = user_notifications_repository.New(usecase.ctx); err != nil {
				return
			}
		}
	}

	usecase.components.Logger.Info().
		Format("A '%s' usecase has been created. ", "user_notifications").
		Field("config", usecase.conf).Write()

	return
}

// GetList - получение списка пользовательских уведомлений.
func (usecase *UseCase) GetList(ctx context.Context,
	userID common_types.ID,
	search *objects.UserNotificationSearch,
	pagination *objects.UserNotificationPagination,
	filters *objects.UserNotificationFilters,
) (count int64, list []*entities.UserNotification, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, userID, search, pagination, filters)
		defer func() { trc.Error(cErr).FunctionCallFinished(count, list) }()
	}

	usecase.components.Logger.Info().
		Text("The process of getting a list of user notifications has started... ").
		Field("user_id", userID).
		Field("search", search).
		Field("pagination", pagination).
		Field("filters", filters).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of getting a list of user notifications is completed. ").
			Field("user_id", userID).
			Field("search", search).
			Field("pagination", pagination).
			Field("filters", filters).Write()
	}()

	// Подготовка данных
	{
		if search != nil {
			search.Global = strings.TrimSpace(search.Global)
		}
	}

	// Валидация
	{
		if userID < 0 {
			usecase.components.Logger.Error().
				Text("An invalid argument value was passed. ").
				Field("user_id", userID).Write()

			cErr = common_errors.InvalidArguments()
			cErr.Details().Set("user_id", "Negative value. ")

			return
		} else if userID == 0 {
			usecase.components.Logger.Error().
				Text("An invalid argument value was passed. ").
				Field("user_id", userID).Write()

			cErr = common_errors.InvalidArguments()
			cErr.Details().Set("user_id", "Zero value. ")

			return
		}

		if filters != nil {
			if filters.Type != nil {
				var v = *filters.Type

				if v != types.NotificationTypeAlerts {
					if cErr == nil {
						cErr = common_errors.InvalidFilterValue()
					}

					usecase.components.Logger.Error().
						Text("An invalid filter value was passed. ").
						Field("value", v).Write()

					cErr.Details().Set("filter_type", "Invalid value. ")
				}
			}

			if filters.SenderID != nil {
				var v = *filters.SenderID

				if v < 0 {
					if cErr == nil {
						cErr = common_errors.InvalidFilterValue()
					}

					usecase.components.Logger.Error().
						Text("An invalid filter value was passed. ").
						Field("value", v).Write()

					cErr.Details().Set("filter_sender_id", "Negative value. ")
				}
			}

			if cErr != nil {
				return
			}
		}
	}

	// Получение
	{
		var err error

		if count, list, err = usecase.repositories.UserNotifications.GetList(ctx, userID, search, pagination, filters); err != nil {
			usecase.components.Logger.Error().
				Format("Failed to get a list of user notifications: '%s'. ", err).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The list of user notifications was successfully received. ").
			Field("list", list).
			Field("count", count).Write()
	}

	return
}

// CreateOne - создание пользовательского уведомления.
func (usecase *UseCase) CreateOne(ctx context.Context, constructor *constructors.UserNotification) (notification *entities.UserNotification, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, constructor)
		defer func() { trc.Error(cErr).FunctionCallFinished(notification) }()
	}

	return
}

// Create - создание пользовательских уведомлений.
func (usecase *UseCase) Create(ctx context.Context, constructors ...*constructors.UserNotification) (notifications []*entities.UserNotification, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, constructors)
		defer func() { trc.Error(cErr).FunctionCallFinished(notifications) }()
	}

	return
}

// RemoveOne - удаление пользовательского уведомления.
func (usecase *UseCase) RemoveOne(ctx context.Context, userID, id common_types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, userID, id)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	return
}

// Remove - удаление пользовательских уведомлений.
func (usecase *UseCase) Remove(ctx context.Context, userID common_types.ID, ids ...common_types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, userID, ids)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	return
}

// ReadOne - чтение пользовательского уведомления.
func (usecase *UseCase) ReadOne(ctx context.Context, userID, id common_types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, userID, id)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	return
}

// Read - чтение пользовательских уведомлений.
func (usecase *UseCase) Read(ctx context.Context, userID common_types.ID, ids ...common_types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, userID, ids)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	return
}
