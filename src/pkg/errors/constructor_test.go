package errors

import (
	grpc_codes "google.golang.org/grpc/codes"
	"reflect"
	"sm-box/pkg/errors/entities/details"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/types"
	"testing"
)

func TestConstructor_Build_WithError(t *testing.T) {
	type testCase[T Error] struct {
		name string
		c    Constructor[T]
		want T
	}

	tests := []testCase[Error]{
		{
			name: "Case 1",
			c: Constructor[Error]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),
			},
			want: ExampleError(),
		},
		{
			name: "Case 2",
			c: Constructor[Error]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),
			},
			want: ExampleErrorWithDetails(),
		},
		{
			name: "Case 3",
			c: Constructor[Error]{
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
			},
			want: ExampleErrorWithDetailsAndFields(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Build()(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor_Build_WithRestAPI(t *testing.T) {
	type testCase[T RestAPI] struct {
		name string
		c    Constructor[T]
		want T
	}

	tests := []testCase[RestAPI]{
		{
			name: "Case 1",
			c: Constructor[RestAPI]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),
			}.RestAPI(
				RestAPIConstructor{
					StatusCode: 500,
				}),
			want: ExampleRestAPIError(),
		},
		{
			name: "Case 2",
			c: Constructor[RestAPI]{
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
				}),
			want: ExampleRestAPIErrorWithDetails(),
		},
		{
			name: "Case 3",
			c: Constructor[RestAPI]{
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
				}),
			want: ExampleRestAPIErrorWithDetailsAndFields(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Build()(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor_Build_WithWebSocket(t *testing.T) {
	type testCase[T WebSocket] struct {
		name string
		c    Constructor[T]
		want T
	}

	tests := []testCase[WebSocket]{
		{
			name: "Case 1",
			c: Constructor[WebSocket]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),
			}.WebSocket(
				WebSocketConstructor{
					StatusCode: 1011,
				}),
			want: ExampleWebSocketError(),
		},
		{
			name: "Case 2",
			c: Constructor[WebSocket]{
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
				}),
			want: ExampleWebSocketErrorWithDetails(),
		},
		{
			name: "Case 3",
			c: Constructor[WebSocket]{
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
				}),
			want: ExampleWebSocketErrorWithDetailsAndFields(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Build()(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor_Build_WithGrpc(t *testing.T) {
	type testCase[T Grpc] struct {
		name string
		c    Constructor[T]
		want T
	}

	tests := []testCase[Grpc]{
		{
			name: "Case 1",
			c: Constructor[Grpc]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),
			}.Grpc(
				GrpcConstructor{
					StatusCode: grpc_codes.Internal,
				}),
			want: ExampleGrpcError(),
		},
		{
			name: "Case 2",
			c: Constructor[Grpc]{
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
				}),
			want: ExampleGrpcErrorWithDetails(),
		},
		{
			name: "Case 3",
			c: Constructor[Grpc]{
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
				}),
			want: ExampleGrpcErrorWithDetailsAndFields(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Build()(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor_RestAPI_WithError(t *testing.T) {
	type args struct {
		cstr RestAPIConstructor
	}

	type testCase[T Error] struct {
		name string
		c    Constructor[T]
		args args
		want Constructor[T]
	}

	tests := []testCase[Error]{
		{
			name: "Case 1",
			c: Constructor[Error]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),
			},
			args: args{
				cstr: RestAPIConstructor{
					StatusCode: 500,
				},
			},
			want: Constructor[Error]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
		},
		{
			name: "Case 2",
			c: Constructor[Error]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),
			},
			args: args{
				cstr: RestAPIConstructor{
					StatusCode: 500,
				},
			},
			want: Constructor[Error]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
		},
		{
			name: "Case 3",
			c: Constructor[Error]{
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
			},
			args: args{
				cstr: RestAPIConstructor{
					StatusCode: 500,
				},
			},
			want: Constructor[Error]{
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

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.RestAPI(tt.args.cstr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RestAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor_RestAPI_WithRestAPI(t *testing.T) {
	type args struct {
		cstr RestAPIConstructor
	}

	type testCase[T Error] struct {
		name string
		c    Constructor[T]
		args args
		want Constructor[T]
	}

	tests := []testCase[RestAPI]{
		{
			name: "Case 1",
			c: Constructor[RestAPI]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 501,
					},
				},
			},
			args: args{
				cstr: RestAPIConstructor{
					StatusCode: 500,
				},
			},
			want: Constructor[RestAPI]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
		},
		{
			name: "Case 2",
			c: Constructor[RestAPI]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 501,
					},
				},
			},
			args: args{
				cstr: RestAPIConstructor{
					StatusCode: 500,
				},
			},
			want: Constructor[RestAPI]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
		},
		{
			name: "Case 3",
			c: Constructor[RestAPI]{
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

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 501,
					},
				},
			},
			args: args{
				cstr: RestAPIConstructor{
					StatusCode: 500,
				},
			},
			want: Constructor[RestAPI]{
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

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.RestAPI(tt.args.cstr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RestAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor_RestAPI_WithWebSocket(t *testing.T) {
	type args struct {
		cstr RestAPIConstructor
	}

	type testCase[T Error] struct {
		name string
		c    Constructor[T]
		args args
		want Constructor[T]
	}

	tests := []testCase[WebSocket]{
		{
			name: "Case 1",
			c: Constructor[WebSocket]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1000,
					},
				},
			},
			args: args{
				cstr: RestAPIConstructor{
					StatusCode: 500,
				},
			},
			want: Constructor[WebSocket]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
					WebSocket: &WebSocketConstructor{
						StatusCode: 1000,
					},
				},
			},
		},
		{
			name: "Case 2",
			c: Constructor[WebSocket]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1000,
					},
				},
			},
			args: args{
				cstr: RestAPIConstructor{
					StatusCode: 500,
				},
			},
			want: Constructor[WebSocket]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
					WebSocket: &WebSocketConstructor{
						StatusCode: 1000,
					},
				},
			},
		},
		{
			name: "Case 3",
			c: Constructor[WebSocket]{
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

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1000,
					},
				},
			},
			args: args{
				cstr: RestAPIConstructor{
					StatusCode: 500,
				},
			},
			want: Constructor[WebSocket]{
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

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
					WebSocket: &WebSocketConstructor{
						StatusCode: 1000,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.RestAPI(tt.args.cstr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RestAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor_RestAPI_WithGrpc(t *testing.T) {
	type args struct {
		cstr RestAPIConstructor
	}

	type testCase[T Grpc] struct {
		name string
		c    Constructor[T]
		args args
		want Constructor[T]
	}

	tests := []testCase[Grpc]{
		{
			name: "Case 1",
			c: Constructor[Grpc]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),
			},
			args: args{
				cstr: RestAPIConstructor{
					StatusCode: 500,
				},
			},
			want: Constructor[Grpc]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
		},
		{
			name: "Case 2",
			c: Constructor[Grpc]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),
			},
			args: args{
				cstr: RestAPIConstructor{
					StatusCode: 500,
				},
			},
			want: Constructor[Grpc]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
		},
		{
			name: "Case 3",
			c: Constructor[Grpc]{
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
			},
			args: args{
				cstr: RestAPIConstructor{
					StatusCode: 500,
				},
			},
			want: Constructor[Grpc]{
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

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.RestAPI(tt.args.cstr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RestAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor_WebSocket_WithError(t *testing.T) {
	type args struct {
		cstr WebSocketConstructor
	}

	type testCase[T Error] struct {
		name string
		c    Constructor[T]
		args args
		want Constructor[T]
	}

	tests := []testCase[Error]{
		{
			name: "Case 1",
			c: Constructor[Error]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),
			},
			args: args{
				cstr: WebSocketConstructor{
					StatusCode: 1001,
				},
			},
			want: Constructor[Error]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
				},
			},
		},
		{
			name: "Case 2",
			c: Constructor[Error]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),
			},
			args: args{
				cstr: WebSocketConstructor{
					StatusCode: 1001,
				},
			},
			want: Constructor[Error]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
				},
			},
		},
		{
			name: "Case 3",
			c: Constructor[Error]{
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
			},
			args: args{
				cstr: WebSocketConstructor{
					StatusCode: 1001,
				},
			},
			want: Constructor[Error]{
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

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.WebSocket(tt.args.cstr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WebSocket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor_WebSocket_WithRestAPI(t *testing.T) {
	type args struct {
		cstr WebSocketConstructor
	}

	type testCase[T Error] struct {
		name string
		c    Constructor[T]
		args args
		want Constructor[T]
	}

	tests := []testCase[RestAPI]{
		{
			name: "Case 1",
			c: Constructor[RestAPI]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
			args: args{
				cstr: WebSocketConstructor{
					StatusCode: 1001,
				},
			},
			want: Constructor[RestAPI]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
		},
		{
			name: "Case 2",
			c: Constructor[RestAPI]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
			args: args{
				cstr: WebSocketConstructor{
					StatusCode: 1001,
				},
			},
			want: Constructor[RestAPI]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
		},
		{
			name: "Case 3",
			c: Constructor[RestAPI]{
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

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
			args: args{
				cstr: WebSocketConstructor{
					StatusCode: 1001,
				},
			},
			want: Constructor[RestAPI]{
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

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.WebSocket(tt.args.cstr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WebSocket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor_WebSocket_WithWebSocket(t *testing.T) {
	type args struct {
		cstr WebSocketConstructor
	}

	type testCase[T Error] struct {
		name string
		c    Constructor[T]
		args args
		want Constructor[T]
	}

	tests := []testCase[WebSocket]{
		{
			name: "Case 1",
			c: Constructor[WebSocket]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1002,
					},
				},
			},
			args: args{
				cstr: WebSocketConstructor{
					StatusCode: 1001,
				},
			},
			want: Constructor[WebSocket]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
				},
			},
		},
		{
			name: "Case 2",
			c: Constructor[WebSocket]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1002,
					},
				},
			},
			args: args{
				cstr: WebSocketConstructor{
					StatusCode: 1001,
				},
			},
			want: Constructor[WebSocket]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
				},
			},
		},
		{
			name: "Case 3",
			c: Constructor[WebSocket]{
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

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1002,
					},
				},
			},
			args: args{
				cstr: WebSocketConstructor{
					StatusCode: 1001,
				},
			},
			want: Constructor[WebSocket]{
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

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.WebSocket(tt.args.cstr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WebSocket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor_WebSocket_WithGrpc(t *testing.T) {
	type args struct {
		cstr WebSocketConstructor
	}

	type testCase[T Grpc] struct {
		name string
		c    Constructor[T]
		args args
		want Constructor[T]
	}

	tests := []testCase[Grpc]{
		{
			name: "Case 1",
			c: Constructor[Grpc]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),
			},
			args: args{
				cstr: WebSocketConstructor{
					StatusCode: 1001,
				},
			},
			want: Constructor[Grpc]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
				},
			},
		},
		{
			name: "Case 2",
			c: Constructor[Grpc]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),
			},
			args: args{
				cstr: WebSocketConstructor{
					StatusCode: 1001,
				},
			},
			want: Constructor[Grpc]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
				},
			},
		},
		{
			name: "Case 3",
			c: Constructor[Grpc]{
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
			},
			args: args{
				cstr: WebSocketConstructor{
					StatusCode: 1001,
				},
			},
			want: Constructor[Grpc]{
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

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.WebSocket(tt.args.cstr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WebSocket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor_Grpc_WithError(t *testing.T) {
	type args struct {
		cstr GrpcConstructor
	}

	type testCase[T Error] struct {
		name string
		c    Constructor[T]
		args args
		want Constructor[T]
	}

	tests := []testCase[Error]{
		{
			name: "Case 1",
			c: Constructor[Error]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),
			},
			args: args{
				cstr: GrpcConstructor{
					StatusCode: grpc_codes.Internal,
				},
			},
			want: Constructor[Error]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),

				addons: &constructorAddons{
					Grpc: &GrpcConstructor{
						StatusCode: grpc_codes.Internal,
					},
				},
			},
		},
		{
			name: "Case 2",
			c: Constructor[Error]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),
			},
			args: args{
				cstr: GrpcConstructor{
					StatusCode: grpc_codes.Internal,
				},
			},
			want: Constructor[Error]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					Grpc: &GrpcConstructor{
						StatusCode: grpc_codes.Internal,
					},
				},
			},
		},
		{
			name: "Case 3",
			c: Constructor[Error]{
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
			},
			args: args{
				cstr: GrpcConstructor{
					StatusCode: grpc_codes.Internal,
				},
			},
			want: Constructor[Error]{
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

				addons: &constructorAddons{
					Grpc: &GrpcConstructor{
						StatusCode: grpc_codes.Internal,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Grpc(tt.args.cstr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grpc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor_Grpc_WithRestAPI(t *testing.T) {
	type args struct {
		cstr GrpcConstructor
	}

	type testCase[T Error] struct {
		name string
		c    Constructor[T]
		args args
		want Constructor[T]
	}

	tests := []testCase[RestAPI]{
		{
			name: "Case 1",
			c: Constructor[RestAPI]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
			args: args{
				cstr: GrpcConstructor{
					StatusCode: grpc_codes.Internal,
				},
			},
			want: Constructor[RestAPI]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),

				addons: &constructorAddons{
					Grpc: &GrpcConstructor{
						StatusCode: grpc_codes.Internal,
					},
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
		},
		{
			name: "Case 2",
			c: Constructor[RestAPI]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
			args: args{
				cstr: GrpcConstructor{
					StatusCode: grpc_codes.Internal,
				},
			},
			want: Constructor[RestAPI]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					Grpc: &GrpcConstructor{
						StatusCode: grpc_codes.Internal,
					},
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
		},
		{
			name: "Case 3",
			c: Constructor[RestAPI]{
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

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
			args: args{
				cstr: GrpcConstructor{
					StatusCode: grpc_codes.Internal,
				},
			},
			want: Constructor[RestAPI]{
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

				addons: &constructorAddons{
					Grpc: &GrpcConstructor{
						StatusCode: grpc_codes.Internal,
					},
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Grpc(tt.args.cstr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grpc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor_Grpc_WithWebSocket(t *testing.T) {
	type args struct {
		cstr GrpcConstructor
	}

	type testCase[T Error] struct {
		name string
		c    Constructor[T]
		args args
		want Constructor[T]
	}

	tests := []testCase[WebSocket]{
		{
			name: "Case 1",
			c: Constructor[WebSocket]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
				},
			},
			args: args{
				cstr: GrpcConstructor{
					StatusCode: grpc_codes.Internal,
				},
			},
			want: Constructor[WebSocket]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),

				addons: &constructorAddons{
					Grpc: &GrpcConstructor{
						StatusCode: grpc_codes.Internal,
					},
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
				},
			},
		},
		{
			name: "Case 2",
			c: Constructor[WebSocket]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
				},
			},
			args: args{
				cstr: GrpcConstructor{
					StatusCode: grpc_codes.Internal,
				},
			},
			want: Constructor[WebSocket]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					Grpc: &GrpcConstructor{
						StatusCode: grpc_codes.Internal,
					},
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
				},
			},
		},
		{
			name: "Case 3",
			c: Constructor[WebSocket]{
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

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
				},
			},
			args: args{
				cstr: GrpcConstructor{
					StatusCode: grpc_codes.Internal,
				},
			},
			want: Constructor[WebSocket]{
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

				addons: &constructorAddons{
					Grpc: &GrpcConstructor{
						StatusCode: grpc_codes.Internal,
					},
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Grpc(tt.args.cstr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grpc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor_Grpc_WithGrpc(t *testing.T) {
	type args struct {
		cstr GrpcConstructor
	}

	type testCase[T Grpc] struct {
		name string
		c    Constructor[T]
		args args
		want Constructor[T]
	}

	tests := []testCase[Grpc]{
		{
			name: "Case 1",
			c: Constructor[Grpc]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),

				addons: &constructorAddons{
					Grpc: &GrpcConstructor{
						StatusCode: grpc_codes.Unknown,
					},
				},
			},
			args: args{
				cstr: GrpcConstructor{
					StatusCode: grpc_codes.Internal,
				},
			},
			want: Constructor[Grpc]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),

				addons: &constructorAddons{
					Grpc: &GrpcConstructor{
						StatusCode: grpc_codes.Internal,
					},
				},
			},
		},
		{
			name: "Case 2",
			c: Constructor[Grpc]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					Grpc: &GrpcConstructor{
						StatusCode: grpc_codes.Unknown,
					},
				},
			},
			args: args{
				cstr: GrpcConstructor{
					StatusCode: grpc_codes.Internal,
				},
			},
			want: Constructor[Grpc]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					Grpc: &GrpcConstructor{
						StatusCode: grpc_codes.Internal,
					},
				},
			},
		},
		{
			name: "Case 3",
			c: Constructor[Grpc]{
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

				addons: &constructorAddons{
					Grpc: &GrpcConstructor{
						StatusCode: grpc_codes.Unknown,
					},
				},
			},
			args: args{
				cstr: GrpcConstructor{
					StatusCode: grpc_codes.Internal,
				},
			},
			want: Constructor[Grpc]{
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

				addons: &constructorAddons{
					Grpc: &GrpcConstructor{
						StatusCode: grpc_codes.Internal,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Grpc(tt.args.cstr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Grpc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor_fillEmptyField_WithError(t *testing.T) {
	type testCase[T Error] struct {
		name string
		c    Constructor[T]
		want *Constructor[T]
	}

	tests := []testCase[Error]{
		{
			name: "Case 1",
			c: Constructor[Error]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),
			},
			want: &Constructor[Error]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),
				Details: new(details.Details),

				addons: new(constructorAddons),
			},
		},
		{
			name: "Case 2",
			c: Constructor[Error]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),
			},
			want: &Constructor[Error]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: new(constructorAddons),
			},
		},
		{
			name: "Case 3",
			c: Constructor[Error]{
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
			},
			want: &Constructor[Error]{
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

				addons: new(constructorAddons),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.fillEmptyField(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fillEmptyField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor_fillEmptyField_WithRestAPI(t *testing.T) {
	type testCase[T Error] struct {
		name string
		c    Constructor[T]
		want *Constructor[T]
	}

	tests := []testCase[RestAPI]{
		{
			name: "Case 1",
			c: Constructor[RestAPI]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
			want: &Constructor[RestAPI]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),
				Details: new(details.Details),

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
		},
		{
			name: "Case 2",
			c: Constructor[RestAPI]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
			want: &Constructor[RestAPI]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
		},
		{
			name: "Case 3",
			c: Constructor[RestAPI]{
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

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
			want: &Constructor[RestAPI]{
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

				addons: &constructorAddons{
					RestAPI: &RestAPIConstructor{
						StatusCode: 500,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.fillEmptyField(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fillEmptyField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor_fillEmptyField_WithWebSocket(t *testing.T) {
	type testCase[T Error] struct {
		name string
		c    Constructor[T]
		want *Constructor[T]
	}

	tests := []testCase[WebSocket]{
		{
			name: "Case 1",
			c: Constructor[WebSocket]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
				},
			},
			want: &Constructor[WebSocket]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),
				Details: new(details.Details),

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
				},
			},
		},
		{
			name: "Case 2",
			c: Constructor[WebSocket]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
				},
			},
			want: &Constructor[WebSocket]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
				},
			},
		},
		{
			name: "Case 3",
			c: Constructor[WebSocket]{
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

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
				},
			},
			want: &Constructor[WebSocket]{
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

				addons: &constructorAddons{
					WebSocket: &WebSocketConstructor{
						StatusCode: 1001,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.fillEmptyField(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fillEmptyField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor_fillEmptyField_WithGrpc(t *testing.T) {
	type testCase[T Grpc] struct {
		name string
		c    Constructor[T]
		want *Constructor[T]
	}

	tests := []testCase[Grpc]{
		{
			name: "Case 1",
			c: Constructor[Grpc]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),
			},
			want: &Constructor[Grpc]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),
				Details: new(details.Details),

				addons: new(constructorAddons),
			},
		},
		{
			name: "Case 2",
			c: Constructor[Grpc]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),
			},
			want: &Constructor[Grpc]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),

				addons: new(constructorAddons),
			},
		},
		{
			name: "Case 3",
			c: Constructor[Grpc]{
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
			},
			want: &Constructor[Grpc]{
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

				addons: new(constructorAddons),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.fillEmptyField(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fillEmptyField() = %v, want %v", got, tt.want)
			}
		})
	}
}
