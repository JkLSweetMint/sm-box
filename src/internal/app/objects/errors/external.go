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
	ProjectNotFound = c_errors.Constructor[c_errors.Error]{
		ID:     "APP-E-100001",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The project was not found. "),
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
	ListUserProjectsCouldNotBeRetrieved = c_errors.Constructor[c_errors.Error]{
		ID:     "APP-E-100002",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The list of user projects could not be retrieved. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusInternalServerError,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseInternalServerErr,
	}).Grpc(c_errors.GrpcConstructor{
		StatusCode: grpc_codes.Internal,
	}).Build()
)

// APP-E-100003
var (
	ProjectHasAlreadyBeenSelected = c_errors.Constructor[c_errors.Error]{
		ID:     "APP-E-100003",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The project has already been selected, it is not possible to re-select it. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusBadRequest,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Grpc(c_errors.GrpcConstructor{
		StatusCode: grpc_codes.Canceled,
	}).Build()
)

// APP-E-100004
var (
	NotAccessToProject = c_errors.Constructor[c_errors.Error]{
		ID:     "APP-E-100004",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("There is no access to the project. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusForbidden,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Grpc(c_errors.GrpcConstructor{
		StatusCode: grpc_codes.PermissionDenied,
	}).Build()
)
