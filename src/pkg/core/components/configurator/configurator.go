package configurator

import "sm-box/pkg/core/env"

var ConfigDir = env.Vars.SystemName

// Configurator - диспетчер конфигураций.
type Configurator[T any] interface {
	Public() Public[T]
	Private() Private[T]
}

// Public - диспетчер публичных конфигураций.
// Для хранения файлов конфигураций используются директория /etc.
type Public[T any] interface {
	Encoder(encoder Encoder) Public[T]
	File(dir, filename string) Public[T]
	Profile(profile PublicProfile) Public[T]

	Init() (err error)
	Write() (err error)
	Read() (err error)
}

// Private - диспетчер приватных конфигураций.
// Для хранения файлов конфигураций используются директория /system.
type Private[T any] interface {
	Encoder(encoder Encoder) Private[T]
	File(dir, filename string) Private[T]
	Profile(profile PrivateProfile) Private[T]

	Init() (err error)
}

// New - создание диспетчера конфигураций.
func New[T any](conf Config[T]) (c Configurator[T], err error) {
	if conf == nil {
		return nil, ErrNilConfigurationInstanceWasPassed
	}

	c = &configurator[T]{
		conf: conf,
	}

	return
}
