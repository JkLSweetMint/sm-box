package grpc_access_system_srv

import (
	"context"
	"google.golang.org/grpc"
	"net"
	common_types "sm-box/internal/common/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	pb "sm-box/transport/proto/pb/golang/users-service"
	"sync"
	"time"
)

// server - grpc сервер для сервиса системы доступа пользователей.
type server struct {
	pb.AccessSystemServer
	listener net.Listener
	grpc     *grpc.Server

	conf *Config
	ctx  context.Context

	controllers *controllers
	components  *components
}

// controllers - контроллеры сервера.
type controllers struct {
	AccessSystem interface {
		CheckUserAccess(ctx context.Context, userID common_types.ID, rolesID, permissionsID []common_types.ID) (allowed bool, cErr c_errors.Grpc)
	}
}

// components - компоненты сервера.
type components struct {
	Logger logger.Logger
}

// Listen - запуск сервера.
func (srv *server) Listen() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelTransport)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	srv.components.Logger.Info().
		Text("The grpc server for access system service is listening... ").
		Field("config", srv.conf).Write()

	if srv.listener, err = net.Listen("tcp", srv.conf.Addr); err != nil {
		srv.components.Logger.Error().
			Format("An error occurred when starting the grpc server for access system service maintenance: '%s'. ", err).Write()
		return
	}

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		if err = srv.grpc.Serve(srv.listener); err != nil {
			srv.components.Logger.Error().
				Format("An error occurred when starting the grpc server for access system service maintenance: '%s'. ", err).Write()
			return
		}
	}()

	time.Sleep(time.Second)

	if err != nil {
		return
	}

	srv.components.Logger.Info().
		Text("The grpc server for access system service is listened. ").Write()

	wg.Wait()

	return
}

// Shutdown - завершение работы сервера.
func (srv *server) Shutdown() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelTransport)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	srv.components.Logger.Info().
		Text("Shutting down the grpc server for access system service... ").Write()

	srv.grpc.Stop()

	srv.components.Logger.Info().
		Text("The grpc server for access system service is turned off. ").Write()

	return
}
