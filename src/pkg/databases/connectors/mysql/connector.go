package mysql

import (
	"github.com/jmoiron/sqlx"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/universal_sqlx"
)

const (
	driverName      = "mysql"
	loggerInitiator = "pkg-[databases]-[connector]=mysql"
)

// connector - коннектор для mysql базы данных.
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

	conn.Components.Logger.Info().
		Format("Connecting to the database of the given '%s' '%s' installed. ", driverName, conn.conf.ConnectionString()).Write()

	return
}
