package user_notifications_usecase

import (
	"context"
	"fmt"
	common_errors "sm-box/internal/common/errors"
	common_types "sm-box/internal/common/types"
	user_notifications_repository "sm-box/internal/services/notifications/infrastructure/repositories/user_notifications"
	"sm-box/internal/services/notifications/objects"
	"sm-box/internal/services/notifications/objects/constructors"
	"sm-box/internal/services/notifications/objects/entities"
	srv_errors "sm-box/internal/services/notifications/objects/errors"
	"sm-box/internal/services/notifications/objects/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	err_details "sm-box/pkg/errors/entities/details"
	err_messages "sm-box/pkg/errors/entities/messages"
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
			recipientID common_types.ID,
			search *objects.UserNotificationSearch,
			pagination *objects.UserNotificationPagination,
			filters *objects.UserNotificationFilters,
		) (count, countNotRead int64, list []*entities.UserNotification, err error)

		CreateOne(ctx context.Context, constructor *constructors.UserNotification) (notification *entities.UserNotification, err error)
		Create(ctx context.Context, constructors ...*constructors.UserNotification) (notifications []*entities.UserNotification, err error)

		RemoveOne(ctx context.Context, recipientID, id common_types.ID) (err error)
		Remove(ctx context.Context, recipientID common_types.ID, ids ...common_types.ID) (err error)

		ReadOne(ctx context.Context, recipientID, id common_types.ID) (err error)
		Read(ctx context.Context, recipientID common_types.ID, ids ...common_types.ID) (err error)

		Exists(ctx context.Context, ids ...common_types.ID) (exists []bool, err error)
		AlreadyRead(ctx context.Context, ids ...common_types.ID) (read []bool, err error)
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
	recipientID common_types.ID,
	search *objects.UserNotificationSearch,
	pagination *objects.UserNotificationPagination,
	filters *objects.UserNotificationFilters,
) (count, countNotRead int64, list []*entities.UserNotification, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, recipientID, search, pagination, filters)
		defer func() { trc.Error(cErr).FunctionCallFinished(count, countNotRead, list) }()
	}

	usecase.components.Logger.Info().
		Text("The process of getting a list of user notifications has started... ").
		Field("recipient_id", recipientID).
		Field("search", search).
		Field("pagination", pagination).
		Field("filters", filters).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of getting a list of user notifications is completed. ").
			Field("recipient_id", recipientID).
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
		// User ID
		{
			if recipientID < 0 {
				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("recipient_id", recipientID).Write()

				cErr = common_errors.InvalidArguments()

				cErr.Details().SetField(
					new(err_details.FieldKey).Add("recipient_id"),
					new(err_messages.TextMessage).Text("Negative value. "),
				)

				return
			} else if recipientID == 0 {
				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("recipient_id", recipientID).Write()

				cErr = common_errors.InvalidArguments()

				cErr.Details().SetField(
					new(err_details.FieldKey).Add("recipient_id"),
					new(err_messages.TextMessage).Text("Zero value. "),
				)

				return
			}
		}

		// Фильтрация
		{
			if filters != nil {
				var tmpCErr c_errors.Error

				if filters.Type != nil {
					var v = *filters.Type

					if v != types.NotificationTypeAlerts {
						if tmpCErr == nil {
							tmpCErr = common_errors.InvalidFilterValue()
						}

						usecase.components.Logger.Error().
							Text("An invalid filter value was passed. ").
							Field("value", v).Write()

						cErr.Details().SetField(
							new(err_details.FieldKey).Add("filter_type"),
							new(err_messages.TextMessage).Text("Invalid value. "),
						)
					}
				}

				if filters.SenderID != nil {
					var v = *filters.SenderID

					if v < 0 {
						if tmpCErr == nil {
							tmpCErr = common_errors.InvalidFilterValue()
						}

						usecase.components.Logger.Error().
							Text("An invalid filter value was passed. ").
							Field("value", v).Write()

						cErr.Details().SetField(
							new(err_details.FieldKey).Add("filter_sender_id"),
							new(err_messages.TextMessage).Text("Negative value. "),
						)
					}
				}

				if tmpCErr != nil {
					cErr = tmpCErr
					return
				}
			}
		}
	}

	// Получение
	{
		var err error

		if count, countNotRead, list, err = usecase.repositories.UserNotifications.GetList(ctx, recipientID, search, pagination, filters); err != nil {
			usecase.components.Logger.Error().
				Format("Failed to get a list of user notifications: '%s'. ", err).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The list of user notifications was successfully received. ").
			Field("list", list).
			Field("count", count).
			Field("count_not_read", countNotRead).Write()
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

	fmt.Printf("\n\n%+v\n\n\n", constructor)

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

	fmt.Printf("\n\n%+v\n\n\n", constructors)

	return
}

// RemoveOne - удаление пользовательского уведомления.
func (usecase *UseCase) RemoveOne(ctx context.Context, recipientID, id common_types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, recipientID, id)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The process of deleting one user notification has started... ").
		Field("recipient_id", recipientID).
		Field("id", id).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of deleting one user notification is completed. ").
			Field("recipient_id", recipientID).
			Field("id", id).Write()
	}()

	// Валидация
	{
		var tempCErr c_errors.Error

		// User ID
		{
			if recipientID < 0 {
				if tempCErr == nil {
					tempCErr = common_errors.InvalidArguments()
				}

				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("recipient_id", recipientID).Write()

				tempCErr.Details().SetField(
					new(err_details.FieldKey).Add("recipient_id"),
					new(err_messages.TextMessage).Text("Negative value. "),
				)

				return
			} else if recipientID == 0 {
				if tempCErr == nil {
					tempCErr = common_errors.InvalidArguments()
				}

				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("recipient_id", recipientID).Write()

				tempCErr.Details().SetField(
					new(err_details.FieldKey).Add("recipient_id"),
					new(err_messages.TextMessage).Text("Zero value. "),
				)

				return
			}
		}

		// ID
		{
			if id < 0 {
				if tempCErr == nil {
					tempCErr = common_errors.InvalidArguments()
				}

				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("id", id).Write()

				tempCErr.Details().SetField(
					new(err_details.FieldKey).Add("id"),
					new(err_messages.TextMessage).Text("Negative value. "),
				)

				return
			} else if id == 0 {
				if tempCErr == nil {
					tempCErr = common_errors.InvalidArguments()
				}

				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("recipient_id", id).Write()

				tempCErr.Details().SetField(
					new(err_details.FieldKey).Add("id"),
					new(err_messages.TextMessage).Text("Zero value. "),
				)

				return
			}
		}

		if tempCErr != nil {
			cErr = tempCErr
			return
		}
	}

	// Проверки
	{
		// Существования
		{
			var (
				err    error
				exists []bool
			)

			if exists, err = usecase.repositories.UserNotifications.Exists(ctx, id); err != nil {
				usecase.components.Logger.Error().
					Format("Failed to verify the existence of user notifications: '%s'. ", err).
					Field("id", id).Write()

				cErr = common_errors.InternalServerError()
				return
			}

			if len(exists) != 1 {
				usecase.components.Logger.Error().
					Text("Invalid value of the data received from the method of checking the existence of user notifications. ").
					Field("id", id).Write()

				cErr = common_errors.InternalServerError()
				return
			}

			if !exists[0] {
				usecase.components.Logger.Error().
					Format("The user notification was not found: '%s'. ", err).
					Field("id", id).Write()

				cErr = srv_errors.UserNotificationNotFound()

				cErr.Details().SetField(
					new(err_details.FieldKey).Add("id"),
					new(err_messages.TextMessage).Text("Not found. "),
				)

				return
			}
		}
	}

	// Удаление
	{
		var err error

		if err = usecase.repositories.UserNotifications.RemoveOne(ctx, recipientID, id); err != nil {
			usecase.components.Logger.Error().
				Format("The user notification could not be deleted: '%s'. ", err).
				Field("recipient_id", recipientID).
				Field("id", id).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("Deleting the user notification has been completed successfully. ").
			Field("recipient_id", recipientID).
			Field("id", id).Write()
	}

	return
}

