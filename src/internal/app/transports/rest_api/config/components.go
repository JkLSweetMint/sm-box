package config

import (
	"sm-box/internal/app/transports/rest_api/components/system_access"
	"sm-box/pkg/core/components/tracer"
)

// Components - конфигурация компонентов http rest api.
type Components struct {
	SystemAccess *system_access.Config `json:"system_access" yaml:"SystemAccess" xml:"SystemAccess"`
}

// FillEmptyFields - заполнение обязательных пустых полей конфигурации
func (conf *Components) FillEmptyFields() *Components {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	if conf.SystemAccess == nil {
		conf.SystemAccess = new(system_access.Config)
	}

	conf.SystemAccess.FillEmptyFields()

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *Components) Default() *Components {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	conf.SystemAccess = new(system_access.Config).Default()

	return conf
}

// Validate - валидация конфигурации.
func (conf *Components) Validate() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	if err = conf.SystemAccess.Validate(); err != nil {
		return
	}

	return
}
