package users_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"sm-box/internal/common/types"
	"sm-box/internal/services/users/objects/db_models"
	"sm-box/internal/services/users/objects/entities"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/postgresql"
)

const (
	loggerInitiator = "infrastructure-[repositories]=users"
)

// Repository - репозиторий для работы с пользователями.
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
func New(ctx context.Context) (repo *Repository, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelRepository)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished(repo) }()
	}

	repo = &Repository{
		ctx: ctx,
	}

	// Конфигурация
	{
		repo.conf = new(Config).Default()

		if err = repo.conf.Read(); err != nil {
			return
		}
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
		Format("A '%s' repository has been created. ", "users").
		Field("config", repo.conf).Write()

	return
}

// Get - получение пользователей по ID.
func (repo *Repository) Get(ctx context.Context, ids []types.ID) (list []*entities.User, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, ids)
		defer func() { trc.Error(err).FunctionCallFinished(list) }()
	}

	var (
		rows  *sqlx.Rows
		query = `
			select
				users.id,
				coalesce(users.email, '') as email,
				users.username
			from
				users.users as users
			where
				users.id = any($1)
			order by users.id
			`
	)

	var ids_ = make(pq.Int64Array, 0, len(ids))

	for _, id := range ids {
		ids_ = append(ids_, int64(id))
	}

	if rows, err = repo.connector.QueryxContext(ctx, query, ids_); err != nil {
		repo.components.Logger.Error().
			Format("Error when retrieving an items from the database: '%s'. ", err).Write()
		return
	}

	list = make([]*entities.User, 0)

	for rows.Next() {
		var model = new(db_models.User)

		if err = rows.StructScan(model); err != nil {
			repo.components.Logger.Error().
				Format("Error while reading item data from the database:: '%s'. ", err).Write()
			return
		}

		var us = &entities.User{
			ID: model.ID,

			Email:    model.Email,
			Username: model.Username,
		}

		// Доступы
		{
			type Model struct {
				*db_models.Role
				*db_models.RoleInheritance
			}

			var models = make([]*Model, 0, 10)

			// Получение
			{
				var (
					rows  *sqlx.Rows
					query = `
				select
					distinct id,
					coalesce(project_id, 0) as project_id,
					name,
					coalesce(parent, 0) as parent
				from
					access_system.get_user_access($1) as (id bigint, project_id bigint, name varchar, parent bigint);
			`
				)

				if rows, err = repo.connector.QueryxContext(ctx, query, us.ID); err != nil {
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
			}

			// Перенос в сущность
			{
				var writeInheritance func(parent *entities.UserAccess)

				writeInheritance = func(parent *entities.UserAccess) {
					for _, model := range models {
						if model.Parent == parent.ID {
							var (
								role = &entities.Role{
									ID:        model.ID,
									ProjectID: model.ProjectID,
									Name:      model.Name,

									Inheritances: make(entities.RoleInheritances, 0),
								}
							)
							role.FillEmptyFields()

							parent.Inheritances = append(parent.Inheritances, &entities.RoleInheritance{
								Role: role,
							})

							writeInheritance(&entities.UserAccess{
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
								Name:      model.Name,

								Inheritances: make(entities.RoleInheritances, 0),
							}
							acc = &entities.UserAccess{
								Role: role.FillEmptyFields(),
							}
						)

						writeInheritance(acc)
						us.Accesses = append(us.Accesses, acc)
					}
				}
			}
		}

		list = append(list, us)
	}

	return
}

// GetOne - получение пользователя по ID.
func (repo *Repository) GetOne(ctx context.Context, id types.ID) (us *entities.User, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(err).FunctionCallFinished(us) }()
	}

	// Подготовка
	{
		us = new(entities.User).FillEmptyFields()
	}

	// Основные данные
	{
		var model = new(db_models.User)

		// Получение
		{
			var query = `
			select
				users.id,
				coalesce(users.email, '') as email,
				users.username
			from
				users.users as users
			where
				users.id = $1
		`

			var row = repo.connector.QueryRowxContext(ctx, query, id)

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

		// Перенос в сущность
		{
			us.ID = model.ID
			us.Email = model.Email
			us.Username = model.Username
		}
	}

	// Доступы
	{
		type Model struct {
			*db_models.Role
			*db_models.RoleInheritance
		}

		var models = make([]*Model, 0, 10)

		// Получение
		{

			var (
				rows  *sqlx.Rows
				query = `
				select
					distinct id,
					coalesce(project_id, 0) as project_id,
					name,
					coalesce(parent, 0) as parent
				from
					access_system.get_user_access($1) as (id bigint, project_id bigint, name varchar, parent bigint);
			`
			)

			if rows, err = repo.connector.QueryxContext(ctx, query, us.ID); err != nil {
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
		}

		// Перенос в сущность
		{
			var writeInheritance func(parent *entities.UserAccess)

			writeInheritance = func(parent *entities.UserAccess) {
				for _, model := range models {
					if model.Parent == parent.ID {
						var (
							role = &entities.Role{
								ID:        model.ID,
								ProjectID: model.ProjectID,
								Name:      model.Name,

								Inheritances: make(entities.RoleInheritances, 0),
							}
						)
						role.FillEmptyFields()

						parent.Inheritances = append(parent.Inheritances, &entities.RoleInheritance{
							Role: role,
						})

						writeInheritance(&entities.UserAccess{
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
							Name:      model.Name,

							Inheritances: make(entities.RoleInheritances, 0),
						}
						acc = &entities.UserAccess{
							Role: role.FillEmptyFields(),
						}
					)

					writeInheritance(acc)
					us.Accesses = append(us.Accesses, acc)
				}
			}
		}
	}

	return
}
