package config

// Config - конфигурация http rest api.
type Config struct {
	Engine      *Engine      `json:"engine"      yaml:"Engine"      xml:"Engine"`
	Middlewares *Middlewares `json:"middlewares" yaml:"Middlewares" xml:"Middlewares"`
}

// FillEmptyFields - заполнение обязательных пустых полей конфигурации
func (conf *Config) FillEmptyFields() *Config {
	if conf.Engine == nil {
		conf.Engine = new(Engine)
	}

	if conf.Middlewares == nil {
		conf.Middlewares = new(Middlewares)
	}

	conf.Engine.FillEmptyFields()
	conf.Middlewares.FillEmptyFields()

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *Config) Default() *Config {
	conf.Engine = new(Engine).Default()
	conf.Middlewares = new(Middlewares).Default()

	return conf
}

// Validate - валидация конфигурации.
func (conf *Config) Validate() (err error) {
	if err = conf.Engine.Validate(); err != nil {
		return
	}

	if err = conf.Middlewares.Validate(); err != nil {
		return
	}

	return
}
