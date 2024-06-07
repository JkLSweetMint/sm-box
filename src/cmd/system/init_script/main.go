package main

import (
	"fmt"
	script "sm-box/internal/system/init_script"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	env_mode "sm-box/pkg/core/env/mode"
)

func init() {
	env.Vars.SystemName = "init-script"
	env.Version = "24.0.16"

	if env.Mode == env_mode.Dev {
		if err := tracer.Init(); err != nil {
			panic(fmt.Sprintf("An error occurred during initialization of the function/method call trace logging component: '%s'. ", err))
		}
	}
}

func main() {
	var (
		src script.Script
		err error
	)

	if src, err = script.New(); err != nil {
		panic(fmt.Sprintf("An error occurred during the creation of the '%s' instance: '%s'. ",
			env.Vars.SystemName,
			err))
	}

	if err = src.Run(); err != nil {
		panic(fmt.Sprintf("An error occurred when starting maintenance of the '%s': '%s'. ",
			env.Vars.SystemName,
			err))
	}
}
