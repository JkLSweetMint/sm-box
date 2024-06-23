package repository

import (
	"context"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/databases/connectors/sqlite3"
)

const (
	loggerInitiator = "transports-[http]-[rest_api]-[components]-[access_system]=repository"
)

// Repository - репозиторий системы доступа http rest api.
type Repository struct {
	*usersRepository
	*tokensRepository
}

// components - компоненты репозитория.
type components struct {
	Logger logger.Logger
}

// New - создание репозитория.
func New(ctx context.Context, conf *Config) (repo *Repository, err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelMain, tracer.LevelRepository)

		trc.FunctionCall(ctx, conf)
		defer func() { trc.Error(err).FunctionCallFinished(repo) }()
	}

	// Конфигурация
	{
		if err = conf.FillEmptyFields().Validate(); err != nil {
			return
		}
	}

	repo = &Repository{
		usersRepository: &usersRepository{
			connector:  nil,
			components: nil,
			conf:       conf,
			ctx:        ctx,
		},
		tokensRepository: &tokensRepository{
			connector:  nil,
			components: nil,
			conf:       conf,
			ctx:        ctx,
		},
	}

	// Компоненты
	{
		var cmps = new(components)

		// Logger
		{
			if cmps.Logger, err = logger.New(loggerInitiator); err != nil {
				return
			}
		}

		repo.usersRepository.components = cmps
		repo.tokensRepository.components = cmps
	}

	// Коннектор базы данных
	{
		var conn sqlite3.Connector

		if conn, err = sqlite3.New(ctx, conf.Connector); err != nil {
			return
		}

		repo.usersRepository.connector = conn
		repo.tokensRepository.connector = conn
	}

	repo.usersRepository.components.Logger.Info().
		Text("The repository of the access system component has been created. ").
		Field("config", conf).Write()

	return
}
