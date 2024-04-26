package app

import "sm-box/pkg/core/components/logger"

// Components - описание компонентов коробки.
type Components interface {
	Logger() logger.Logger
}

// components - компоненты коробки.
type components struct {
	logger logger.Logger
}

// Logger - получение компонента ведения журнала.
func (c *components) Logger() logger.Logger {
	return c.logger
}
