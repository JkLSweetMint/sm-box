package config

import (
	"sm-box/pkg/core/env"
	"strings"
	"time"
)

// Config - конфигурация компонента ведения журнала трессировки.
type Config struct {
	Terminal *TerminalLog `json:"terminal" yaml:"Terminal" xml:"Terminal"`   // Конфигурация терминала.
	Files    FilesLog     `json:"files"    yaml:"Files"    xml:"Files>File"` // Конфигурация файлов.
}

// FillEmptyFields - заполнение пустых полей конфигурации.
func (conf *Config) FillEmptyFields() *Config {
	if conf.Terminal == nil {
		conf.Terminal = new(TerminalLog)
	}

	if conf.Terminal.Levels == nil {
		conf.Terminal.Levels = new(TerminalLogLevels)
	}

	if conf.Terminal.Options == nil {
		conf.Terminal.Options = new(TerminalLogOptions)
	}

	if strings.TrimSpace(conf.Terminal.Options.TimeFormat) == "" {
		conf.Terminal.Options.TimeFormat = time.RFC3339
	}

	if conf.Files == nil {
		conf.Files = make(FilesLog, 0)
	}

	return conf
}

// Default - запись стандартной конфигурации.
func (conf *Config) Default() *Config {
	// Terminal
	{
		conf.Terminal = &TerminalLog{
			Levels: &TerminalLogLevels{
				Info: &TerminalLogLevel{
					Enable: false,
					Options: &TerminalLogLevelOptions{
						Encoder: "raw",
						Format:  "capital_color",
					},
				},
				Error: &TerminalLogLevel{
					Enable: false,
					Options: &TerminalLogLevelOptions{
						Encoder: "raw",
						Format:  "capital_color",
					},
				},
			},
			Options: &TerminalLogOptions{
				TimeFormat: time.RFC3339,
			},
		}
	}

	// Files
	{
		conf.Files = []*FileLog{
			// Global
			{
				FileName: "%s.trace.log",
				Path:     env.Paths.Var.Logs,

				Options: &FilesLogFileOptions{
					TimeFormat: time.RFC3339,
					Rotation: &FilesLogFileOptionRotation{
						Enable:   true,
						FileSize: "4GB",
					},
				},

				Levels: &FileLogLevels{
					Info: &FileLogLevel{
						Enable: true,
						Options: &FilesLogLevelOptions{
							Encoder: "raw",
							Format:  "capital",
						},
					},
					Error: &FileLogLevel{
						Enable: true,
						Options: &FilesLogLevelOptions{
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
