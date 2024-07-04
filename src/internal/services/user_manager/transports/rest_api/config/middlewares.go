package config

import "sm-box/pkg/core/components/tracer"

// Middlewares - конфигурация промежуточного программного обеспечения http rest api.
type Middlewares struct {
	Compress *MiddlewareCompress `json:"compress" yaml:"Compress" xml:"Compress"`
	Cache    *MiddlewareCache    `json:"cache"    yaml:"Cache"    xml:"Cache"`
	Cors     *MiddlewareCors     `json:"cors"     yaml:"Cors"     xml:"Cors"`
}

// FillEmptyFields - заполнение пустых полей конфигурации
func (conf *Middlewares) FillEmptyFields() *Middlewares {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	if conf.Compress == nil {
		conf.Compress = new(MiddlewareCompress)
	}

	if conf.Cache == nil {
		conf.Cache = new(MiddlewareCache)
	}

	if conf.Cors == nil {
		conf.Cors = new(MiddlewareCors)
	}

	conf.Compress.FillEmptyFields()
	conf.Cache.FillEmptyFields()
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

	conf.Compress = new(MiddlewareCompress).Default()
	conf.Cache = new(MiddlewareCache).Default()
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

	if err = conf.Compress.Validate(); err != nil {
		return
	}

	if err = conf.Cache.Validate(); err != nil {
		return
	}

	if err = conf.Cors.Validate(); err != nil {
		return
	}

	return
}