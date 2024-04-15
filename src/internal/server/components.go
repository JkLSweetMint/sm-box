package server

import "sm-box/src/core/components/logger"

// Components - описание компонентов сервера.
type Components interface {
	Logger() logger.Logger
}

// components - компоненты сервера.
type components struct {
	logger logger.Logger
}

// Logger - получение компонента ведения журнала.
func (c *components) Logger() logger.Logger {
	return c.logger
}
