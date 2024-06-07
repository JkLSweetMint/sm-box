package init_script

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"errors"
	"github.com/jmoiron/sqlx"
	"os"
	"path"
	"sm-box/embed"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	"sm-box/pkg/databases/connectors/sqlite3"
)

// serve - внутренний метод для запуска скрипта.
func (scr *script) serve(ctx context.Context) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}
	// Завершение работы
	{
		defer func() {
			var failed = err != nil

			if err = scr.core.Shutdown(); err != nil {
				scr.components.Logger.Error().
					Format("An error occurred when starting maintenance of the '%s': '%s'. ",
						env.Vars.SystemName,
						err).Write()
				return
			}

			if failed {
				os.Exit(1)
			}
		}()
	}

	scr.components.Logger.Info().
		Format("Starting the '%s'... ", env.Vars.SystemName).Write()

	// Логика
	{
		// Проверки что уже инициализировано
		{
			if _, err = os.Stat(path.Join(env.Paths.SystemLocation, env.Paths.Var.Lib, env.Files.Var.Lib.SystemDB)); errors.Is(err, os.ErrNotExist) {
				err = nil
			} else {
				if err == nil {
					scr.components.Logger.Info().
						Format("'%s' is initialized, no reinitialization is required. ", env.Vars.SystemName).Write()
				} else {
					scr.components.Logger.Error().
						Format("Failed to initialize '%s': '%s'. ", env.Vars.SystemName, err).Write()
				}

				return
			}
		}

		scr.components.Logger.Info().
			Format("Starting initialization '%s'... ", env.Vars.SystemName).Write()

		if err = scr.initSystemDB(ctx); err != nil {
			scr.components.Logger.Error().
				Format("Failed to initialize the system database: '%s'. ", err).Write()
			return
		}

		scr.components.Logger.Info().
			Format("The initialization '%s' has been successfully finished. ", env.Vars.SystemName).Write()
	}

	scr.components.Logger.Info().
		Format("The '%s' has been successfully started. ", env.Vars.SystemName).Write()

	return
}

// initSystemDB - инициализация системной базы данных.
func (scr *script) initSystemDB(ctx context.Context) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelInternal)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished(ctx) }()
	}

	scr.components.Logger.Info().
		Text("Starting initialization system db... ").Write()

	var connector sqlite3.Connector

	// Подключение/создание файла
	{
		var conf = new(sqlite3.Config).Default()

		conf.Database = path.Join(env.Paths.Var.Lib, env.Files.Var.Lib.SystemDB)

		if connector, err = sqlite3.New(ctx, conf); err != nil {
			scr.components.Logger.Error().
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
				scr.components.Logger.Error().
					Format("The migration file for the system database could not be read: '%s'. ", err).Write()
				return
			}

			query = string(data)
		}

		if _, err = connector.Exec(query); err != nil {
			scr.components.Logger.Error().
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
			scr.components.Logger.Error().
				Format("The encryption of the root user's password failed: '%s'. ", err).Write()
			return
		}

		var tx *sqlx.Tx

		if tx, err = connector.Beginx(); err != nil {
			scr.components.Logger.Error().
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
				scr.components.Logger.Error().
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
				scr.components.Logger.Error().
					Format("Error inserting an item from the database: '%s'. ", err).Write()
				return
			}
		}

		if err = tx.Commit(); err != nil {
			scr.components.Logger.Error().
				Format("The transaction for the database failed: '%s'. ", err).Write()
			return
		}
	}

	scr.components.Logger.Info().
		Text("The initialization system db has been successfully finished. ").Write()

	return
}
