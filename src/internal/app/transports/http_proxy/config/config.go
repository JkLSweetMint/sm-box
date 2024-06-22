package config

import (
	"sm-box/pkg/core/components/configurator"
	"sm-box/pkg/core/components/tracer"
)

// Config - конфигурация http proxy.
type Config struct {
	Engine  *Engine  `json:"engine"  yaml:"Engine"  xml:"Engine"`
	Postman *Postman `json:"postman" yaml:"Postman" xml:"Postman"`
}

// Read - чтение конфигурации.
func (conf *Config) Read() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var (
		c       configurator.Configurator[*Config]
		profile = configurator.PrivateProfile{
			Dir:      "/transports",
			Filename: "http_proxy.xml",
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
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	if conf.Engine == nil {
		conf.Engine = new(Engine)
	}

	if conf.Postman == nil {
		conf.Postman = new(Postman)
	}

	conf.Engine.FillEmptyFields()
	conf.Postman.FillEmptyFields()

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

	conf.Engine = new(Engine).Default()
	conf.Postman = new(Postman).Default()

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

	if err = conf.Engine.Validate(); err != nil {
		return
	}

	if err = conf.Postman.Validate(); err != nil {
		return
	}

	return
}
