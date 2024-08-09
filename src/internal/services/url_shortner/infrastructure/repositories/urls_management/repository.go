package urls_management_repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/url_shortner/objects"
	"sm-box/internal/services/url_shortner/objects/constructors"
	"sm-box/internal/services/url_shortner/objects/db_models"
	"sm-box/internal/services/url_shortner/objects/entities"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/postgresql"
	"strings"
	"time"
)

const (
	loggerInitiator = "infrastructure-[repositories]=urls_management"
)

// Repository - репозиторий управления сокращенными url запросов.
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

// GetList - получение списка сокращенных url.
func (repo *Repository) GetList(ctx context.Context,
	search *objects.ShortUrlsListSearch,
	sort *objects.ShortUrlsListSort,
	pagination *objects.ShortUrlsListPagination,
	filters *objects.ShortUrlsListFilters,
) (count int64, list []*entities.ShortUrl, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, search, sort, pagination, filters)
		defer func() { trc.Error(err).FunctionCallFinished(count, list) }()
	}

	// Основные данные
	{
		var rows *sqlx.Rows

		// Выполнение запроса
		{
			var query = new(strings.Builder)

			query.WriteString(`
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
				properties.active,
				array(select accesses.role_id from public.accesses as accesses where accesses.url_id = urls.id and accesses.role_id is not null) as roles_id,
				array(select accesses.permission_id from public.accesses as accesses where accesses.url_id = urls.id and accesses.permission_id is not null) as permissions_id
			from
				public.urls as urls
					left join public.properties properties on urls.id = properties.url_id
			where 
			    
		`)

			// Доработки запроса
			{
				if search != nil {
					query.WriteString(fmt.Sprintf("(lower(urls.source) like lower('%%%s%%') or lower(urls.reduction) like lower('%%%s%%'))", search.Global, search.Global))
				}

				if filters != nil {
					if filters.Active != nil {
						var v = *filters.Active
						query.WriteString(fmt.Sprintf("\nand active=%t", v))
					}

					if filters.Type != nil {
						var v = *filters.Type
						query.WriteString(fmt.Sprintf("\nand type='%s'", v))
					}

					if filters.NumberOfUses != nil {
						var (
							v        = *filters.NumberOfUses
							operator common_types.ComparisonOperators
						)

						if filters.NumberOfUsesType != nil {
							operator = common_types.ParseComparisonOperators(*filters.NumberOfUsesType)
						}

						query.WriteString(fmt.Sprintf("\nand number_of_uses * -1 %s%d", operator.TranslateToSign(), v))
					}

					if filters.StartActive != nil {
						var (
							v        = *filters.StartActive
							operator common_types.ComparisonOperators
						)

						if filters.StartActiveType != nil {
							operator = common_types.ParseComparisonOperators(*filters.StartActiveType)
						}

						query.WriteString(fmt.Sprintf("\nand properties.start_active%s'%s'", operator.TranslateToSign(), v.Format(time.RFC3339Nano)))
					}

					if filters.EndActive != nil {
						var (
							v        = *filters.EndActive
							operator common_types.ComparisonOperators
						)

						if filters.EndActiveType != nil {
							operator = common_types.ParseComparisonOperators(*filters.EndActiveType)
						}

						query.WriteString(fmt.Sprintf("\nand properties.end_active%s'%s'", operator.TranslateToSign(), v.Format(time.RFC3339Nano)))
					}
				}

				if sort != nil {
					if sort.Key != "" {
						query.WriteString(fmt.Sprintf("\norder by %s %s", sort.Key, sort.Type))
					}
				}

				if pagination != nil {
					if pagination.Limit != nil {
						query.WriteString(fmt.Sprintf("\nlimit %d", *pagination.Limit))
					}

					if pagination.Offset != nil {
						query.WriteString(fmt.Sprintf("\noffset %d", *pagination.Offset))
					}
				}
			}

			if rows, err = repo.connector.QueryxContext(ctx, query.String()); err != nil {
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
					&model2.Active,
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
	}

	// Получение кол-во элементов
	{
		var row *sqlx.Row

		// Выполнение запроса
		{
			var query = new(strings.Builder)

			query.WriteString(`
			select
				count(urls.*)	
			from
				public.urls as urls
					left join public.properties properties on urls.id = properties.url_id
			where 
			    
		`)

			// Доработки запроса
			{
				if search != nil {
					query.WriteString(fmt.Sprintf("(urls.source like '%%%s%%' or urls.reduction like '%%%s%%')", search.Global, search.Global))
				}

				if filters != nil {
					if filters.Active != nil {
						var v = *filters.Active
						query.WriteString(fmt.Sprintf("\nand active=%t", v))
					}

					if filters.Type != nil {
						var v = *filters.Type
						query.WriteString(fmt.Sprintf("\nand type='%s'", v))
					}

					if filters.NumberOfUses != nil {
						var (
							v        = *filters.NumberOfUses
							operator common_types.ComparisonOperators
						)

						if filters.NumberOfUsesType != nil {
							operator = common_types.ParseComparisonOperators(*filters.NumberOfUsesType)
						}

						query.WriteString(fmt.Sprintf("\nand number_of_uses * -1 %s%d", operator.TranslateToSign(), v))
					}

					if filters.StartActive != nil {
						var (
							v        = *filters.StartActive
							operator common_types.ComparisonOperators
						)

						if filters.StartActiveType != nil {
							operator = common_types.ParseComparisonOperators(*filters.StartActiveType)
						}

						query.WriteString(fmt.Sprintf("\nand properties.start_active%s'%s'", operator.TranslateToSign(), v.Format(time.RFC3339Nano)))
					}

					if filters.EndActive != nil {
						var (
							v        = *filters.EndActive
							operator common_types.ComparisonOperators
						)

						if filters.EndActiveType != nil {
							operator = common_types.ParseComparisonOperators(*filters.EndActiveType)
						}

						query.WriteString(fmt.Sprintf("\nand properties.end_active%s'%s'", operator.TranslateToSign(), v.Format(time.RFC3339Nano)))
					}
				}

				if sort != nil {
					if sort.Key != "" {
						query.WriteString(fmt.Sprintf("\norder by %s %s", sort.Key, sort.Type))
					}
				}
			}

			row = repo.connector.QueryRowxContext(ctx, query.String())

			if err = row.Err(); err != nil {
				repo.components.Logger.Error().
					Format("Error when retrieving an item from the database: '%s'. ", err).Write()
				return
			}
		}

		// Чтение данных
		{
			if err = row.Scan(&count); err != nil {
				repo.components.Logger.Error().
					Format("Error while reading item data from the database:: '%s'. ", err).Write()
				return
			}
		}
	}

	return
}

// GetOne - получение сокращенного url по id.
func (repo *Repository) GetOne(ctx context.Context, id common_types.ID) (url *entities.ShortUrl, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(err).FunctionCallFinished(url) }()
	}

	var (
		model1                 = new(db_models.ShortUrl)
		model2                 = new(db_models.ShortUrlProperties)
		rolesID, permissionsID pq.Int64Array
	)

	// Получение
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
				properties.active,
				array(select accesses.role_id from public.accesses as accesses where accesses.url_id = urls.id and accesses.role_id is not null) as roles_id,
				array(select accesses.permission_id from public.accesses as accesses where accesses.url_id = urls.id and accesses.permission_id is not null) as permissions_id
			from
				public.urls as urls
					left join public.properties properties on urls.id = properties.url_id
			where 
			    urls.id = $1
		`

		var row = repo.connector.QueryRowxContext(ctx, query, id)

		if err = row.Err(); err != nil {
			repo.components.Logger.Error().
				Format("Error when retrieving an item from the database: '%s'. ", err).Write()
			return
		}

		if err = row.Scan(
			&model1.ID,
			&model1.Source,
			&model1.Reduction,
			&model2.Type,
			&model2.NumberOfUses,
			&model2.RemainedNumberOfUses,
			&model2.StartActive,
			&model2.EndActive,
			&model2.Active,
			&rolesID,
			&permissionsID); err != nil {
			repo.components.Logger.Error().
				Format("Error while reading item data from the database:: '%s'. ", err).Write()
			return
		}
	}

	// Перенос в сущность
	{
		url = &entities.ShortUrl{
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
	}

	return
}

// GetOneByReduction - получение сокращенного url по сокращению.
func (repo *Repository) GetOneByReduction(ctx context.Context, reduction string) (url *entities.ShortUrl, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, reduction)
		defer func() { trc.Error(err).FunctionCallFinished(url) }()
	}

	var (
		model1                 = new(db_models.ShortUrl)
		model2                 = new(db_models.ShortUrlProperties)
		rolesID, permissionsID pq.Int64Array
	)

	// Получение
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
				properties.active,
				array(select accesses.role_id from public.accesses as accesses where accesses.url_id = urls.id and accesses.role_id is not null) as roles_id,
				array(select accesses.permission_id from public.accesses as accesses where accesses.url_id = urls.id and accesses.permission_id is not null) as permissions_id
			from
				public.urls as urls
					left join public.properties properties on urls.id = properties.url_id
			where 
			    urls.reduction = $1
		`

		var row = repo.connector.QueryRowxContext(ctx, query, reduction)

		if err = row.Err(); err != nil {
			repo.components.Logger.Error().
				Format("Error when retrieving an item from the database: '%s'. ", err).Write()
			return
		}

		if err = row.Scan(
			&model1.ID,
			&model1.Source,
			&model1.Reduction,
			&model2.Type,
			&model2.NumberOfUses,
			&model2.RemainedNumberOfUses,
			&model2.StartActive,
			&model2.EndActive,
			&model2.Active,
			&rolesID,
			&permissionsID); err != nil {
			repo.components.Logger.Error().
				Format("Error while reading item data from the database:: '%s'. ", err).Write()
			return
		}
	}

	// Перенос в сущность
	{
		url = &entities.ShortUrl{
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
	}

	return
}

