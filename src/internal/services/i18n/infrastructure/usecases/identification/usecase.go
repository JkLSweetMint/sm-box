package identification_usecase

import (
	"context"
	"database/sql"
	"errors"
	error_list "sm-box/internal/common/errors"
	common_entities "sm-box/internal/common/objects/entities"
	identification_repository "sm-box/internal/services/i18n/infrastructure/repositories/identification"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator = "infrastructure-[usecases]=identification"
)

// UseCase - логика идентификации пользователей.
type UseCase struct {
	components   *components
	repositories *repositories

	conf *Config
	ctx  context.Context
}

// repositories - репозитории логики.
type repositories struct {
	Identification interface {
		GetToken(ctx context.Context, data string) (tok *common_entities.JwtToken, err error)
	}
}

// components - компоненты логики.
type components struct {
	Logger logger.Logger
}

// New - создание логики.
func New(ctx context.Context) (usecase *UseCase, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelMain, tracer.LevelUseCase)

		trace.FunctionCall(ctx)

		defer func() { trace.Error(err).FunctionCallFinished(usecase) }()
	}

	usecase = new(UseCase)
	usecase.ctx = ctx

	// Конфигурация
	{
		usecase.conf = new(Config).Default()

		if err = usecase.conf.Read(); err != nil {
			return
		}
	}

	// Компоненты
	{
		usecase.components = new(components)

		// Logger
		{
			if usecase.components.Logger, err = logger.New(loggerInitiator); err != nil {
				return
			}
		}
	}

	// Репозитории
	{
		usecase.repositories = new(repositories)

		// Identification
		{
			if usecase.repositories.Identification, err = identification_repository.New(ctx); err != nil {
				return
			}
		}
	}

	usecase.components.Logger.Info().
		Format("A '%s' usecase has been created. ", "identification").
		Field("config", usecase.conf).Write()

	return
}

// GetToken - получение jwt токена.
func (usecase *UseCase) GetToken(ctx context.Context, data string) (tok *common_entities.JwtToken, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, data)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The receipt of the user's token has begun... ").
		Field("data", data).Write()

	var err error

	if tok, err = usecase.repositories.Identification.GetToken(ctx, data); err != nil {
		tok = nil

		usecase.components.Logger.Error().
			Format("Failed to get token: '%s'. ", err).
			Field("data", data).Write()

		if errors.Is(err, sql.ErrNoRows) {
			cErr = error_list.TokenNotFound()
			cErr.SetError(err)
			return
		}

		cErr = error_list.InternalServerError()
		cErr.SetError(err)
		return
	}

	usecase.components.Logger.Info().
		Text("The receipt of the user's token has been successfully completed. ").
		Field("token", tok).Write()

	return
}
