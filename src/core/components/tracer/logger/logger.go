package logger

import (
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
func New(conf ...*Config) (log Logger, err error) {
	once.Do(func() {
		var c *Config

		if conf != nil {
			c = conf[0]
		}

		if instance, err = newLogger(c); err != nil {
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
