package app

import (
	rest_api_conf "sm-box/src/internal/app/transports/rest_api/config"
	"sm-box/src/pkg/core/components/configurator"
)

var confProfile = configurator.PublicProfile{
	Dir:      "/",
	Filename: "box.yaml",
}

// Config - конфигурация коробки.
type Config struct {
	Transports *ConfigTransports `json:"transports" yaml:"Transports" xml:"Transports"`
}

// ConfigTransports - конфигурация транспортной части коробки
type ConfigTransports struct {
	RestAPI *rest_api_conf.Config `json:"rest_api" yaml:"RestAPI" xml:"RestAPI"`
}

// FillEmptyFields - заполнение обязательных пустых полей конфигурации
func (conf *Config) FillEmptyFields() *Config {
	if conf.Transports == nil {
		conf.Transports = new(ConfigTransports)
	}

	conf.Transports.FillEmptyFields()

	return conf
}

// FillEmptyFields - заполнение обязательных пустых полей конфигурации
func (conf *ConfigTransports) FillEmptyFields() *ConfigTransports {
	if conf.RestAPI == nil {
		conf.RestAPI = new(rest_api_conf.Config)
	}

	conf.RestAPI.FillEmptyFields()

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *Config) Default() *Config {
	conf.Transports = new(ConfigTransports).Default()

	return conf
}

// Default - ConfigTransports стандартной конфигурации.
func (conf *ConfigTransports) Default() *ConfigTransports {
	conf.RestAPI = new(rest_api_conf.Config).Default()

	return conf
}

// Validate - валидация конфигурации.
func (conf *Config) Validate() (err error) {
	if err = conf.Transports.Validate(); err != nil {
		return
	}

	return
}

// Validate - валидация конфигурации.
func (conf *ConfigTransports) Validate() (err error) {
	if err = conf.RestAPI.Validate(); err != nil {
		return
	}

	return
}
