package internal

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"reflect"
	"sm-box/pkg/errors/entities/details"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/types"
	"testing"
)

func Test_Internal_MarshalJSON(t *testing.T) {
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
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "Case 1",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Message. "),
				details: new(details.Details),

				ctx: context.Background(),
			},
			want:    []byte{123, 34, 105, 100, 34, 58, 34, 84, 45, 48, 48, 48, 48, 48, 49, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 115, 121, 115, 116, 101, 109, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 116, 97, 108, 34, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 77, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 100, 101, 116, 97, 105, 108, 115, 34, 58, 123, 125, 125},
			wantErr: false,
		},
		{
			name: "Case 2",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: errors.New("Test. "),
				message: new(messages.TextMessage).
					Text("Message. "),
				details: new(details.Details),

				ctx: context.Background(),
			},
			want:    []byte{123, 34, 105, 100, 34, 58, 34, 84, 45, 48, 48, 48, 48, 48, 49, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 115, 121, 115, 116, 101, 109, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 116, 97, 108, 34, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 77, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 100, 101, 116, 97, 105, 108, 115, 34, 58, 123, 125, 125},
			wantErr: false,
		},
		{
			name: "Case 3",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Message. "),
				details: new(details.Details).
					Set("key", "value"),

				ctx: context.Background(),
			},
			want:    []byte{123, 34, 105, 100, 34, 58, 34, 84, 45, 48, 48, 48, 48, 48, 49, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 115, 121, 115, 116, 101, 109, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 116, 97, 108, 34, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 77, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 100, 101, 116, 97, 105, 108, 115, 34, 58, 123, 34, 107, 101, 121, 34, 58, 34, 118, 97, 108, 117, 101, 34, 125, 125},
			wantErr: false,
		},
		{
			name: "Case 4",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: errors.New("Test. "),
				message: new(messages.TextMessage).
					Text("Message. "),
				details: new(details.Details).
					Set("key", "value"),

				ctx: context.Background(),
			},
			want:    []byte{123, 34, 105, 100, 34, 58, 34, 84, 45, 48, 48, 48, 48, 48, 49, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 115, 121, 115, 116, 101, 109, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 116, 97, 108, 34, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 77, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 100, 101, 116, 97, 105, 108, 115, 34, 58, 123, 34, 107, 101, 121, 34, 58, 34, 118, 97, 108, 117, 101, 34, 125, 125},
			wantErr: false,
		},
		{
			name: "Case 5",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Message. "),
				details: new(details.Details).
					Set("key", "value"),

				ctx: context.Background(),
			},
			want:    []byte{123, 34, 105, 100, 34, 58, 34, 84, 45, 48, 48, 48, 48, 48, 49, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 115, 121, 115, 116, 101, 109, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 116, 97, 108, 34, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 77, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 100, 101, 116, 97, 105, 108, 115, 34, 58, 123, 34, 107, 101, 121, 34, 58, 34, 118, 97, 108, 117, 101, 34, 125, 125},
			wantErr: false,
		},
		{
			name: "Case 6",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: errors.New("Test. "),
				message: new(messages.TextMessage).
					Text("Message. "),
				details: new(details.Details).
					Set("key", "value"),

				ctx: context.Background(),
			},
			want:    []byte{123, 34, 105, 100, 34, 58, 34, 84, 45, 48, 48, 48, 48, 48, 49, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 115, 121, 115, 116, 101, 109, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 116, 97, 108, 34, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 77, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 100, 101, 116, 97, 105, 108, 115, 34, 58, 123, 34, 107, 101, 121, 34, 58, 34, 118, 97, 108, 117, 101, 34, 125, 125},
			wantErr: false,
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

			got, err := json.Marshal(i)

			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Internal_MarshalXML(t *testing.T) {
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
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "Case 1",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Message. "),
				details: new(details.Details),

				ctx: context.Background(),
			},
			want:    []byte{60, 69, 114, 114, 111, 114, 32, 105, 100, 61, 34, 84, 45, 48, 48, 48, 48, 48, 49, 34, 32, 116, 121, 112, 101, 61, 34, 115, 121, 115, 116, 101, 109, 34, 32, 115, 116, 97, 116, 117, 115, 61, 34, 102, 97, 116, 97, 108, 34, 62, 60, 77, 101, 115, 115, 97, 103, 101, 62, 77, 101, 115, 115, 97, 103, 101, 46, 32, 60, 47, 77, 101, 115, 115, 97, 103, 101, 62, 60, 68, 101, 116, 97, 105, 108, 115, 62, 60, 47, 68, 101, 116, 97, 105, 108, 115, 62, 60, 47, 69, 114, 114, 111, 114, 62},
			wantErr: false,
		},
		{
			name: "Case 2",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: errors.New("Test. "),
				message: new(messages.TextMessage).
					Text("Message. "),
				details: new(details.Details),

				ctx: context.Background(),
			},
			want:    []byte{60, 69, 114, 114, 111, 114, 32, 105, 100, 61, 34, 84, 45, 48, 48, 48, 48, 48, 49, 34, 32, 116, 121, 112, 101, 61, 34, 115, 121, 115, 116, 101, 109, 34, 32, 115, 116, 97, 116, 117, 115, 61, 34, 102, 97, 116, 97, 108, 34, 62, 60, 77, 101, 115, 115, 97, 103, 101, 62, 77, 101, 115, 115, 97, 103, 101, 46, 32, 60, 47, 77, 101, 115, 115, 97, 103, 101, 62, 60, 68, 101, 116, 97, 105, 108, 115, 62, 60, 47, 68, 101, 116, 97, 105, 108, 115, 62, 60, 47, 69, 114, 114, 111, 114, 62},
			wantErr: false,
		},
		{
			name: "Case 3",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Message. "),
				details: new(details.Details).
					Set("key", "value"),

				ctx: context.Background(),
			},
			want:    []byte{60, 69, 114, 114, 111, 114, 32, 105, 100, 61, 34, 84, 45, 48, 48, 48, 48, 48, 49, 34, 32, 116, 121, 112, 101, 61, 34, 115, 121, 115, 116, 101, 109, 34, 32, 115, 116, 97, 116, 117, 115, 61, 34, 102, 97, 116, 97, 108, 34, 62, 60, 77, 101, 115, 115, 97, 103, 101, 62, 77, 101, 115, 115, 97, 103, 101, 46, 32, 60, 47, 77, 101, 115, 115, 97, 103, 101, 62, 60, 68, 101, 116, 97, 105, 108, 115, 62, 60, 73, 116, 101, 109, 32, 107, 101, 121, 61, 34, 107, 101, 121, 34, 62, 118, 97, 108, 117, 101, 60, 47, 73, 116, 101, 109, 62, 60, 47, 68, 101, 116, 97, 105, 108, 115, 62, 60, 47, 69, 114, 114, 111, 114, 62},
			wantErr: false,
		},
		{
			name: "Case 4",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: errors.New("Test. "),
				message: new(messages.TextMessage).
					Text("Message. "),
				details: new(details.Details).
					Set("key", "value"),

				ctx: context.Background(),
			},
			want:    []byte{60, 69, 114, 114, 111, 114, 32, 105, 100, 61, 34, 84, 45, 48, 48, 48, 48, 48, 49, 34, 32, 116, 121, 112, 101, 61, 34, 115, 121, 115, 116, 101, 109, 34, 32, 115, 116, 97, 116, 117, 115, 61, 34, 102, 97, 116, 97, 108, 34, 62, 60, 77, 101, 115, 115, 97, 103, 101, 62, 77, 101, 115, 115, 97, 103, 101, 46, 32, 60, 47, 77, 101, 115, 115, 97, 103, 101, 62, 60, 68, 101, 116, 97, 105, 108, 115, 62, 60, 73, 116, 101, 109, 32, 107, 101, 121, 61, 34, 107, 101, 121, 34, 62, 118, 97, 108, 117, 101, 60, 47, 73, 116, 101, 109, 62, 60, 47, 68, 101, 116, 97, 105, 108, 115, 62, 60, 47, 69, 114, 114, 111, 114, 62},
			wantErr: false,
		},
		{
			name: "Case 5",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: nil,
				message: new(messages.TextMessage).
					Text("Message. "),
				details: new(details.Details).
					Set("key", "value"),

				ctx: context.Background(),
			},
			want:    []byte{60, 69, 114, 114, 111, 114, 32, 105, 100, 61, 34, 84, 45, 48, 48, 48, 48, 48, 49, 34, 32, 116, 121, 112, 101, 61, 34, 115, 121, 115, 116, 101, 109, 34, 32, 115, 116, 97, 116, 117, 115, 61, 34, 102, 97, 116, 97, 108, 34, 62, 60, 77, 101, 115, 115, 97, 103, 101, 62, 77, 101, 115, 115, 97, 103, 101, 46, 32, 60, 47, 77, 101, 115, 115, 97, 103, 101, 62, 60, 68, 101, 116, 97, 105, 108, 115, 62, 60, 73, 116, 101, 109, 32, 107, 101, 121, 61, 34, 107, 101, 121, 34, 62, 118, 97, 108, 117, 101, 60, 47, 73, 116, 101, 109, 62, 60, 47, 68, 101, 116, 97, 105, 108, 115, 62, 60, 47, 69, 114, 114, 111, 114, 62},
			wantErr: false,
		},
		{
			name: "Case 6",
			fields: fields{
				id:     "T-000001",
				t:      types.TypeSystem,
				status: types.StatusFatal,

				err: errors.New("Test. "),
				message: new(messages.TextMessage).
					Text("Message. "),
				details: new(details.Details).
					Set("key", "value"),

				ctx: context.Background(),
			},
			want:    []byte{60, 69, 114, 114, 111, 114, 32, 105, 100, 61, 34, 84, 45, 48, 48, 48, 48, 48, 49, 34, 32, 116, 121, 112, 101, 61, 34, 115, 121, 115, 116, 101, 109, 34, 32, 115, 116, 97, 116, 117, 115, 61, 34, 102, 97, 116, 97, 108, 34, 62, 60, 77, 101, 115, 115, 97, 103, 101, 62, 77, 101, 115, 115, 97, 103, 101, 46, 32, 60, 47, 77, 101, 115, 115, 97, 103, 101, 62, 60, 68, 101, 116, 97, 105, 108, 115, 62, 60, 73, 116, 101, 109, 32, 107, 101, 121, 61, 34, 107, 101, 121, 34, 62, 118, 97, 108, 117, 101, 60, 47, 73, 116, 101, 109, 62, 60, 47, 68, 101, 116, 97, 105, 108, 115, 62, 60, 47, 69, 114, 114, 111, 114, 62},
			wantErr: false,
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

			got, err := xml.Marshal(i)

			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalXML() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Internal_UnmarshalJSON(t *testing.T) {
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
		bytes []byte
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Internal
		wantErr bool
	}{
		{
			name:   "Case 1",
			fields: fields{},
			args: args{
				bytes: []byte{123, 34, 105, 100, 34, 58, 34, 84, 45, 48, 48, 48, 48, 48, 49, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 115, 121, 115, 116, 101, 109, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 116, 97, 108, 34, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 77, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 100, 101, 116, 97, 105, 108, 115, 34, 58, 123, 125, 125},
			},
			wantErr: false,
			want: &Internal{
				Store: &Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Err: nil,
					Message: new(messages.TextMessage).
						Text("Message. "),
					Details: new(details.Details),
				},

				ctx: context.Background(),
			},
		},
		{
			name:   "Case 2",
			fields: fields{},
			args: args{
				bytes: []byte{123, 34, 105, 100, 34, 58, 34, 84, 45, 48, 48, 48, 48, 48, 49, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 115, 121, 115, 116, 101, 109, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 116, 97, 108, 34, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 77, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 100, 101, 116, 97, 105, 108, 115, 34, 58, 123, 125, 125},
			},
			wantErr: false,
			want: &Internal{
				Store: &Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Err: errors.New("Test. "),
					Message: new(messages.TextMessage).
						Text("Message. "),
					Details: new(details.Details),
				},

				ctx: context.Background(),
			},
		},
		{
			name:   "Case 3",
			fields: fields{},
			args: args{
				bytes: []byte{123, 34, 105, 100, 34, 58, 34, 84, 45, 48, 48, 48, 48, 48, 49, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 115, 121, 115, 116, 101, 109, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 116, 97, 108, 34, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 77, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 100, 101, 116, 97, 105, 108, 115, 34, 58, 123, 34, 107, 101, 121, 34, 58, 34, 118, 97, 108, 117, 101, 34, 125, 125},
			},
			wantErr: false,
			want: &Internal{
				Store: &Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Err: nil,
					Message: new(messages.TextMessage).
						Text("Message. "),
					Details: new(details.Details).
						Set("key", "value"),
				},

				ctx: context.Background(),
			},
		},
		{
			name:   "Case 4",
			fields: fields{},
			args: args{
				bytes: []byte{123, 34, 105, 100, 34, 58, 34, 84, 45, 48, 48, 48, 48, 48, 49, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 115, 121, 115, 116, 101, 109, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 116, 97, 108, 34, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 77, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 100, 101, 116, 97, 105, 108, 115, 34, 58, 123, 34, 107, 101, 121, 34, 58, 34, 118, 97, 108, 117, 101, 34, 125, 125},
			},
			wantErr: false,
			want: &Internal{
				Store: &Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Err: errors.New("Test. "),
					Message: new(messages.TextMessage).
						Text("Message. "),
					Details: new(details.Details).
						Set("key", "value"),
				},

				ctx: context.Background(),
			},
		},
		{
			name:   "Case 5",
			fields: fields{},
			args: args{
				bytes: []byte{123, 34, 105, 100, 34, 58, 34, 84, 45, 48, 48, 48, 48, 48, 49, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 115, 121, 115, 116, 101, 109, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 116, 97, 108, 34, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 77, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 100, 101, 116, 97, 105, 108, 115, 34, 58, 123, 34, 107, 101, 121, 34, 58, 34, 118, 97, 108, 117, 101, 34, 125, 125},
			},
			wantErr: false,
			want: &Internal{
				Store: &Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Err: nil,
					Message: new(messages.TextMessage).
						Text("Message. "),
					Details: new(details.Details).
						Set("key", "value"),
				},

				ctx: context.Background(),
			},
		},
		{
			name:   "Case 6",
			fields: fields{},
			args: args{
				bytes: []byte{123, 34, 105, 100, 34, 58, 34, 84, 45, 48, 48, 48, 48, 48, 49, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 115, 121, 115, 116, 101, 109, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 116, 97, 108, 34, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 77, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 100, 101, 116, 97, 105, 108, 115, 34, 58, 123, 34, 107, 101, 121, 34, 58, 34, 118, 97, 108, 117, 101, 34, 125, 125},
			},
			wantErr: false,
			want: &Internal{
				Store: &Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Err: errors.New("Test. "),
					Message: new(messages.TextMessage).
						Text("Message. "),
					Details: new(details.Details).
						Set("key", "value"),
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

			if err := json.Unmarshal(tt.args.bytes, i); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
