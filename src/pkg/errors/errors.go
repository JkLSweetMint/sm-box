package errors

import (
	grpc_codes "google.golang.org/grpc/codes"
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

	// Grpc - описание grpc ошибки.
	Grpc interface {
		Error

		StatusCode() (c grpc_codes.Code)
	}
)
