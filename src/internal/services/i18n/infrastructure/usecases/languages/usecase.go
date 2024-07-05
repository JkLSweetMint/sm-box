package languages_usecase

import (
	"context"
	error_list "sm-box/internal/common/errors"
	"sm-box/internal/services/i18n/infrastructure/objects/entities"
	languages_repository "sm-box/internal/services/i18n/infrastructure/repositories/languages"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
)

const (
	loggerInitiator = "infrastructure-[usecases]=languages"
)

// UseCase - логика языков локализации.
type UseCase struct {
	components   *components
	repositories *repositories

	conf *Config
	ctx  context.Context
}

// repositories - репозитории логики.
type repositories struct {
	Languages interface {
		GetList(ctx context.Context) (list []*entities.Language, err error)
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

		// Languages
		{
			if usecase.repositories.Languages, err = languages_repository.New(ctx); err != nil {
				return
			}
		}
	}

	usecase.components.Logger.Info().
		Format("A '%s' usecase has been created. ", "languages").
		Field("config", usecase.conf).Write()

	return
}

// GetList - получение списка языков.
func (usecase *UseCase) GetList(ctx context.Context) (list []*entities.Language, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx)
		defer func() { trc.Error(cErr).FunctionCallFinished(list) }()
	}

	// Получение
	{
		var err error

		if list, err = usecase.repositories.Languages.GetList(ctx); err != nil {
			list = nil

			usecase.components.Logger.Error().
				Format("Couldn't get localization languages: '%s'. ", err).Write()

			cErr = error_list.InternalServerError()
			cErr.SetError(err)
			return
		}
	}

	return
}

// Remove - удаление языка.
// Текста и ресурсы локализации так же удаляются.
func (usecase *UseCase) Remove(ctx context.Context, code string) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, code)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	return
}

// Update - обновление данных языка.
func (usecase *UseCase) Update(ctx context.Context, code, name string) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, code, name)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	return
}

// Create - создание языка.
func (usecase *UseCase) Create(ctx context.Context, code string, name string) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, code, name)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	return
}

// Activate - активировать язык.
func (usecase *UseCase) Activate(ctx context.Context, code string) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, code)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	return
}

// Deactivate - деактивировать язык.
func (usecase *UseCase) Deactivate(ctx context.Context, code string) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, code)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	return
}
