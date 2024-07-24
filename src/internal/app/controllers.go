package app

import (
	"context"
	projects_controller "sm-box/internal/app/infrastructure/controllers/projects"
	"sm-box/internal/app/objects/models"
	common_types "sm-box/internal/common/types"
	c_errors "sm-box/pkg/errors"
)

// Controllers - описание контроллеров приложения.
type Controllers interface {
	Projects() interface {
		Get(ctx context.Context, ids ...common_types.ID) (list models.ProjectList, cErr c_errors.Error)
		GetOne(ctx context.Context, id common_types.ID) (project *models.ProjectInfo, cErr c_errors.Error)
	}
}

// controllers - контроллеры приложения.
type controllers struct {
	projects *projects_controller.Controller
}

// Projects - получение контроллера сервиса.
func (controllers *controllers) Projects() interface {
	Get(ctx context.Context, ids ...common_types.ID) (list models.ProjectList, cErr c_errors.Error)
	GetOne(ctx context.Context, id common_types.ID) (project *models.ProjectInfo, cErr c_errors.Error)
} {
	return controllers.projects
}
