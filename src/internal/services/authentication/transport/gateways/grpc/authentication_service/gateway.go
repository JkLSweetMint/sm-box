package authentication_service_gateway

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
	loggerInitiator = "transports-[gateways]-[grpc]=authentication_service"
)

// Gateway - шлюз для работы с grpc сервером сервиса аутентификации пользователей.
type Gateway struct {
	conf *Config
	ctx  context.Context

	components *components
	client     pb.AuthenticationClient
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

		gw.client = pb.NewAuthenticationClient(conn)
	}

	gw.components.Logger.Info().
		Format("A '%s' grpc gateway has been created. ", "system access agent").
		Field("config", gw.conf).Write()

	return
}

// BasicAuth - базовая авторизация пользователя в системе.
// Для авторизации используется имя пользователя и пароль.
func (gw *Gateway) BasicAuth(ctx context.Context, username, password string) (user *users_models.UserInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelGateway)

		trc.FunctionCall(ctx, username, password)
		defer func() { trc.Error(cErr).FunctionCallFinished(user) }()
	}

	var (
		err      error
		response *pb.AuthenticationBasicAuthResponse
		request  = &pb.AuthenticationBasicAuthRequest{
			Username: username,
			Password: password,
		}
	)

	if response, err = gw.client.BasicAuth(ctx, request); err != nil {
		gw.components.Logger.Error().
			Format("Authorization failed on the remote service: '%s'. ", err).Write()

		cErr = c_errors.ToError(c_errors.ParseGrpc(status.Convert(err)))

		return
	}

	// Проверки ответа
	{
		if response == nil || response.User == nil {
			gw.components.Logger.Error().
				Text("Authorization failed on the remote service. ").Write()

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
