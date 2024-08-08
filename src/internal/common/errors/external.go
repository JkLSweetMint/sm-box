package common_errors

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	grpc_codes "google.golang.org/grpc/codes"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/types"
)

// E-000001
var (
	InvalidSortValue = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000001",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Invalid sorting value. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusBadRequest,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Grpc(c_errors.GrpcConstructor{
		StatusCode: grpc_codes.InvalidArgument,
	}).Build()
)

// E-000002
var (
	InvalidFilterValue = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000002",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Invalid filter value. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusBadRequest,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Grpc(c_errors.GrpcConstructor{
		StatusCode: grpc_codes.InvalidArgument,
	}).Build()
)

// E-000003
var (
	InvalidArguments = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000003",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Invalid arguments. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusBadRequest,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Grpc(c_errors.GrpcConstructor{
		StatusCode: grpc_codes.InvalidArgument,
	}).Build()
)
