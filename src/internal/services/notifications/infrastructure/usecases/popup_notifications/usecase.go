package popup_notifications_usecase

import (
	"context"
	"database/sql"
	"errors"
	common_errors "sm-box/internal/common/errors"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/notifications/components/notification_notifier"
	popup_notifications_repository "sm-box/internal/services/notifications/infrastructure/repositories/popup_notifications"
	"sm-box/internal/services/notifications/objects/constructors"
	"sm-box/internal/services/notifications/objects/entities"
	srv_errors "sm-box/internal/services/notifications/objects/errors"
	"sm-box/internal/services/notifications/objects/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	c_errors "sm-box/pkg/errors"
	err_details "sm-box/pkg/errors/entities/details"
	err_messages "sm-box/pkg/errors/entities/messages"
	"strings"
)

const (
	loggerInitiator = "infrastructure-[usecases]=popup_notifications"
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
	PopupNotifications interface {
		GetOne(ctx context.Context, id common_types.ID) (notification *entities.PopupNotification, err error)
		Get(ctx context.Context, ids ...common_types.ID) (list []*entities.PopupNotification, err error)

		CreateOne(ctx context.Context, constructor *constructors.PopupNotification) (id common_types.ID, err error)
		Create(ctx context.Context, constructors ...*constructors.PopupNotification) (ids []common_types.ID, err error)
	}
}

// components - компоненты логики.
type components struct {
	Logger               logger.Logger
	NotificationNotifier notification_notifier.Notifier
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

		// NotificationNotifier
		{
			if usecase.components.NotificationNotifier, err = notification_notifier.New(ctx); err != nil {
				return
			}
		}
	}

	// Репозитории
	{
		usecase.repositories = new(repositories)

		// PopupNotifications
		{
			if usecase.repositories.PopupNotifications, err = popup_notifications_repository.New(usecase.ctx); err != nil {
				return
			}
		}
	}

	usecase.components.Logger.Info().
		Format("A '%s' usecase has been created. ", "popup_notifications").
		Field("config", usecase.conf).Write()

	return
}

