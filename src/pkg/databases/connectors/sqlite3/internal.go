package sqlite3

import (
	"context"
	"github.com/jmoiron/sqlx"
	"path"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	"sm-box/pkg/databases/connectors/universal_sqlx"
)

const (
	driverName      = "sqlite3_with_go_func"
	loggerInitiator = "pkg-[databases]-[connector]=sqlite3"
)

// connector - коннектор для sqlite3 базы данных.
type connector struct {
	*universal_sqlx.UniversalConnector
	conf *Config
}

// New - создание нового коннектора.
func New(ctx context.Context, conf *Config) (conn Connector, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelMain, tracer.LevelDatabaseConnector)

		trace.FunctionCall(ctx, conf)
		defer func() { trace.Error(err).FunctionCallFinished(conn) }()
	}

	// Конфигурация
	{
		if err = conf.FillEmptyFields().Validate(); err != nil {
			return
		}
	}

	var ref = &connector{
		UniversalConnector: &universal_sqlx.UniversalConnector{
			Components: new(universal_sqlx.Components),
			Concurrency: &universal_sqlx.Concurrency{
				GlobCtx: ctx,
			},
		},
		conf: conf,
	}

	// Компоненты
	{
		// Logger
		{
			if ref.Components.Logger, err = logger.New(loggerInitiator); err != nil {
				return
			}
		}
	}

	ref.conf.Database = path.Join(env.Paths.SystemLocation, ref.conf.Database)

	if err = ref.connect(); err != nil {
		ref.Components.Logger.Error().
			Format("Failed to connect to the database: '%s'. ", err).Write()
		return
	}

	ref.DB.SetConnMaxLifetime(ref.conf.ConnMaxLifetime)
	ref.DB.SetMaxOpenConns(ref.conf.MaxOpenConns)
	ref.DB.SetMaxIdleConns(ref.conf.MaxIdleConns)

	ref.Concurrency.Ctx, ref.Concurrency.Cancel = context.WithCancel(context.Background())

	env.Synchronization.WaitGroup.Add(1)

	go func(conn *connector) {
		defer env.Synchronization.WaitGroup.Done()

		select {
		case <-conn.Concurrency.GlobCtx.Done():
			if err = conn.Close(); err != nil {
				conn.Components.Logger.Error().
					Format("An error occurred while closing the database connection: '%s'. ", err).Write()
				return
			}
		case <-conn.Concurrency.Ctx.Done():
			if err = conn.Close(); err != nil {
				conn.Components.Logger.Error().
					Format("An error occurred while closing the database connection: '%s'. ", err).Write()
				return
			}
		}
	}(ref)

	conn = ref

	ref.Components.Logger.Info().
		Format("The database connector '%s' has been created...  ", driverName).
		Field("config", ref.conf).Write()

	return
}

// connect - подключение к базе данных.
func (conn *connector) connect() (err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelInternal, tracer.LevelDatabaseConnector)

		trace.FunctionCall()
		defer func() { trace.Error(err).FunctionCallFinished() }()
	}

	conn.Components.Logger.Info().
		Format("Connecting to the database '%s' '%s'... ", driverName, conn.conf.ConnectionString()).Write()

	if conn.DB, err = sqlx.Open(driverName, conn.conf.ConnectionString()); err != nil {
		conn.Components.Logger.Error().
			Format("An error occurred while connecting to the database: '%s'. ", err)
		return
	}

	conn.Components.Logger.Info().
		Format("Connecting to the database of the given '%s' '%s' installed. ", driverName, conn.conf.ConnectionString()).Write()

	return
}
