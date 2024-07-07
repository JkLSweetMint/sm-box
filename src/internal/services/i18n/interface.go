package service

import (
	"context"
	"sm-box/internal/services/i18n/transport/servers/http_rest_api"
	"sm-box/pkg/core"
	"sm-box/pkg/core/addons/pid"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	"sm-box/pkg/core/tools/task_scheduler"
)

// Service - описание функционала сервиса.
type Service interface {
	Serve() (err error)
	Shutdown() (err error)

	State() (state core.State)
	Ctx() (ctx context.Context)

	Components() Components
	Controllers() Controllers
	Transport() Transport
}

// New - создание сервиса.
func New() (srv_ Service, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished(srv_) }()
	}

	var ref = new(service)

	// Конфигурация
	{
		ref.conf = new(Config).Default()

		if err = ref.conf.Read(); err != nil {
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
		ref.transport = new(transport)
		ref.transport.servers = new(transportServers)
		ref.transport.gateways = new(transportGateways)

		// Сервера
		{
			if ref.transport.servers.httpRestApi, err = http_rest_api.New(ref.Ctx()); err != nil {
				return
			}
		}
	}

	// Регистрация задач сервиса
	{
		// Дополнения ядра
		{
			if err = ref.core.Tools().TaskScheduler().Register(pid.TaskCreatePIDFile); err != nil {
				ref.Components().Logger().Error().
					Format("Failed to register a task in task scheduler: '%s'. ", err).Write()
			}

			if err = ref.core.Tools().TaskScheduler().Register(pid.TaskRemovePIDFile); err != nil {
				ref.Components().Logger().Error().
					Format("Failed to register a task in task scheduler: '%s'. ", err).Write()
			}
		}

		// Основные
		{
			if err = ref.core.Tools().TaskScheduler().Register(&task_scheduler.ImmediateTask{
				Name:  "Starting the server maintenance. ",
				Event: task_scheduler.EventServe,
				Func:  ref.serve,
			}); err != nil {
				ref.Components().Logger().Error().
					Format("Failed to register a task in task scheduler: '%s'. ", err).Write()
			}

			if err = ref.core.Tools().TaskScheduler().Register(&task_scheduler.ImmediateTask{
				Name:  "Completion of server maintenance. ",
				Event: task_scheduler.EventShutdown,
				Func:  ref.shutdown,
			}); err != nil {
				ref.Components().Logger().Error().
					Format("Failed to register a task in task scheduler: '%s'. ", err).Write()
			}
		}
	}

	// Построение ядра
	{
		if err = ref.core.Boot(); err != nil {
			return
		}
	}

	srv_ = ref

	srv_.Components().Logger().Info().
		Format("A '%s' has been created. ", env.Vars.SystemName).
		Field("config", ref.conf).Write()

	return
}