// GetUsageHistory - получение истории использования сокращенного url.
func (repo *Repository) GetUsageHistory(ctx context.Context, id common_types.ID,
	sort *objects.ShortUrlsUsageHistoryListSort,
	pagination *objects.ShortUrlsUsageHistoryListPagination,
	filters *objects.ShortUrlsUsageHistoryListFilters,
) (count int64, history []*entities.ShortUrlUsageHistory, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, id, sort, pagination, filters)
		defer func() { trc.Error(err).FunctionCallFinished(count, history) }()
	}

	// Основные данные
	{
		var rows *sqlx.Rows

		// Выполнение запроса
		{
			var query = new(strings.Builder)

			query.WriteString(`
				select
					usage_history.status,
					usage_history.timestamp,
					usage_history.token_info
				from
					public.usage_history as usage_history
				where
					usage_history.url_id = $1
			    
		`)

			// Доработки запроса
			{
				if filters != nil {
					if filters.Status != nil {
						var v = *filters.Status
						query.WriteString(fmt.Sprintf("\nand status='%s'", v))
					}

					if filters.Timestamp != nil {
						var (
							v        = *filters.Timestamp
							operator common_types.ComparisonOperators
						)

						if filters.TimestampType != nil {
							operator = common_types.ParseComparisonOperators(*filters.TimestampType)
						}

						query.WriteString(fmt.Sprintf("\nand timestamp%s'%s'", operator.TranslateToSign(), v.Format(time.RFC3339Nano)))
					}
				}

				if sort != nil {
					if sort.Key != "" {
						query.WriteString(fmt.Sprintf("\norder by %s %s", sort.Key, sort.Type))
					}
				}

				if pagination != nil {
					if pagination.Limit != nil {
						query.WriteString(fmt.Sprintf("\nlimit %d", *pagination.Limit))
					}

					if pagination.Offset != nil {
						query.WriteString(fmt.Sprintf("\noffset %d", *pagination.Offset))
					}
				}
			}

			if rows, err = repo.connector.QueryxContext(ctx, query.String(), id); err != nil {
				repo.components.Logger.Error().
					Format("Error when retrieving an items from the database: '%s'. ", err).Write()
				return
			}
		}

		// Чтение данных
		{
			history = make([]*entities.ShortUrlUsageHistory, 0)

			for rows.Next() {
				var model = new(db_models.ShortUrlUsageHistory)

				if err = rows.StructScan(model); err != nil {
					repo.components.Logger.Error().
						Format("Error while reading item data from the database:: '%s'. ", err).Write()
					return
				}

				history = append(history, &entities.ShortUrlUsageHistory{
					Status:    model.Status,
					Timestamp: model.Timestamp,
					TokenInfo: model.TokenInfo,
				})
			}
		}
	}

	// Получение кол-во элементов
	{
		var row *sqlx.Row

		// Выполнение запроса
		{
			var query = new(strings.Builder)

			query.WriteString(`
				select
					count(usage_history.*)
				from
					public.usage_history as usage_history
				where
					usage_history.url_id = $1
			    
		`)

			// Доработки запроса
			{
				if filters != nil {
					if filters.Status != nil {
						var v = *filters.Status
						query.WriteString(fmt.Sprintf("\nand status='%s'", v))
					}

					if filters.Timestamp != nil {
						var (
							v        = *filters.Timestamp
							operator common_types.ComparisonOperators
						)

						if filters.TimestampType != nil {
							operator = common_types.ParseComparisonOperators(*filters.TimestampType)
						}

						query.WriteString(fmt.Sprintf("\nand timestamp%s'%s'", operator.TranslateToSign(), v.Format(time.RFC3339Nano)))
					}
				}

				if sort != nil {
					if sort.Key != "" {
						query.WriteString(fmt.Sprintf("\norder by %s %s", sort.Key, sort.Type))
					}
				}

				if pagination != nil {
					if pagination.Limit != nil {
						query.WriteString(fmt.Sprintf("\nlimit %d", *pagination.Limit))
					}

					if pagination.Offset != nil {
						query.WriteString(fmt.Sprintf("\noffset %d", *pagination.Offset))
					}
				}
			}

			row = repo.connector.QueryRowxContext(ctx, query.String(), id)

			if err = row.Err(); err != nil {
				repo.components.Logger.Error().
					Format("Error when retrieving an item from the database: '%s'. ", err).Write()
				return
			}
		}

		// Чтение данных
		{
			if err = row.Scan(&count); err != nil {
				repo.components.Logger.Error().
					Format("Error while reading item data from the database:: '%s'. ", err).Write()
				return
			}
		}
	}

	return
}

