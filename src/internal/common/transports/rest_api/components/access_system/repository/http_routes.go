package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	db_models2 "sm-box/internal/common/objects/db_models"
	"sm-box/internal/common/objects/entities"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/sqlite3"
	"strings"
	"time"
)

// httpRoutesRepository - часть репозитория с управлением http маршрутов.
type httpRoutesRepository struct {
	connector  sqlite3.Connector
	components *components

	conf *Config
	ctx  context.Context
}

// GetRoute - получение http маршрута.
func (repo *httpRoutesRepository) GetRoute(ctx context.Context, method, path string) (route *entities.HttpRoute, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, method, path)
		defer func() { trc.Error(err).FunctionCallFinished(route) }()
	}

	// Подготовка
	{
		route = new(entities.HttpRoute).FillEmptyFields()
		method = strings.ToUpper(method)
	}

	// Основные данные
	{
		var model = new(db_models2.HttpRoute)

		var query = `
			select
				routes.id,
				routes.register_time,
				routes.active,
				routes.authorize
			from
				transports_http_routes as routes
			where
				routes.method = $1 and
				routes.path = $2
		`

		var row = repo.connector.QueryRowxContext(ctx, query, method, path)

		if err = row.Err(); err != nil {
			repo.components.Logger.Error().
				Format("Error when retrieving an item from the database: '%s'. ", err).Write()
			return
		}

		if err = row.StructScan(model); err != nil {
			repo.components.Logger.Error().
				Format("Error while reading item data from the database:: '%s'. ", err).Write()
			return
		}

		route.ID = model.ID
		route.Active = model.Active == "on"
		route.Authorize = model.Authorize == "on"
		route.Method = method
		route.Path = path

		if route.RegisterTime, err = time.Parse(time.RFC3339Nano, model.RegisterTime); err != nil {
			repo.components.Logger.Error().
				Format("Error while reading item data from the database:: '%s'. ", err).Write()
			return
		}
	}

	// Доступы
	{
		type Model struct {
			*db_models2.Role
			*db_models2.RoleInheritance
		}

		var (
			rows  *sqlx.Rows
			query = `
				WITH RECURSIVE cte_roles (id, project_id, title, parent) AS (
					select
						roles.id,
						roles.project_id,
						roles.title,
						0 as parent
					from
						system_access_roles as roles
					where
						roles.id in (
							select
								route_accesses.role_id as id
							from
								transports_http_routes as routes
									left join transports_http_route_accesses as route_accesses on routes.id = route_accesses.route_id
							where
								routes.method = $1 and
								routes.path = $2
						)
				
					UNION ALL
				
					select
						roles.id,
						roles.project_id,
						roles.title,
						role_inheritance.parent as parent
					from
						system_access_roles as roles
							left join system_access_role_inheritance role_inheritance on (role_inheritance.heir = roles.id)
							JOIN cte_roles cte ON cte.id = role_inheritance.parent
				)
				
				select
					distinct id,
							 coalesce(project_id, 0) as project_id,
							 title,
							 coalesce(parent, 0) as parent
				from
					cte_roles;
			`
			models = make([]*Model, 0, 10)
		)

		if rows, err = repo.connector.QueryxContext(ctx, query, method, path); err != nil {
			repo.components.Logger.Error().
				Format("Error when retrieving an items from the database: '%s'. ", err).Write()
			return
		}

		for rows.Next() {
			var model = new(Model)

			if err = rows.StructScan(model); err != nil {
				repo.components.Logger.Error().
					Format("Error while reading item data from the database:: '%s'. ", err).Write()
				return
			}

			models = append(models, model)
		}

		var writeInheritance func(parent *entities.HttpRouteAccess)

		writeInheritance = func(parent *entities.HttpRouteAccess) {
			for _, model := range models {
				if model.Parent == parent.ID {
					var (
						role = &entities.Role{
							ID:        model.ID,
							ProjectID: model.ProjectID,
							Title:     model.Title,

							Inheritances: make(entities.RoleInheritances, 0),
						}
					)
					role.FillEmptyFields()

					parent.Inheritances = append(parent.Inheritances, &entities.RoleInheritance{
						Role: role,
					})

					writeInheritance(&entities.HttpRouteAccess{
						Role: role,
					})
				}
			}
		}

		for _, model := range models {
			if model.Parent == 0 {
				var (
					role = &entities.Role{
						ID:        model.ID,
						ProjectID: model.ProjectID,
						Title:     model.Title,

						Inheritances: make(entities.RoleInheritances, 0),
					}
					acc = &entities.HttpRouteAccess{
						Role: role.FillEmptyFields(),
					}
				)

				writeInheritance(acc)
				route.Accesses = append(route.Accesses, acc)
			}
		}
	}

	return
}

