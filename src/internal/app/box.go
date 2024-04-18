package app

import (
	"context"
	"sm-box/src/internal/app/transports/graphql"
	"sm-box/src/pkg/core"
	"sm-box/src/pkg/core/components/configurator"
	"sm-box/src/pkg/core/components/logger"
	"sm-box/src/pkg/core/components/tracer"
	"sm-box/src/pkg/core/env"
	"sm-box/src/pkg/core/tools/task_scheduler"
)

// Box - описание функционала коробки.
type Box interface {
	Serve() (err error)
	Shutdown() (err error)

	State() (state core.State)
	Ctx() (ctx context.Context)

	Components() Components
	Controllers() Controllers
	Transports() Transports
}

// New - создание коробки.
func New() (box_ Box, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain)

		trc.FunctionCall()
		trc.Error(err).FunctionCallFinished(box_)
	}

	var bx = new(box)

	// Конфигурация
	{
		var (
			conf = new(Config)
			c    configurator.Configurator[*Config]
		)

		if c, err = configurator.New[*Config](conf); err != nil {
			return
		} else if err = c.Private().Profile(confProfile).Read(); err != nil {
			return
		}

		bx.conf = conf
	}

	// Ядро
	{
		if bx.core, err = core.New(); err != nil {
			return
		}
	}

	// Компоненты
	{
		bx.components = new(components)

		// Logger
		{
			if bx.components.logger, err = logger.New(env.Vars.SystemName); err != nil {
				return
			}
		}
	}

	// Контроллеры
	{
		bx.controllers = new(controllers)
	}

	// Транспортная часть
	{
		bx.transports = new(transports)

		if bx.transports.graphql, err = graphql.New(bx.Ctx()); err != nil {
			return
		}
	}

	// Регистрация задач коробки
	{
		if err = bx.core.Tools().TaskScheduler().Register(task_scheduler.Task{
			Name: "Starting the server maintenance. ",
			Type: task_scheduler.TaskServe,
			Func: bx.serve,
		}); err != nil {
			bx.Components().Logger().Error().
				Format("Failed to register a task in task scheduler: '%s'. ", err)
		}

		if err = bx.core.Tools().TaskScheduler().Register(task_scheduler.Task{
			Name: "Completion of server maintenance. ",
			Type: task_scheduler.TaskShutdown,
			Func: bx.shutdown,
		}); err != nil {
			bx.Components().Logger().Error().
				Format("Failed to register a task in task scheduler: '%s'. ", err)
		}
	}

	// Построение ядра
	{
		if err = bx.core.Boot(); err != nil {
			return
		}
	}

	box_ = bx

	box_.Components().Logger().Info().
		Format("A '%s' has been created. ", env.Vars.SystemName).
		Field("config", bx.conf).Write()

	return
}