// GetUsageHistoryByReduction - получение истории использования сокращенного url по сокращению.
func (repo *Repository) GetUsageHistoryByReduction(ctx context.Context, reduction string,
	sort *objects.ShortUrlsUsageHistoryListSort,
	pagination *objects.ShortUrlsUsageHistoryListPagination,
	filters *objects.ShortUrlsUsageHistoryListFilters,
) (count int64, history []*entities.ShortUrlUsageHistory, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, reduction, sort, pagination, filters)
		defer func() { trc.Error(err).FunctionCallFinished(count, history) }()
	}

	// Основные данные
	{
		var rows *sqlx.Rows

		// Выполнение запроса
		{
			var query = new(strings.Builder)

			query.WriteString(`
			select
				usage_history.status,
				usage_history.timestamp,
				usage_history.token_info
			from
				public.usage_history as usage_history
			where
				usage_history.url_id = (select urls.id from public.urls as urls where urls.reduction = $1)
			    
		`)

			// Доработки запроса
			{
				if filters != nil {
					if filters.Status != nil {
						var v = *filters.Status
						query.WriteString(fmt.Sprintf("\nand status='%s'", v))
					}

					if filters.Timestamp != nil {
						var (
							v        = *filters.Timestamp
							operator common_types.ComparisonOperators
						)

						if filters.TimestampType != nil {
							operator = common_types.ParseComparisonOperators(*filters.TimestampType)
						}

						query.WriteString(fmt.Sprintf("\nand timestamp%s'%s'", operator.TranslateToSign(), v.Format(time.RFC3339Nano)))
					}
				}

				if sort != nil {
					if sort.Key != "" {
						query.WriteString(fmt.Sprintf("\norder by %s %s", sort.Key, sort.Type))
					}
				}

				if pagination != nil {
					if pagination.Limit != nil {
						query.WriteString(fmt.Sprintf("\nlimit %d", *pagination.Limit))
					}

					if pagination.Offset != nil {
						query.WriteString(fmt.Sprintf("\noffset %d", *pagination.Offset))
					}
				}
			}

			if rows, err = repo.connector.QueryxContext(ctx, query.String(), reduction); err != nil {
				repo.components.Logger.Error().
					Format("Error when retrieving an items from the database: '%s'. ", err).Write()
				return
			}
		}

		// Чтение данных
		{
			history = make([]*entities.ShortUrlUsageHistory, 0)

			for rows.Next() {
				var model = new(db_models.ShortUrlUsageHistory)

				if err = rows.StructScan(model); err != nil {
					repo.components.Logger.Error().
						Format("Error while reading item data from the database:: '%s'. ", err).Write()
					return
				}

				history = append(history, &entities.ShortUrlUsageHistory{
					Status:    model.Status,
					Timestamp: model.Timestamp,
					TokenInfo: model.TokenInfo,
				})
			}
		}
	}

	// Получение кол-во элементов
	{
		var row *sqlx.Row

		// Выполнение запроса
		{
			var query = new(strings.Builder)

			query.WriteString(`
				select
					count(usage_history.*)
				from
					public.usage_history as usage_history
				where
					usage_history.url_id = (select urls.id from public.urls as urls where urls.reduction = $1)
			    
		`)

			// Доработки запроса
			{
				if filters != nil {
					if filters.Status != nil {
						var v = *filters.Status
						query.WriteString(fmt.Sprintf("\nand status='%s'", v))
					}

					if filters.Timestamp != nil {
						var (
							v        = *filters.Timestamp
							operator common_types.ComparisonOperators
						)

						if filters.TimestampType != nil {
							operator = common_types.ParseComparisonOperators(*filters.TimestampType)
						}

						query.WriteString(fmt.Sprintf("\nand timestamp%s'%s'", operator.TranslateToSign(), v.Format(time.RFC3339Nano)))
					}
				}

				if sort != nil {
					if sort.Key != "" {
						query.WriteString(fmt.Sprintf("\norder by %s %s", sort.Key, sort.Type))
					}
				}

				if pagination != nil {
					if pagination.Limit != nil {
						query.WriteString(fmt.Sprintf("\nlimit %d", *pagination.Limit))
					}

					if pagination.Offset != nil {
						query.WriteString(fmt.Sprintf("\noffset %d", *pagination.Offset))
					}
				}
			}

			row = repo.connector.QueryRowxContext(ctx, query.String(), reduction)

			if err = row.Err(); err != nil {
				repo.components.Logger.Error().
					Format("Error when retrieving an item from the database: '%s'. ", err).Write()
				return
			}
		}

		// Чтение данных
		{
			if err = row.Scan(&count); err != nil {
				repo.components.Logger.Error().
					Format("Error while reading item data from the database:: '%s'. ", err).Write()
				return
			}
		}
	}

	return
}

