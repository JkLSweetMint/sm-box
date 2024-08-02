package user_notifications_repository

import (
	"context"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/notifications/objects"
	"sm-box/internal/services/notifications/objects/constructors"
	"sm-box/internal/services/notifications/objects/entities"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/postgresql"
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

// GetList - получение списка пользовательских уведомлений.
func (repo *Repository) GetList(ctx context.Context,
	userID common_types.ID,
	search *objects.UserNotificationSearch,
	pagination *objects.UserNotificationPagination,
	filters *objects.UserNotificationFilters,
) (count int64, list []*entities.UserNotification, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, userID, search, pagination, filters)
		defer func() { trc.Error(err).FunctionCallFinished(count, list) }()
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
func (repo *Repository) RemoveOne(ctx context.Context, userID, id common_types.ID) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, userID, id)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	return
}

// Remove - удаление пользовательских уведомлений.
func (repo *Repository) Remove(ctx context.Context, userID common_types.ID, ids ...common_types.ID) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, userID, ids)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	return
}

// ReadOne - чтение пользовательского уведомления.
func (repo *Repository) ReadOne(ctx context.Context, userID, id common_types.ID) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, userID, id)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	return
}

// Read - чтение пользовательских уведомлений.
func (repo *Repository) Read(ctx context.Context, userID common_types.ID, ids ...common_types.ID) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelRepository)

		trc.FunctionCall(ctx, userID, ids)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	return
}
