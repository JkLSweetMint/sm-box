package grpc_access_system_srv

import (
	"context"
	"google.golang.org/grpc"
	access_system_adapter "sm-box/internal/services/users/infrastructure/adapters/access_system"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	pb "sm-box/transport/proto/pb/golang/users-service"
)

const (
	loggerInitiator = "transport-[servers]-[grpc]=access_system-service"
)

// Server - описание grpc сервера для сервиса системы доступа пользователей.
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

		// AccessSystem
		{
			if ref.controllers.AccessSystem, err = access_system_adapter.New_Grpc(ref.ctx); err != nil {
				return
			}
		}
	}

	// grpc
	{
		var opts []grpc.ServerOption

		ref.grpc = grpc.NewServer(opts...)

		pb.RegisterAccessSystemServiceServer(ref.grpc, ref)
	}

	ref.components.Logger.Info().
		Text("The grpc server access system service has been created. ").
		Field("config", ref.conf).Write()

	srv = ref

	return
}