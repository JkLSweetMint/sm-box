package languages_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"sm-box/internal/services/i18n/objects/db_models"
	"sm-box/internal/services/i18n/objects/entities"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/postgresql"
)

const (
	loggerInitiator = "infrastructure-[repositories]=languages"
)

// Repository - репозиторий языков локализации.
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
		Format("A '%s' repository has been created. ", "languages").
		Field("config", repo.conf).Write()

	return
}

// GetList - получение списка языков.
func (repo *Repository) GetList(ctx context.Context) (list []*entities.Language, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished(list) }()
	}

	var (
		rows  *sqlx.Rows
		query = `
				select
					languages.code,
					languages.name,
					languages.active
				from
					public.languages as languages
				order by languages.name;
			`
	)

	if rows, err = repo.connector.QueryxContext(ctx, query); err != nil {
		repo.components.Logger.Error().
			Format("Error when retrieving an items from the database: '%s'. ", err).Write()
		return
	}

	list = make([]*entities.Language, 0)

	for rows.Next() {
		var model = new(db_models.Language)

		if err = rows.StructScan(model); err != nil {
			repo.components.Logger.Error().
				Format("Error while reading item data from the database:: '%s'. ", err).Write()
			return
		}

		list = append(list, &entities.Language{
			Code:   model.Code,
			Name:   model.Name,
			Active: model.Active,
		})
	}

	return
}
