package popup_notifications_repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/notifications/objects/constructors"
	"sm-box/internal/services/notifications/objects/db_models"
	"sm-box/internal/services/notifications/objects/entities"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/postgresql"
)

const (
	loggerInitiator = "infrastructure-[repositories]=popup_notifications"
)

// Repository - репозиторий управления всплывающими уведомлениями.
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
		Format("A '%s' repository has been created. ", "popup_notifications").
		Field("config", repo.conf).Write()

	return
}

// GetOne - получение всплывающего уведомления по id.
func (repo *Repository) GetOne(ctx context.Context, id common_types.ID) (notification *entities.PopupNotification, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(err).FunctionCallFinished(notification) }()
	}

	var model = new(db_models.PopupNotification)

	// Получение
	{
		var query = `
				select
					notifications.id,
					notifications.type,
					coalesce(notifications.sender_id, 0) as sender_id,
					notifications.recipient_id,
					notifications.title,
					coalesce(notifications.title_i18n, '00000000-0000-0000-0000-000000000000') as title_i18n,
					notifications.text,
					coalesce(notifications.text_i18n, '00000000-0000-0000-0000-000000000000') as text_i18n,
					coalesce(notifications.created_timestamp, '0001-01-01 00:00:0.000000 +00:00') as created_timestamp
				from
					popups.notifications as notifications
				where
					notifications.id = $1
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
		notification = &entities.PopupNotification{
			ID:   model.ID,
			Type: model.Type,

			SenderID:    model.SenderID,
			RecipientID: model.RecipientID,

			Title:     model.Title,
			TitleI18n: model.TitleI18n,

			Text:     model.Text,
			TextI18n: model.TextI18n,

			CreatedTimestamp: model.CreatedTimestamp,
		}
	}

	return
}

// Get - получение всплывающих уведомлений по списку id.
func (repo *Repository) Get(ctx context.Context, ids ...common_types.ID) (list []*entities.PopupNotification, err error) {
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
					notifications.id,
					notifications.type,
					coalesce(notifications.sender_id, 0) as sender_id,
					notifications.recipient_id,
					notifications.title,
					coalesce(notifications.title_i18n, '00000000-0000-0000-0000-000000000000') as title_i18n,
					notifications.text,
					coalesce(notifications.text_i18n, '00000000-0000-0000-0000-000000000000') as text_i18n,
					coalesce(notifications.created_timestamp, '0001-01-01 00:00:0.000000 +00:00') as created_timestamp
				from
					popups.notifications as notifications
				where
				    notifications.id = any($1)
`
		ids_ = make(pq.Int64Array, 0, len(ids))
	)

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
		list = make([]*entities.PopupNotification, 0)

		for rows.Next() {
			var (
				model = new(db_models.PopupNotification)
			)

			if err = rows.StructScan(model); err != nil {
				repo.components.Logger.Error().
					Format("Error while reading item data from the database:: '%s'. ", err).Write()
				return
			}

			list = append(list, &entities.PopupNotification{
				ID:   model.ID,
				Type: model.Type,

				SenderID:    model.SenderID,
				RecipientID: model.RecipientID,

				Title:     model.Title,
				TitleI18n: model.TitleI18n,

				Text:     model.Text,
				TextI18n: model.TextI18n,

				CreatedTimestamp: model.CreatedTimestamp,
			})
		}
	}

	return
}

// CreateOne - создание всплывающего уведомления.
func (repo *Repository) CreateOne(ctx context.Context, constructor *constructors.PopupNotification) (id common_types.ID, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, constructor)
		defer func() { trc.Error(err).FunctionCallFinished(id) }()
	}

	var query = `
			insert into
				popups.notifications(
					type,
					sender_id,
					recipient_id,
					title,
					title_i18n,
					text,
					text_i18n
				) 
			values (
					$1,
					$2,
					$3,
					$4,
					$5,
					$6,
					$7
			)
			returning id;
		`

	var row = repo.connector.QueryRowxContext(ctx, query,
		constructor.Type,
		constructor.SenderID,
		constructor.RecipientID,
		constructor.Title,
		constructor.TitleI18n,
		constructor.Text,
		constructor.TextI18n)

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

// Create - создание всплывающих уведомлений.
func (repo *Repository) Create(ctx context.Context, constructors ...*constructors.PopupNotification) (ids []common_types.ID, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, constructors)
		defer func() { trc.Error(err).FunctionCallFinished(ids) }()
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

	// Добавлений инструкций
	{
		var query = `
			insert into
				popups.notifications(
					type,
					sender_id,
					recipient_id,
					title,
					title_i18n,
					text,
					text_i18n
				) 
			values (
					$1,
					$2,
					$3,
					$4,
					$5,
					$6,
					$7
			)
			returning id;
		`

		for _, constructor := range constructors {
			var id common_types.ID

			var row = tx.QueryRowxContext(ctx, query,
				constructor.Type,
				constructor.SenderID,
				constructor.RecipientID,
				constructor.Title,
				constructor.TitleI18n,
				constructor.Text,
				constructor.TextI18n)

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

			ids = append(ids, id)
		}
	}

	// Выполнение транзакции
	{
		if err = tx.Commit(); err != nil {
			repo.components.Logger.Error().
				Format("An error occurred during the execution of the transaction: '%s'. ", err).Write()
			return
		}
	}

	return
}
