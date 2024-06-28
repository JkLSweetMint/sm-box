package errors

import (
	"reflect"
	"sm-box/pkg/errors/internal"
	"sm-box/pkg/errors/internal/rest_api"
	"sm-box/pkg/errors/internal/ws"
)

// ToError - преобразование ошибки в ошибку Error.
func ToError[T Error](err T) (newErr Error) {
	switch reflect.TypeOf(new(T)).String() {
	case "*errors.Error":
		{
			var e = interface{}(err).(*internal.Internal)

			newErr = e
		}
	case "*errors.RestAPI":
		{
			var e = interface{}(err).(*rest_api.Internal)

			newErr = e.Internal
		}
	case "*errors.WebSocket":
		{
			var e = interface{}(err).(*ws.Internal)

			newErr = e.Internal
		}
	}

	return
}

// ToRestAPI - преобразование ошибки в ошибку RestAPI.
func ToRestAPI[T Error](err T) (newErr RestAPI) {
	switch reflect.TypeOf(new(T)).String() {
	case "*errors.Error":
		{
			var e = interface{}(err).(*internal.Internal)

			newErr = &rest_api.Internal{
				Internal: e,
			}
		}
	case "*errors.RestAPI":
		{
			var e = interface{}(err).(*rest_api.Internal)

			newErr = e
		}
	case "*errors.WebSocket":
		{
			var e = interface{}(err).(*ws.Internal)

			newErr = &rest_api.Internal{
				Internal: e.Internal,
			}
		}
	}

	return
}

// ToWebSocket - преобразование ошибки в ошибку WebSocket.
func ToWebSocket[T Error](err T) (newErr WebSocket) {
	switch reflect.TypeOf(new(T)).String() {
	case "*errors.Error":
		{
			var e = interface{}(err).(*internal.Internal)

			newErr = &ws.Internal{
				Internal: e,
			}
		}
	case "*errors.RestAPI":
		{
			var e = interface{}(err).(*rest_api.Internal)

			newErr = &ws.Internal{
				Internal: e.Internal,
			}
		}
	case "*errors.WebSocket":
		{
			var e = interface{}(err).(*ws.Internal)

			newErr = e
		}
	}

	return
}
