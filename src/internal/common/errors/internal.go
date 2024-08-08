package common_errors

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	grpc_codes "google.golang.org/grpc/codes"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/errors/entities/details"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/types"
)

// I-000000
var (
	Unknown = c_errors.Constructor[c_errors.Error]{
		ID:     "I-000000",
		Type:   types.TypeUnknown,
		Status: types.StatusUnknown,

		Message: new(messages.TextMessage).
			Text("Unknown error. "),
		Details: new(details.Details).Set("key", "value"),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusInternalServerError,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseInternalServerErr,
	}).Grpc(c_errors.GrpcConstructor{
		StatusCode: grpc_codes.Internal,
	}).Build()
)

// I-000001
var (
	InternalServerError = c_errors.Constructor[c_errors.Error]{
		ID:     "I-000001",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Internal server error. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusInternalServerError,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseInternalServerErr,
	}).Grpc(c_errors.GrpcConstructor{
		StatusCode: grpc_codes.Internal,
	}).Build()
)
