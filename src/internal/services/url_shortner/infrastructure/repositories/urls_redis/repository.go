package urls_redis_repository

import (
	"context"
	"fmt"
	"sm-box/internal/services/url_shortner/objects/db_models"
	"sm-box/internal/services/url_shortner/objects/entities"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/redis"
	"time"
)

const (
	loggerInitiator = "infrastructure-[repositories]=urls_redis"
)

// Repository - репозиторий для работы с сокращенными url запросов в базе данных Redis.
type Repository struct {
	connector  redis.Connector
	components *components

	conf *Config
	ctx  context.Context
}

// components - компоненты репозитория.
type components struct {
	Logger logger.Logger
}

// New - создание репозитория.
func New(ctx context.Context) (repo *Repository, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelRepository)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished(repo) }()
	}

	repo = &Repository{
		ctx: ctx,
	}

	// Конфигурация
	{
		repo.conf = new(Config).Default()

		if err = repo.conf.Read(); err != nil {
			return
		}
	}

	// Компоненты
	{
		repo.components = new(components)

		// Logger
		{
			if repo.components.Logger, err = logger.New(loggerInitiator); err != nil {
				return
			}
		}
	}

	// Коннектор базы данных
	{
		if repo.connector, err = redis.New(ctx, repo.conf.Connector); err != nil {
			return
		}
	}

	repo.components.Logger.Info().
		Format("A '%s' repository has been created. ", "urls_redis").
		Field("config", repo.conf).Write()

	return
}

// Set - установить значение коротких маршрутов.
func (repo *Repository) Set(ctx context.Context, list ...*entities.ShortUrl) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, list)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var emptyTm time.Time

	for _, url := range list {
		var (
			key        string = fmt.Sprintf("short_url:%s", url.Reduction)
			value      any    = url.ToRedisDbModel()
			expiration time.Duration
		)

		if !url.Properties.EndActive.Equal(emptyTm) {
			expiration = url.Properties.EndActive.Sub(time.Now())
		}

		var result = repo.connector.Set(ctx, key, value, expiration)

		if err = result.Err(); err != nil {
			repo.components.Logger.Error().
				Format("Error inserting an item from the database: '%s'. ", err).Write()
			return
		}
	}

	return
}

// GetOneByReduction - получение короткого маршрута по сокращению.
func (repo *Repository) GetOneByReduction(ctx context.Context, reduction string) (url *entities.ShortUrl, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, reduction)
		defer func() { trc.Error(err).FunctionCallFinished(url) }()
	}

	var (
		key   string = fmt.Sprintf("short_url:%s", reduction)
		value        = new(db_models.ShortUrlInfo)

		result = repo.connector.Get(ctx, key)
	)

	// Выполнение запроса
	{
		if err = result.Err(); err != nil {
			repo.components.Logger.Error().
				Format("Error while reading item data from the database:: '%s'. ", err).Write()
			return
		}

		if err = result.Scan(value); err != nil {
			repo.components.Logger.Error().
				Format("Error while reading item data from the database:: '%s'. ", err).Write()
			return
		}
	}

	// Преобразование в сущность
	{
		url = &entities.ShortUrl{
			ID:        value.ID,
			Source:    value.Source,
			Reduction: value.Reduction,

			Properties: &entities.ShortUrlProperties{
				Type:                 value.Properties.Type,
				NumberOfUses:         value.Properties.NumberOfUses,
				RemainedNumberOfUses: value.Properties.RemainedNumberOfUses,
				StartActive:          value.Properties.StartActive,
				EndActive:            value.Properties.EndActive,
				Active:               true,
			},
		}
	}

	return
}

// RemoveByReduction - удаление короткого маршрута по сокращению.
func (repo *Repository) RemoveByReduction(ctx context.Context, reduction string) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, reduction)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var (
		key string = fmt.Sprintf("short_url:%s", reduction)

		result = repo.connector.Del(ctx, key)
	)

	if err = result.Err(); err != nil {
		repo.components.Logger.Error().
			Format("Error while reading item data from the database:: '%s'. ", err).Write()
		return
	}

	return
}
