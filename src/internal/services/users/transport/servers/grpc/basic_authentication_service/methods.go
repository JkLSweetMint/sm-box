package grpc_basic_authentication_srv

import (
	"context"
	"sm-box/internal/services/users/objects/models"
	"sm-box/pkg/core/components/tracer"
	pb "sm-box/transport/proto/pb/golang/users-service"
)

// BasicAuth - базовая авторизация пользователя в системе.
// Для авторизации используется имя пользователя и пароль.
func (srv *server) Auth(ctx context.Context, request *pb.BasicAuthenticationAuthRequest) (response *pb.BasicAuthenticationAuthResponse, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelTransportGrpc)

		trc.FunctionCall(ctx, request)
		defer func() { trc.Error(err).FunctionCallFinished(response) }()
	}

	response = new(pb.BasicAuthenticationAuthResponse)

	var user *models.UserInfo

	// Получение данных
	{
		if user, err = srv.controllers.BasicAuthentication.Auth(ctx, request.Username, request.Password); err != nil {
			srv.components.Logger.Error().
				Format("User authorization failed: '%s'. ", err).Write()

			return
		}
	}

	// Преобразование данных в структуры grpc
	{
		response.User = &pb.User{
			ID: uint64(user.ID),

			Email:    user.Email,
			Username: user.Username,

			Accesses: make([]*pb.Role, 0),
		}

		var writeInheritance func(parent *pb.Role, inheritances models.RoleInfoInheritances)

		writeInheritance = func(parent *pb.Role, inheritances models.RoleInfoInheritances) {
			for _, rl := range inheritances {
				var child = &pb.Role{
					ID:        uint64(rl.ID),
					ProjectID: uint64(rl.ProjectID),

					Name:     rl.Name,
					IsSystem: rl.IsSystem,

					Inheritances: make([]*pb.Role, 0),
				}

				parent.Inheritances = append(parent.Inheritances, child)

				writeInheritance(child, rl.Inheritances)
			}
		}

		for _, rl := range user.Accesses {
			var parent = &pb.Role{
				ID:        uint64(rl.ID),
				ProjectID: uint64(rl.ProjectID),

				Name:     rl.Name,
				IsSystem: rl.IsSystem,

				Inheritances: make([]*pb.Role, 0),
			}

			response.User.Accesses = append(response.User.Accesses, parent)

			writeInheritance(parent, rl.Inheritances)
		}
	}

	return
}
