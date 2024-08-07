package http_rest_api

import (
	"context"
	"sm-box/internal/services/notifications/components/notification_notifier"
	popup_notifications_adapter "sm-box/internal/services/notifications/infrastructure/adapters/popup_notifications"
	user_notifications_adapter "sm-box/internal/services/notifications/infrastructure/adapters/user_notifications"
	"sm-box/internal/services/notifications/transport/servers/http/rest_api/config"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/http/postman"
)

const (
	loggerInitiator = "transports-[servers]-[http]=rest_api"
)

// Server - описание http rest api сервера.
type Server interface {
	Listen() (err error)
	Shutdown() (err error)
}

// New - создание сервера.
func New(ctx context.Context) (srv Server, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelTransport)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished(srv) }()
	}

	var ref = &server{
		ctx: ctx,
	}

	// Конфигурация
	{
		ref.conf = new(config.Config).Default()

		if err = ref.conf.Read(); err != nil {
			return
		}
	}

	// Компоненты
	{
		ref.components = new(components)

		// Logger
		{
			if ref.components.Logger, err = logger.New(loggerInitiator); err != nil {
				return
			}
		}

		// NotificationNotifier
		{
			if ref.components.NotificationNotifier, err = notification_notifier.New(ref.ctx); err != nil {
				return
			}
		}
	}

	// Контроллеры
	{
		ref.controllers = new(controllers)

		if ref.controllers.UserNotifications, err = user_notifications_adapter.New_RestAPI(ref.ctx); err != nil {
			return
		}

		if ref.controllers.PopupNotifications, err = popup_notifications_adapter.New_RestAPI(ref.ctx); err != nil {
			return
		}
	}

	// Postman
	{
		ref.postman = postman.NewCollection(ref.conf.Server.AppName, "")
	}

	// fiber
	{
		if err = ref.initFiberServer(); err != nil {
			return
		}
	}

	ref.components.Logger.Info().
		Text("The http rest api server has been created. ").
		Field("config", ref.conf).Write()

	srv = ref

	return
}
