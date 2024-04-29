package init_script

import (
	"sm-box/pkg/core"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
)

type script struct {
	conf *Config
	core core.Core

	components *components
}

// components - компоненты скрипта.
type components struct {
	Logger logger.Logger
}

// Run - запуск скрипта.
func (scr *script) Run() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	if err = scr.core.Serve(); err != nil {
		scr.components.Logger.Error().
			Format("An error occurred when starting maintenance of the '%s': '%s'. ",
				env.Vars.SystemName,
				err).Write()
	}

	return
}
