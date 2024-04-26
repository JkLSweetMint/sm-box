package format

import (
	"fmt"
	"strings"
)

// TextOption - опции для форматирования текста.
type TextOption struct {
	Key   string
	Value any
}

// Text - форматирование текста.
func Text(t string, opts ...TextOption) (text string) {
	text = t

	for _, opt := range opts {
		text = strings.Replace(text, fmt.Sprintf("{{%s}}", opt.Key), fmt.Sprintf("%+v", opt.Value), -1)
	}

	return
}
