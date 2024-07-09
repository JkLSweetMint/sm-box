package http_rest_api

import (
	"context"
	languages_adapter "sm-box/internal/services/i18n/infrastructure/adapters/languages"
	texts_adapter "sm-box/internal/services/i18n/infrastructure/adapters/texts"
	"sm-box/internal/services/i18n/transport/servers/http/rest_api/config"
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

		// Texts
		{
			if ref.controllers.Texts, err = texts_adapter.New_HttpRestAPI(ctx); err != nil {
				return
			}
		}

		// Languages
		{
			if ref.controllers.Languages, err = languages_adapter.New_HttpRestAPI(ctx); err != nil {
				return
			}
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
