package ws

import (
	"reflect"
	"sm-box/pkg/errors/entities/details"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/internal"
	"sm-box/pkg/errors/types"
	"testing"
)

func TestInternal_StatusCode(t *testing.T) {
	type fields struct {
		Internal   *internal.Internal
		statusCode int
	}

	tests := []struct {
		name   string
		fields fields
		wantC  int
	}{
		{
			name: "Case 1",
			fields: fields{
				Internal: internal.New(&internal.Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).
						Text("Example error. "),
				}),
				statusCode: 500,
			},
			wantC: 500,
		},
		{
			name: "Case 2",
			fields: fields{
				Internal: internal.New(&internal.Store{
					ID:     "T-000002",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).
						Text("Example error with details. "),
					Details: new(details.Details).
						Set("key", "value"),
				}),
				statusCode: 500,
			},
			wantC: 500,
		},
		{
			name: "Case 3",
			fields: fields{
				Internal: internal.New(&internal.Store{
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
				}),
				statusCode: 500,
			},
			wantC: 500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Internal{
				Internal: tt.fields.Internal,
			}

			if tt.fields.Internal.Store.Others == nil {
				tt.fields.Internal.Store.Others = new(internal.StoreOthers)
			}

			if tt.fields.Internal.Store.Others.WebSocket == nil {
				tt.fields.Internal.Store.Others.WebSocket = new(internal.WebSocketStore)
			}

			tt.fields.Internal.Store.Others.WebSocket.StatusCode = tt.fields.statusCode

			if gotC := i.StatusCode(); gotC != tt.wantC {
				t.Errorf("StatusCode() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		store *internal.Store
	}

	tests := []struct {
		name  string
		args  args
		wantI *Internal
	}{
		{
			name: "Case 1",
			args: args{
				store: &internal.Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).
						Text("Example error. "),

					Others: &internal.StoreOthers{
						WebSocket: &internal.WebSocketStore{
							StatusCode: 500,
						},
					},
				},
			},
			wantI: &Internal{
				Internal: internal.New(&internal.Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).
						Text("Example error. "),

					Others: &internal.StoreOthers{
						WebSocket: &internal.WebSocketStore{
							StatusCode: 500,
						},
					},
				}),
			},
		},
		{
			name: "Case 2",
			args: args{
				store: &internal.Store{
					ID:     "T-000002",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).
						Text("Example error with details. "),
					Details: new(details.Details).
						Set("key", "value"),

					Others: &internal.StoreOthers{
						WebSocket: &internal.WebSocketStore{
							StatusCode: 500,
						},
					},
				},
			},
			wantI: &Internal{
				Internal: internal.New(&internal.Store{
					ID:     "T-000002",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).
						Text("Example error with details. "),
					Details: new(details.Details).
						Set("key", "value"),

					Others: &internal.StoreOthers{
						WebSocket: &internal.WebSocketStore{
							StatusCode: 500,
						},
					},
				}),
			},
		},
		{
			name: "Case 3",
			args: args{
				store: &internal.Store{
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

					Others: &internal.StoreOthers{
						WebSocket: &internal.WebSocketStore{
							StatusCode: 500,
						},
					},
				},
			},
			wantI: &Internal{
				Internal: internal.New(&internal.Store{
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

					Others: &internal.StoreOthers{
						WebSocket: &internal.WebSocketStore{
							StatusCode: 500,
						},
					},
				}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotI := New(tt.args.store); !reflect.DeepEqual(gotI, tt.wantI) {
				t.Errorf("New() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}
