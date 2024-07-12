package service

import "sm-box/pkg/core/components/logger"

// Components - описание компонентов сервиса.
type Components interface {
	Logger() logger.Logger
}

// components - компоненты сервиса.
type components struct {
	logger logger.Logger
}

// Logger - получение компонента ведения журнала.
func (c *components) Logger() logger.Logger {
	return c.logger
}
