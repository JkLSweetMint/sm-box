package authentication_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	common_db_models "sm-box/internal/common/objects/db_models"
	common_entities "sm-box/internal/common/objects/entities"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
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

// GetToken - получение jwt токена.
func (repo *Repository) GetToken(ctx context.Context, data string) (tok *common_entities.JwtToken, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, data)
		defer func() { trc.Error(err).FunctionCallFinished(tok) }()
	}

	var model = new(common_db_models.JwtToken)

	// Получение данных
	{
		var query = `
			select
				tokens.id,
				coalesce(tokens.user_id, 0) as user_id,
				coalesce(tokens.project_id, 0) as project_id,
				tokens.language,
				tokens.issued_at,
				tokens.not_before,
				tokens.expires_at
			from
				access_system.jwt_tokens as tokens
			where
				tokens.data = $1
		`

		var row = repo.connector.QueryRowxContext(ctx, query, data)

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
		tok = new(common_entities.JwtToken)
		tok.FillEmptyFields()

		tok.ID = model.ID
		tok.UserID = model.UserID
		tok.ProjectID = model.ProjectID

		tok.Language = model.Language
		tok.Data = data

		tok.IssuedAt = model.IssuedAt
		tok.NotBefore = model.NotBefore
		tok.ExpiresAt = model.ExpiresAt
	}

	return
}

// GetTokenByID - получение jwt токена по ID.
func (repo *Repository) GetTokenByID(ctx context.Context, id types.ID) (tok *common_entities.JwtToken, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(err).FunctionCallFinished(tok) }()
	}

	var model = new(common_db_models.JwtToken)

	// Получение данных
	{
		var query = `
			select
				tokens.id,
				coalesce(tokens.user_id, 0) as user_id,
				coalesce(tokens.project_id, 0) as project_id,
				tokens.language,
				tokens.data,
				tokens.issued_at,
				tokens.not_before,
				tokens.expires_at
			from
				access_system.jwt_tokens as tokens
			where
				tokens.id = $1
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
		tok = new(common_entities.JwtToken)
		tok.FillEmptyFields()

		tok.ID = model.ID
		tok.UserID = model.UserID
		tok.ProjectID = model.ProjectID

		tok.Language = model.Language
		tok.Data = model.Data

		tok.IssuedAt = model.IssuedAt
		tok.NotBefore = model.NotBefore
		tok.ExpiresAt = model.ExpiresAt
	}

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
			access_system.jwt_tokens
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

// SetTokenProject - установить проект для токена.
func (repo *Repository) SetTokenProject(ctx context.Context, tokenID, projectID types.ID) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, tokenID, projectID)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var query = `
		update 
			access_system.jwt_tokens
		set
		    project_id = $1
		where
		    id = $2
	`

	if _, err = repo.connector.ExecContext(ctx, query, projectID, tokenID); err != nil {
		repo.components.Logger.Error().
			Format("Error updating an item from the database: '%s'. ", err).Write()
		return
	}

	return
}

// BasicAuth - базовая авторизация пользователя в системе.
// Для авторизации используется имя пользователя и пароль.
func (repo *Repository) BasicAuth(ctx context.Context, username, password string) (us *common_entities.User, err error) {
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

		// Получение
		{
			var query = `
			select
				users.id,
				coalesce(users.project_id, 0) as project_id,
				coalesce(users.email, '') as email,
				users.username
			from
				users.users as users
			where
				users.username = $1 and 
				users.password = $2
		`

			var row = repo.connector.QueryRowxContext(ctx, query, username, password)

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
			us.ProjectID = model.ProjectID
			us.Email = model.Email
			us.Username = model.Username
		}
	}

	// Доступы
	{
		type Model struct {
			*common_db_models.Role
			*common_db_models.RoleInheritance
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
	}

	return
}
