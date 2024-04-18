package task_scheduler

import (
	"context"
	"fmt"
	"sm-box/src/pkg/core/components/logger"
	"sm-box/src/pkg/core/env"
	"sync"
	"testing"
	"time"
)

func Test_scheduler_Register(t *testing.T) {
	type fields struct {
		aggregate  aggregate
		channel    chan TaskType
		components *components
	}

	type args struct {
		t Task
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
				aggregate: &baseShelf{
					Tasks: []*Task{},
					rwMx:  new(sync.RWMutex),
				},
				channel:    make(chan TaskType),
				components: new(components),
			},
			args: args{
				t: Task{
					Name: "Test task 1",
					Type: TaskServe,
					Func: nil,
				},
			},
			wantErr: false,
		},
		{
			name: "Case 2",
			fields: fields{
				aggregate: &baseShelf{
					Tasks: []*Task{},
					rwMx:  new(sync.RWMutex),
				},
				channel:    make(chan TaskType),
				components: new(components),
			},
			args: args{
				t: Task{
					Name: "Test task 1",
					Type: minTaskType,
					Func: nil,
				},
			},
			wantErr: true,
		},
		{
			name: "Case 3",
			fields: fields{
				aggregate: &baseShelf{
					Tasks: []*Task{},
					rwMx:  new(sync.RWMutex),
				},
				channel:    make(chan TaskType),
				components: new(components),
			},
			args: args{
				t: Task{
					Name: "Test task 1",
					Type: maxTaskType,
					Func: nil,
				},
			},
			wantErr: true,
		},
		{
			name: "Case 4",
			fields: fields{
				aggregate: &baseShelf{
					Tasks: []*Task{},
					rwMx:  new(sync.RWMutex),
				},
				channel:    make(chan TaskType),
				components: new(components),
			},
			args: args{
				t: Task{
					Name: "Test task 1",
					Type: -1,
					Func: nil,
				},
			},
			wantErr: true,
		},
		{
			name: "Case 5",
			fields: fields{
				aggregate: &baseShelf{
					Tasks: []*Task{
						{
							Name: "Test task 1",
							Type: TaskBeforeNew,
							Func: nil,
						}, {
							Name: "Test task 2",
							Type: TaskAfterBoot,
							Func: nil,
						},
					},
					rwMx: new(sync.RWMutex),
				},
				channel:    make(chan TaskType),
				components: new(components),
			},
			args: args{
				t: Task{
					Name: "Test task 3",
					Type: TaskServe,
					Func: nil,
				},
			},
			wantErr: false,
		},
		{
			name: "Case 6",
			fields: fields{
				aggregate: &baseShelf{
					Tasks: []*Task{
						{
							Name: "Test task 1",
							Type: TaskAfterServe,
							Func: nil,
						},
						{
							Name: "Test task 2",
							Type: TaskAfterShutdown,
							Func: nil,
						},
					},
					rwMx: new(sync.RWMutex),
				},
				channel:    make(chan TaskType),
				components: new(components),
			},
			args: args{
				t: Task{
					Name: "Test task 3",
					Type: minTaskType,
					Func: nil,
				},
			},
			wantErr: true,
		},
		{
			name: "Case 7",
			fields: fields{
				aggregate: &baseShelf{
					Tasks: []*Task{
						{
							Name: "Test task 1",
							Type: TaskBoot,
							Func: nil,
						},
						{
							Name: "Test task 2",
							Type: TaskBeforeServe,
							Func: nil,
						},
					},
					rwMx: new(sync.RWMutex),
				},
				channel:    make(chan TaskType),
				components: new(components),
			},
			args: args{
				t: Task{
					Name: "Test task 3",
					Type: maxTaskType,
					Func: nil,
				},
			},
			wantErr: true,
		},
		{
			name: "Case 8",
			fields: fields{
				aggregate: &baseShelf{
					Tasks: []*Task{
						{
							Name: "Test task 1",
							Type: TaskBeforeBoot,
							Func: nil,
						},
						{
							Name: "Test task 2",
							Type: TaskAfterBoot,
							Func: nil,
						},
					},
					rwMx: new(sync.RWMutex),
				},
				channel:    make(chan TaskType),
				components: new(components),
			},
			args: args{
				t: Task{
					Name: "Test task 3",
					Type: -1,
					Func: nil,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt.fields.components.Logger, _ = logger.New(loggerInitiator)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scheduler{
				aggregate:  tt.fields.aggregate,
				channel:    tt.fields.channel,
				components: tt.fields.components,
			}

			if err := s.Register(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scheduler_tracking(t *testing.T) {
	var (
		ctx, cancel = context.WithCancel(context.Background())
		sc, c, err  = New(ctx)
	)

	if err != nil {
		t.Errorf("Failed to create a task scheduler: '%s'. ", err)
	}

	if err = sc.Register(Task{
		Name: "Test task 1",
		Type: TaskServe,
		Func: func(ctx context.Context) (err error) {
			fmt.Println("Test task 1")
			return
		},
	}); err != nil {
		t.Errorf("Failed to register a task in task scheduler: '%s'. ", err)
	}

	if err = sc.Register(Task{
		Name: "Test task 2",
		Type: TaskShutdown,
		Func: func(ctx context.Context) (err error) {
			fmt.Println("Test task 2")
			return
		},
	}); err != nil {
		t.Errorf("Failed to register a task in task scheduler: '%s'. ", err)
	}

	if err = sc.Register(Task{
		Name: "Test task 3",
		Type: TaskServe,
		Func: func(ctx context.Context) (err error) {
			fmt.Println("Test task 3")
			return
		},
	}); err != nil {
		t.Errorf("Failed to register a task in task scheduler: '%s'. ", err)
	}

	var completed = make(chan struct{})

	go func() {
		c <- TaskServe

		time.Sleep(time.Second * 3)

		c <- TaskShutdown

		time.Sleep(time.Second * 3)

		cancel()

		env.Synchronization.WaitGroup.Wait()

		completed <- struct{}{}
	}()

	select {
	case <-completed:
		return
	case <-time.NewTimer(time.Second * 10).C:
		t.Error("The test execution time has exceeded the allowed value. ")
	}
}
