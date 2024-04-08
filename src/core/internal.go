package core

import (
	"context"
	"sm-box/src/core/components/logger"
)

// core - ядро системы.
type core struct {
	components *components

	ctx  context.Context
	conf *Config

	instance interface {
		Shutdown() (err error)
		Boot() (err error)
		Serve() (err error)

		State() (state State)
	}
}

// Shutdown - завершение работы ядра.
// Остановить можно только ядро со статусом StateServed.
// Состояние ядра будет изменено на StateOff.
func (c *core) Shutdown() (err error) {
	return c.instance.Shutdown()
}

// Boot - загрузка ядра, построение системы, создание компонентов.
// Загрузить можно только ядро со статусом StateNew.
// Состояние ядра будет изменено на StateBooted.
func (c *core) Boot() (err error) {
	return c.instance.Boot()
}

// Serve - запуск обслуживания ядра.
// Запустить обслуживание можно только ядро со статусом StateBooted.
// Состояние ядра будет изменено на StateServed.
func (c *core) Serve() (err error) {
	return c.instance.Serve()
}

// State - получение состояния ядра системы.
//
// Возможные варианты состояния ядра:
//  1. StateNew    - "New";
//  2. StateBooted - "Booted";
//  3. StateServed - "Served";
//  4. StateOff    - "Off";
func (c *core) State() (state State) {
	return c.instance.State()
}

// Ctx - получение контекста ядра системы.
func (c *core) Ctx() (ctx context.Context) {
	return c.ctx
}

// Components - получение компонентов ядра системы.
func (c *core) Components() interface {
	Logger() logger.Logger
} {
	return c.components
}
