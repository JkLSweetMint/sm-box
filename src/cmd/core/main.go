package main

import (
	"fmt"
	"sm-box/src/core"
	"sm-box/src/core/components/tracer"
	"sm-box/src/core/env"
)

func init() {
	env.Vars.SystemName = "test"

	if err := tracer.Init(); err != nil {
		panic(err)
	}
}

func main() {
	var (
		cr  core.Core
		err error
	)

	// Создание и проверки экземпляра ядра
	{
		if cr, err = core.New(); err != nil {
			panic(err)
		} else if cr == nil {
			panic("Core instance is nil. ")
		}

		fmt.Printf("%+v\n", cr)
	}
}
