package repository

import (
	"context"
	"sm-box/internal/common/entities"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/sqlite3"
)

// usersRepository - часть репозитория с управлением пользователями.
type usersRepository struct {
	connector  sqlite3.Connector
	components *components

	conf *Config
	ctx  context.Context
}

// GetUser - получение пользователя по идентификатору.
func (repo *usersRepository) GetUser(ctx context.Context, id types.ID) (us *entities.User, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(err).FunctionCallFinished(us) }()
	}

	return
}

// BasicAuth - базовая авторизация пользователя в системе.
// Для авторизации используется имя пользователя и пароль.
func (repo *usersRepository) BasicAuth(ctx context.Context, username, password string) (us *entities.User, tok *entities.JwtToken, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, username, password)
		defer func() { trc.Error(err).FunctionCallFinished(us, tok) }()
	}

	return
}
