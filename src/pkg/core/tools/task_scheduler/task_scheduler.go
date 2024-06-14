package task_scheduler

import (
	"context"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sync"
)

const (
	loggerInitiator = "core-[tools]=task_scheduler"
)

// Scheduler - описание инструмента ядра системы для выполнения запланированных задач.
type Scheduler interface {
	Register(t Task) (err error)
	Call(e Event) (err error)
}

// New - создание инструмента ядра системы для выполнения запланированных задач.
func New(ctx context.Context) (sc Scheduler, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelCoreTool)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished(sc) }()
	}

	var s = &scheduler{
		aggregate: &baseShelf{
			Tasks: make([]Task, 0),
			rwMx:  new(sync.RWMutex),
		},
	}

	// Компоненты
	{
		s.components = new(components)

		// Logger
		{
			if s.components.Logger, err = logger.New(loggerInitiator); err != nil {
				return
			}
		}
	}

	sc = s

	return
}
