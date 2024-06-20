package error_list

import (
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/types"
)

var (
	ErrInternalServerError = c_errors.Constructor[c_errors.Error]{
		ID:     "I-000001",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Internal server error. "),
	}.Build()

	ErrFailedToInitializeSystem = c_errors.Constructor[c_errors.Error]{
		ID:     "I-000002",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Failed to initialize the system. "),
	}.Build()

	ErrSystemCleanupError = c_errors.Constructor[c_errors.Error]{
		ID:     "I-000003",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("System cleanup error. "),
	}.Build()
)
