package internal

import (
	"context"
	"errors"
	"reflect"
	"sm-box/pkg/errors/entities/details"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/types"
	"testing"
)

func Test_Internal_Details(t *testing.T) {
	type fields struct {
		id     types.ID
		t      types.ErrorType
		status types.Status

		err     error
		message types.Message
		details types.Details

		ctx context.Context
	}

	tests := []struct {
		name        string
		fields      fields
		wantDetails types.Details
	}{
		{
			name: "Case 1",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details),

				ctx: context.Background(),
			},
			wantDetails: new(details.Details),
		},
		{
			name: "Case 2",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details).
					Set("key", "value"),

				ctx: context.Background(),
			},
			wantDetails: new(details.Details).
				Set("key", "value"),
		},
		{
			name: "Case 3",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details).
					Set("key", "value").
					SetFields(types.DetailsField{
						Key:     new(details.FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					}),

				ctx: context.Background(),
			},
			wantDetails: new(details.Details).
				Set("key", "value").
				SetFields(types.DetailsField{
					Key:     new(details.FieldKey).Add("test"),
					Message: new(messages.TextMessage).Text("123"),
				}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Internal{
				Store: &Store{
					ID:     tt.fields.id,
					Type:   tt.fields.t,
					Status: tt.fields.status,

					Err:     tt.fields.err,
					Message: tt.fields.message,
					Details: tt.fields.details,
				},
				ctx: tt.fields.ctx,
			}

			if gotM := i.Details(); !reflect.DeepEqual(gotM, tt.wantDetails) {
				t.Errorf("Details() = %v, want %v", gotM, tt.wantDetails)
			}
		})
	}
}

func Test_Internal_Error(t *testing.T) {
	type fields struct {
		id      types.ID
		t       types.ErrorType
		status  types.Status
		err     error
		message types.Message
		details types.Details
		ctx     context.Context
	}

	tests := []struct {
		name   string
		fields fields
		wantS  string
	}{
		{
			name: "Case 1",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details),

				ctx: context.Background(),
			},
			wantS: "Test error. ",
		},
		{
			name: "Case 2",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details).
					Set("key", "value"),

				ctx: context.Background(),
			},
			wantS: "Test error. ",
		},
		{
			name: "Case 3",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details).
					Set("key", "value").
					SetFields(types.DetailsField{
						Key:     new(details.FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					}),

				ctx: context.Background(),
			},
			wantS: "Test error. ",
		},
		{
			name: "Case 4",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: errors.New("Test. "),
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details),

				ctx: context.Background(),
			},
			wantS: "Test. ",
		},
		{
			name: "Case 5",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: errors.New("Test. "),
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details).
					Set("key", "value"),

				ctx: context.Background(),
			},
			wantS: "Test. ",
		},
		{
			name: "Case 6",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: errors.New("Test. "),
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details).
					Set("key", "value").
					SetFields(types.DetailsField{
						Key:     new(details.FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					}),

				ctx: context.Background(),
			},
			wantS: "Test. ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Internal{
				Store: &Store{
					ID:     tt.fields.id,
					Type:   tt.fields.t,
					Status: tt.fields.status,

					Err:     tt.fields.err,
					Message: tt.fields.message,
					Details: tt.fields.details,
				},
				ctx: tt.fields.ctx,
			}

			if gotS := i.Error(); gotS != tt.wantS {
				t.Errorf("Error() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func Test_Internal_ID(t *testing.T) {
	type fields struct {
		id      types.ID
		t       types.ErrorType
		status  types.Status
		err     error
		message types.Message
		details types.Details
		ctx     context.Context
	}

	tests := []struct {
		name   string
		fields fields
		wantId types.ID
	}{
		{
			name: "Case 1",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details),

				ctx: context.Background(),
			},
			wantId: "T-000001",
		},
		{
			name: "Case 2",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details).
					Set("key", "value"),

				ctx: context.Background(),
			},
			wantId: "T-000001",
		},
		{
			name: "Case 3",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details).
					Set("key", "value").
					SetFields(types.DetailsField{
						Key:     new(details.FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					}),

				ctx: context.Background(),
			},
			wantId: "T-000001",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Internal{
				Store: &Store{
					ID:     tt.fields.id,
					Type:   tt.fields.t,
					Status: tt.fields.status,

					Err:     tt.fields.err,
					Message: tt.fields.message,
					Details: tt.fields.details,
				},
				ctx: tt.fields.ctx,
			}

			if gotId := i.ID(); gotId != tt.wantId {
				t.Errorf("ID() = %v, want %v", gotId, tt.wantId)
			}
		})
	}
}

