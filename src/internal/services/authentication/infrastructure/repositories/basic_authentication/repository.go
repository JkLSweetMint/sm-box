package basic_authentication_repository

import (
	"context"
	"sm-box/internal/services/authentication/objects/db_models"
	"sm-box/internal/services/authentication/objects/entities"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/postgresql"
)

const (
	loggerInitiator = "infrastructure-[repositories]=basic_authentication"
)

// Repository - репозиторий базовой аутентификации пользователей.
type Repository struct {
	connector  postgresql.Connector
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
		if repo.connector, err = postgresql.New(ctx, repo.conf.Connector); err != nil {
			return
		}
	}

	repo.components.Logger.Info().
		Format("A '%s' repository has been created. ", "basic_authentication").
		Field("config", repo.conf).Write()

	return
}

// GetToken - получение jwt токена.
func (repo *Repository) GetToken(ctx context.Context, raw string) (tok *entities.JwtToken, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, raw)
		defer func() { trc.Error(err).FunctionCallFinished(tok) }()
	}

	// Основные данные
	{
		var model = new(db_models.JwtToken)

		// Получение
		{
			var query = `
			select
				tokens.id,
				tokens.parent_id,
				coalesce(tokens.user_id, 0) as user_id,
				coalesce(tokens.project_id, 0) as project_id,
				tokens.type,
				tokens.issued_at,
				tokens.not_before,
				tokens.expires_at
			from
				tokens.jwt as tokens
			where
				tokens.raw = $1
		`

			var row = repo.connector.QueryRowxContext(ctx, query, raw)

			if err = row.Err(); err != nil {
				repo.components.Logger.Error().
					Format("Error when retrieving an item from the database: '%s'. ", err).Write()
				return
			}

			if err = row.StructScan(model); err != nil {
				repo.components.Logger.Error().
					Format("Error while reading item data from the database:: '%s'. ", err).Write()
				return
			}
		}

		// Перенос в сущность
		{
			tok = new(entities.JwtToken)
			tok.FillEmptyFields()

			tok.ID = model.ID
			tok.ParentID = model.ParentID

			tok.UserID = model.UserID
			tok.ProjectID = model.ProjectID

			tok.Raw = raw

			tok.IssuedAt = model.IssuedAt
			tok.NotBefore = model.NotBefore
			tok.ExpiresAt = model.ExpiresAt
		}
	}

	// Параметры
	{
		var model = new(db_models.JwtTokenParams)

		// Получение
		{
			var query = `
			select
				params.remote_addr,
				params.user_agent
			from
				tokens.jwt_params as params
			where
				params.token_id = $1
		`

			var row = repo.connector.QueryRowxContext(ctx, query, tok.ID)

			if err = row.Err(); err != nil {
				repo.components.Logger.Error().
					Format("Error when retrieving an item from the database: '%s'. ", err).Write()
				return
			}

			if err = row.StructScan(model); err != nil {
				repo.components.Logger.Error().
					Format("Error while reading item data from the database:: '%s'. ", err).Write()
				return
			}
		}

		// Перенос в сущность
		{
			tok.Params = new(entities.JwtTokenParams)

			tok.Params.RemoteAddr = model.RemoteAddr
			tok.Params.UserAgent = model.UserAgent
		}
	}

	return
}

// Register - регистрация jwt токена.
func (repo *Repository) Register(ctx context.Context, tok *entities.JwtToken) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, tok)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	// Основные данные
	{
		var (
			model = tok.ToDbModel()
			query = `
			insert into 
				tokens.jwt (
						id,
						parent_id,
				        user_id,
						project_id,
						type,
						raw, 
						expires_at, 
						not_before,
						issued_at
					) values (
						$1,
						$2,
						nullif($3, 0),
						nullif($4, 0),
						$5,
						$6,
						$7,
						$8,
						$9
					);
		`
		)

		if _, err = repo.connector.ExecContext(ctx, query,
			model.ID,
			model.ParentID,
			model.UserID,
			model.ProjectID,
			model.Type,
			model.Raw,
			model.ExpiresAt,
			model.NotBefore,
			model.IssuedAt); err != nil {
			repo.components.Logger.Error().
				Format("Error inserting an item from the database: '%s'. ", err).Write()
			return
		}
	}

	// Параметры
	{
		var (
			model = tok.Params.ToDbModel()
			query = `
			insert into 
				tokens.jwt_params (
						token_id, 
						remote_addr, 
						user_agent
					) values (
						$1,
						$2,
						$3
					)
		`
		)

		model.TokenID = tok.ID

		if _, err = repo.connector.ExecContext(ctx, query,
			model.TokenID,
			model.RemoteAddr,
			model.UserAgent); err != nil {
			repo.components.Logger.Error().
				Format("Error inserting an item from the database: '%s'. ", err).Write()
			return
		}
	}

	return
}
