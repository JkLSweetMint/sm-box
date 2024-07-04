package postgresql

import (
	"github.com/jmoiron/sqlx"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/universal_sqlx"
	"strings"
	"time"
)

const (
	driverName      = "postgres"
	loggerInitiator = "pkg-[databases]-[connector]=postgresql"
)

// connector - коннектор для postgresql базы данных.
type connector struct {
	*universal_sqlx.UniversalConnector
	conf *Config
}

// connect - подключение к базе данных.
func (conn *connector) connect() (err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall()
		defer func() { trace.Error(err).FunctionCallFinished() }()
	}

	conn.Components.Logger.Info().
		Format("Connecting to the database '%s' '%s'... ", driverName, conn.conf.ConnectionString()).Write()

	if conn.DB, err = sqlx.Open(driverName, conn.conf.ConnectionString()); err != nil {
		conn.Components.Logger.Error().
			Format("An error occurred while connecting to the database: '%s'. ", err).Write()
		return
	}

	for attempt := 1; ; attempt++ {
		if err = conn.Ping(); err != nil {
			conn.Components.Logger.Error().
				Format("An error occurred while connecting to the database: '%s'. ", err).
				Field("database", conn.conf.ConnectionString()).Write()

			if (!strings.HasSuffix(err.Error(), "connect: connection refused") &&
				!strings.HasSuffix(err.Error(), "pq: the database system is starting up")) ||
				attempt == 5 {
				return
			}

			conn.Components.Logger.Error().
				Text("Trying to connect to the database again...").
				Field("database", conn.conf.ConnectionString()).Write()

			time.Sleep(2 * time.Second)
			continue
		}

		break
	}

	conn.Components.Logger.Info().
		Format("Connecting to the database of the given '%s' '%s' installed. ", driverName, conn.conf.ConnectionString()).Write()

	return
}
