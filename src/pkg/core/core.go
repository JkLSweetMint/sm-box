package core

import (
	"context"
	"sm-box/pkg/core/components/tracer"
)

// core - ядро системы.
type core struct {
	components *components
	tools      *tools

	ctx  context.Context
	conf *Config

	state interface {
		Shutdown() (err error)
		Boot() (err error)
		Serve() (err error)

		State() (state State)
	}
}

// Boot - загрузка ядра, построение системы, создание компонентов.
// Загрузить можно только ядро со статусом StateNew.
// Состояние ядра будет изменено на StateBooted.
func (c *core) Boot() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelCore)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	if err = c.state.Boot(); err != nil {
		c.Components().Logger().Error().
			Format("An error occurred while booting the system core: '%s'.  ", err).Write()
	}

	return
}

// Serve - запуск обслуживания ядра.
// Запустить обслуживание можно только ядро со статусом StateBooted.
// Состояние ядра будет изменено на StateServed.
func (c *core) Serve() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelCore)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	if err = c.state.Serve(); err != nil {
		c.Components().Logger().Error().
			Format("An error occurred while serving the system core: '%s'.  ", err).Write()
	}

	return
}

// Shutdown - завершение работы ядра.
// Остановить можно только ядро со статусом StateServed.
// Состояние ядра будет изменено на StateOff.
func (c *core) Shutdown() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelCore)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	if err = c.state.Shutdown(); err != nil {
		c.Components().Logger().Error().
			Format("An error occurred during the shutdown of the system core: '%s'.  ", err).Write()
	}

	return
}

// State - получение состояния ядра системы.
//
// Возможные варианты состояния ядра:
//  0. StateNil    - "Nil";
//  1. StateNew    - "New";
//  2. StateBooted - "Booted";
//  3. StateServed - "Served";
//  4. StateOff    - "Off";
func (c *core) State() (state State) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelCore)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(state) }()
	}

	if c.state == nil {
		return StateNil
	}

	return c.state.State()
}

// Ctx - получение контекста ядра системы.
func (c *core) Ctx() (ctx context.Context) {
	return c.ctx
}

// Components - получение компонентов ядра системы.
func (c *core) Components() Components {
	return c.components
}

// Tools - получение внутренних инструментов ядра системы.
func (c *core) Tools() Tools {
	return c.tools
}
