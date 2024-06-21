package projects

import (
	"context"
	g_uuid "github.com/google/uuid"
	"sm-box/internal/common/db_models"
	"sm-box/internal/common/entities"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	"sm-box/pkg/databases/connectors/sqlite3"
)

// Repository - репозиторий для управления проектами.
type Repository struct {
	connector  sqlite3.Connector
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
		var trace = tracer.New(tracer.LevelMain)

		trace.FunctionCall(ctx)

		defer func() { trace.Error(err).FunctionCallFinished(repo) }()
	}

	repo = new(Repository)
	repo.ctx = ctx

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
			if repo.components.Logger, err = logger.New(env.Vars.SystemName); err != nil {
				return
			}
		}
	}

	// Коннектор базы данных
	{
		if repo.connector, err = sqlite3.New(ctx, repo.conf.Connector); err != nil {
			return
		}
	}

	repo.components.Logger.Info().
		Format("A '%s' repository has been created. ", "projects").
		Field("config", repo.conf).Write()

	return
}

// Create - создание проекта.
func (repo *Repository) Create(ctx context.Context, uuid, name, description, version string) (id types.ID, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, uuid, name, description, version)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var (
		query = `
			insert into 
				projects (
						uuid, 
						name, 
						version,
					    description
					) values (
						$1,
						$2,
						$3,
						$4
					)
			returning id;
		`
	)

	var row = repo.connector.QueryRowxContext(ctx, query,
		uuid,
		name,
		description,
		version)

	if err = row.Err(); err != nil {
		repo.components.Logger.Error().
			Format("Error when retrieving an item from the database: '%s'. ", err).Write()
		return
	}

	if err = row.Scan(&id); err != nil {
		repo.components.Logger.Error().
			Format("Error while reading item data from the database:: '%s'. ", err).Write()
		return
	}

	return
}

// RemoveByUUID - удаление проекта по UUID.
func (repo *Repository) RemoveByUUID(ctx context.Context, uuid string) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, uuid)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var query = `
		delete from
			projects
		where
		    uuid = $1
	`

	if _, err = repo.connector.ExecContext(ctx, query, uuid); err != nil {
		repo.components.Logger.Error().
			Format("Error when deleting an item from the database: '%s'. ", err).Write()
		return
	}

	return
}

// GetByID - получение данных проекта по ID.
func (repo *Repository) GetByID(ctx context.Context, id types.ID) (project *entities.Project, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	// Основные данные
	{
		var model = new(db_models.Project)

		// Получение данных
		{
			var query = `
			select
				projects.id,
				projects.uuid,
				projects.owner_id,
				projects.name,
				projects.description,
				projects.version
			from
				projects
			where
				projects.id = $1
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
			project = new(entities.Project)
			project.FillEmptyFields()

			project.ID = model.ID

			if project.UUID, err = g_uuid.Parse(model.UUID); err != nil {
				repo.components.Logger.Error().
					Format("Error while reading item data from the database:: '%s'. ", err).Write()
				return
			}

			project.Owner.ID = model.OwnerID

			project.Name = model.Name
			project.Description = model.Description
			project.Version = model.Version
		}
	}

	// Данные владельца
	{
		var model = new(db_models.User)

		// Получение данных
		{
			var query = `
			select
				users.id,
				coalesce(users.project_id, 0) as project_id,
				coalesce(users.email, '') as email,
				users.username
			from
				users
			where
				users.id = $1
		`

			var row = repo.connector.QueryRowxContext(ctx, query, project.Owner.ID)

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
			project.Owner.ID = model.ID
			project.Owner.ProjectID = model.ProjectID

			project.Owner.Email = model.Email
			project.Owner.Username = model.Username
		}
	}

	return
}

// GetByUUID - получение данных проекта по UUID.
func (repo *Repository) GetByUUID(ctx context.Context, uuid string) (project *entities.Project, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, uuid)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	// Основные данные
	{
		var model = new(db_models.Project)

		// Получение данных
		{
			var query = `
			select
				projects.id,
				projects.uuid,
				projects.owner_id,
				projects.name,
				projects.description,
				projects.version
			from
				projects
			where
				projects.uuid = $1
		`

			var row = repo.connector.QueryRowxContext(ctx, query, uuid)

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
			project = new(entities.Project)
			project.FillEmptyFields()

			project.ID = model.ID

			if project.UUID, err = g_uuid.Parse(model.UUID); err != nil {
				repo.components.Logger.Error().
					Format("Error while reading item data from the database:: '%s'. ", err).Write()
				return
			}

			project.Owner.ID = model.OwnerID

			project.Name = model.Name
			project.Description = model.Description
			project.Version = model.Version
		}
	}

	// Данные владельца
	{
		var model = new(db_models.User)

		// Получение данных
		{
			var query = `
			select
				users.id,
				coalesce(users.project_id, 0) as project_id,
				coalesce(users.email, '') as email,
				users.username
			from
				users
			where
				users.id = $1
		`

			var row = repo.connector.QueryRowxContext(ctx, query, project.Owner.ID)

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
			project.Owner.ID = model.ID
			project.Owner.ProjectID = model.ProjectID

			project.Owner.Email = model.Email
			project.Owner.Username = model.Username
		}
	}

	return
}
