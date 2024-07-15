package service

import (
	"context"
	app_models "sm-box/internal/app/objects/models"
	"sm-box/internal/common/types"
	basic_authentication_controller "sm-box/internal/services/authentication/infrastructure/controllers/basic_authentication"
	"sm-box/internal/services/authentication/objects/models"
	c_errors "sm-box/pkg/errors"
)

// Controllers - описание контроллеров сервиса.
type Controllers interface {
	BasicAuthentication() interface {
		Auth(ctx context.Context, rawSessionToken, username, password string) (sessionToken *models.JwtTokenInfo, cErr c_errors.Error)
		GetUserProjectList(ctx context.Context, rawSessionToken string) (list app_models.ProjectList, cErr c_errors.Error)
		SetTokenProject(ctx context.Context, rawSessionToken string, projectID types.ID) (
			sessionToken, accessToken, refreshToken *models.JwtTokenInfo, cErr c_errors.Error)
		Logout(ctx context.Context, rawToken string) (cErr c_errors.Error)
	}
}

// controllers - контроллеры сервиса.
type controllers struct {
	basicAuthentication *basic_authentication_controller.Controller
}

// BasicAuthentication - получение контроллера сервиса.
func (controllers *controllers) BasicAuthentication() interface {
	Auth(ctx context.Context, rawSessionToken, username, password string) (sessionToken *models.JwtTokenInfo, cErr c_errors.Error)
	GetUserProjectList(ctx context.Context, rawSessionToken string) (list app_models.ProjectList, cErr c_errors.Error)
	SetTokenProject(ctx context.Context, rawSessionToken string, projectID types.ID) (
		sessionToken, accessToken, refreshToken *models.JwtTokenInfo, cErr c_errors.Error)
	Logout(ctx context.Context, rawToken string) (cErr c_errors.Error)
} {
	return controllers.basicAuthentication
}
