package repository

import (
	"context"
	"sm-box/internal/common/entities"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/sqlite3"
	"strings"
)

// httpRoutesRepository - часть репозитория с управлением http маршрутов.
type httpRoutesRepository struct {
	connector  sqlite3.Connector
	components *components

	conf *Config
	ctx  context.Context
}

// GetRoute - получение http маршрута.
func (repo *httpRoutesRepository) GetRoute(ctx context.Context, method, path string) (info *entities.HttpRoute, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, method, path)
		defer func() { trc.Error(err).FunctionCallFinished(info) }()
	}

	return
}

// RegisterRoute - регистрация http маршрута.
func (repo *httpRoutesRepository) RegisterRoute(ctx context.Context, route *entities.HttpRoute) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, route)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var query = `
		insert into 
			transports_http_routes (
					active, 
					method, 
					path
				) values (
				  	$1,
				    $2,
				    $3
				) 
				  on conflict (method, path) do update 
				      set active = 1
	`

	if _, err = repo.connector.Exec(query, route.ActiveInt(), strings.ToUpper(route.Method), route.Path); err != nil {
		repo.components.Logger.Error().
			Format("Error inserting an item from the database: '%s'. ", err).Write()
		return
	}

	return
}

// SetInactiveRoutes - установить все http маршруты как не активные.
func (repo *httpRoutesRepository) SetInactiveRoutes(ctx context.Context) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var query = `
		update 
			transports_http_routes
		set
		    active = 0
	`

	if _, err = repo.connector.Exec(query); err != nil {
		repo.components.Logger.Error().
			Format("Error updating an item from the database: '%s'. ", err).Write()
		return
	}

	return
}
