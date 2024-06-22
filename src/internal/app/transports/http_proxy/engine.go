package http_proxy

import (
	"context"
	"sm-box/internal/app/transports/http_proxy/config"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/http/postman"
)

const (
	loggerInitiator = "transports-[http]=http_proxy"
)

// Engine - описание движка http proxy коробки.
type Engine interface {
	Listen() (err error)
	Shutdown() (err error)
}

// New - создание движка http proxy коробки.
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
		ref.conf = new(config.Config)

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
		Text("The http proxy engine has been created. ").
		Field("config", ref.conf).Write()

	eng = ref

	return
}
