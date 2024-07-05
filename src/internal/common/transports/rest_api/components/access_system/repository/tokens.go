package repository

import (
	"context"
	common_db_models "sm-box/internal/common/objects/db_models"
	"sm-box/internal/common/objects/entities"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/postgresql"
)

// tokensRepository - часть репозитория с управлением токенами.
type tokensRepository struct {
	connector  postgresql.Connector
	components *components

	conf *Config
	ctx  context.Context
}

// GetToken - получение jwt токена.
func (repo *tokensRepository) GetToken(ctx context.Context, data string) (tok *common_entities.JwtToken, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, data)
		defer func() { trc.Error(err).FunctionCallFinished(tok) }()
	}

	// Основные данные
	{
		var model = new(common_db_models.JwtToken)

		// Получение данных
		{
			var query = `
			select
				tokens.id,
				coalesce(tokens.user_id, 0) as user_id,
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

			tok.Data = data

			tok.IssuedAt = model.IssuedAt
			tok.NotBefore = model.NotBefore
			tok.ExpiresAt = model.ExpiresAt
		}
	}

	// Параметры
	{
		var model = new(common_db_models.JwtTokenParams)

		// Получение данных
		{
			var query = `
			select
				params.remote_addr,
				params.user_agent
			from
				access_system.jwt_token_params as params
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
			tok.Params = new(common_entities.JwtTokenParams)

			tok.Params.RemoteAddr = model.RemoteAddr
			tok.Params.UserAgent = model.UserAgent
		}
	}

	return
}

// RegisterToken - регистрация jwt токена.
func (repo *tokensRepository) RegisterToken(ctx context.Context, tok *common_entities.JwtToken) (err error) {
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
				access_system.jwt_tokens (
						data, 
						expires_at, 
						not_before,
						issued_at
					) values (
						$1,
						$2,
						$3,
						$4
					)
			returning id;
		`
		)

		var row = repo.connector.QueryRowxContext(ctx, query,
			model.Data,
			model.ExpiresAt,
			model.NotBefore,
			model.IssuedAt)

		if err = row.Err(); err != nil {
			repo.components.Logger.Error().
				Format("Error when retrieving an item from the database: '%s'. ", err).Write()
			return
		}

		if err = row.Scan(&tok.ID); err != nil {
			repo.components.Logger.Error().
				Format("Error while reading item data from the database:: '%s'. ", err).Write()
			return
		}
	}

	// Параметры
	{
		var (
			model = tok.Params.ToDbModel()
			query = `
			insert into 
				access_system.jwt_token_params (
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
