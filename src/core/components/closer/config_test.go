package closer

import (
	"os"
	"reflect"
	"testing"
)

func TestConfig_Default(t *testing.T) {
	type fields struct {
		Signals []os.Signal
	}

	tests := []struct {
		name   string
		fields fields
		want   *Config
	}{
		{
			name: "Case 1",
			want: &Config{
				Signals: defaultSignals,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{
				Signals: tt.fields.Signals,
			}

			if got := conf.Default(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Default() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_FillEmptyFields(t *testing.T) {
	type fields struct {
		Signals []os.Signal
	}

	tests := []struct {
		name   string
		fields fields
		want   *Config
	}{
		{
			name: "Case 1",
			want: &Config{
				Signals: make([]os.Signal, 0),
			},
		},
		{
			name: "Case 2",
			fields: fields{
				Signals: defaultSignals,
			},
			want: &Config{
				Signals: defaultSignals,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{
				Signals: tt.fields.Signals,
			}

			if got := conf.FillEmptyFields(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FillEmptyFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_Validate(t *testing.T) {
	type fields struct {
		Signals []os.Signal
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{

		{
			name: "Case 1",
			fields: fields{
				Signals: make([]os.Signal, 0),
			},
			wantErr: false,
		},
		{
			name: "Case 2",
			fields: fields{
				Signals: defaultSignals,
			},
			wantErr: false,
		},
		{
			name:    "Case 3",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{
				Signals: tt.fields.Signals,
			}

			if err := conf.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
