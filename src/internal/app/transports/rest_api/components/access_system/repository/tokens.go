package repository

import (
	"context"
	"sm-box/internal/common/db_models"
	"sm-box/internal/common/entities"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/sqlite3"
	"time"
)

// tokensRepository - часть репозитория с управлением токенами.
type tokensRepository struct {
	connector  sqlite3.Connector
	components *components

	conf *Config
	ctx  context.Context
}

// GetToken - получение jwt токена.
func (repo *tokensRepository) GetToken(ctx context.Context, data string) (tok *entities.JwtToken, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, data)
		defer func() { trc.Error(err).FunctionCallFinished(tok) }()
	}

	// Основные данные
	{
		var model = new(db_models.JwtToken)

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
				system_access_jwt_tokens as tokens
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
			tok = new(entities.JwtToken)
			tok.FillEmptyFields()

			tok.ID = model.ID
			tok.UserID = model.UserID

			tok.Data = data

			if tok.IssuedAt, err = time.Parse(time.RFC3339Nano, model.IssuedAt); err != nil {
				repo.components.Logger.Error().
					Format("Error while reading item data from the database:: '%s'. ", err).Write()
				return
			}

			if tok.NotBefore, err = time.Parse(time.RFC3339Nano, model.NotBefore); err != nil {
				repo.components.Logger.Error().
					Format("Error while reading item data from the database:: '%s'. ", err).Write()
				return
			}

			if tok.ExpiresAt, err = time.Parse(time.RFC3339Nano, model.ExpiresAt); err != nil {
				repo.components.Logger.Error().
					Format("Error while reading item data from the database:: '%s'. ", err).Write()
				return
			}
		}
	}

	// Параметры
	{
		var model = new(db_models.JwtTokenParams)

		// Получение данных
		{
			var query = `
			select
				params.remote_addr,
				params.user_agent
			from
				system_access_jwt_token_params as params
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

// RegisterToken - регистрация jwt токена.
func (repo *tokensRepository) RegisterToken(ctx context.Context, tok *entities.JwtToken) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, tok)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	// Основные данные
	{
		var (
			model = tok.DbModel()
			query = `
			insert into 
				system_access_jwt_tokens (
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
			model = tok.Params.DbModel()
			query = `
			insert into 
				system_access_jwt_token_params (
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

// SetTokenOwner - установить владельца токена.
func (repo *tokensRepository) SetTokenOwner(ctx context.Context, tokenID, ownerID types.ID) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, tokenID, ownerID)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var query = `
		update 
			system_access_jwt_tokens
		set
		    user_id = $1
		where
		    id = $2
	`

	if _, err = repo.connector.ExecContext(ctx, query, ownerID, tokenID); err != nil {
		repo.components.Logger.Error().
			Format("Error updating an item from the database: '%s'. ", err).Write()
		return
	}

	return
}
