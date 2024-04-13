package task_scheduler

import (
	"reflect"
	"sync"
	"testing"
)

func Test_baseShelf_Add(t *testing.T) {
	type fields struct {
		Tasks []*Task
		rwMx  *sync.RWMutex
	}

	type args struct {
		t Task
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   aggregate
	}{
		{
			name: "Case 1",
			fields: fields{
				Tasks: []*Task{},
				rwMx:  new(sync.RWMutex),
			},
			args: args{
				t: Task{
					Name: "Test task 1",
					Type: 0,
					Func: nil,
				},
			},
			want: &baseShelf{
				Tasks: []*Task{
					{
						Name: "Test task 1",
						Type: 0,
						Func: nil,
					},
				},
				rwMx: new(sync.RWMutex),
			},
		},
		{
			name: "Case 2",
			fields: fields{
				Tasks: []*Task{
					{
						Name: "Test task 1",
						Type: 0,
						Func: nil,
					},
				},
				rwMx: new(sync.RWMutex),
			},
			args: args{
				t: Task{
					Name: "Test task 2",
					Type: 0,
					Func: nil,
				},
			},
			want: &baseShelf{
				Tasks: []*Task{
					{
						Name: "Test task 1",
						Type: 0,
						Func: nil,
					},
					{
						Name: "Test task 2",
						Type: 0,
						Func: nil,
					},
				},
				rwMx: new(sync.RWMutex),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &baseShelf{
				Tasks: tt.fields.Tasks,
				rwMx:  tt.fields.rwMx,
			}

			s.Add(tt.args.t)

			if !reflect.DeepEqual(s.Tasks, tt.want.(*baseShelf).Tasks) {
				t.Errorf("End() = %v, want %v", s, tt.want)
			}
		})
	}
}

func Test_baseShelf_Iterator(t *testing.T) {
	type fields struct {
		Tasks []*Task
		rwMx  *sync.RWMutex
	}

	type args struct {
		tt TaskType
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   iterator
	}{
		{
			name: "Case 1",
			fields: fields{
				Tasks: []*Task{
					{
						Name: "Test task 1",
						Type: TaskServe,
						Func: nil,
					},
					{
						Name: "Test task 2",
						Type: TaskAfterBoot,
						Func: nil,
					},
					{
						Name: "Test task 3",
						Type: TaskServe,
						Func: nil,
					},
					{
						Name: "Test task 4",
						Type: TaskBeforeShutdown,
						Func: nil,
					},
				},
				rwMx: new(sync.RWMutex),
			},
			args: args{
				tt: TaskServe,
			},
			want: &baseIterator{
				shelf: &baseShelf{
					Tasks: []*Task{
						{
							Name: "Test task 1",
							Type: TaskServe,
							Func: nil,
						},
						{
							Name: "Test task 3",
							Type: TaskServe,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    0,
				internal: 0,
			},
		},
		{
			name: "Case 2",
			fields: fields{
				Tasks: []*Task{
					{
						Name: "Test task 1",
						Type: TaskServe,
						Func: nil,
					},
					{
						Name: "Test task 2",
						Type: TaskAfterBoot,
						Func: nil,
					},
					{
						Name: "Test task 3",
						Type: TaskServe,
						Func: nil,
					},
					{
						Name: "Test task 4",
						Type: TaskBeforeShutdown,
						Func: nil,
					},
				},
				rwMx: new(sync.RWMutex),
			},
			args: args{
				tt: TaskAfterBoot,
			},
			want: &baseIterator{
				shelf: &baseShelf{
					Tasks: []*Task{
						{
							Name: "Test task 2",
							Type: TaskAfterBoot,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    0,
				internal: 0,
			},
		},
		{
			name: "Case 3",
			fields: fields{
				Tasks: []*Task{
					{
						Name: "Test task 1",
						Type: TaskServe,
						Func: nil,
					},
					{
						Name: "Test task 2",
						Type: TaskAfterBoot,
						Func: nil,
					},
					{
						Name: "Test task 3",
						Type: TaskServe,
						Func: nil,
					},
					{
						Name: "Test task 4",
						Type: TaskBeforeShutdown,
						Func: nil,
					},
				},
				rwMx: new(sync.RWMutex),
			},
			args: args{
				tt: TaskShutdown,
			},
			want: &baseIterator{
				shelf: &baseShelf{
					Tasks: []*Task{},
					rwMx:  nil,
				},
				index:    0,
				internal: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &baseShelf{
				Tasks: tt.fields.Tasks,
				rwMx:  tt.fields.rwMx,
			}

			gotIter := s.Iterator(tt.args.tt)
			gotIter.(*baseIterator).shelf.rwMx = nil

			if !reflect.DeepEqual(gotIter, tt.want) {
				t.Errorf("Iterator() = %v, want %v", gotIter, tt.want)
			}
		})
	}
}