// Remove - удаление пользовательских уведомлений.
func (usecase *UseCase) Remove(ctx context.Context, recipientID common_types.ID, ids ...common_types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, recipientID, ids)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The process of deleting several user notifications has started... ").
		Field("recipient_id", recipientID).
		Field("ids", ids).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of deleting several user notifications is completed. ").
			Field("recipient_id", recipientID).
			Field("ids", ids).Write()
	}()

	// Валидация
	{
		var tempCErr c_errors.Error

		// User ID
		{
			if recipientID < 0 {
				if tempCErr == nil {
					tempCErr = common_errors.InvalidArguments()
				}

				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("recipient_id", recipientID).Write()

				tempCErr.Details().SetField(
					new(err_details.FieldKey).Add("recipient_id"),
					new(err_messages.TextMessage).Text("Negative value. "),
				)

				return
			} else if recipientID == 0 {
				if tempCErr == nil {
					tempCErr = common_errors.InvalidArguments()
				}

				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("recipient_id", recipientID).Write()

				tempCErr.Details().SetField(
					new(err_details.FieldKey).Add("recipient_id"),
					new(err_messages.TextMessage).Text("Zero value. "),
				)

				return
			}
		}

		// IDs
		{
			if len(ids) == 0 {
				if tempCErr == nil {
					tempCErr = common_errors.InvalidArguments()
				}

				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("ids", ids).Write()

				tempCErr.Details().SetField(
					new(err_details.FieldKey).Add("ids"),
					new(err_messages.TextMessage).Text("Is empty. "),
				)
			} else {
				for index, id := range ids {
					if id < 0 {
						if tempCErr == nil {
							tempCErr = common_errors.InvalidArguments()
						}

						usecase.components.Logger.Error().
							Text("An invalid argument value was passed. ").
							Field(fmt.Sprintf("ids.%d", index), id).Write()

						tempCErr.Details().SetField(
							new(err_details.FieldKey).AddArray("ids", index),
							new(err_messages.TextMessage).Text("Negative value. "),
						)

						return
					} else if id == 0 {
						if tempCErr == nil {
							tempCErr = common_errors.InvalidArguments()
						}

						usecase.components.Logger.Error().
							Text("An invalid argument value was passed. ").
							Field(fmt.Sprintf("ids.%d", index), id).Write()

						tempCErr.Details().SetField(
							new(err_details.FieldKey).AddArray("ids", index),
							new(err_messages.TextMessage).Text("Zero value. "),
						)

						return
					}
				}
			}
		}

		if tempCErr != nil {
			cErr = tempCErr
			return
		}
	}

	// Проверки
	{
		// Существования
		{
			var (
				tempCErr c_errors.Error
				err      error
				exists   []bool
			)

			if exists, err = usecase.repositories.UserNotifications.Exists(ctx, ids...); err != nil {
				usecase.components.Logger.Error().
					Format("Failed to verify the existence of user notifications: '%s'. ", err).
					Field("ids", ids).Write()

				cErr = common_errors.InternalServerError()
				return
			}

			if len(exists) != len(ids) {
				usecase.components.Logger.Error().
					Text("Invalid value of the data received from the method of checking the existence of user notifications. ").
					Field("ids", ids).Write()

				cErr = common_errors.InternalServerError()
				return
			}

			for index, id := range ids {
				if !exists[index] {
					if tempCErr == nil {
						tempCErr = srv_errors.UserNotificationNotFound()
					}

					usecase.components.Logger.Error().
						Text("The user notification was not found. ").
						Field("id", id).Write()

					tempCErr.Details().SetField(
						new(err_details.FieldKey).AddArray("ids", index),
						new(err_messages.TextMessage).Text("Not found. "),
					)
				}
			}

			if tempCErr != nil {
				cErr = tempCErr
				return
			}
		}
	}

	// Удаление
	{
		var err error

		if err = usecase.repositories.UserNotifications.Remove(ctx, recipientID, ids...); err != nil {
			usecase.components.Logger.Error().
				Format("Several user notifications could not be deleted: '%s'. ", err).
				Field("recipient_id", recipientID).
				Field("ids", ids).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("Deleting multiple user notifications has been completed successfully. ").
			Field("recipient_id", recipientID).
			Field("ids", ids).Write()
	}

	return
}

