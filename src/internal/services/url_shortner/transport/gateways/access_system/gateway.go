package access_system_service_gateway

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	common_errors "sm-box/internal/common/errors"
	common_types "sm-box/internal/common/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	pb "sm-box/transport/proto/pb/golang/users-service"
)

const (
	loggerInitiator = "transports-[gateways]-[grpc]=access_system_service"
)

// Gateway - шлюз для работы с grpc сервером сервиса пользователей.
type Gateway struct {
	conf *Config
	ctx  context.Context

	components *components
	client     pb.AccessSystemServiceClient
}

type (
	// components - компоненты компонента.
	components struct {
		Logger logger.Logger
	}
)

// New - создание шлюза.
func New(ctx context.Context) (gw *Gateway, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelTransportGateway)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished(gw) }()
	}

	gw = &Gateway{
		ctx: ctx,
	}

	// Конфигурация
	{
		gw.conf = new(Config).Default()

		if err = gw.conf.Read(); err != nil {
			return
		}
	}

	// Компоненты
	{
		gw.components = new(components)

		// Logger
		{
			if gw.components.Logger, err = logger.New(loggerInitiator); err != nil {
				return
			}
		}
	}

	// Client
	{
		var conn *grpc.ClientConn

		if conn, err = grpc.NewClient(gw.conf.Addr, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
			return
		}

		gw.client = pb.NewAccessSystemServiceClient(conn)
	}

	gw.components.Logger.Info().
		Text("A access system service grpc gateway has been created. ").
		Field("config", gw.conf).Write()

	return
}

// CheckUserAccess - проверка доступов пользователя.
func (gw *Gateway) CheckUserAccess(ctx context.Context, userID common_types.ID, rolesID, permissionsID []common_types.ID) (allowed bool, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelTransportGatewayGrpc)

		trc.FunctionCall(ctx, userID, rolesID, permissionsID)
		defer func() { trc.Error(cErr).FunctionCallFinished(allowed) }()
	}

	var (
		response *pb.AccessSystemCheckUserAccessResponse
		request  *pb.AccessSystemCheckUserAccessRequest
	)

	// Подготовка запроса
	{
		request = &pb.AccessSystemCheckUserAccessRequest{
			UserID:        uint64(userID),
			RolesID:       make([]uint64, 0, len(rolesID)),
			PermissionsID: make([]uint64, 0, len(permissionsID)),
		}

		for _, id := range rolesID {
			request.RolesID = append(request.RolesID, uint64(id))
		}

		for _, id := range permissionsID {
			request.PermissionsID = append(request.PermissionsID, uint64(id))
		}
	}

	// Выполнение запроса
	{
		var err error

		if response, err = gw.client.CheckUserAccess(ctx, request); err != nil {
			gw.components.Logger.Error().
				Format("User access verification failed: '%s'. ", err).
				Field("user_id", userID).
				Field("roles_id", rolesID).
				Field("permissions_id", permissionsID).Write()

			cErr = c_errors.ToError(c_errors.ParseGrpc(status.Convert(err)))

			return
		}
	}

	// Проверки ответа
	{
		if response == nil {
			gw.components.Logger.Error().
				Text("User access verification failed. ").
				Field("user_id", userID).
				Field("roles_id", rolesID).
				Field("permissions_id", permissionsID).Write()

			cErr = common_errors.InternalServerError()
			cErr.SetError(errors.New("Response instance is nil. "))

			return
		}
	}

	allowed = response.Allowed

	return
}
