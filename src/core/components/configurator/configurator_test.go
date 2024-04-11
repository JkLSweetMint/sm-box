package configurator

import (
	"reflect"
	"testing"
)

type TestConfig struct {
	Elements []*TestConfigElement `json:"elements" yaml:"Elements" xml:"Elements>Element"`
	Date     string               `json:"date"     yaml:"Date"     xml:"date,attr"`
}

type TestConfigElement struct {
	Value string `json:"value" yaml:"Value" xml:"value,attr"`
}

func (conf *TestConfig) FillEmptyFields() *TestConfig {
	if conf.Elements == nil {
		conf.Elements = make([]*TestConfigElement, 0)
	}

	return conf
}

func (conf *TestConfig) Default() *TestConfig {
	conf.Elements = []*TestConfigElement{
		{Value: "1"},
		{Value: "2"},
		{Value: "3"},
		{Value: "4"},
		{Value: "5"},
	}

	conf.Date = "0001-01-01"

	return conf
}

func (conf *TestConfig) Validate() (err error) {
	return
}

func TestNew(t *testing.T) {
	type args[T any] struct {
		conf Config[*TestConfig]
	}

	type testCase[T any] struct {
		name    string
		args    args[T]
		wantC   Configurator[*TestConfig]
		wantErr bool
	}

	tests := []testCase[*TestConfig]{
		{
			name: "Case 1",
			args: args[*TestConfig]{
				conf: new(TestConfig),
			},
			wantC: &configurator[*TestConfig]{
				conf: new(TestConfig),
			},
			wantErr: false,
		},
		{
			name: "Case 2",
			args: args[*TestConfig]{
				conf: new(TestConfig).Default(),
			},
			wantC: &configurator[*TestConfig]{
				conf: new(TestConfig).Default(),
			},
			wantErr: false,
		},
		{
			name: "Case 3",
			args: args[*TestConfig]{
				conf: nil,
			},
			wantC:   nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotC, err := New(tt.args.conf)

			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("New() gotC = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
