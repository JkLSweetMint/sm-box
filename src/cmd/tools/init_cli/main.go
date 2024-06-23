package main

import (
	"fmt"
	"path"
	"sm-box/internal/tools/init_cli"
	"sm-box/pkg/core/components/configurator"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	env_mode "sm-box/pkg/core/env/mode"
)

func init() {
	env.Vars.SystemName = "init-cli"
	env.Version = "24.0.23"

	configurator.PbDir = path.Join(configurator.PbDir, "/tools", env.Vars.SystemName)
	configurator.PrtDir = path.Join(configurator.PrtDir, "/tools", env.Vars.SystemName)

	if env.Mode == env_mode.Dev {
		if err := tracer.Init(); err != nil {
			panic(fmt.Sprintf("An error occurred during initialization of the function/method call trace logging component: '%s'. ", err))
		}
	}
}

func main() {
	var (
		cli init_cli.CLI
		err error
	)

	if cli, err = init_cli.New(); err != nil {
		panic(fmt.Sprintf("An error occurred during the creation of the '%s' instance: '%s'. ",
			env.Vars.SystemName,
			err))
	}

	if err = cli.Exec(); err != nil {
		panic(fmt.Sprintf("An error occurred when starting maintenance of the '%s': '%s'. ",
			env.Vars.SystemName,
			err))
	}
}
