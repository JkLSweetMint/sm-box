package server

import (
	"context"
	"sm-box/src/core"
	"sm-box/src/core/components/tracer"
	"sm-box/src/core/env"
)

// server - сервер.
type server struct {
	conf *Config
	core core.Core

	components  *components
	controllers *controllers
}

// Serve - запуск сервера.
// Состояние сервера будет изменено на core.StateServed.
func (s *server) Serve() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain)

		trc.FunctionCall()
		trc.Error(err).FunctionCallFinished()
	}

	if err = s.core.Serve(); err != nil {
		s.Components().Logger().Error().
			Format("An error occurred when starting maintenance of the '%s': '%s'. ",
				env.Vars.SystemName,
				err)
	}

	return
}

// Shutdown - завершение работы сервера.
// Остановить можно только сервер со статусом core.StateServed.
// Состояние ядра будет изменено на core.StateOff.
func (s *server) Shutdown() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain)

		trc.FunctionCall()
		trc.Error(err).FunctionCallFinished()
	}

	if err = s.core.Shutdown(); err != nil {
		s.Components().Logger().Error().
			Format("An error occurred when starting maintenance of the '%s': '%s'. ",
				env.Vars.SystemName,
				err)
	}

	return
}

// State - получение состояния сервера.
//
// Возможные варианты состояния:
//  1. StateNew    - "New";
//  2. StateBooted - "Booted";
//  3. StateServed - "Served";
//  4. StateOff    - "Off";
func (s *server) State() (state core.State) {
	return s.core.State()
}

// Ctx - получение контекста сервера.
func (s *server) Ctx() (ctx context.Context) {
	return s.core.Ctx()
}

// Components - получение компонентов сервера.
func (s *server) Components() Components {
	return s.components
}

// Controllers - получение контроллеров сервера.
func (s *server) Controllers() Controllers {
	return s.controllers
}

// serve - внутренний метод для запуска сервера.
func (s *server) serve(ctx context.Context) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal)

		trc.FunctionCall(ctx)
		trc.Error(err).FunctionCallFinished()
	}

	s.Components().Logger().Info().
		Format("Starting the '%s'... ", env.Vars.SystemName).Write()

	s.Components().Logger().Info().
		Format("The '%s' has been successfully started. ", env.Vars.SystemName).Write()

	return
}

// shutdown - внутренний метод для завершения работы сервера.
func (s *server) shutdown(ctx context.Context) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal)

		trc.FunctionCall(ctx)
		trc.Error(err).FunctionCallFinished()
	}

	s.Components().Logger().Info().
		Format("Shutting down the '%s'... ", env.Vars.SystemName).Write()

	s.Components().Logger().Info().
		Format("The '%s' has finished its work. ", env.Vars.SystemName).Write()

	return
}
