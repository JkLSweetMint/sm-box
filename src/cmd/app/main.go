package main

import (
	"fmt"
	"path"
	"sm-box/internal/app"
	"sm-box/pkg/core/components/configurator"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	env_mode "sm-box/pkg/core/env/mode"
)

func init() {
	env.Vars.SystemName = "box"
	env.Version = "24.0.37"

	configurator.PbDir = path.Join(configurator.PbDir, env.Vars.SystemName)
	configurator.PrtDir = path.Join(configurator.PrtDir, env.Vars.SystemName)

	if env.Mode == env_mode.Dev {
		if err := tracer.Init(); err != nil {
			panic(fmt.Sprintf("An error occurred during initialization of the function/method call trace logging component: '%s'. ", err))
		}
	}
}

func main() {
	var (
		bx  app.Box
		err error
	)

	if bx, err = app.New(); err != nil {
		panic(fmt.Sprintf("An error occurred during the creation of the '%s' instance: '%s'. ",
			env.Vars.SystemName,
			err))
	}

	if err = bx.Serve(); err != nil {
		panic(fmt.Sprintf("An error occurred when starting maintenance of the '%s': '%s'. ",
			env.Vars.SystemName,
			err))
	}

	if err = bx.Shutdown(); err != nil {
		panic(fmt.Sprintf("An error occurred while completing the maintenance of the '%s': '%s'. ",
			env.Vars.SystemName,
			err))
	}
}
