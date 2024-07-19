package jwt_tokens_repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"sm-box/internal/services/authentication/objects/db_models"
	"sm-box/internal/services/authentication/objects/entities"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/redis"
	"time"
)

const (
	loggerInitiator = "infrastructure-[repositories]=jwt_tokens"
)

// Repository - репозиторий управления токенами пользователей.
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
		Format("A '%s' repository has been created. ", "jwt_tokens").
		Field("config", repo.conf).Write()

	return
}

// RegisterJwtRefreshToken - регистрация jwt токена обновления.
func (repo *Repository) RegisterJwtRefreshToken(ctx context.Context, tok *entities.JwtRefreshToken) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, tok)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var (
		key        string = fmt.Sprintf("jwt_token:%s", tok.ID)
		value      any    = tok.ToRedisDbModel()
		expiration        = tok.ExpiresAt.Sub(time.Now())

		result = repo.connector.Set(ctx, key, value, expiration)
	)

	if err = result.Err(); err != nil {
		repo.components.Logger.Error().
			Format("Error inserting an item from the database: '%s'. ", err).Write()
		return
	}

	return
}

// GetJwtRefreshToken - получение jwt токена обновления.
func (repo *Repository) GetJwtRefreshToken(ctx context.Context, id uuid.UUID) (tok *entities.JwtRefreshToken, err error) {
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

// RegisterJwtAccessToken - регистрация jwt токена доступа.
func (repo *Repository) RegisterJwtAccessToken(ctx context.Context, tok *entities.JwtAccessToken) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, tok)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var (
		key        string = fmt.Sprintf("jwt_token:%s", tok.ID)
		value      any    = tok.ToRedisDbModel()
		expiration        = tok.ExpiresAt.Sub(time.Now())

		result = repo.connector.Set(ctx, key, value, expiration)
	)

	if err = result.Err(); err != nil {
		repo.components.Logger.Error().
			Format("Error inserting an item from the database: '%s'. ", err).Write()
		return
	}

	return
}

// GetJwtAccessToken - получение jwt токена доступа.
func (repo *Repository) GetJwtAccessToken(ctx context.Context, id uuid.UUID) (tok *entities.JwtAccessToken, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(err).FunctionCallFinished(tok) }()
	}

	var (
		key   string = fmt.Sprintf("jwt_token:%s", id)
		value        = new(db_models.JwtAccessToken)

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
		tok = &entities.JwtAccessToken{
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
			UserInfo: &entities.JwtAccessTokenUserInfo{
				Accesses: value.UserInfo.Accesses,
			},
		}
	}

	return
}

// Remove - удаление токена.
func (repo *Repository) Remove(ctx context.Context, id uuid.UUID) (err error) {
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