// CreateOne - создание сокращенного url.
func (repo *Repository) CreateOne(ctx context.Context, constructor *constructors.ShortUrl) (id common_types.ID, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, constructor)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	// Основные данные
	{
		var query = `
			select
				public.create_short_url(
					$1,
					$2,
					$3,
					$4,
					$5,
					$6
				) as id;
		`

		var row = repo.connector.QueryRowxContext(ctx, query,
			constructor.Source,
			constructor.Properties.Type,
			constructor.Properties.NumberOfUses,
			constructor.Properties.StartActive,
			constructor.Properties.EndActive,
			constructor.Properties.Active)

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
	}

	// Доступы
	{
		if constructor.Accesses != nil &&
			(len(constructor.Accesses.RolesID) > 0 || len(constructor.Accesses.PermissionsID) > 0) {
			var tx *sqlx.Tx

			// Создание транзакции
			{
				if tx, err = repo.connector.BeginTxx(ctx, nil); err != nil {
					repo.components.Logger.Error().
						Format("An error occurred during the creation of the transaction: '%s'. ", err).Write()
					return
				}
			}

			// Загрузка новых
			{
				// Роли
				{
					if constructor.Accesses.RolesID != nil {
						var query = `
						insert into 
							   public.accesses(
								   url_id,
								   role_id
							   )
							   values (
								   $1,
								   $2
							   )
					`

						for _, role := range constructor.Accesses.RolesID {
							if _, err = tx.Exec(query, id, role); err != nil {
								repo.components.Logger.Error().
									Format("Error inserting an item from the database: '%s'. ", err).Write()
								return
							}
						}
					}
				}

				// Права
				{
					if constructor.Accesses.PermissionsID != nil {
						var query = `
				insert into 
					   public.accesses(
						   url_id,
						   permission_id
					   )
					   values (
						   $1,
						   $2
					   )
			`

						for _, permission := range constructor.Accesses.PermissionsID {
							if _, err = tx.Exec(query, id, permission); err != nil {
								repo.components.Logger.Error().
									Format("Error inserting an item from the database: '%s'. ", err).Write()
								return
							}
						}
					}
				}
			}

			if err = tx.Commit(); err != nil {
				repo.components.Logger.Error().
					Format("An error occurred during the execution of the transaction: '%s'. ", err).Write()
				return
			}
		}
	}

	return
}

