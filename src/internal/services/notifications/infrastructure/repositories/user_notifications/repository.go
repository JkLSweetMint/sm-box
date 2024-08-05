package user_notifications_repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/notifications/objects"
	"sm-box/internal/services/notifications/objects/constructors"
	"sm-box/internal/services/notifications/objects/db_models"
	"sm-box/internal/services/notifications/objects/entities"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/postgresql"
	"strings"
)

const (
	loggerInitiator = "infrastructure-[repositories]=user_notifications"
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

// Exists - проверка существования.
func (repo *Repository) Exists(ctx context.Context, ids ...common_types.ID) (exists []bool, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, ids)
		defer func() { trc.Error(err).FunctionCallFinished(exists) }()
	}

	var (
		rows *sqlx.Rows
		ids_ = make(pq.Int64Array, 0, len(ids))
	)

	// Подготовка
	{
		for _, id := range ids {
			ids_ = append(ids_, int64(id))
		}
	}

	// Выполнение запроса
	{
		var query = `
			select
				notifications.id is not null as exist
			from
				(
					select
						*
					from
						unnest($1::bigint[]) as id
				) as ids
					left join users.notifications notifications on notifications.id = ids.id and
					                                               notifications.removed_timestamp is null
			order by ids.id;
`
		if rows, err = repo.connector.QueryxContext(ctx, query, ids_); err != nil {
			repo.components.Logger.Error().
				Format("Error when retrieving an items from the database: '%s'. ", err).Write()
			return
		}
	}

	// Чтение данных
	{
		exists = make([]bool, 0)

		for rows.Next() {
			var exist bool

			if err = rows.Scan(&exist); err != nil {
				repo.components.Logger.Error().
					Format("Error while reading item data from the database:: '%s'. ", err).Write()
				return
			}

			exists = append(exists, exist)
		}
	}

	return
}

// AlreadyRead - проверка что уже прочитаны.
func (repo *Repository) AlreadyRead(ctx context.Context, ids ...common_types.ID) (read []bool, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, ids)
		defer func() { trc.Error(err).FunctionCallFinished(read) }()
	}

	var (
		rows *sqlx.Rows
		ids_ = make(pq.Int64Array, 0, len(ids))
	)

	// Подготовка
	{
		for _, id := range ids {
			ids_ = append(ids_, int64(id))
		}
	}

	// Выполнение запроса
	{
		var query = `
			select
				notifications.id is not null as exist
			from
				(
					select
						*
					from
						unnest($1::bigint[]) as id
				) as ids
					left join users.notifications notifications on notifications.id = ids.id
																	and notifications.removed_timestamp is null
																	and notifications.read_timestamp is not null
			order by ids.id;
`
		if rows, err = repo.connector.QueryxContext(ctx, query, ids_); err != nil {
			repo.components.Logger.Error().
				Format("Error when retrieving an items from the database: '%s'. ", err).Write()
			return
		}
	}

	// Чтение данных
	{
		read = make([]bool, 0)

		for rows.Next() {
			var r bool

			if err = rows.Scan(&r); err != nil {
				repo.components.Logger.Error().
					Format("Error while reading item data from the database:: '%s'. ", err).Write()
				return
			}

			read = append(read, r)
		}
	}

	return
}

