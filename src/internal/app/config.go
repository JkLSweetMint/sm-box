package app

import "sm-box/src/pkg/core/components/configurator"

var confProfile = configurator.PrivateProfile{
	Dir:      "/",
	Filename: "box.xml",
}

// Config - конфигурация коробки.
type Config struct {
}

// FillEmptyFields - заполнение обязательных пустых полей конфигурации
func (conf *Config) FillEmptyFields() *Config {
	return conf
}

// Default - запись стандартной конфигурации.
func (conf *Config) Default() *Config {
	return conf
}

// Validate - валидация конфигурации.
func (conf *Config) Validate() (err error) {
	return
}
