package texts_repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"sm-box/internal/services/i18n/infrastructure/objects/db_models"
	"sm-box/internal/services/i18n/infrastructure/objects/entities"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/postgresql"
	"strings"
)

const (
	loggerInitiator = "infrastructure-[repositories]=texts"
)

// Repository - репозиторий текстов локализации.
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
		Format("A '%s' repository has been created. ", "texts").
		Field("config", repo.conf).Write()

	return
}

// AssembleDictionary - собрать словарь локализации.
func (repo *Repository) AssembleDictionary(ctx context.Context, lang string, paths []string) (dictionary entities.Dictionary, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, lang, paths)
		defer func() { trc.Error(err).FunctionCallFinished(dictionary) }()
	}

	dictionary = make(entities.Dictionary, 0)

	// Получение данных
	{
		var (
			rows  *sqlx.Rows
			query = new(strings.Builder)
		)

		for i, path := range paths {
			if i > 0 {
				query.WriteString("union all")
			}

			query.WriteString(fmt.Sprintf(`
				select
					key,
					value
				from
					i18n.assemble_dictionary(
							case
									 when (select count(*) from i18n.languages where code = $1) = 1
										 then $1
									 else get_default_language()
								 end,
							'%s') as (key varchar(1024), value varchar(4096))
				
				`, path))
		}

		if rows, err = repo.connector.QueryxContext(ctx, query.String(), lang); err != nil {
			repo.components.Logger.Error().
				Format("Error when retrieving an items from the database: '%s'. ", err).Write()
			return
		}

		for rows.Next() {
			var model = new(db_models.Text)

			if err = rows.StructScan(model); err != nil {
				repo.components.Logger.Error().
					Format("Error while reading item data from the database:: '%s'. ", err).Write()
				return
			}

			dictionary = append(dictionary, &entities.Text{
				Key:   model.Key,
				Value: model.Value,
			})
		}
	}

	return
}
