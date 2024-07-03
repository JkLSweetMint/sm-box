package rest_api

import (
	"context"
	"sm-box/internal/common/transports/rest_api/components/access_system"
	"sm-box/internal/services/i18n/transports/rest_api/config"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/http/postman"
)

const (
	loggerInitiator = "transports-[http]=rest_api"
)

// Engine - описание движка http rest api сервиса.
type Engine interface {
	Listen() (err error)
	Shutdown() (err error)
}

// New - создание движка http rest api сервиса.
func New(ctx context.Context) (eng Engine, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelTransport)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished(eng) }()
	}

	var ref = &engine{
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

		// AccessSystem
		{
			if ref.components.AccessSystem, err = access_system.New(ref.ctx, ref.conf.Components.AccessSystem); err != nil {
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
		ref.postman = postman.NewCollection(ref.conf.Engine.AppName, "")
	}

	// fiber
	{
		if err = ref.initFiberApp(); err != nil {
			return
		}
	}

	ref.components.Logger.Info().
		Text("The http rest engine has been created. ").
		Field("config", ref.conf).Write()

	eng = ref

	return
}
