package repository

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"database/sql"
	"github.com/jmoiron/sqlx"
	common_db_models "sm-box/internal/common/objects/db_models"
	"sm-box/internal/common/objects/entities"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	"sm-box/pkg/databases/connectors/postgresql"
)

// usersRepository - часть репозитория с управлением пользователями.
type usersRepository struct {
	connector  postgresql.Connector
	components *components

	conf *Config
	ctx  context.Context
}

// GetUser - получение пользователя по идентификатору.
func (repo *usersRepository) GetUser(ctx context.Context, id types.ID) (us *common_entities.User, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(err).FunctionCallFinished(us) }()
	}

	// Подготовка
	{
		us = new(common_entities.User).FillEmptyFields()
	}

	// Основные данные
	{
		var model = new(common_db_models.User)

		var query = `
			select
				users.id,
				coalesce(users.project_id, 0) as project_id,
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

		us.ID = model.ID
		us.ProjectID = model.ProjectID
		us.Email = model.Email
		us.Username = model.Username
	}

	// Доступы
	{
		type Model struct {
			*common_db_models.Role
			*common_db_models.RoleInheritance
		}

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

		var writeInheritance func(parent *common_entities.UserAccess)

		writeInheritance = func(parent *common_entities.UserAccess) {
			for _, model := range models {
				if model.Parent == parent.ID {
					var (
						role = &common_entities.Role{
							ID:        model.ID,
							ProjectID: model.ProjectID,
							Name:      model.Name,

							Inheritances: make(common_entities.RoleInheritances, 0),
						}
					)
					role.FillEmptyFields()

					parent.Inheritances = append(parent.Inheritances, &common_entities.RoleInheritance{
						Role: role,
					})

					writeInheritance(&common_entities.UserAccess{
						Role: role,
					})
				}
			}
		}

		for _, model := range models {
			if model.Parent == 0 {
				var (
					role = &common_entities.Role{
						ID:        model.ID,
						ProjectID: model.ProjectID,
						Name:      model.Name,

						Inheritances: make(common_entities.RoleInheritances, 0),
					}
					acc = &common_entities.UserAccess{
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

// BasicAuth - базовая авторизация пользователя в системе.
// Для авторизации используется имя пользователя и пароль.
func (repo *usersRepository) BasicAuth(ctx context.Context, username, password string) (us *common_entities.User, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, username, password)
		defer func() { trc.Error(err).FunctionCallFinished(us) }()
	}

	// Подготовка
	{
		us = new(common_entities.User).FillEmptyFields()
	}

	// Основные данные
	{
		var model = new(common_db_models.User)

		var query = `
			select
				users.id,
				coalesce(users.project_id, 0) as project_id,
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
			*common_db_models.Role
			*common_db_models.RoleInheritance
		}

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

		var writeInheritance func(parent *common_entities.UserAccess)

		writeInheritance = func(parent *common_entities.UserAccess) {
			for _, model := range models {
				if model.Parent == parent.ID {
					var (
						role = &common_entities.Role{
							ID:        model.ID,
							ProjectID: model.ProjectID,
							Name:      model.Name,

							Inheritances: make(common_entities.RoleInheritances, 0),
						}
					)
					role.FillEmptyFields()

					parent.Inheritances = append(parent.Inheritances, &common_entities.RoleInheritance{
						Role: role,
					})

					writeInheritance(&common_entities.UserAccess{
						Role: role,
					})
				}
			}
		}

		for _, model := range models {
			if model.Parent == 0 {
				var (
					role = &common_entities.Role{
						ID:        model.ID,
						ProjectID: model.ProjectID,
						Name:      model.Name,

						Inheritances: make(common_entities.RoleInheritances, 0),
					}
					acc = &common_entities.UserAccess{
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
