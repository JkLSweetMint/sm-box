package tracer

import (
	"sm-box/pkg/core/components/configurator"
	"sm-box/pkg/core/components/tracer/logger"
)

var confProfile = configurator.PrivateProfile{
	Dir:      "/components",
	Filename: "tracer.xml",
}

// Config - конфигурация компонента трессировки.
type Config struct {
	Levels []Level        `json:"levels" yaml:"Levels" xml:"Levels>Level"`
	Logger *logger.Config `json:"logger" yaml:"Logger" xml:"Logger"`
}

// FillEmptyFields - заполнение пустых полей конфигурации
func (conf *Config) FillEmptyFields() *Config {
	if conf.Levels == nil {
		conf.Levels = make([]Level, 0)
	}

	if conf.Logger == nil {
		conf.Logger = new(logger.Config)
	}

	conf.Logger.FillEmptyFields()

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *Config) Default() *Config {
	conf.Levels = allLevels

	conf.Logger = new(logger.Config).Default()

	return conf
}

// Validate - валидация конфигурации.
func (conf *Config) Validate() (err error) {
	return
}
