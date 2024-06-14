package rest_api

import (
	"context"
	"sm-box/internal/app/transports/rest_api/components/access_system"
	"sm-box/internal/app/transports/rest_api/config"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/http/postman"
)

const (
	loggerInitiator = "transports-[http]=rest_api"
)

// Engine - описание движка http rest api коробки.
type Engine interface {
	Listen() (err error)
	Shutdown() (err error)
}

// New - создание движка http rest api коробки.
func New(ctx context.Context, conf *config.Config) (eng Engine, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelTransport)

		trc.FunctionCall(ctx, conf)
		defer func() { trc.Error(err).FunctionCallFinished(eng) }()
	}

	var ref = &engine{
		conf: conf,
		ctx:  ctx,
	}

	// Конфигурация
	{
		if err = ref.conf.FillEmptyFields().Validate(); err != nil {
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

	// Postman
	{
		ref.postman = postman.NewCollection(conf.Engine.AppName, "")
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
