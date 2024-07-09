package grpc_authentication_srv

import (
	"context"
	"sm-box/internal/services/users/objects/models"
	c_errors "sm-box/pkg/errors"
	pb "sm-box/transport/proto/pb/golang/users-service"
)

// BasicAuth - базовая авторизация пользователя в системе.
// Для авторизации используется имя пользователя и пароль.
func (srv *server) BasicAuth(ctx context.Context, request *pb.AuthenticationBasicAuthRequest) (response *pb.AuthenticationBasicAuthResponse, err error) {
	response = new(pb.AuthenticationBasicAuthResponse)

	var (
		cErr c_errors.Error
		user *models.UserInfo
	)

	if user, cErr = srv.controllers.Authentication.BasicAuth(ctx, request.Username, request.Password); cErr != nil {
		srv.components.Logger.Error().
			Format("User authorization failed: '%s'. ", cErr).Write()

		err = cErr
		return
	}

	// Преобразование данных в структуры grpc
	{
		response.User = &pb.User{
			ID:        uint64(user.ID),
			ProjectID: uint64(user.ProjectID),

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