// Remove - удаление сокращенного url.
func (repo *Repository) Remove(ctx context.Context, id common_types.ID) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var query = `
			delete from
				   public.urls as urls
			where
			    urls.id = $1
		`

	if _, err = repo.connector.Exec(query, id); err != nil {
		repo.components.Logger.Error().
			Format("Error removing an item from the database: '%s'. ", err).Write()
		return
	}

	return
}

// RemoveByReduction - удаление сокращенного url по сокращению.
func (repo *Repository) RemoveByReduction(ctx context.Context, reduction string) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, reduction)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var query = `
			delete from
				   public.urls as urls
			where
			    urls.reduction = $1
		`

	if _, err = repo.connector.Exec(query, reduction); err != nil {
		repo.components.Logger.Error().
			Format("Error removing an item from the database: '%s'. ", err).Write()
		return
	}

	return
}

// UpdateActive - обновления данных по активации сокращенного url.
func (repo *Repository) UpdateActive(ctx context.Context, id common_types.ID, active bool) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, id, active)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var query = `
			update
				public.properties
			set
			    active = $2
			where
			    url_id = $1
		`

	if _, err = repo.connector.Exec(query, id, active); err != nil {
		repo.components.Logger.Error().
			Format("Error updating an item from the database: '%s'. ", err).Write()
		return
	}

	return
}

