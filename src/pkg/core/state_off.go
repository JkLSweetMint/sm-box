package core

import (
	"context"
	"sm-box/pkg/core/components/tracer"
)

// stateOff - реализация ядра системы для состояния  StateOff - "Off".
type stateOff struct {
	components *components
	tools      *tools

	ctx  context.Context
	conf *Config
}

// Boot - загрузка ядра, построение системы, создание компонентов.
//
// Вызов завершится с ошибкой т.к ядро уже построено.
func (c *stateOff) Boot() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelCore)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
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
// Вызов завершится с ошибкой т.к обслуживание системы ядром уже завершено и запуск не возможен.
func (c *stateOff) Serve() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelCore)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	c.Components().Logger().Info().
		Text("The core starts system maintenance... ").Write()

	err = ErrSystemCoreAlreadyClosedStartIsNotPossible

	c.Components().Logger().Error().
		Format("The core failed to start system maintenance: '%s'. ", err).Write()

	return
}

// Shutdown - завершение работы ядра.
//
// Вызов завершится с ошибкой т.к обслуживание системы ядром уже завершено.
func (c *stateOff) Shutdown() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelCore)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	c.Components().Logger().Info().
		Text("The core completes system maintenance... ").Write()

	err = ErrISystemCoresAlreadyTurnedOff

	c.Components().Logger().Error().
		Format("The core failed to complete system maintenance: '%s'. ", err).Write()

	return
}

// State - получение состояния ядра системы.
//
// Будет возвращено значение StateOff - "Off".
func (c *stateOff) State() (state State) {
	return StateOff
}

// Components - получение компонентов ядра системы.
func (c *stateOff) Components() Components {
	return c.components
}

// Tools - получение внутренних инструментов ядра системы.
func (c *stateOff) Tools() Tools {
	return c.tools
}
