package projects

import (
	"context"
	"fmt"
	g_uuid "github.com/google/uuid"
	"path"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	"sm-box/pkg/databases/connectors/sqlite3"
)

// connectorByProject - получение коннектора для проекта.
func (repo *Repository) connectorByProject(ctx context.Context, uuid g_uuid.UUID) (connector sqlite3.Connector, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepositoryInternal)

		trc.FunctionCall(ctx, uuid)
		defer func() { trc.Error(err).FunctionCallFinished(connector) }()
	}

	var conf = new(sqlite3.Config).Default()

	conf.Database = path.Join(env.Paths.Var.Lib.Projects, fmt.Sprintf("%s.db", uuid.String()))

	if connector, err = sqlite3.New(ctx, conf); err != nil {
		return
	}

	return
}
