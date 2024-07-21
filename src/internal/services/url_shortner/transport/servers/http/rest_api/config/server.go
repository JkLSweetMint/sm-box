package config

import (
	"github.com/gofiber/fiber/v3"
	"sm-box/pkg/core/components/tracer"
	"time"
)

// Server - конфигурация http сервера.
type Server struct {
	// Включает HTTP-заголовок "Server: value".
	//
	// По умолчанию: ""
	ServerHeader string `json:"server_header" yaml:"ServerHeader" xml:"ServerHeader"`

	// Если установлено значение true, маршрутизатор обрабатывает "/foo" и "/foo/" как разные.
	// По умолчанию это отключено, и оба "/foo" и "/foo/" будут выполнять один и тот же обработчик.
	//
	// По умолчанию: false
	StrictRouting bool `json:"strict_routing" yaml:"StrictRouting" xml:"StrictRouting"`

	// Если установлено значение true, включается маршрутизация с учетом регистра.
	// Например, "/FoO" и "/foo" рассматриваются как разные маршруты.
	// По умолчанию это отключено, и оба "/FoO" и "/foo" будут выполнять один и тот же обработчик.
	//
	// По умолчанию: false
	CaseSensitive bool `json:"case_sensitive" yaml:"CaseSensitive" xml:"CaseSensitive"`

	// Если установлено значение true, это отменяет обещание выделения 0 в определенных
	// случаях, чтобы получить доступ к значениям обработчика (например, к телам запросов)
	// неизменяемым образом, чтобы эти значения были доступны, даже если вы возвращаете из обработчика.
	//
	// По умолчанию: false
	Immutable bool `json:"immutable" yaml:"Immutable" xml:"Immutable"`

	// Если установлено значение true, преобразует все закодированные символы в маршруте обратно
	// перед установкой пути для контекста, так что маршрутизация,
	// возврат текущего URL-адреса из контекста `ctx.Path()`
	// и параметры `ctx.Params(%key%)` с декодированными символами будут работать
	//
	// По умолчанию: false
	UnescapePath bool `json:"unescape_path" yaml:"UnescapePath" xml:"UnescapePath"`

	// Максимальный размер тела, который принимает сервер.
	// -1 уменьшит любой размер тела
	//
	// По умолчанию: 4 * 1024 * 1024
	BodyLimit int `json:"body_limit" yaml:"BodyLimit" xml:"BodyLimit"`

	// Максимальное количество одновременных подключений.
	//
	// По умолчанию: 256 * 1024
	Concurrency int `json:"concurrency" yaml:"Concurrency" xml:"Concurrency"`

	// Количество времени, отведенное для чтения всего запроса, включая тело.
	// Оно сбрасывается после возврата обработчика запроса.
	// Крайний срок чтения соединения сбрасывается при открытии соединения.
	//
	// По умолчанию: неограниченный
	ReadTimeout time.Duration `json:"read_timeout" yaml:"ReadTimeout" xml:"ReadTimeout"`

	// Максимальная продолжительность до истечения времени ожидания записи ответа.
	// Она сбрасывается после возврата обработчика запроса.
	//
	// По умолчанию: неограниченно
	WriteTimeout time.Duration `json:"write_timeout" yaml:"WriteTimeout" xml:"WriteTimeout"`

	// Максимальное время ожидания следующего запроса при включенной функции keep-alive.
	// Если idleTimeout равно нулю, используется значение ReadTimeout.
	//
	// По умолчанию: неограниченно
	IdleTimeout time.Duration `json:"idle_timeout" yaml:"IdleTimeout" xml:"IdleTimeout"`

	// Размер буфера для чтения запросов для каждого соединения.
	// Это также ограничивает максимальный размер заголовка.
	// Увеличьте этот буфер, если ваши клиенты отправляют запросы размером в несколько килобайт
	// и/или заголовки размером в несколько килобайт (например, БОЛЬШИЕ файлы cookie).
	//
	// По умолчанию: 4096 байт.
	ReadBufferSize int `json:"read_buffer_size" yaml:"ReadBufferSize" xml:"ReadBufferSize"`

	// Размер буфера для каждого соединения для записи ответов.
	//
	// По умолчанию: 4096 байт.
	WriteBufferSize int `json:"write_buffer_size" yaml:"WriteBufferSize" xml:"WriteBufferSize"`

	// Суффикс сжатого файла добавляет суффикс к исходному имени файла и
	// пытается сохранить полученный сжатый файл под новым именем файла.
	//
	// По умолчанию: ".fiber.gz"
	CompressedFileSuffix string `json:"compressed_file_suffix" yaml:"CompressedFileSuffix" xml:"CompressedFileSuffix"`

	// Заголовок прокси позволит c.IP() возвращать значение заданного ключа заголовка
	// По умолчанию c.IP() возвращает удаленный IP из TCP-соединения
	// Это свойство может быть полезно, если вы используете балансировщик нагрузки: X-Forwarded-*
	// ПРИМЕЧАНИЕ: заголовки легко подделываются, а обнаруженные IP-адреса ненадежны.
	//
	// По умолчанию: ""
	ProxyHeader string `json:"proxy_header" yaml:"ProxyHeader" xml:"ProxyHeader"`

	// GET отклоняет все запросы, не связанные с GET, только в том случае, если установлено значение true.
	// Эта опция полезна для защиты серверов от DoS
	// принимает только запросы GET. Размер запроса ограничен
	// размером буфера чтения, если установлено значение GETOnly.
	//
	// По умолчанию: false
	GETOnly bool `json:"get_only" yaml:"GETOnly" xml:"GETOnly"`

	// Если установлено значение true, отключает поддерживаемые соединения.
	// Сервер закроет входящие соединения после отправки первого ответа клиенту.
	//
	// По умолчанию: false
	DisableKeepalive bool `json:"disable_keepalive" yaml:"DisableKeepalive" xml:"DisableKeepalive"`

	// Если установлено значение true, приводит к исключению заголовка даты по умолчанию из ответа.
	//
	// По умолчанию: false
	DisableDefaultDate bool `json:"disable_default_date" yaml:"DisableDefaultDate" xml:"DisableDefaultDate"`

	// Если установлено значение true, приводит к исключению заголовка Content-Type по умолчанию из ответа.
	//
	// По умолчанию: false
	DisableDefaultContentType bool `json:"disable_default_content_type" yaml:"DisableDefaultContentType" xml:"DisableDefaultContentType"`

	// Если установлено значение true, нормализация заголовка отключается.
	// По умолчанию нормализуются все имена заголовков: conteNT-tYPE -> Content-Type.
	//
	// По умолчанию: false
	DisableHeaderNormalizing bool `json:"disable_header_normalizing" yaml:"DisableHeaderNormalizing" xml:"DisableHeaderNormalizing"`

	// Название приложения.
	AppName string `json:"app_name" yaml:"AppName" xml:"AppName"`

	// Тело запроса Stream включает потоковую передачу тела запроса,
	// и вызывает обработчик раньше, когда заданное тело
	// превышает текущее ограничение.
	StreamRequestBody bool `json:"stream_request_body" yaml:"StreamRequestBody" xml:"StreamRequestBody"`

	// Не будет выполнять предварительный анализ данных составной формы, если установлено значение true.
	//
	// Этот параметр полезен для серверов, которые хотят обрабатывать
	// данные составной формы как двоичный объект или выбирать, когда выполнять синтаксический анализ данных.
	//
	// Сервер выполняет предварительный анализ данных составной формы по умолчанию.
	DisablePreParseMultipartForm bool `json:"disable_pre_parse_multipart_form" yaml:"DisablePreParseMultipartForm" xml:"DisablePreParseMultipartForm"`

	// Агрессивно сокращает использование памяти за счет более высокой загрузки процессора
	// если установлено значение true.
	//
	// Попробуйте включить эту опцию, только если сервер потребляет слишком много памяти
	// обслуживает в основном незанятые соединения для поддержания активности. Это может сократить использование памяти
	// более чем на 50%.
	//
	// По умолчанию: false
	ReduceMemoryUsage bool `json:"reduce_memory_usage" yaml:"ReduceMemoryUsage" xml:"ReduceMemoryUsage"`

	// Известными сетями являются "tcp", "tcp4" (только для IPv4), "tcp6" (только для IPv6)
	// ПРЕДУПРЕЖДЕНИЕ: Если для предварительной настройки установлено значение true, можно выбрать только "tcp4" и "tcp6".
	//
	// По умолчанию: NetworkTCP4
	Network string `json:"network" yaml:"Network" xml:"Network"`

	// Если вы окажетесь за каким-либо прокси, например, за балансировщиком нагрузки,
	// тогда определенная информация заголовка может быть отправлена вам с использованием специальных заголовков X-Forwarded-* или заголовка Forwarded.
	// Например, HTTP-заголовок Host обычно используется для возврата запрошенного хоста.
	// Но когда вы находитесь за прокси, фактический хост может храниться в заголовке X-Forwarded-Host.
	//
	// Если вы находитесь за прокси, вам следует включить TrustedProxyCheck, чтобы предотвратить подмену заголовка.
	// Если вы включите EnableTrustedProxyCheck и оставите TrustedProxies пустыми, Fiber пропустит
	// все заголовки, которые могут быть подделаны.
	// Если запрос ip в белом списке TrustedProxies, то:
	// 1. c.Protocol() получает значение из заголовка X-Forwarded-Proto, X-Forwarded-Protocol, X-Forwarded-Ssl или X-Url-Scheme
	// 2. c.IP() получает значение из заголовка ProxyHeader.
	// 3. c.Hostname() получает значение из заголовка X-Forwarded-Host
	// Но если ip запроса НЕТ в белом списке доверенных прокси, то:
	// 1. c.Protocol() не получит значение из заголовка X-Forwarded-Proto, X-Forwarded-Protocol, X-Forwarded-Ssl или X-Url-Scheme,
	// вернет https в случае, когда приложение обрабатывает tls-соединение, или http в противном случае
	// 2. c.IP() НЕ получит значение из заголовка ProxyHeader, вернет RemoteIP() из контекста fasthttp
	// 3. c.Hostname() НЕ получит значение из заголовка X-Forwarded-Host, fasthttp.Запрос.URI().Host()
	// будет использоваться для получения имени хоста.
	//
	// По умолчанию: false
	EnableTrustedProxyCheck bool `json:"enable_trusted_proxy_check" yaml:"EnableTrustedProxyCheck" xml:"EnableTrustedProxyCheck"`

	// Прочитайте документ EnableTrustedProxyCheck.
	//
	// По умолчанию: []строка
	TrustedProxies []string `json:"trusted_proxies" yaml:"TrustedProxies" xml:"TrustedProxies"`

	// Если установлено значение true, c.IP() и c.IPs() будут проверять IP-адреса перед их возвратом.
	// Кроме того, c.IP() вернет только первый действительный IP, а не только необработанный заголовок
	// ПРЕДУПРЕЖДЕНИЕ: с этим связаны затраты на производительность.
	//
	// По умолчанию: false
	EnableIPValidation bool `json:"enable_ip_validation" yaml:"EnableIPValidation" xml:"EnableIPValidation"`

	// Методы запроса обеспечивают возможность настройки HTTP-методов. Вы можете добавлять/удалять методы по своему усмотрению.
	//
	// Необязательно, по умолчанию все методы.
	RequestMethods []string `json:"request_methods" yaml:"RequestMethods" xml:"RequestMethods"`

	// Включение разделения в парсерах разбивает параметры запроса/тела/заголовка на запятую, когда это значение равно true.
	// Например, вы можете использовать его для разбора нескольких значений из параметра запроса, подобного этому:
	// /api?foo=bar,baz == foo[]=bar&foo[]=baz
	//
	// Необязательно. По умолчанию: false
	EnableSplittingOnParsers bool `json:"enable_splitting_on_parsers" yaml:"EnableSplittingOnParsers" xml:"EnableSplittingOnParsers"`

	// Адрес сервера.
	Addr string `json:"addr" yaml:"Addr" xml:"addr,attr"`

	// Расположение сервера.
	Location string `json:"location" yaml:"Location" xml:"location,attr"`

	// Название сервера.
	Name string `json:"name" yaml:"Name" xml:"name,attr"`

	// Версия сервера
	Version string `json:"version" yaml:"Version" xml:"version,attr"`

	// Домен сервера
	Domain string `json:"domain" yaml:"Domain" xml:"domain,attr"`
}

