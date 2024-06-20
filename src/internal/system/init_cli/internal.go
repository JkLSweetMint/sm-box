package init_cli

import (
	"context"
	"sm-box/pkg/core"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	c_errors "sm-box/pkg/errors"
)

// cli - реализация инструмента для инициализации системы.
type cli struct {
	conf *Config
	core core.Core

	controllers *controllers
	components  *components
}

// components - компоненты cli.
type components struct {
	Logger logger.Logger
}

// controllers - контроллеры коробки.
type controllers struct {
	Initialization interface {
		Initialize(ctx context.Context) (cErr c_errors.Error)
		Clear(ctx context.Context) (cErr c_errors.Error)
	}
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
