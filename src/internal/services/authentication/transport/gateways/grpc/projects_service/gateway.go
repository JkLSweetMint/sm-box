package projects_service_gateway

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
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

// Get - получение проектов по ID.
func (gw *Gateway) Get(ctx context.Context, ids ...types.ID) (list app_models.ProjectList, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelGateway)

		trc.FunctionCall(ctx, ids)
		defer func() { trc.Error(cErr).FunctionCallFinished(list) }()
	}

	var ids_ = make([]uint64, 0, len(ids))

	for _, id := range ids {
		ids_ = append(ids_, uint64(id))
	}

	var (
		err      error
		response *pb.ProjectsGetResponse
		request  = &pb.ProjectsGetRequest{
			IDs: ids_,
		}
	)

	if response, err = gw.client.Get(ctx, request); err != nil {
		gw.components.Logger.Error().
			Format("Failed to get the projects: '%s'. ", err).
			Field("ids", ids).Write()

		cErr = c_errors.ToError(c_errors.ParseGrpc(status.Convert(err)))

		return
	}

	// Проверки ответа
	{
		if response == nil {
			gw.components.Logger.Error().
				Text("Failed to get users. ").Write()

			cErr = error_list.InternalServerError()
			cErr.SetError(errors.New("Response instance is nil. "))

			return
		}
	}

	// Преобразование в модель
	{
		list = make(app_models.ProjectList, 0)

		for _, project := range response.List {
			list = append(list, &app_models.ProjectInfo{
				ID: types.ID(project.ID),

				Name:        project.Name,
				Description: project.Description,
				Version:     project.Version,
			})
		}
	}

	return
}

// GetOne - получение проекта по ID.
func (gw *Gateway) GetOne(ctx context.Context, id types.ID) (project *app_models.ProjectInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelGateway)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(cErr).FunctionCallFinished(project) }()
	}

	var (
		err      error
		response *pb.ProjectsGetOneResponse
		request  = &pb.ProjectsGetOneRequest{
			ID: uint64(id),
		}
	)

	if response, err = gw.client.GetOne(ctx, request); err != nil {
		gw.components.Logger.Error().
			Format("Failed to get the project: '%s'. ", err).
			Field("id", id).Write()

		cErr = c_errors.ToError(c_errors.ParseGrpc(status.Convert(err)))

		return
	}

	// Проверки ответа
	{
		if response == nil || response.Project == nil {
			gw.components.Logger.Error().
				Text("Failed to get project. ").Write()

			cErr = error_list.InternalServerError()
			cErr.SetError(errors.New("Project instance is nil. "))

			return
		}
	}

	// Преобразование в модель
	{
		project = &app_models.ProjectInfo{
			ID: types.ID(response.Project.ID),

			Name:        response.Project.Name,
			Description: response.Project.Description,
			Version:     response.Project.Version,
		}
	}

	return
}