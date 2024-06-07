package details

import (
	"reflect"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/types"
	"sync"
	"testing"
)

func TestDetails_init(t *testing.T) {
	type fields struct {
		fields  Fields
		storage map[string]any
		rwMux   *sync.RWMutex
	}

	tests := []struct {
		name   string
		fields fields
		want   types.Details
	}{
		{
			name: "Case 1",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("1"),
					},
				},
				storage: map[string]any{
					"test1": "1",
					"test2": false,
				},
				rwMux: new(sync.RWMutex),
			},
			want: &Details{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("1"),
					},
				},
				storage: map[string]any{
					"test1": "1",
					"test2": false,
				},
				rwMux: new(sync.RWMutex),
			},
		},
		{
			name: "Case 2",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("1"),
					},
				},
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			want: &Details{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("1"),
					},
				},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
		},
		{
			name: "Case 3",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("1"),
					},
				},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
			want: &Details{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("1"),
					},
				},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &Details{
				fields:  tt.fields.fields,
				storage: tt.fields.storage,
				rwMux:   tt.fields.rwMux,
			}

			if ds.init(); !reflect.DeepEqual(ds, tt.want) {
				t.Errorf("Set() = %v, want %v", ds, tt.want)
			}
		})
	}
}

func TestDetails_Peek(t *testing.T) {
	type fields struct {
		fields  Fields
		storage map[string]any
		rwMux   *sync.RWMutex
	}

	type args struct {
		k string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		wantV  any
	}{
		{
			name: "Case 1",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("1"),
					},
				},
				storage: map[string]any{
					"test1": "1",
					"test2": false,
				},
				rwMux: new(sync.RWMutex),
			},
			args: args{
				k: "test1",
			},
			wantV: "1",
		},
		{
			name: "Case 2",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("1"),
					},
				},
				storage: map[string]any{
					"test1": "1",
					"test2": false,
				},
				rwMux: new(sync.RWMutex),
			},
			args: args{
				k: "test2",
			},
			wantV: false,
		},
		{
			name: "Case 3",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("1"),
					},
				},
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			args: args{
				k: "test",
			},
			wantV: nil,
		},
		{
			name: "Case 4",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("1"),
					},
				},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
			args: args{
				k: "test",
			},
			wantV: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &Details{
				fields:  tt.fields.fields,
				storage: tt.fields.storage,
				rwMux:   tt.fields.rwMux,
			}

			if gotV := ds.Peek(tt.args.k); !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("Peek() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestDetails_Set(t *testing.T) {
	type fields struct {
		fields  Fields
		storage map[string]any
		rwMux   *sync.RWMutex
	}

	type args struct {
		k string
		v any
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   types.Details
	}{
		{
			name: "Case 1",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("1"),
					},
				},
				storage: map[string]any{
					"test1": "1",
					"test2": false,
				},
				rwMux: new(sync.RWMutex),
			},
			args: args{
				k: "test3",
				v: "3",
			},
			want: &Details{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("1"),
					},
				},
				storage: map[string]any{
					"test1": "1",
					"test2": false,
					"test3": "3",
				},
				rwMux: new(sync.RWMutex),
			},
		},
		{
			name: "Case 2",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("1"),
					},
				},
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			args: args{
				k: "test1",
				v: "1",
			},
			want: &Details{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("1"),
					},
				},
				storage: map[string]any{
					"test1": "1",
				},
				rwMux: new(sync.RWMutex),
			},
		},
		{
			name: "Case 3",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("1"),
					},
				},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
			args: args{
				k: "test1",
				v: "1",
			},
			want: &Details{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("1"),
					},
				},
				storage: map[string]any{
					"test1": "1",
				},
				rwMux: new(sync.RWMutex),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &Details{
				fields:  tt.fields.fields,
				storage: tt.fields.storage,
				rwMux:   tt.fields.rwMux,
			}

			if got := ds.Set(tt.args.k, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetails_Reset(t *testing.T) {
	type fields struct {
		fields  Fields
		storage map[string]any
		rwMux   *sync.RWMutex
	}

	tests := []struct {
		name   string
		fields fields
		want   types.Details
	}{
		{
			name: "Case 1",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("123"),
					},
					{
						Key:     new(FieldKey).Add("test2"),
						Message: new(messages.TextMessage).Text("321"),
					},
				},
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			want: &Details{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("123"),
					},
					{
						Key:     new(FieldKey).Add("test2"),
						Message: new(messages.TextMessage).Text("321"),
					},
				},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
		},
		{
			name: "Case 2",
			fields: fields{
				fields:  Fields{},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
			want: &Details{
				fields:  Fields{},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
		},
		{
			name: "Case 3",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("123"),
					},
					{
						Key:     new(FieldKey).Add("test2"),
						Message: new(messages.TextMessage).Text("321"),
					},
				},
				storage: map[string]any{
					"test": "123",
				},
				rwMux: new(sync.RWMutex),
			},
			want: &Details{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("123"),
					},
					{
						Key:     new(FieldKey).Add("test2"),
						Message: new(messages.TextMessage).Text("321"),
					},
				},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
		},
		{
			name: "Case 4",
			fields: fields{
				fields: nil,
				storage: map[string]any{
					"test": true,
					"key":  "value",
				},
				rwMux: new(sync.RWMutex),
			},
			want: &Details{
				fields:  Fields{},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &Details{
				fields:  tt.fields.fields,
				storage: tt.fields.storage,
				rwMux:   tt.fields.rwMux,
			}

			if got := ds.Reset(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetails_PeekFieldMessage(t *testing.T) {
	type fields struct {
		fields  Fields
		storage map[string]any
		rwMux   *sync.RWMutex
	}

	type args struct {
		k string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		wantM  types.DetailsFieldMessage
	}{
		{
			name: "Case 1",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("123"),
					},
					{
						Key:     new(FieldKey).Add("test2"),
						Message: new(messages.TextMessage).Text("321"),
					},
				},
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			args: args{
				k: "test1",
			},
			wantM: new(messages.TextMessage).Text("123"),
		},
		{
			name: "Case 2",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("123"),
					},
					{
						Key:     new(FieldKey).Add("test2"),
						Message: new(messages.TextMessage).Text("321"),
					},
				},
				storage: map[string]any{
					"test1": "312",
				},
				rwMux: new(sync.RWMutex),
			},
			args: args{
				k: "test1",
			},
			wantM: new(messages.TextMessage).Text("123"),
		},
		{
			name: "Case 3",
			fields: fields{
				fields:  Fields{},
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			args: args{
				k: "test1",
			},
			wantM: nil,
		},
		{
			name: "Case 4",
			fields: fields{
				fields:  nil,
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			args: args{
				k: "test1",
			},
			wantM: nil,
		},
		{
			name: "Case 5",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("123"),
					},
					{
						Key:     new(FieldKey).Add("test2"),
						Message: new(messages.TextMessage).Text("321"),
					},
				},
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			args: args{
				k: "test3",
			},
			wantM: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &Details{
				fields:  tt.fields.fields,
				storage: tt.fields.storage,
				rwMux:   tt.fields.rwMux,
			}

			if gotM := ds.PeekFieldMessage(tt.args.k); !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("PeekFieldMessage() = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

func TestDetails_SetField(t *testing.T) {
	type fields struct {
		fields  Fields
		storage map[string]any
		rwMux   *sync.RWMutex
	}

	type args struct {
		k types.DetailsFieldKey
		m types.DetailsFieldMessage
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   types.Details
	}{
		{
			name: "Case 1",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
				storage: map[string]any{
					"test": "123",
				},
				rwMux: new(sync.RWMutex),
			},
			args: args{
				k: new(FieldKey).Add("test2"),
				m: new(messages.TextMessage).Text("321"),
			},
			want: &Details{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
					{
						Key:     new(FieldKey).Add("test2"),
						Message: new(messages.TextMessage).Text("321"),
					},
				},
				storage: map[string]any{
					"test": "123",
				},
				rwMux: new(sync.RWMutex),
			},
		},
		{
			name: "Case 2",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			args: args{
				k: new(FieldKey).Add("test2"),
				m: new(messages.TextMessage).Text("321"),
			},
			want: &Details{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
					{
						Key:     new(FieldKey).Add("test2"),
						Message: new(messages.TextMessage).Text("321"),
					},
				},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
		},
		{
			name: "Case 3",
			fields: fields{
				fields:  Fields{},
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			args: args{
				k: new(FieldKey).Add("test"),
				m: new(messages.TextMessage).Text("123"),
			},
			want: &Details{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
		},
		{
			name: "Case 4",
			fields: fields{
				fields:  nil,
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			args: args{
				k: new(FieldKey).Add("test"),
				m: new(messages.TextMessage).Text("123"),
			},
			want: &Details{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &Details{
				fields:  tt.fields.fields,
				storage: tt.fields.storage,
				rwMux:   tt.fields.rwMux,
			}

			if got := ds.SetField(tt.args.k, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetails_SetFields(t *testing.T) {
	type fields struct {
		fields  Fields
		storage map[string]any
		rwMux   *sync.RWMutex
	}

	type args struct {
		fields []types.DetailsField
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   types.Details
	}{
		{
			name: "Case 1",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
				storage: map[string]any{
					"test": "123",
				},
				rwMux: new(sync.RWMutex),
			},
			args: args{
				fields: []types.DetailsField{
					{
						Key:     new(FieldKey).Add("test2"),
						Message: new(messages.TextMessage).Text("321"),
					},
				},
			},
			want: &Details{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
					{
						Key:     new(FieldKey).Add("test2"),
						Message: new(messages.TextMessage).Text("321"),
					},
				},
				storage: map[string]any{
					"test": "123",
				},
				rwMux: new(sync.RWMutex),
			},
		},
		{
			name: "Case 2",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			args: args{
				fields: []types.DetailsField{
					{
						Key:     new(FieldKey).Add("test2"),
						Message: new(messages.TextMessage).Text("321"),
					},
				},
			},
			want: &Details{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
					{
						Key:     new(FieldKey).Add("test2"),
						Message: new(messages.TextMessage).Text("321"),
					},
				},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
		},
		{
			name: "Case 3",
			fields: fields{
				fields:  Fields{},
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			args: args{
				fields: []types.DetailsField{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
			},
			want: &Details{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
		},
		{
			name: "Case 4",
			fields: fields{
				fields:  nil,
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			args: args{
				fields: []types.DetailsField{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
			},
			want: &Details{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
		},
		{
			name: "Case 5",
			fields: fields{
				fields:  nil,
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			args: args{
				fields: []types.DetailsField{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("123"),
					},
					{
						Key:     new(FieldKey).Add("test2"),
						Message: new(messages.TextMessage).Text("321"),
					},
				},
			},
			want: &Details{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("123"),
					},
					{
						Key:     new(FieldKey).Add("test2"),
						Message: new(messages.TextMessage).Text("321"),
					},
				},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &Details{
				fields:  tt.fields.fields,
				storage: tt.fields.storage,
				rwMux:   tt.fields.rwMux,
			}

			if got := ds.SetFields(tt.args.fields...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetails_ResetFields(t *testing.T) {
	type fields struct {
		fields  Fields
		storage map[string]any
		rwMux   *sync.RWMutex
	}

	tests := []struct {
		name   string
		fields fields
		want   types.Details
	}{
		{
			name: "Case 1",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("123"),
					},
					{
						Key:     new(FieldKey).Add("test2"),
						Message: new(messages.TextMessage).Text("321"),
					},
				},
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			want: &Details{
				fields:  Fields{},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
		},
		{
			name: "Case 2",
			fields: fields{
				fields:  Fields{},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
			want: &Details{
				fields:  Fields{},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
		},
		{
			name: "Case 3",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test1"),
						Message: new(messages.TextMessage).Text("123"),
					},
					{
						Key:     new(FieldKey).Add("test2"),
						Message: new(messages.TextMessage).Text("321"),
					},
				},
				storage: map[string]any{
					"test": "123",
				},
				rwMux: new(sync.RWMutex),
			},
			want: &Details{
				fields: Fields{},
				storage: map[string]any{
					"test": "123",
				},
				rwMux: new(sync.RWMutex),
			},
		},
		{
			name: "Case 4",
			fields: fields{
				fields: nil,
				storage: map[string]any{
					"test": "123",
				},
				rwMux: new(sync.RWMutex),
			},
			want: &Details{
				fields: Fields{},
				storage: map[string]any{
					"test": "123",
				},
				rwMux: new(sync.RWMutex),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &Details{
				fields:  tt.fields.fields,
				storage: tt.fields.storage,
				rwMux:   tt.fields.rwMux,
			}

			if got := ds.ResetFields(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResetFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetails_Clone(t *testing.T) {
	type fields struct {
		fields  Fields
		storage map[string]any
		rwMux   *sync.RWMutex
	}

	tests := []struct {
		name   string
		fields fields
		want   types.Details
	}{
		{
			name: "Case 1",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
				storage: map[string]any{
					"test": "123",
				},
				rwMux: new(sync.RWMutex),
			},
			want: &Details{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
				storage: map[string]any{
					"test": "123",
				},
				rwMux: new(sync.RWMutex),
			},
		},
		{
			name: "Case 2",
			fields: fields{
				fields: nil,
				storage: map[string]any{
					"test": "123",
				},
				rwMux: new(sync.RWMutex),
			},
			want: &Details{
				fields: Fields{},
				storage: map[string]any{
					"test": "123",
				},
				rwMux: new(sync.RWMutex),
			},
		},
		{
			name: "Case 3",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			want: &Details{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
		},
		{
			name: "Case 4",
			fields: fields{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
				storage: map[string]any{
					"test": "123",
				},
				rwMux: nil,
			},
			want: &Details{
				fields: Fields{
					{
						Key:     new(FieldKey).Add("test"),
						Message: new(messages.TextMessage).Text("123"),
					},
				},
				storage: map[string]any{
					"test": "123",
				},
				rwMux: new(sync.RWMutex),
			},
		},
		{
			name: "Case 5",
			fields: fields{
				fields:  nil,
				storage: nil,
				rwMux:   new(sync.RWMutex),
			},
			want: &Details{
				fields:  Fields{},
				storage: map[string]any{},
				rwMux:   new(sync.RWMutex),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &Details{
				fields:  tt.fields.fields,
				storage: tt.fields.storage,
				rwMux:   tt.fields.rwMux,
			}

			if got := ds.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}
