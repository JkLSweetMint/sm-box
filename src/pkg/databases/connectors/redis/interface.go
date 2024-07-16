package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
)

// Connector - описание коннектора для redis базы данных.
type Connector interface {
	redis.UniversalClient
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
		concurrency: &concurrency{
			GlobCtx: ctx,
		},

		conf: conf,
	}

	// Компоненты
	{
		ref.components = new(components)

		// Logger
		{
			if ref.components.Logger, err = logger.New(loggerInitiator); err != nil {
				return
			}
		}
	}

	ref.concurrency.Ctx, ref.concurrency.Cancel = context.WithCancel(context.Background())

	if err = ref.connect(); err != nil {
		ref.components.Logger.Error().
			Format("Failed to connect to the database: '%s'. ", err).Write()
		return
	}

	env.Synchronization.WaitGroup.Add(1)

	go func(conn *connector) {
		defer env.Synchronization.WaitGroup.Done()

		select {
		case <-conn.concurrency.GlobCtx.Done():
			if err = conn.Close(); err != nil {
				conn.components.Logger.Error().
					Format("An error occurred while closing the database connection: '%s'. ", err).Write()
				return
			}
		case <-conn.concurrency.Ctx.Done():
			if err = conn.Close(); err != nil {
				conn.components.Logger.Error().
					Format("An error occurred while closing the database connection: '%s'. ", err).Write()
				return
			}
		}
	}(ref)

	conn = ref

	ref.components.Logger.Info().
		Format("The database connector '%s' has been created...  ", driverName).
		Field("config", ref.conf).Write()

	return
}
