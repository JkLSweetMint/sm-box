package projects_service_gateway

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	app_models "sm-box/internal/app/objects/models"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	pb "sm-box/transport/proto/pb/golang/app"
)

const (
	loggerInitiator = "transports-[gateways]-[grpc]=projects_service"
)

// Gateway - шлюз для работы с grpc сервером сервиса пользователей.
type Gateway struct {
	conf *Config
	ctx  context.Context

	components *components
	client     pb.ProjectsClient
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
		var trc = tracer.New(tracer.LevelGateway)

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

		gw.client = pb.NewProjectsClient(conn)
	}

	gw.components.Logger.Info().
		Format("A '%s' grpc gateway has been created. ", "system access agent").
		Field("config", gw.conf).Write()

	return
}

// GetListByUser - получение списка проектов пользователя.
func (gw *Gateway) GetListByUser(ctx context.Context, userID types.ID) (list app_models.ProjectList, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelGateway)

		trc.FunctionCall(ctx, userID)
		defer func() { trc.Error(cErr).FunctionCallFinished(list) }()
	}

	var (
		err      error
		response *pb.ProjectsGetListByUserResponse
		request  = &pb.ProjectsGetListByUserRequest{
			ID: uint64(userID),
		}
	)

	if response, err = gw.client.GetListByUser(ctx, request); err != nil {
		gw.components.Logger.Error().
			Format("Authorization failed on the remote service: '%s'. ", err).Write()

		cErr = error_list.UserCouldNotBeAuthorizedOnRemoteService()
		cErr.SetError(err)

		return
	}

	// Преобразование в модель
	{
		list = make(app_models.ProjectList, 0)

		if response != nil {
			for _, project := range response.List {
				list = append(list, &app_models.ProjectInfo{
					ID:      types.ID(project.ID),
					OwnerID: types.ID(project.OwnerID),

					Name:        project.Name,
					Description: project.Description,
					Version:     project.Version,
				})
			}
		}
	}

	return
}