// ReadOne - чтение пользовательского уведомления.
func (usecase *UseCase) ReadOne(ctx context.Context, recipientID, id common_types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, recipientID, id)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The process of reading one user notification has started... ").
		Field("recipient_id", recipientID).
		Field("id", id).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of reading one user notification is completed. ").
			Field("recipient_id", recipientID).
			Field("id", id).Write()
	}()

	// Валидация
	{
		var tempCErr c_errors.Error

		// User ID
		{
			if recipientID < 0 {
				if tempCErr == nil {
					tempCErr = common_errors.InvalidArguments()
				}

				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("recipient_id", recipientID).Write()

				tempCErr.Details().SetField(
					new(err_details.FieldKey).Add("recipient_id"),
					new(err_messages.TextMessage).Text("Negative value. "),
				)

				return
			} else if recipientID == 0 {
				if tempCErr == nil {
					tempCErr = common_errors.InvalidArguments()
				}

				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("recipient_id", recipientID).Write()

				tempCErr.Details().SetField(
					new(err_details.FieldKey).Add("recipient_id"),
					new(err_messages.TextMessage).Text("Zero value. "),
				)

				return
			}
		}

		// ID
		{
			if id < 0 {
				if tempCErr == nil {
					tempCErr = common_errors.InvalidArguments()
				}

				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("id", id).Write()

				tempCErr.Details().SetField(
					new(err_details.FieldKey).Add("id"),
					new(err_messages.TextMessage).Text("Negative value. "),
				)

				return
			} else if id == 0 {
				if tempCErr == nil {
					tempCErr = common_errors.InvalidArguments()
				}

				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("recipient_id", id).Write()

				tempCErr.Details().SetField(
					new(err_details.FieldKey).Add("id"),
					new(err_messages.TextMessage).Text("Zero value. "),
				)

				return
			}
		}

		if tempCErr != nil {
			cErr = tempCErr
			return
		}
	}

	// Проверки
	{
		// Существования
		{
			var (
				err    error
				exists []bool
			)

			if exists, err = usecase.repositories.UserNotifications.Exists(ctx, id); err != nil {
				usecase.components.Logger.Error().
					Format("Failed to verify the existence of user notifications: '%s'. ", err).
					Field("id", id).Write()

				cErr = common_errors.InternalServerError()
				return
			}

			if len(exists) != 1 {
				usecase.components.Logger.Error().
					Text("Invalid value of the data received from the method of checking the existence of user notifications. ").
					Field("id", id).Write()

				cErr = common_errors.InternalServerError()
				return
			}

			if !exists[0] {
				usecase.components.Logger.Error().
					Text("The user notification was not found. ").
					Field("id", id).Write()

				cErr = srv_errors.UserNotificationNotFound()

				cErr.Details().SetField(
					new(err_details.FieldKey).Add("id"),
					new(err_messages.TextMessage).Text("Not found. "),
				)

				return
			}
		}

		// Уже прочитано
		{
			var (
				err  error
				read []bool
			)

			if read, err = usecase.repositories.UserNotifications.AlreadyRead(ctx, id); err != nil {
				usecase.components.Logger.Error().
					Format("It was not possible to check whether the user's notifications were read: '%s'. ", err).
					Field("id", id).Write()

				cErr = common_errors.InternalServerError()
				return
			}

			if len(read) != 1 {
				usecase.components.Logger.Error().
					Text("Invalid value of the data received from the method of checking read user notifications. ").
					Field("id", id).Write()

				cErr = common_errors.InternalServerError()
				return
			}

			if read[0] {
				usecase.components.Logger.Error().
					Text("The user notification already read. ").
					Field("id", id).Write()

				cErr = srv_errors.UserNotificationAlreadyRead()

				cErr.Details().SetField(
					new(err_details.FieldKey).Add("id"),
					new(err_messages.TextMessage).Text("Already read. "),
				)

				return
			}
		}
	}

	// Чтение
	{
		var err error

		if err = usecase.repositories.UserNotifications.ReadOne(ctx, recipientID, id); err != nil {
			usecase.components.Logger.Error().
				Format("The user notification could not be read: '%s'. ", err).
				Field("recipient_id", recipientID).
				Field("id", id).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("Reading the user notification has been completed successfully. ").
			Field("recipient_id", recipientID).
			Field("id", id).Write()
	}

	return
}

