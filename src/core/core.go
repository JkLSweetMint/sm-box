package core

import (
	"context"
	"sm-box/src/core/components/logger"
	"sync"
)

var (
	once     = new(sync.Once)
	instance Core
)

// Core - описание ядра системы.
type Core interface {
	Shutdown() (err error)
	Boot() (err error)
	Serve() (err error)

	State() (state State)
	Ctx() (ctx context.Context)

	Components() interface {
		Logger() logger.Logger
	}
}

// New - создание ядра системы.
// Может быть создан только один объект ядра!
//
// Ядро может быть в следующих состояний:
//   - StateNew    - "New"
//   - StateBooted - "Booted"
//   - StateServed - "Served"
//   - StateOff    - "Off"
func New() (c Core, err error) {
	once.Do(func() {
		var instance_ = &core{
			ctx: context.Background(),
		}

		if instance_.conf, err = BuildConfig(); err != nil {
			return
		}

		instance = instance_
	})

	if err != nil {
		return
	}

	c = instance

	return
}
