package logger

import (
	"sm-box/pkg/core/components/tracer/logger/config"
	"sync"
)

var (
	once     = new(sync.Once)
	instance *logger
)

// Logger - описание компонента ведения журнала трессировки.
type Logger interface {
	Info() (msg Message)
	Error() (msg Message)

	Close() (err error)
}

// New - создание компонента ведения журнала.
func New(configs ...*config.Config) (log Logger, err error) {
	once.Do(func() {
		var conf *config.Config

		if configs != nil {
			conf = configs[0]

			if err = conf.FillEmptyFields().Validate(); err != nil {
				return
			}
		}

		if instance, err = newLogger(conf); err != nil {
			return
		}

		log = instance
	})

	if err != nil {
		return
	}

	if instance == nil {
		return nil, ErrInstanceNoFound
	}

	if log == nil {
		if log = instance.Copy(); log == nil {
			return nil, ErrComponentCouldNotBeCopied
		}
	}

	return
}
