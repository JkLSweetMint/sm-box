package core

import (
	"context"
	"sm-box/src/core/components/tracer"
	"sm-box/src/core/env"
	"sm-box/src/core/tools/task_scheduler"
)

// stateServed - реализация ядра системы для состояния StateServed - "Served".
type stateServed struct {
	components *components
	tools      *tools
	channels   *channels

	ctx  context.Context
	conf *Config
}

// Boot - загрузка ядра, построение системы, создание компонентов.
//
// Вызов завершится с ошибкой т.к ядро уже построено.
func (c *stateServed) Boot() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelCore)

		trc.FunctionCall()
		trc.Error(err).FunctionCallFinished()
	}

	c.Components().Logger().Info().
		Text("The system core has started booting... ").Write()

	err = ErrSystemCoreIsAlreadyBooted

	c.Components().Logger().Error().
		Format("The core failed to boot system maintenance: '%s'. ", err).Write()

	return
}

// Serve - запуск обслуживания ядра.
//
// Вызов завершится с ошибкой т.к обслуживание системы ядром уже запущено.
func (c *stateServed) Serve() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelCore)

		trc.FunctionCall()
		trc.Error(err).FunctionCallFinished()
	}

	c.Components().Logger().Info().
		Text("The core starts system maintenance... ").Write()

	err = ErrSystemCoreIsAlreadyServed

	c.Components().Logger().Error().
		Format("The core failed to start system maintenance: '%s'. ", err).Write()

	return
}

// Shutdown - завершение работы ядра.
// Состояние ядра будет изменено на StateOff.
func (c *stateServed) Shutdown() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelCore)

		trc.FunctionCall()
		trc.Error(err).FunctionCallFinished()
	}

	// Вызов задачи планировщика - 'BeforeShutdown'.
	{
		c.channels.taskScheduler <- task_scheduler.TaskBeforeShutdown
	}

	c.Components().Logger().Info().
		Text("The core completes system maintenance... ").Write()

	// Вызов задачи планировщика - 'Shutdown'.
	{
		c.channels.taskScheduler <- task_scheduler.TaskShutdown
	}

	// Завершение работы
	{
		c.tools.closer.Cancel()
	}

	// Изменение состояния
	{
		if err = instance.(*core).updateState(StateOff); err != nil {
			c.Components().Logger().Error().
				Format("An error occurred during the core status update: '%s'. ", err).Write()
			return
		}
	}

	c.Components().Logger().Info().
		Text("The core has completed system maintenance. ").Write()

	// Вызов задачи планировщика - 'AfterShutdown'.
	{
		c.channels.taskScheduler <- task_scheduler.TaskAfterShutdown
	}

	env.Synchronization.WaitGroup.Wait()

	return
}

// State - получение состояния ядра системы.
//
// Будет возвращено значение StateServed - "Served".
func (c *stateServed) State() (state State) {
	return StateServed
}

// Components - получение компонентов ядра системы.
func (c *stateServed) Components() Components {
	return c.components
}

// Tools - получение внутренних инструментов ядра системы.
func (c *stateServed) Tools() Tools {
	return c.tools
}