// FillEmptyFields - заполнение пустых полей конфигурации
func (conf *Server) FillEmptyFields() *Server {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *Server) Default() *Server {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	conf.ServerHeader = "Url shortner service"
	conf.StrictRouting = true
	conf.CaseSensitive = false
	conf.Immutable = false
	conf.UnescapePath = false
	conf.BodyLimit = 4 * 1024 * 1024
	conf.Concurrency = 256 * 1024
	conf.ReadTimeout = 0
	conf.WriteTimeout = 0
	conf.IdleTimeout = 0
	conf.ReadBufferSize = 4096
	conf.WriteBufferSize = 4096
	conf.CompressedFileSuffix = ".service.gz"
	conf.ProxyHeader = ""
	conf.GETOnly = false
	conf.DisableKeepalive = false
	conf.DisableDefaultDate = false
	conf.DisableDefaultContentType = false
	conf.DisableHeaderNormalizing = false
	conf.AppName = "Url shortner service"
	conf.DisablePreParseMultipartForm = false
	conf.StreamRequestBody = false
	conf.ReduceMemoryUsage = false
	conf.Network = "tcp4"
	conf.EnableTrustedProxyCheck = false
	conf.TrustedProxies = make([]string, 0)
	conf.EnableIPValidation = false
	conf.RequestMethods = make([]string, 0)
	conf.EnableSplittingOnParsers = false
	conf.Addr = "0.0.0.0:8080"
	conf.Location = "/url-shortner"
	conf.Name = "api"
	conf.Version = "v1.0"
	conf.Domain = "box.samgk.ru"

	return conf
}

