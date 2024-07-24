package grpc_users_srv

import (
	"context"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/users/objects/models"
	"sm-box/pkg/core/components/tracer"
	pb "sm-box/transport/proto/pb/golang/users-service"
)

// GetOne - получение пользователя по ID.
func (srv *server) GetOne(ctx context.Context, request *pb.UsersGetOneRequest) (response *pb.UsersGetOneResponse, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelTransportGrpc)

		trc.FunctionCall(ctx, request)
		defer func() { trc.Error(err).FunctionCallFinished(response) }()
	}

	response = new(pb.UsersGetOneResponse)

	var user *models.UserInfo

	// Получение данных
	{
		if user, err = srv.controllers.Users.GetOne(ctx, common_types.ID(request.ID)); err != nil {
			srv.components.Logger.Error().
				Format("User data could not be retrieved: '%s'. ", err).Write()

			return
		}
	}

	// Преобразование данных в структуры grpc
	{
		response.User = &pb.User{
			ID: uint64(user.ID),

			Email:    user.Email,
			Username: user.Username,

			Accesses: &pb.UserAccesses{
				Permissions: make([]*pb.Permission, 0),
				Roles:       make([]*pb.Role, 0),
			},
		}

		// Доступы
		{
			if user.Accesses != nil {
				// Права
				{
					for _, permission := range user.Accesses.Permissions {
						response.User.Accesses.Permissions = append(response.User.Accesses.Permissions, &pb.Permission{
							ID:        uint64(permission.ID),
							ProjectID: uint64(permission.ProjectID),

							Name:            permission.Name,
							NameI18N:        permission.NameI18n.String(),
							Description:     permission.Description,
							DescriptionI18N: permission.DescriptionI18n.String(),

							IsSystem: permission.IsSystem,
						})
					}
				}

				// Роли
				{
					var writeInheritance func(parent *pb.Role, inheritances models.RoleInfoInheritances)

					writeInheritance = func(parent *pb.Role, inheritances models.RoleInfoInheritances) {
						for _, rl := range inheritances {
							var child = &pb.Role{
								ID:        uint64(rl.ID),
								ProjectID: uint64(rl.ProjectID),

								Name:            rl.Name,
								NameI18N:        rl.NameI18n.String(),
								Description:     rl.Description,
								DescriptionI18N: rl.DescriptionI18n.String(),

								IsSystem: rl.IsSystem,

								Permissions:  make([]*pb.Permission, 0),
								Inheritances: make([]*pb.Role, 0),
							}

							for _, permission := range rl.Permissions {
								child.Permissions = append(child.Permissions, &pb.Permission{
									ID:        uint64(permission.ID),
									ProjectID: uint64(permission.ProjectID),

									Name:            permission.Name,
									NameI18N:        permission.NameI18n.String(),
									Description:     permission.Description,
									DescriptionI18N: permission.DescriptionI18n.String(),

									IsSystem: permission.IsSystem,
								})
							}

							parent.Inheritances = append(parent.Inheritances, child)

							writeInheritance(child, rl.Inheritances)
						}
					}

					for _, rl := range user.Accesses.Roles {
						var parent = &pb.Role{
							ID:        uint64(rl.ID),
							ProjectID: uint64(rl.ProjectID),

							Name:            rl.Name,
							NameI18N:        rl.NameI18n.String(),
							Description:     rl.Description,
							DescriptionI18N: rl.DescriptionI18n.String(),

							IsSystem: rl.IsSystem,

							Permissions:  make([]*pb.Permission, 0),
							Inheritances: make([]*pb.Role, 0),
						}

						// Права
						{
							for _, permission := range rl.Permissions {
								parent.Permissions = append(parent.Permissions, &pb.Permission{
									ID:        uint64(permission.ID),
									ProjectID: uint64(permission.ProjectID),

									Name:            permission.Name,
									NameI18N:        permission.NameI18n.String(),
									Description:     permission.Description,
									DescriptionI18N: permission.DescriptionI18n.String(),

									IsSystem: permission.IsSystem,
								})
							}
						}

						response.User.Accesses.Roles = append(response.User.Accesses.Roles, parent)

						writeInheritance(parent, rl.Inheritances)
					}
				}
			}
		}
	}

	return
}

