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
