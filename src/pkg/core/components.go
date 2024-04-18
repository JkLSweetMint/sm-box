package core

import (
	"sm-box/src/pkg/core/components/logger"
)

// Components - описание компонентов ядра системы.
type Components interface {
	Logger() logger.Logger
}

// components - компоненты ядра системы.
type components struct {
	logger logger.Logger
}

// Logger - получение компонента ведения журнала.
func (c *components) Logger() logger.Logger {
	return c.logger
}
