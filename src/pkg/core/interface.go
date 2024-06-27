package core

import (
	"context"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/tools/closer"
	"sm-box/pkg/core/tools/task_scheduler"
	"sync"
)

const (
	loggerInitiator = "core"
)

var (
	once     = new(sync.Once)
	instance Core
)

// Core - описание ядра системы.
type Core interface {
	Boot() (err error)
	Serve() (err error)
	Shutdown() (err error)

	State() (state State)
	Ctx() (ctx context.Context)

	Components() Components
	Tools() Tools
}

// New - создание ядра системы.
// Может быть создан только один объект ядра!
//
// Ядро может быть в следующих состояний:
//   - StateNew    - "New"
//   - StateBooted - "Booted"
//   - StateServed - "Served"
//   - StateOff    - "Off"
func New() (cr Core, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelCore)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished(cr) }()
	}

	var created bool

	once.Do(func() {
		var ref = &core{
			ctx: context.Background(),
		}

		// Конфигурация
		{
			ref.conf = new(Config)

			if err = ref.conf.Read(); err != nil {
				return
			}
		}

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

		// Инструменты
		{
			ref.tools = new(tools)

			// Closer
			{
				if ref.tools.closer, ref.ctx = closer.New(ref.ctx, ref.conf.Tools.Closer); err != nil {
					return
				}
			}

			// TaskScheduler
			{
				if ref.tools.taskScheduler, err = task_scheduler.New(ref.ctx); err != nil {
					return
				}
			}
		}

		instance = ref

		// Состояние
		{
			if err = ref.updateState(StateNew); err != nil {
				return
			}
		}

		instance.Components().Logger().Info().
			Text("The system core instance has been created. ").
			Field("state", ref.State()).
			Field("config", ref.conf).Write()

		// Вызов задачи планировщика - 'BeforeNew'.
		{
			if err = ref.tools.taskScheduler.Call(task_scheduler.EventBeforeNew); err != nil {
				return
			}
		}

		created = true
	})

	if err != nil {
		return
	}

	cr = instance

	if !created {
		instance.Components().Logger().Info().
			Text("The system core instance was received. ").
			Field("state", instance.State()).Write()
	}

	return
}
