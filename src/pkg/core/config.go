package core

import (
	"sm-box/pkg/core/components/configurator"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/tools/closer"
)

var confProfile = configurator.PrivateProfile{
	Dir:      "/",
	Filename: "core.xml",
}

// Config - конфигурация ядра системы.
type Config struct {
	Tools *ConfigTools `json:"tools" yaml:"Tools" xml:"Tools"`
}

// ConfigTools - конфигурация инструментов ядра системы.
type ConfigTools struct {
	Closer *closer.Config `json:"closer" yaml:"Closer" xml:"Closer"`
}

// FillEmptyFields - заполнение обязательных пустых полей конфигурации
func (conf *Config) FillEmptyFields() *Config {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	if conf.Tools == nil {
		conf.Tools = new(ConfigTools)
	}

	conf.Tools.FillEmptyFields()

	return conf
}

// FillEmptyFields - заполнение обязательных пустых полей конфигурации
func (conf *ConfigTools) FillEmptyFields() *ConfigTools {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	if conf.Closer == nil {
		conf.Closer = new(closer.Config)
	}

	conf.Closer.FillEmptyFields()

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

	conf.Tools = new(ConfigTools).Default()

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *ConfigTools) Default() *ConfigTools {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	conf.Closer = new(closer.Config).Default()

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

// Validate - валидация конфигурации.
func (conf *ConfigTools) Validate() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	return
}
