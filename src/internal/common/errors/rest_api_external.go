package error_list

import (
	"github.com/gofiber/fiber/v3"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/types"
)

// ERA-000001
var (
	RouteNotFound_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "ERA-000001",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The route was not found. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusNotFound,
	}).Build()
)

// ERA-000002
var (
	RequestBodyDataCouldNotBeRead_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "ERA-000002",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The request body data could not be read. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusBadRequest,
	}).Build()
)

// ERA-000003
var (
	Unauthorized_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "ERA-000003",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Not authorized. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusUnauthorized,
	}).Build()
)

// ERA-000004
var (
	TokenHasNotBeenTransferred_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "ERA-000004",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The token has not been transferred. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusNotFound,
	}).Build()
)

// ERA-000005
var (
	Forbidden_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "ERA-000005",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("You do not have access to visit this route. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusForbidden,
	}).Build()
)
