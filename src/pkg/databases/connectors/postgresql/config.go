package postgresql

import (
	"fmt"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/tools/format"
	"strings"
	"time"
)

// templateConnectionString - шаблон строки для подключения к базе.
const templateConnectionString = "host={{host}} port={{port}} user={{user}} password={{password}} dbname={{dbname}}{{tags}}"

// Config - конфигурация коннектора к базе данных.
type Config struct {
	// DbName - название базы данных.
	DbName string `json:"db_name" yaml:"DbName" xml:"db_name,attr"`

	// Host - адрес базы данных.
	Host string `json:"host" yaml:"Host" xml:"host,attr"`

	// Port - порт базы данных.
	Port uint `json:"port" yaml:"Port" xml:"port,attr"`

	// Auth - конфигурация авторизации в базе данных.
	Auth *ConfigAuth `json:"auth" yaml:"Auth" xml:"Auth"`

	// Tags - теги для подключения
	Tags Tags `json:"tags" yaml:"Tags" xml:"Tags"`

	// MaxOpenConns - макс. кол-во. подключений к бд.
	MaxOpenConns int `json:"max_open_conns" yaml:"MaxOpenConns" xml:"MaxOpenConns"`

	// MaxIdleConns - макс. ожидаемых подключений к бд.
	MaxIdleConns int `json:"max_idle_conns" yaml:"MaxIdleConns" xml:"MaxIdleConns"`

	// ConnMaxLifetime - макс. время жизни соединения.
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime" yaml:"ConnMaxLifetime" xml:"ConnMaxLifetime"`
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

	if conf.Tags == nil {
		conf.Tags = make(Tags)
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

	conf.DbName = ""
	conf.Host = "127.0.0.1"
	conf.Port = 5432

	conf.Tags = make(Tags)

	conf.MaxOpenConns = 50
	conf.MaxIdleConns = 50

	conf.ConnMaxLifetime = time.Minute * 2

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

	if strings.ReplaceAll(conf.DbName, " ", "") == "" {
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

	var tags string

	// tags
	{
		if len(conf.Tags) > 0 {
			for key, value := range conf.Tags {
				tags += fmt.Sprintf(" %s=%s", key, value)
			}
		}
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
			Key:   "dbname",
			Value: conf.DbName,
		},
		{
			Key:   "tags",
			Value: tags,
		},
	}

	str = format.Text(strings.Clone(templateConnectionString), textOpts...)

	return
}
