package rest_api_io

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"github.com/gofiber/fiber/v3"
	"gopkg.in/yaml.v3"
	c_errors "sm-box/pkg/errors"
	http_tools "sm-box/pkg/tools/http"
)

// ResponseWrapper - описание обертки данных ответа.
type ResponseWrapper interface {
	Body() (body []byte)
}

// Write - запись ответа.
func Write(ctx fiber.Ctx, resp any) (err error) {
	var wrapper = &ResponseDataWrapper{
		XMLName: xml.Name{
			Space: "",
			Local: "Response",
		},

		Code:        ctx.Response().StatusCode(),
		CodeMessage: httpStatusToString(ctx.Response().StatusCode()),
		Data:        resp,
		Status:      statusFromStatusCode(ctx.Response().StatusCode()),
	}

	if err = write(ctx, wrapper); err != nil {
		return
	}

	return
}

// WriteError - запись ошибки.
func WriteError(ctx fiber.Ctx, cErr c_errors.RestAPI) (err error) {
	ctx.Status(cErr.StatusCode())

	var wrapper = &ResponseErrorWrapper{
		XMLName: xml.Name{
			Space: "",
			Local: "Response",
		},

		Code:        ctx.Response().StatusCode(),
		CodeMessage: httpStatusToString(ctx.Response().StatusCode()),
		Status:      statusFromStatusCode(ctx.Response().StatusCode()),

		Error: cErr,
	}

	if err = write(ctx, wrapper); err != nil {
		return
	}

	return
}

// write - внутренняя функция для записи данных через обертку.
func write(ctx fiber.Ctx, wrapper ResponseWrapper) (err error) {
	var data []byte

	switch {
	case acceptMimeType(ctx, []byte(http_tools.MIMEApplicationXML)):
		{
			ctx.Response().Header.SetContentType(http_tools.MIMEApplicationXML)

			if data, err = xml.Marshal(wrapper); err != nil {
				return
			}
		}
	case acceptMimeType(ctx, []byte(http_tools.MIMETextXML)):
		{
			ctx.Response().Header.SetContentType(http_tools.MIMETextXML)

			if data, err = xml.Marshal(wrapper); err != nil {
				return
			}
		}
	case acceptMimeType(ctx, []byte(http_tools.MIMEApplicationJSON)):
		{
			ctx.Response().Header.SetContentType(http_tools.MIMEApplicationJSON)

			if data, err = json.Marshal(wrapper); err != nil {
				return
			}
		}
	case acceptMimeType(ctx, []byte(http_tools.MIMETextJSON)):
		{
			ctx.Response().Header.SetContentType(http_tools.MIMETextJSON)

			if data, err = json.Marshal(wrapper); err != nil {
				return
			}
		}
	case acceptMimeType(ctx, []byte(http_tools.MIMEApplicationYAML)):
		{
			ctx.Response().Header.SetContentType(http_tools.MIMEApplicationYAML)

			if data, err = yaml.Marshal(wrapper); err != nil {
				return
			}
		}
	case acceptMimeType(ctx, []byte(http_tools.MIMETextYAML)):
		{
			ctx.Response().Header.SetContentType(http_tools.MIMETextYAML)

			if data, err = yaml.Marshal(wrapper); err != nil {
				return
			}
		}
	case acceptMimeType(ctx, []byte(http_tools.MIMEAll)):
		{
			ctx.Response().Header.SetContentType(http_tools.MIMEApplicationJSON)

			if data, err = json.Marshal(wrapper); err != nil {
				return
			}
		}
	default:
		ctx.Response().Header.SetContentType(http_tools.MIMETextPlain)

		data = wrapper.Body()
	}

	if _, err = ctx.Write(data); err != nil {
		return
	}

	return
}

// acceptMimeType - проверить является ли кодировка принимаемой.
func acceptMimeType(ctx fiber.Ctx, acceptEncoding []byte) (accepted bool) {
	for _, ae := range ctx.Request().Header.PeekAll("Accept") {
		if bytes.Index(ae, acceptEncoding) >= 0 {
			accepted = true
			break
		}
	}

	return
}

