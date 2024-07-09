package app

import "sm-box/pkg/core/components/logger"

// Components - описание компонентов приложения.
type Components interface {
	Logger() logger.Logger
}

// components - компоненты приложения.
type components struct {
	logger logger.Logger
}

// Logger - получение компонента ведения журнала.
func (c *components) Logger() logger.Logger {
	return c.logger
}
