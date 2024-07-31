package users_service_gateway

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	common_errors "sm-box/internal/common/errors"
	common_types "sm-box/internal/common/types"
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

		gw.client = pb.NewUsersClient(conn)
	}

	gw.components.Logger.Info().
		Text("A users service grpc gateway has been created. ").
		Field("config", gw.conf).Write()

	return
}

// Get - получение проектов по ID.
func (gw *Gateway) Get(ctx context.Context, ids ...common_types.ID) (list []*users_models.UserInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelTransportGatewayGrpc)

		trc.FunctionCall(ctx, ids)
		defer func() { trc.Error(cErr).FunctionCallFinished(list) }()
	}

	var (
		response *pb.UsersGetResponse
		request  *pb.UsersGetRequest
	)

	// Подготовка запроса
	{
		request = new(pb.UsersGetRequest)

		request.IDs = make([]uint64, 0, len(ids))

		for _, id := range ids {
			request.IDs = append(request.IDs, uint64(id))
		}
	}

	// Выполнение запроса
	{
		var err error

		if response, err = gw.client.Get(ctx, request); err != nil {
			gw.components.Logger.Error().
				Format("Failed to get the users: '%s'. ", err).
				Field("ids", ids).Write()

			cErr = c_errors.ToError(c_errors.ParseGrpc(status.Convert(err)))

			return
		}
	}

	// Проверки ответа
	{
		if response == nil {
			gw.components.Logger.Error().
				Text("Failed to get users. ").Write()

			cErr = common_errors.InternalServerError()
			cErr.SetError(errors.New("Response instance is nil. "))

			return
		}
	}

	// Преобразование в модель
	{
		list = make([]*users_models.UserInfo, 0)

		var writeInheritance func(parent *users_models.RoleInfo, inheritances []*pb.Role)

		writeInheritance = func(parent *users_models.RoleInfo, inheritances []*pb.Role) {
			for _, rl := range inheritances {
				var child = &users_models.RoleInfo{
					ID:        common_types.ID(rl.ID),
					ProjectID: common_types.ID(rl.ProjectID),

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

		for _, us := range response.List {
			var user = &users_models.UserInfo{
				ID: common_types.ID(us.ID),

				Email:    us.Email,
				Username: us.Username,

				Accesses: &users_models.UserInfoAccesses{
					Roles:       make([]*users_models.RoleInfo, 0),
					Permissions: make([]*users_models.PermissionInfo, 0),
				},
			}

			// Права
			{
				if us.Accesses != nil && us.Accesses.Permissions != nil {
					for _, perm := range us.Accesses.Permissions {
						var permission = &users_models.PermissionInfo{
							ID:        common_types.ID(perm.ID),
							ProjectID: common_types.ID(perm.ProjectID),

							Name:            perm.Name,
							NameI18n:        uuid.UUID{},
							Description:     perm.Description,
							DescriptionI18n: uuid.UUID{},

							IsSystem: perm.IsSystem,
						}

						permission.NameI18n, _ = uuid.Parse(perm.NameI18N)
						permission.DescriptionI18n, _ = uuid.Parse(perm.DescriptionI18N)

						user.Accesses.Permissions = append(user.Accesses.Permissions, permission)
					}
				}
			}

			// Роли
			{
				if us.Accesses != nil && us.Accesses.Roles != nil {
					for _, rl := range us.Accesses.Roles {
						var parent = &users_models.RoleInfo{
							ID:        common_types.ID(rl.ID),
							ProjectID: common_types.ID(rl.ProjectID),

							Name:            rl.Name,
							NameI18n:        uuid.UUID{},
							Description:     rl.Description,
							DescriptionI18n: uuid.UUID{},

							IsSystem: rl.IsSystem,

							Permissions:  make([]*users_models.PermissionInfo, 0),
							Inheritances: make(users_models.RoleInfoInheritances, 0),
						}

						parent.NameI18n, _ = uuid.Parse(rl.NameI18N)
						parent.DescriptionI18n, _ = uuid.Parse(rl.DescriptionI18N)

						for _, permission := range rl.Permissions {
							parent.Permissions = append(parent.Permissions, &users_models.PermissionInfo{
								ID:        common_types.ID(permission.ID),
								ProjectID: common_types.ID(permission.ProjectID),

								Name:            permission.Name,
								NameI18n:        uuid.UUID{},
								Description:     permission.Description,
								DescriptionI18n: uuid.UUID{},

								IsSystem: permission.IsSystem,
							})

							parent.NameI18n, _ = uuid.Parse(permission.NameI18N)
							parent.DescriptionI18n, _ = uuid.Parse(permission.DescriptionI18N)
						}

						user.Accesses.Roles = append(user.Accesses.Roles, parent)

						writeInheritance(parent, rl.Inheritances)
					}
				}
			}
		}
	}

	return
}

// GetOne - получение проекта по ID.
func (gw *Gateway) GetOne(ctx context.Context, id common_types.ID) (user *users_models.UserInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelTransportGatewayGrpc)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(cErr).FunctionCallFinished(user) }()
	}

	var (
		response *pb.UsersGetOneResponse
		request  = &pb.UsersGetOneRequest{
			ID: uint64(id),
		}
	)

	// Выполнение запроса
	{
		var err error

		if response, err = gw.client.GetOne(ctx, request); err != nil {
			gw.components.Logger.Error().
				Format("Failed to get the project: '%s'. ", err).
				Field("id", id).Write()

			cErr = c_errors.ToError(c_errors.ParseGrpc(status.Convert(err)))

			return
		}
	}

	// Проверки ответа
	{
		if response == nil || response.User == nil {
			gw.components.Logger.Error().
				Text("Failed to get user. ").Write()

			cErr = common_errors.InternalServerError()
			cErr.SetError(errors.New("User instance is nil. "))

			return
		}
	}

	// Преобразование в модель
	{
		user = &users_models.UserInfo{
			ID: common_types.ID(response.User.ID),

			Email:    response.User.Email,
			Username: response.User.Username,

			Accesses: &users_models.UserInfoAccesses{
				Roles:       make([]*users_models.RoleInfo, 0),
				Permissions: make([]*users_models.PermissionInfo, 0),
			},
		}

		var writeInheritance func(parent *users_models.RoleInfo, inheritances []*pb.Role)

		writeInheritance = func(parent *users_models.RoleInfo, inheritances []*pb.Role) {
			for _, rl := range inheritances {
				var child = &users_models.RoleInfo{
					ID:        common_types.ID(rl.ID),
					ProjectID: common_types.ID(rl.ProjectID),

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

		// Права
		{
			if response.User.Accesses != nil && response.User.Accesses.Permissions != nil {
				for _, perm := range response.User.Accesses.Permissions {
					var permission = &users_models.PermissionInfo{
						ID:        common_types.ID(perm.ID),
						ProjectID: common_types.ID(perm.ProjectID),

						Name:            perm.Name,
						NameI18n:        uuid.UUID{},
						Description:     perm.Description,
						DescriptionI18n: uuid.UUID{},

						IsSystem: perm.IsSystem,
					}

					permission.NameI18n, _ = uuid.Parse(perm.NameI18N)
					permission.DescriptionI18n, _ = uuid.Parse(perm.DescriptionI18N)

					user.Accesses.Permissions = append(user.Accesses.Permissions, permission)
				}
			}
		}

		// Роли
		{
			if response.User.Accesses != nil && response.User.Accesses.Roles != nil {
				for _, rl := range response.User.Accesses.Roles {
					var parent = &users_models.RoleInfo{
						ID:        common_types.ID(rl.ID),
						ProjectID: common_types.ID(rl.ProjectID),

						Name:            rl.Name,
						NameI18n:        uuid.UUID{},
						Description:     rl.Description,
						DescriptionI18n: uuid.UUID{},

						IsSystem: rl.IsSystem,

						Permissions:  make([]*users_models.PermissionInfo, 0),
						Inheritances: make(users_models.RoleInfoInheritances, 0),
					}

					parent.NameI18n, _ = uuid.Parse(rl.NameI18N)
					parent.DescriptionI18n, _ = uuid.Parse(rl.DescriptionI18N)

					for _, permission := range rl.Permissions {
						parent.Permissions = append(parent.Permissions, &users_models.PermissionInfo{
							ID:        common_types.ID(permission.ID),
							ProjectID: common_types.ID(permission.ProjectID),

							Name:            permission.Name,
							NameI18n:        uuid.UUID{},
							Description:     permission.Description,
							DescriptionI18n: uuid.UUID{},

							IsSystem: permission.IsSystem,
						})

						parent.NameI18n, _ = uuid.Parse(permission.NameI18N)
						parent.DescriptionI18n, _ = uuid.Parse(permission.DescriptionI18N)
					}

					user.Accesses.Roles = append(user.Accesses.Roles, parent)

					writeInheritance(parent, rl.Inheritances)
				}
			}
		}
	}

	return
}
