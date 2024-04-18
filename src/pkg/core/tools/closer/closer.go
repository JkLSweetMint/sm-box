package closer

import (
	"context"
	"sm-box/src/pkg/core/components/tracer"
	"sm-box/src/pkg/core/env"
)

// Closer - описание инструмента ядра системы отвечающий за корректное завершение работы системы.
type Closer interface {
	Wait()
	Cancel()
}

// New - создание инструмента ядра системы отвечающий за корректное завершение работы системы.
func New(conf *Config, ctx context.Context) (cl Closer, ct context.Context) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelCoreTool)

		trc.FunctionCall(conf, ctx)
		trc.FunctionCallFinished(cl, ct)
	}

	var c = &closer{
		conf: conf,

		stop: make(chan struct{}, 5),
	}

	c.ctx, c.ctxCancel = context.WithCancel(ctx)

	cl = c
	ct = c.ctx

	env.Synchronization.WaitGroup.Add(1)

	go func() {
		defer env.Synchronization.WaitGroup.Done()

		c.tracking()
	}()

	return
}
