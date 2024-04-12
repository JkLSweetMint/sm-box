package core

import (
	"sm-box/src/core/components/closer"
	"sm-box/src/core/components/configurator"
)

var confProfile = configurator.PrivateProfile{
	Dir:      "/",
	Filename: "core.xml",
}

// Config - конфигурация ядра системы.
type Config struct {
	Closer *closer.Config `json:"closer" yaml:"Closer" xml:"Closer"`
}

// FillEmptyFields - заполнение обязательных пустых полей конфигурации
func (conf *Config) FillEmptyFields() *Config {
	if conf.Closer == nil {
		conf.Closer = new(closer.Config)
	}

	conf.Closer.FillEmptyFields()

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *Config) Default() *Config {
	conf.Closer = new(closer.Config).Default()

	return conf
}

// Validate - валидация конфигурации.
func (conf *Config) Validate() (err error) {
	return
}
