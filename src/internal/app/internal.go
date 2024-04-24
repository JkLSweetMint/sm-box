package app

import (
	"context"
	"sm-box/src/pkg/core"
	"sm-box/src/pkg/core/components/tracer"
	"sm-box/src/pkg/core/env"
)

// box - реализация коробки.
type box struct {
	conf *Config
	core core.Core

	components  *components
	controllers *controllers
	transports  *transports
}

// Serve - запуск коробки.
// Состояние коробки будет изменено на core.StateServed.
func (bx *box) Serve() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain)

		trc.FunctionCall()
		trc.Error(err).FunctionCallFinished()
	}

	if err = bx.core.Serve(); err != nil {
		bx.Components().Logger().Error().
			Format("An error occurred when starting maintenance of the '%s': '%s'. ",
				env.Vars.SystemName,
				err)
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
		trc.Error(err).FunctionCallFinished()
	}

	if err = bx.core.Shutdown(); err != nil {
		bx.Components().Logger().Error().
			Format("An error occurred when starting maintenance of the '%s': '%s'. ",
				env.Vars.SystemName,
				err)
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

// Transports - получение транспортной части коробки.
func (bx *box) Transports() Transports {
	return bx.transports
}

// serve - внутренний метод для запуска коробки.
func (bx *box) serve(ctx context.Context) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal)

		trc.FunctionCall(ctx)
		trc.Error(err).FunctionCallFinished()
	}

	bx.Components().Logger().Info().
		Format("Starting the '%s'... ", env.Vars.SystemName).Write()

	// Транспортная часть
	{
		go func() {
			if err = bx.Transports().RestApi().Listen(); err != nil {
				bx.Components().Logger().Error().
					Format("Failed to launch 'http rest api': '%s'. ", err)
			}
		}()
	}

	bx.Components().Logger().Info().
		Format("The '%s' has been successfully started. ", env.Vars.SystemName).Write()

	return
}

// shutdown - внутренний метод для завершения работы коробки.
func (bx *box) shutdown(ctx context.Context) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal)

		trc.FunctionCall(ctx)
		trc.Error(err).FunctionCallFinished()
	}

	bx.Components().Logger().Info().
		Format("Shutting down the '%s'... ", env.Vars.SystemName).Write()

	// Транспортная часть
	{
		if err = bx.Transports().RestApi().Shutdown(); err != nil {
			bx.Components().Logger().Error().
				Format("Failed to stop 'http rest api': '%s'. ", err)
		}
	}

	bx.Components().Logger().Info().
		Format("The '%s' has finished its work. ", env.Vars.SystemName).Write()

	return
}
