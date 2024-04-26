package pid

import (
	"errors"
	"fmt"
	"os"
	"path"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	"strconv"
)

// NewFile - создание нового PID файла.
func NewFile() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelCoreAddon)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var p = path.Join(env.Paths.SystemLocation, env.Paths.Var.Run, fmt.Sprintf("%s.pid", env.Vars.SystemName))

	// Проверка существования
	{
		var exist = true

		if _, err = os.Stat(p); errors.Is(err, os.ErrNotExist) {
			err = nil
			exist = false
		} else if err != nil {
			return
		}

		if exist {
			err = ErrPidFileAlreadyExist
			return
		}
	}

	// Создание файла
	{
		if err = os.WriteFile(p, []byte(strconv.Itoa(env.Vars.PID)), 0644); err != nil {
			return
		}
	}

	return
}

// RemoveFile - удаление pid файла.
func RemoveFile() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelCoreAddon)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var p = path.Join(env.Paths.SystemLocation, env.Paths.Var.Run, fmt.Sprintf("%s.pid", env.Vars.SystemName))

	// Проверка существования
	{
		if _, err = os.Stat(p); errors.Is(err, os.ErrNotExist) {
			err = ErrPidFileNotExist
			return
		} else if err != nil {
			return
		}
	}

	// Удаление файла
	{
		if err = os.Remove(p); err != nil {
			return
		}
	}

	return
}
