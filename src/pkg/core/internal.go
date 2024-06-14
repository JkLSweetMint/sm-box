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
		var trc = tracer.New(tracer.LevelMain, tracer.LevelCore)

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
		var trc = tracer.New(tracer.LevelMain, tracer.LevelCore)

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
		var trc = tracer.New(tracer.LevelMain, tracer.LevelCore)

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

// updateState - обновление состояния ядра.
func (c *core) updateState(state State) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelCore)

		trc.FunctionCall(state)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var old = c.State()

	defer func() {
		if err != nil {
			c.Components().Logger().Error().
				Format("An error occurred during the core status update: '%s'. ", err).
				Field("old_state", old).
				Field("new_state", state).Write()
			return
		}

		c.Components().Logger().Info().
			Format("The state of the system core has been changed from '%s' to '%s'. ",
				old,
				instance.State(),
			).Write()
	}()

	switch state {
	case StateNew:
		{
			if old == StateNil {
				c.state = &stateNew{
					components: c.components,
					tools:      c.tools,

					ctx:  c.ctx,
					conf: c.conf,
				}
				return
			}
		}
	case StateBooted:
		{
			if old == StateNew {
				c.state = &stateBooted{
					components: c.components,
					tools:      c.tools,

					ctx:  c.ctx,
					conf: c.conf,
				}
				return
			}
		}
	case StateServed:
		{
			if old == StateBooted {
				c.state = &stateServed{
					components: c.components,
					tools:      c.tools,

					ctx:  c.ctx,
					conf: c.conf,
				}
				return
			}
		}
	case StateOff:
		{
			if old == StateServed {
				c.state = &stateOff{
					components: c.components,
					tools:      c.tools,

					ctx:  c.ctx,
					conf: c.conf,
				}
				return
			}
		}
	}

	err = ErrInvalidSystemCoreState
	return
}
