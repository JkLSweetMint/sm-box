package synchronization

import (
	"reflect"
	"sync"
	"testing"
)

func TestDev_Build(t *testing.T) {
	type fields struct {
		WaitGroup *sync.WaitGroup
	}

	tests := []struct {
		name   string
		fields fields
		want   *Dev
	}{
		{
			name:   "Case 1",
			fields: fields{},
			want: &Dev{
				WaitGroup: new(sync.WaitGroup),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &Dev{
				WaitGroup: tt.fields.WaitGroup,
			}

			if got := storage.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProd_Build(t *testing.T) {
	type fields struct {
		WaitGroup *sync.WaitGroup
	}

	tests := []struct {
		name   string
		fields fields
		want   *Prod
	}{
		{
			name:   "Case 1",
			fields: fields{},
			want: &Prod{
				WaitGroup: new(sync.WaitGroup),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := &Prod{
				WaitGroup: tt.fields.WaitGroup,
			}

			if got := storage.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
