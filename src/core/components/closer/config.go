package closer

import (
	"os"
	"syscall"
)

var defaultSignals = []os.Signal{
	syscall.SIGINT,
	syscall.SIGTERM,
	syscall.SIGQUIT,
	syscall.SIGKILL,
}

// Config - конфигурация компонента ядра системы отвечающий за корректное завершение работы системы.
type Config struct {
	Signals []os.Signal `json:"signals" yaml:"Signals" xml:"Signals>Signal"`
}

// FillEmptyFields - заполнение обязательных пустых полей конфигурации
func (conf *Config) FillEmptyFields() *Config {
	if conf.Signals == nil {
		conf.Signals = make([]os.Signal, 0)
	}

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *Config) Default() *Config {
	conf.Signals = defaultSignals

	return conf
}

// Validate - валидация конфигурации.
func (conf *Config) Validate() (err error) {
	return
}
