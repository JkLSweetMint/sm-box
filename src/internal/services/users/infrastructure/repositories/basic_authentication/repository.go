package basic_authentication_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/users/objects/db_models"
	"sm-box/internal/services/users/objects/entities"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/postgresql"
)

const (
	loggerInitiator = "infrastructure-[repositories]=basic_authentication"
)

// Repository - репозиторий для аутентификации пользователей.
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
		Format("A '%s' repository has been created. ", "basic_authentication").
		Field("config", repo.conf).Write()

	return
}

// GetByUsername - получение пользователя по имени.
func (repo *Repository) GetByUsername(ctx context.Context, username string) (us *entities.User, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, username)
		defer func() { trc.Error(err).FunctionCallFinished(us) }()
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
				users.username,
				users.password
			from
				users.users as users
			where
				users.username = $1
		`

			var row = repo.connector.QueryRowxContext(ctx, query, username)

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
			us = new(entities.User).FillEmptyFields()

			us.ID = model.ID
			us.Email = model.Email
			us.Username = model.Username
			us.Password = model.Password
		}
	}

	// Доступы
	{
		type RoleModel struct {
			*db_models.Role
			*db_models.RoleInheritance
		}

		type PermissionModel struct {
			*db_models.Permission

			RoleID common_types.ID `db:"role_id"`
		}

		var (
			writeRoleInheritance func(parent *entities.Role, models []*RoleModel)
			writePermission      func(role *entities.Role, models []*PermissionModel)
		)

		// writeRoleInheritance
		{
			writeRoleInheritance = func(parent *entities.Role, models []*RoleModel) {
				for _, model := range models {
					if model.ParentID == parent.ID {
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

						writeRoleInheritance(role, models)
					}
				}
			}

			writePermission = func(role *entities.Role, models []*PermissionModel) {
				for _, model := range models {
					if model.RoleID == role.ID {
						role.Permissions = append(role.Permissions, &entities.Permission{
							ID:        model.ID,
							ProjectID: model.ProjectID,

							Name:            model.Name,
							NameI18n:        model.NameI18n,
							Description:     model.Description,
							DescriptionI18n: model.DescriptionI18n,

							IsSystem: model.IsSystem,
						})
					}
				}

				for _, child := range role.Inheritances {
					writePermission(child.Role, models)
				}
			}
		}

		// Роли
		{
			var (
				models = make([]*RoleModel, 0, 10)
				rows   *sqlx.Rows
				query  = `
					select
						distinct id,
								 coalesce(project_id, 0) as project_id,
								 coalesce(parent_id, 0) as parent_id,
								 coalesce(name, '') as name,
								 coalesce(name_i18n, '00000000-0000-0000-0000-000000000000') as name_i18n,
								 coalesce(description, '') as description,
								 coalesce(description_i18n, '00000000-0000-0000-0000-000000000000') as description_i18n,
								 is_system
					from
						access_system.get_user_roles($1) as (
							id bigint, 
							project_id bigint,
							parent_id bigint,
							name varchar,
							name_i18n uuid,
							description varchar,
							description_i18n uuid,
							is_system boolean);
								`
			)

			// Выполнение запроса
			{
				if rows, err = repo.connector.QueryxContext(ctx, query, us.ID); err != nil {
					repo.components.Logger.Error().
						Format("Error when retrieving an items from the database: '%s'. ", err).Write()
					return
				}
			}

			// Чтение данных
			{
				for rows.Next() {
					var model = new(RoleModel)

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
				for _, model := range models {
					if model.ParentID == 0 {
						var (
							role = &entities.Role{
								ID:        model.ID,
								ProjectID: model.ProjectID,
								Name:      model.Name,

								Inheritances: make(entities.RoleInheritances, 0),
							}
							acc = role.FillEmptyFields()
						)

						writeRoleInheritance(acc, models)
						us.Accesses.Roles = append(us.Accesses.Roles, acc)
					}
				}
			}
		}

		// Права
		{
			var (
				models = make([]*PermissionModel, 0, 10)
				rows   *sqlx.Rows
				query  = `
					select
						distinct id,
								 coalesce(project_id, 0) as project_id,
								 coalesce(role_id, 0) as role_id,
								 coalesce(name, '') as name,
								 coalesce(name_i18n, '00000000-0000-0000-0000-000000000000') as name_i18n,
								 coalesce(description, '') as description,
								 coalesce(description_i18n, '00000000-0000-0000-0000-000000000000') as description_i18n,
								 is_system
					from
						access_system.get_user_permissions($1) as (
							id bigint,
							project_id bigint,
							role_id bigint,
							name varchar,
							name_i18n uuid,
							description varchar,
							description_i18n uuid,
							is_system boolean);
					`
			)

			// Выполнение запроса
			{
				if rows, err = repo.connector.QueryxContext(ctx, query, us.ID); err != nil {
					repo.components.Logger.Error().
						Format("Error when retrieving an items from the database: '%s'. ", err).Write()
					return
				}
			}

			// Чтение данных
			{
				for rows.Next() {
					var model = new(PermissionModel)

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
				for _, model := range models {
					if model.RoleID == 0 {
						us.Accesses.Permissions = append(us.Accesses.Permissions, &entities.Permission{
							ID:        model.ID,
							ProjectID: model.ProjectID,

							Name:            model.Name,
							NameI18n:        model.NameI18n,
							Description:     model.Description,
							DescriptionI18n: model.DescriptionI18n,

							IsSystem: model.IsSystem,
						})

						continue
					}
				}

				for _, role := range us.Accesses.Roles {
					writePermission(role, models)
				}
			}
		}
	}

	return
}
