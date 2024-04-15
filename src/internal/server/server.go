package server

import (
	"context"
	"sm-box/src/core"
	"sm-box/src/core/components/configurator"
	"sm-box/src/core/components/logger"
	"sm-box/src/core/components/tracer"
	"sm-box/src/core/env"
	"sm-box/src/core/tools/task_scheduler"
)

// Server - описание сервера.
type Server interface {
	Serve() (err error)
	Shutdown() (err error)

	State() (state core.State)
	Ctx() (ctx context.Context)

	Components() Components
	Controllers() Controllers
}

// New - создание сервера.
func New() (srv Server, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain)

		trc.FunctionCall()
		trc.Error(err).FunctionCallFinished(srv)
	}

	var s = new(server)

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

		s.conf = conf
	}

	// Ядро
	{
		if s.core, err = core.New(); err != nil {
			return
		}
	}

	// Компоненты
	{
		s.components = new(components)

		// Logger
		{
			if s.components.logger, err = logger.New(env.Vars.SystemName); err != nil {
				return
			}
		}
	}

	// Контроллеры
	{
		s.controllers = new(controllers)
	}

	// Регистрация задач сервера
	{
		if err = s.core.Tools().TaskScheduler().Register(task_scheduler.Task{
			Name: "Starting the server maintenance. ",
			Type: task_scheduler.TaskServe,
			Func: s.serve,
		}); err != nil {
			s.Components().Logger().Error().
				Format("Failed to register a task in task scheduler: '%s'. ", err)
		}

		if err = s.core.Tools().TaskScheduler().Register(task_scheduler.Task{
			Name: "Completion of server maintenance. ",
			Type: task_scheduler.TaskShutdown,
			Func: s.shutdown,
		}); err != nil {
			s.Components().Logger().Error().
				Format("Failed to register a task in task scheduler: '%s'. ", err)
		}
	}

	// Построение ядра
	{
		if err = s.core.Boot(); err != nil {
			return
		}
	}

	srv = s

	srv.Components().Logger().Info().
		Format("A '%s' has been created. ", env.Vars.SystemName).
		Field("config", s.conf).Write()

	return
}
