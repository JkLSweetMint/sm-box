package urls_usecase

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	error_list "sm-box/internal/common/errors"
	common_types "sm-box/internal/common/types"
	authentication_entities "sm-box/internal/services/authentication/objects/entities"
	urls_repository "sm-box/internal/services/url_shortner/infrastructure/repositories/urls"
	urls_redis_repository "sm-box/internal/services/url_shortner/infrastructure/repositories/urls_redis"
	"sm-box/internal/services/url_shortner/objects/entities"
	"sm-box/internal/services/url_shortner/objects/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator = "infrastructure-[usecases]=urls"
)

// UseCase - логика управления сокращениями url запросов.
type UseCase struct {
	components   *components
	repositories *repositories

	conf *Config
	ctx  context.Context
}

// repositories - репозитории логики.
type repositories struct {
	Urls interface {
		GetActive(ctx context.Context) (list []*entities.ShortUrl, err error)
		WriteCallToHistory(ctx context.Context, id common_types.ID, status types.ShortUrlUsageHistoryStatus, token *authentication_entities.JwtSessionToken) (err error)
	}
	UrlsRedis interface {
		Set(ctx context.Context, list ...*entities.ShortUrl) (err error)
		GetByReduce(ctx context.Context, reduce string) (url *entities.ShortUrl, err error)
		RemoveByReduce(ctx context.Context, reduce string) (err error)
	}
}

// components - компоненты логики.
type components struct {
	Logger logger.Logger
}

// New - создание логики.
func New(ctx context.Context) (usecase *UseCase, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelMain, tracer.LevelUseCase)

		trace.FunctionCall(ctx)

		defer func() { trace.Error(err).FunctionCallFinished(usecase) }()
	}

	usecase = new(UseCase)
	usecase.ctx = ctx

	// Конфигурация
	{
		usecase.conf = new(Config).Default()

		if err = usecase.conf.Read(); err != nil {
			return
		}
	}

	// Компоненты
	{
		usecase.components = new(components)

		// Logger
		{
			if usecase.components.Logger, err = logger.New(loggerInitiator); err != nil {
				return
			}
		}
	}

	// Репозитории
	{
		usecase.repositories = new(repositories)

		// Urls
		{
			if usecase.repositories.Urls, err = urls_repository.New(ctx); err != nil {
				return
			}
		}

		// UrlsRedis
		{
			if usecase.repositories.UrlsRedis, err = urls_redis_repository.New(ctx); err != nil {
				return
			}
		}
	}

	usecase.components.Logger.Info().
		Format("A '%s' usecase has been created. ", "urls").
		Field("config", usecase.conf).Write()

	return
}

// RegisterToRedisDB - регистрация сокращений url в базе данных redis.
func (usecase *UseCase) RegisterToRedisDB(ctx context.Context) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The process of obtaining all active url abbreviations has been started... ").Write()

	var list []*entities.ShortUrl

	// Получение
	{
		var err error

		if list, err = usecase.repositories.Urls.GetActive(ctx); err != nil {
			usecase.components.Logger.Error().
				Format("Could not get all active url abbreviations completed: '%s'. ", err).Write()

			cErr = error_list.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("Short urls have been successfully collected. ").
			Field("list", list).Write()
	}

	// Регистрация в redis
	{
		if len(list) > 0 {
			if err := usecase.repositories.UrlsRedis.Set(ctx, list...); err != nil {
				usecase.components.Logger.Error().
					Format("Failed to register short urls in the redis database: '%s'. ", err).
					Field("list", list).Write()

				cErr = error_list.InternalServerError()
				return
			}
		}
	}

	usecase.components.Logger.Info().
		Text("The process of getting all active url abbreviations is completed. ").Write()

	return
}

// GetByReduceFromRedisDB - получение короткого маршрута по сокращению из базы данных redis.
func (usecase *UseCase) GetByReduceFromRedisDB(ctx context.Context, reduce string) (url *entities.ShortUrl, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, reduce)
		defer func() { trc.Error(cErr).FunctionCallFinished(url) }()
	}

	usecase.components.Logger.Info().
		Text("The process of obtaining a short url for reduction has been launched... ").
		Field("reduce", reduce).Write()

	// Получение
	{
		var err error

		if url, err = usecase.repositories.UrlsRedis.GetByReduce(ctx, reduce); err != nil {
			usecase.components.Logger.Error().
				Format("Could not get all active url abbreviations completed: '%s'. ", err).Write()

			if errors.Is(err, redis.Nil) {
				cErr = error_list.ShortUrlNotFound()
				return
			}

			cErr = error_list.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("Short urls have been successfully collected. ").
			Field("url", url).Write()
	}

	usecase.components.Logger.Info().
		Text("The process of obtaining a short url by reduction is completed. ").
		Field("reduce", reduce).
		Field("url", url).Write()

	return
}

// UpdateInRedisDB - обновление короткого маршрута в базу данных redis.
func (usecase *UseCase) UpdateInRedisDB(ctx context.Context, url *entities.ShortUrl) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, url)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The process of updating the short url has been started... ").
		Field("url", url).Write()

	// Обновление
	{
		var err error

		if err = usecase.repositories.UrlsRedis.Set(ctx, url); err != nil {
			usecase.components.Logger.Error().
				Format("The short url data could not be updated: '%s'. ", err).Write()

			cErr = error_list.InternalServerError()
			return
		}
	}

	usecase.components.Logger.Info().
		Text("The process of updating the short url is completed. ").
		Field("url", url).Write()

	return
}

// RemoveByReduceFromRedisDB - удаление короткого маршрута по сокращению из базы данных redis.
func (usecase *UseCase) RemoveByReduceFromRedisDB(ctx context.Context, reduce string) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, reduce)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The process of deleting a short url has been started... ").
		Field("reduce", reduce).Write()

	// Удаление
	{
		var err error

		if err = usecase.repositories.UrlsRedis.RemoveByReduce(ctx, reduce); err != nil {
			usecase.components.Logger.Error().
				Format("The short url could not be deleted: '%s'. ", err).Write()

			cErr = error_list.InternalServerError()
			return
		}
	}

	usecase.components.Logger.Info().
		Text("The process of deleting the short url is completed. ").
		Field("reduce", reduce).Write()

	return
}

// WriteCallToHistory - записать обращение по короткой ссылке в историю.
func (usecase *UseCase) WriteCallToHistory(ctx context.Context, id common_types.ID, status types.ShortUrlUsageHistoryStatus, token *authentication_entities.JwtSessionToken) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, id, status, token)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The process of recording a short route call in the history has been started... ").
		Field("id", id).
		Field("status", status).
		Field("token", token).Write()

	// Запись в историю
	{
		var err error

		if err = usecase.repositories.Urls.WriteCallToHistory(ctx, id, status, token); err != nil {
			usecase.components.Logger.Error().
				Format("The call data could not be recorded in the history: '%s'. ", err).Write()

			cErr = error_list.InternalServerError()
			return
		}
	}

	usecase.components.Logger.Info().
		Text("The process of recording a short route call in the history is completed. ").
		Field("id", id).
		Field("status", status).
		Field("token", token).Write()

	return
}
