package error_list

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v3"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/errors/entities/details"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/types"
)

// E-000001
var (
	TokenNotFound = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000001",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The token was not found. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusNotFound,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Build()
)

// E-000002
var (
	AlreadyAuthorized = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000002",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Already authorized. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusBadRequest,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Build()
)

// E-000003
var (
	Unauthorized = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000003",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Not authorized. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusUnauthorized,
	}).Build()
)

// E-000004
var (
	ValidityPeriodOfUserTokenHasNotStarted = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000004",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The validity period of the user's token has not started yet. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusNotFound,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Build()
)

// E-000005
var (
	UserNotFound = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000005",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The user was not found. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusNotFound,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Build()
)

// E-000006
var (
	ListUserProjectsCouldNotBeRetrieved = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000006",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The list of user projects could not be retrieved. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusInternalServerError,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseInternalServerErr,
	}).Build()
)

// E-000007
var (
	ProjectHasAlreadyBeenSelected = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000007",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The project has already been selected, it is not possible to re-select it. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusBadRequest,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Build()
)

// E-000008
var (
	ProjectNotFound = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000008",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The project was not found. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusNotFound,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Build()
)

// E-000009
var (
	NotAccessToProject = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000009",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("There is no access to the project. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusBadRequest,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Build()
)

// E-000010
var (
	InvalidTextLocalizationPaths = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000010",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Invalid value of text localization paths. "),

		Details: new(details.Details).Set("paths", new(messages.TextMessage).Text("Invalid value. ").String()),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusBadRequest,
	}).WebSocket(c_errors.WebSocketConstructor{
		StatusCode: websocket.CloseNormalClosure,
	}).Build()
)
