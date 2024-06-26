package authentication_repository

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"sm-box/internal/common/objects/db_models"
	"sm-box/internal/common/objects/entities"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	"sm-box/pkg/databases/connectors/postgresql"
)

const (
	loggerInitiator = "infrastructure-[repositories]=authentication"
)

// Repository - репозиторий аутентификации пользователей.
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
		Format("A '%s' repository has been created. ", "authentication").
		Field("config", repo.conf).Write()

	return
}

// SetTokenOwner - установить владельца токена.
func (repo *Repository) SetTokenOwner(ctx context.Context, tokenID, ownerID types.ID) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, tokenID, ownerID)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var query = `
		update 
			system_access.jwt_tokens
		set
		    user_id = $1
		where
		    id = $2
	`

	if _, err = repo.connector.ExecContext(ctx, query, ownerID, tokenID); err != nil {
		repo.components.Logger.Error().
			Format("Error updating an item from the database: '%s'. ", err).Write()
		return
	}

	return
}

// BasicAuth - базовая авторизация пользователя в системе.
// Для авторизации используется имя пользователя и пароль.
func (repo *Repository) BasicAuth(ctx context.Context, username, password string) (us *entities.User, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, username, password)
		defer func() { trc.Error(err).FunctionCallFinished(us) }()
	}

	// Подготовка
	{
		us = new(entities.User).FillEmptyFields()
	}

	// Основные данные
	{
		var model = new(db_models.User)

		var query = `
			select
				users.id,
				coalesce(users.project_id, 0) as project_id,
				coalesce(users.email, '') as email,
				users.username,
				users.password
			from
				public.users as users
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

		us.ID = model.ID
		us.ProjectID = model.ProjectID
		us.Email = model.Email
		us.Username = model.Username

		// Проверка пароля
		{
			var passwordDB []byte

			if passwordDB, err = rsa.DecryptOAEP(
				sha256.New(),
				rand.Reader,
				env.Vars.EncryptionKeys.Private,
				model.Password,
				[]byte("password")); err != nil {
				repo.components.Logger.Error().
					Format("The decryption of the user's password failed: '%s'. ", err).Write()
				return
			}

			if string(passwordDB) != password {
				repo.components.Logger.Error().
					Text("The user's password does not match. ").Write()

				err = sql.ErrNoRows
				return
			}
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
								user_accesses.role_id as id
							from
								users
									left join user_accesses on users.id = user_accesses.user_id
				
							where
								users.id = $1
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

	return
}
