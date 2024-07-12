package internal

import (
	"context"
	"fmt"
	grpc_codes "google.golang.org/grpc/codes"
	"sm-box/pkg/errors/types"
)

type (
	// Internal - внутренняя реализация ошибки.
	Internal struct {
		Store *Store
		ctx   context.Context
	}
)

type (
	// Store - хранилище для построения ошибки.
	Store struct {
		ID     types.ID
		Type   types.ErrorType
		Status types.Status

		Err     error
		Message types.Message
		Details types.Details

		Others *StoreOthers
	}

	// StoreOthers -другие обьекты хранилище.
	StoreOthers struct {
		RestAPI   *RestAPIStore
		WebSocket *WebSocketStore
		Grpc      *GrpcStore
	}

	// RestAPIStore - хранилище для построения ошибок rest api.
	RestAPIStore struct {
		StatusCode int
	}

	// WebSocketStore - хранилище для построения ошибок web socket.
	WebSocketStore struct {
		StatusCode int
	}

	// GrpcStore - хранилище для построения ошибок grpc.
	GrpcStore struct {
		StatusCode grpc_codes.Code
	}
)

// New - создание внутренней реализации ошибки.
func New(store *Store) (i *Internal) {
	i = &Internal{
		Store: store,
		ctx:   context.Background(),
	}

	return
}

// ID - получение идентификатора ошибки.
func (i *Internal) ID() (id types.ID) {
	id = i.Store.ID
	return
}

// Type - получение типа ошибки.
func (i *Internal) Type() (t types.ErrorType) {
	t = i.Store.Type

	if t < types.TypeUnknown || t > types.TypeSystem {
		t = types.TypeUnknown
	}

	return
}

// Status - получение статуса ошибки.
func (i *Internal) Status() (s types.Status) {
	s = i.Store.Status

	if s < types.StatusUnknown || s > types.StatusFatal {
		s = types.StatusUnknown
		return
	}

	return
}

// Message - получение сообщения ошибки.
func (i *Internal) Message() (m string) {
	m = i.Store.Message.String()
	return
}

// Details - получение деталей ошибки.
func (i *Internal) Details() (m types.Details) {
	m = i.Store.Details
	return
}

// Error - получение текста исходной ошибки.
func (i *Internal) Error() (s string) {
	if i.Store.Err != nil {
		s = i.Store.Err.Error()
		return
	}

	s = i.Store.Message.String()

	return
}

// String - получение строкового представления ошибки.
func (i *Internal) String() (s string) {
	if i.Store.Message == nil {
		return
	}

	if i.Store.Err == nil {
		return i.Store.Message.String()
	}

	s = fmt.Sprintf("%s: '%s'. ", i.Store.Message.String(), i.Store.Err.Error())
	return
}

// SetError - установить значение исходной ошибки.
func (i *Internal) SetError(err error) {
	i.Store.Err = err
	return
}

// SetMessage - установить значение сообщения ошибки.
func (i *Internal) SetMessage(m types.Message) {
	i.Store.Message = m
	return
}
