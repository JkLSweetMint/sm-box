package errors

import (
	"reflect"
	"sm-box/pkg/errors/entities/details"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/internal"
	"sm-box/pkg/errors/internal/rest_api"
	"sm-box/pkg/errors/internal/ws"
	"sm-box/pkg/errors/types"
)

// Базовая реализация
type (
	// Builder - функция дла построения базовой ошибки.
	Builder[T Error] func() T

	// Constructor - конструктор для построения ошибки.
	Constructor[T Error] struct {
		ID     types.ID
		Type   types.ErrorType
		Status types.Status

		Err     error
		Message types.Message
		Details types.Details

		addons *constructorAddons
	}

	// constructorAddons - дополнение к базовому конструктору.
	constructorAddons struct {
		RestAPI   *RestAPIConstructor
		WebSocket *WebSocketConstructor
	}

	// RestAPIConstructor - конструктор для построения ошибок rest api.
	RestAPIConstructor struct {
		StatusCode int
	}

	// WebSocketConstructor - конструктор для построения ошибок web socket.
	WebSocketConstructor struct {
		StatusCode int
	}
)

// Build - построение ошибки.
func (c Constructor[T]) Build() (fn Builder[T]) {
	c.fillEmptyField()

	fn = func() (e T) {
		switch reflect.TypeOf(new(T)).String() {
		case "*errors.Error":
			{
				var i = internal.New(internal.Constructor{
					ID:     c.ID,
					Type:   c.Type,
					Status: c.Status,

					Err:     c.Err,
					Message: c.Message.Clone(),
					Details: c.Details.Clone(),
				})

				e = interface{}(i).(T)
			}
		case "*errors.RestAPI":
			{
				var i = rest_api.New(rest_api.Constructor{
					Constructor: internal.Constructor{
						ID:     c.ID,
						Type:   c.Type,
						Status: c.Status,

						Err:     c.Err,
						Message: c.Message.Clone(),
						Details: c.Details.Clone(),
					},
					StatusCode: c.addons.RestAPI.StatusCode,
				})

				e = interface{}(i).(T)
			}
		case "*errors.WebSocket":
			{
				var i = ws.New(ws.Constructor{
					Constructor: internal.Constructor{
						ID:     c.ID,
						Type:   c.Type,
						Status: c.Status,

						Err:     c.Err,
						Message: c.Message.Clone(),
						Details: c.Details.Clone(),
					},
					StatusCode: c.addons.WebSocket.StatusCode,
				})

				e = interface{}(i).(T)
			}
		}

		return
	}

	return
}

// RestAPI - записать данные конструктора rest api ошибок.
func (c Constructor[T]) RestAPI(cstr RestAPIConstructor) Constructor[T] {
	if c.addons == nil {
		c.addons = new(constructorAddons)
	}

	c.addons.RestAPI = &cstr
	return c
}

// WebSocket - записать данные конструктора web socket ошибок.
func (c Constructor[T]) WebSocket(cstr WebSocketConstructor) Constructor[T] {
	if c.addons == nil {
		c.addons = new(constructorAddons)
	}

	c.addons.WebSocket = &cstr
	return c
}

// fillEmptyField - заполнение пустых полей структуры.
func (c *Constructor[T]) fillEmptyField() *Constructor[T] {
	if c.Message == nil {
		c.Message = new(messages.TextMessage)
	}

	if c.Details == nil {
		c.Details = new(details.Details)
	}

	if c.addons == nil {
		c.addons = new(constructorAddons)
	}

	return c
}
