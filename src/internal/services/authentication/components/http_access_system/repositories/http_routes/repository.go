package http_routes_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"sm-box/internal/services/authentication/objects/db_models"
	"sm-box/internal/services/authentication/objects/entities"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/postgresql"
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

// GetAll - получение всех http маршрутов.
func (repo *Repository) GetAll(ctx context.Context) (list []*entities.HttpRoute, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished(list) }()
	}

	list = make([]*entities.HttpRoute, 0)

	// Получение данных
	{
		var (
			rows  *sqlx.Rows
			query = `
				select
					routes.id,
					routes.system_name,
					routes.name,
					routes.description,
					routes.protocols,
					routes.method,
					coalesce(routes.path, '') as path,
					coalesce(routes.regexp_path, '') as regexp_path,
					routes.authorize,
					routes.active
				from
					transports.http_routes as routes
			`
		)

		if rows, err = repo.connector.QueryxContext(ctx, query); err != nil {
			repo.components.Logger.Error().
				Format("Error when retrieving an items from the database: '%s'. ", err).Write()
			return
		}

		for rows.Next() {
			var model = new(db_models.HttpRoute)

			if err = rows.StructScan(model); err != nil {
				repo.components.Logger.Error().
					Format("Error while reading item data from the database:: '%s'. ", err).Write()
				return
			}

			var route = new(entities.HttpRoute)

			// Преобразование в сущность
			{
				route.ID = model.ID

				route.SystemName = model.SystemName
				route.Name = model.Name
				route.Description = model.Description

				route.Protocols = model.Protocols
				route.Method = model.Method
				route.Path = model.Path
				route.RegexpPath = model.RegexpPath

				route.Active = model.Active
				route.Authorize = model.Authorize
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

			list = append(list, route)
		}
	}

	return
}
