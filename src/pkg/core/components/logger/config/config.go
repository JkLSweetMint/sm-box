package config

import (
	"sm-box/pkg/core/components/configurator"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	"strings"
	"time"
)

// Config - конфигурация компонента ведения журнала системы.
type Config struct {
	Terminal *TerminalLog `json:"terminal" yaml:"Terminal" xml:"Terminal"`   // Конфигурация терминала.
	Files    FilesLog     `json:"files"    yaml:"Files"    xml:"Files>File"` // Конфигурация файлов.
}

// Read - чтение конфигурации.
func (conf *Config) Read() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var (
		c       configurator.Configurator[*Config]
		profile = configurator.PrivateProfile{
			Dir:      "/components",
			Filename: "logger.xml",
		}
	)

	if c, err = configurator.New[*Config](conf); err != nil {
		return
	} else if err = c.Private().Profile(profile).Init(); err != nil {
		return
	}

	if err = conf.FillEmptyFields().Validate(); err != nil {
		return
	}

	return
}

// FillEmptyFields - заполнение пустых полей конфигурации.
func (conf *Config) FillEmptyFields() *Config {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

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
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(conf) }()
	}

	// Terminal
	{
		conf.Terminal = &TerminalLog{
			Levels: &TerminalLogLevels{
				Debug: &TerminalLogLevel{
					Enable: true,
					Options: &TerminalLogLevelOptions{
						Encoder:     "raw",
						Format:      "capital_color",
						EnableTrace: true,
					},
				},
				Info: &TerminalLogLevel{
					Enable: true,
					Options: &TerminalLogLevelOptions{
						Encoder:     "raw",
						Format:      "capital_color",
						EnableTrace: false,
					},
				},
				Warn: &TerminalLogLevel{
					Enable: true,
					Options: &TerminalLogLevelOptions{
						Encoder:     "raw",
						Format:      "capital_color",
						EnableTrace: true,
					},
				},
				Error: &TerminalLogLevel{
					Enable: true,
					Options: &TerminalLogLevelOptions{
						Encoder:     "raw",
						Format:      "capital_color",
						EnableTrace: true,
					},
				},
				Panic: &TerminalLogLevel{
					Enable: true,
					Options: &TerminalLogLevelOptions{
						Encoder:     "raw",
						Format:      "capital_color",
						EnableTrace: true,
					},
				},
				DPanic: &TerminalLogLevel{
					Enable: true,
					Options: &TerminalLogLevelOptions{
						Encoder:     "raw",
						Format:      "capital_color",
						EnableTrace: true,
					},
				},
				Fatal: &TerminalLogLevel{
					Enable: true,
					Options: &TerminalLogLevelOptions{
						Encoder:     "raw",
						Format:      "capital_color",
						EnableTrace: true,
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
				FileName: "%s.log",
				Path:     env.Paths.Var.Logs,

				Options: &FilesLogFileOptions{
					TimeFormat: time.RFC3339,
					Rotation: &FilesLogFileOptionRotation{
						Enable:   true,
						FileSize: "4GB",
					},
				},

				Levels: &FileLogLevels{
					Debug: &FileLogLevel{
						Enable: true,
						Options: &FilesLogLevelOptions{
							Encoder:     "raw",
							Format:      "capital",
							EnableTrace: true,
						},
					},
					Info: &FileLogLevel{
						Enable: true,
						Options: &FilesLogLevelOptions{
							Encoder:     "raw",
							Format:      "capital",
							EnableTrace: false,
						},
					},
					Warn: &FileLogLevel{
						Enable: true,
						Options: &FilesLogLevelOptions{
							Encoder:     "raw",
							Format:      "capital",
							EnableTrace: true,
						},
					},
					Error: &FileLogLevel{
						Enable: true,
						Options: &FilesLogLevelOptions{
							Encoder:     "raw",
							Format:      "capital",
							EnableTrace: true,
						},
					},
					DPanic: &FileLogLevel{
						Enable: true,
						Options: &FilesLogLevelOptions{
							Encoder:     "raw",
							Format:      "capital",
							EnableTrace: true,
						},
					},
					Panic: &FileLogLevel{
						Enable: true,
						Options: &FilesLogLevelOptions{
							Encoder:     "raw",
							Format:      "capital",
							EnableTrace: true,
						},
					},
					Fatal: &FileLogLevel{
						Enable: true,
						Options: &FilesLogLevelOptions{
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

				Options: &FilesLogFileOptions{
					TimeFormat: time.RFC3339,
					Rotation: &FilesLogFileOptionRotation{
						Enable:   true,
						FileSize: "4GB",
					},
				},

				Levels: &FileLogLevels{
					Debug: &FileLogLevel{
						Enable: true,
						Options: &FilesLogLevelOptions{
							Encoder:     "raw",
							Format:      "capital",
							EnableTrace: true,
						},
					},
					DPanic: &FileLogLevel{
						Enable: true,
						Options: &FilesLogLevelOptions{
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

				Options: &FilesLogFileOptions{
					TimeFormat: time.RFC3339,
					Rotation: &FilesLogFileOptionRotation{
						Enable:   true,
						FileSize: "4GB",
					},
				},

				Levels: &FileLogLevels{
					Error: &FileLogLevel{
						Enable: true,
						Options: &FilesLogLevelOptions{
							Encoder:     "raw",
							Format:      "capital",
							EnableTrace: true,
						},
					},
					Panic: &FileLogLevel{
						Enable: true,
						Options: &FilesLogLevelOptions{
							Encoder:     "raw",
							Format:      "capital",
							EnableTrace: true,
						},
					},
					Fatal: &FileLogLevel{
						Enable: true,
						Options: &FilesLogLevelOptions{
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
	// tracer
	{
		var trc = tracer.New(tracer.LevelConfig)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	return
}
