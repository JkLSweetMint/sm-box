package initialization

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"errors"
	"github.com/jmoiron/sqlx"
	"os"
	"path"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/system/init_cli/embed"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	"sm-box/pkg/databases/connectors/sqlite3"
	c_errors "sm-box/pkg/errors"
)

var (
	initFile = path.Join(env.Paths.SystemLocation, env.Paths.System.Path, ".init")
)

// UseCase - логика инициализации.
type UseCase struct {
	components *components

	conf *Config
	ctx  context.Context
}

// components - компоненты логики.
type components struct {
	Logger logger.Logger
}

// New - создание логики.
func New(ctx context.Context) (usecase *UseCase, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelMain)

		trace.FunctionCall(ctx)

		defer func() { trace.Error(err).FunctionCallFinished(usecase) }()
	}

	usecase = new(UseCase)
	usecase.ctx = ctx

	// Конфигурация
	{
		usecase.conf = new(Config).Default()

		if err = usecase.conf.Read(); err != nil {
			return
		}
	}

	// Компоненты
	{
		usecase.components = new(components)

		// Logger
		{
			if usecase.components.Logger, err = logger.New(env.Vars.SystemName); err != nil {
				return
			}
		}
	}

	usecase.components.Logger.Info().
		Format("A '%s' usecase has been created. ", "initialization").
		Field("config", usecase.conf).Write()

	return
}

// Initialize - инициализировать систему.
func (usecase *UseCase) Initialize(ctx context.Context) (cErr c_errors.Error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelController)

		trace.FunctionCall(ctx)

		defer func() { trace.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Format("Starting initialization '%s' CLI... ", env.Vars.SystemName).Write()

	// Проверка что уже инициализировано (существует init файл)
	{
		if _, err := os.Stat(initFile); errors.Is(err, os.ErrNotExist) {
			err = nil
		} else {
			if err == nil {
				usecase.components.Logger.Error().
					Text("The system is initialized, no initialization is required. ").Write()
			} else {
				cErr = error_list.FailedToInitializeSystem()
				cErr.SetError(err)

				usecase.components.Logger.Error().
					Format("Failed to initialize system: '%s'. ", err).Write()
			}

			return
		}
	}

	// Процесс создания базы данных
	{
		var (
			connector sqlite3.Connector
			err       error
		)

		usecase.components.Logger.Info().
			Text("Starting initialization system db... ").Write()

		// Подключение/создание файла
		{
			var conf = new(sqlite3.Config).Default()

			conf.Database = path.Join(env.Paths.Var.Lib.Path, env.Files.Var.Lib.SystemDB)

			if connector, err = sqlite3.New(ctx, conf); err != nil {
				cErr = error_list.FailedToInitializeSystem()
				cErr.SetError(err)

				usecase.components.Logger.Error().
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
					cErr = error_list.FailedToInitializeSystem()
					cErr.SetError(err)

					usecase.components.Logger.Error().
						Format("The migration file for the system database could not be read: '%s'. ", err).Write()
					return
				}

				query = string(data)
			}

			if _, err = connector.Exec(query); err != nil {
				cErr = error_list.FailedToInitializeSystem()
				cErr.SetError(err)

				usecase.components.Logger.Error().
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
				[]byte("toor"),
				[]byte("password")); err != nil {
				cErr = error_list.FailedToInitializeSystem()
				cErr.SetError(err)

				usecase.components.Logger.Error().
					Format("The encryption of the root user's password failed: '%s'. ", err).Write()
				return
			}

			var tx *sqlx.Tx

			if tx, err = connector.Beginx(); err != nil {
				cErr = error_list.FailedToInitializeSystem()
				cErr.SetError(err)

				usecase.components.Logger.Error().
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
					cErr = error_list.FailedToInitializeSystem()
					cErr.SetError(err)

					usecase.components.Logger.Error().
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
					cErr = error_list.FailedToInitializeSystem()
					cErr.SetError(err)

					usecase.components.Logger.Error().
						Format("Error inserting an item from the database: '%s'. ", err).Write()
					return
				}
			}

			if err = tx.Commit(); err != nil {
				cErr = error_list.FailedToInitializeSystem()
				cErr.SetError(err)

				usecase.components.Logger.Error().
					Format("The transaction for the database failed: '%s'. ", err).Write()
				return
			}
		}

		usecase.components.Logger.Info().
			Text("The initialization system db has been successfully finished. ").Write()
	}

	// Создание init файла
	{
		if err := os.WriteFile(initFile, []byte{}, 0666); err != nil {
			cErr = error_list.FailedToInitializeSystem()
			cErr.SetError(err)

			usecase.components.Logger.Error().
				Format("Failed to initialize system: '%s'. ", err).Write()
			return
		}
	}

	usecase.components.Logger.Info().
		Format("The initialization '%s' CLI has been successfully finished. ", env.Vars.SystemName).Write()

	return
}

// Clear - очистить систему.
func (usecase *UseCase) Clear(ctx context.Context) (cErr c_errors.Error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelController)

		trace.FunctionCall(ctx)

		defer func() { trace.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("Cleaning the system... ")

	// Очистка системной базы данных
	{
		if err := os.Remove(path.Join(env.Paths.SystemLocation, env.Paths.Var.Lib.Path, env.Files.Var.Lib.SystemDB)); err != nil {
			cErr = error_list.SystemCleanupError()
			cErr.SetError(err)

			usecase.components.Logger.Error().
				Format("The system could not be cleaned: '%s'. ", err).Write()
			return
		}
	}

	// Очистка проектов
	{
		if err := os.RemoveAll(path.Join(env.Paths.SystemLocation, env.Paths.Var.Lib.Projects)); err != nil {
			cErr = error_list.SystemCleanupError()
			cErr.SetError(err)

			usecase.components.Logger.Error().
				Format("The system could not be cleaned: '%s'. ", err).Write()
			return
		}
	}

	// Удалить .init файл
	{
		if err := os.Remove(initFile); err != nil {
			cErr = error_list.SystemCleanupError()
			cErr.SetError(err)

			usecase.components.Logger.Error().
				Format("The system could not be cleaned: '%s'. ", err).Write()
			return
		}
	}

	usecase.components.Logger.Info().
		Text("System cleanup is complete. ")

	return
}
