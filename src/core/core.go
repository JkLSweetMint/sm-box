package core

import (
	"context"
	"sm-box/src/core/components/configurator"
	"sm-box/src/core/components/logger"
	"sm-box/src/core/components/tracer"
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
	Shutdown() (err error)
	Boot() (err error)
	Serve() (err error)

	State() (state State)
	Ctx() (ctx context.Context)

	Components() interface {
		Logger() logger.Logger
	}
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
		var trc = tracer.New(tracer.LevelMain, tracer.LevelCore)

		trc.FunctionCall()
		trc.Error(err).FunctionCallFinished(cr)
	}

	var created bool

	once.Do(func() {
		var instance_ = &core{
			ctx: context.Background(),
		}

		// Конфигурация
		{
			var conf = new(Config)

			var c configurator.Configurator[*Config]

			if c, err = configurator.New[*Config](conf); err != nil {
				return
			} else if err = c.Private().Profile(confProfile).Read(); err != nil {
				return
			}

			instance_.conf = conf
		}

		// Компоненты
		{
			instance_.components = new(components)

			// Logger
			{
				if instance_.components.logger, err = logger.New(loggerInitiator); err != nil {
					return
				}
			}
		}

		// Состояние
		{
			instance_.state = &stateNew{
				components: instance_.components,
				ctx:        instance_.ctx,
				conf:       instance_.conf,
			}
		}

		instance = instance_

		instance.Components().Logger().Info().
			Text("The system core instance has been created. ").
			Field("state", instance_.State()).
			Field("config", instance_.conf).Write()

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
