package grpc_projects_srv

import (
	"context"
	"sm-box/internal/app/objects/models"
	"sm-box/internal/common/types"
	pb "sm-box/transport/proto/pb/golang/app"
)

// GetListByUser - получение списка проектов пользователя.
func (srv *server) GetListByUser(ctx context.Context, request *pb.ProjectsGetListByUserRequest) (response *pb.ProjectsGetListByUserResponse, err error) {
	response = new(pb.ProjectsGetListByUserResponse)

	var list models.ProjectList

	if list, err = srv.controllers.Projects.GetListByUser(ctx, types.ID(request.UserID)); err != nil {
		srv.components.Logger.Error().
			Format("The list of user's projects could not be retrieved: '%s'. ", err).Write()

		return
	}

	// Преобразование данных в структуры grpc
	{
		response.List = make([]*pb.Project, 0)

		for _, project := range list {
			response.List = append(response.List, &pb.Project{
				ID:          uint64(project.ID),
				OwnerID:     uint64(project.OwnerID),
				Name:        project.Name,
				Description: project.Description,
				Version:     project.Version,
			})
		}
	}

	return
}

// Get - получение проекта.
func (srv *server) Get(ctx context.Context, request *pb.ProjectsGetRequest) (response *pb.ProjectsGetResponse, err error) {
	response = new(pb.ProjectsGetResponse)

	var project *models.ProjectInfo

	if project, err = srv.controllers.Projects.Get(ctx, types.ID(request.ID)); err != nil {
		srv.components.Logger.Error().
			Format("Failed to get the project: '%s'. ", err).Write()

		return
	}

	// Преобразование данных в структуры grpc
	{
		response.Project = &pb.Project{
			ID:      uint64(project.ID),
			OwnerID: uint64(project.OwnerID),

			Name:        project.Name,
			Description: project.Description,
			Version:     project.Version,
		}
	}

	return
}