// CreateOne - создание всплывающего уведомления.
func (usecase *UseCase) CreateOne(ctx context.Context, constructor *constructors.PopupNotification) (notification *entities.PopupNotification, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, constructor)
		defer func() { trc.Error(cErr).FunctionCallFinished(notification) }()
	}

	usecase.components.Logger.Info().
		Text("The process of creating a popup notification has been started... ").
		Field("constructor", constructor).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of creating a custom notification is completed. ").
			Field("constructor", constructor).Write()
	}()

	var id common_types.ID

	// Валидация
	{
		// Пустой конструктор
		{
			if constructor == nil {
				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("constructor", constructor).Write()

				cErr = common_errors.InvalidArguments()
				cErr.Details().SetField(
					new(err_details.FieldKey).Add("constructor"),
					new(err_messages.TextMessage).Text("Is empty. "),
				)

				return
			}
		}

		// Данные конструктора
		{
			var tempCErr c_errors.Error

			if strings.TrimSpace(constructor.RecipientID) == "" {
				if tempCErr == nil {
					tempCErr = common_errors.InvalidArguments()
				}

				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("constructor", constructor).Write()

				tempCErr.Details().SetField(
					new(err_details.FieldKey).Add("recipient_id"),
					new(err_messages.TextMessage).Text("Invalid value. "),
				)
			}

			if constructor.Type != types.NotificationTypePopup {
				if tempCErr == nil {
					tempCErr = common_errors.InvalidArguments()
				}

				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("constructor", constructor).Write()

				tempCErr.Details().SetField(
					new(err_details.FieldKey).Add("type"),
					new(err_messages.TextMessage).Text("Invalid value. "),
				)
			}

			if len(constructor.Title) == 0 && constructor.TitleI18n.String() == "00000000-0000-0000-0000-000000000000" {
				if tempCErr == nil {
					tempCErr = common_errors.InvalidArguments()
				}

				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("constructor", constructor).Write()

				tempCErr.Details().SetField(
					new(err_details.FieldKey).Add("title"),
					new(err_messages.TextMessage).Text("Is empty. "),
				)
				tempCErr.Details().SetField(
					new(err_details.FieldKey).Add("title_i18n"),
					new(err_messages.TextMessage).Text("Is empty. "),
				)
			}

			if len(constructor.Text) == 0 && constructor.TextI18n.String() == "00000000-0000-0000-0000-000000000000" {
				if tempCErr == nil {
					tempCErr = common_errors.InvalidArguments()
				}

				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("constructor", constructor).Write()

				tempCErr.Details().SetField(
					new(err_details.FieldKey).Add("text"),
					new(err_messages.TextMessage).Text("Is empty. "),
				)
				tempCErr.Details().SetField(
					new(err_details.FieldKey).Add("text_i18n"),
					new(err_messages.TextMessage).Text("Is empty. "),
				)
			}

			if tempCErr != nil {
				cErr = tempCErr
				return
			}
		}
	}

	// Создание
	{
		var err error

		if id, err = usecase.repositories.PopupNotifications.CreateOne(ctx, constructor); err != nil {
			usecase.components.Logger.Error().
				Format("Failed to create a popup notification: '%s'. ", err).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The popup notification has been successfully created. ").
			Field("id", id).Write()
	}

	// Получение
	{
		var err error

		if notification, err = usecase.repositories.PopupNotifications.GetOne(ctx, id); err != nil {
			usecase.components.Logger.Error().
				Format("Could not get the popup notification by id: '%s'. ", err).Write()

			if errors.Is(err, sql.ErrNoRows) {
				cErr = srv_errors.UserNotificationNotFound()
				return
			}

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The popup notification have been successfully collected. ").
			Field("notification", notification).Write()
	}

	// Рассылка
	{
		env.Synchronization.WaitGroup.Add(1)

		go func() {
			defer env.Synchronization.WaitGroup.Done()

			usecase.components.NotificationNotifier.Notify(&notification_notifier.Notification{
				Type:      notification_notifier.NotificationTypeCreated,
				Recipient: notification.RecipientID,
				Data:      notification,
			})
		}()
	}

	return
}

// Create - создание всплывающих уведомлений.
func (usecase *UseCase) Create(ctx context.Context, constructors ...*constructors.PopupNotification) (notifications []*entities.PopupNotification, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, constructors)
		defer func() { trc.Error(cErr).FunctionCallFinished(notifications) }()
	}

	usecase.components.Logger.Info().
		Text("The process of creating several custom notifications has started... ").
		Field("constructors", constructors).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of creating multiple popup notifications is completed. ").
			Field("constructors", constructors).Write()
	}()

	var ids []common_types.ID

	// Валидация
	{
		// Пустые конструктора
		{
			if len(constructors) == 0 {
				usecase.components.Logger.Error().
					Text("An invalid argument value was passed. ").
					Field("constructors", constructors).Write()

				cErr = common_errors.InvalidArguments()
				cErr.Details().SetField(
					new(err_details.FieldKey).Add("constructors"),
					new(err_messages.TextMessage).Text("Is empty. "),
				)

				return
			}
		}

		// Данные конструкторов
		{
			var tempCErr c_errors.Error

			for index, constructor := range constructors {
				if strings.TrimSpace(constructor.RecipientID) == "" {
					if tempCErr == nil {
						tempCErr = common_errors.InvalidArguments()
					}

					usecase.components.Logger.Error().
						Text("An invalid argument value was passed. ").
						Field("constructor", constructor).Write()

					tempCErr.Details().SetField(
						new(err_details.FieldKey).Add("recipient_id"),
						new(err_messages.TextMessage).Text("Invalid value. "),
					)
				}

				if constructor.Type != types.NotificationTypePopup {
					if tempCErr == nil {
						tempCErr = common_errors.InvalidArguments()
					}

					usecase.components.Logger.Error().
						Text("An invalid argument value was passed. ").
						Field("constructor", constructor).Write()

					tempCErr.Details().SetField(
						new(err_details.FieldKey).AddArray("constructors", index).Add("type"),
						new(err_messages.TextMessage).Text("Invalid value. "),
					)
				}

				if len(constructor.Title) == 0 && constructor.TitleI18n.String() == "00000000-0000-0000-0000-000000000000" {
					if tempCErr == nil {
						tempCErr = common_errors.InvalidArguments()
					}

					usecase.components.Logger.Error().
						Text("An invalid argument value was passed. ").
						Field("constructor", constructor).Write()

					tempCErr.Details().SetField(
						new(err_details.FieldKey).AddArray("constructors", index).Add("title"),
						new(err_messages.TextMessage).Text("Is empty. "),
					)
					tempCErr.Details().SetField(
						new(err_details.FieldKey).AddArray("constructors", index).Add("title_i18n"),
						new(err_messages.TextMessage).Text("Is empty. "),
					)
				}

				if len(constructor.Text) == 0 && constructor.TextI18n.String() == "00000000-0000-0000-0000-000000000000" {
					if tempCErr == nil {
						tempCErr = common_errors.InvalidArguments()
					}

					usecase.components.Logger.Error().
						Text("An invalid argument value was passed. ").
						Field("constructor", constructor).Write()

					tempCErr.Details().SetField(
						new(err_details.FieldKey).AddArray("constructors", index).Add("text"),
						new(err_messages.TextMessage).Text("Is empty. "),
					)
					tempCErr.Details().SetField(
						new(err_details.FieldKey).AddArray("constructors", index).Add("text_i18n"),
						new(err_messages.TextMessage).Text("Is empty. "),
					)
				}
			}

			if tempCErr != nil {
				cErr = tempCErr
				return
			}
		}
	}

	// Создание
	{
		var err error

		if ids, err = usecase.repositories.PopupNotifications.Create(ctx, constructors...); err != nil {
			usecase.components.Logger.Error().
				Format("Failed to create a popup notifications: '%s'. ", err).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The popup notification has been successfully created. ").
			Field("ids", ids).Write()
	}

	// Получение
	{
		var err error

		if notifications, err = usecase.repositories.PopupNotifications.Get(ctx, ids...); err != nil {
			usecase.components.Logger.Error().
				Format("Could not get the popup notifications by ids: '%s'. ", err).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The popup notifications have been successfully collected. ").
			Field("notifications", notifications).Write()
	}

	// Рассылка
	{
		env.Synchronization.WaitGroup.Add(1)

		go func() {
			defer env.Synchronization.WaitGroup.Done()

			var list = make([]*notification_notifier.Notification, 0, len(notifications))

			for _, notification := range notifications {
				list = append(list, &notification_notifier.Notification{
					Type:      notification_notifier.NotificationTypeCreated,
					Recipient: notification.RecipientID,
					Data:      notification,
				})
			}

			usecase.components.NotificationNotifier.Notify(list...)
		}()
	}

	return
}
