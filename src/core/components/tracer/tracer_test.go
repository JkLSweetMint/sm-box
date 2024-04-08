package tracer

import (
	"errors"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		levels []Level
	}

	tests := []struct {
		name  string
		args  args
		wantT Tracer
	}{
		{
			name: "Case 1",
			args: args{
				levels: allLevels,
			},
			wantT: &tracer{
				params:    make([]any, maxParams),
				levels:    allLevels,
				inputArgs: nil,
				err:       nil,
			},
		},
		{
			name: "Case 2",
			args: args{
				levels: []Level{
					LevelMain,
					LevelDebug,
					LevelInternal,
					LevelEvent,
				},
			},
			wantT: &tracer{
				params: make([]any, maxParams),
				levels: []Level{
					LevelMain,
					LevelDebug,
					LevelInternal,
					LevelEvent,
				},
				inputArgs: nil,
				err:       nil,
			},
		},
		{
			name: "Case 3",
			args: args{
				levels: nil,
			},
			wantT: &tracer{
				params:    make([]any, maxParams),
				levels:    nil,
				inputArgs: nil,
				err:       nil,
			},
		},
		{
			name: "Case 4",
			args: args{
				levels: []Level{},
			},
			wantT: &tracer{
				params:    make([]any, maxParams),
				levels:    []Level{},
				inputArgs: nil,
				err:       nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotT := New(tt.args.levels...); !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("New() = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}

func Test_tracer_Error(t *testing.T) {
	type fields struct {
		params    []any
		levels    []Level
		inputArgs []any
		err       error
	}

	type args struct {
		err error
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   Tracer
	}{
		{
			name: "Case 1",
			fields: fields{
				params:    make([]any, maxParams),
				levels:    allLevels,
				inputArgs: nil,
				err:       nil,
			},
			args: args{
				err: nil,
			},
			want: &tracer{
				params:    make([]any, maxParams),
				levels:    allLevels,
				inputArgs: nil,
				err:       nil,
			},
		},
		{
			name: "Case 2",
			fields: fields{
				params:    make([]any, maxParams),
				levels:    allLevels,
				inputArgs: nil,
				err:       nil,
			},
			args: args{
				err: errors.New("Test. "),
			},
			want: &tracer{
				params:    make([]any, maxParams),
				levels:    allLevels,
				inputArgs: nil,
				err:       errors.New("Test. "),
			},
		},
		{
			name: "Case 3",
			fields: fields{
				params:    make([]any, maxParams),
				levels:    allLevels,
				inputArgs: nil,
				err:       errors.New("Test 1. "),
			},
			args: args{
				err: errors.New("Test 2. "),
			},
			want: &tracer{
				params:    make([]any, maxParams),
				levels:    allLevels,
				inputArgs: nil,
				err:       errors.New("Test 2. "),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trc := &tracer{
				params:    tt.fields.params,
				levels:    tt.fields.levels,
				inputArgs: tt.fields.inputArgs,
				err:       tt.fields.err,
			}

			if got := trc.Error(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tracer_FunctionCall(t *testing.T) {
	type fields struct {
		params    []any
		levels    []Level
		inputArgs []any
		err       error
	}

	type args struct {
		args []any
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Case 1",
			fields: fields{
				params:    make([]any, maxParams),
				levels:    allLevels,
				inputArgs: nil,
				err:       nil,
			},
			args: args{
				args: []any{123},
			},
		},
		{
			name: "Case 2",
			fields: fields{
				params:    make([]any, maxParams),
				levels:    allLevels,
				inputArgs: nil,
				err:       errors.New("Test. "),
			},
			args: args{
				args: []any{123},
			},
		},
		{
			name: "Case 3",
			fields: fields{
				params:    make([]any, maxParams),
				levels:    allLevels,
				inputArgs: nil,
				err:       errors.New("Test. "),
			},
			args: args{},
		},
		{
			name: "Case 4",
			fields: fields{
				params:    make([]any, maxParams),
				levels:    allLevels,
				inputArgs: nil,
				err:       nil,
			},
			args: args{},
		},
		{
			name: "Case 5",
			fields: fields{
				params:    make([]any, maxParams),
				levels:    allLevels,
				inputArgs: []any{321},
				err:       errors.New("Test. "),
			},
			args: args{},
		},
		{
			name: "Case 6",
			fields: fields{
				params:    make([]any, maxParams),
				levels:    allLevels,
				inputArgs: []any{321},
				err:       nil,
			},
			args: args{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trc := &tracer{
				params:    tt.fields.params,
				levels:    tt.fields.levels,
				inputArgs: tt.fields.inputArgs,
				err:       tt.fields.err,
			}

			trc.FunctionCall(tt.args.args...)
		})
	}
}

func Test_tracer_FunctionCallFinished(t *testing.T) {
	type fields struct {
		params    []any
		levels    []Level
		inputArgs []any
		err       error
	}

	type args struct {
		args []any
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Case 1",
			fields: fields{
				params:    make([]any, maxParams),
				levels:    allLevels,
				inputArgs: nil,
				err:       nil,
			},
			args: args{
				args: []any{123},
			},
		},
		{
			name: "Case 2",
			fields: fields{
				params:    make([]any, maxParams),
				levels:    allLevels,
				inputArgs: nil,
				err:       errors.New("Test. "),
			},
			args: args{
				args: []any{123},
			},
		},
		{
			name: "Case 3",
			fields: fields{
				params:    make([]any, maxParams),
				levels:    allLevels,
				inputArgs: nil,
				err:       errors.New("Test. "),
			},
			args: args{},
		},
		{
			name: "Case 4",
			fields: fields{
				params:    make([]any, maxParams),
				levels:    allLevels,
				inputArgs: nil,
				err:       nil,
			},
			args: args{},
		},
		{
			name: "Case 5",
			fields: fields{
				params:    make([]any, maxParams),
				levels:    allLevels,
				inputArgs: []any{321},
				err:       errors.New("Test. "),
			},
			args: args{},
		},
		{
			name: "Case 6",
			fields: fields{
				params:    make([]any, maxParams),
				levels:    allLevels,
				inputArgs: []any{321},
				err:       nil,
			},
			args: args{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trc := &tracer{
				params:    tt.fields.params,
				levels:    tt.fields.levels,
				inputArgs: tt.fields.inputArgs,
				err:       tt.fields.err,
			}

			trc.FunctionCallFinished(tt.args.args...)
		})
	}
}
