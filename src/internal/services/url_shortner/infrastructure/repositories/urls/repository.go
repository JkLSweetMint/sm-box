package urls_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	common_types "sm-box/internal/common/types"
	authentication_entities "sm-box/internal/services/authentication/objects/entities"
	"sm-box/internal/services/url_shortner/objects/db_models"
	"sm-box/internal/services/url_shortner/objects/entities"
	"sm-box/internal/services/url_shortner/objects/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/postgresql"
)

const (
	loggerInitiator = "infrastructure-[repositories]=urls"
)

// Repository - репозиторий для работы с сокращенными url запросов.
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
				case when properties.number_of_uses = 0 then
					0
				else
					properties.number_of_uses - (
						select
							count(usage_history.*)
						from
							public.usage_history as usage_history
						where
							usage_history.url_id = urls.id
					)
				end as remained_number_of_uses,
				coalesce(properties.start_active, '0001-01-01 00:00:0.000000 +00:00') as start_active,
				coalesce(properties.end_active, '0001-01-01 00:00:0.000000 +00:00') as end_active,
				array(select accesses.role_id from public.accesses as accesses where accesses.url_id = urls.id and accesses.role_id is not null) as roles_id,
				array(select accesses.permission_id from public.accesses as accesses where accesses.url_id = urls.id and accesses.permission_id is not null) as permissions_id
			from
				public.urls as urls
					left join public.properties properties on urls.id = properties.url_id
			where
				(
					(properties.number_of_uses = 0)
						or
					(properties.number_of_uses > (
						select
							count(usage_history.*)
						from
							public.usage_history as usage_history
						where
							usage_history.url_id = urls.id
					)))
			  and
				(
					(properties.start_active is null or properties.start_active = '0001-01-01 00:00:00.000000 +00:00' or properties.start_active <= now()) and
					(properties.end_active is null or properties.start_active = '0001-01-01 00:00:00.000000 +00:00' or properties.end_active >= now())
					)
			  and
				properties.active;
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
				model1                 = new(db_models.ShortUrl)
				model2                 = new(db_models.ShortUrlProperties)
				rolesID, permissionsID pq.Int64Array
			)

			if err = rows.Scan(
				&model1.ID,
				&model1.Source,
				&model1.Reduction,
				&model2.Type,
				&model2.NumberOfUses,
				&model2.RemainedNumberOfUses,
				&model2.StartActive,
				&model2.EndActive,
				&rolesID,
				&permissionsID); err != nil {
				repo.components.Logger.Error().
					Format("Error while reading item data from the database:: '%s'. ", err).Write()
				return
			}

			var url = &entities.ShortUrl{
				ID:        model1.ID,
				Source:    model1.Source,
				Reduction: model1.Reduction,

				Accesses: &entities.ShortUrlAccesses{
					RolesID:       make([]common_types.ID, 0),
					PermissionsID: make([]common_types.ID, 0),
				},
				Properties: &entities.ShortUrlProperties{
					Type:                 model2.Type,
					NumberOfUses:         model2.NumberOfUses,
					RemainedNumberOfUses: model2.RemainedNumberOfUses,
					StartActive:          model2.StartActive,
					EndActive:            model2.EndActive,
					Active:               model2.Active,
				},
			}

			for _, id := range rolesID {
				url.Accesses.RolesID = append(url.Accesses.RolesID, common_types.ID(id))
			}

			for _, id := range permissionsID {
				url.Accesses.PermissionsID = append(url.Accesses.PermissionsID, common_types.ID(id))
			}

			list = append(list, url)
		}
	}

	return
}

// WriteCallToHistory - записать обращение по короткой ссылке в историю.
func (repo *Repository) WriteCallToHistory(ctx context.Context, id common_types.ID, status types.ShortUrlUsageHistoryStatus, token *authentication_entities.JwtSessionToken) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, id, status, token)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var query = `
			insert into
				usage_history(
					url_id,
					status,
					token_info
				) 
			values (
					$1,
					$2,
					$3
			)
		`

	if _, err = repo.connector.Exec(query, id, status, token); err != nil {
		repo.components.Logger.Error().
			Format("Error inserting an item from the database: '%s'. ", err).Write()
		return
	}

	return
}
