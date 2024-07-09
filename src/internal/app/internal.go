package app

import (
	"context"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
)

// serve - внутренний метод для запуска приложения.
func (app *box) serve(ctx context.Context) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	app.Components().Logger().Info().
		Format("Starting the '%s'... ", env.Vars.SystemName).Write()

	// Транспортная часть
	{
		env.Synchronization.WaitGroup.Add(2)

		go func() {
			defer env.Synchronization.WaitGroup.Done()

			if err = app.Transport().Servers().Http().RestApi().Listen(); err != nil {
				app.Components().Logger().Error().
					Format("Failed to launch 'http rest api server': '%s'. ", err).Write()
			}
		}()

		go func() {
			defer env.Synchronization.WaitGroup.Done()

			if err = app.Transport().Servers().Grpc().ProjectsService().Listen(); err != nil {
				app.Components().Logger().Error().
					Format("Failed to launch 'grpc server for projects service': '%s'. ", err).Write()
			}
		}()
	}

	app.Components().Logger().Info().
		Format("The '%s' has been successfully started. ", env.Vars.SystemName).Write()

	return
}

// shutdown - внутренний метод для завершения работы приложения.
func (app *box) shutdown(ctx context.Context) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	app.Components().Logger().Info().
		Format("Shutting down the '%s'... ", env.Vars.SystemName).Write()

	// Транспортная часть
	{
		if err = app.Transport().Servers().Http().RestApi().Shutdown(); err != nil {
			app.Components().Logger().Error().
				Format("Failed to stop 'http rest api server': '%s'. ", err).Write()
		}

		if err = app.Transport().Servers().Grpc().ProjectsService().Shutdown(); err != nil {
			app.Components().Logger().Error().
				Format("grpc server for projects service': '%s'. ", err).Write()
		}
	}

	app.Components().Logger().Info().
		Format("The '%s' has finished its work. ", env.Vars.SystemName).Write()

	return
}
