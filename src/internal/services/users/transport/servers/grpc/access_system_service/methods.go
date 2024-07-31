package grpc_access_system_srv

import (
	"context"
	common_types "sm-box/internal/common/types"
	"sm-box/pkg/core/components/tracer"
	pb "sm-box/transport/proto/pb/golang/users-service"
)

// CheckUserAccess - проверка доступов пользователя.
func (srv *server) CheckUserAccess(ctx context.Context, request *pb.AccessSystemCheckUserAccessRequest) (response *pb.AccessSystemCheckUserAccessResponse, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelTransportGrpc)

		trc.FunctionCall(ctx, request)
		defer func() { trc.Error(err).FunctionCallFinished(response) }()
	}

	response = new(pb.AccessSystemCheckUserAccessResponse)

	var (
		rolesID       = make([]common_types.ID, 0)
		permissionsID = make([]common_types.ID, 0)
		userID        = common_types.ID(request.UserID)
	)

	// Подготовка входных аргументов
	{
		if request.RolesID != nil {
			for _, id := range request.RolesID {
				rolesID = append(rolesID, common_types.ID(id))
			}
		}

		if request.PermissionsID != nil {
			for _, id := range request.PermissionsID {
				permissionsID = append(permissionsID, common_types.ID(id))
			}
		}
	}

	// Проверка доступов
	{
		if response.Allowed, err = srv.controllers.AccessSystem.CheckUserAccess(ctx, userID, rolesID, permissionsID); err != nil {
			srv.components.Logger.Error().
				Format("User access verification failed: '%s'. ", err).
				Field("user_id", userID).
				Field("roles_id", rolesID).
				Field("permissions_id", permissionsID).Write()

			return
		}
	}

	return
}
