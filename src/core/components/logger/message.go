package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

// Message - общее описание сообщения компонента ведения журнала.
type Message interface {
	// Initiator - установить значения инициатора вызова.
	Initiator(name string) (msg Message)

	// Text - установить текст сообщения.
	Text(text string) (msg Message)

	// Format - установить текст сообщения с форматированием, под аналогии с fmt.Sprintf.
	Format(format string, a ...any) (msg Message)

	// Field - установить значение поля с дополнительной информацией сообщения.
	Field(key string, val any) (msg Message)

	// Write - запись сообщения в журналы.
	Write()
}

// message - реализация сообщения компонента ведения журнала.
type message struct {
	initiator string
	text      string

	lvl    zapcore.Level
	fields []zap.Field

	write func(lvl zapcore.Level, msg string, fields ...zap.Field)
}

// Initiator - установить значения инициатора вызова.
func (msg *message) Initiator(name string) Message {
	if name = strings.TrimSpace(name); name != "" {
		msg.initiator = name
	}

	return msg
}

// Text - установить текст сообщения.
func (msg *message) Text(text string) Message {
	msg.text = text
	return msg
}

// Format - установить текст сообщения с форматированием, под аналогии с fmt.Sprintf.
func (msg *message) Format(format string, a ...any) Message {
	msg.text = fmt.Sprintf(format, a...)
	return msg
}

// Field - установить значение поля с дополнительной информацией сообщения.
func (msg *message) Field(key string, val any) Message {
	msg.fields = append(msg.fields, zap.Any(key, val))
	return msg
}

// Write - запись сообщения в журналы.
func (msg *message) Write() {
	var initiator string

	if msg.initiator != "" {
		initiator = msg.initiator + ":\t"
	}

	msg.write(msg.lvl, initiator+msg.text, msg.fields...)
}
