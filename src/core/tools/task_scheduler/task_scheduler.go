package task_scheduler

import (
	"context"
	"sm-box/src/core/components/logger"
	"sm-box/src/core/components/tracer"
	"sm-box/src/core/env"
	"sync"
)

const (
	loggerInitiator = "core-[tools]=task_scheduler"
)

// Scheduler - описание инструмента ядра системы для выполнения запланированных задач.
type Scheduler interface {
	Register(t Task) (err error)
}

// New - создание инструмента ядра системы для выполнения запланированных задач.
func New(ctx context.Context) (sc Scheduler, c chan<- TaskType, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelCoreTool)

		trc.FunctionCall(ctx)
		trc.Error(err).FunctionCallFinished(sc, c)
	}

	var s = &scheduler{
		aggregate: &baseShelf{
			Tasks: make([]*Task, 0),
			rwMx:  new(sync.RWMutex),
		},
		channel: make(chan TaskType, 1),
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

	c = s.channel
	sc = s

	env.Synchronization.WaitGroup.Add(1)
	go s.tracking(ctx)

	return
}
