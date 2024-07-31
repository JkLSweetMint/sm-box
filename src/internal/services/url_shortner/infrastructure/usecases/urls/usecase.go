package urls_usecase

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	common_errors "sm-box/internal/common/errors"
	common_types "sm-box/internal/common/types"
	authentication_entities "sm-box/internal/services/authentication/objects/entities"
	urls_repository "sm-box/internal/services/url_shortner/infrastructure/repositories/urls"
	urls_redis_repository "sm-box/internal/services/url_shortner/infrastructure/repositories/urls_redis"
	"sm-box/internal/services/url_shortner/objects/entities"
	srv_errors "sm-box/internal/services/url_shortner/objects/errors"
	"sm-box/internal/services/url_shortner/objects/types"
	access_system_service_gateway "sm-box/internal/services/url_shortner/transport/gateways/access_system"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	"time"
)

const (
	loggerInitiator = "infrastructure-[usecases]=urls"
)

// UseCase - логика управления сокращениями url запросов.
type UseCase struct {
	components   *components
	gateways     *gateways
	repositories *repositories

	conf *Config
	ctx  context.Context
}

// gateways - шлюзы логики.
type gateways struct {
	AccessSystem interface {
		CheckUserAccess(ctx context.Context, userID common_types.ID, rolesID, permissionsID []common_types.ID) (allowed bool, cErr c_errors.Error)
	}
}

// repositories - репозитории логики.
type repositories struct {
	Urls interface {
		GetActive(ctx context.Context) (list []*entities.ShortUrl, err error)
		WriteCallToHistory(ctx context.Context, id common_types.ID, status types.ShortUrlUsageHistoryStatus, token *authentication_entities.JwtSessionToken) (err error)
	}
	UrlsRedis interface {
		Set(ctx context.Context, list ...*entities.ShortUrl) (err error)
		GetOneByReduction(ctx context.Context, reduction string) (url *entities.ShortUrl, err error)
		RemoveByReduction(ctx context.Context, reduction string) (err error)
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

	// Шлюзы
	{
		usecase.gateways = new(gateways)

		// AccessSystem
		{
			if usecase.gateways.AccessSystem, err = access_system_service_gateway.New(ctx); err != nil {
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
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	var list []*entities.ShortUrl

	usecase.components.Logger.Info().
		Text("The process of obtaining all active url abbreviations has been started... ").Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of getting all active url abbreviations is completed. ").
			Field("list", list).Write()
	}()

	// Получение
	{
		var err error

		if list, err = usecase.repositories.Urls.GetActive(ctx); err != nil {
			usecase.components.Logger.Error().
				Format("Could not get all active url abbreviations completed: '%s'. ", err).Write()

			cErr = common_errors.InternalServerError()
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

				cErr = common_errors.InternalServerError()
				return
			}

			usecase.components.Logger.Info().
				Text("Short URLs have been successfully registered. ").
				Field("list", list).Write()
		}
	}

	return
}

// GetByReductionFromRedisDB - получение короткого маршрута по сокращению из базы данных redis.
func (usecase *UseCase) GetByReductionFromRedisDB(ctx context.Context, reduction string) (url *entities.ShortUrl, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, reduction)
		defer func() { trc.Error(cErr).FunctionCallFinished(url) }()
	}

	usecase.components.Logger.Info().
		Text("The process of obtaining a short url for reduction has been launched... ").
		Field("reduction", reduction).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of obtaining a short url by reduction is completed. ").
			Field("reduction", reduction).
			Field("url", url).Write()
	}()

	// Получение
	{
		var err error

		if url, err = usecase.repositories.UrlsRedis.GetOneByReduction(ctx, reduction); err != nil {
			usecase.components.Logger.Error().
				Format("Could not get the shortened url by reduction: '%s'. ", err).Write()

			if errors.Is(err, redis.Nil) {
				cErr = srv_errors.ShortUrlNotFound()
				return
			}

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("Short url have been successfully collected. ").
			Field("url", url).Write()
	}

	return
}

