package grpc_users_srv

import (
	"context"
	"sm-box/internal/common/types"
	"sm-box/internal/services/users/objects/models"
	pb "sm-box/transport/proto/pb/golang/users-service"
)

// GetOne - получение пользователя по ID.
func (srv *server) GetOne(ctx context.Context, request *pb.UsersGetOneRequest) (response *pb.UsersGetOneResponse, err error) {
	response = new(pb.UsersGetOneResponse)

	var user *models.UserInfo

	if user, err = srv.controllers.Users.GetOne(ctx, types.ID(request.ID)); err != nil {
		srv.components.Logger.Error().
			Format("User data could not be retrieved: '%s'. ", err).Write()

		return
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

// Get - получение пользователей по ID.
func (srv *server) Get(ctx context.Context, request *pb.UsersGetRequest) (response *pb.UsersGetResponse, err error) {
	response = new(pb.UsersGetResponse)

	var (
		users []*models.UserInfo
		ids   = make([]types.ID, 0, len(request.IDs))
	)

	for _, id := range request.IDs {
		ids = append(ids, types.ID(id))
	}

	if users, err = srv.controllers.Users.Get(ctx, ids...); err != nil {
		srv.components.Logger.Error().
			Format("Users data could not be retrieved: '%s'. ", err).Write()

		return
	}

	// Преобразование данных в структуры grpc
	{
		for _, user := range users {
			var us = &pb.User{
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

				us.Accesses = append(us.Accesses, parent)

				writeInheritance(parent, rl.Inheritances)
			}

			response.List = append(response.List, us)
		}
	}

	return
}
