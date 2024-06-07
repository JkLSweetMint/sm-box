package rest_api

import (
	"sm-box/pkg/errors/internal"
)

type (
	// Internal - внутренняя реализация ошибки rest api.
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

// New - создание внутренней реализаци ошибки rest api.
func New(cnst Constructor) (i *Internal) {
	i = &Internal{
		Internal: internal.New(cnst.Constructor),

		statusCode: cnst.StatusCode,
	}

	return
}

// StatusCode - получение статус кода http rest api ошибки.
func (i *Internal) StatusCode() (c int) {
	c = i.statusCode
	return
}
