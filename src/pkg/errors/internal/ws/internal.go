package ws

import (
	"sm-box/pkg/errors/internal"
)

type (
	// Internal - внутренняя реализация ошибки web socket.
	Internal struct {
		*internal.Internal
	}
)

// New - создание внутренней реализации ошибки web socket.
func New(store *internal.Store) (i *Internal) {
	i = &Internal{
		Internal: internal.New(store),
	}

	return
}

// StatusCode - получение статус кода http web socket ошибки.
func (i *Internal) StatusCode() (c int) {
	c = i.Internal.Store.Others.WebSocket.StatusCode
	return
}
