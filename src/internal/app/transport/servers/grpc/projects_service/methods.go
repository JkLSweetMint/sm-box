package grpc_projects_srv

import (
	"context"
	"sm-box/internal/app/objects/models"
	"sm-box/internal/common/types"
	pb "sm-box/transport/proto/pb/golang/app"
)

// GetOne - получение проекта по ID.
func (srv *server) GetOne(ctx context.Context, request *pb.ProjectsGetOneRequest) (response *pb.ProjectsGetOneResponse, err error) {
	response = new(pb.ProjectsGetOneResponse)

	var project *models.ProjectInfo

	if project, err = srv.controllers.Projects.GetOne(ctx, types.ID(request.ID)); err != nil {
		srv.components.Logger.Error().
			Format("Project data could not be retrieved: '%s'. ", err).Write()

		return
	}

	// Преобразование данных в структуры grpc
	{
		response.Project = &pb.Project{
			ID: uint64(project.ID),

			Name:        project.Name,
			Description: project.Description,
			Version:     project.Version,
		}
	}

	return
}

// Get - получение проектов по ID.
func (srv *server) Get(ctx context.Context, request *pb.ProjectsGetRequest) (response *pb.ProjectsGetResponse, err error) {
	response = &pb.ProjectsGetResponse{
		List: make([]*pb.Project, 0),
	}

	var (
		projects models.ProjectList
		ids      = make([]types.ID, 0, len(request.IDs))
	)

	for _, id := range request.IDs {
		ids = append(ids, types.ID(id))
	}

	if projects, err = srv.controllers.Projects.Get(ctx, ids...); err != nil {
		srv.components.Logger.Error().
			Format("Projects data could not be retrieved: '%s'. ", err).Write()

		return
	}

	// Преобразование данных в структуры grpc
	{
		for _, project := range projects {
			response.List = append(response.List, &pb.Project{
				ID: uint64(project.ID),

				Name:        project.Name,
				Description: project.Description,
				Version:     project.Version,
			})
		}
	}

	return
}
