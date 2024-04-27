package access_system

import (
	"sm-box/internal/app/transports/rest_api/components/access_system/repository"
	"sm-box/pkg/core/components/tracer"
)

// Config - конфигурация компонента системы доступа.
type Config struct {
	Repository *repository.Config `json:"repository" yaml:"Repository" xml:"Repository"`
}

// FillEmptyFields - заполнение пустых полей конфигурации
func (conf *Config) FillEmptyFields() *Config {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	if conf.Repository == nil {
		conf.Repository = new(repository.Config)
	}

	conf.Repository.FillEmptyFields()

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

	conf.Repository = new(repository.Config).Default()

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

	if err = conf.Repository.Validate(); err != nil {
		return
	}

	return
}
