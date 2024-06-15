package init_cli

import (
	"sm-box/pkg/core"
	"sm-box/pkg/core/addons/encryption_keys"
	"sm-box/pkg/core/addons/pid"
	"sm-box/pkg/core/components/configurator"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	"sm-box/pkg/core/tools/task_scheduler"
)

// CLI - описание функционала CLI для управления инициализации системы.
type CLI interface {
	Exec() (err error)
}

// New - создание CLI.
func New() (cli_ CLI, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished(cli_) }()
	}

	var ref = new(cli)

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
			if ref.components.Logger, err = logger.New(env.Vars.SystemName); err != nil {
				return
			}
		}
	}

	// Регистрация задач
	{
		// Дополнения ядра
		{
			if err = ref.core.Tools().TaskScheduler().Register(pid.TaskCreatePIDFile); err != nil {
				ref.components.Logger.Error().
					Format("Failed to register a task in task scheduler: '%s'. ", err).Write()
			}

			if err = ref.core.Tools().TaskScheduler().Register(pid.TaskRemovePIDFile); err != nil {
				ref.components.Logger.Error().
					Format("Failed to register a task in task scheduler: '%s'. ", err).Write()
			}

			if err = ref.core.Tools().TaskScheduler().Register(encryption_keys.TaskInitEncryptionKeys); err != nil {
				ref.components.Logger.Error().
					Format("Failed to register a task in task scheduler: '%s'. ", err).Write()
			}
		}

		// Основные
		{
			if err = ref.core.Tools().TaskScheduler().Register(&task_scheduler.ImmediateTask{
				Name:  "Starting the CLI maintenance. ",
				Event: task_scheduler.EventAfterServe,
				Func:  ref.exec,
			}); err != nil {
				ref.components.Logger.Error().
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

	cli_ = ref

	ref.components.Logger.Info().
		Format("A '%s' has been created. ", env.Vars.SystemName).
		Field("config", ref.conf).Write()

	return
}
