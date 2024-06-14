package postman

import (
	"encoding/json"
	"errors"
	"fmt"
)

// URL - представляет собой структуру, которая содержит URL-адрес в "разобранном виде".
// Raw содержит полный URL-адрес.
type URL struct {
	version   version
	Raw       string        `json:"raw"`
	Protocol  string        `json:"protocol,omitempty"`
	Host      []string      `json:"host,omitempty"`
	Path      []string      `json:"path,omitempty"`
	Port      string        `json:"port,omitempty"`
	Query     []*QueryParam `json:"query,omitempty"`
	Hash      string        `json:"hash,omitempty"`
	Variables []*Variable   `json:"variable,omitempty" mapstructure:"variable"`
}

// mURL - используется для сортировки/разгрузки данных.
type mURL URL

type QueryParam struct {
	Key         string  `json:"key"`
	Value       string  `json:"value"`
	Description *string `json:"description"`
}

// String - возвращает исходную версию URL-адреса.
func (u URL) String() string {
	return u.Raw
}

func (u *URL) setVersion(v version) {
	u.version = v
}

// MarshalJSON - возвращает кодировку URL-адреса в формате JSON.
// URL-адрес кодируется как строка, если он не содержит какой-либо переменной.
// Если он содержит какую-либо переменную, он кодируется как объект.
func (u URL) MarshalJSON() ([]byte, error) {
	// Коллекция Postman всегда является объектами в версии 2.1.0, но может быть строками в версии 2.0.0.
	if u.version == V200 && u.Variables == nil {
		return []byte(fmt.Sprintf("\"%s\"", u.Raw)), nil
	}

	return json.Marshal(mURL{
		Raw:       u.Raw,
		Protocol:  u.Protocol,
		Host:      u.Host,
		Path:      u.Path,
		Port:      u.Port,
		Query:     u.Query,
		Hash:      u.Hash,
		Variables: u.Variables,
	})
}

// UnmarshalJSON - анализирует данные, закодированные в формате JSON, и создает из них URL-адрес.
// URL-адрес может быть создан из объекта или строки.
// Если это строка, то предполагается, что значение является исходным атрибутом URL-адреса.
func (u *URL) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' {
		u.Raw = string(b[1 : len(b)-1])
	} else if b[0] == '{' {
		tmp := (*mURL)(u)
		err = json.Unmarshal(b, &tmp)
	} else {
		err = errors.New("Unsupported type")
	}

	return
}
