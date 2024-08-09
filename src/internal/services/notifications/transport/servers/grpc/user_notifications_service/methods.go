package user_notifications_srv

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	common_errors "sm-box/internal/common/errors"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/notifications/objects/constructors"
	"sm-box/internal/services/notifications/objects/models"
	"sm-box/internal/services/notifications/objects/types"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	err_details "sm-box/pkg/errors/entities/details"
	err_messages "sm-box/pkg/errors/entities/messages"
	pb "sm-box/transport/proto/pb/golang/notifications"
)

// CreateOne - создание пользовательского уведомления.
func (srv *server) CreateOne(ctx context.Context, request *pb.UserNotificationsCreateOneRequest) (response *pb.UserNotificationsCreateOneResponse, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelTransportGrpc)

		trc.FunctionCall(ctx, request)
		defer func() { trc.Error(err).FunctionCallFinished(response) }()
	}

	response = new(pb.UserNotificationsCreateOneResponse)

	// Проверка данных
	{
		if request.Constructor == nil {
			srv.components.Logger.Error().
				Text("Invalid arguments were received. ").Write()

			var cErr = common_errors.InvalidArguments()
			cErr.Details().SetField(
				new(err_details.FieldKey).Add("constructor"),
				new(err_messages.TextMessage).Text("Is empty. "),
			)

			err = cErr
			return
		}
	}

	var notification *models.UserNotificationInfo

	// Обработка
	{
		var constructor *constructors.UserNotification

		// Подготовка конструктора
		{
			constructor = &constructors.UserNotification{
				Type: types.NotificationType(request.Constructor.Type),

				SenderID:    common_types.ID(request.Constructor.SenderID),
				RecipientID: common_types.ID(request.Constructor.RecipientID),

				Title:     request.Constructor.Title,
				TitleI18n: uuid.UUID{},

				Text:     request.Constructor.Text,
				TextI18n: uuid.UUID{},
			}

			constructor.TitleI18n, _ = uuid.Parse(request.Constructor.TitleI18N)
			constructor.TextI18n, _ = uuid.Parse(request.Constructor.TextI18N)

			constructor.FillEmptyFields()
		}

		if notification, err = srv.controllers.UserNotifications.CreateOne(ctx, constructor); err != nil {
			srv.components.Logger.Error().
				Format("Failed to create a user notification: '%s'. ", err).Write()

			return
		}
	}

	// Преобразование данных в структуры grpc
	{
		if notification != nil {
			response.Notification = &pb.UserNotification{
				ID:   uint64(notification.ID),
				Type: string(notification.Type),

				SenderID:    uint64(notification.SenderID),
				RecipientID: uint64(notification.RecipientID),

				Title:     notification.Title,
				TitleI18N: notification.TitleI18n.String(),

				Text:     notification.Text,
				TextI18N: notification.TextI18n.String(),

				CreatedTimestamp: timestamppb.New(notification.CreatedTimestamp),
				ReadTimestamp:    timestamppb.New(notification.ReadTimestamp),
				RemovedTimestamp: timestamppb.New(notification.RemovedTimestamp),
			}
		}
	}

	return
}

// Create - создание пользовательских уведомлений.
func (srv *server) Create(ctx context.Context, request *pb.UserNotificationsCreateRequest) (response *pb.UserNotificationsCreateResponse, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelTransportGrpc)

		trc.FunctionCall(ctx, request)
		defer func() { trc.Error(err).FunctionCallFinished(response) }()
	}

	response = new(pb.UserNotificationsCreateResponse)

	// Проверка данных
	{
		if len(request.Constructors) == 0 {
			srv.components.Logger.Error().
				Text("Invalid arguments were received. ").Write()

			var cErr = common_errors.InvalidArguments()
			cErr.Details().SetField(
				new(err_details.FieldKey).Add("constructors"),
				new(err_messages.TextMessage).Text("Is empty. "),
			)

			err = cErr
			return
		}

		var cErr c_errors.Grpc

		for index, constructor := range request.Constructors {
			if constructor == nil {
				srv.components.Logger.Error().
					Text("Invalid arguments were received. ").Write()

				if cErr == nil {
					cErr = c_errors.ToGrpc(common_errors.InvalidArguments())
				}

				cErr.Details().SetField(
					new(err_details.FieldKey).AddArray("constructors", index),
					new(err_messages.TextMessage).Text("Is empty. "),
				)
			}
		}

		if cErr != nil {
			err = cErr
			return
		}
	}

	var notifications []*models.UserNotificationInfo

	// Обработка
	{
		var list = make([]*constructors.UserNotification, 0)

		// Подготовка конструктора
		{
			for _, c := range request.Constructors {
				var constructor = &constructors.UserNotification{
					Type: types.NotificationType(c.Type),

					SenderID:    common_types.ID(c.SenderID),
					RecipientID: common_types.ID(c.RecipientID),

					Title:     c.Title,
					TitleI18n: uuid.UUID{},

					Text:     c.Text,
					TextI18n: uuid.UUID{},
				}

				constructor.TitleI18n, _ = uuid.Parse(c.TitleI18N)
				constructor.TextI18n, _ = uuid.Parse(c.TextI18N)

				constructor.FillEmptyFields()

				list = append(list, constructor)
			}
		}

		if notifications, err = srv.controllers.UserNotifications.Create(ctx, list...); err != nil {
			srv.components.Logger.Error().
				Format("Failed to create a user notifications: '%s'. ", err).Write()

			return
		}
	}

	// Преобразование данных в структуры grpc
	{
		response.Notifications = make([]*pb.UserNotification, 0, len(notifications))

		for _, notification := range notifications {
			response.Notifications = append(response.Notifications, &pb.UserNotification{
				ID:   uint64(notification.ID),
				Type: string(notification.Type),

				SenderID:    uint64(notification.SenderID),
				RecipientID: uint64(notification.RecipientID),

				Title:     notification.Title,
				TitleI18N: notification.TitleI18n.String(),

				Text:     notification.Text,
				TextI18N: notification.TextI18n.String(),

				CreatedTimestamp: timestamppb.New(notification.CreatedTimestamp),
				ReadTimestamp:    timestamppb.New(notification.ReadTimestamp),
				RemovedTimestamp: timestamppb.New(notification.RemovedTimestamp),
			})
		}
	}

	return
}
