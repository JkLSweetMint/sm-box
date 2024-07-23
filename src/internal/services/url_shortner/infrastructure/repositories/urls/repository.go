package urls_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"sm-box/internal/services/url_shortner/objects/db_models"
	"sm-box/internal/services/url_shortner/objects/entities"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/postgresql"
)

const (
	loggerInitiator = "infrastructure-[repositories]=urls"
)

// Repository - репозиторий управления сокращениями url запросов.
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
		Format("A '%s' repository has been created. ", "urls").
		Field("config", repo.conf).Write()

	return
}

// GetActive - получить все активные сокращение url.
func (repo *Repository) GetActive(ctx context.Context) (list []*entities.ShortUrl, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(err).FunctionCallFinished(list) }()
	}

	var rows *sqlx.Rows

	// Выполнение запроса
	{
		var query = `
			select
				urls.id,
				urls.source,
				urls.reduction,
				properties.type,
				properties.number_of_uses,
				coalesce(properties.start_active, '0001-01-01 00:00:0.000000 +00:00') as start_active,
				coalesce(properties.end_active, '0001-01-01 00:00:0.000000 +00:00') as end_active
			from
				public.urls as urls
					left join public.properties properties on urls.id = properties.url
		`

		if rows, err = repo.connector.QueryxContext(ctx, query); err != nil {
			repo.components.Logger.Error().
				Format("Error when retrieving an items from the database: '%s'. ", err).Write()
			return
		}
	}

	// Чтение данных
	{
		list = make([]*entities.ShortUrl, 0)

		for rows.Next() {
			var (
				model1 = new(db_models.ShortUrl)
				model2 = new(db_models.ShortUrlProperties)
			)

			if err = rows.Scan(
				&model1.ID,
				&model1.Source,
				&model1.Reduction,
				&model2.Type,
				&model2.NumberOfUses,
				&model2.StartActive,
				&model2.EndActive); err != nil {
				repo.components.Logger.Error().
					Format("Error while reading item data from the database:: '%s'. ", err).Write()
				return
			}

			list = append(list, &entities.ShortUrl{
				ID:        model1.ID,
				Source:    model1.Source,
				Reduction: model1.Reduction,

				Properties: &entities.ShortUrlProperties{
					Type:         model2.Type,
					NumberOfUses: model2.NumberOfUses,
					StartActive:  model2.StartActive,
					EndActive:    model2.EndActive,
				},
			})
		}
	}

	return
}
