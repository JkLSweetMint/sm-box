package grpc_authentication_srv

import (
	"context"
	"google.golang.org/grpc"
	authentication_adapter "sm-box/internal/app/infrastructure/adapters/authentication"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	pb "sm-box/transport/proto/pb/golang/app"
)

const (
	loggerInitiator = "transport-[servers]-[grpc]=authentication-service"
)

// Server - описание grpc сервера для сервиса аутентификации пользователей.
type Server interface {
	Listen() (err error)
	Shutdown() (err error)
}

// New - создание сервера.
func New(ctx context.Context) (srv Server, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelTransport)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished(srv) }()
	}

	var ref = &server{
		ctx: ctx,
	}

	// Конфигурация
	{
		ref.conf = new(Config).Default()

		if err = ref.conf.Read(); err != nil {
			return
		}
	}

	// Компоненты
	{
		ref.components = new(components)

		// Logger
		{
			if ref.components.Logger, err = logger.New(loggerInitiator); err != nil {
				return
			}
		}
	}

	// Контроллеры
	{
		ref.controllers = new(controllers)

		// Authentication
		{
			if ref.controllers.Authentication, err = authentication_adapter.New_Grpc(ref.ctx); err != nil {
				return
			}
		}
	}

	// grpc
	{
		var opts []grpc.ServerOption

		ref.grpc = grpc.NewServer(opts...)

		pb.RegisterAuthenticationServer(ref.grpc, ref)
	}

	ref.components.Logger.Info().
		Text("The grpc server authentication service has been created. ").
		Field("config", ref.conf).Write()

	srv = ref

	return
}
