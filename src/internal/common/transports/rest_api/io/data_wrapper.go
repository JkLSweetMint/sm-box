package rest_api_io

import (
	"encoding/xml"
	"fmt"
)

// ResponseDataWrapper - обертка для записи данных ответа.
type ResponseDataWrapper struct {
	XMLName xml.Name `json:"-"`

	Code        int    `json:"code"         xml:"code,attr"`
	CodeMessage string `json:"code_message" xml:"code_message,attr"`
	Status      string `json:"status"       xml:"status,attr"`

	Data any `json:"data,omitempty" xml:"Data,omitempty"`
}

// Body - получение тела ответа.
func (wrapper *ResponseDataWrapper) Body() (body []byte) {
	if wrapper.Data != nil {
		body = []byte(fmt.Sprintf("%s", wrapper.Data))
	}

	return
}
