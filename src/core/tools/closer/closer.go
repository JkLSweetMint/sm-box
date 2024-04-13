package closer

import (
	"context"
	"sm-box/src/core/components/tracer"
	"sync"
)

// Closer - описание инструмента ядра системы отвечающий за корректное завершение работы системы.
type Closer interface {
	Wait()
	Cancel()
}

// New - создание инструмента ядра системы отвечающий за корректное завершение работы системы.
func New(conf *Config, ctx context.Context, wg *sync.WaitGroup) (cl Closer, ct context.Context) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelCoreTool)

		trc.FunctionCall(conf, ctx, wg)
		trc.FunctionCallFinished(cl, ct)
	}

	var c = &closer{
		conf: conf,

		wg: wg,
	}

	c.ctx, c.ctxCancel = context.WithCancel(ctx)

	cl = c
	ct = c.ctx

	return
}
