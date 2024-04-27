package config

import (
	"github.com/gofiber/fiber/v3"
	cache_middleware "github.com/gofiber/fiber/v3/middleware/cache"
	"github.com/gofiber/utils/v2"
	"sm-box/pkg/core/components/tracer"
	"time"
)

// MiddlewareCache - конфигурация промежуточного программного обеспечения для сжатия ответа с сервера.
type MiddlewareCache struct {
	// Enable - включить промежуточный слой.
	Enable bool `json:"enable" yaml:"Enable" xml:"enable,attr"`

	// Истечение срока действия - это время, в течение которого сохранится кэшированный ответ
	//
	// Необязательно. По умолчанию: 1 * time.Minute
	Expiration time.Duration `json:"expiration" yaml:"Expiration" xml:"Expiration"`

	// CacheHeader заголовок ответа указывает состояние кэша со следующим возможным возвращаемым значением
	//
	// hit, miss, unreachable
	//
	// Необязательно. По умолчанию: X-Cache
	CacheHeader string `json:"cache_header" yaml:"CacheHeader" xml:"CacheHeader"`

	// Включает кэширование на стороне клиента, если установлено значение true
	//
	// Необязательно. По умолчанию: false
	CacheControl bool `json:"cache_control" yaml:"CacheControl" xml:"CacheControl"`

	// Позволяет сохранять дополнительные заголовки, сгенерированные промежуточными слоями и обработчиком
	//
	// По умолчанию: false
	StoreResponseHeaders bool `json:"store_response_headers" yaml:"StoreResponseHeaders" xml:"StoreResponseHeaders"`

	// Максимальное количество байт текста ответа, одновременно хранящегося в кэше. При достижении лимита
	// записи с ближайшим сроком действия удаляются, чтобы освободить место для новых.
	// 0 означает отсутствие ограничения
	//
	// По умолчанию: 0
	MaxBytes uint `json:"max_bytes" yaml:"MaxBytes" xml:"MaxBytes"`

	// HTTP-методы для кэширования.
	// Промежуточное программное обеспечение просто кэширует маршруты своих методов в этом фрагменте.
	//
	// По умолчанию: []string{fiber.MethodGet, fiber.MethodHead}
	Methods []string `json:"methods" yaml:"Methods" xml:"Methods>Method"`
}

// FillEmptyFields - заполнение пустых полей конфигурации
func (conf *MiddlewareCache) FillEmptyFields() *MiddlewareCache {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *MiddlewareCache) Default() *MiddlewareCache {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	conf.Enable = true
	conf.Expiration = 1 * time.Minute
	conf.CacheHeader = "X-Cache"
	conf.CacheControl = false
	conf.StoreResponseHeaders = false
	conf.MaxBytes = 0
	conf.Methods = []string{fiber.MethodGet, fiber.MethodHead}

	return conf
}

// Validate - валидация конфигурации.
func (conf *MiddlewareCache) Validate() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	return
}

// ToFiberConfig - преобразовать конфигурацию в формат cache.Config.
func (conf *MiddlewareCache) ToFiberConfig() (c cache_middleware.Config) {
	c = cache_middleware.Config{
		Next:         nil,
		Expiration:   conf.Expiration,
		CacheHeader:  conf.CacheHeader,
		CacheControl: conf.CacheControl,
		KeyGenerator: func(c fiber.Ctx) string {
			return utils.CopyString(c.Path())
		},
		ExpirationGenerator:  nil,
		StoreResponseHeaders: conf.StoreResponseHeaders,
		Storage:              nil,
		MaxBytes:             conf.MaxBytes,
		Methods:              conf.Methods,
	}

	return
}
