package error_list

import (
	"github.com/gofiber/fiber/v3"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/errors/entities/details"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/types"
	"time"
)

var (
	ErrUnknown_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "I-000000",
		Type:   types.TypeUnknown,
		Status: types.StatusUnknown,

		Message: new(messages.TextMessage).
			Text("Unknown error. "),
		Details: new(details.Details).Set("timestamp", time.Now().Unix()),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusInternalServerError,
	}).Build()

	ErrInternalServerError_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "I-000001",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Internal server error. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusInternalServerError,
	}).Build()

	ErrRequestBodyDataCouldNotBeRead_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "I-000003",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The request body data could not be read. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusBadRequest,
	}).Build()

	ErrResponseCouldNotBeRecorded_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "I-000002",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The response could not be recorded. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusInternalServerError,
	}).Build()

	ErrRouteNotFound_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "E-000100",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The route was not found. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusNotFound,
	}).Build()

	ErrTokenNotFound_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "E-000101",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The token was not found. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusNotFound,
	}).Build()

	ErrUnauthorized_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "E-000102",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Not authorized. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusUnauthorized,
	}).Build()

	ErrAlreadyAuthorized_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "E-000103",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Already authorized. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusBadRequest,
	}).Build()

	ErrUserNotFound_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "E-000104",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The user was not found. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusNotFound,
	}).Build()

	ErrValidityPeriodOfUserTokenHasNotStarted_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "E-000105",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The validity period of the user's token has not started yet. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusNotFound,
	}).Build()

	ErrForbidden_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "E-000106",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("You do not have access to visit this route. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusForbidden,
	}).Build()
)
