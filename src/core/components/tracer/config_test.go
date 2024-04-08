package tracer

import (
	"reflect"
	"sm-box/src/core/components/tracer/logger"
	"testing"
)

func TestConfig_Default(t *testing.T) {
	type fields struct {
		Levels []Level
		Logger *logger.Config
	}

	tests := []struct {
		name   string
		fields fields
		want   *Config
	}{
		{
			name: "Case 1",
			want: &Config{
				Levels: allLevels,
				Logger: new(logger.Config).Default(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{
				Levels: tt.fields.Levels,
				Logger: tt.fields.Logger,
			}

			if got := conf.Default(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Default() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_FillEmptyFields(t *testing.T) {
	type fields struct {
		Levels []Level
		Logger *logger.Config
	}

	tests := []struct {
		name   string
		fields fields
		want   *Config
	}{
		{
			name: "Case 1",
			fields: fields{
				Levels: nil,
				Logger: nil,
			},
			want: &Config{
				Levels: make([]Level, 0),
				Logger: new(logger.Config).FillEmptyFields(),
			},
		},
		{
			name: "Case 2",
			fields: fields{
				Levels: allLevels,
				Logger: new(logger.Config).Default(),
			},
			want: &Config{
				Levels: allLevels,
				Logger: new(logger.Config).Default(),
			},
		},
		{
			name: "Case 3",
			fields: fields{
				Levels: []Level{LevelComponent},
				Logger: nil,
			},
			want: &Config{
				Levels: []Level{LevelComponent},
				Logger: new(logger.Config).FillEmptyFields(),
			},
		},
		{
			name: "Case 4",
			fields: fields{
				Levels: nil,
				Logger: new(logger.Config).Default(),
			},
			want: &Config{
				Levels: []Level{},
				Logger: new(logger.Config).Default(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{
				Levels: tt.fields.Levels,
				Logger: tt.fields.Logger,
			}

			if got := conf.FillEmptyFields(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FillEmptyFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_Validate(t *testing.T) {
	type fields struct {
		Levels []Level
		Logger *logger.Config
	}

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "Case 1. ",
			fields:  fields{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{
				Levels: tt.fields.Levels,
				Logger: tt.fields.Logger,
			}

			if err := conf.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
