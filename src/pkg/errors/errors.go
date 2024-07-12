package errors

import (
	"errors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	grpc_codes "google.golang.org/grpc/codes"
	grpc_status "google.golang.org/grpc/status"
	"sm-box/pkg/errors/entities/details"
	"sm-box/pkg/errors/entities/messages"
	"sm-box/pkg/errors/helpers"
	"sm-box/pkg/errors/types"
)

// Описание ошибок.
type (
	// Error - описание базовой ошибки.
	Error interface {
		ID() (id types.ID)
		Type() (t types.ErrorType)
		Status() (s types.Status)
		Message() (m string)
		Details() (details types.Details)

		SetError(err error)
		SetMessage(m types.Message)

		helpers.Error
		helpers.Stringer
	}

	// RestAPI - описание rest api ошибки.
	RestAPI interface {
		Error

		StatusCode() (c int)
	}

	// WebSocket - описание web socket ошибки.
	WebSocket interface {
		Error

		StatusCode() (c int)
	}

	// Grpc - описание grpc ошибки.
	Grpc interface {
		Error

		StatusCode() (c grpc_codes.Code)
	}
)

// ParseGrpc - парсинг grpc ошибки.
func ParseGrpc(st *grpc_status.Status) (cErr Grpc) {
	var constructor = &Constructor[Grpc]{
		ID: "",

		Type:   0,
		Status: 0,

		Err:     nil,
		Message: new(messages.TextMessage).Text(st.Message()),
		Details: new(details.Details),

		addons: &constructorAddons{
			Grpc: &GrpcConstructor{
				StatusCode: st.Code(),
			},
		},
	}

	// Парсинг.
	{
		for _, v := range st.Details() {
			switch info := v.(type) {
			case *errdetails.ErrorInfo:
				{
					if info.Reason != st.Message() {
						constructor.Err = errors.New(info.Reason)
					}

					for k, v := range info.Metadata {
						switch k {
						case "id":
							constructor.ID = types.ID(v)
						case "type":
							constructor.Type = types.ParseErrorType(v)
						case "status":
							constructor.Status = types.ParseStatus(v)
						default:
							constructor.Details.Set(k, v)
						}
					}
				}
			case *errdetails.BadRequest:
				{
					for _, field := range info.FieldViolations {
						constructor.Details.SetField(
							new(details.FieldKey).Add(field.Field),
							new(messages.TextMessage).Text(field.Description),
						)
					}
				}
			}
		}
	}

	cErr = constructor.Build()()
	return
}
