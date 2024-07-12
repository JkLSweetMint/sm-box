package grpc_authentication_srv

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"sm-box/internal/app/objects/models"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	pb "sm-box/transport/proto/pb/golang/app"
	"sync"
	"time"
)

// server - grpc сервер для сервиса аутентификации пользователей.
type server struct {
	pb.AuthenticationServer
	listener net.Listener
	grpc     *grpc.Server

	conf *Config
	ctx  context.Context

	controllers *controllers
	components  *components
}

// controllers - контроллеры сервера.
type controllers struct {
	Authentication interface {
		BasicAuth(ctx context.Context, username, password string) (user *models.UserInfo, cErr c_errors.Grpc)
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
		Text("The grpc server for authentication service is listening... ").
		Field("config", srv.conf).Write()

	if srv.listener, err = net.Listen("tcp", srv.conf.Addr); err != nil {
		srv.components.Logger.Error().
			Format("An error occurred when starting the grpc server for authentication service maintenance: '%s'. ", err).Write()
		return
	}

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		if err = srv.grpc.Serve(srv.listener); err != nil {
			srv.components.Logger.Error().
				Format("An error occurred when starting the grpc server for authentication service maintenance: '%s'. ", err).Write()
			return
		}
	}()

	time.Sleep(time.Second)

	if err != nil {
		return
	}

	srv.components.Logger.Info().
		Text("The grpc server for authentication service is listened. ").Write()

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
		Text("Shutting down the grpc server for authentication service... ").Write()

	srv.grpc.Stop()

	srv.components.Logger.Info().
		Text("The grpc server for authentication service is turned off. ").Write()

	return
}
