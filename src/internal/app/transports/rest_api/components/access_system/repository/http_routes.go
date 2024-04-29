package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"sm-box/internal/common/db_models"
	"sm-box/internal/common/entities"
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
		var model = new(db_models.HttpRoute)

		var query = `
			select
				routes.id,
				routes.active,
				routes.register_time
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
		route.Active = model.Active
		route.Method = method
		route.Path = path

		if route.RegisterTime, err = time.ParseInLocation("2006-01-02 15:04:05", model.RegisterTime, time.Local); err != nil {
			repo.components.Logger.Error().
				Format("Error while reading item data from the database:: '%s'. ", err).Write()
			return
		}
	}

	// Доступы
	{
		type Model struct {
			*db_models.Role
			*db_models.RoleInheritance
		}

		var (
			rows  *sqlx.Rows
			query = `
				WITH RECURSIVE cte_roles AS (
					select
						roles.id,
						roles.project_id,
						roles.title,
						role_inheritance.parent as parent,
						role_inheritance.heir as heir
					from
						system_access_roles as roles
							left join system_access_role_inheritance role_inheritance on (role_inheritance.heir = roles.id)
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
						role_inheritance.parent as parent,
						role_inheritance.heir as heir
					from
						system_access_roles as roles
							left join system_access_role_inheritance role_inheritance on (role_inheritance.heir = roles.id)
							join cte_roles cte on cte.parent = roles.id
				)
				
				select 
					distinct id, 
					coalesce(project_id, 0) as project_id,
					title,
					coalesce(parent, 0) as parent
				from 
					cte_roles 
				order by id;
			`
			list = make([]*Model, 0, 10)
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

			list = append(list, model)
		}

		var writeInheritance func(parent *entities.HttpRouteAccess)

		writeInheritance = func(parent *entities.HttpRouteAccess) {
			for _, model := range list {
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

		for _, model := range list {
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

	if _, err = repo.connector.ExecContext(ctx, query, route.ActiveInt(), strings.ToUpper(route.Method), route.Path); err != nil {
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

	if _, err = repo.connector.ExecContext(ctx, query); err != nil {
		repo.components.Logger.Error().
			Format("Error updating an item from the database: '%s'. ", err).Write()
		return
	}

	return
}
