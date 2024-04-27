package config

import (
	cors_middleware "github.com/gofiber/fiber/v3/middleware/cors"
	"sm-box/pkg/core/components/tracer"
	"strings"
)

// MiddlewareCors - конфигурация промежуточного программного обеспечения по кросс-доменным запросам.
type MiddlewareCors struct {
	// Enable - включить промежуточный слой.
	Enable bool `json:"enable" yaml:"Enable" xml:"enable,attr"`

	// Разрешить источник определяет список источников, разделенных запятыми, которые могут обращаться к ресурсу.
	//
	// Необязательно. Значение по умолчанию "*"
	AllowOrigins []string `json:"allow_origins" yaml:"AllowOrigins" xml:"AllowOrigins>AllowOrigin"`

	// Разрешить методы определяет список методов, разрешенных при доступе к ресурсу.
	// Используется в ответ на предполетный запрос.
	//
	// Необязательный. Значение по умолчанию: "GET,POST,HEAD,PUT,DELETE,PATCH"
	AllowMethods []string `json:"allow_methods" yaml:"AllowMethods" xml:"AllowMethods>AllowMethod"`

	// Разрешить заголовки определяет список заголовков запроса, которые могут быть использованы при
	// отправке фактического запроса. Это ответ на предварительный запрос.
	//
	// Необязательно. Значение по умолчанию "".
	AllowHeaders []string `json:"allow_headers" yaml:"AllowHeaders" xml:"AllowHeaders>AllowHeader"`

	// Разрешить использование учетных данных указывает, может ли быть получен ответ на запрос,
	// если флаг учетных данных имеет значение true. При использовании в качестве части
	// ответа на предполетный запрос, это указывает, может ли
	// фактический запрос быть выполнен с использованием учетных данных. Примечание: Если значение равно true,
	// для параметра AllowOrigins
	// не может быть установлен подстановочный знак ("*") для предотвращения уязвимостей в системе безопасности.
	//
	// Необязательно. Значение по умолчанию false.
	AllowCredentials bool `json:"allow_credentials" yaml:"AllowCredentials" xml:"AllowCredentials"`

	// Expose Headers определяет заголовки белого списка, к которым клиентам разрешен доступ.
	//
	// Необязательно. Значение по умолчанию "".
	ExposeHeaders []string `json:"expose_headers" yaml:"ExposeHeaders" xml:"ExposeHeaders>ExposeHeader"`

	// Максимальный возраст указывает, как долго (в секундах) могут кэшироваться результаты предполетного запрос.
	//
	// Если вы укажете значение MaxAge 0, заголовок Access-Control-Max-Age добавляться не будет, и
	// браузер по умолчанию будет использовать 5 секунд.
	// Чтобы полностью отключить кэширование, передайте значение MaxAge отрицательным. Это установит
	// заголовок Access-Control-Max-Age равным 0.
	//
	// Необязательно. Значение по умолчанию равно 0.
	MaxAge int `json:"max_age" yaml:"MaxAge" xml:"MaxAge"`

	// Разрешить частную сеть указывает, следует ли установить для заголовка ответа Access-Control-Allow-Private-Network
	// значение true, разрешающее запросы из частных сетей.
	//
	// Необязательно. Значение по умолчанию false.
	AllowPrivateNetwork bool `json:"allow_private_network" yaml:"AllowPrivateNetwork" xml:"AllowPrivateNetwork"`
}

// FillEmptyFields - заполнение пустых полей конфигурации
func (conf *MiddlewareCors) FillEmptyFields() *MiddlewareCors {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	if conf.AllowOrigins == nil {
		conf.AllowOrigins = make([]string, 0)
	}

	if conf.AllowMethods == nil {
		conf.AllowMethods = make([]string, 0)
	}

	if conf.AllowHeaders == nil {
		conf.AllowHeaders = make([]string, 0)
	}

	if conf.ExposeHeaders == nil {
		conf.ExposeHeaders = make([]string, 0)
	}

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *MiddlewareCors) Default() *MiddlewareCors {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	conf.Enable = true
	conf.AllowOrigins = make([]string, 0)
	conf.AllowMethods = []string{
		"GET",
		"POST",
		"HEAD",
		"PUT",
		"DELETE",
		"PATCH",
		"OPTIONS",
	}
	conf.AllowHeaders = []string{
		"Origin",
		"Access-Control-Allow-Headers",
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Credentials",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"Cookie",
		"Authorization",
		"RequestURL",
		"Set-Cookie",
		"X-Requested-With",
	}
	conf.AllowCredentials = false
	conf.ExposeHeaders = []string{
		"Content-Disposition",
		"Content-Length",
	}
	conf.MaxAge = 36000
	conf.AllowPrivateNetwork = true

	return conf
}

// Validate - валидация конфигурации.
func (conf *MiddlewareCors) Validate() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	return
}

// ToFiberConfig - преобразовать конфигурацию в формат cors.Config.
func (conf *MiddlewareCors) ToFiberConfig() (c cors_middleware.Config) {
	c = cors_middleware.Config{
		Next:                nil,
		AllowOriginsFunc:    nil,
		AllowOrigins:        strings.Join(conf.AllowOrigins, ", "),
		AllowMethods:        strings.Join(conf.AllowMethods, ", "),
		AllowHeaders:        strings.Join(conf.AllowHeaders, ", "),
		AllowCredentials:    conf.AllowCredentials,
		ExposeHeaders:       strings.Join(conf.ExposeHeaders, ", "),
		MaxAge:              conf.MaxAge,
		AllowPrivateNetwork: conf.AllowPrivateNetwork,
	}

	return
}
