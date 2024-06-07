package errors

import (
	"sm-box/pkg/errors/helpers"
	"sm-box/pkg/errors/types"
)

// Описание ошибок.
type (
	// Error - описание базовой ошибки.
	Error interface {
		ID() (id types.ID)
		Type() (t types.ErrorType)
		Status() (s types.Status)
		Message() (m string)
		Details() (details types.Details)

		SetError(err error)
		SetMessage(m types.Message)

		helpers.Error
		helpers.Stringer
		helpers.Serialization
	}

	// RestAPI - описание rest api ошибки.
	RestAPI interface {
		Error

		StatusCode() (c int)
	}

	// WebSocket - описание web socket ошибки.
	WebSocket interface {
		Error

		StatusCode() (c int)
	}
)
