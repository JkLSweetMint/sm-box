package core

import (
	"sm-box/src/core/components/logger"
)

// components - компоненты ядра системы.
type components struct {
	logger logger.Logger
}

// Logger - получение компонента ведения журнала.
func (c *components) Logger() logger.Logger {
	return c.logger
}
