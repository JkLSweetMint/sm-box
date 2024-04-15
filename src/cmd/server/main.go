package main

import (
	"fmt"
	"sm-box/src/core/components/tracer"
	"sm-box/src/core/env"
	env_mode "sm-box/src/core/env/mode"
	"sm-box/src/internal/server"
)

func init() {
	env.Vars.SystemName = "server"

	if env.Mode == env_mode.Dev {
		if err := tracer.Init(); err != nil {
			panic(fmt.Sprintf("An error occurred during initialization of the function/method call trace logging component: '%s'. ", err))
		}
	}
}

func main() {
	var (
		srv server.Server
		err error
	)

	if srv, err = server.New(); err != nil {
		panic(fmt.Sprintf("An error occurred during the creation of the '%s' instance: '%s'. ",
			env.Vars.SystemName,
			err))
	}

	if err = srv.Serve(); err != nil {
		panic(fmt.Sprintf("An error occurred when starting maintenance of the '%s': '%s'. ",
			env.Vars.SystemName,
			err))
	}

	if err = srv.Shutdown(); err != nil {
		panic(fmt.Sprintf("An error occurred while completing the maintenance of the '%s': '%s'. ",
			env.Vars.SystemName,
			err))
	}
}
