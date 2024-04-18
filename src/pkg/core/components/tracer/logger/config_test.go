package logger

import (
	"reflect"
	"sm-box/src/pkg/core/env"
	"testing"
	"time"
)

func TestConfig_Default(t *testing.T) {
	type fields struct {
		Terminal *ConfigTerminalLog
		Files    ConfigFilesLog
	}

	tests := []struct {
		name   string
		fields fields
		want   *Config
	}{
		{
			name: "Case 1. ",
			want: &Config{
				Terminal: &ConfigTerminalLog{
					Levels: &ConfigTerminalLogLevels{
						Info: &ConfigTerminalLogLevel{
							Enable: true,
							Options: &ConfigTerminalLogLevelOptions{
								Encoder: "raw",
								Format:  "capital_color",
							},
						},
						Error: &ConfigTerminalLogLevel{
							Enable: true,
							Options: &ConfigTerminalLogLevelOptions{
								Encoder: "raw",
								Format:  "capital_color",
							},
						},
					},
					Options: &ConfigTerminalLogOptions{
						TimeFormat: time.RFC3339,
					},
				},

				Files: []*ConfigFileLog{
					// Global
					{
						FileName: "%s.trace.log",
						Path:     env.Paths.Var.Logs,

						Options: &ConfigFilesLogFileOptions{
							TimeFormat: time.RFC3339,
							Rotation: &ConfigFilesLogFileOptionRotation{
								Enable:   true,
								FileSize: "4GB",
							},
						},

						Levels: &ConfigFileLogLevels{
							Info: &ConfigFileLogLevel{
								Enable: true,
								Options: &ConfigFilesLogLevelOptions{
									Encoder: "raw",
									Format:  "capital",
								},
							},
							Error: &ConfigFileLogLevel{
								Enable: true,
								Options: &ConfigFilesLogLevelOptions{
									Encoder: "raw",
									Format:  "capital",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{
				Terminal: tt.fields.Terminal,
				Files:    tt.fields.Files,
			}

			if got := conf.Default(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Default() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_FillEmptyFields(t *testing.T) {
	type fields struct {
		Terminal *ConfigTerminalLog
		Files    ConfigFilesLog
	}

	tests := []struct {
		name   string
		fields fields
		want   *Config
	}{
		{
			name: "Case 1. ",
			fields: fields{
				Terminal: nil,
				Files:    nil,
			},
			want: &Config{
				Terminal: &ConfigTerminalLog{
					Levels: new(ConfigTerminalLogLevels),
					Options: &ConfigTerminalLogOptions{
						TimeFormat: time.RFC3339,
					},
				},
				Files: make(ConfigFilesLog, 0),
			},
		},
		{
			name: "Case 2. ",
			fields: fields{
				Terminal: &ConfigTerminalLog{
					Levels: &ConfigTerminalLogLevels{
						Info: &ConfigTerminalLogLevel{
							Enable: true,
							Options: &ConfigTerminalLogLevelOptions{
								Encoder: "raw",
								Format:  "capital_color",
							},
						},
						Error: &ConfigTerminalLogLevel{
							Enable: true,
							Options: &ConfigTerminalLogLevelOptions{
								Encoder: "raw",
								Format:  "capital_color",
							},
						},
					},
					Options: &ConfigTerminalLogOptions{
						TimeFormat: time.RFC3339,
					},
				},
			},
			want: &Config{
				Terminal: &ConfigTerminalLog{
					Levels: &ConfigTerminalLogLevels{
						Info: &ConfigTerminalLogLevel{
							Enable: true,
							Options: &ConfigTerminalLogLevelOptions{
								Encoder: "raw",
								Format:  "capital_color",
							},
						},
						Error: &ConfigTerminalLogLevel{
							Enable: true,
							Options: &ConfigTerminalLogLevelOptions{
								Encoder: "raw",
								Format:  "capital_color",
							},
						},
					},
					Options: &ConfigTerminalLogOptions{
						TimeFormat: time.RFC3339,
					},
				},
				Files: make(ConfigFilesLog, 0),
			},
		},
		{
			name: "Case 3. ",
			fields: fields{
				Files: []*ConfigFileLog{
					// Global
					{
						FileName: "%s.trace.log",
						Path:     env.Paths.Var.Logs,

						Options: &ConfigFilesLogFileOptions{
							TimeFormat: time.RFC3339,
							Rotation: &ConfigFilesLogFileOptionRotation{
								Enable:   true,
								FileSize: "4GB",
							},
						},

						Levels: &ConfigFileLogLevels{
							Info: &ConfigFileLogLevel{
								Enable: true,
								Options: &ConfigFilesLogLevelOptions{
									Encoder: "raw",
									Format:  "capital",
								},
							},
							Error: &ConfigFileLogLevel{
								Enable: true,
								Options: &ConfigFilesLogLevelOptions{
									Encoder: "raw",
									Format:  "capital",
								},
							},
						},
					},
				},
			},
			want: &Config{
				Terminal: &ConfigTerminalLog{
					Levels: new(ConfigTerminalLogLevels),
					Options: &ConfigTerminalLogOptions{
						TimeFormat: time.RFC3339,
					},
				},
				Files: []*ConfigFileLog{
					// Global
					{
						FileName: "%s.trace.log",
						Path:     env.Paths.Var.Logs,

						Options: &ConfigFilesLogFileOptions{
							TimeFormat: time.RFC3339,
							Rotation: &ConfigFilesLogFileOptionRotation{
								Enable:   true,
								FileSize: "4GB",
							},
						},

						Levels: &ConfigFileLogLevels{
							Info: &ConfigFileLogLevel{
								Enable: true,
								Options: &ConfigFilesLogLevelOptions{
									Encoder: "raw",
									Format:  "capital",
								},
							},
							Error: &ConfigFileLogLevel{
								Enable: true,
								Options: &ConfigFilesLogLevelOptions{
									Encoder: "raw",
									Format:  "capital",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Case 4. ",
			want: &Config{
				Terminal: &ConfigTerminalLog{
					Levels: new(ConfigTerminalLogLevels),
					Options: &ConfigTerminalLogOptions{
						TimeFormat: time.RFC3339,
					},
				},
				Files: make(ConfigFilesLog, 0),
			},
		},
		{
			name: "Case 5. ",
			fields: fields{
				Terminal: &ConfigTerminalLog{
					Levels: &ConfigTerminalLogLevels{
						Info: &ConfigTerminalLogLevel{
							Enable: true,
							Options: &ConfigTerminalLogLevelOptions{
								Encoder: "raw",
								Format:  "capital_color",
							},
						},
						Error: &ConfigTerminalLogLevel{
							Enable: true,
							Options: &ConfigTerminalLogLevelOptions{
								Encoder: "raw",
								Format:  "capital_color",
							},
						},
					},
					Options: &ConfigTerminalLogOptions{
						TimeFormat: time.RFC3339,
					},
				},
				Files: []*ConfigFileLog{
					// Global
					{
						FileName: "%s.trace.log",
						Path:     env.Paths.Var.Logs,

						Options: &ConfigFilesLogFileOptions{
							TimeFormat: time.RFC3339,
							Rotation: &ConfigFilesLogFileOptionRotation{
								Enable:   true,
								FileSize: "4GB",
							},
						},

						Levels: &ConfigFileLogLevels{
							Info: &ConfigFileLogLevel{
								Enable: true,
								Options: &ConfigFilesLogLevelOptions{
									Encoder: "raw",
									Format:  "capital",
								},
							},
							Error: &ConfigFileLogLevel{
								Enable: true,
								Options: &ConfigFilesLogLevelOptions{
									Encoder: "raw",
									Format:  "capital",
								},
							},
						},
					},
				},
			},
			want: &Config{
				Terminal: &ConfigTerminalLog{
					Levels: &ConfigTerminalLogLevels{
						Info: &ConfigTerminalLogLevel{
							Enable: true,
							Options: &ConfigTerminalLogLevelOptions{
								Encoder: "raw",
								Format:  "capital_color",
							},
						},
						Error: &ConfigTerminalLogLevel{
							Enable: true,
							Options: &ConfigTerminalLogLevelOptions{
								Encoder: "raw",
								Format:  "capital_color",
							},
						},
					},
					Options: &ConfigTerminalLogOptions{
						TimeFormat: time.RFC3339,
					},
				},
				Files: []*ConfigFileLog{
					// Global
					{
						FileName: "%s.trace.log",
						Path:     env.Paths.Var.Logs,

						Options: &ConfigFilesLogFileOptions{
							TimeFormat: time.RFC3339,
							Rotation: &ConfigFilesLogFileOptionRotation{
								Enable:   true,
								FileSize: "4GB",
							},
						},

						Levels: &ConfigFileLogLevels{
							Info: &ConfigFileLogLevel{
								Enable: true,
								Options: &ConfigFilesLogLevelOptions{
									Encoder: "raw",
									Format:  "capital",
								},
							},
							Error: &ConfigFileLogLevel{
								Enable: true,
								Options: &ConfigFilesLogLevelOptions{
									Encoder: "raw",
									Format:  "capital",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{
				Terminal: tt.fields.Terminal,
				Files:    tt.fields.Files,
			}

			got := conf.FillEmptyFields()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FillEmptyFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_Validate(t *testing.T) {
	type fields struct {
		Terminal *ConfigTerminalLog
		Files    ConfigFilesLog
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
		{
			name: "Case 2. ",
			fields: fields{
				Terminal: &ConfigTerminalLog{
					Levels: &ConfigTerminalLogLevels{
						Info: &ConfigTerminalLogLevel{
							Enable: true,
							Options: &ConfigTerminalLogLevelOptions{
								Encoder: "raw",
								Format:  "capital_color",
							},
						},
						Error: &ConfigTerminalLogLevel{
							Enable: true,
							Options: &ConfigTerminalLogLevelOptions{
								Encoder: "raw",
								Format:  "capital_color",
							},
						},
					},
					Options: &ConfigTerminalLogOptions{
						TimeFormat: time.RFC3339,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Case 3. ",
			fields: fields{
				Files: []*ConfigFileLog{
					// Global
					{
						FileName: "%s.trace.log",
						Path:     env.Paths.Var.Logs,

						Options: &ConfigFilesLogFileOptions{
							TimeFormat: time.RFC3339,
							Rotation: &ConfigFilesLogFileOptionRotation{
								Enable:   true,
								FileSize: "4GB",
							},
						},

						Levels: &ConfigFileLogLevels{
							Info: &ConfigFileLogLevel{
								Enable: true,
								Options: &ConfigFilesLogLevelOptions{
									Encoder: "raw",
									Format:  "capital",
								},
							},
							Error: &ConfigFileLogLevel{
								Enable: true,
								Options: &ConfigFilesLogLevelOptions{
									Encoder: "raw",
									Format:  "capital",
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name:    "Case 4. ",
			wantErr: false,
		},
		{
			name: "Case 5. ",
			fields: fields{
				Terminal: &ConfigTerminalLog{
					Levels: &ConfigTerminalLogLevels{
						Info: &ConfigTerminalLogLevel{
							Enable: true,
							Options: &ConfigTerminalLogLevelOptions{
								Encoder: "raw",
								Format:  "capital_color",
							},
						},
						Error: &ConfigTerminalLogLevel{
							Enable: true,
							Options: &ConfigTerminalLogLevelOptions{
								Encoder: "raw",
								Format:  "capital_color",
							},
						},
					},
					Options: &ConfigTerminalLogOptions{
						TimeFormat: time.RFC3339,
					},
				},
				Files: []*ConfigFileLog{
					// Global
					{
						FileName: "%s.trace.log",
						Path:     env.Paths.Var.Logs,

						Options: &ConfigFilesLogFileOptions{
							TimeFormat: time.RFC3339,
							Rotation: &ConfigFilesLogFileOptionRotation{
								Enable:   true,
								FileSize: "4GB",
							},
						},

						Levels: &ConfigFileLogLevels{
							Info: &ConfigFileLogLevel{
								Enable: true,
								Options: &ConfigFilesLogLevelOptions{
									Encoder: "raw",
									Format:  "capital",
								},
							},
							Error: &ConfigFileLogLevel{
								Enable: true,
								Options: &ConfigFilesLogLevelOptions{
									Encoder: "raw",
									Format:  "capital",
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{
				Terminal: tt.fields.Terminal,
				Files:    tt.fields.Files,
			}

			if err := conf.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
