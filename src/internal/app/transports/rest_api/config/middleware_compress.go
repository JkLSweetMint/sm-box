package config

import (
	compress_middleware "github.com/gofiber/fiber/v3/middleware/compress"
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

// FillEmptyFields - заполнение обязательных пустых полей конфигурации
func (conf *MiddlewareCompress) FillEmptyFields() *MiddlewareCompress {
	return conf
}

// Default - запись стандартной конфигурации.
func (conf *MiddlewareCompress) Default() *MiddlewareCompress {
	conf.Enable = true
	conf.Level = compress_middleware.LevelBestSpeed

	return conf
}

// Validate - валидация конфигурации.
func (conf *MiddlewareCompress) Validate() (err error) {
	return
}

// ToFiberConfig - преобразовать конфигурацию в формат compress.Config.
func (conf *MiddlewareCompress) ToFiberConfig() (c compress_middleware.Config) {
	c = compress_middleware.Config{
		Level: conf.Level,
	}

	return
}
