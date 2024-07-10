package projects_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"sm-box/internal/app/objects/db_models"
	"sm-box/internal/app/objects/entities"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/postgresql"
)

const (
	loggerInitiator = "infrastructure-[repositories]=projects"
)

// Repository - репозиторий проектов системы.
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
		Format("A '%s' repository has been created. ", "projects").
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
							access_system.get_user_access($1) as (id bigint, project_id bigint, name varchar, parent bigint)
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
		var model = new(db_models.Project)

		if err = rows.StructScan(model); err != nil {
			repo.components.Logger.Error().
				Format("Error while reading item data from the database:: '%s'. ", err).Write()
			return
		}

		list = append(list, &entities.Project{
			ID:      model.ID,
			OwnerID: model.OwnerID,

			Name:        model.Name,
			Description: model.Description,
			Version:     model.Version,
		})
	}

	return
}
