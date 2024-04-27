package sqlite3

import (
	"sm-box/pkg/core/components/tracer"
	"time"
)

// Config - конфигурация коннектора к базе данных.
type Config struct {
	// Database - путь к файлу базы данных.
	Database string `json:"database" yaml:"Database" xml:"database,attr"`

	// MaxOpenConns - макс. кол-во. подключений к бд.
	MaxOpenConns int `json:"max_open_conns"    yaml:"MaxOpenConns"    xml:"MaxOpenConns"`

	// MaxIdleConns - макс. ожидаемых подключений к бд.
	MaxIdleConns int `json:"max_idle_conns"    yaml:"MaxIdleConns"    xml:"MaxIdleConns"`

	// ConnMaxLifetime - макс. время жизни соединения.
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime" yaml:"ConnMaxLifetime" xml:"ConnMaxLifetime"`
}

// FillEmptyFields - заполнение пустых полей конфигурации
func (conf *Config) FillEmptyFields() *Config {
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

	conf.Database = ""
	conf.MaxOpenConns = 50
	conf.MaxIdleConns = 50
	conf.ConnMaxLifetime = time.Minute * 2

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

	str = conf.Database

	return
}
