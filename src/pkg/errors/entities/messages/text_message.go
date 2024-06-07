package messages

import (
	"fmt"
	"sm-box/pkg/errors/types"
	"strings"
)

// TextMessage - текстовое сообщение.
type TextMessage struct {
	content string
}

// String - получение строкового представления сообщения.
func (m *TextMessage) String() (str string) {
	return m.content
}

// Text - установить текст сообщения.
func (m *TextMessage) Text(content string) *TextMessage {
	m.content = content
	return m
}

// Format - установить текст сообщения с форматированием, по аналогии с fmt.Sprintf.
func (m *TextMessage) Format(format string, a ...any) *TextMessage {
	m.content = fmt.Sprintf(format, a...)
	return m
}

// Clone - копирование сообщения.
func (m *TextMessage) Clone() types.Message {
	return &TextMessage{
		content: strings.Clone(m.content),
	}
}
