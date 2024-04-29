package init_script

import (
	"sm-box/pkg/core/components/configurator"
	"sm-box/pkg/core/components/tracer"
)

var confProfile = configurator.PrivateProfile{
	Dir:      "/system/",
	Filename: "init_script.xml",
}

// Config - конфигурация скрипта инициализации.
type Config struct {
}

// FillEmptyFields - заполнение пустых полей конфигурации
func (conf *Config) FillEmptyFields() *Config {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *Config) Default() *Config {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	return conf
}

// Validate - валидация конфигурации.
func (conf *Config) Validate() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	return
}
