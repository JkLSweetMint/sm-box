package http_rest_api

import (
	"context"
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
	}

	// Контроллеры
	{
		ref.controllers = new(controllers)
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