// UpdateInRedisDB - обновление короткого маршрута в базу данных redis.
func (usecase *UseCase) UpdateInRedisDB(ctx context.Context, url *entities.ShortUrl) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, url)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The process of updating the short url has been started... ").
		Field("url", url).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of updating the short url is completed. ").
			Field("url", url).Write()
	}()

	// Обновление
	{
		if err := usecase.repositories.UrlsRedis.Set(ctx, url); err != nil {
			usecase.components.Logger.Error().
				Format("The short url data could not be updated: '%s'. ", err).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The short url data has been successfully recorded. ").
			Field("url", url).Write()
	}

	return
}

// RemoveByReductionFromRedisDB - удаление короткого маршрута по сокращению из базы данных redis.
func (usecase *UseCase) RemoveByReductionFromRedisDB(ctx context.Context, reduction string) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, reduction)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The process of deleting a short url has been started... ").
		Field("reduction", reduction).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of deleting the short url is completed. ").
			Field("reduction", reduction).Write()
	}()

	// Удаление
	{
		if err := usecase.repositories.UrlsRedis.RemoveByReduction(ctx, reduction); err != nil {
			usecase.components.Logger.Error().
				Format("The short url could not be deleted: '%s'. ", err).
				Field("reduction", reduction).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The short url has been successfully deleted by reduction. ").
			Field("reduction", reduction).Write()
	}

	return
}

