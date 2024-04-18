package task_scheduler

import (
	"context"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Case 1",
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSc, gotC, err := New(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			switch {
			case gotC == nil:
				t.Error("Call channel is nil. ")
			case gotSc == nil:
				t.Error("Scheduler is nil. ")
			default:
				{
					var sc = gotSc.(*scheduler)

					switch {
					case sc.channel == nil:
						t.Error("Scheduler call channel is nil. ")
					case sc.aggregate == nil:
						t.Error("Scheduler aggregate is nil. ")
					case sc.components == nil:
						t.Error("Scheduler components is nil. ")
					case sc.components.Logger == nil:
						t.Error("Scheduler logger component is nil. ")
					}
				}
			}
		})
	}
}
