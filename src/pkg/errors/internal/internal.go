package internal

import (
	"context"
	"fmt"
	"sm-box/pkg/errors/types"
)

type (
	// Internal - внутренняя реализация ошибки.
	Internal struct {
		id     types.ID
		t      types.ErrorType
		status types.Status

		err     error
		message types.Message
		details types.Details

		ctx context.Context
	}

	// Constructor - конструктор для построения ошибки.
	Constructor struct {
		ID     types.ID
		Type   types.ErrorType
		Status types.Status

		Err     error
		Message types.Message
		Details types.Details
	}
)

// New - создание внутренней реализаци ошибки.
func New(cnst Constructor) (i *Internal) {
	i = &Internal{
		id:     cnst.ID,
		t:      cnst.Type,
		status: cnst.Status,

		err:     cnst.Err,
		message: cnst.Message,
		details: cnst.Details,

		ctx: context.Background(),
	}

	return
}

// ID - получение идентификатора ошибки.
func (i *Internal) ID() (id types.ID) {
	id = i.id
	return
}

// Type - получение типа ошибки.
func (i *Internal) Type() (t types.ErrorType) {
	t = i.t

	if t < types.TypeUnknown || t > types.TypeSystem {
		t = types.TypeUnknown
	}

	return
}

// Status - получение статуса ошибки.
func (i *Internal) Status() (s types.Status) {
	s = i.status

	if s < types.StatusUnknown || s > types.StatusFatal {
		s = types.StatusUnknown
		return
	}

	return
}

// Message - получение сообщения ошибки.
func (i *Internal) Message() (m string) {
	m = i.message.String()
	return
}

// Details - получение деталей ошибки.
func (i *Internal) Details() (m types.Details) {
	m = i.details
	return
}

// Error - получение текста исходной ошибки.
func (i *Internal) Error() (s string) {
	if i.err != nil {
		s = i.err.Error()
		return
	}

	s = i.message.String()

	return
}

// String - получение строкового представления ошибки.
func (i *Internal) String() (s string) {
	if i.message == nil {
		return
	}

	if i.err == nil {
		return i.message.String()
	}

	s = fmt.Sprintf("%s: '%s'. ", i.message.String(), i.err.Error())
	return
}

// SetError - установить значение исходной ошибки.
func (i *Internal) SetError(err error) {
	i.err = err
	return
}

// SetMessage - установить значение сообщения ошибки.
func (i *Internal) SetMessage(m types.Message) {
	i.message = m
	return
}
