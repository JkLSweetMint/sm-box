package repository

import (
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/postgresql"
)

// Config - конфигурация репозитория.
type Config struct {
	Connector *postgresql.Config `json:"connector" yaml:"Connector" xml:"Connector"`
}

// FillEmptyFields - заполнение пустых полей конфигурации
func (conf *Config) FillEmptyFields() *Config {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	if conf.Connector == nil {
		conf.Connector = new(postgresql.Config)
	}

	conf.Connector.FillEmptyFields()

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

	conf.Connector = new(postgresql.Config).Default()

	conf.Connector.DbName = "box"
	conf.Connector.Auth.User = "root"
	conf.Connector.Auth.Password = "3bxMue16ztXPR635"
	conf.Connector.Host = "postgres"
	conf.Connector.Port = 5432

	conf.Connector.Tags["sslmode"] = "disable"

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

	if err = conf.Connector.Validate(); err != nil {
		return
	}

	return
}