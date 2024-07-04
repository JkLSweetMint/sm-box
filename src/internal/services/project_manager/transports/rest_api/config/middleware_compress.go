package config

import (
	compress_middleware "github.com/gofiber/fiber/v3/middleware/compress"
	"sm-box/pkg/core/components/tracer"
)

// MiddlewareCompress - конфигурация промежуточного программного обеспечения для сжатия ответа с сервера.
type MiddlewareCompress struct {
	// Enable - включить промежуточный слой.
	Enable bool `json:"enable" yaml:"Enable" xml:"enable,attr"`

	// Уровень определяет алгоритм сжатия
	//
	// Необязательно. По умолчанию: LevelDefault
	// LevelDisabled: -1
	// LevelDefault: 0
	// LevelBestSpeed: 1
	// LevelBestCompression: 2
	Level compress_middleware.Level `json:"level" yaml:"Level" xml:"Level"`
}

// FillEmptyFields - заполнение пустых полей конфигурации
func (conf *MiddlewareCompress) FillEmptyFields() *MiddlewareCompress {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *MiddlewareCompress) Default() *MiddlewareCompress {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	conf.Enable = true
	conf.Level = compress_middleware.LevelBestSpeed

	return conf
}

// Validate - валидация конфигурации.
func (conf *MiddlewareCompress) Validate() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	return
}

// ToFiberConfig - преобразовать конфигурацию в формат compress.Config.
func (conf *MiddlewareCompress) ToFiberConfig() (c compress_middleware.Config) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(c) }()
	}

	c = compress_middleware.Config{
		Level: conf.Level,
	}

	return
}