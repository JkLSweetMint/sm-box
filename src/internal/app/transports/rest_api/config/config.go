package config

import (
	"path"
	"sm-box/pkg/core/components/configurator"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
)

// Config - конфигурация http rest api.
type Config struct {
	Engine      *Engine      `json:"engine"      yaml:"Engine"      xml:"Engine"`
	Components  *Components  `json:"components"  yaml:"Components"  xml:"Components"`
	Middlewares *Middlewares `json:"middlewares" yaml:"Middlewares" xml:"Middlewares"`
	Postman     *Postman     `json:"postman"     yaml:"Postman"     xml:"Postman"`
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
			Dir:      path.Join(env.Vars.SystemName, "/transports"),
			Filename: "rest_api.xml",
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

	if conf.Components == nil {
		conf.Components = new(Components)
	}

	if conf.Middlewares == nil {
		conf.Middlewares = new(Middlewares)
	}

	if conf.Postman == nil {
		conf.Postman = new(Postman)
	}

	conf.Engine.FillEmptyFields()
	conf.Components.FillEmptyFields()
	conf.Middlewares.FillEmptyFields()
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
	conf.Components = new(Components).Default()
	conf.Middlewares = new(Middlewares).Default()
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

	if err = conf.Components.Validate(); err != nil {
		return
	}

	if err = conf.Middlewares.Validate(); err != nil {
		return
	}

	if err = conf.Postman.Validate(); err != nil {
		return
	}

	return
}
