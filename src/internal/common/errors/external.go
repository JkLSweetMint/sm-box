package error_list

import (
	"github.com/gofiber/fiber/v3"
	c_errors "sm-box/pkg/errors"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/types"
)

// E-000100
var (
	RouteNotFound_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "E-000100",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The route was not found. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusNotFound,
	}).Build()
)

// E-000101
var (
	TokenNotFound_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "E-000101",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The token was not found. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusNotFound,
	}).Build()

	TokenNotFound = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000101",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The token was not found. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusNotFound,
	}).Build()
)

// E-000102
var (
	Unauthorized_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "E-000102",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Not authorized. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusUnauthorized,
	}).Build()
)

// E-000103
var (
	AlreadyAuthorized_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "E-000103",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Already authorized. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusBadRequest,
	}).Build()

	AlreadyAuthorized = c_errors.Constructor[c_errors.Error]{
		ID:     "E-000103",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Already authorized. "),
	}.Build()
)

// E-000104
var (
	ValidityPeriodOfUserTokenHasNotStarted_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "E-000104",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The validity period of the user's token has not started yet. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusNotFound,
	}).Build()
)

// E-000105
var (
	Forbidden_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "E-000105",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("You do not have access to visit this route. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusForbidden,
	}).Build()
)

// E-000106
var (
	InvalidArgumentsValue_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "E-000106",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Invalid arguments value. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusBadRequest,
	}).Build()
)

// E-000300
var (
	InvalidFlag = c_errors.Constructor[c_errors.Error]{
		ID:     "E-100010",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Invalid flag value. "),
	}.Build()
)

// E-100000
var (
	UserNotFound_RestAPI = c_errors.Constructor[c_errors.RestAPI]{
		ID:     "E-100000",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The user was not found. "),
	}.RestAPI(c_errors.RestAPIConstructor{
		StatusCode: fiber.StatusNotFound,
	}).Build()

	UserNotFound = c_errors.Constructor[c_errors.Error]{
		ID:     "E-100000",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The user was not found. "),
	}.Build()
)

// E-100002
var (
	ProjectOwnerNotFound = c_errors.Constructor[c_errors.Error]{
		ID:     "E-100002",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The project owner was not found. "),
	}.Build()
)

// E-100003
var (
	FailedCreateProject = c_errors.Constructor[c_errors.Error]{
		ID:     "E-100003",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("An error occurred during the creation of the project. "),
	}.Build()
)

// E-100004
var (
	FailedRemoveProject = c_errors.Constructor[c_errors.Error]{
		ID:     "E-100004",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("An error occurred while deleting the project. "),
	}.Build()
)

// E-100005
var (
	ProjectDataCouldNotBeRetrieved = c_errors.Constructor[c_errors.Error]{
		ID:     "E-100005",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("Project data could not be retrieved. "),
	}.Build()
)

// E-100006
var (
	ProjectNotFound = c_errors.Constructor[c_errors.Error]{
		ID:     "E-100006",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("The project was not found. "),
	}.Build()
)

// E-100007
var (
	ReceivingTheProjects = c_errors.Constructor[c_errors.Error]{
		ID:     "E-100007",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("An error occurred while receiving the projects. "),
	}.Build()
)

// E-100008
var (
	FailedSetProjectEnv = c_errors.Constructor[c_errors.Error]{
		ID:     "E-100008",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("An error occurred while deleting the project. "),
	}.Build()
)

// E-100009
var (
	FailedGetProjectEnv = c_errors.Constructor[c_errors.Error]{
		ID:     "E-100009",
		Type:   types.TypeSystem,
		Status: types.StatusError,

		Message: new(messages.TextMessage).
			Text("An error occurred while retrieving the values of the project environment variables. "),
	}.Build()
)
