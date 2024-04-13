package task_scheduler

import (
	"reflect"
	"testing"
)

func Test_baseIterator_End(t *testing.T) {
	type fields struct {
		shelf    *baseShelf
		index    int
		internal int
	}

	tests := []struct {
		name   string
		fields fields
		want   iterator
	}{
		{
			name: "Case 1",
			fields: fields{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    0,
				internal: 0,
			},
			want: &baseIterator{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    2,
				internal: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := &baseIterator{
				shelf:    tt.fields.shelf,
				index:    tt.fields.index,
				internal: tt.fields.internal,
			}
			iter.End()

			if !reflect.DeepEqual(iter, tt.want) {
				t.Errorf("End() = %v, want %v", iter, tt.want)
			}
		})
	}
}

func Test_baseIterator_Has(t *testing.T) {
	type fields struct {
		shelf    *baseShelf
		index    int
		internal int
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Case 1",
			fields: fields{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    0,
				internal: 1,
			},
			want: true,
		},
		{
			name: "Case 2",
			fields: fields{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    1,
				internal: 2,
			},
			want: true,
		},
		{
			name: "Case 3",
			fields: fields{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    2,
				internal: 3,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := &baseIterator{
				shelf:    tt.fields.shelf,
				index:    tt.fields.index,
				internal: tt.fields.internal,
			}

			if gotHas := iter.Has(); gotHas != tt.want {
				t.Errorf("Has() = %v, want %v", gotHas, tt.want)
			}
		})
	}
}

func Test_baseIterator_Index(t *testing.T) {
	type fields struct {
		shelf    *baseShelf
		index    int
		internal int
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Case 1",
			fields: fields{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    0,
				internal: 0,
			},
			want: 0,
		},
		{
			name: "Case 2",
			fields: fields{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    1,
				internal: 1,
			},
			want: 1,
		},
		{
			name: "Case 3",
			fields: fields{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    2,
				internal: 2,
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := &baseIterator{
				shelf:    tt.fields.shelf,
				index:    tt.fields.index,
				internal: tt.fields.internal,
			}

			if got := iter.Index(); got != tt.want {
				t.Errorf("Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseIterator_Len(t *testing.T) {
	type fields struct {
		shelf    *baseShelf
		index    int
		internal int
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{

		{
			name: "Case 1",
			fields: fields{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    0,
				internal: 0,
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := &baseIterator{
				shelf:    tt.fields.shelf,
				index:    tt.fields.index,
				internal: tt.fields.internal,
			}

			if gotL := iter.Len(); gotL != tt.want {
				t.Errorf("Len() = %v, want %v", gotL, tt.want)
			}
		})
	}
}

func Test_baseIterator_Next(t *testing.T) {
	type fields struct {
		shelf    *baseShelf
		index    int
		internal int
	}

	tests := []struct {
		name   string
		fields fields
		want   iterator
	}{
		{
			name: "Case 1",
			fields: fields{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    0,
				internal: 0,
			},
			want: &baseIterator{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    1,
				internal: 1,
			},
		},
		{
			name: "Case 2",
			fields: fields{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    1,
				internal: 1,
			},
			want: &baseIterator{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    2,
				internal: 2,
			},
		},
		{
			name: "Case 3",
			fields: fields{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    2,
				internal: 2,
			},
			want: &baseIterator{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    2,
				internal: 3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := &baseIterator{
				shelf:    tt.fields.shelf,
				index:    tt.fields.index,
				internal: tt.fields.internal,
			}

			iter.Next()

			if !reflect.DeepEqual(iter, tt.want) {
				t.Errorf("Next() = %v, want %v", iter, tt.want)
			}
		})
	}
}

func Test_baseIterator_Prev(t *testing.T) {
	type fields struct {
		shelf    *baseShelf
		index    int
		internal int
	}

	tests := []struct {
		name   string
		fields fields
		want   iterator
	}{
		{
			name: "Case 1",
			fields: fields{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    0,
				internal: 0,
			},
			want: &baseIterator{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    0,
				internal: -1,
			},
		},
		{
			name: "Case 2",
			fields: fields{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    1,
				internal: 1,
			},
			want: &baseIterator{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
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
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    2,
				internal: 2,
			},
			want: &baseIterator{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    1,
				internal: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := &baseIterator{
				shelf:    tt.fields.shelf,
				index:    tt.fields.index,
				internal: tt.fields.internal,
			}
			iter.Prev()

			if !reflect.DeepEqual(iter, tt.want) {
				t.Errorf("Prev() = %v, want %v", iter, tt.want)
			}
		})
	}
}

func Test_baseIterator_Reset(t *testing.T) {
	type fields struct {
		shelf    *baseShelf
		index    int
		internal int
	}

	tests := []struct {
		name   string
		fields fields
		want   iterator
	}{
		{
			name: "Case 1",
			fields: fields{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    0,
				internal: 0,
			},
			want: &baseIterator{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
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
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    1,
				internal: 1,
			},
			want: &baseIterator{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
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
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    2,
				internal: 2,
			},
			want: &baseIterator{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    0,
				internal: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := &baseIterator{
				shelf:    tt.fields.shelf,
				index:    tt.fields.index,
				internal: tt.fields.internal,
			}
			iter.Reset()

			if !reflect.DeepEqual(iter, tt.want) {
				t.Errorf("Reset() = %v, want %v", iter, tt.want)
			}
		})
	}
}

func Test_baseIterator_Value(t *testing.T) {
	type fields struct {
		shelf    *baseShelf
		index    int
		internal int
	}

	tests := []struct {
		name   string
		fields fields
		want   *Task
	}{
		{
			name: "Case 1",
			fields: fields{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    0,
				internal: 0,
			},
			want: &Task{
				Name: "Test task 1",
				Type: 0,
				Func: nil,
			},
		},
		{
			name: "Case 2",
			fields: fields{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    1,
				internal: 1,
			},
			want: &Task{
				Name: "Test task 2",
				Type: 0,
				Func: nil,
			},
		},
		{
			name: "Case 3",
			fields: fields{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    2,
				internal: 2,
			},
			want: &Task{
				Name: "Test task 3",
				Type: 0,
				Func: nil,
			},
		},
		{
			name: "Case 4",
			fields: fields{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    3,
				internal: 3,
			},
			want: nil,
		},
		{
			name: "Case 5",
			fields: fields{
				shelf: &baseShelf{
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
						{
							Name: "Test task 3",
							Type: 0,
							Func: nil,
						},
					},
					rwMx: nil,
				},
				index:    -1,
				internal: -1,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := &baseIterator{
				shelf:    tt.fields.shelf,
				index:    tt.fields.index,
				internal: tt.fields.internal,
			}

			if gotT := iter.Value(); !reflect.DeepEqual(gotT, tt.want) {
				t.Errorf("Value() = %v, want %v", gotT, tt.want)
			}
		})
	}
}
