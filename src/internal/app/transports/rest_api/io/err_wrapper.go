package rest_api_io

import (
	"encoding/xml"
	"fmt"
	c_errors "sm-box/pkg/errors"
)

// ResponseErrorWrapper - обертка для записи ошибки.
type ResponseErrorWrapper struct {
	XMLName xml.Name `json:"-"`

	Code        int    `json:"code"         xml:"code,attr"`
	CodeMessage string `json:"code_message" xml:"code_message,attr"`
	Status      string `json:"status"       xml:"status,attr"`

	Error c_errors.RestAPI `json:"error" xml:"Error"`
}

// Body - получение тела ответа.
func (wrapper *ResponseErrorWrapper) Body() (body []byte) {
	if wrapper.Error != nil {
		body = []byte(fmt.Sprintf("%s", wrapper.Error))
	}

	return
}
