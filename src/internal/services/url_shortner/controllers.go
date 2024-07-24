package service

import (
	"context"
	common_types "sm-box/internal/common/types"
	authentication_entities "sm-box/internal/services/authentication/objects/entities"
	urls_controller "sm-box/internal/services/url_shortner/infrastructure/controllers/urls"
	"sm-box/internal/services/url_shortner/objects/models"
	"sm-box/internal/services/url_shortner/objects/types"
	c_errors "sm-box/pkg/errors"
)

// Controllers - описание контроллеров сервиса.
type Controllers interface {
	Urls() interface {
		RegisterToRedisDB(ctx context.Context) (cErr c_errors.Error)
		GetByReduceFromRedisDB(ctx context.Context, reduce string) (url *models.ShortUrlInfo, cErr c_errors.Error)
		UpdateInRedisDB(ctx context.Context, url *models.ShortUrlInfo) (cErr c_errors.Error)
		RemoveByReduceFromRedisDB(ctx context.Context, reduce string) (cErr c_errors.Error)

		WriteCallToHistory(ctx context.Context, id common_types.ID, status types.ShortUrlUsageHistoryStatus, token *authentication_entities.JwtSessionToken) (cErr c_errors.Error)
	}
}

// controllers - контроллеры сервиса.
type controllers struct {
	urls *urls_controller.Controller
}

// Urls - получение контроллера сервиса.
func (controllers *controllers) Urls() interface {
	RegisterToRedisDB(ctx context.Context) (cErr c_errors.Error)
	GetByReduceFromRedisDB(ctx context.Context, reduce string) (url *models.ShortUrlInfo, cErr c_errors.Error)
	UpdateInRedisDB(ctx context.Context, url *models.ShortUrlInfo) (cErr c_errors.Error)
	RemoveByReduceFromRedisDB(ctx context.Context, reduce string) (cErr c_errors.Error)

	WriteCallToHistory(ctx context.Context, id common_types.ID, status types.ShortUrlUsageHistoryStatus, token *authentication_entities.JwtSessionToken) (cErr c_errors.Error)
} {
	return controllers.urls
}
