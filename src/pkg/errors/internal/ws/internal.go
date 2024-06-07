package ws

import (
	"sm-box/pkg/errors/internal"
)

type (
	// Internal - внутренняя реализация ошибки web socket.
	Internal struct {
		*internal.Internal

		statusCode int
	}

	// Constructor - конструктор для построения ошибки.
	Constructor struct {
		internal.Constructor

		StatusCode int
	}
)

// New - создание внутренней реализаци ошибки web socket.
func New(cnst Constructor) (i *Internal) {
	i = &Internal{
		Internal: internal.New(cnst.Constructor),

		statusCode: cnst.StatusCode,
	}

	return
}

// StatusCode - получение статус кода http web socket ошибки.
func (i *Internal) StatusCode() (c int) {
	c = i.statusCode
	return
}
