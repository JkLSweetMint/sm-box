package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"runtime"
	env_mode "sm-box/src/core/env/mode"
	env_paths "sm-box/src/core/env/paths"
	env_synchronization "sm-box/src/core/env/synchronization"
	env_vars "sm-box/src/core/env/vars"
	"testing"
)

var (
	// Mode - текущий режим работы системы.
	Mode env_mode.Mode = env_mode.Dev
	// Version - текущая версия системы.
	Version = "24.0.6"
	// OS - версия операционной системы.
	OS = fmt.Sprintf("%s - %s", runtime.GOOS, runtime.GOARCH)
)

var (
	// Vars - переменные окружения системы.
	Vars = new(env_vars.Dev).Build()

	// Paths - хранилища путей системы.
	Paths = new(env_paths.Dev).Build(env_paths.BuildOptions{ID: Vars.ID})

	// Synchronization - инструменты синхронизации системы.
	Synchronization = new(env_synchronization.Dev).Build()
)

const (
	testSystemLocation = "F:/projects/SweetMint/sm-box"
)

// init - инициализация окружения системы.
func init() {
	var err error

	testing.Init()
	godotenv.Load()

	if Paths.SystemLocation, err = getSystemLocation(testSystemLocation); err != nil {
		panic(err)
	}

	if err = initSystemDir(); err != nil {
		panic(err)
	}
}
