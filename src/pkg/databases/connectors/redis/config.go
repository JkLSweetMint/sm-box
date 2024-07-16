package redis

import (
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/tools/format"
	"strings"
)

// templateConnectionString - шаблон строки для подключения к базе.
const templateConnectionString = "redis://{{user}}:{{password}}@{{host}}:{{port}}/{{db}}"

// Config - конфигурация коннектора к базе данных.
type Config struct {
	// Db - название базы данных.
	Db string `json:"db" yaml:"Db" xml:"db,attr"`

	// Host - адрес базы данных.
	Host string `json:"host" yaml:"Host" xml:"host,attr"`

	// Port - порт базы данных.
	Port uint `json:"port" yaml:"Port" xml:"port,attr"`

	// Auth - конфигурация авторизации в базе данных.
	Auth *ConfigAuth `json:"auth" yaml:"Auth" xml:"Auth"`
}

// ConfigAuth - конфигурация авторизации в базе данных.
type ConfigAuth struct {
	// User - имя пользователя.
	User string `json:"user" yaml:"User" xml:"user,attr"`

	// Password - пароль пользователя.
	Password string `json:"password" yaml:"Password" xml:"password,attr"`
}

// FillEmptyFields - заполнение пустых полей конфигурации
func (conf *Config) FillEmptyFields() *Config {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	if conf.Auth == nil {
		conf.Auth = new(ConfigAuth)
	}

	conf.Auth.FillEmptyFields()

	return conf
}

// FillEmptyFields - заполнение пустых полей конфигурации
func (conf *ConfigAuth) FillEmptyFields() *ConfigAuth {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

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

	conf.Db = "0"
	conf.Host = "127.0.0.1"
	conf.Port = 6379

	conf.Auth = new(ConfigAuth).Default()

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *ConfigAuth) Default() *ConfigAuth {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	conf.User = ""
	conf.Password = ""

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

	if strings.ReplaceAll(conf.Db, " ", "") == "" {
		err = ErrDatabaseNameIsNotSpecified
		return
	}

	if strings.ReplaceAll(conf.Host, " ", "") == "" {
		err = ErrHostIsNotSpecified
		return
	}

	if conf.Port == 0 {
		err = ErrPortIsNotSpecified
		return
	}

	if err = conf.Auth.Validate(); err != nil {
		return
	}

	return
}

// Validate - валидация конфигурации.
func (conf *ConfigAuth) Validate() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	if strings.ReplaceAll(conf.User, " ", "") == "" {
		err = ErrUserIsNotSpecified
		return
	}

	if strings.ReplaceAll(conf.Password, " ", "") == "" {
		err = ErrUserPasswordIsNotSpecified
		return
	}

	return
}

// ConnectionString - получение строки для подключения к базе данных.
func (conf *Config) ConnectionString() (str string) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelConfig)

		trace.FunctionCall()
		defer func() { trace.FunctionCallFinished(str) }()
	}

	var textOpts = []format.TextOption{
		{
			Key:   "host",
			Value: conf.Host,
		},
		{
			Key:   "port",
			Value: conf.Port,
		},
		{
			Key:   "user",
			Value: conf.Auth.User,
		},
		{
			Key:   "password",
			Value: conf.Auth.Password,
		},
		{
			Key:   "db",
			Value: conf.Db,
		},
	}

	str = format.Text(strings.Clone(templateConnectionString), textOpts...)

	return
}
