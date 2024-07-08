package config

import "sm-box/pkg/core/components/tracer"

// Middlewares - конфигурация промежуточного программного обеспечения http rest api.
type Middlewares struct {
	Cors *MiddlewareCors `json:"cors" yaml:"Cors" xml:"Cors"`
}

// FillEmptyFields - заполнение пустых полей конфигурации
func (conf *Middlewares) FillEmptyFields() *Middlewares {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	if conf.Cors == nil {
		conf.Cors = new(MiddlewareCors)
	}

	conf.Cors.FillEmptyFields()

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *Middlewares) Default() *Middlewares {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	conf.Cors = new(MiddlewareCors).Default()

	return conf
}

// Validate - валидация конфигурации.
func (conf *Middlewares) Validate() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	if err = conf.Cors.Validate(); err != nil {
		return
	}

	return
}
