package logger

import (
	"sync"
)

var (
	once     = new(sync.Once)
	instance *logger
)

// Logger - описание компонента ведения журнала системы.
type Logger interface {
	Debug() (msg Message)
	Info() (msg Message)
	Warn() (msg Message)
	Error() (msg Message)
	Panic() (msg Message)
	DPanic() (msg Message)
	Fatal() (msg Message)

	Close() (err error)
}

// New - создание компонента ведения журнала.
func New(initiator string, conf ...*Config) (log Logger, err error) {
	once.Do(func() {
		var c *Config

		if conf != nil {
			c = conf[0]
		}

		if instance, err = newLogger(initiator, c); err != nil {
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
		if log = instance.Copy(initiator); log == nil {
			return nil, ErrComponentCouldNotBeCopied
		}
	}

	return
}
