package core

import (
	"context"
	"sm-box/src/core/components/logger"
	"sm-box/src/core/components/tracer"
	"sm-box/src/core/tools/task_scheduler"
)

// stateNew - реализация ядра системы для состояния StateNew - "New".
type stateNew struct {
	components *components
	tools      *tools

	ctx  context.Context
	conf *Config
}

// Boot - загрузка ядра, построение системы, создание компонентов.
// Состояние ядра будет изменено на StateBooted.
func (c *stateNew) Boot() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelCore)

		trc.FunctionCall()
		trc.Error(err).FunctionCallFinished()
	}

	c.Components().Logger().Info().
		Text("The system core has started booting... ").Write()

	// Изменение состояния
	{
		if err = instance.(*core).updateState(StateBooted); err != nil {
			c.Components().Logger().Error().
				Format("An error occurred during the core status update: '%s'. ", err).Write()
			return
		}
	}

	c.Components().Logger().Info().
		Text("The system core has been booted. ").Write()

	return
}

// Serve - запуск обслуживания ядра.
//
// Вызов завершится с ошибкой т.к ядро не было построено.
func (c *stateNew) Serve() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelCore)

		trc.FunctionCall()
		trc.Error(err).FunctionCallFinished()
	}

	c.Components().Logger().Info().
		Text("The core starts system maintenance... ").Write()

	err = ErrSystemCoreIsNotBooted

	c.Components().Logger().Error().
		Format("The core failed to start system maintenance: '%s'. ", err).Write()

	return
}

// Shutdown - завершение работы ядра.
//
// Вызов завершится с ошибкой т.к ядро не было запущено.
func (c *stateNew) Shutdown() (err error) {
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
// Будет возвращено значение StateNew - "New".
func (c *stateNew) State() (state State) {
	return StateNew
}

// Components - получение компонентов ядра системы.
func (c *stateNew) Components() interface {
	Logger() logger.Logger
} {
	return c.components
}

// Tools - получение внутренних инструментов ядра системы.
func (c *stateNew) Tools() interface {
	TaskScheduler() task_scheduler.Scheduler
} {
	return c.tools
}
