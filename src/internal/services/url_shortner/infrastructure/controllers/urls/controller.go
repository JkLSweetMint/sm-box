package urls_controller

import (
	"context"
	common_types "sm-box/internal/common/types"
	authentication_entities "sm-box/internal/services/authentication/objects/entities"
	urls_usecase "sm-box/internal/services/url_shortner/infrastructure/usecases/urls"
	"sm-box/internal/services/url_shortner/objects/entities"
	"sm-box/internal/services/url_shortner/objects/models"
	"sm-box/internal/services/url_shortner/objects/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator = "infrastructure-[controllers]=urls"
)

// Controller - контроллер базовой аутентификации пользователей.
type Controller struct {
	components *components
	usecases   *usecases

	conf *Config
	ctx  context.Context
}

// usecases - логика контроллера.
type usecases struct {
	Urls interface {
		RegisterToRedisDB(ctx context.Context) (cErr c_errors.Error)
		GetByReductionFromRedisDB(ctx context.Context, reduction string) (url *entities.ShortUrl, cErr c_errors.Error)
		UpdateInRedisDB(ctx context.Context, url *entities.ShortUrl) (cErr c_errors.Error)
		RemoveByReductionFromRedisDB(ctx context.Context, reduction string) (cErr c_errors.Error)

		WriteCallToHistory(ctx context.Context, id common_types.ID, status types.ShortUrlUsageHistoryStatus, token *authentication_entities.JwtSessionToken) (cErr c_errors.Error)
	}
}

// components - компоненты контроллера.
type components struct {
	Logger logger.Logger
}

// New - создание контроллера.
func New(ctx context.Context) (controller *Controller, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelMain, tracer.LevelController)

		trace.FunctionCall(ctx)

		defer func() { trace.Error(err).FunctionCallFinished(controller) }()
	}

	controller = new(Controller)
	controller.ctx = ctx

	// Конфигурация
	{
		controller.conf = new(Config).Default()

		if err = controller.conf.Read(); err != nil {
			return
		}
	}

	// Компоненты
	{
		controller.components = new(components)

		// Logger
		{
			if controller.components.Logger, err = logger.New(loggerInitiator); err != nil {
				return
			}
		}
	}

	// Логика
	{
		controller.usecases = new(usecases)

		// Urls
		{
			if controller.usecases.Urls, err = urls_usecase.New(ctx); err != nil {
				return
			}
		}
	}

	controller.components.Logger.Info().
		Format("A '%s' controller has been created. ", "urls").
		Field("config", controller.conf).Write()

	return
}

// RegisterToRedisDB - регистрация сокращений url в базе данных redis.
func (controller *Controller) RegisterToRedisDB(ctx context.Context) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	// Выполнения инструкций
	{
		if cErr = controller.usecases.Urls.RegisterToRedisDB(ctx); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	return
}

// GetByReductionFromRedisDB - получение короткого маршрута по сокращению из базы данных redis.
func (controller *Controller) GetByReductionFromRedisDB(ctx context.Context, reduction string) (url *models.ShortUrlInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, reduction)
		defer func() { trc.Error(cErr).FunctionCallFinished(url) }()
	}

	// Выполнения инструкций
	{
		var url_ *entities.ShortUrl

		if url_, cErr = controller.usecases.Urls.GetByReductionFromRedisDB(ctx, reduction); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}

		// Преобразование в модель
		{
			url = url_.ToModel()
		}
	}

	return
}

// UpdateInRedisDB - обновление короткого маршрута в базу данных redis.
func (controller *Controller) UpdateInRedisDB(ctx context.Context, url *models.ShortUrlInfo) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, url)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	// Выполнения инструкций
	{
		var url_ *entities.ShortUrl

		// Преобразовать в сущность
		{
			url_ = &entities.ShortUrl{
				ID:        url.ID,
				Source:    url.Source,
				Reduction: url.Reduction,

				Properties: &entities.ShortUrlProperties{
					Type:                 url.Properties.Type,
					NumberOfUses:         url.Properties.NumberOfUses,
					RemainedNumberOfUses: url.Properties.RemainedNumberOfUses,
					StartActive:          url.Properties.StartActive,
					EndActive:            url.Properties.EndActive,
				},
			}
		}

		if cErr = controller.usecases.Urls.UpdateInRedisDB(ctx, url_); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	return
}

// RemoveByReductionFromRedisDB - удаление короткого маршрута по сокращению из базы данных redis.
func (controller *Controller) RemoveByReductionFromRedisDB(ctx context.Context, reduction string) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, reduction)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	// Выполнения инструкций
	{
		if cErr = controller.usecases.Urls.RemoveByReductionFromRedisDB(ctx, reduction); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	return
}

// WriteCallToHistory - записать обращение по короткой ссылке в историю.
func (controller *Controller) WriteCallToHistory(ctx context.Context, id common_types.ID, status types.ShortUrlUsageHistoryStatus, token *authentication_entities.JwtSessionToken) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, id, status, token)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	// Выполнения инструкций
	{
		if cErr = controller.usecases.Urls.WriteCallToHistory(ctx, id, status, token); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	return
}