// UpdateActiveByReduction - обновления данных по активации сокращенного url по сокращению.
func (repo *Repository) UpdateActiveByReduction(ctx context.Context, reduction string, active bool) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, reduction, active)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var query = `
			update
				public.properties
			set
			    active = $2
			where
			    url_id = (select urls.id from public.urls as urls where urls.reduction = $1)
		`

	if _, err = repo.connector.Exec(query, reduction, active); err != nil {
		repo.components.Logger.Error().
			Format("Error updating an item from the database: '%s'. ", err).Write()
		return
	}

	return
}

// UpdateAccesses - обновления данных доступов к сокращенному url.
func (repo *Repository) UpdateAccesses(ctx context.Context, id common_types.ID, rolesID, permissionsID []common_types.ID) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, id, rolesID, permissionsID)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var tx *sqlx.Tx

	// Создание транзакции
	{
		if tx, err = repo.connector.BeginTxx(ctx, nil); err != nil {
			repo.components.Logger.Error().
				Format("An error occurred during the creation of the transaction: '%s'. ", err).Write()
			return
		}
	}

	// Очистка текущих
	{
		// Роли
		{
			if rolesID != nil {
				var query = `
					delete from
						   public.accesses as accesses
					where
						accesses.url_id = $1 and
					    accesses.role_id is not null
				`

				if _, err = tx.Exec(query, id); err != nil {
					repo.components.Logger.Error().
						Format("Error removing an item from the database: '%s'. ", err).Write()
					return
				}
			}
		}

		// Права
		{
			if permissionsID != nil {
				var query = `
					delete from
						   public.accesses as accesses
					where
						accesses.url_id = $1 and
					    accesses.permission_id is not null
				`

				if _, err = tx.Exec(query, id); err != nil {
					repo.components.Logger.Error().
						Format("Error removing an item from the database: '%s'. ", err).Write()
					return
				}
			}
		}
	}

	// Загрузка новых
	{
		// Роли
		{
			if rolesID != nil {
				var query = `
				insert into 
					   public.accesses(
						   url_id,
						   role_id
					   )
					   values (
						   $1,
						   $2
					   )
			`

				for _, role := range rolesID {
					if _, err = tx.Exec(query, id, role); err != nil {
						repo.components.Logger.Error().
							Format("Error inserting an item from the database: '%s'. ", err).Write()
						return
					}
				}
			}
		}

		// Права
		{
			if permissionsID != nil {
				var query = `
				insert into 
					   public.accesses(
						   url_id,
						   permission_id
					   )
					   values (
						   $1,
						   $2
					   )
			`

				for _, permission := range permissionsID {
					if _, err = tx.Exec(query, id, permission); err != nil {
						repo.components.Logger.Error().
							Format("Error inserting an item from the database: '%s'. ", err).Write()
						return
					}
				}
			}
		}
	}

	if err = tx.Commit(); err != nil {
		repo.components.Logger.Error().
			Format("An error occurred during the execution of the transaction: '%s'. ", err).Write()
		return
	}

	return
}

