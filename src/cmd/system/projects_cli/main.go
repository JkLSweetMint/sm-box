package main

import (
	"fmt"
	"sm-box/internal/system/projects_cli"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	env_mode "sm-box/pkg/core/env/mode"
)

func init() {
	env.Vars.SystemName = "project-cli"
	env.Version = "24.0.21"

	if env.Mode == env_mode.Dev {
		if err := tracer.Init(); err != nil {
			panic(fmt.Sprintf("An error occurred during initialization of the function/method call trace logging component: '%s'. ", err))
		}
	}
}

func main() {
	var (
		cli projects_cli.CLI
		err error
	)

	if cli, err = projects_cli.New(); err != nil {
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