// GetActiveRoute - получение активного http маршрута.
func (repo *httpRoutesRepository) GetActiveRoute(ctx context.Context, method, path string) (route *entities.HttpRoute, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, method, path)
		defer func() { trc.Error(err).FunctionCallFinished(route) }()
	}

	// Подготовка
	{
		route = new(entities.HttpRoute).FillEmptyFields()
		method = strings.ToUpper(method)
	}

	// Основные данные
	{
		var model = new(db_models2.HttpRoute)

		var query = `
			select
				routes.id,
				routes.register_time,
				routes.active,
				routes.authorize
			from
				transports_http_routes as routes
			where
				routes.method = $1 and
				routes.path = $2 and
				routes.active = 'on'
		`

		var row = repo.connector.QueryRowxContext(ctx, query, method, path)

		if err = row.Err(); err != nil {
			repo.components.Logger.Error().
				Format("Error when retrieving an item from the database: '%s'. ", err).Write()
			return
		}

		if err = row.StructScan(model); err != nil {
			repo.components.Logger.Error().
				Format("Error while reading item data from the database:: '%s'. ", err).Write()
			return
		}

		route.ID = model.ID
		route.Active = model.Active == "on"
		route.Authorize = model.Authorize == "on"
		route.Method = method
		route.Path = path

		if route.RegisterTime, err = time.Parse(time.RFC3339Nano, model.RegisterTime); err != nil {
			repo.components.Logger.Error().
				Format("Error while reading item data from the database:: '%s'. ", err).Write()
			return
		}
	}

	// Доступы
	{
		type Model struct {
			*db_models2.Role
			*db_models2.RoleInheritance
		}

		var (
			rows  *sqlx.Rows
			query = `
				WITH RECURSIVE cte_roles (id, project_id, title, parent) AS (
					select
						roles.id,
						roles.project_id,
						roles.title,
						0 as parent
					from
						system_access_roles as roles
					where
						roles.id in (
							select
								route_accesses.role_id as id
							from
								transports_http_routes as routes
									left join transports_http_route_accesses as route_accesses on routes.id = route_accesses.route_id
							where
								routes.method = $1 and
								routes.path = $2
						)
				
					UNION ALL
				
					select
						roles.id,
						roles.project_id,
						roles.title,
						role_inheritance.parent as parent
					from
						system_access_roles as roles
							left join system_access_role_inheritance role_inheritance on (role_inheritance.heir = roles.id)
							JOIN cte_roles cte ON cte.id = role_inheritance.parent
				)
				
				select
					distinct id,
							 coalesce(project_id, 0) as project_id,
							 title,
							 coalesce(parent, 0) as parent
				from
					cte_roles;
			`
			models = make([]*Model, 0, 10)
		)

		if rows, err = repo.connector.QueryxContext(ctx, query, method, path); err != nil {
			repo.components.Logger.Error().
				Format("Error when retrieving an items from the database: '%s'. ", err).Write()
			return
		}

		for rows.Next() {
			var model = new(Model)

			if err = rows.StructScan(model); err != nil {
				repo.components.Logger.Error().
					Format("Error while reading item data from the database:: '%s'. ", err).Write()
				return
			}

			models = append(models, model)
		}

		var writeInheritance func(parent *entities.HttpRouteAccess)

		writeInheritance = func(parent *entities.HttpRouteAccess) {
			for _, model := range models {
				if model.Parent == parent.ID {
					var (
						role = &entities.Role{
							ID:        model.ID,
							ProjectID: model.ProjectID,
							Title:     model.Title,

							Inheritances: make(entities.RoleInheritances, 0),
						}
					)
					role.FillEmptyFields()

					parent.Inheritances = append(parent.Inheritances, &entities.RoleInheritance{
						Role: role,
					})

					writeInheritance(&entities.HttpRouteAccess{
						Role: role,
					})
				}
			}
		}

		for _, model := range models {
			if model.Parent == 0 {
				var (
					role = &entities.Role{
						ID:        model.ID,
						ProjectID: model.ProjectID,
						Title:     model.Title,

						Inheritances: make(entities.RoleInheritances, 0),
					}
					acc = &entities.HttpRouteAccess{
						Role: role.FillEmptyFields(),
					}
				)

				writeInheritance(acc)
				route.Accesses = append(route.Accesses, acc)
			}
		}
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

	var (
		model = route.DbModel()
		query = `
			insert into 
				transports_http_routes (
						method, 
						path,
						register_time,
						active, 
						authorize
					) values (
						$1,
						$2,
						$3,
						$4,
						$5
					) 
					on conflict (method, path) do update 
						set active = 'on'
		`
	)

	if _, err = repo.connector.ExecContext(ctx, query,
		model.Method,
		model.Path,
		model.RegisterTime,
		model.Active,
		model.Authorize); err != nil {
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
		    active = 'off'
	`

	if _, err = repo.connector.ExecContext(ctx, query); err != nil {
		repo.components.Logger.Error().
			Format("Error updating an item from the database: '%s'. ", err).Write()
		return
	}

	return
}