func Test_Internal_Message(t *testing.T) {
	type fields struct {
		id      types.ID
		t       types.ErrorType
		status  types.Status
		err     error
		message types.Message
		details types.Details
		ctx     context.Context
	}

	tests := []struct {
		name   string
		fields fields
		wantM  string
	}{
		{
			name: "Case 1",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details),

				ctx: context.Background(),
			},
			wantM: "Test error. ",
		},
		{
			name: "Case 2",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details).
					Set("key", "value"),

				ctx: context.Background(),
			},
			wantM: "Test error. ",
		},
		{
			name: "Case 3",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details).
					Set("key", "value").
					SetFields(types.DetailsField{
						Key:     new(details.FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					}),

				ctx: context.Background(),
			},
			wantM: "Test error. ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Internal{
				Store: &Store{
					ID:     tt.fields.id,
					Type:   tt.fields.t,
					Status: tt.fields.status,

					Err:     tt.fields.err,
					Message: tt.fields.message,
					Details: tt.fields.details,
				},
				ctx: tt.fields.ctx,
			}

			if gotM := i.Message(); gotM != tt.wantM {
				t.Errorf("Message() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

func Test_Internal_SetError(t *testing.T) {
	type fields struct {
		id      types.ID
		t       types.ErrorType
		status  types.Status
		err     error
		message types.Message
		details types.Details
		ctx     context.Context
	}

	type args struct {
		err error
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Internal
	}{
		{
			name: "Case 1",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details),

				ctx: context.Background(),
			},
			args: args{
				err: errors.New("Test 1. "),
			},
			want: &Internal{
				Store: &Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Err: errors.New("Test 1. "),
					Message: new(messages.TextMessage).
						Text("Test error. "),
					Details: new(details.Details),
				},

				ctx: context.Background(),
			},
		},
		{
			name: "Case 2",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: errors.New("Test 1. "),
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details),

				ctx: context.Background(),
			},
			args: args{
				err: errors.New("Test 2. "),
			},
			want: &Internal{
				Store: &Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Err: errors.New("Test 2. "),
					Message: new(messages.TextMessage).
						Text("Test error. "),
					Details: new(details.Details),
				},

				ctx: context.Background(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Internal{
				Store: &Store{
					ID:     tt.fields.id,
					Type:   tt.fields.t,
					Status: tt.fields.status,

					Err:     tt.fields.err,
					Message: tt.fields.message,
					Details: tt.fields.details,
				},
				ctx: tt.fields.ctx,
			}

			i.SetError(tt.args.err)

			if !reflect.DeepEqual(i, tt.want) {
				t.Errorf("SetError() = %v, want %v", i, tt.want)
			}
		})
	}
}

func Test_Internal_SetMessage(t *testing.T) {
	type fields struct {
		id      types.ID
		t       types.ErrorType
		status  types.Status
		err     error
		message types.Message
		details types.Details
		ctx     context.Context
	}

	type args struct {
		m types.Message
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Internal
	}{
		{
			name: "Case 1",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err:     errors.New("Test. "),
				message: nil,
				details: new(details.Details),

				ctx: context.Background(),
			},
			args: args{
				m: new(messages.TextMessage).
					Text("Test error. "),
			},
			want: &Internal{
				Store: &Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Err: errors.New("Test. "),
					Message: new(messages.TextMessage).
						Text("Test error. "),
					Details: new(details.Details),
				},

				ctx: context.Background(),
			},
		},
		{
			name: "Case 2",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: errors.New("Test. "),
				message: new(messages.TextMessage).
					Text("Test message 1. "),
				details: new(details.Details),

				ctx: context.Background(),
			},
			args: args{
				m: new(messages.TextMessage).
					Text("Test message 2. "),
			},
			want: &Internal{
				Store: &Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Err: errors.New("Test. "),
					Message: new(messages.TextMessage).
						Text("Test message 2. "),
					Details: new(details.Details),
				},

				ctx: context.Background(),
			},
		},
		{
			name: "Case 3",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err:     nil,
				message: nil,
				details: new(details.Details),

				ctx: context.Background(),
			},
			args: args{
				m: new(messages.TextMessage).
					Text("Test error. "),
			},
			want: &Internal{
				Store: &Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Err: nil,
					Message: new(messages.TextMessage).
						Text("Test error. "),
					Details: new(details.Details),
				},

				ctx: context.Background(),
			},
		},
		{
			name: "Case 4",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test message 1. "),
				details: new(details.Details),

				ctx: context.Background(),
			},
			args: args{
				m: new(messages.TextMessage).
					Text("Test message 2. "),
			},
			want: &Internal{
				Store: &Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Err: nil,
					Message: new(messages.TextMessage).
						Text("Test message 2. "),
					Details: new(details.Details),
				},

				ctx: context.Background(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Internal{
				Store: &Store{
					ID:     tt.fields.id,
					Type:   tt.fields.t,
					Status: tt.fields.status,

					Err:     tt.fields.err,
					Message: tt.fields.message,
					Details: tt.fields.details,
				},
				ctx: tt.fields.ctx,
			}

			i.SetMessage(tt.args.m)

			if !reflect.DeepEqual(i, tt.want) {
				t.Errorf("SetMessage() = %v, want %v", i, tt.want)
			}
		})
	}
}

