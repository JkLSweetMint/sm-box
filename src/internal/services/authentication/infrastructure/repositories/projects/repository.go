package projects_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"sm-box/internal/common/objects/db_models"
	"sm-box/internal/common/objects/entities"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/postgresql"
)

const (
	loggerInitiator = "infrastructure-[repositories]=projects"
)

// Repository - репозиторий проектов пользователей.
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

// GetListByUser - получение списка проектов пользователя.
func (repo *Repository) GetListByUser(ctx context.Context, userID types.ID) (list entities.ProjectList, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, userID)
		defer func() { trc.Error(err).FunctionCallFinished(list) }()
	}

	var (
		rows  *sqlx.Rows
		query = `
				select
					projects.id,
					projects.name,
					projects.version
				from
					projects
				where
					projects.owner_id = $1 or
					projects.id = any(
						select
							distinct coalesce(project_id, 0) as project_id
						from
							system_access.get_user_access($1) as (id bigint, project_id bigint, name varchar, parent bigint)
						where
							project_id != 0
					)
				order by projects.name;
			`
	)

	if rows, err = repo.connector.QueryxContext(ctx, query, userID); err != nil {
		repo.components.Logger.Error().
			Format("Error when retrieving an items from the database: '%s'. ", err).Write()
		return
	}

	list = make(entities.ProjectList, 0)

	for rows.Next() {
		var model = new(db_models.ProjectListItem)

		if err = rows.StructScan(model); err != nil {
			repo.components.Logger.Error().
				Format("Error while reading item data from the database:: '%s'. ", err).Write()
			return
		}

		list = append(list, &entities.ProjectListItem{
			ID:      model.ID,
			Name:    model.Name,
			Version: model.Version,
		})
	}

	return
}

// GetByID - получение проекта по ID.
func (repo *Repository) GetByID(ctx context.Context, id types.ID) (project *entities.Project, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(err).FunctionCallFinished(project) }()
	}

	// Основные данные
	{
		var model = new(db_models.Project)

		// Получение
		{
			var query = `
			select
					projects.id,
					projects.owner_id,
					projects.name,
					projects.description,
					projects.version
				from
					projects
				where
					projects.id = $1
				order by projects.name;
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

			project.Name = model.Name
			project.Description = model.Description
			project.Version = model.Version

			project.Owner.ID = model.ID
		}
	}

	// Владелец
	{
		var model = new(db_models.User)

		// Получение
		{
			var query = `
			select
				coalesce(users.project_id, 0) as project_id,
				coalesce(users.email, '') as email,
				users.username
			from
				users.users as users
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
			project.Owner.ProjectID = model.ProjectID
			project.Owner.Email = model.Email
			project.Owner.Username = model.Username
		}
	}

	return
}

// CheckAccess - проверка доступа пользователя к проекту.
func (repo *Repository) CheckAccess(ctx context.Context, userID, projectID types.ID) (exist bool, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, userID, projectID)
		defer func() { trc.Error(err).FunctionCallFinished(exist) }()
	}

	var query = `
			select
				$2 = any(
						select
							distinct coalesce(project_id, 0) as project_id
						from
							system_access.get_user_access($1) as (id bigint, project_id bigint, name varchar, parent bigint)
						where
							project_id != 0
					) as exist
		`

	var row = repo.connector.QueryRowxContext(ctx, query, userID, projectID)

	if err = row.Err(); err != nil {
		repo.components.Logger.Error().
			Format("Error when retrieving an item from the database: '%s'. ", err).Write()
		return
	}

	if err = row.Scan(&exist); err != nil {
		repo.components.Logger.Error().
			Format("Error while reading item data from the database:: '%s'. ", err).Write()
		return
	}

	return
}
