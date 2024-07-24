package projects_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"sm-box/internal/app/objects/db_models"
	"sm-box/internal/app/objects/entities"
	common_types "sm-box/internal/common/types"
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

// Get - получение проектов по ID.
func (repo *Repository) Get(ctx context.Context, ids []common_types.ID) (list entities.ProjectList, err error) {
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
				projects.id,
				projects.name,
				projects.version,
				projects.description
			from
				public.projects as projects
			where
				projects.id = any($1)
		order by projects.id
			`
	)

	var ids_ = make(pq.Int64Array, 0, len(ids))

	// Подготовка данных
	{
		for _, id := range ids {
			ids_ = append(ids_, int64(id))
		}
	}

	// Выполнение запроса
	{
		if rows, err = repo.connector.QueryxContext(ctx, query, ids_); err != nil {
			repo.components.Logger.Error().
				Format("Error when retrieving an items from the database: '%s'. ", err).Write()
			return
		}
	}

	// Чтение данных
	{
		list = make(entities.ProjectList, 0)

		for rows.Next() {
			var model = new(db_models.Project)

			if err = rows.StructScan(model); err != nil {
				repo.components.Logger.Error().
					Format("Error while reading item data from the database:: '%s'. ", err).Write()
				return
			}

			list = append(list, &entities.Project{
				ID: model.ID,

				Name:        model.Name,
				Description: model.Description,
				Version:     model.Version,
			})
		}
	}

	return
}

// GetOne - получение проекта по ID.
func (repo *Repository) GetOne(ctx context.Context, id common_types.ID) (project *entities.Project, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(err).FunctionCallFinished(project) }()
	}

	var model = new(db_models.Project)

	// Получение
	{
		var query = `
				select
					projects.id,
					projects.name,
					projects.version,
					projects.description
				from
					public.projects as projects
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

		project.Name = model.Name
		project.Version = model.Version
		project.Description = model.Description
	}

	return
}
