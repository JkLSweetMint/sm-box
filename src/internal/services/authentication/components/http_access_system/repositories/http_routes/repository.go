package http_routes_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	common_types "sm-box/internal/common/types"
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

	// Основные данные
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

		// Выполнение запроса
		{
			if rows, err = repo.connector.QueryxContext(ctx, query); err != nil {
				repo.components.Logger.Error().
					Format("Error when retrieving an items from the database: '%s'. ", err).Write()
				return
			}
		}

		// Чтение данных
		{
			list = make([]*entities.HttpRoute, 0)

			for rows.Next() {
				var model = new(db_models.HttpRoute)

				if err = rows.StructScan(model); err != nil {
					repo.components.Logger.Error().
						Format("Error while reading item data from the database:: '%s'. ", err).Write()
					return
				}

				var route = &entities.HttpRoute{
					ID: model.ID,

					SystemName:  model.SystemName,
					Name:        model.Name,
					Description: model.Description,

					Protocols:  model.Protocols,
					Method:     model.Method,
					Path:       model.Path,
					RegexpPath: model.RegexpPath,

					Active:    model.Active,
					Authorize: model.Authorize,
				}

				route.FillEmptyFields()

				list = append(list, route)
			}
		}
	}

	// Доступы
	{
		// Роли
		{
			type Model struct {
				RouteID common_types.ID `db:"route_id"`
				RoleID  common_types.ID `db:"role_id"`
			}

			var (
				rows  *sqlx.Rows
				query = `
						select
						    route_id,
							role_id
						from
							transports.http_route_accesses
						where
							role_id is not null
			`
			)

			// Выполнение запроса
			{
				if rows, err = repo.connector.QueryxContext(ctx, query); err != nil {
					repo.components.Logger.Error().
						Format("Error when retrieving an items from the database: '%s'. ", err).Write()
					return
				}
			}

			// Чтение данных
			{
				for rows.Next() {
					var model = new(Model)

					if err = rows.StructScan(&model); err != nil {
						repo.components.Logger.Error().
							Format("Error while reading item data from the database:: '%s'. ", err).Write()
						return
					}

					for _, route := range list {
						if route.ID == model.RouteID {
							route.Accesses.Roles = append(route.Accesses.Roles, model.RoleID)
							break
						}
					}
				}
			}
		}

		// Права
		{
			type Model struct {
				RouteID      common_types.ID `db:"route_id"`
				PermissionID common_types.ID `db:"permission_id"`
			}

			var (
				rows  *sqlx.Rows
				query = `
						select
						    route_id,
							permission_id
						from
							transports.http_route_accesses
						where
							permission_id is not null
			`
			)

			// Выполнение запроса
			{
				if rows, err = repo.connector.QueryxContext(ctx, query); err != nil {
					repo.components.Logger.Error().
						Format("Error when retrieving an items from the database: '%s'. ", err).Write()
					return
				}
			}

			// Чтение данных
			{
				for rows.Next() {
					var model = new(Model)

					if err = rows.StructScan(&model); err != nil {
						repo.components.Logger.Error().
							Format("Error while reading item data from the database:: '%s'. ", err).Write()
						return
					}

					for _, route := range list {
						if route.ID == model.RouteID {
							route.Accesses.Permissions = append(route.Accesses.Permissions, model.PermissionID)
							break
						}
					}
				}
			}
		}
	}

	return
}