// GetList - получение списка пользовательских уведомлений.
func (repo *Repository) GetList(ctx context.Context,
	recipientID common_types.ID,
	search *objects.UserNotificationSearch,
	pagination *objects.UserNotificationPagination,
	filters *objects.UserNotificationFilters,
) (count, countNotRead int64, list []*entities.UserNotification, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, recipientID, search, pagination, filters)
		defer func() { trc.Error(err).FunctionCallFinished(count, countNotRead, list) }()
	}

	// Основные данные
	{
		var rows *sqlx.Rows

		// Выполнение запроса
		{
			var query = new(strings.Builder)

			query.WriteString(`
				select
					notifications.id,
					notifications.type,
					coalesce(notifications.sender_id, 0) as sender_id,
					notifications.recipient_id,
					notifications.title,
					coalesce(notifications.title_i18n, '00000000-0000-0000-0000-000000000000') as title_i18n,
					notifications.text,
					coalesce(notifications.text_i18n, '00000000-0000-0000-0000-000000000000') as text_i18n,
					coalesce(notifications.created_timestamp, '0001-01-01 00:00:0.000000 +00:00') as created_timestamp,
					coalesce(notifications.read_timestamp, '0001-01-01 00:00:0.000000 +00:00') as read_timestamp,
					coalesce(notifications.removed_timestamp, '0001-01-01 00:00:0.000000 +00:00') as removed_timestamp
				from
					users.notifications as notifications
				where
					notifications.recipient_id = $1 and
					notifications.removed_timestamp is null
			`)

			// Доработки запроса
			{
				if search != nil {
					query.WriteString(fmt.Sprintf("\nand (lower(notifications.title) like lower('%%%s%%') or lower(notifications.text) like lower('%%%s%%'))", search.Global, search.Global))
				}

				if filters != nil {
					if filters.Type != nil {
						var v = *filters.Type
						query.WriteString(fmt.Sprintf("\nand type='%s'", v))
					}

					if filters.NotRead != nil {
						if v := *filters.NotRead; v {
							query.WriteString("\nand notifications.read_timestamp is null")
						} else {
							query.WriteString("\nand notifications.read_timestamp is not null")
						}
					}

					if filters.SenderID != nil {
						var v = *filters.SenderID
						query.WriteString(fmt.Sprintf("\nand sender_id='%d'", v))
					}
				}

				query.WriteString("\n order by notifications.read_timestamp desc,  notifications.created_timestamp desc")

				if pagination != nil {
					if pagination.Limit != nil {
						query.WriteString(fmt.Sprintf("\nlimit %d", *pagination.Limit))
					}

					if pagination.Offset != nil {
						query.WriteString(fmt.Sprintf("\noffset %d", *pagination.Offset))
					}
				}
			}

			if rows, err = repo.connector.QueryxContext(ctx, query.String(), recipientID); err != nil {
				repo.components.Logger.Error().
					Format("Error when retrieving an items from the database: '%s'. ", err).Write()
				return
			}
		}

		// Чтение данных
		{
			list = make([]*entities.UserNotification, 0)

			for rows.Next() {
				var (
					model = new(db_models.UserNotification)
				)

				if err = rows.StructScan(model); err != nil {
					repo.components.Logger.Error().
						Format("Error while reading item data from the database:: '%s'. ", err).Write()
					return
				}

				list = append(list, &entities.UserNotification{
					ID:   model.ID,
					Type: model.Type,

					SenderID:    model.SenderID,
					RecipientID: model.RecipientID,

					Title:     model.Title,
					TitleI18n: model.TitleI18n,

					Text:     model.Text,
					TextI18n: model.TextI18n,

					CreatedTimestamp: model.CreatedTimestamp,
					ReadTimestamp:    model.ReadTimestamp,
					RemovedTimestamp: model.RemovedTimestamp,
				})
			}
		}
	}

	// Получение кол-во элементов
	{
		var row *sqlx.Row

		// Выполнение запроса
		{
			var (
				subQuery1 = new(strings.Builder)
				subQuery2 = new(strings.Builder)
			)

			// Общее кол-во
			{
				subQuery1.WriteString(`
					select
						count(notifications.*)
					from
						users.notifications as notifications
					where
						notifications.recipient_id = $1 and
						notifications.removed_timestamp is null
				`)

				// Доработки запроса
				{
					if search != nil {
						subQuery1.WriteString(fmt.Sprintf("\nand (lower(notifications.title) like lower('%%%s%%') or lower(notifications.text) like lower('%%%s%%'))", search.Global, search.Global))
					}

					if filters != nil {
						if filters.Type != nil {
							var v = *filters.Type
							subQuery1.WriteString(fmt.Sprintf("\nand type='%s'", v))
						}

						if filters.NotRead != nil {
							if v := *filters.NotRead; v {
								subQuery1.WriteString("\nand notifications.read_timestamp is null")
							} else {
								subQuery1.WriteString("\nand notifications.read_timestamp is not null")
							}
						}

						if filters.SenderID != nil {
							var v = *filters.SenderID
							subQuery1.WriteString(fmt.Sprintf("\nand sender_id='%d'", v))
						}
					}
				}
			}

			// Не прочитано
			{
				subQuery2.WriteString(`
					select
						count(notifications.*)
					from
						users.notifications as notifications
					where
						notifications.recipient_id = $1 and
						notifications.read_timestamp is null and
						notifications.removed_timestamp is null
				`)
			}

			var query = fmt.Sprintf(`
			select
                (%s),
				(%s)
			`, subQuery1.String(), subQuery2.String())

			row = repo.connector.QueryRowxContext(ctx, query, recipientID)

			if err = row.Err(); err != nil {
				repo.components.Logger.Error().
					Format("Error when retrieving an item from the database: '%s'. ", err).Write()
				return
			}
		}

		// Чтение данных
		{
			if err = row.Scan(&count, &countNotRead); err != nil {
				repo.components.Logger.Error().
					Format("Error while reading item data from the database:: '%s'. ", err).Write()
				return
			}
		}
	}

	return
}

