package app

import (
	"context"
	"sm-box/pkg/core"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
)

// box - реализация приложения.
type box struct {
	conf *Config
	core core.Core

	components  *components
	controllers *controllers
	transport   *transport
}

// Serve - запуск приложения.
// Состояние приложения будет изменено на core.StateServed.
func (app *box) Serve() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	if err = app.core.Serve(); err != nil {
		app.Components().Logger().Error().
			Format("An error occurred when starting maintenance of the '%s': '%s'. ",
				env.Vars.SystemName,
				err).Write()
	}

	return
}

// Shutdown - завершение работы приложения.
// Остановить можно только приложения со статусом core.StateServed.
// Состояние ядра будет изменено на core.StateOff.
func (app *box) Shutdown() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	if err = app.core.Shutdown(); err != nil {
		app.Components().Logger().Error().
			Format("An error occurred when starting maintenance of the '%s': '%s'. ",
				env.Vars.SystemName,
				err).Write()
	}

	return
}

// State - получение состояния приложения.
//
// Возможные варианты состояния:
//  1. StateNew    - "New";
//  2. StateBooted - "Booted";
//  3. StateServed - "Served";
//  4. StateOff    - "Off";
func (app *box) State() (state core.State) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(state) }()
	}

	return app.core.State()
}

// Ctx - получение контекста приложения.
func (app *box) Ctx() (ctx context.Context) {
	return app.core.Ctx()
}

// Components - получение компонентов приложения.
func (app *box) Components() Components {
	return app.components
}

// Controllers - получение контроллеров приложения.
func (app *box) Controllers() Controllers {
	return app.controllers
}

// Transport - получение транспортной части приложения.
func (app *box) Transport() Transport {
	return app.transport
}
