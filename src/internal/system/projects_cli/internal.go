package projects_cli

import (
	"context"
	"sm-box/pkg/core"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	c_errors "sm-box/pkg/errors"
)

// cli - реализация инструмента для управления проектами.
type cli struct {
	conf *Config
	core core.Core

	controllers *controllers
	components  *components
}

// controllers - контроллеры cli.
type controllers struct {
	Projects interface {
		Create(ctx context.Context, title, version, description string) (cErr c_errors.Error)
		Remove(ctx context.Context, id string) (cErr c_errors.Error)
	}
}

// components - компоненты cli.
type components struct {
	Logger logger.Logger
}

// Exec - выполнить cli.
func (cli_ *cli) Exec() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	if err = cli_.core.Serve(); err != nil {
		cli_.components.Logger.Error().
			Format("An error occurred when starting maintenance of the '%s': '%s'. ",
				env.Vars.SystemName,
				err).Write()
	}

	return
}
