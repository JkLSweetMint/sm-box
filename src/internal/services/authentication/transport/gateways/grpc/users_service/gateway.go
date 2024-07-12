package users_service_gateway

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/common/types"
	users_models "sm-box/internal/services/users/objects/models"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	pb "sm-box/transport/proto/pb/golang/users-service"
)

const (
	loggerInitiator = "transports-[gateways]-[grpc]=users_service"
)

// Gateway - шлюз для работы с grpc сервером сервиса пользователей.
type Gateway struct {
	conf *Config
	ctx  context.Context

	components *components
	client     pb.UsersClient
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

		gw.client = pb.NewUsersClient(conn)
	}

	gw.components.Logger.Info().
		Format("A '%s' grpc gateway has been created. ", "system access agent").
		Field("config", gw.conf).Write()

	return
}

// Get - получение проектов по ID.
func (gw *Gateway) Get(ctx context.Context, ids ...types.ID) (list []*users_models.UserInfo, cErr c_errors.Error) {
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
		response *pb.UsersGetResponse
		request  = &pb.UsersGetRequest{
			IDs: ids_,
		}
	)

	if response, err = gw.client.Get(ctx, request); err != nil {
		gw.components.Logger.Error().
			Format("Failed to get the users: '%s'. ", err).
			Field("ids", ids).Write()

		cErr = c_errors.ToError(c_errors.ParseGrpc(status.Convert(err)))

		return
	}

	// Проверки ответа
	{
		if response == nil || response.List == nil {
			gw.components.Logger.Error().
				Text("Failed to get users. ").Write()

			cErr = error_list.InternalServerError()
			cErr.SetError(errors.New("Users list instance is nil. "))

			return
		}
	}

	// Преобразование в модель
	{
		list = make([]*users_models.UserInfo, 0)

		for _, us := range response.List {
			var user = &users_models.UserInfo{
				ID: types.ID(us.ID),

				Email:    us.Email,
				Username: us.Username,

				Accesses: make(users_models.UserInfoAccesses, 0),
			}

			var writeInheritance func(parent *users_models.RoleInfo, inheritances []*pb.Role)

			writeInheritance = func(parent *users_models.RoleInfo, inheritances []*pb.Role) {
				for _, rl := range inheritances {
					var child = &users_models.RoleInfo{
						ID:        types.ID(rl.ID),
						ProjectID: types.ID(rl.ProjectID),

						Name:     rl.Name,
						IsSystem: rl.IsSystem,

						Inheritances: make(users_models.RoleInfoInheritances, 0),
					}

					parent.Inheritances = append(parent.Inheritances, &users_models.RoleInfoInheritance{
						RoleInfo: child,
					})

					writeInheritance(child, rl.Inheritances)
				}
			}

			for _, rl := range us.Accesses {
				var parent = &users_models.RoleInfo{
					ID:        types.ID(rl.ID),
					ProjectID: types.ID(rl.ProjectID),

					Name:     rl.Name,
					IsSystem: rl.IsSystem,

					Inheritances: make(users_models.RoleInfoInheritances, 0),
				}

				user.Accesses = append(user.Accesses, &users_models.UserInfoAccess{
					RoleInfo: parent,
				})

				writeInheritance(parent, rl.Inheritances)
			}
		}
	}

	return
}

// GetOne - получение проекта по ID.
func (gw *Gateway) GetOne(ctx context.Context, id types.ID) (user *users_models.UserInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelGateway)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(cErr).FunctionCallFinished(user) }()
	}

	var (
		err      error
		response *pb.UsersGetOneResponse
		request  = &pb.UsersGetOneRequest{
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
		if response == nil || response.User == nil {
			gw.components.Logger.Error().
				Text("Failed to get user. ").Write()

			cErr = error_list.InternalServerError()
			cErr.SetError(errors.New("User instance is nil. "))

			return
		}
	}

	// Преобразование в модель
	{
		user = &users_models.UserInfo{
			ID: types.ID(response.User.ID),

			Email:    response.User.Email,
			Username: response.User.Username,

			Accesses: make(users_models.UserInfoAccesses, 0),
		}

		var writeInheritance func(parent *users_models.RoleInfo, inheritances []*pb.Role)

		writeInheritance = func(parent *users_models.RoleInfo, inheritances []*pb.Role) {
			for _, rl := range inheritances {
				var child = &users_models.RoleInfo{
					ID:        types.ID(rl.ID),
					ProjectID: types.ID(rl.ProjectID),

					Name:     rl.Name,
					IsSystem: rl.IsSystem,

					Inheritances: make(users_models.RoleInfoInheritances, 0),
				}

				parent.Inheritances = append(parent.Inheritances, &users_models.RoleInfoInheritance{
					RoleInfo: child,
				})

				writeInheritance(child, rl.Inheritances)
			}
		}

		for _, rl := range response.User.Accesses {
			var parent = &users_models.RoleInfo{
				ID:        types.ID(rl.ID),
				ProjectID: types.ID(rl.ProjectID),

				Name:     rl.Name,
				IsSystem: rl.IsSystem,

				Inheritances: make(users_models.RoleInfoInheritances, 0),
			}

			user.Accesses = append(user.Accesses, &users_models.UserInfoAccess{
				RoleInfo: parent,
			})

			writeInheritance(parent, rl.Inheritances)
		}
	}

	return
}
