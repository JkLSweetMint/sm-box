package rest_api

import (
	"context"
	"sm-box/src/internal/app/transports/rest_api/config"
	"sm-box/src/pkg/core/components/logger"
	"sm-box/src/pkg/core/components/tracer"
)

const (
	loggerInitiator = "transports-[http]=rest_api"
)

type Engine interface {
	Listen() (err error)
	Shutdown() (err error)
}

func New(ctx context.Context, conf *config.Config) (eng Engine, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelTransport)

		trc.FunctionCall(ctx, conf)
		trc.Error(err).FunctionCallFinished(eng)
	}

	var ref = &engine{
		conf: conf,
		ctx:  ctx,
	}

	ref.conf.FillEmptyFields()

	// Компоненты
	{
		ref.components = new(components)

		// Logger
		{
			if ref.components.logger, err = logger.New(loggerInitiator); err != nil {
				return
			}
		}
	}

	// fiber
	{
		if err = ref.initFiberApp(); err != nil {
			return
		}
	}

	ref.components.logger.Info().
		Text("The http rest engine has been created. ").
		Field("config", ref.conf).Write()

	eng = ref

	return
}
