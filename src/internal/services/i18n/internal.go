package service

import (
	"context"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
)

// serve - внутренний метод для запуска сервиса.
func (srv *service) serve(ctx context.Context) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	srv.Components().Logger().Info().
		Format("Starting the '%s'... ", env.Vars.SystemName).Write()

	// Транспортная часть
	{
		env.Synchronization.WaitGroup.Add(1)

		go func() {
			defer env.Synchronization.WaitGroup.Done()

			if err = srv.Transport().Servers().HttpRestApi().Listen(); err != nil {
				srv.Components().Logger().Error().
					Format("Failed to launch 'http rest api server': '%s'. ", err).Write()
			}
		}()
	}

	srv.Components().Logger().Info().
		Format("The '%s' has been successfully started. ", env.Vars.SystemName).Write()

	return
}

// shutdown - внутренний метод для завершения работы сервиса.
func (srv *service) shutdown(ctx context.Context) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	srv.Components().Logger().Info().
		Format("Shutting down the '%s'... ", env.Vars.SystemName).Write()

	// Транспортная часть
	{
		if err = srv.Transport().Servers().HttpRestApi().Shutdown(); err != nil {
			srv.Components().Logger().Error().
				Format("Failed to stop 'http rest api server': '%s'. ", err).Write()
		}
	}

	srv.Components().Logger().Info().
		Format("The '%s' has finished its work. ", env.Vars.SystemName).Write()

	return
}
