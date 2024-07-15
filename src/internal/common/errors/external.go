package error_list

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v3"
	grpc_codes "google.golang.org/grpc/codes"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/errors/entities/details"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/types"
)

// E-000001
var (
	AnUnregisteredTokenWasTransferred = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000001",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("An unregistered token was transferred. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusForbidden,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Grpc(c_errors.GrpcConstructor{
		StatusCode: grpc_codes.PermissionDenied,
	}).Build()
)

// E-000002
var (
	TokenWasNotTransferred = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000002",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The token was not transferred. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusForbidden,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Grpc(c_errors.GrpcConstructor{
		StatusCode: grpc_codes.PermissionDenied,
	}).Build()
)

// E-000003
var (
	AlreadyAuthorized = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000003",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Already authorized. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusBadRequest,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Grpc(c_errors.GrpcConstructor{
		StatusCode: grpc_codes.Canceled,
	}).Build()
)

// E-000004
var (
	Unauthorized = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000004",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Not authorized. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusUnauthorized,
	}).Grpc(c_errors.GrpcConstructor{
		StatusCode: grpc_codes.Unauthenticated,
	}).Build()
)

// E-000005
var (
	NotAccess = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000005",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Not access. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusForbidden,
	}).Grpc(c_errors.GrpcConstructor{
		StatusCode: grpc_codes.Unauthenticated,
	}).Build()
)

// E-000006
var (
	ValidityPeriodOfUserTokenHasNotStarted = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000006",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The validity period of the user's token has not started yet. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusNotFound,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Grpc(c_errors.GrpcConstructor{
		StatusCode: grpc_codes.NotFound,
	}).Build()
)

// E-000007
var (
	UserNotFound = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000007",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The user was not found. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusNotFound,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Grpc(c_errors.GrpcConstructor{
		StatusCode: grpc_codes.NotFound,
	}).Build()
)

// E-000008
var (
	InvalidAuthorizationDataWasTransferred = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000008",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Invalid authorization data was transferred. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusBadRequest,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Grpc(c_errors.GrpcConstructor{
		StatusCode: grpc_codes.Canceled,
	}).Build()
)

// E-000009
var (
	InvalidDataWasTransmitted = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000009",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Invalid data was transmitted. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusBadRequest,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Grpc(c_errors.GrpcConstructor{
		StatusCode: grpc_codes.InvalidArgument,
	}).Build()
)

// E-000010
var (
	InvalidToken = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000010",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("An invalid token was transferred. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusForbidden,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Grpc(c_errors.GrpcConstructor{
		StatusCode: grpc_codes.PermissionDenied,
	}).Build()
)

// E-010001
var (
	ListUserProjectsCouldNotBeRetrieved = c_errors.Constructor[c_errors.Error]{
		ID:     "E-010001",
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

// E-010002
var (
	ProjectHasAlreadyBeenSelected = c_errors.Constructor[c_errors.Error]{
		ID:     "E-010002",
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

// E-010003
var (
	ProjectNotFound = c_errors.Constructor[c_errors.Error]{
		ID:     "E-010003",
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

// E-010004
var (
	NotAccessToProject = c_errors.Constructor[c_errors.Error]{
		ID:     "E-010004",
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

// E-020001
var (
	InvalidTextLocalizationPaths = c_errors.Constructor[c_errors.Error]{
		ID:     "E-020001",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Invalid value of text localization paths. "),

		Details: new(details.Details).Set("paths", new(messages.TextMessage).Text("Invalid value. ").String()),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusBadRequest,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Grpc(c_errors.GrpcConstructor{
		StatusCode: grpc_codes.InvalidArgument,
	}).Build()
)
