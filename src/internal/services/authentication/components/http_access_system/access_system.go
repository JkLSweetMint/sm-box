package http_access_system

import (
	"context"
	"github.com/google/uuid"
	app_models "sm-box/internal/app/objects/models"
	"sm-box/internal/common/types"
	"sm-box/internal/services/authentication/objects/entities"
	users_models "sm-box/internal/services/users/objects/models"
	"sm-box/pkg/core/components/logger"
	c_errors "sm-box/pkg/errors"
)

// accessSystem - компонент системы доступа http маршрутов.
type accessSystem struct {
	conf *Config
	ctx  context.Context

	components   *components
	gateways     *gateways
	repositories *repositories
}

type (
	// components - компоненты компонента.
	components struct {
		Logger logger.Logger
	}

	// gateways - шлюзы компонента.
	gateways struct {
		Projects interface {
			Get(ctx context.Context, ids ...types.ID) (list app_models.ProjectList, cErr c_errors.Error)
			GetOne(ctx context.Context, id types.ID) (project *app_models.ProjectInfo, cErr c_errors.Error)
		}
		Users interface {
			Get(ctx context.Context, ids ...types.ID) (list []*users_models.UserInfo, cErr c_errors.Error)
			GetOne(ctx context.Context, id types.ID) (project *users_models.UserInfo, cErr c_errors.Error)
		}
	}

	// repositories - репозитории компонента.
	repositories struct {
		JwtTokens interface {
			RegisterJwtRefreshToken(ctx context.Context, tok *entities.JwtRefreshToken) (err error)
			GetJwtRefreshToken(ctx context.Context, id uuid.UUID) (tok *entities.JwtRefreshToken, err error)

			RegisterJwtAccessToken(ctx context.Context, tok *entities.JwtAccessToken) (err error)
			GetJwtAccessToken(ctx context.Context, id uuid.UUID) (tok *entities.JwtAccessToken, err error)

			Remove(ctx context.Context, id uuid.UUID) (err error)
		}
	}
)
