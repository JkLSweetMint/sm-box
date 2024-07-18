package http_routes_redis_repository

import (
	"context"
	"fmt"
	"regexp"
	"sm-box/internal/services/authentication/objects/entities"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/redis"
	"time"
)

const (
	loggerInitiator = "system-[components]-[access_system]-[repositories]=http_routes_redis"
)

// Repository - репозиторий управления http маршрутами.
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
func New(ctx context.Context, conf *Config) (repo *Repository, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelRepository)

		trc.FunctionCall(ctx, conf)
		defer func() { trc.Error(err).FunctionCallFinished(repo) }()
	}

	repo = &Repository{
		ctx:  ctx,
		conf: conf,
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
		Format("A '%s' repository has been created. ", "http_routes_redis").
		Field("config", repo.conf).Write()

	return
}

// Register - регистрация http маршрута.
func (repo *Repository) Register(ctx context.Context, routes ...*entities.HttpRoute) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, routes)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var expiration = 24 * time.Hour

	for _, route := range routes {
		var (
			key   string
			value any = route
		)

		for _, protocol := range route.Protocols {
			if route.Path != "" {
				key = fmt.Sprintf("http_route:%s:%s:%s", protocol, route.Method, route.Path)
			} else if route.RegexpPath != "" {
				key = fmt.Sprintf("http_route:%s:%s:%s", protocol, route.Method, route.RegexpPath)
			}

			var result = repo.connector.Set(ctx, key, value, expiration)

			if err = result.Err(); err != nil {
				repo.components.Logger.Error().
					Format("Error inserting an item from the database: '%s'. ", err).Write()
				return
			}
		}
	}

	return
}

// Get - получение http маршрута.
func (repo *Repository) Get(ctx context.Context, protocol, method, path string) (route *entities.HttpRoute, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, protocol, method, path)
		defer func() { trc.Error(err).FunctionCallFinished(route) }()
	}

	var (
		prefix  = fmt.Sprintf("http_route:%s:%s:", protocol, method)
		pattern = fmt.Sprintf("%s*", prefix)
		key     = fmt.Sprintf("%s%s", prefix, path)
		result  = repo.connector.Keys(ctx, pattern)
	)

	if err = result.Err(); err != nil {
		repo.components.Logger.Error().
			Format("Error while reading item data from the database:: '%s'. ", err).Write()
		return
	}

	var keys = result.Val()

	for _, k := range keys {
		if ok, _ := regexp.MatchString(fmt.Sprintf("^%s$", k), key); ok {
			var (
				key    = k
				value  = new(entities.HttpRoute)
				result = repo.connector.Get(ctx, key)
			)

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

			route = value

			break
		}
	}

	return
}

// Clear - очистка всех http маршрутов.
func (repo *Repository) Clear(ctx context.Context) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var result = repo.connector.FlushDB(ctx)

	if err = result.Err(); err != nil {
		repo.components.Logger.Error().
			Format("Error while reading item data from the database:: '%s'. ", err).Write()
		return
	}

	return
}
