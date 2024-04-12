package closer

import (
	"context"
	"os"
	"os/signal"
	"sm-box/src/core/components/tracer"
	"sync"
)

// closer - компонент ядра системы отвечающий за корректное завершение работы системы.
type closer struct {
	conf *Config

	wg        *sync.WaitGroup
	ctx       context.Context
	ctxCancel context.CancelFunc
}

// Wait - ожидание завершения работы.
// Вызов этого метода замораживает выполнение до завершения работы всех сценариев использующих
// глобальный инструмент синхронизации WaitGroup.
func (c *closer) Wait() {
	// tracer
	{
		var trc = tracer.New(tracer.LevelCoreComponent)

		trc.FunctionCall()
		trc.FunctionCallFinished()
	}

	c.wg.Wait()
	c.ctxCancel()
}

// Cancel - сообщает системе о завершении работы.
func (c *closer) Cancel() {
	// tracer
	{
		var trc = tracer.New(tracer.LevelCoreComponent)

		trc.FunctionCall()
		trc.FunctionCallFinished()
	}

	c.ctxCancel()
}

// tracking - отслеживание сигналов для завершения работы.
func (c *closer) tracking() {
	var ch = make(chan os.Signal, 1)

	signal.Notify(ch, c.conf.Signals...)

	select {
	case <-ch:
		{
			c.Cancel()
		}
	case <-c.ctx.Done():
		return
	}
}
