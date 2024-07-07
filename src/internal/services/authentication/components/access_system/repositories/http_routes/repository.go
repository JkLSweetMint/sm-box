package http_routes_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"sm-box/internal/services/authentication/objects/db_models"
	"sm-box/internal/services/authentication/objects/entities"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/postgresql"
	"strings"
)

const (
	loggerInitiator = "system-[components]-[access_system]-[repositories]=http_routes"
)

// Repository - репозиторий управления http маршрутами.
type Repository struct {
	connector  postgresql.Connector
	components *components

	conf *Config
	ctx  context.Context
}

// components - компоненты репозитория.
type components struct {
	Logger logger.Logger
}

// New - создание репозитория.
func New(ctx context.Context, conf *Config) (repo *Repository, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelRepository)

		trc.FunctionCall(ctx, conf)
		defer func() { trc.Error(err).FunctionCallFinished(repo) }()
	}

	repo = &Repository{
		ctx:  ctx,
		conf: conf,
	}

	// Компоненты
	{
		repo.components = new(components)

		// Logger
		{
			if repo.components.Logger, err = logger.New(loggerInitiator); err != nil {
				return
			}
		}
	}

	// Коннектор базы данных
	{
		if repo.connector, err = postgresql.New(ctx, repo.conf.Connector); err != nil {
			return
		}
	}

	repo.components.Logger.Info().
		Format("A '%s' repository has been created. ", "http_routes").
		Field("config", repo.conf).Write()

	return
}

// Get - получение http маршрута.
func (repo *Repository) Get(ctx context.Context, method, path string) (route *entities.HttpRoute, err error) {
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

		// Получение
		{
			var query = `
			select
				routes.id,
				routes.name,
				routes.description,
				routes.active,
				routes.authorize
			from
				transports.http_routes as routes
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
		}

		// Преобразование в сущность
		{
			route.ID = model.ID
			route.Name = model.Name
			route.Description = model.Description

			route.Method = method
			route.Path = path

			route.Active = model.Active
			route.Authorize = model.Authorize
		}
	}

	// Доступы
	{
		var models = make([]entities.HttpRouteAccess, 0, 10)

		// Получение
		{
			var (
				rows  *sqlx.Rows
				query = `
				select
					accesses.role_id
				from
				    transports.http_route_accesses as accesses
				where
				    accesses.route_id = $1
			`
			)

			if rows, err = repo.connector.QueryxContext(ctx, query, route.ID); err != nil {
				repo.components.Logger.Error().
					Format("Error when retrieving an items from the database: '%s'. ", err).Write()
				return
			}

			for rows.Next() {
				var model entities.HttpRouteAccess

				if err = rows.Scan(&model); err != nil {
					repo.components.Logger.Error().
						Format("Error while reading item data from the database:: '%s'. ", err).Write()
					return
				}

				models = append(models, model)
			}
		}

		// Преобразование в сущность
		{
			for _, model := range models {
				route.Accesses = append(route.Accesses, model)
			}
		}
	}

	return
}

// GetActive - получение активного http маршрута.
func (repo *Repository) GetActive(ctx context.Context, method, path string) (route *entities.HttpRoute, err error) {
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

		// Получение
		{
			var query = `
				select
					routes.id,
					routes.name,
					routes.description,
					routes.authorize
				from
					transports.http_routes as routes
				where
					routes.method = $1 and
					routes.path = $2 and
					routes.active
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
		}

		// Преобразование в сущность
		{
			route.ID = model.ID
			route.Name = model.Name
			route.Description = model.Description

			route.Method = method
			route.Path = path

			route.Active = true
			route.Authorize = model.Authorize
		}
	}

	// Доступы
	{
		var models = make([]entities.HttpRouteAccess, 0, 10)

		// Получение
		{
			var (
				rows  *sqlx.Rows
				query = `
				select
					accesses.role_id
				from
				    transports.http_route_accesses as accesses
				where
				    accesses.route_id = $1
			`
			)

			if rows, err = repo.connector.QueryxContext(ctx, query, route.ID); err != nil {
				repo.components.Logger.Error().
					Format("Error when retrieving an items from the database: '%s'. ", err).Write()
				return
			}

			for rows.Next() {
				var model entities.HttpRouteAccess

				if err = rows.Scan(&model); err != nil {
					repo.components.Logger.Error().
						Format("Error while reading item data from the database:: '%s'. ", err).Write()
					return
				}

				models = append(models, model)
			}
		}

		// Преобразование в сущность
		{
			for _, model := range models {
				route.Accesses = append(route.Accesses, model)
			}
		}
	}

	return
}