// Use - использование короткой ссылки.
func (usecase *UseCase) Use(ctx context.Context, reduction string, token *authentication_entities.JwtSessionToken) (url *entities.ShortUrl, status types.ShortUrlUsageHistoryStatus, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, reduction, token)
		defer func() { trc.Error(cErr).FunctionCallFinished(url, status) }()
	}

	usecase.components.Logger.Info().
		Text("The process of using a short url has been started... ").
		Field("reduction", reduction).
		Field("token", token).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of using the short url is completed. ").
			Field("reduction", reduction).
			Field("token", token).Write()
	}()

	// Запись в историю
	{
		defer func() {
			go func() {
				if cErr != nil {
					url = nil
				}

				if url != nil {
					if err := usecase.repositories.Urls.WriteCallToHistory(ctx, url.ID, status, token); err != nil {
						usecase.components.Logger.Error().
							Format("The call data could not be recorded in the history: '%s'. ", err).Write()

						cErr = common_errors.InternalServerError()
						return
					}

					usecase.components.Logger.Info().
						Text("The use of the short url has been successfully recorded in the history. ").
						Field("url", url).
						Field("token", token).Write()
				}
			}()
		}()
	}

	// Получение и обработка короткого url
	{
		// Получение
		{
			var err error

			if url, err = usecase.repositories.UrlsRedis.GetOneByReduction(ctx, reduction); err != nil {
				usecase.components.Logger.Error().
					Format("Could not get the shortened url by reduction: '%s'. ", err).Write()

				if errors.Is(err, redis.Nil) {
					cErr = srv_errors.ShortUrlNotFound()

					status = types.ShortUrlUsageHistoryStatusForbidden

					return
				}

				status = types.ShortUrlUsageHistoryStatusFailed

				cErr = common_errors.InternalServerError()
				return
			}

			usecase.components.Logger.Info().
				Text("The short url data was successfully received. ").
				Field("url", url).Write()
		}

		// Проверки
		{
			var (
				tm      = time.Now()
				emptyTm time.Time
			)

			// Ещё не начал действовать
			{
				if tm.Before(url.Properties.StartActive) && !url.Properties.StartActive.Equal(emptyTm) {
					usecase.components.Logger.Warn().
						Text("The validity period of the short url has not yet begun. ").Write()

					return
				}
			}

			// Уже закончился
			{
				if tm.After(url.Properties.EndActive) && !url.Properties.EndActive.Equal(emptyTm) {
					usecase.components.Logger.Warn().
						Text("The validity period of the short url has already been completed. ").Write()

					// Удаление из базы данных Redis
					{
						if err := usecase.repositories.UrlsRedis.RemoveByReduction(ctx, url.Reduction); err != nil {
							usecase.components.Logger.Error().
								Format("The short url could not be deleted: '%s'. ", err).
								Field("reduction", reduction).Write()

							cErr = common_errors.InternalServerError()
							return
						}

						usecase.components.Logger.Info().
							Text("The short url has been successfully deleted by reduction. ").
							Field("reduction", reduction).Write()
					}

					return
				}
			}

			// Кол-во использований превышено
			{
				if url.Properties.RemainedNumberOfUses == 0 && url.Properties.NumberOfUses > 0 {
					usecase.components.Logger.Warn().
						Text("The number of uses of the short url  is overestimated. ").Write()

					// Удаление из базы данных Redis
					{
						if err := usecase.repositories.UrlsRedis.RemoveByReduction(ctx, url.Reduction); err != nil {
							usecase.components.Logger.Error().
								Format("The short url could not be deleted: '%s'. ", err).
								Field("reduction", reduction).Write()

							cErr = common_errors.InternalServerError()
							return
						}

						usecase.components.Logger.Info().
							Text("The short url has been successfully deleted by reduction. ").
							Field("reduction", reduction).Write()
					}

					status = types.ShortUrlUsageHistoryStatusForbidden
					return
				}
			}

			// Доступов пользователя
			{
				var allowed bool

				if url != nil && url.Accesses != nil && token != nil {
					if allowed, cErr = usecase.gateways.AccessSystem.CheckUserAccess(ctx, token.UserID, url.Accesses.RolesID, url.Accesses.PermissionsID); cErr != nil {
						usecase.components.Logger.Error().
							Format("The user's access to the short url could not be verified: '%s'. ", cErr).Write()

						status = types.ShortUrlUsageHistoryStatusFailed

						cErr = common_errors.InternalServerError()
						return
					}

					usecase.components.Logger.Info().
						Text("Verification of the user's access to the short url has been completed successfully. ").
						Field("allowed", allowed).
						Field("url", url).
						Field("token", token).Write()
				}

				if !allowed {
					usecase.components.Logger.Warn().
						Text("The user does not have access to the short url. ").
						Field("token", token).
						Field("url", url).Write()

					status = types.ShortUrlUsageHistoryStatusForbidden
					return
				}
			}
		}
	}

	// Выполнение инструкций
	{
		// Обновление данных в базе если кол-во использований не бесконечное
		{
			if url.Properties.NumberOfUses > 0 {
				url.Properties.RemainedNumberOfUses--

				if url.Properties.RemainedNumberOfUses == 0 {
					if err := usecase.repositories.UrlsRedis.RemoveByReduction(ctx, url.Reduction); err != nil {
						usecase.components.Logger.Error().
							Format("The short url could not be deleted: '%s'. ", err).
							Field("reduction", reduction).Write()

						status = types.ShortUrlUsageHistoryStatusFailed

						cErr = common_errors.InternalServerError()
						return
					}

					usecase.components.Logger.Info().
						Text("The short url has been successfully deleted by reduction. ").
						Field("reduction", reduction).Write()
				} else {
					if err := usecase.repositories.UrlsRedis.Set(ctx, url); err != nil {
						usecase.components.Logger.Error().
							Format("The short url data could not be updated: '%s'. ", err).Write()

						status = types.ShortUrlUsageHistoryStatusFailed

						cErr = common_errors.InternalServerError()
						return
					}

					usecase.components.Logger.Info().
						Text("The short url data has been successfully recorded. ").
						Field("url", url).Write()
				}
			}
		}
	}

	status = types.ShortUrlUsageHistoryStatusSuccess

	return
}
