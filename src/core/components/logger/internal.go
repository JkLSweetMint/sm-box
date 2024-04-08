package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"path"
	"sm-box/src/core/env"
	"strings"
)

// logger - компонент ведения журнала системы.
type logger struct {
	initiator string
	*internal
}

// internal - внутренняя реалихация компонента ведения журнала системы.
type internal struct {
	conf     *Config
	instance *zap.Logger
	sources  []io.WriteCloser
}

// newLogger - создание компонента ведения журнала.
func newLogger(initiator string, conf *Config) (log *logger, err error) {
	if initiator = strings.TrimSpace(initiator); initiator == "" {
		initiator = "unknown"
	}

	if conf == nil {
		conf = new(Config).Default()
	}

	log = &logger{
		initiator: initiator,
		internal: &internal{
			conf:    conf,
			sources: make([]io.WriteCloser, 0),
		},
	}

	if err = log.conf.FillEmptyFields().Validate(); err != nil {
		return
	}

	// zap
	{
		var cores = make([]zapcore.Core, 0)

		// Terminal
		{
			var (
				levels = []zapcore.Level{
					zapcore.DebugLevel,
					zapcore.InfoLevel,
					zapcore.WarnLevel,
					zapcore.ErrorLevel,
					zapcore.DPanicLevel,
					zapcore.PanicLevel,
					zapcore.FatalLevel,
				}
				configs = []*ConfigTerminalLogLevel{
					log.conf.Terminal.Levels.Debug,
					log.conf.Terminal.Levels.Info,
					log.conf.Terminal.Levels.Warn,
					log.conf.Terminal.Levels.Error,
					log.conf.Terminal.Levels.DPanic,
					log.conf.Terminal.Levels.Panic,
					log.conf.Terminal.Levels.Fatal,
				}
			)

			for i, c := range configs {
				if c == nil || !c.Enable {
					continue
				}

				var (
					lvl         = levels[i]
					encoderConf = zapcore.EncoderConfig{
						MessageKey:    "message",
						LevelKey:      "level",
						TimeKey:       "time",
						NameKey:       "name",
						CallerKey:     "trace",
						FunctionKey:   "func",
						StacktraceKey: "stack_trace",

						LineEnding: zapcore.DefaultLineEnding,
						EncodeTime: zapcore.TimeEncoderOfLayout(log.conf.Terminal.Options.TimeFormat),
					}
				)

				switch strings.ToLower(c.Options.Format) {
				case "capital":
					encoderConf.EncodeLevel = zapcore.CapitalLevelEncoder
				case "capital_color":
					encoderConf.EncodeLevel = zapcore.CapitalColorLevelEncoder
				case "lowercase":
					encoderConf.EncodeLevel = zapcore.LowercaseLevelEncoder
				case "lowercase_color":
					encoderConf.EncodeLevel = zapcore.LowercaseColorLevelEncoder
				}

				var (
					zapFn zap.LevelEnablerFunc = func(level zapcore.Level) bool {
						return level == lvl
					}
					encoder zapcore.Encoder
				)

				switch strings.ToLower(c.Options.Encoder) {
				case "raw":
					encoder = zapcore.NewConsoleEncoder(encoderConf)
				case "json":
					encoder = zapcore.NewJSONEncoder(encoderConf)
				}

				cores = append(cores,
					zapcore.NewCore(
						encoder,
						zapcore.Lock(os.Stdout),
						zapFn,
					),
				)
			}
		}

		// Files
		{
			for _, flConf := range log.conf.Files {
				var (
					levels = []zapcore.Level{
						zapcore.DebugLevel,
						zapcore.InfoLevel,
						zapcore.WarnLevel,
						zapcore.ErrorLevel,
						zapcore.DPanicLevel,
						zapcore.PanicLevel,
						zapcore.FatalLevel,
					}
					configs = []*ConfigFileLogLevel{
						flConf.Levels.Debug,
						flConf.Levels.Info,
						flConf.Levels.Warn,
						flConf.Levels.Error,
						flConf.Levels.DPanic,
						flConf.Levels.Panic,
						flConf.Levels.Fatal,
					}
				)

				for i, c := range configs {
					if c == nil || !c.Enable {
						continue
					}

					var (
						lvl         = levels[i]
						encoderConf = zapcore.EncoderConfig{
							MessageKey:    "message",
							LevelKey:      "level",
							TimeKey:       "time",
							NameKey:       "name",
							CallerKey:     "trace",
							FunctionKey:   "func",
							StacktraceKey: "stack_trace",

							LineEnding: zapcore.DefaultLineEnding,
							EncodeTime: zapcore.TimeEncoderOfLayout(log.conf.Terminal.Options.TimeFormat),
						}
					)

					switch strings.ToLower(c.Options.Format) {
					case "capital":
						encoderConf.EncodeLevel = zapcore.CapitalLevelEncoder
					case "capital_color":
						encoderConf.EncodeLevel = zapcore.CapitalColorLevelEncoder
					case "lowercase":
						encoderConf.EncodeLevel = zapcore.LowercaseLevelEncoder
					case "lowercase_color":
						encoderConf.EncodeLevel = zapcore.LowercaseColorLevelEncoder
					}

					var (
						writeSyncer zapcore.WriteSyncer
						zapFn       zap.LevelEnablerFunc = func(level zapcore.Level) bool {
							return level == lvl
						}
						encoder zapcore.Encoder
					)

					switch strings.ToLower(c.Options.Encoder) {
					case "raw":
						encoder = zapcore.NewConsoleEncoder(encoderConf)
					case "json":
						encoder = zapcore.NewJSONEncoder(encoderConf)
					}

					if flConf.Options != nil && flConf.Options.Rotation != nil && flConf.Options.Rotation.Enable {
						var rotator *rotatelogs.RotateLogs

						if rotator, err = rotatelogs.New(
							path.Join(env.Paths.SystemLocation, flConf.Path, flConf.FileName),
							rotatelogs.WithRotationSize(flConf.Options.Rotation.FileSize.Int64())); err != nil {
							return
						}

						writeSyncer = zapcore.AddSync(rotator)
						log.sources = append(log.sources, rotator)
					} else {
						var fl *os.File

						if fl, err = os.OpenFile(path.Join(env.Paths.SystemLocation, flConf.Path, flConf.FileName), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err != nil {
							return
						}

						writeSyncer = zapcore.Lock(fl)
						log.sources = append(log.sources, fl)
					}

					cores = append(cores,
						zapcore.NewCore(
							encoder,
							writeSyncer,
							zapFn,
						),
					)
				}
			}
		}

		log.instance = zap.New(zapcore.NewTee(cores...))
	}

	return
}

// Debug - создает сообщение 'Debug' уровня журнала.
func (log *logger) Debug() (msg Message) {
	if log.instance == nil {
		return
	}

	var m = &message{
		initiator: log.initiator,
		text:      "",

		lvl:    zap.DebugLevel,
		fields: nil,

		write: log.instance.Log,
	}

	msg = m

	return
}

// Info - создает сообщение 'Info' уровня журнала.
func (log *logger) Info() (msg Message) {
	if log.instance == nil {
		return
	}

	var m = &message{
		initiator: log.initiator,
		text:      "",

		lvl:    zap.InfoLevel,
		fields: nil,

		write: log.instance.Log,
	}

	msg = m

	return
}

// Warn - создает сообщение 'Warn' уровня журнала.
func (log *logger) Warn() (msg Message) {
	if log.instance == nil {
		return
	}

	var m = &message{
		initiator: log.initiator,
		text:      "",

		lvl:    zap.WarnLevel,
		fields: nil,

		write: log.instance.Log,
	}

	msg = m

	return
}

// Error - создает сообщение 'Error' уровня журнала.
func (log *logger) Error() (msg Message) {
	if log.instance == nil {
		return
	}

	var m = &message{
		initiator: log.initiator,
		text:      "",

		lvl:    zap.ErrorLevel,
		fields: nil,

		write: log.instance.Log,
	}

	msg = m

	return
}

// Panic - создает сообщение 'Panic' уровня журнала.
//
// Вызов всегда будет завершаться с вызовом паники.
func (log *logger) Panic() (msg Message) {
	if log.instance == nil {
		return
	}

	var m = &message{
		initiator: log.initiator,
		text:      "",

		lvl:    zap.PanicLevel,
		fields: nil,

		write: log.instance.Log,
	}

	msg = m

	return
}

// DPanic - создает сообщение 'DPanic' уровня журнала.
//
// Вызов завершится с вызовом паники если уровень журнала стоит на 'Debug'.
func (log *logger) DPanic() (msg Message) {
	if log.instance == nil {
		return
	}

	var m = &message{
		initiator: log.initiator,
		text:      "",

		lvl:    zap.DPanicLevel,
		fields: nil,

		write: log.instance.Log,
	}

	msg = m

	return
}

// Fatal - создает сообщение 'Fatal' уровня журнала.
//
// Вызов всегда будет завершаться с вызовом os.Exit(1).
func (log *logger) Fatal() (msg Message) {
	if log.instance == nil {
		return
	}

	var m = &message{
		initiator: log.initiator,
		text:      "",

		lvl:    zap.FatalLevel,
		fields: nil,

		write: log.instance.Log,
	}

	msg = m

	return
}

// Close - закрытие компонента и журналов.
func (log *logger) Close() (err error) {
	log.instance = nil

	for _, src := range log.sources {
		if err = src.Close(); err != nil {
			return
		}
	}

	return
}

// Copy - копирование компонента и журналов для инициатора.
func (log *logger) Copy(initiator string) Logger {
	if initiator = strings.TrimSpace(initiator); initiator == "" {
		initiator = "unknown"
	}

	return &logger{
		initiator: initiator,
		internal:  log.internal,
	}
}
