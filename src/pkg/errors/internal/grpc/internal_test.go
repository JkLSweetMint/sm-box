package grpc

import (
	grpc_codes "google.golang.org/grpc/codes"
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
		statusCode grpc_codes.Code
	}

	tests := []struct {
		name   string
		fields fields
		wantC  grpc_codes.Code
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
				statusCode: grpc_codes.Internal,
			},
			wantC: grpc_codes.Internal,
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
				statusCode: grpc_codes.Internal,
			},
			wantC: grpc_codes.Internal,
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
				statusCode: grpc_codes.Internal,
			},
			wantC: grpc_codes.Internal,
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

			if tt.fields.Internal.Store.Others.Grpc == nil {
				tt.fields.Internal.Store.Others.Grpc = new(internal.GrpcStore)
			}

			tt.fields.Internal.Store.Others.Grpc.StatusCode = tt.fields.statusCode

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
				},
			},
			wantI: &Internal{
				Internal: internal.New(&internal.Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).
						Text("Example error. "),
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
