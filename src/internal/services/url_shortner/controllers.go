package service

import (
	"context"
	urls_controller "sm-box/internal/services/url_shortner/infrastructure/controllers/urls"
	"sm-box/internal/services/url_shortner/objects/models"
	c_errors "sm-box/pkg/errors"
)

// Controllers - описание контроллеров сервиса.
type Controllers interface {
	Urls() interface {
		RegisterToRedisDB(ctx context.Context) (cErr c_errors.Error)
		GetByReduceFromRedisDB(ctx context.Context, reduce string) (url *models.ShortUrlInfo, cErr c_errors.Error)
		UpdateInRedisDB(ctx context.Context, url *models.ShortUrlInfo) (cErr c_errors.Error)
		RemoveByReduceFromRedisDB(ctx context.Context, reduce string) (cErr c_errors.Error)
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
} {
	return controllers.urls
}