// Validate - валидация конфигурации.
func (conf *Server) Validate() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	return
}

// ToFiberConfig - преобразовать конфигурацию в формат fiber.Config.
func (conf *Server) ToFiberConfig() (c fiber.Config) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(c) }()
	}

	c = fiber.Config{
		ServerHeader:                 conf.ServerHeader,
		StrictRouting:                conf.StrictRouting,
		CaseSensitive:                conf.CaseSensitive,
		Immutable:                    conf.Immutable,
		UnescapePath:                 conf.UnescapePath,
		BodyLimit:                    conf.BodyLimit,
		Concurrency:                  conf.Concurrency,
		Views:                        nil,
		ViewsLayout:                  "",
		PassLocalsToViews:            false,
		ReadTimeout:                  conf.ReadTimeout,
		WriteTimeout:                 conf.WriteTimeout,
		IdleTimeout:                  conf.IdleTimeout,
		ReadBufferSize:               conf.ReadBufferSize,
		WriteBufferSize:              conf.WriteBufferSize,
		CompressedFileSuffix:         conf.CompressedFileSuffix,
		ProxyHeader:                  conf.ProxyHeader,
		GETOnly:                      conf.GETOnly,
		ErrorHandler:                 nil,
		DisableKeepalive:             conf.DisableKeepalive,
		DisableDefaultDate:           conf.DisableDefaultDate,
		DisableDefaultContentType:    conf.DisableDefaultContentType,
		DisableHeaderNormalizing:     conf.DisableHeaderNormalizing,
		AppName:                      conf.AppName,
		StreamRequestBody:            conf.StreamRequestBody,
		DisablePreParseMultipartForm: conf.DisablePreParseMultipartForm,
		ReduceMemoryUsage:            conf.ReduceMemoryUsage,
		JSONEncoder:                  nil,
		JSONDecoder:                  nil,
		XMLEncoder:                   nil,
		EnableTrustedProxyCheck:      conf.EnableTrustedProxyCheck,
		TrustedProxies:               conf.TrustedProxies,
		EnableIPValidation:           conf.EnableIPValidation,
		ColorScheme:                  fiber.Colors{},
		StructValidator:              nil,
		RequestMethods:               conf.RequestMethods,
		EnableSplittingOnParsers:     conf.EnableSplittingOnParsers,
	}

	return
}

// ToFiberListenConfig - преобразовать конфигурацию в формат fiber.ListenConfig.
func (conf *Server) ToFiberListenConfig() (c fiber.ListenConfig) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(c) }()
	}

	c = fiber.ListenConfig{
		DisableStartupMessage: true,
	}

	return
}
