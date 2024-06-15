package init_cli

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"github.com/jmoiron/sqlx"
	"path"
	"sm-box/internal/system/init_cli/embed"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	"sm-box/pkg/databases/connectors/sqlite3"
)

// initSystemDB - инициализация системной базы данных.
func (cli_ *cli) initSystemDB(ctx context.Context) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished(ctx) }()
	}

	cli_.components.Logger.Info().
		Text("Starting initialization system db... ").Write()

	var connector sqlite3.Connector

	// Подключение/создание файла
	{
		var conf = new(sqlite3.Config).Default()

		conf.Database = path.Join(env.Paths.Var.Lib, env.Files.Var.Lib.SystemDB)

		if connector, err = sqlite3.New(ctx, conf); err != nil {
			cli_.components.Logger.Error().
				Format("The system database file could not be created: '%s'. ", err).Write()
			return
		}
	}

	// Выполнение миграций
	{
		var query string

		// Чтение файла миграций
		{
			var data []byte

			if data, err = embed.Dir.ReadFile("migrations/system.sql"); err != nil {
				cli_.components.Logger.Error().
					Format("The migration file for the system database could not be read: '%s'. ", err).Write()
				return
			}

			query = string(data)
		}

		if _, err = connector.Exec(query); err != nil {
			cli_.components.Logger.Error().
				Format("Migrations for the system database failed: '%s'. ", err).Write()
			return
		}
	}

	// Создание root пользователя
	{
		var password []byte

		if password, err = rsa.EncryptOAEP(
			sha256.New(),
			rand.Reader,
			env.Vars.EncryptionKeys.Public,
			[]byte("whoifnotme"),
			[]byte("password")); err != nil {
			cli_.components.Logger.Error().
				Format("The encryption of the root user's password failed: '%s'. ", err).Write()
			return
		}

		var tx *sqlx.Tx

		if tx, err = connector.Beginx(); err != nil {
			cli_.components.Logger.Error().
				Format("Failed to create a transaction for the database: '%s'. ", err).Write()
			return
		}

		// Добавление root пользователя
		{
			var query = `
				insert into
					users(
						username,
						password
					) values (
						'root',
						$1
				);
			`

			if _, err = tx.Exec(query, password); err != nil {
				cli_.components.Logger.Error().
					Format("Error inserting an item from the database: '%s'. ", err).Write()
				return
			}
		}

		// Добавление роли для root пользователя
		{
			var query = `
				insert into
					user_accesses(
						  user_id,
						  role_id
					) values (
						  1,
						  1
					);
			`

			if _, err = tx.Exec(query); err != nil {
				cli_.components.Logger.Error().
					Format("Error inserting an item from the database: '%s'. ", err).Write()
				return
			}
		}

		if err = tx.Commit(); err != nil {
			cli_.components.Logger.Error().
				Format("The transaction for the database failed: '%s'. ", err).Write()
			return
		}
	}

	cli_.components.Logger.Info().
		Text("The initialization system db has been successfully finished. ").Write()

	return
}
