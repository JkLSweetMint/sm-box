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
		env.Synchronization.WaitGroup.Add(4)

		go func() {
			defer env.Synchronization.WaitGroup.Done()

			if err = srv.Transport().Servers().Http().RestApi().Listen(); err != nil {
				srv.Components().Logger().Error().
					Format("Failed to launch 'http rest api server': '%s'. ", err).Write()
			}
		}()

		go func() {
			defer env.Synchronization.WaitGroup.Done()

			if err = srv.Transport().Servers().Grpc().BasicAuthenticationService().Listen(); err != nil {
				srv.Components().Logger().Error().
					Format("Failed to launch 'grpc server for authentication service': '%s'. ", err).Write()
			}
		}()

		go func() {
			defer env.Synchronization.WaitGroup.Done()

			if err = srv.Transport().Servers().Grpc().UsersService().Listen(); err != nil {
				srv.Components().Logger().Error().
					Format("Failed to launch 'grpc server for users service': '%s'. ", err).Write()
			}
		}()

		go func() {
			defer env.Synchronization.WaitGroup.Done()

			if err = srv.Transport().Servers().Grpc().AccessSystemService().Listen(); err != nil {
				srv.Components().Logger().Error().
					Format("Failed to launch 'grpc server for access system service': '%s'. ", err).Write()
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
		if err = srv.Transport().Servers().Http().RestApi().Shutdown(); err != nil {
			srv.Components().Logger().Error().
				Format("Failed to stop 'http rest api server': '%s'. ", err).Write()
		}

		if err = srv.Transport().Servers().Grpc().BasicAuthenticationService().Shutdown(); err != nil {
			srv.Components().Logger().Error().
				Format("Failed to stop 'grpc server for authentication service': '%s'. ", err).Write()
		}

		if err = srv.Transport().Servers().Grpc().UsersService().Shutdown(); err != nil {
			srv.Components().Logger().Error().
				Format("Failed to stop 'grpc server for users service': '%s'. ", err).Write()
		}

		if err = srv.Transport().Servers().Grpc().AccessSystemService().Shutdown(); err != nil {
			srv.Components().Logger().Error().
				Format("Failed to stop 'grpc server for access system service': '%s'. ", err).Write()
		}
	}

	srv.Components().Logger().Info().
		Format("The '%s' has finished its work. ", env.Vars.SystemName).Write()

	return
}
