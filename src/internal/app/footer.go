package app

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"path"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
)

// init - инициализация коробки.
func (bx *box) init() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	// Проверки что уже инициализировано
	{
		if _, err = os.Stat(path.Join(env.Paths.SystemLocation, env.Paths.Var.Lib, env.Files.Var.Lib.SystemDB)); errors.Is(err, os.ErrNotExist) {
			err = nil
		} else {
			if err == nil {
				bx.Components().Logger().Info().
					Format("'%s' is initialized, no reinitialization is required. ", env.Vars.SystemName).Write()
			} else {
				bx.Components().Logger().Error().
					Format("Failed to initialize '%s': '%s'. ", env.Vars.SystemName, err).Write()
			}

			return
		}
	}

	bx.Components().Logger().Info().
		Format("Starting initialization '%s'... ", env.Vars.SystemName).Write()

	// Запуск скрипта для инициализации
	{
		var cmd = exec.Command(path.Join("./", env.Paths.SystemBin, "init.exe"))
		cmd.Dir = env.Paths.SystemLocation

		if err = cmd.Start(); err != nil {
			bx.Components().Logger().Error().
				Format("Failed to call the script to initialize '%s': '%s'. ",
					env.Vars.SystemName,
					err).Write()
			return
		}

		if err = cmd.Wait(); err != nil {
			if exiterr, ok := err.(*exec.ExitError); ok {
				if code := exiterr.ExitCode(); code != 0 {
					bx.Components().Logger().Error().
						Format("The initialization script '%s' failed with the status code %d. ",
							env.Vars.SystemName,
							code).Write()
					return
				}
			} else {
				bx.Components().Logger().Error().
					Format("Failed to call the script to initialize '%s': '%s'. ",
						env.Vars.SystemName,
						err).Write()
				return
			}
		}
	}

	bx.Components().Logger().Info().
		Format("The initialization '%s' has been successfully finished. ", env.Vars.SystemName).Write()

	return
}

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
