package closer

import (
	"context"
	"os"
	"os/signal"
	"sm-box/src/pkg/core/components/tracer"
)

// closer - инструмент ядра системы отвечающий за корректное завершение работы системы.
type closer struct {
	conf *Config

	stop      chan struct{}
	ctx       context.Context
	ctxCancel context.CancelFunc
}

// Wait - ожидание завершения работы.
// Вызов этого метода замораживает выполнение до завершения работы всех сценариев использующих
// глобальный инструмент синхронизации WaitGroup.
func (c *closer) Wait() {
	// tracer
	{
		var trc = tracer.New(tracer.LevelCoreTool)

		trc.FunctionCall()
		trc.FunctionCallFinished()
	}

	<-c.stop
}

// Cancel - сообщает системе о завершении работы.
func (c *closer) Cancel() {
	// tracer
	{
		var trc = tracer.New(tracer.LevelCoreTool)

		trc.FunctionCall()
		trc.FunctionCallFinished()
	}

	c.ctxCancel()

	c.stop <- struct{}{}
}

// tracking - отслеживание сигналов для завершения работы.
func (c *closer) tracking() {
	// tracer
	{
		var trc = tracer.New(tracer.LevelCoreTool)

		trc.FunctionCall()
		trc.FunctionCallFinished()
	}

	var (
		ch      = make(chan os.Signal, 1)
		signals = make([]os.Signal, len(c.conf.Signals))
	)

	for i, sig := range c.conf.Signals {
		signals[i] = sig
	}

	signal.Notify(ch, signals...)

	select {
	case <-ch:
		c.Cancel()
	case <-c.ctx.Done():
		c.Cancel()
	}
}
