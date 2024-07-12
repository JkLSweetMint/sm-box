package grpc

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	grpc_codes "google.golang.org/grpc/codes"
	grpc_status "google.golang.org/grpc/status"
	"google.golang.org/protobuf/protoadapt"
	"sm-box/pkg/errors/internal"
)

type (
	// Internal - внутренняя реализация ошибки grpc.
	Internal struct {
		*internal.Internal
	}
)

// New - создание внутренней реализации ошибки grpc.
func New(store *internal.Store) (i *Internal) {
	i = &Internal{
		Internal: internal.New(store),
	}

	return
}

// StatusCode - получение статус кода grpc ошибки.
func (i *Internal) StatusCode() (c grpc_codes.Code) {
	c = i.Internal.Store.Others.Grpc.StatusCode
	return
}

// GRPCStatus - упаковка ошибки для передачи по grpc.
func (i *Internal) GRPCStatus() *grpc_status.Status {
	var status = grpc_status.New(i.Store.Others.Grpc.StatusCode, i.Message())

	var details = make([]protoadapt.MessageV1, 0)

	// Информация об ошибке.
	{
		var info = &errdetails.ErrorInfo{
			Reason: i.Error(),
			Metadata: map[string]string{
				"id":     string(i.ID()),
				"type":   i.Type().String(),
				"status": i.Status().String(),
			},
		}

		if list := i.Details().PeekAll(); len(list) > 0 {
			for k, v := range list {
				if k != "id" && k != "type" && k != "status" {
					info.Metadata[k] = v
				}
			}
		}

		details = append(details, info)
	}

	// Поля из деталей
	{
		if list := i.Details().PeekFields(); len(list) > 0 {
			var info = &errdetails.BadRequest{
				FieldViolations: make([]*errdetails.BadRequest_FieldViolation, 0),
			}

			for _, field := range list {
				info.FieldViolations = append(info.FieldViolations, &errdetails.BadRequest_FieldViolation{
					Field:       field.Key.String(),
					Description: field.Message.String(),
				})
			}

			details = append(details, info)
		}
	}

	status, _ = status.WithDetails(details...)

	return status
}