// CreateOne - создание пользовательского уведомления.
func (repo *Repository) CreateOne(ctx context.Context, constructor *constructors.UserNotification) (notification *entities.UserNotification, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, constructor)
		defer func() { trc.Error(err).FunctionCallFinished(notification) }()
	}

	return
}

// Create - создание пользовательских уведомлений.
func (repo *Repository) Create(ctx context.Context, constructors ...*constructors.UserNotification) (notifications []*entities.UserNotification, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, constructors)
		defer func() { trc.Error(err).FunctionCallFinished(notifications) }()
	}

	return
}

// RemoveOne - удаление пользовательского уведомления.
func (repo *Repository) RemoveOne(ctx context.Context, recipientID, id common_types.ID) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, recipientID, id)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var query = `
			update
				users.notifications
			set
			    removed_timestamp = now()
			where
			    recipient_id = $1 and
			    id = $2
		`

	if _, err = repo.connector.Exec(query, recipientID, id); err != nil {
		repo.components.Logger.Error().
			Format("Error updating an item from the database: '%s'. ", err).Write()
		return
	}

	return
}

// Remove - удаление пользовательских уведомлений.
func (repo *Repository) Remove(ctx context.Context, recipientID common_types.ID, ids ...common_types.ID) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, recipientID, ids)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var ids_ = make(pq.Int64Array, 0, len(ids))

	// Подготовка
	{
		for _, id := range ids {
			ids_ = append(ids_, int64(id))
		}
	}

	// Выполнение запроса
	{
		var query = `
			update
				users.notifications
			set
			    removed_timestamp = now()
			where
			    recipient_id = $1 and
			    id = any($2)
		`

		if _, err = repo.connector.Exec(query, recipientID, ids_); err != nil {
			repo.components.Logger.Error().
				Format("Error updating an item from the database: '%s'. ", err).Write()
			return
		}
	}

	return
}

// ReadOne - чтение пользовательского уведомления.
func (repo *Repository) ReadOne(ctx context.Context, recipientID, id common_types.ID) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, recipientID, id)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var query = `
			update
				users.notifications
			set
			    read_timestamp = now()
			where
			    recipient_id = $1 and
			    id = $2
		`

	if _, err = repo.connector.Exec(query, recipientID, id); err != nil {
		repo.components.Logger.Error().
			Format("Error updating an item from the database: '%s'. ", err).Write()
		return
	}

	return
}

// Read - чтение пользовательских уведомлений.
func (repo *Repository) Read(ctx context.Context, recipientID common_types.ID, ids ...common_types.ID) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, recipientID, ids)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var ids_ = make(pq.Int64Array, 0, len(ids))

	// Подготовка
	{
		for _, id := range ids {
			ids_ = append(ids_, int64(id))
		}
	}

	// Выполнение запроса
	{
		var query = `
			update
				users.notifications
			set
			    read_timestamp = now()
			where
			    recipient_id = $1 and
			    id = any($2)
		`

		if _, err = repo.connector.Exec(query, recipientID, ids_); err != nil {
			repo.components.Logger.Error().
				Format("Error updating an item from the database: '%s'. ", err).Write()
			return
		}
	}

	return
}
