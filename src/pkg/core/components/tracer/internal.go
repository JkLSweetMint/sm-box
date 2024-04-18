package tracer

import (
	"encoding/json"
	"reflect"
	"sm-box/src/pkg/core/components/tracer/logger"
	"time"
)

const (
	paramFunctionStartTimeStart = iota

	maxParams
)

var config *Config

// tracer - внутренняя реализация компонента.
type tracer struct {
	params    []any
	levels    []Level
	inputArgs []any
	err       error
}

// FunctionCall - начало отслеживание работы функции/метода.
func (t *tracer) FunctionCall(args ...any) {
	var log, _ = logger.New()

	// Проверки
	{
		if log == nil || config == nil {
			return
		}

		if !existLogLevel(config.Levels, t.levels) {
			return
		}
	}

	t.params[paramFunctionStartTimeStart] = time.Now()

	// args
	{
		for i, arg := range args {
			if buff, err := json.Marshal(arg); err != nil {
				args[i] = reflect.TypeOf(arg).String()
			} else {
				args[i] = string(buff)
			}
		}

		t.inputArgs = args
	}

	log.Info().Format("--> |                  | '%s'", functionName()).Field("input", args).Write()
}

// Error - запись ошибки в журнал.
func (t *tracer) Error(err error) Tracer {
	t.err = err
	return t
}

// FunctionCallFinished - конец отслеживания вызова функции/метода.
func (t *tracer) FunctionCallFinished(args ...any) {
	var log, _ = logger.New()

	// Проверки
	{
		if log == nil || config == nil {
			return
		}

		if !existLogLevel(config.Levels, t.levels) {
			return
		}
	}

	// outputArgs
	{
		for i, arg := range args {
			if buff, err := json.Marshal(arg); err != nil {
				args[i] = reflect.TypeOf(arg).String()
			} else {
				args[i] = string(buff)
			}
		}
	}

	if t.err != nil {
		var workTm string

		if tm, ok := t.params[paramFunctionStartTimeStart].(time.Time); ok {
			workTm = formatWorkTime(time.Now().Sub(tm))
		}

		log.Error().Format("<-x %s '%s'", workTm, functionName()).
			Field("error", t.err).
			Field("output", args).Write()

	} else {
		var workTm string

		if tm, ok := t.params[paramFunctionStartTimeStart].(time.Time); ok {
			workTm = formatWorkTime(time.Now().Sub(tm))
		}

		log.Info().Format("<-- %s '%s'", workTm, functionName()).Field("output", args).Write()
	}
}
