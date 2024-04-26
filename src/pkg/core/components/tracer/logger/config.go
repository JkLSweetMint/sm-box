package logger

import (
	"sm-box/pkg/core/env"
	"strings"
	"time"
)

// Config - конфигурация компонента ведения журнала трессировки.
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
				Info: &ConfigTerminalLogLevel{
					Enable: false,
					Options: &ConfigTerminalLogLevelOptions{
						Encoder: "raw",
						Format:  "capital_color",
					},
				},
				Error: &ConfigTerminalLogLevel{
					Enable: false,
					Options: &ConfigTerminalLogLevelOptions{
						Encoder: "raw",
						Format:  "capital_color",
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
		}
	}

	return conf
}

// Validate - валидация конфигурации.
func (conf *Config) Validate() (err error) {
	return
}
