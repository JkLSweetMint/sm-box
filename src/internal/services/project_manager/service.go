package service

import (
	"context"
	"sm-box/pkg/core"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
)

// service - реализация сервиса.
type service struct {
	conf *Config
	core core.Core

	components  *components
	controllers *controllers
	transports  *transports
}

// Serve - запуск сервиса.
// Состояние сервиса будет изменено на core.StateServed.
func (srv *service) Serve() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	if err = srv.core.Serve(); err != nil {
		srv.Components().Logger().Error().
			Format("An error occurred when starting maintenance of the '%s': '%s'. ",
				env.Vars.SystemName,
				err).Write()
	}

	return
}

// Shutdown - завершение работы сервиса.
// Остановить можно только сервиса со статусом core.StateServed.
// Состояние ядра будет изменено на core.StateOff.
func (srv *service) Shutdown() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	if err = srv.core.Shutdown(); err != nil {
		srv.Components().Logger().Error().
			Format("An error occurred when starting maintenance of the '%s': '%s'. ",
				env.Vars.SystemName,
				err).Write()
	}

	return
}

// State - получение состояния сервиса.
//
// Возможные варианты состояния:
//  1. StateNew    - "New";
//  2. StateBooted - "Booted";
//  3. StateServed - "Served";
//  4. StateOff    - "Off";
func (srv *service) State() (state core.State) {
	return srv.core.State()
}

// Ctx - получение контекста сервиса.
func (srv *service) Ctx() (ctx context.Context) {
	return srv.core.Ctx()
}

// Components - получение компонентов сервиса.
func (srv *service) Components() Components {
	return srv.components
}

// Controllers - получение контроллеров сервиса.
func (srv *service) Controllers() Controllers {
	return srv.controllers
}

// Transports - получение транспортной части сервиса.
func (srv *service) Transports() Transports {
	return srv.transports
}
