package service

import (
	"context"
	common_types "sm-box/internal/common/types"
	basic_authentication_controller "sm-box/internal/services/users/infrastructure/controllers/basic_authentication"
	users_controller "sm-box/internal/services/users/infrastructure/controllers/users"
	"sm-box/internal/services/users/objects/models"
	c_errors "sm-box/pkg/errors"
)

// Controllers - описание контроллеров сервиса.
type Controllers interface {
	BasicAuthentication() interface {
		Auth(ctx context.Context, username, password string) (user *models.UserInfo, cErr c_errors.Error)
	}

	Users() interface {
		Get(ctx context.Context, ids ...common_types.ID) (list []*models.UserInfo, cErr c_errors.Error)
		GetOne(ctx context.Context, id common_types.ID) (user *models.UserInfo, cErr c_errors.Error)
	}
}

// controllers - контроллеры сервиса.
type controllers struct {
	basicAuthentication *basic_authentication_controller.Controller
	users               *users_controller.Controller
}

// BasicAuthentication - получение контроллера сервиса.
func (controllers *controllers) BasicAuthentication() interface {
	Auth(ctx context.Context, username, password string) (user *models.UserInfo, cErr c_errors.Error)
} {
	return controllers.basicAuthentication
}

// Users - получение контроллера сервиса.
func (controllers *controllers) Users() interface {
	Get(ctx context.Context, ids ...common_types.ID) (list []*models.UserInfo, cErr c_errors.Error)
	GetOne(ctx context.Context, id common_types.ID) (user *models.UserInfo, cErr c_errors.Error)
} {
	return controllers.users
}
