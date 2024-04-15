package server

import (
	"reflect"
	"testing"
)

func TestConfig_Default(t *testing.T) {
	type fields struct{}

	tests := []struct {
		name   string
		fields fields
		want   *Config
	}{
		{
			name:   "Case 1",
			fields: fields{},
			want:   &Config{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{}

			if got := conf.Default(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Default() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_FillEmptyFields(t *testing.T) {
	type fields struct{}

	tests := []struct {
		name   string
		fields fields
		want   *Config
	}{
		{
			name:   "Case 1",
			fields: fields{},
			want:   &Config{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{}

			if got := conf.FillEmptyFields(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FillEmptyFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_Validate(t *testing.T) {
	type fields struct{}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "Case 1",
			fields:  fields{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{}

			if err := conf.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
