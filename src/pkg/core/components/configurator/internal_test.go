package configurator

import (
	"reflect"
	"testing"
)

func Test_configurator_Private(t *testing.T) {
	type testCase[T any] struct {
		name string
		c    configurator[T]
		want Private[*TestConfig]
	}

	tests := []testCase[*TestConfig]{
		{
			name: "Case 1",
			c: configurator[*TestConfig]{
				conf: new(TestConfig),
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  prtDefaultEncoder,
				dir:      "",
				filename: "",
			},
		},
		{
			name: "Case 2",
			c: configurator[*TestConfig]{
				conf: new(TestConfig).Default(),
			},
			want: &privateConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  prtDefaultEncoder,
				dir:      "",
				filename: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Private(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Private() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_configurator_Public(t *testing.T) {
	type testCase[T any] struct {
		name string
		c    configurator[T]
		want Public[*TestConfig]
	}

	tests := []testCase[*TestConfig]{
		{
			name: "Case 1",
			c: configurator[*TestConfig]{
				conf: new(TestConfig),
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig),
				encoder:  pbDefaultEncoder,
				dir:      "",
				filename: "",
			},
		},
		{
			name: "Case 2",
			c: configurator[*TestConfig]{
				conf: new(TestConfig).Default(),
			},
			want: &publicConfigurator[*TestConfig]{
				conf:     new(TestConfig).Default(),
				encoder:  pbDefaultEncoder,
				dir:      "",
				filename: "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Public(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Public() = %v, want %v", got, tt.want)
			}
		})
	}
}
