package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"runtime"
	env_files "sm-box/pkg/core/env/files"
	env_mode "sm-box/pkg/core/env/mode"
	env_paths "sm-box/pkg/core/env/paths"
	env_synchronization "sm-box/pkg/core/env/synchronization"
	env_vars "sm-box/pkg/core/env/vars"
	"testing"
)

var (
	// Mode - текущий режим работы системы.
	Mode env_mode.Mode = env_mode.Dev
	// Version - текущая версия системы.
	Version = ""
	// OS - версия операционной системы.
	OS = fmt.Sprintf("%s - %s", runtime.GOOS, runtime.GOARCH)
)

var (
	// Vars - переменные окружения системы.
	Vars = new(env_vars.Dev).Build()

	// Paths - хранилища путей системы.
	Paths = new(env_paths.Dev).Build(env_paths.BuildOptions{ID: Vars.ID})

	// Files - хранилища файлов системы.
	Files = new(env_files.Dev).Build()

	// Synchronization - инструменты синхронизации системы.
	Synchronization = new(env_synchronization.Dev).Build()
)

// init - инициализация окружения системы.
func init() {
	var err error

	testing.Init()
	godotenv.Load()

	if Paths.SystemLocation, err = getSystemLocation(); err != nil {
		panic(err)
	}

	if err = initSystemDir(); err != nil {
		panic(err)
	}

	if err = readEncryptionKeys(); err != nil {
		panic(err)
	}
}
