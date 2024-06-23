package config

import "sm-box/pkg/core/components/tracer"

// Proxy - конфигурация проксирования.
type Proxy struct {
	Sources []*ProxySource `json:"sources" yaml:"Sources" xml:"Sources"`
}

// ProxySource - конфигурация источников проксирования.
type ProxySource struct {
	Path       string `json:"path"        yaml:"Path"       xml:"path,attr"`
	RemoteAddr string `json:"remote_addr" yaml:"RemoteAddr" xml:"remote_addr,attr"`
}

// FillEmptyFields - заполнение пустых полей конфигурации
func (conf *Proxy) FillEmptyFields() *Proxy {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	if conf.Sources == nil {
		conf.Sources = make([]*ProxySource, 0)
	}

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *Proxy) Default() *Proxy {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	conf.Sources = []*ProxySource{
		{
			Path:       "/authentication/v1.0/*",
			RemoteAddr: "http://localhost:8001",
		},
	}

	return conf
}

// Validate - валидация конфигурации.
func (conf *Proxy) Validate() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	return
}
