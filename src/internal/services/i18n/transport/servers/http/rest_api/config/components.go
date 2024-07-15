package config

import (
	"sm-box/pkg/core/components/tracer"
)

// Components - конфигурация компонентов http rest api.
type Components struct {
	AccessSystem *ComponentAccessSystem `json:"access_system" yaml:"AccessSystem" xml:"AccessSystem"`
}

// ComponentAccessSystem - конфигурация компонента системы доступа.
type ComponentAccessSystem struct {
	CookieKeyForSessionToken string `json:"cookie_key_for_session_token"  yaml:"CookieKeyForSessionToken"  xml:"cookie_key_for_session_token,attr"`
}

// FillEmptyFields - заполнение пустых полей конфигурации
func (conf *Components) FillEmptyFields() *Components {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	if conf.AccessSystem == nil {
		conf.AccessSystem = new(ComponentAccessSystem)
	}

	conf.AccessSystem.FillEmptyFields()

	return conf
}

// FillEmptyFields - заполнение пустых полей конфигурации
func (conf *ComponentAccessSystem) FillEmptyFields() *ComponentAccessSystem {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *Components) Default() *Components {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	conf.AccessSystem = new(ComponentAccessSystem).Default()

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *ComponentAccessSystem) Default() *ComponentAccessSystem {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	conf.CookieKeyForSessionToken = "box_session"

	return conf
}

// Validate - валидация конфигурации.
func (conf *Components) Validate() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	if err = conf.AccessSystem.Validate(); err != nil {
		return
	}

	return
}

// Validate - валидация конфигурации.
func (conf *ComponentAccessSystem) Validate() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	return
}