// Read - чтение пользовательских уведомлений.
func (usecase *UseCase) Read(ctx context.Context, recipientID common_types.ID, ids ...common_types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, recipientID, ids)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The process of reading several user notifications has started... ").
		Field("recipient_id", recipientID).
		Field("ids", ids).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of reading several user notifications is completed. ").
			Field("recipient_id", recipientID).
			Field("ids", ids).Write()
	}()

	// Валидация
	{
		var tempCErr c_errors.Error

		// User ID
		{
			if recipientID < 0 {
				if tempCErr == nil {
					tempCErr = common_errors.InvalidArguments()
				}

				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("recipient_id", recipientID).Write()

				tempCErr.Details().SetField(
					new(err_details.FieldKey).Add("recipient_id"),
					new(err_messages.TextMessage).Text("Negative value. "),
				)

				return
			} else if recipientID == 0 {
				if tempCErr == nil {
					tempCErr = common_errors.InvalidArguments()
				}

				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("recipient_id", recipientID).Write()

				tempCErr.Details().SetField(
					new(err_details.FieldKey).Add("recipient_id"),
					new(err_messages.TextMessage).Text("Zero value. "),
				)

				return
			}
		}

		// IDs
		{
			if len(ids) == 0 {
				if tempCErr == nil {
					tempCErr = common_errors.InvalidArguments()
				}

				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("ids", ids).Write()

				tempCErr.Details().SetField(
					new(err_details.FieldKey).Add("ids"),
					new(err_messages.TextMessage).Text("Is empty. "),
				)
			} else {
				for index, id := range ids {
					if id < 0 {
						if tempCErr == nil {
							tempCErr = common_errors.InvalidArguments()
						}

						usecase.components.Logger.Error().
							Text("An invalid argument value was passed. ").
							Field(fmt.Sprintf("ids.%d", index), id).Write()

						tempCErr.Details().SetField(
							new(err_details.FieldKey).AddArray("ids", index),
							new(err_messages.TextMessage).Text("Negative value. "),
						)

						return
					} else if id == 0 {
						if tempCErr == nil {
							tempCErr = common_errors.InvalidArguments()
						}

						usecase.components.Logger.Error().
							Text("An invalid argument value was passed. ").
							Field(fmt.Sprintf("ids.%d", index), id).Write()

						tempCErr.Details().SetField(
							new(err_details.FieldKey).AddArray("ids", index),
							new(err_messages.TextMessage).Text("Zero value. "),
						)

						return
					}
				}
			}
		}

		if tempCErr != nil {
			cErr = tempCErr
			return
		}
	}

	// Проверки
	{
		// Существования
		{
			var (
				tempCErr c_errors.Error
				err      error
				exists   []bool
			)

			if exists, err = usecase.repositories.UserNotifications.Exists(ctx, ids...); err != nil {
				usecase.components.Logger.Error().
					Format("Failed to verify the existence of user notifications: '%s'. ", err).
					Field("ids", ids).Write()

				cErr = common_errors.InternalServerError()
				return
			}

			if len(exists) != len(ids) {
				usecase.components.Logger.Error().
					Text("Invalid value of the data received from the method of checking the existence of user notifications. ").
					Field("ids", ids).Write()

				cErr = common_errors.InternalServerError()
				return
			}

			for index, id := range ids {
				if !exists[index] {
					if tempCErr == nil {
						tempCErr = srv_errors.UserNotificationNotFound()
					}

					usecase.components.Logger.Error().
						Text("The user notification was not found. ").
						Field("id", id).Write()

					tempCErr.Details().SetField(
						new(err_details.FieldKey).AddArray("ids", index),
						new(err_messages.TextMessage).Text("Not found. "),
					)
				}
			}

			if tempCErr != nil {
				cErr = tempCErr
				return
			}
		}

		// Уже прочитано
		{
			var (
				tempCErr c_errors.Error
				err      error
				read     []bool
			)

			if read, err = usecase.repositories.UserNotifications.AlreadyRead(ctx, ids...); err != nil {
				usecase.components.Logger.Error().
					Format("It was not possible to check whether the user's notifications were read: '%s'. ", err).
					Field("ids", ids).Write()

				cErr = common_errors.InternalServerError()
				return
			}

			if len(read) != len(ids) {
				usecase.components.Logger.Error().
					Text("Invalid value of the data received from the method of checking read user notifications. ").
					Field("ids", ids).Write()

				cErr = common_errors.InternalServerError()
				return
			}

			for index, id := range ids {
				if read[index] {
					if tempCErr == nil {
						tempCErr = srv_errors.UserNotificationAlreadyRead()
					}

					usecase.components.Logger.Error().
						Text("The user notification already read. ").
						Field("id", id).Write()

					tempCErr.Details().SetField(
						new(err_details.FieldKey).AddArray("ids", index),
						new(err_messages.TextMessage).Text("Already read. "),
					)
				}
			}

			if tempCErr != nil {
				cErr = tempCErr
				return
			}
		}
	}

	// Чтение
	{
		var err error

		if err = usecase.repositories.UserNotifications.Read(ctx, recipientID, ids...); err != nil {
			usecase.components.Logger.Error().
				Format("Several user notifications could not be read: '%s'. ", err).
				Field("recipient_id", recipientID).
				Field("ids", ids).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("Reading multiple user notifications has been completed successfully. ").
			Field("recipient_id", recipientID).
			Field("ids", ids).Write()
	}

	return
}
