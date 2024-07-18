package jwt_tokens_repository

import (
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/redis"
)

// Config - конфигурация.
type Config struct {
	Connector *redis.Config `json:"connector" yaml:"Connector" xml:"Connector"`
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
		conf.Connector = new(redis.Config)
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

	conf.Connector = new(redis.Config).Default()

	conf.Connector.Db = "0"
	conf.Connector.Auth.User = "root"
	conf.Connector.Auth.Password = "T4b4g9)53(W)l(SM"
	conf.Connector.Host = "authentication-redis"
	conf.Connector.Port = 6379

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
