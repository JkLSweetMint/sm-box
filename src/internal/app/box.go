package app

import (
	"context"
	"sm-box/internal/app/transports/rest_api"
	"sm-box/pkg/core"
	"sm-box/pkg/core/addons/encryption_keys"
	"sm-box/pkg/core/addons/pid"
	"sm-box/pkg/core/components/configurator"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	"sm-box/pkg/core/tools/task_scheduler"
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
		defer func() { trc.Error(err).FunctionCallFinished(box_) }()
	}

	var ref = new(box)

	// Конфигурация
	{
		var c configurator.Configurator[*Config]

		ref.conf = new(Config).Default()

		if c, err = configurator.New[*Config](ref.conf); err != nil {
			return
		} else if err = c.Private().Profile(confProfile).Init(); err != nil {
			return
		}

		if err = ref.conf.FillEmptyFields().Validate(); err != nil {
			return
		}
	}

	// Ядро
	{
		if ref.core, err = core.New(); err != nil {
			return
		}
	}

	// Компоненты
	{
		ref.components = new(components)

		// Logger
		{
			if ref.components.logger, err = logger.New(env.Vars.SystemName); err != nil {
				return
			}
		}
	}

	// Контроллеры
	{
		ref.controllers = new(controllers)
	}

	// Транспортная часть
	{
		ref.transports = new(transports)

		if ref.transports.restApi, err = rest_api.New(ref.Ctx(), ref.conf.Transports.RestAPI); err != nil {
			return
		}
	}

	// Регистрация задач коробки
	{
		// Дополнения ядра
		{
			if err = ref.core.Tools().TaskScheduler().Register(pid.TaskCreatePIDFile); err != nil {
				ref.Components().Logger().Error().
					Format("Failed to register a task in task scheduler: '%s'. ", err)
			}

			if err = ref.core.Tools().TaskScheduler().Register(pid.TaskRemovePIDFile); err != nil {
				ref.Components().Logger().Error().
					Format("Failed to register a task in task scheduler: '%s'. ", err)
			}

			if err = ref.core.Tools().TaskScheduler().Register(encryption_keys.TaskInitEncryptionKeys); err != nil {
				ref.Components().Logger().Error().
					Format("Failed to register a task in task scheduler: '%s'. ", err)
			}
		}

		// Основные
		{
			if err = ref.core.Tools().TaskScheduler().Register(task_scheduler.Task{
				Name: "Starting the server maintenance. ",
				Type: task_scheduler.TaskServe,
				Func: ref.serve,
			}); err != nil {
				ref.Components().Logger().Error().
					Format("Failed to register a task in task scheduler: '%s'. ", err)
			}

			if err = ref.core.Tools().TaskScheduler().Register(task_scheduler.Task{
				Name: "Completion of server maintenance. ",
				Type: task_scheduler.TaskShutdown,
				Func: ref.shutdown,
			}); err != nil {
				ref.Components().Logger().Error().
					Format("Failed to register a task in task scheduler: '%s'. ", err)
			}
		}
	}

	// Построение ядра
	{
		if err = ref.core.Boot(); err != nil {
			return
		}
	}

	box_ = ref

	box_.Components().Logger().Info().
		Format("A '%s' has been created. ", env.Vars.SystemName).
		Field("config", ref.conf).Write()

	return
}
