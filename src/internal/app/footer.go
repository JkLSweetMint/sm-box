package app

import (
	"context"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
)

// serve - внутренний метод для запуска коробки.
func (bx *box) serve(ctx context.Context) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	bx.Components().Logger().Info().
		Format("Starting the '%s'... ", env.Vars.SystemName).Write()

	// Транспортная часть
	{
		go func() {
			if err = bx.Transports().RestApi().Listen(); err != nil {
				bx.Components().Logger().Error().
					Format("Failed to launch 'http rest api': '%s'. ", err).Write()
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
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	bx.Components().Logger().Info().
		Format("Shutting down the '%s'... ", env.Vars.SystemName).Write()

	// Транспортная часть
	{
		if err = bx.Transports().RestApi().Shutdown(); err != nil {
			bx.Components().Logger().Error().
				Format("Failed to stop 'http rest api': '%s'. ", err).Write()
		}
	}

	bx.Components().Logger().Info().
		Format("The '%s' has finished its work. ", env.Vars.SystemName).Write()

	return
}
