package service

import (
	"context"
	languages_controller "sm-box/internal/services/i18n/infrastructure/controllers/languages"
	texts_controller "sm-box/internal/services/i18n/infrastructure/controllers/texts"
	"sm-box/internal/services/i18n/objects/models"
	c_errors "sm-box/pkg/errors"
)

// Controllers - описание контроллеров сервиса.
type Controllers interface {
	Texts() interface {
		AssembleDictionary(ctx context.Context, lang string, paths []string) (dictionary models.Dictionary, cErr c_errors.Error)
	}
	Languages() interface {
		GetList(ctx context.Context) (list []*models.Language, cErr c_errors.Error)
		Remove(ctx context.Context, code string) (cErr c_errors.Error)
		Update(ctx context.Context, code, name string) (cErr c_errors.Error)
		Create(ctx context.Context, code string, name string) (cErr c_errors.Error)

		Activate(ctx context.Context, code string) (cErr c_errors.Error)
		Deactivate(ctx context.Context, code string) (cErr c_errors.Error)
	}
}

// controllers - контроллеры сервиса.
type controllers struct {
	texts     *texts_controller.Controller
	languages *languages_controller.Controller
}

// Texts - получение контроллера сервиса.
func (controllers *controllers) Texts() interface {
	AssembleDictionary(ctx context.Context, lang string, paths []string) (dictionary models.Dictionary, cErr c_errors.Error)
} {
	return controllers.texts
}

// Languages - получение контроллера сервиса.
func (controllers *controllers) Languages() interface {
	GetList(ctx context.Context) (list []*models.Language, cErr c_errors.Error)
	Remove(ctx context.Context, code string) (cErr c_errors.Error)
	Update(ctx context.Context, code, name string) (cErr c_errors.Error)
	Create(ctx context.Context, code string, name string) (cErr c_errors.Error)

	Activate(ctx context.Context, code string) (cErr c_errors.Error)
	Deactivate(ctx context.Context, code string) (cErr c_errors.Error)
} {
	return controllers.languages
}
