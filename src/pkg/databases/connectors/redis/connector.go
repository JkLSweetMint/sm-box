package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
)

const (
	driverName      = "redis"
	loggerInitiator = "pkg-[databases]-[connector]=redis"
)

// connector - коннектор для redis базы данных.
type connector struct {
	*redis.Client

	conf *Config

	components  *components
	concurrency *concurrency
}

// concurrency - управление конкурентностью.
type concurrency struct {
	Ctx     context.Context
	GlobCtx context.Context
	Cancel  context.CancelFunc
}

// components - компоненты коннектора.
type components struct {
	Logger logger.Logger
}

// connect - подключение к базе данных.
func (conn *connector) connect() (err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall()
		defer func() { trace.Error(err).FunctionCallFinished() }()
	}

	conn.components.Logger.Info().
		Format("Connecting to the database '%s' '%s'... ", driverName, conn.conf.ConnectionString()).Write()

	var opt *redis.Options

	if opt, err = redis.ParseURL(conn.conf.ConnectionString()); err != nil {
		conn.components.Logger.Error().
			Format("The database connection string could not be parsed: '%s'. ", err).Write()
		return
	}

	conn.Client = redis.NewClient(opt)

	conn.components.Logger.Info().
		Format("Connecting to the database of the given '%s' '%s' installed. ", driverName, conn.conf.ConnectionString()).Write()

	return
}

// Close - закрывает базу данных и предотвращает запуск новых запросов,
// затем ожидает завершения всех запросов, которые начали обрабатываться на сервере.
//
// Закрытие базы данных происходит редко, так как дескриптор базы данных должен быть
// долговечным и совместно использоваться многими подпрограммами.
func (conn *connector) Close() (err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall()
		defer func() { trace.Error(err).FunctionCallFinished() }()
	}

	conn.concurrency.Cancel()

	if err = conn.Client.Close(); err != nil {
		conn.components.Logger.Error().
			Format("Error closing the database connection: '%s'. ", err).Write()
		return
	}

	return
}
