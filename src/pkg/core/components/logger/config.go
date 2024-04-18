package logger

import (
	"sm-box/src/pkg/core/components/configurator"
	"sm-box/src/pkg/core/env"
	"strings"
	"time"
)

var confProfile = configurator.PrivateProfile{
	Dir:      "/components",
	Filename: "logger.xml",
}

// Config - конфигурация компонента ведения журнала системы.
type Config struct {
	Terminal *ConfigTerminalLog `json:"terminal" yaml:"Terminal" xml:"Terminal"`   // Конфигурация терминала.
	Files    ConfigFilesLog     `json:"files"    yaml:"Files"    xml:"Files>File"` // Конфигурация файлов.
}

// FillEmptyFields - заполнение обязательных пустых полей конфигурации.
func (conf *Config) FillEmptyFields() *Config {
	if conf.Terminal == nil {
		conf.Terminal = new(ConfigTerminalLog)
	}

	if conf.Terminal.Levels == nil {
		conf.Terminal.Levels = new(ConfigTerminalLogLevels)
	}

	if conf.Terminal.Options == nil {
		conf.Terminal.Options = new(ConfigTerminalLogOptions)
	}

	if strings.TrimSpace(conf.Terminal.Options.TimeFormat) == "" {
		conf.Terminal.Options.TimeFormat = time.RFC3339
	}

	if conf.Files == nil {
		conf.Files = make(ConfigFilesLog, 0)
	}

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *Config) Default() *Config {
	// Terminal
	{
		conf.Terminal = &ConfigTerminalLog{
			Levels: &ConfigTerminalLogLevels{
				Debug: &ConfigTerminalLogLevel{
					Enable: true,
					Options: &ConfigTerminalLogLevelOptions{
						Encoder:     "raw",
						Format:      "capital_color",
						EnableTrace: true,
					},
				},
				Info: &ConfigTerminalLogLevel{
					Enable: true,
					Options: &ConfigTerminalLogLevelOptions{
						Encoder:     "raw",
						Format:      "capital_color",
						EnableTrace: false,
					},
				},
				Warn: &ConfigTerminalLogLevel{
					Enable: true,
					Options: &ConfigTerminalLogLevelOptions{
						Encoder:     "raw",
						Format:      "capital_color",
						EnableTrace: true,
					},
				},
				Error: &ConfigTerminalLogLevel{
					Enable: true,
					Options: &ConfigTerminalLogLevelOptions{
						Encoder:     "raw",
						Format:      "capital_color",
						EnableTrace: true,
					},
				},
				Panic: &ConfigTerminalLogLevel{
					Enable: true,
					Options: &ConfigTerminalLogLevelOptions{
						Encoder:     "raw",
						Format:      "capital_color",
						EnableTrace: true,
					},
				},
				DPanic: &ConfigTerminalLogLevel{
					Enable: true,
					Options: &ConfigTerminalLogLevelOptions{
						Encoder:     "raw",
						Format:      "capital_color",
						EnableTrace: true,
					},
				},
				Fatal: &ConfigTerminalLogLevel{
					Enable: true,
					Options: &ConfigTerminalLogLevelOptions{
						Encoder:     "raw",
						Format:      "capital_color",
						EnableTrace: true,
					},
				},
			},
			Options: &ConfigTerminalLogOptions{
				TimeFormat: time.RFC3339,
			},
		}
	}

	// Files
	{
		conf.Files = []*ConfigFileLog{
			// Global
			{
				FileName: "%s.log",
				Path:     env.Paths.Var.Logs,

				Options: &ConfigFilesLogFileOptions{
					TimeFormat: time.RFC3339,
					Rotation: &ConfigFilesLogFileOptionRotation{
						Enable:   true,
						FileSize: "4GB",
					},
				},

				Levels: &ConfigFileLogLevels{
					Debug: &ConfigFileLogLevel{
						Enable: true,
						Options: &ConfigFilesLogLevelOptions{
							Encoder:     "raw",
							Format:      "capital",
							EnableTrace: true,
						},
					},
					Info: &ConfigFileLogLevel{
						Enable: true,
						Options: &ConfigFilesLogLevelOptions{
							Encoder:     "raw",
							Format:      "capital",
							EnableTrace: false,
						},
					},
					Warn: &ConfigFileLogLevel{
						Enable: true,
						Options: &ConfigFilesLogLevelOptions{
							Encoder:     "raw",
							Format:      "capital",
							EnableTrace: true,
						},
					},
					Error: &ConfigFileLogLevel{
						Enable: true,
						Options: &ConfigFilesLogLevelOptions{
							Encoder:     "raw",
							Format:      "capital",
							EnableTrace: true,
						},
					},
					DPanic: &ConfigFileLogLevel{
						Enable: true,
						Options: &ConfigFilesLogLevelOptions{
							Encoder:     "raw",
							Format:      "capital",
							EnableTrace: true,
						},
					},
					Panic: &ConfigFileLogLevel{
						Enable: true,
						Options: &ConfigFilesLogLevelOptions{
							Encoder:     "raw",
							Format:      "capital",
							EnableTrace: true,
						},
					},
					Fatal: &ConfigFileLogLevel{
						Enable: true,
						Options: &ConfigFilesLogLevelOptions{
							Encoder:     "raw",
							Format:      "capital",
							EnableTrace: true,
						},
					},
				},
			},

			// Debug
			{
				FileName: "%s.debug.log",
				Path:     env.Paths.Var.Logs,

				Options: &ConfigFilesLogFileOptions{
					TimeFormat: time.RFC3339,
					Rotation: &ConfigFilesLogFileOptionRotation{
						Enable:   true,
						FileSize: "4GB",
					},
				},

				Levels: &ConfigFileLogLevels{
					Debug: &ConfigFileLogLevel{
						Enable: true,
						Options: &ConfigFilesLogLevelOptions{
							Encoder:     "raw",
							Format:      "capital",
							EnableTrace: true,
						},
					},
					DPanic: &ConfigFileLogLevel{
						Enable: true,
						Options: &ConfigFilesLogLevelOptions{
							Encoder:     "raw",
							Format:      "capital",
							EnableTrace: true,
						},
					},
				},
			},

			// Error
			{
				FileName: "%s.error.log",
				Path:     env.Paths.Var.Logs,

				Options: &ConfigFilesLogFileOptions{
					TimeFormat: time.RFC3339,
					Rotation: &ConfigFilesLogFileOptionRotation{
						Enable:   true,
						FileSize: "4GB",
					},
				},

				Levels: &ConfigFileLogLevels{
					Error: &ConfigFileLogLevel{
						Enable: true,
						Options: &ConfigFilesLogLevelOptions{
							Encoder:     "raw",
							Format:      "capital",
							EnableTrace: true,
						},
					},
					Panic: &ConfigFileLogLevel{
						Enable: true,
						Options: &ConfigFilesLogLevelOptions{
							Encoder:     "raw",
							Format:      "capital",
							EnableTrace: true,
						},
					},
					Fatal: &ConfigFileLogLevel{
						Enable: true,
						Options: &ConfigFilesLogLevelOptions{
							Encoder:     "raw",
							Format:      "capital",
							EnableTrace: true,
						},
					},
				},
			},
		}
	}

	return conf
}

// Validate - валидация конфигурации.
func (conf *Config) Validate() (err error) {
	return
}
