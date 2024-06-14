package postman

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Request - представляет собой HTTP-запрос.
type Request struct {
	URL         *URL      `json:"url"`
	Auth        *Auth     `json:"auth,omitempty"`
	Proxy       any       `json:"proxy,omitempty"`
	Certificate any       `json:"certificate,omitempty"`
	Method      Method    `json:"method"`
	Description any       `json:"description,omitempty"`
	Header      []*Header `json:"header,omitempty"`
	Body        *Body     `json:"body,omitempty"`
}

// mRequest - используется для marshalling/unmarshalling.
type mRequest Request

// MarshalJSON - возвращает кодировку запроса в формате JSON.
// Если запрос содержит только URL-адрес с использованием метода Get HTTP, он возвращается в виде строки.
func (r Request) MarshalJSON() ([]byte, error) {
	if r.Auth == nil && r.Proxy == nil && r.Certificate == nil && r.Description == nil && r.Header == nil && r.Body == nil && r.Method == Get {
		return []byte(fmt.Sprintf("\"%s\"", r.URL)), nil
	}

	return json.Marshal(mRequest{
		URL:         r.URL,
		Auth:        r.Auth,
		Proxy:       r.Proxy,
		Certificate: r.Certificate,
		Method:      r.Method,
		Description: r.Description,
		Header:      r.Header,
		Body:        r.Body,
	})
}

// UnmarshalJSON - анализирует данные, закодированные в формате JSON, и создает на их основе запрос.
// Запрос может быть создан из объекта или строки.
// Если строка, то предполагается, что строка является URL-адресом запроса, а метод - "GET".
func (r *Request) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' {
		r.Method = Get
		r.URL = &URL{
			Raw: string(string(b[1 : len(b)-1])),
		}
	} else if b[0] == '{' {
		tmp := (*mRequest)(r)
		err = json.Unmarshal(b, &tmp)
	} else {
		err = errors.New("Unsupported type")
	}

	return
}
