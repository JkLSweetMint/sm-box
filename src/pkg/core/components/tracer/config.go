package tracer

import (
	"sm-box/pkg/core/components/configurator"
	config2 "sm-box/pkg/core/components/tracer/logger/config"
)

// Config - конфигурация компонента трессировки.
type Config struct {
	Levels []Level         `json:"levels" yaml:"Levels" xml:"Levels>Level"`
	Logger *config2.Config `json:"logger" yaml:"Logger" xml:"Logger"`
}

// Read - чтение конфигурации.
func (conf *Config) Read() (err error) {
	var (
		c       configurator.Configurator[*Config]
		profile = configurator.PrivateProfile{
			Dir:      "/components",
			Filename: "tracer.xml",
		}
	)

	if c, err = configurator.New[*Config](conf); err != nil {
		return
	} else if err = c.Private().Profile(profile).Init(); err != nil {
		return
	}

	if err = conf.FillEmptyFields().Validate(); err != nil {
		return
	}

	return
}

// FillEmptyFields - заполнение пустых полей конфигурации
func (conf *Config) FillEmptyFields() *Config {
	if conf.Levels == nil {
		conf.Levels = make([]Level, 0)
	}

	if conf.Logger == nil {
		conf.Logger = new(config2.Config)
	}

	conf.Logger.FillEmptyFields()

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *Config) Default() *Config {
	conf.Levels = allLevels

	conf.Logger = new(config2.Config).Default()

	return conf
}

// Validate - валидация конфигурации.
func (conf *Config) Validate() (err error) {
	return
}
