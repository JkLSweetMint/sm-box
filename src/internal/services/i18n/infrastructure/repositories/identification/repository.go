package identification_repository

import (
	"context"
	common_db_models "sm-box/internal/common/objects/db_models"
	common_entities "sm-box/internal/common/objects/entities"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/postgresql"
)

const (
	loggerInitiator = "infrastructure-[repositories]=identification"
)

// Repository - репозиторий идентификации пользователей.
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
		Format("A '%s' repository has been created. ", "identification").
		Field("config", repo.conf).Write()

	return
}

// GetToken - получение jwt токена.
func (repo *Repository) GetToken(ctx context.Context, data string) (tok *common_entities.JwtToken, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, data)
		defer func() { trc.Error(err).FunctionCallFinished(tok) }()
	}

	var model = new(common_db_models.JwtToken)

	// Получение данных
	{
		var query = `
			select
				tokens.id,
				coalesce(tokens.user_id, 0) as user_id,
				coalesce(tokens.project_id, 0) as project_id,
				tokens.language,
				tokens.issued_at,
				tokens.not_before,
				tokens.expires_at
			from
				access_system.jwt_tokens as tokens
			where
				tokens.data = $1
		`

		var row = repo.connector.QueryRowxContext(ctx, query, data)

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
		tok = new(common_entities.JwtToken)
		tok.FillEmptyFields()

		tok.ID = model.ID
		tok.UserID = model.UserID
		tok.ProjectID = model.ProjectID

		tok.Language = model.Language
		tok.Data = data

		tok.IssuedAt = model.IssuedAt
		tok.NotBefore = model.NotBefore
		tok.ExpiresAt = model.ExpiresAt
	}

	return
}