// UpdateAccessesByReduction - обновления данных доступов к сокращенному url по сокращению.
func (repo *Repository) UpdateAccessesByReduction(ctx context.Context, reduction string, rolesID, permissionsID []common_types.ID) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, reduction, rolesID, permissionsID)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var id common_types.ID

	// Получение ID
	{
		var query = `
			select
				urls.id
			from
				public.urls as urls
			where 
			    urls.reduction = $1
		`

		var row = repo.connector.QueryRowxContext(ctx, query, reduction)

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
	}

	var tx *sqlx.Tx

	// Создание транзакции
	{
		if tx, err = repo.connector.BeginTxx(ctx, nil); err != nil {
			repo.components.Logger.Error().
				Format("An error occurred during the creation of the transaction: '%s'. ", err).Write()
			return
		}
	}

	// Очистка текущих
	{
		// Роли
		{
			if rolesID != nil {
				var query = `
					delete from
						   public.accesses as accesses
					where
						accesses.url_id = $1 and
					    accesses.role_id is not null
				`

				if _, err = tx.Exec(query, id); err != nil {
					repo.components.Logger.Error().
						Format("Error removing an item from the database: '%s'. ", err).Write()
					return
				}
			}
		}

		// Права
		{
			if permissionsID != nil {
				var query = `
					delete from
						   public.accesses as accesses
					where
						accesses.url_id = $1 and
					    accesses.permission_id is not null
				`

				if _, err = tx.Exec(query, id); err != nil {
					repo.components.Logger.Error().
						Format("Error removing an item from the database: '%s'. ", err).Write()
					return
				}
			}
		}
	}

	// Загрузка новых
	{
		// Роли
		{
			if rolesID != nil {
				var query = `
				insert into 
					   public.accesses(
						   url_id,
						   role_id
					   )
					   values (
						   $1,
						   $2
					   )
			`

				for _, role := range rolesID {
					if _, err = tx.Exec(query, id, role); err != nil {
						repo.components.Logger.Error().
							Format("Error inserting an item from the database: '%s'. ", err).Write()
						return
					}
				}
			}
		}

		// Права
		{
			if permissionsID != nil {
				var query = `
				insert into 
					   public.accesses(
						   url_id,
						   permission_id
					   )
					   values (
						   $1,
						   $2
					   )
			`

				for _, permission := range permissionsID {
					if _, err = tx.Exec(query, id, permission); err != nil {
						repo.components.Logger.Error().
							Format("Error inserting an item from the database: '%s'. ", err).Write()
						return
					}
				}
			}
		}
	}

	if err = tx.Commit(); err != nil {
		repo.components.Logger.Error().
			Format("An error occurred during the execution of the transaction: '%s'. ", err).Write()
		return
	}

	return
}
