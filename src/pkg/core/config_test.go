package core

import (
	"reflect"
	"sm-box/src/pkg/core/tools/closer"
	"testing"
)

func TestConfigTools_Default(t *testing.T) {
	type fields struct {
		Closer *closer.Config
	}

	tests := []struct {
		name   string
		fields fields
		want   *ConfigTools
	}{
		{
			name: "Case 1",
			fields: fields{
				Closer: nil,
			},
			want: &ConfigTools{
				Closer: new(closer.Config).Default(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &ConfigTools{
				Closer: tt.fields.Closer,
			}

			if got := conf.Default(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Default() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigTools_FillEmptyFields(t *testing.T) {
	type fields struct {
		Closer *closer.Config
	}

	tests := []struct {
		name   string
		fields fields
		want   *ConfigTools
	}{
		{
			name: "Case 1",
			fields: fields{
				Closer: nil,
			},
			want: &ConfigTools{
				Closer: new(closer.Config).FillEmptyFields(),
			},
		},
		{
			name: "Case 2",
			fields: fields{
				Closer: new(closer.Config).Default(),
			},
			want: &ConfigTools{
				Closer: new(closer.Config).Default(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &ConfigTools{
				Closer: tt.fields.Closer,
			}

			if got := conf.FillEmptyFields(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FillEmptyFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigTools_Validate(t *testing.T) {
	type fields struct {
		Closer *closer.Config
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Case 1",
			fields: fields{
				Closer: nil,
			},
			wantErr: false,
		},
		{
			name: "Case 2",
			fields: fields{
				Closer: new(closer.Config).Default(),
			},
			wantErr: false,
		},
		{
			name: "Case 3",
			fields: fields{
				Closer: new(closer.Config).FillEmptyFields(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &ConfigTools{
				Closer: tt.fields.Closer,
			}

			if err := conf.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfig_Default(t *testing.T) {
	type fields struct {
		Tools *ConfigTools
	}

	tests := []struct {
		name   string
		fields fields
		want   *Config
	}{

		{
			name: "Case 1",
			fields: fields{
				Tools: &ConfigTools{
					Closer: new(closer.Config).Default(),
				},
			},
			want: &Config{
				Tools: &ConfigTools{
					Closer: new(closer.Config).Default(),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{
				Tools: tt.fields.Tools,
			}

			if got := conf.Default(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Default() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_FillEmptyFields(t *testing.T) {
	type fields struct {
		Tools *ConfigTools
	}

	tests := []struct {
		name   string
		fields fields
		want   *Config
	}{

		{
			name: "Case 1",
			fields: fields{
				Tools: nil,
			},
			want: &Config{
				Tools: &ConfigTools{
					Closer: new(closer.Config).FillEmptyFields(),
				},
			},
		},
		{
			name: "Case 2",
			fields: fields{
				Tools: &ConfigTools{
					Closer: new(closer.Config).Default(),
				},
			},
			want: &Config{
				Tools: &ConfigTools{
					Closer: new(closer.Config).Default(),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{
				Tools: tt.fields.Tools,
			}

			if got := conf.FillEmptyFields(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FillEmptyFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_Validate(t *testing.T) {
	type fields struct {
		Tools *ConfigTools
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Case 1",
			fields: fields{
				Tools: nil,
			},
			wantErr: false,
		},
		{
			name: "Case 2",
			fields: fields{
				Tools: &ConfigTools{
					Closer: new(closer.Config).Default(),
				},
			},
			wantErr: false,
		},
		{
			name: "Case 3",
			fields: fields{
				Tools: &ConfigTools{
					Closer: new(closer.Config).FillEmptyFields(),
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{
				Tools: tt.fields.Tools,
			}

			if err := conf.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
