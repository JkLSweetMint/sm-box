package postman

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Description - содержит описание коллекции.
type Description struct {
	Content string `json:"content,omitempty"`
	Type    string `json:"type,omitempty"`
	Version string `json:"version,omitempty"`
}

// mDescription - используется для сортировки/разгрузки данных.
type mDescription Description

// MarshalJSON - возвращает кодировку описания в формате JSON.
// Если Описание содержит только содержимое, оно возвращается в виде строки.
func (d Description) MarshalJSON() ([]byte, error) {
	if d.Type == "" && d.Version == "" {
		return []byte(fmt.Sprintf("\"%s\"", d.Content)), nil
	}

	return json.Marshal(mDescription{
		Content: d.Content,
		Type:    d.Type,
		Version: d.Version,
	})
}

// UnmarshalJSON - анализирует данные, закодированные в формате JSON, и создает на их основе описание.
// Описание может быть создано из объекта или строки.
func (d *Description) UnmarshalJSON(b []byte) (err error) {
	if len(b) == 0 {
		return nil
	} else if len(b) >= 2 && b[0] == '"' && b[len(b)-1] == '"' {
		d.Content = string(string(b[1 : len(b)-1]))
	} else if len(b) >= 2 && b[0] == '{' && b[len(b)-1] == '}' {
		tmp := (*mDescription)(d)
		err = json.Unmarshal(b, &tmp)
	} else {
		err = errors.New("unsupported type for description")
	}

	return
}
