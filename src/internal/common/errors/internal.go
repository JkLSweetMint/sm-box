package errors

import (
	"github.com/gofiber/fiber/v3"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/errors/entities/details"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/types"
	"time"
)

// I-000000
var (
	Unknown_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "I-000000",
		Type:   types.TypeUnknown,
		Status: types.StatusUnknown,

		Message: new(messages.TextMessage).
			Text("Unknown error. "),
		Details: new(details.Details).Set("timestamp", time.Now().Unix()),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusInternalServerError,
	}).Build()
)

// I-000001
var (
	InternalServerError_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "I-000001",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Internal server error. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusInternalServerError,
	}).Build()

	InternalServerError = c_errors.Constructor[c_errors.Error]{
		ID:     "I-000001",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Internal server error. "),
	}.Build()
)

// I-000100
var (
	RequestBodyDataCouldNotBeRead_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "I-000100",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The request body data could not be read. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusBadRequest,
	}).Build()
)

// I-000100
var (
	ResponseCouldNotBeRecorded_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "I-000101",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The response could not be recorded. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusInternalServerError,
	}).Build()
)

// I-000002
var (
	FailedToInitializeSystem = c_errors.Constructor[c_errors.Error]{
		ID:     "I-000002",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Failed to initialize the system. "),
	}.Build()
)

// I-000003
var (
	SystemCleanupError = c_errors.Constructor[c_errors.Error]{
		ID:     "I-000003",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("System cleanup error. "),
	}.Build()
)