// Get - получение пользователей по ID.
func (srv *server) Get(ctx context.Context, request *pb.UsersGetRequest) (response *pb.UsersGetResponse, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelTransportGrpc)

		trc.FunctionCall(ctx, request)
		defer func() { trc.Error(err).FunctionCallFinished(response) }()
	}

	response = &pb.UsersGetResponse{
		List: make([]*pb.User, 0),
	}

	var users []*models.UserInfo

	// Получение данных
	{
		var ids = make([]common_types.ID, 0, len(request.IDs))

		// Сбор id
		{
			for _, id := range request.IDs {
				ids = append(ids, common_types.ID(id))
			}
		}

		if users, err = srv.controllers.Users.Get(ctx, ids...); err != nil {
			srv.components.Logger.Error().
				Format("Users data could not be retrieved: '%s'. ", err).Write()

			return
		}
	}

	// Преобразование данных в структуры grpc
	{
		for _, user := range users {
			var us = &pb.User{
				ID: uint64(user.ID),

				Email:    user.Email,
				Username: user.Username,

				Accesses: &pb.UserAccesses{
					Permissions: make([]*pb.Permission, 0),
					Roles:       make([]*pb.Role, 0),
				},
			}

			// Доступы
			{
				if user.Accesses != nil {
					// Права
					{
						for _, permission := range user.Accesses.Permissions {
							us.Accesses.Permissions = append(us.Accesses.Permissions, &pb.Permission{
								ID:        uint64(permission.ID),
								ProjectID: uint64(permission.ProjectID),

								Name:            permission.Name,
								NameI18N:        permission.NameI18n.String(),
								Description:     permission.Description,
								DescriptionI18N: permission.DescriptionI18n.String(),

								IsSystem: permission.IsSystem,
							})
						}
					}

					// Роли
					{
						var writeInheritance func(parent *pb.Role, inheritances models.RoleInfoInheritances)

						writeInheritance = func(parent *pb.Role, inheritances models.RoleInfoInheritances) {
							for _, rl := range inheritances {
								var child = &pb.Role{
									ID:        uint64(rl.ID),
									ProjectID: uint64(rl.ProjectID),

									Name:            rl.Name,
									NameI18N:        rl.NameI18n.String(),
									Description:     rl.Description,
									DescriptionI18N: rl.DescriptionI18n.String(),

									IsSystem: rl.IsSystem,

									Permissions:  make([]*pb.Permission, 0),
									Inheritances: make([]*pb.Role, 0),
								}

								for _, permission := range rl.Permissions {
									child.Permissions = append(child.Permissions, &pb.Permission{
										ID:        uint64(permission.ID),
										ProjectID: uint64(permission.ProjectID),

										Name:            permission.Name,
										NameI18N:        permission.NameI18n.String(),
										Description:     permission.Description,
										DescriptionI18N: permission.DescriptionI18n.String(),

										IsSystem: permission.IsSystem,
									})
								}

								parent.Inheritances = append(parent.Inheritances, child)

								writeInheritance(child, rl.Inheritances)
							}
						}

						for _, rl := range user.Accesses.Roles {
							var parent = &pb.Role{
								ID:        uint64(rl.ID),
								ProjectID: uint64(rl.ProjectID),

								Name:            rl.Name,
								NameI18N:        rl.NameI18n.String(),
								Description:     rl.Description,
								DescriptionI18N: rl.DescriptionI18n.String(),

								IsSystem: rl.IsSystem,

								Permissions:  make([]*pb.Permission, 0),
								Inheritances: make([]*pb.Role, 0),
							}

							// Права
							{
								for _, permission := range rl.Permissions {
									parent.Permissions = append(parent.Permissions, &pb.Permission{
										ID:        uint64(permission.ID),
										ProjectID: uint64(permission.ProjectID),

										Name:            permission.Name,
										NameI18N:        permission.NameI18n.String(),
										Description:     permission.Description,
										DescriptionI18N: permission.DescriptionI18n.String(),

										IsSystem: permission.IsSystem,
									})
								}
							}

							us.Accesses.Roles = append(us.Accesses.Roles, parent)

							writeInheritance(parent, rl.Inheritances)
						}
					}
				}
			}

			response.List = append(response.List, us)
		}
	}

	return
}
