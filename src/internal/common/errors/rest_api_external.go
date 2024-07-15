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
		StatusCode: fiber.StatusForbidden,
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
