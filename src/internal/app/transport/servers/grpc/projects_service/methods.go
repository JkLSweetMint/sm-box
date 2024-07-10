package grpc_projects_srv

import (
	"context"
	"sm-box/internal/app/objects/models"
	"sm-box/internal/common/types"
	c_errors "sm-box/pkg/errors"
	pb "sm-box/transport/proto/pb/golang/app"
)

// GetListByUser - получение списка проектов пользователя.
func (srv *server) GetListByUser(ctx context.Context, request *pb.ProjectsGetListByUserRequest) (response *pb.ProjectsGetListByUserResponse, err error) {
	response = new(pb.ProjectsGetListByUserResponse)

	var (
		cErr c_errors.Error
		list models.ProjectList
	)

	if list, cErr = srv.controllers.Projects.GetListByUser(ctx, types.ID(request.ID)); cErr != nil {
		srv.components.Logger.Error().
			Format("The list of user's projects could not be retrieved: '%s'. ", cErr).Write()

		err = cErr
		return
	}

	// Преобразование данных в структуры grpc
	{
		response.List = make([]*pb.Project, 0)

		for _, project := range list {
			response.List = append(response.List, &pb.Project{
				ID:          uint64(project.ID),
				OwnerID:     0,
				Name:        project.Name,
				Description: "",
				Version:     project.Version,
			})
		}
	}

	return
}
