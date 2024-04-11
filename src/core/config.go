package core

import "sm-box/src/core/components/configurator"

var confProfile = configurator.PrivateProfile{
	Dir:      "/",
	Filename: "core.xml",
}

// Config - конфигурация ядра системы.
type Config struct{}

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
