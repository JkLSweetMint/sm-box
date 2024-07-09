package config

import (
	"sm-box/internal/services/authentication/transport/servers/http/rest_api/components/access_system"
	"sm-box/pkg/core/components/tracer"
)

// Components - конфигурация компонентов http rest api.
type Components struct {
	AccessSystem *access_system.Config `json:"access_system" yaml:"AccessSystem" xml:"AccessSystem"`
}

// FillEmptyFields - заполнение пустых полей конфигурации
func (conf *Components) FillEmptyFields() *Components {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	if conf.AccessSystem == nil {
		conf.AccessSystem = new(access_system.Config)
	}

	conf.AccessSystem.FillEmptyFields()

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

	conf.AccessSystem = new(access_system.Config).Default()

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

	if err = conf.AccessSystem.Validate(); err != nil {
		return
	}

	return
}
