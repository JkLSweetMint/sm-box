package rest_api

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
				Internal: internal.New(internal.Constructor{
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
				Internal: internal.New(internal.Constructor{
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
				Internal: internal.New(internal.Constructor{
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
				Internal:   tt.fields.Internal,
				statusCode: tt.fields.statusCode,
			}
			if gotC := i.StatusCode(); gotC != tt.wantC {
				t.Errorf("StatusCode() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		cnst Constructor
	}

	tests := []struct {
		name  string
		args  args
		wantI *Internal
	}{
		{
			name: "Case 1",
			args: args{
				cnst: Constructor{
					Constructor: internal.Constructor{
						ID:     "T-000001",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error. "),
					},
					StatusCode: 500,
				},
			},
			wantI: &Internal{
				Internal: internal.New(internal.Constructor{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).
						Text("Example error. "),
				}),
				statusCode: 500,
			},
		},
		{
			name: "Case 2",
			args: args{
				cnst: Constructor{
					Constructor: internal.Constructor{
						ID:     "T-000002",
						Type:   types.TypeSystem,
						Status: types.StatusFatal,

						Message: new(messages.TextMessage).
							Text("Example error with details. "),
						Details: new(details.Details).
							Set("key", "value"),
					},
					StatusCode: 500,
				},
			},
			wantI: &Internal{
				Internal: internal.New(internal.Constructor{
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
		},
		{
			name: "Case 3",
			args: args{
				cnst: Constructor{
					Constructor: internal.Constructor{
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
					StatusCode: 500,
				},
			},
			wantI: &Internal{
				Internal: internal.New(internal.Constructor{
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotI := New(tt.args.cnst); !reflect.DeepEqual(gotI, tt.wantI) {
				t.Errorf("New() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}