func Test_Internal_Status(t *testing.T) {
	type fields struct {
		id      types.ID
		t       types.ErrorType
		status  types.Status
		err     error
		message types.Message
		details types.Details
		ctx     context.Context
	}

	tests := []struct {
		name   string
		fields fields
		wantS  types.Status
	}{
		{
			name: "Case 1",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details),

				ctx: context.Background(),
			},
			wantS: types.StatusFatal,
		},
		{
			name: "Case 2",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusUnknown,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details).
					Set("key", "value"),

				ctx: context.Background(),
			},
			wantS: types.StatusUnknown,
		},
		{
			name: "Case 3",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFailed,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details).
					Set("key", "value").
					SetFields(types.DetailsField{
						Key:     new(details.FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					}),

				ctx: context.Background(),
			},
			wantS: types.StatusFailed,
		},
		{
			name: "Case 4",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusError,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details).
					Set("key", "value").
					SetFields(types.DetailsField{
						Key:     new(details.FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					}),

				ctx: context.Background(),
			},
			wantS: types.StatusError,
		},
		{
			name: "Case 5",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: -1,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details).
					Set("key", "value").
					SetFields(types.DetailsField{
						Key:     new(details.FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					}),

				ctx: context.Background(),
			},
			wantS: types.StatusUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Internal{
				Store: &Store{
					ID:     tt.fields.id,
					Type:   tt.fields.t,
					Status: tt.fields.status,

					Err:     tt.fields.err,
					Message: tt.fields.message,
					Details: tt.fields.details,
				},
				ctx: tt.fields.ctx,
			}

			if gotS := i.Status(); gotS != tt.wantS {
				t.Errorf("Status() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func Test_Internal_String(t *testing.T) {
	type fields struct {
		id      types.ID
		t       types.ErrorType
		status  types.Status
		err     error
		message types.Message
		details types.Details
		ctx     context.Context
	}

	tests := []struct {
		name   string
		fields fields
		wantS  string
	}{
		{
			name: "Case 1",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details),

				ctx: context.Background(),
			},
			wantS: "Test error. ",
		},
		{
			name: "Case 2",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: errors.New("Test. "),
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details).
					Set("key", "value"),

				ctx: context.Background(),
			},
			wantS: "Test error. : 'Test. '. ",
		},
		{
			name: "Case 3",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err:     errors.New("Test error. "),
				message: nil,
				details: new(details.Details).
					Set("key", "value").
					SetFields(types.DetailsField{
						Key:     new(details.FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					}),

				ctx: context.Background(),
			},
			wantS: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Internal{
				Store: &Store{
					ID:     tt.fields.id,
					Type:   tt.fields.t,
					Status: tt.fields.status,

					Err:     tt.fields.err,
					Message: tt.fields.message,
					Details: tt.fields.details,
				},
				ctx: tt.fields.ctx,
			}

			if gotS := i.String(); gotS != tt.wantS {
				t.Errorf("String() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func Test_Internal_Type(t *testing.T) {
	type fields struct {
		id      types.ID
		t       types.ErrorType
		status  types.Status
		err     error
		message types.Message
		details types.Details
		ctx     context.Context
	}

	tests := []struct {
		name   string
		fields fields
		wantT  types.ErrorType
	}{
		{
			name: "Case 1",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details),

				ctx: context.Background(),
			},
			wantT: types.TypeSystem,
		},
		{
			name: "Case 2",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeUnknown,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details).
					Set("key", "value"),

				ctx: context.Background(),
			},
			wantT: types.TypeUnknown,
		},
		{
			name: "Case 3",
			fields: fields{
				id:     "T-000001",
				t:      -1,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Test error. "),
				details: new(details.Details).
					Set("key", "value").
					SetFields(types.DetailsField{
						Key:     new(details.FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					}),

				ctx: context.Background(),
			},
			wantT: types.TypeUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Internal{
				Store: &Store{
					ID:     tt.fields.id,
					Type:   tt.fields.t,
					Status: tt.fields.status,

					Err:     tt.fields.err,
					Message: tt.fields.message,
					Details: tt.fields.details,
				},
				ctx: tt.fields.ctx,
			}

			if gotT := i.Type(); gotT != tt.wantT {
				t.Errorf("Type() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}
