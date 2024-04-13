package closer

import (
	"syscall"
)

var defaultSignals = []syscall.Signal{
	syscall.SIGINT,
	syscall.SIGTERM,
	syscall.SIGQUIT,
	syscall.SIGKILL,
}

// Config - конфигурация инструмента ядра системы отвечающий за корректное завершение работы системы.
type Config struct {
	Signals []syscall.Signal `json:"signals" yaml:"Signals" xml:"Signals>Signal"`
}

// FillEmptyFields - заполнение обязательных пустых полей конфигурации
func (conf *Config) FillEmptyFields() *Config {
	if conf.Signals == nil {
		conf.Signals = make([]syscall.Signal, 0)
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
