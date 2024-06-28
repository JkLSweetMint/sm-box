package rest_api

import (
	"sm-box/pkg/errors/internal"
)

type (
	// Internal - внутренняя реализация ошибки rest api.
	Internal struct {
		*internal.Internal
	}
)

// New - создание внутренней реализации ошибки rest api.
func New(store *internal.Store) (i *Internal) {
	i = &Internal{
		Internal: internal.New(store),
	}

	return
}

// StatusCode - получение статус кода http rest api ошибки.
func (i *Internal) StatusCode() (c int) {
	c = i.Internal.Store.Others.RestAPI.StatusCode
	return
}
