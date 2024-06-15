package project_cli

import (
	"sm-box/pkg/core"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
)

// cli - реализация скрипта для управления проектами.
type cli struct {
	conf *Config
	core core.Core

	components *components
}

// components - компоненты скрипта.
type components struct {
	Logger logger.Logger
}

// Exec - выполнить скрипт.
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
