package errors

import (
	"encoding/xml"
	"errors"
	"fmt"
	"sm-box/pkg/errors/entities/details"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/types"
	"testing"
)

// Примеры базовой ошибки.
var (
	ExampleError = Constructor[Error]{
		ID:     "T-000001",
		Type:   types.TypeSystem,
		Status: types.StatusFatal,

		Message: new(messages.TextMessage).
			Text("Example error. "),
	}.Build()

	ExampleErrorWithDetails = Constructor[Error]{
		ID:     "T-000002",
		Type:   types.TypeSystem,
		Status: types.StatusFatal,

		Message: new(messages.TextMessage).
			Text("Example error with details. "),
		Details: new(details.Details).
			Set("key", "value"),
	}.Build()

	ExampleErrorWithDetailsAndFields = Constructor[Error]{
		ID:     "T-000003",
		Type:   types.TypeSystem,
		Status: types.StatusFatal,

		Message: new(messages.TextMessage).Text("Example error with details and fields. "),
		Details: new(details.Details).
			Set("key", "value").
			SetFields(types.DetailsField{
				Key:     new(details.FieldKey).Add("test"),
				Message: new(messages.TextMessage).Text("123"),
			}),
	}.Build()
)

// Примеры rest api ошибок.
var (
	ExampleRestAPIError = Constructor[RestAPI]{
		ID:     "T-000001",
		Type:   types.TypeSystem,
		Status: types.StatusFatal,

		Message: new(messages.TextMessage).
			Text("Example error. "),
	}.RestAPI(
		RestAPIConstructor{
			StatusCode: 500,
		},
	).Build()

	ExampleRestAPIErrorWithDetails = Constructor[RestAPI]{
		ID:     "T-000002",
		Type:   types.TypeSystem,
		Status: types.StatusFatal,

		Message: new(messages.TextMessage).
			Text("Example error with details. "),
		Details: new(details.Details).
			Set("key", "value"),
	}.RestAPI(
		RestAPIConstructor{
			StatusCode: 500,
		},
	).Build()

	ExampleRestAPIErrorWithDetailsAndFields = Constructor[RestAPI]{
		ID:     "T-000003",
		Type:   types.TypeSystem,
		Status: types.StatusFatal,

		Message: new(messages.TextMessage).Text("Example error with details and fields. "),
		Details: new(details.Details).
			Set("key", "value").
			SetFields(types.DetailsField{
				Key:     new(details.FieldKey).Add("test"),
				Message: new(messages.TextMessage).Text("123"),
			}),
	}.RestAPI(
		RestAPIConstructor{
			StatusCode: 500,
		},
	).Build()
)

// Примеры web socket ошибок.
var (
	ExampleWebSocketError = Constructor[WebSocket]{
		ID:     "T-000001",
		Type:   types.TypeSystem,
		Status: types.StatusFatal,

		Message: new(messages.TextMessage).
			Text("Example error. "),
	}.WebSocket(
		WebSocketConstructor{
			StatusCode: 1011,
		},
	).Build()

	ExampleWebSocketErrorWithDetails = Constructor[WebSocket]{
		ID:     "T-000002",
		Type:   types.TypeSystem,
		Status: types.StatusFatal,

		Message: new(messages.TextMessage).
			Text("Example error with details. "),
		Details: new(details.Details).
			Set("key", "value"),
	}.WebSocket(
		WebSocketConstructor{
			StatusCode: 1011,
		},
	).Build()

	ExampleWebSocketErrorWithDetailsAndFields = Constructor[WebSocket]{
		ID:     "T-000003",
		Type:   types.TypeSystem,
		Status: types.StatusFatal,

		Message: new(messages.TextMessage).Text("Example error with details and fields. "),
		Details: new(details.Details).
			Set("key", "value").
			SetFields(types.DetailsField{
				Key:     new(details.FieldKey).Add("test"),
				Message: new(messages.TextMessage).Text("123"),
			}),
	}.WebSocket(
		WebSocketConstructor{
			StatusCode: 1011,
		},
	).Build()
)

func Test(t *testing.T) {
	var e = ExampleWebSocketError()
	e.SetError(errors.New("Test error. "))

	if data, err := xml.MarshalIndent(e, "", "\t"); err != nil {
		t.Fatal(err)
	} else {
		fmt.Printf("%s\n", data)
	}
}
