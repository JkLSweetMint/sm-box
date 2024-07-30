package common_errors

import (
	"github.com/gofiber/fiber/v3"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/types"
)

// IRA-000001
var (
	ResponseCouldNotBeRecorded_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "IRA-000001",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The response could not be recorded. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusInternalServerError,
	}).Build()
)
