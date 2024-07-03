package service

import (
	"context"
	"sm-box/internal/common/objects/models"
	"sm-box/internal/common/types"
	c_errors "sm-box/pkg/errors"
)

// Controllers - описание контроллеров сервиса.
type Controllers interface {
	Authentication() authenticationController
}

// authenticationController - описание контроллера аутентификации.
type authenticationController interface {
	BasicAuth(ctx context.Context, tokenData, username, password string) (token *models.JwtTokenInfo, user *models.UserInfo, cErr c_errors.Error)
	SetTokenProject(ctx context.Context, tokenID, projectID types.ID) (cErr c_errors.Error)
	GetUserProjectsList(ctx context.Context, tokenID, userID types.ID) (list models.ProjectList, cErr c_errors.Error)
	GetToken(ctx context.Context, data string) (token *models.JwtTokenInfo, cErr c_errors.Error)
}

// controllers - контроллеры сервиса.
type controllers struct {
	authentication authenticationController
}

// Authentication - получение контроллера аутентификации.
func (c *controllers) Authentication() authenticationController {
	return c.authentication
}
