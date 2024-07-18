package http_routes_redis_repository

import (
	"context"
	"fmt"
	"sm-box/internal/services/authentication/objects/db_models"
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
func (repo *Repository) Register(ctx context.Context, route *entities.HttpRoute) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, route)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var (
		key        string = fmt.Sprintf("http_route:%s:%s", route.Method, route.Path)
		value      any    = route.ToDbModel()
		expiration        = 24 * time.Hour

		result = repo.connector.Set(ctx, key, value, expiration)
	)

	if err = result.Err(); err != nil {
		repo.components.Logger.Error().
			Format("Error inserting an item from the database: '%s'. ", err).Write()
		return
	}

	return
}

// Get - получение http маршрута.
func (repo *Repository) Get(ctx context.Context, method, path string) (route *entities.HttpRoute, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(err).FunctionCallFinished(tok) }()
	}

	var (
		key   string = fmt.Sprintf("jwt_token:%s", id)
		value        = new(db_models.JwtRefreshToken)

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

	// Преобразование в сущность
	{
		tok = &entities.JwtRefreshToken{
			JwtToken: &entities.JwtToken{
				ID:       value.ID,
				ParentID: value.ParentID,

				UserID:    value.UserID,
				ProjectID: value.ProjectID,

				Type: entities.JwtTokenType(value.Type),
				Raw:  "",

				ExpiresAt: value.ExpiresAt,
				NotBefore: value.NotBefore,
				IssuedAt:  value.IssuedAt,

				Params: nil,
			},
		}
	}

	return
}

// Remove - удаление http маршрута.
func (repo *Repository) Remove(ctx context.Context, method, path string) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var (
		key string = fmt.Sprintf("jwt_token:%s", id)

		result = repo.connector.Del(ctx, key)
	)

	if err = result.Err(); err != nil {
		repo.components.Logger.Error().
			Format("Error while reading item data from the database:: '%s'. ", err).Write()
		return
	}

	return
}