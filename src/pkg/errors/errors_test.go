package errors

import (
	"encoding/xml"
	"errors"
	"fmt"
	grpc_codes "google.golang.org/grpc/codes"
	grpc_status "google.golang.org/grpc/status"
	"reflect"
	"sm-box/pkg/errors/entities/details"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/internal/grpc"
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

// Примеры grpc ошибок.
var (
	ExampleGrpcError = Constructor[Grpc]{
		ID:     "T-000001",
		Type:   types.TypeSystem,
		Status: types.StatusFatal,

		Message: new(messages.TextMessage).
			Text("Example error. "),
	}.Grpc(
		GrpcConstructor{
			StatusCode: grpc_codes.Internal,
		},
	).Build()

	ExampleGrpcErrorWithDetails = Constructor[Grpc]{
		ID:     "T-000002",
		Type:   types.TypeSystem,
		Status: types.StatusFatal,

		Message: new(messages.TextMessage).
			Text("Example error with details. "),
		Details: new(details.Details).
			Set("key", "value"),
	}.Grpc(
		GrpcConstructor{
			StatusCode: grpc_codes.Internal,
		},
	).Build()

	ExampleGrpcErrorWithDetailsAndFields = Constructor[Grpc]{
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
	}.Grpc(
		GrpcConstructor{
			StatusCode: grpc_codes.Internal,
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

func TestParseGrpc(t *testing.T) {
	type args struct {
		st *grpc_status.Status
	}

	tests := []struct {
		name     string
		args     args
		wantCErr Grpc
	}{
		{
			name: "Case 1",
			args: args{
				st: ExampleGrpcError().(*grpc.Internal).GRPCStatus(),
			},
			wantCErr: ExampleGrpcError(),
		},
		{
			name: "Case 2",
			args: args{
				st: ExampleGrpcErrorWithDetails().(*grpc.Internal).GRPCStatus(),
			},
			wantCErr: ExampleGrpcErrorWithDetails(),
		},
		{
			name: "Case 3",
			args: args{
				st: ExampleGrpcErrorWithDetailsAndFields().(*grpc.Internal).GRPCStatus(),
			},
			wantCErr: ExampleGrpcErrorWithDetailsAndFields(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCErr := ParseGrpc(tt.args.st); !reflect.DeepEqual(gotCErr, tt.wantCErr) {
				t.Errorf("ParseGrpc() = %v, want %v", gotCErr, tt.wantCErr)
			}
		})
	}
}
