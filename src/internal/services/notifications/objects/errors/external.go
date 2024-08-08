package srv_errors

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	grpc_codes "google.golang.org/grpc/codes"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/types"
)

// APP-E-100001
var (
	UserNotificationNotFound = c_errors.Constructor[c_errors.Error]{
		ID:     "NTF-E-100001",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The user notification was not found. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusNotFound,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Grpc(c_errors.GrpcConstructor{
		StatusCode: grpc_codes.NotFound,
	}).Build()
)

// APP-E-100002
var (
	UserNotificationAlreadyRead = c_errors.Constructor[c_errors.Error]{
		ID:     "NTF-E-100002",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The user notification already read. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusBadRequest,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Grpc(c_errors.GrpcConstructor{
		StatusCode: grpc_codes.InvalidArgument,
	}).Build()
)
