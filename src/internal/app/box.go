package app

import (
	"context"
	"sm-box/pkg/core"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
)

// box - реализация коробки.
type box struct {
	conf *Config
	core core.Core

	components  *components
	controllers *controllers
}

// Serve - запуск коробки.
// Состояние коробки будет изменено на core.StateServed.
func (bx *box) Serve() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	if err = bx.core.Serve(); err != nil {
		bx.Components().Logger().Error().
			Format("An error occurred when starting maintenance of the '%s': '%s'. ",
				env.Vars.SystemName,
				err).Write()
	}

	return
}

// Shutdown - завершение работы коробки.
// Остановить можно только коробки со статусом core.StateServed.
// Состояние ядра будет изменено на core.StateOff.
func (bx *box) Shutdown() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	if err = bx.core.Shutdown(); err != nil {
		bx.Components().Logger().Error().
			Format("An error occurred when starting maintenance of the '%s': '%s'. ",
				env.Vars.SystemName,
				err).Write()
	}

	return
}

// State - получение состояния коробки.
//
// Возможные варианты состояния:
//  1. StateNew    - "New";
//  2. StateBooted - "Booted";
//  3. StateServed - "Served";
//  4. StateOff    - "Off";
func (bx *box) State() (state core.State) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(state) }()
	}

	return bx.core.State()
}

// Ctx - получение контекста коробки.
func (bx *box) Ctx() (ctx context.Context) {
	return bx.core.Ctx()
}

// Components - получение компонентов коробки.
func (bx *box) Components() Components {
	return bx.components
}

// Controllers - получение контроллеров коробки.
func (bx *box) Controllers() Controllers {
	return bx.controllers
}
