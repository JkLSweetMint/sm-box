package texts_usecase

import (
	"context"
	common_errors "sm-box/internal/common/errors"
	texts_repository "sm-box/internal/services/i18n/infrastructure/repositories/texts"
	"sm-box/internal/services/i18n/objects/entities"
	srv_errors "sm-box/internal/services/i18n/objects/errors"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	"strings"
)

const (
	loggerInitiator = "infrastructure-[usecases]=texts"
)

// UseCase - логика текстов локализации.
type UseCase struct {
	components   *components
	repositories *repositories

	conf *Config
	ctx  context.Context
}

// repositories - репозитории логики.
type repositories struct {
	Texts interface {
		AssembleDictionary(ctx context.Context, lang string, paths []string) (dictionary entities.Dictionary, err error)
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

		// Texts
		{
			if usecase.repositories.Texts, err = texts_repository.New(ctx); err != nil {
				return
			}
		}
	}

	usecase.components.Logger.Info().
		Format("A '%s' usecase has been created. ", "texts").
		Field("config", usecase.conf).Write()

	return
}

// AssembleDictionary - собрать словарь локализации.
func (usecase *UseCase) AssembleDictionary(ctx context.Context, lang string, paths []string) (dictionary entities.Dictionary, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, lang, paths)
		defer func() { trc.Error(cErr).FunctionCallFinished(dictionary) }()
	}

	usecase.components.Logger.Info().
		Text("The collection of localization texts has been launched... ").
		Field("paths", paths).
		Field("lang", lang).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The collection of localization texts is completed. ").
			Field("paths", paths).
			Field("lang", lang).
			Field("dictionary", dictionary).Write()
	}()

	// Обработка входных данных
	{
		if len(paths) > 0 {
			var newPaths = make([]string, 0, len(paths))

			for _, path := range paths {
				path = strings.TrimSpace(path)

				if len(path) > 0 {
					newPaths = append(newPaths, path)
				}
			}

			paths = newPaths
		}
	}

	// Проверки
	{
		if len(paths) == 0 {
			usecase.components.Logger.Error().
				Text("Invalid value of text localization paths. ").
				Field("paths", paths).
				Field("lang", lang).Write()

			cErr = srv_errors.InvalidTextLocalizationPaths()
			return
		}

		if len(lang) == 0 {
			usecase.components.Logger.Warn().
				Text("Invalid language value, the standard language will be used. ").
				Field("lang", lang).Write()
		}
	}

	// Получение
	{
		var err error

		if dictionary, err = usecase.repositories.Texts.AssembleDictionary(ctx, lang, paths); err != nil {
			dictionary = nil

			usecase.components.Logger.Error().
				Format("Failed to get dictionary: '%s'. ", err).
				Field("paths", paths).
				Field("lang", lang).Write()

			cErr = common_errors.InternalServerError()
			cErr.SetError(err)
			return
		}
	}

	return
}
