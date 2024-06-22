package config

import "sm-box/pkg/core/components/tracer"

// Postman - конфигурация для генерации Postman коллекции.
type Postman struct {
	Protocol string `json:"protocol" yaml:"Protocol" xml:"Protocol"`
	Host     string `json:"host"     yaml:"Host"     xml:"Host"`
}

// FillEmptyFields - заполнение пустых полей конфигурации
func (conf *Postman) FillEmptyFields() *Postman {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *Postman) Default() *Postman {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	conf.Protocol = "https"
	conf.Host = "box.samgk.ru"

	return conf
}

// Validate - валидация конфигурации.
func (conf *Postman) Validate() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	return
}
