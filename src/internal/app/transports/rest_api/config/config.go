package config

import "sm-box/pkg/core/components/tracer"

// Config - конфигурация http rest api.
type Config struct {
	Engine      *Engine      `json:"engine"      yaml:"Engine"      xml:"Engine"`
	Components  *Components  `json:"components"  yaml:"Components"  xml:"Components"`
	Middlewares *Middlewares `json:"middlewares" yaml:"Middlewares" xml:"Middlewares"`
}

// FillEmptyFields - заполнение обязательных пустых полей конфигурации
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

	if conf.Middlewares == nil {
		conf.Middlewares = new(Middlewares)
	}

	if conf.Components == nil {
		conf.Components = new(Components)
	}

	conf.Engine.FillEmptyFields()
	conf.Middlewares.FillEmptyFields()
	conf.Components.FillEmptyFields()

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
	conf.Middlewares = new(Middlewares).Default()
	conf.Components = new(Components).Default()

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

	if err = conf.Middlewares.Validate(); err != nil {
		return
	}

	if err = conf.Components.Validate(); err != nil {
		return
	}

	return
}
