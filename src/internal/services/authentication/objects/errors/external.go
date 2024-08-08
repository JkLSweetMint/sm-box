package srv_errors

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	grpc_codes "google.golang.org/grpc/codes"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/types"
)

// E-100001
var (
	AnUnregisteredTokenWasTransferred = c_errors.Constructor[c_errors.Error]{
		ID:     "AUTH-E-100001",
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

// E-100002
var (
	TokenWasNotTransferred = c_errors.Constructor[c_errors.Error]{
		ID:     "AUTH-E-100002",
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

// E-100003
var (
	AlreadyAuthorized = c_errors.Constructor[c_errors.Error]{
		ID:     "AUTH-E-100003",
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

// E-100004
var (
	Unauthorized = c_errors.Constructor[c_errors.Error]{
		ID:     "AUTH-E-100004",
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

// E-100005
var (
	NotAccess = c_errors.Constructor[c_errors.Error]{
		ID:     "AUTH-E-100005",
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

// E-100006
var (
	ValidityPeriodOfUserTokenHasNotStarted = c_errors.Constructor[c_errors.Error]{
		ID:     "AUTH-E-100006",
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

// E-100007
var (
	InvalidAuthorizationDataWasTransferred = c_errors.Constructor[c_errors.Error]{
		ID:     "AUTH-E-100007",
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

// E-100008
var (
	InvalidDataWasTransmitted = c_errors.Constructor[c_errors.Error]{
		ID:     "AUTH-E-100008",
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

// E-100009
var (
	InvalidToken = c_errors.Constructor[c_errors.Error]{
		ID:     "AUTH-E-100009",
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
