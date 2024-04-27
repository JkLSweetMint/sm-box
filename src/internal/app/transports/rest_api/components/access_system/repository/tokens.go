package repository

import (
	"context"
	"sm-box/internal/common/entities"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/sqlite3"
)

// tokensRepository - часть репозитория с управлением токенами.
type tokensRepository struct {
	connector  sqlite3.Connector
	components *components

	conf *Config
	ctx  context.Context
}

// GetToken - получение jwt токена.
func (repo *tokensRepository) GetToken(ctx context.Context, data []byte) (tok *entities.JwtToken, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, data)
		defer func() { trc.Error(err).FunctionCallFinished(tok) }()
	}

	return
}

// GetUserToken - получение jwt токена по идентификатору пользователя.
func (repo *tokensRepository) GetUserToken(ctx context.Context, userID types.ID) (tok *entities.JwtToken, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, userID)
		defer func() { trc.Error(err).FunctionCallFinished(tok) }()
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

	return
}