// httpStatusToString - конвертация http статус кода в строку.
func httpStatusToString(code int) (value string) {
	var (
		unknownStatusCode = "Unknown Status Code"

		statusMessages = []string{
			fiber.StatusContinue:           "Continue",
			fiber.StatusSwitchingProtocols: "Switching Protocols",
			fiber.StatusProcessing:         "Processing",
			fiber.StatusEarlyHints:         "Early Hints",

			fiber.StatusOK:                          "OK",
			fiber.StatusCreated:                     "Created",
			fiber.StatusAccepted:                    "Accepted",
			fiber.StatusNonAuthoritativeInformation: "Non-Authoritative Information",
			fiber.StatusNoContent:                   "No Content",
			fiber.StatusResetContent:                "Reset Content",
			fiber.StatusPartialContent:              "Partial Content",
			fiber.StatusMultiStatus:                 "Multi-Status",
			fiber.StatusAlreadyReported:             "Already Reported",
			fiber.StatusIMUsed:                      "IM Used",

			fiber.StatusMultipleChoices:   "Multiple Choices",
			fiber.StatusMovedPermanently:  "Moved Permanently",
			fiber.StatusFound:             "Found",
			fiber.StatusSeeOther:          "See Other",
			fiber.StatusNotModified:       "Not Modified",
			fiber.StatusUseProxy:          "Use Proxy",
			fiber.StatusTemporaryRedirect: "Temporary Redirect",
			fiber.StatusPermanentRedirect: "Permanent Redirect",

			fiber.StatusBadRequest:                   "Bad Request",
			fiber.StatusUnauthorized:                 "Unauthorized",
			fiber.StatusPaymentRequired:              "Payment Required",
			fiber.StatusForbidden:                    "Forbidden",
			fiber.StatusNotFound:                     "Not Found",
			fiber.StatusMethodNotAllowed:             "Method Not Allowed",
			fiber.StatusNotAcceptable:                "Not Acceptable",
			fiber.StatusProxyAuthRequired:            "Proxy Authentication Required",
			fiber.StatusRequestTimeout:               "Request Timeout",
			fiber.StatusConflict:                     "Conflict",
			fiber.StatusGone:                         "Gone",
			fiber.StatusLengthRequired:               "Length Required",
			fiber.StatusPreconditionFailed:           "Precondition Failed",
			fiber.StatusRequestEntityTooLarge:        "Request Entity Too Large",
			fiber.StatusRequestURITooLong:            "Request URI Too Long",
			fiber.StatusUnsupportedMediaType:         "Unsupported Media Type",
			fiber.StatusRequestedRangeNotSatisfiable: "Requested Range Not Satisfiable",
			fiber.StatusExpectationFailed:            "Expectation Failed",
			fiber.StatusTeapot:                       "I'm a teapot",
			fiber.StatusMisdirectedRequest:           "Misdirected Request",
			fiber.StatusUnprocessableEntity:          "Unprocessable Entity",
			fiber.StatusLocked:                       "Locked",
			fiber.StatusFailedDependency:             "Failed Dependency",
			fiber.StatusUpgradeRequired:              "Upgrade Required",
			fiber.StatusPreconditionRequired:         "Precondition Required",
			fiber.StatusTooManyRequests:              "Too Many Requests",
			fiber.StatusRequestHeaderFieldsTooLarge:  "Request Header Fields Too Large",
			fiber.StatusUnavailableForLegalReasons:   "Unavailable For Legal Reasons",

			fiber.StatusInternalServerError:           "Internal Server Error",
			fiber.StatusNotImplemented:                "Not Implemented",
			fiber.StatusBadGateway:                    "Bad Gateway",
			fiber.StatusServiceUnavailable:            "Service Unavailable",
			fiber.StatusGatewayTimeout:                "Gateway Timeout",
			fiber.StatusHTTPVersionNotSupported:       "HTTP Version Not Supported",
			fiber.StatusVariantAlsoNegotiates:         "Variant Also Negotiates",
			fiber.StatusInsufficientStorage:           "Insufficient Storage",
			fiber.StatusLoopDetected:                  "Loop Detected",
			fiber.StatusNotExtended:                   "Not Extended",
			fiber.StatusNetworkAuthenticationRequired: "Network Authentication Required",
		}
	)

	if code < 100 || code > 511 {
		return unknownStatusCode
	}

	if s := statusMessages[code]; s != "" {
		return s
	}

	return unknownStatusCode
}

// statusFromStatusCode - получение статуса запрсоа на основе http статус кода.
func statusFromStatusCode(code int) (value string) {
	if code >= 100 && code < 400 {
		value = "success"
	} else if code >= 400 && code < 500 {
		value = "failed"
	} else if code >= 500 && code < 600 {
		value = "error"
	}

	return
}
