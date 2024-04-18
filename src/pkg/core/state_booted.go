package core

import (
	"context"
	"sm-box/src/pkg/core/components/tracer"
	"sm-box/src/pkg/core/tools/task_scheduler"
)

// stateBooted - реализация ядра системы для состояния StateBooted - "Booted".
type stateBooted struct {
	components *components
	tools      *tools
	channels   *channels

	ctx  context.Context
	conf *Config
}

// Boot - загрузка ядра, построение системы, создание компонентов.
//
// Вызов завершится с ошибкой т.к ядро уже построено.
func (c *stateBooted) Boot() (err error) {
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
// Состояние ядра будет изменено на StateServed.
func (c *stateBooted) Serve() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelCore)

		trc.FunctionCall()
		trc.Error(err).FunctionCallFinished()
	}

	// Вызов задачи планировщика - 'BeforeServe'.
	{
		c.channels.taskScheduler <- task_scheduler.TaskBeforeServe
	}

	c.Components().Logger().Info().
		Text("The core starts system maintenance... ").Write()

	// Вызов задачи планировщика - 'Serve'.
	{
		c.channels.taskScheduler <- task_scheduler.TaskServe
	}

	// Изменение состояния
	{
		if err = instance.(*core).updateState(StateServed); err != nil {
			c.Components().Logger().Error().
				Format("An error occurred during the core status update: '%s'. ", err).Write()
			return
		}
	}

	c.Components().Logger().Info().
		Text("The core has started system maintenance. ").Write()

	// Вызов задачи планировщика - 'AfterServe'.
	{
		c.channels.taskScheduler <- task_scheduler.TaskAfterServe
	}

	c.tools.closer.Wait()

	return
}

// Shutdown - завершение работы ядра.
//
// Вызов завершится с ошибкой т.к ядро не было запущено.
func (c *stateBooted) Shutdown() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelCore)

		trc.FunctionCall()
		trc.Error(err).FunctionCallFinished()
	}

	c.Components().Logger().Info().
		Text("The core completes system maintenance... ").Write()

	err = ErrSystemCoreIsNotServe

	c.Components().Logger().Error().
		Format("The core failed to complete system maintenance: '%s'. ", err).Write()

	return
}

// State - получение состояния ядра системы.
//
// Будет возвращено значение StateBooted - "Booted".
func (c *stateBooted) State() (state State) {
	return StateBooted
}

// Components - получение компонентов ядра системы.
func (c *stateBooted) Components() Components {
	return c.components
}

// Tools - получение внутренних инструментов ядра системы.
func (c *stateBooted) Tools() Tools {
	return c.tools
}
