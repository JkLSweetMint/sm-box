package task_scheduler

import (
	"context"
	"errors"
	"testing"
)

func TestTaskType_String(t *testing.T) {
	tests := []struct {
		name    string
		e       TaskType
		wantVal string
	}{
		{
			name:    "Case 1",
			e:       TaskBeforeNew,
			wantVal: allTaskTypesString[TaskBeforeNew-1],
		},
		{
			name:    "Case 2",
			e:       TaskBeforeBoot,
			wantVal: allTaskTypesString[TaskBeforeBoot-1],
		},
		{
			name:    "Case 3",
			e:       TaskBeforeServe,
			wantVal: allTaskTypesString[TaskBeforeServe-1],
		},
		{
			name:    "Case 4",
			e:       TaskBeforeShutdown,
			wantVal: allTaskTypesString[TaskBeforeShutdown-1],
		},
		{
			name:    "Case 5",
			e:       TaskAfterServe,
			wantVal: allTaskTypesString[TaskAfterServe-1],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVal := tt.e.String(); gotVal != tt.wantVal {
				t.Errorf("String() = %v, want %v", gotVal, tt.wantVal)
			}
		})
	}
}

func TestTask_Exec(t1 *testing.T) {
	type fields struct {
		Name string
		Type TaskType
		Func TaskFunc
	}

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Case 1",
			fields: fields{
				Name: "Test task 1",
				Type: TaskServe,
				Func: func(ctx context.Context) (err error) {
					return
				},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
		{
			name: "Case 2",
			fields: fields{
				Name: "Test task 2",
				Type: TaskServe,
				Func: func(ctx context.Context) (err error) {
					return errors.New("Test. ")
				},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Task{
				Name: tt.fields.Name,
				Type: tt.fields.Type,
				Func: tt.fields.Func,
			}

			if err := t.Exec(tt.args.ctx); (err != nil) != tt.wantErr {
				t1.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
