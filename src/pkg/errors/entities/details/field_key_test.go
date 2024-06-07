package details

import (
	"reflect"
	"sm-box/pkg/errors/types"
	"testing"
)

func TestFieldKey_Add(t *testing.T) {
	type fields struct {
		path []string
	}

	type args struct {
		name string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   types.DetailsFieldKey
	}{
		{
			name: "Case 1",
			fields: fields{
				path: []string{},
			},
			args: args{
				name: "test",
			},
			want: &FieldKey{
				path: []string{
					"test",
				},
			},
		},
		{
			name: "Case 2",
			fields: fields{
				path: []string{
					"test",
				},
			},
			args: args{
				name: "123",
			},
			want: &FieldKey{
				path: []string{
					"test",
					"123",
				},
			},
		},
		{
			name:   "Case 3",
			fields: fields{},
			args: args{
				name: "test",
			},
			want: &FieldKey{
				path: []string{
					"test",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fk := &FieldKey{
				path: tt.fields.path,
			}

			if got := fk.Add(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldKey_AddArray(t *testing.T) {
	type fields struct {
		path []string
	}

	type args struct {
		name  string
		index int
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   types.DetailsFieldKey
	}{
		{
			name: "Case 1",
			fields: fields{
				path: []string{},
			},
			args: args{
				name:  "test",
				index: 0,
			},
			want: &FieldKey{
				path: []string{
					"test[0]",
				},
			},
		},
		{
			name: "Case 2",
			fields: fields{
				path: []string{
					"test",
				},
			},
			args: args{
				name:  "arr",
				index: 0,
			},
			want: &FieldKey{
				path: []string{
					"test",
					"arr[0]",
				},
			},
		},
		{
			name:   "Case 3",
			fields: fields{},
			args: args{
				name:  "arr",
				index: 0,
			},
			want: &FieldKey{
				path: []string{
					"arr[0]",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fk := &FieldKey{
				path: tt.fields.path,
			}

			if got := fk.AddArray(tt.args.name, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldKey_AddMap(t *testing.T) {
	type fields struct {
		path []string
	}

	type args struct {
		name string
		key  any
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   types.DetailsFieldKey
	}{
		{
			name: "Case 1",
			fields: fields{
				path: []string{},
			},
			args: args{
				name: "test",
				key:  "key",
			},
			want: &FieldKey{
				path: []string{
					"test[key]",
				},
			},
		},
		{
			name: "Case 2",
			fields: fields{
				path: []string{
					"test",
				},
			},
			args: args{
				name: "map",
				key:  "key",
			},
			want: &FieldKey{
				path: []string{
					"test",
					"map[key]",
				},
			},
		},
		{
			name:   "Case 3",
			fields: fields{},
			args: args{
				name: "test",
				key:  "key",
			},
			want: &FieldKey{
				path: []string{
					"test[key]",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fk := &FieldKey{
				path: tt.fields.path,
			}

			if got := fk.AddMap(tt.args.name, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldKey_Clone(t *testing.T) {
	type fields struct {
		path []string
	}

	tests := []struct {
		name   string
		fields fields
		want   types.DetailsFieldKey
	}{
		{
			name:   "Case 1",
			fields: fields{},
			want: &FieldKey{
				path: []string{},
			},
		},
		{
			name: "Case 2",
			fields: fields{
				path: []string{},
			},
			want: &FieldKey{
				path: []string{},
			},
		},
		{
			name: "Case 3",
			fields: fields{
				path: []string{
					"test",
				},
			},
			want: &FieldKey{
				path: []string{
					"test",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fk := &FieldKey{
				path: tt.fields.path,
			}

			if got := fk.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldKey_String(t *testing.T) {
	type fields struct {
		path []string
	}

	tests := []struct {
		name    string
		fields  fields
		wantStr string
	}{
		{
			name: "Case 1",
			fields: fields{
				[]string{
					"test",
				},
			},
			wantStr: "test",
		},
		{
			name: "Case 2",
			fields: fields{
				[]string{
					"test",
					"arr[0]",
				},
			},
			wantStr: "test.arr[0]",
		},
		{
			name: "Case 3",
			fields: fields{
				[]string{
					"test",
					"map[key]",
				},
			},
			wantStr: "test.map[key]",
		},
		{
			name: "Case 4",
			fields: fields{
				[]string{},
			},
			wantStr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fk := &FieldKey{
				path: tt.fields.path,
			}

			if gotStr := fk.String(); gotStr != tt.wantStr {
				t.Errorf("String() = %v, want %v", gotStr, tt.wantStr)
			}
		})
	}
}

func TestFieldKey_init(t *testing.T) {
	type fields struct {
		path []string
	}

	tests := []struct {
		name   string
		fields fields
		want   types.DetailsFieldKey
	}{
		{
			name: "Case 1",
			fields: fields{
				path: []string{},
			},
			want: &FieldKey{
				path: []string{},
			},
		},
		{
			name:   "Case 2",
			fields: fields{},
			want: &FieldKey{
				path: []string{},
			},
		},
		{
			name: "Case 3",
			fields: fields{
				path: []string{
					"test",
				},
			},
			want: &FieldKey{
				path: []string{
					"test",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fk := &FieldKey{
				path: tt.fields.path,
			}

			fk.init()

			if !reflect.DeepEqual(fk, tt.want) {
				t.Errorf("Clone() = %v, want %v", fk, tt.want)
			}
		})
	}
}
