package urls_management_usecase

import (
	"context"
	"database/sql"
	"errors"
	common_errors "sm-box/internal/common/errors"
	common_types "sm-box/internal/common/types"
	urls_management_repository "sm-box/internal/services/url_shortner/infrastructure/repositories/urls_management"
	urls_redis_repository "sm-box/internal/services/url_shortner/infrastructure/repositories/urls_redis"
	"sm-box/internal/services/url_shortner/objects"
	"sm-box/internal/services/url_shortner/objects/constructors"
	"sm-box/internal/services/url_shortner/objects/entities"
	srv_errors "sm-box/internal/services/url_shortner/objects/errors"
	"sm-box/internal/services/url_shortner/objects/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	"strings"
)

const (
	loggerInitiator = "infrastructure-[usecases]=urls_management"
)

// UseCase - логика управления сокращениями url запросов.
type UseCase struct {
	components   *components
	repositories *repositories

	conf *Config
	ctx  context.Context
}

// repositories - репозитории логики.
type repositories struct {
	UrlsManagement interface {
		GetList(ctx context.Context,
			search *objects.ShortUrlsListSearch,
			sort *objects.ShortUrlsListSort,
			pagination *objects.ShortUrlsListPagination,
			filters *objects.ShortUrlsListFilters,
		) (count int64, list []*entities.ShortUrl, err error)
		GetOne(ctx context.Context, id common_types.ID) (url *entities.ShortUrl, err error)
		GetOneByReduction(ctx context.Context, reduction string) (url *entities.ShortUrl, err error)

		GetUsageHistory(ctx context.Context, id common_types.ID,
			sort *objects.ShortUrlsUsageHistoryListSort,
			pagination *objects.ShortUrlsUsageHistoryListPagination,
			filters *objects.ShortUrlsUsageHistoryListFilters,
		) (count int64, history []*entities.ShortUrlUsageHistory, err error)
		GetUsageHistoryByReduction(ctx context.Context, reduction string,
			sort *objects.ShortUrlsUsageHistoryListSort,
			pagination *objects.ShortUrlsUsageHistoryListPagination,
			filters *objects.ShortUrlsUsageHistoryListFilters,
		) (count int64, history []*entities.ShortUrlUsageHistory, err error)

		Create(ctx context.Context, constructor *constructors.ShortUrl) (id common_types.ID, err error)

		Remove(ctx context.Context, id common_types.ID) (err error)
		RemoveByReduction(ctx context.Context, reduction string) (err error)

		UpdateActive(ctx context.Context, id common_types.ID, active bool) (err error)
		UpdateActiveByReduction(ctx context.Context, reduction string, active bool) (err error)

		UpdateAccesses(ctx context.Context, id common_types.ID, rolesID, permissionsID []common_types.ID) (err error)
		UpdateAccessesByReduction(ctx context.Context, reduction string, rolesID, permissionsID []common_types.ID) (err error)
	}
	UrlsRedis interface {
		Set(ctx context.Context, list ...*entities.ShortUrl) (err error)
		RemoveByReduction(ctx context.Context, reduction string) (err error)
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

		// UrlsManagement
		{
			if usecase.repositories.UrlsManagement, err = urls_management_repository.New(ctx); err != nil {
				return
			}
		}

		// UrlsRedis
		{
			if usecase.repositories.UrlsRedis, err = urls_redis_repository.New(ctx); err != nil {
				return
			}
		}
	}

	usecase.components.Logger.Info().
		Format("A '%s' usecase has been created. ", "urls_management").
		Field("config", usecase.conf).Write()

	return
}

// GetList - получение списка сокращенных url.
func (usecase *UseCase) GetList(ctx context.Context,
	search *objects.ShortUrlsListSearch,
	sort *objects.ShortUrlsListSort,
	pagination *objects.ShortUrlsListPagination,
	filters *objects.ShortUrlsListFilters,
) (count int64, list []*entities.ShortUrl, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, search, sort, pagination, filters)
		defer func() { trc.Error(cErr).FunctionCallFinished(count, list) }()
	}

	usecase.components.Logger.Info().
		Text("The process of getting a list of short urls has started... ").
		Field("search", search).
		Field("sort", sort).
		Field("pagination", pagination).
		Field("filters", filters).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of getting a list of short urls is completed. ").
			Field("search", search).
			Field("sort", sort).
			Field("pagination", pagination).
			Field("filters", filters).Write()
	}()

	// Подготовка данных
	{
		if search != nil {
			search.Global = strings.TrimSpace(search.Global)
		}

		if sort != nil {
			sort.Key = strings.TrimSpace(sort.Key)

			if sort.Key == "" {
				sort = nil
			}
		}
	}

	// Валидация
	{
		if sort != nil {
			sort.Type = strings.ToLower(strings.TrimSpace(sort.Type))

			if sort.Type != "asc" && sort.Type != "desc" {
				usecase.components.Logger.Error().
					Text("An invalid sort type value was passed. ").Write()

				cErr = common_errors.InvalidSortValue()
				cErr.Details().Set("sort_type", "Invalid value. ")

				return
			}
		}

		if filters != nil {
			if filters.StartActiveType != nil {
				var v = *filters.StartActiveType

				if v != common_types.ComparisonOperatorsEqual &&
					v != common_types.ComparisonOperatorsNotEqual &&
					v != common_types.ComparisonOperatorsGreater &&
					v != common_types.ComparisonOperatorsLess &&
					v != common_types.ComparisonOperatorsGreaterThanOrEqual &&
					v != common_types.ComparisonOperatorsLessThanOrEqual {
					usecase.components.Logger.Error().
						Text("An invalid filter value was passed. ").Write()

					cErr = common_errors.InvalidFilterValue()
					cErr.Details().Set("filter_start_active_type", "Invalid value. ")

					return
				}
			}

			if filters.EndActiveType != nil {
				var v = *filters.EndActiveType

				if v != common_types.ComparisonOperatorsEqual &&
					v != common_types.ComparisonOperatorsNotEqual &&
					v != common_types.ComparisonOperatorsGreater &&
					v != common_types.ComparisonOperatorsLess &&
					v != common_types.ComparisonOperatorsGreaterThanOrEqual &&
					v != common_types.ComparisonOperatorsLessThanOrEqual {
					usecase.components.Logger.Error().
						Text("An invalid filter value was passed. ").Write()

					cErr = common_errors.InvalidFilterValue()
					cErr.Details().Set("filter_end_active_type", "Invalid value. ")

					return
				}
			}

			if filters.NumberOfUsesType != nil {
				var v = *filters.NumberOfUsesType

				if v != common_types.ComparisonOperatorsEqual &&
					v != common_types.ComparisonOperatorsNotEqual &&
					v != common_types.ComparisonOperatorsGreater &&
					v != common_types.ComparisonOperatorsLess &&
					v != common_types.ComparisonOperatorsGreaterThanOrEqual &&
					v != common_types.ComparisonOperatorsLessThanOrEqual {
					usecase.components.Logger.Error().
						Text("An invalid filter value was passed. ").Write()

					cErr = common_errors.InvalidFilterValue()
					cErr.Details().Set("filter_end_active_type", "Invalid value. ")

					return
				}
			}
		}
	}

	// Получение
	{
		var err error

		if count, list, err = usecase.repositories.UrlsManagement.GetList(ctx, search, sort, pagination, filters); err != nil {
			usecase.components.Logger.Error().
				Format("Failed to get a list of short urls: '%s'. ", err).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The list of short URLs was successfully received. ").
			Field("list", list).
			Field("count", count).Write()
	}

	return
}

// GetOne - получение сокращенного url.
func (usecase *UseCase) GetOne(ctx context.Context, id common_types.ID) (url *entities.ShortUrl, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(cErr).FunctionCallFinished(url) }()
	}

	usecase.components.Logger.Info().
		Text("The process of obtaining a short url for id has been launched... ").
		Field("id", id).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of obtaining a short url by id is completed. ").
			Field("id", id).
			Field("url", url).Write()
	}()

	// Получение
	{
		var err error

		if url, err = usecase.repositories.UrlsManagement.GetOne(ctx, id); err != nil {
			usecase.components.Logger.Error().
				Format("Could not get the shortened url by id: '%s'. ", err).Write()

			if errors.Is(err, sql.ErrNoRows) {
				cErr = srv_errors.ShortUrlNotFound()
				return
			}

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("Short url have been successfully collected. ").
			Field("url", url).Write()
	}

	return
}

// GetOneByReduction - получение сокращенного url по сокращению.
func (usecase *UseCase) GetOneByReduction(ctx context.Context, reduction string) (url *entities.ShortUrl, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, reduction)
		defer func() { trc.Error(cErr).FunctionCallFinished(url) }()
	}

	usecase.components.Logger.Info().
		Text("The process of obtaining a short url for reduction has been launched... ").
		Field("reduction", reduction).Write()

	// Получение
	{
		var err error

		if url, err = usecase.repositories.UrlsManagement.GetOneByReduction(ctx, reduction); err != nil {
			usecase.components.Logger.Error().
				Format("Could not get the shortened url by reduction: '%s'. ", err).Write()

			if errors.Is(err, sql.ErrNoRows) {
				cErr = srv_errors.ShortUrlNotFound()
				return
			}

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("Short url have been successfully collected. ").
			Field("url", url).Write()
	}

	usecase.components.Logger.Info().
		Text("The process of obtaining a short url by reduction is completed. ").
		Field("reduction", reduction).
		Field("url", url).Write()

	return
}

// GetUsageHistory - получение истории использования сокращенного url.
func (usecase *UseCase) GetUsageHistory(ctx context.Context, id common_types.ID,
	sort *objects.ShortUrlsUsageHistoryListSort,
	pagination *objects.ShortUrlsUsageHistoryListPagination,
	filters *objects.ShortUrlsUsageHistoryListFilters,
) (count int64, history []*entities.ShortUrlUsageHistory, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, id, sort, pagination, filters)
		defer func() { trc.Error(cErr).FunctionCallFinished(count, history) }()
	}

	usecase.components.Logger.Info().
		Text("The process of getting the history of using a short url by id has been started... ").
		Field("id", id).
		Field("sort", sort).
		Field("pagination", pagination).
		Field("filters", filters).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of getting the history of using the short url by id is completed. ").
			Field("id", id).
			Field("sort", sort).
			Field("pagination", pagination).
			Field("filters", filters).Write()
	}()

	// Валидация
	{
		if sort != nil {
			sort.Key = strings.TrimSpace(sort.Key)

			if sort.Key == "" {
				sort = nil
			} else {
				sort.Type = strings.ToLower(strings.TrimSpace(sort.Type))

				if sort.Type != "asc" && sort.Type != "desc" {
					usecase.components.Logger.Error().
						Text("An invalid sort type value was passed. ").Write()

					cErr = common_errors.InvalidSortValue()
					cErr.Details().Set("sort_type", "Invalid value. ")

					return
				}
			}
		}

		if filters != nil {
			if filters.Timestamp != nil {
				var v = *filters.TimestampType

				if v != common_types.ComparisonOperatorsEqual &&
					v != common_types.ComparisonOperatorsNotEqual &&
					v != common_types.ComparisonOperatorsGreater &&
					v != common_types.ComparisonOperatorsLess &&
					v != common_types.ComparisonOperatorsGreaterThanOrEqual &&
					v != common_types.ComparisonOperatorsLessThanOrEqual {
					usecase.components.Logger.Error().
						Text("An invalid filter value was passed. ").Write()

					cErr = common_errors.InvalidFilterValue()
					cErr.Details().Set("timestamp_type", "Invalid value. ")

					return
				}
			}

			if filters.Status != nil {
				var v = *filters.Status

				if v != types.ShortUrlUsageHistoryStatusFailed &&
					v != types.ShortUrlUsageHistoryStatusSuccess &&
					v != types.ShortUrlUsageHistoryStatusForbidden {
					usecase.components.Logger.Error().
						Text("An invalid filter value was passed. ").Write()

					cErr = common_errors.InvalidFilterValue()
					cErr.Details().Set("status", "Invalid value. ")

					return
				}
			}
		}
	}

	// Получение
	{
		var err error

		if count, history, err = usecase.repositories.UrlsManagement.GetUsageHistory(ctx, id, sort, pagination, filters); err != nil {
			usecase.components.Logger.Error().
				Format("The short url usage history by id could not be retrieved: '%s'. ", err).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The history of using the short url by id has been successfully obtained. ").
			Field("history", history).
			Field("count", count).Write()
	}

	return
}

// GetUsageHistoryByReduction - получение истории использования сокращенного url по сокращению.
func (usecase *UseCase) GetUsageHistoryByReduction(ctx context.Context, reduction string,
	sort *objects.ShortUrlsUsageHistoryListSort,
	pagination *objects.ShortUrlsUsageHistoryListPagination,
	filters *objects.ShortUrlsUsageHistoryListFilters,
) (count int64, history []*entities.ShortUrlUsageHistory, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, reduction, sort, pagination, filters)
		defer func() { trc.Error(cErr).FunctionCallFinished(history) }()
	}

	usecase.components.Logger.Info().
		Text("The process of getting the history of using a short url by reduction has been started... ").
		Field("reduction", reduction).
		Field("sort", sort).
		Field("pagination", pagination).
		Field("filters", filters).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of getting the history of using the short url by reduction is completed. ").
			Field("reduction", reduction).
			Field("sort", sort).
			Field("pagination", pagination).
			Field("filters", filters).Write()
	}()

	// Валидация
	{
		if sort != nil {
			sort.Key = strings.TrimSpace(sort.Key)

			if sort.Key == "" {
				sort = nil
			} else {
				sort.Type = strings.ToLower(strings.TrimSpace(sort.Type))

				if sort.Type != "asc" && sort.Type != "desc" {
					usecase.components.Logger.Error().
						Text("An invalid sort type value was passed. ").Write()

					cErr = common_errors.InvalidSortValue()
					cErr.Details().Set("sort_type", "Invalid value. ")

					return
				}
			}
		}

		if filters != nil {
			if filters.Timestamp != nil {
				var v = *filters.TimestampType

				if v != common_types.ComparisonOperatorsEqual &&
					v != common_types.ComparisonOperatorsNotEqual &&
					v != common_types.ComparisonOperatorsGreater &&
					v != common_types.ComparisonOperatorsLess &&
					v != common_types.ComparisonOperatorsGreaterThanOrEqual &&
					v != common_types.ComparisonOperatorsLessThanOrEqual {
					usecase.components.Logger.Error().
						Text("An invalid filter value was passed. ").Write()

					cErr = common_errors.InvalidFilterValue()
					cErr.Details().Set("timestamp_type", "Invalid value. ")

					return
				}
			}

			if filters.Status != nil {
				var v = *filters.Status

				if v != types.ShortUrlUsageHistoryStatusFailed &&
					v != types.ShortUrlUsageHistoryStatusSuccess &&
					v != types.ShortUrlUsageHistoryStatusForbidden {
					usecase.components.Logger.Error().
						Text("An invalid filter value was passed. ").Write()

					cErr = common_errors.InvalidFilterValue()
					cErr.Details().Set("status", "Invalid value. ")

					return
				}
			}
		}
	}

	// Получение
	{
		var err error

		if count, history, err = usecase.repositories.UrlsManagement.GetUsageHistoryByReduction(ctx, reduction, sort, pagination, filters); err != nil {
			usecase.components.Logger.Error().
				Format("The short url usage history by reduction could not be retrieved: '%s'. ", err).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The history of using the short url by reduction has been successfully obtained. ").
			Field("history", history).
			Field("count", count).Write()
	}

	return
}

// Create - создание сокращенного url.
func (usecase *UseCase) Create(ctx context.Context, constructor *constructors.ShortUrl) (url *entities.ShortUrl, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, constructor)
		defer func() { trc.Error(cErr).FunctionCallFinished(url) }()
	}

	usecase.components.Logger.Info().
		Text("The process of creating a short url has been started... ").
		Field("constructor", constructor).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of creating a short url is completed. ").
			Field("constructor", constructor).
			Field("url", url).Write()
	}()

	var id common_types.ID

	// Подготовка входных данных
	{
		constructor.FillEmptyFields()
	}

	// Валидация
	{
		if constructor.Source = strings.TrimSpace(constructor.Source); constructor.Source == "" {
			usecase.components.Logger.Error().
				Text("An invalid type value was passed. ").Write()

			cErr = common_errors.InvalidArguments()
			cErr.Details().Set("source", "Is empty. ")

			return
		}

		if constructor.Properties.NumberOfUses < 0 {
			usecase.components.Logger.Error().
				Text("An invalid type value was passed. ").Write()

			cErr = common_errors.InvalidArguments()
			cErr.Details().Set("number_of_uses", "Negative value. ")

			return
		}

		if constructor.Properties.Type != types.ShortUrlTypeProxy && constructor.Properties.Type != types.ShortUrlTypeRedirect {
			usecase.components.Logger.Error().
				Text("An invalid type value was passed. ").Write()

			cErr = common_errors.InvalidArguments()
			cErr.Details().Set("type", "Invalid value. ")

			return
		}
	}

	// Создание
	{
		var err error

		if id, err = usecase.repositories.UrlsManagement.Create(ctx, constructor); err != nil {
			usecase.components.Logger.Error().
				Format("Could not get the shortened url by id: '%s'. ", err).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The short url was successfully created. ").
			Field("id", id).Write()
	}

	// Получение
	{
		var err error

		if url, err = usecase.repositories.UrlsManagement.GetOne(ctx, id); err != nil {
			usecase.components.Logger.Error().
				Format("Could not get the shortened url by id: '%s'. ", err).Write()

			if errors.Is(err, sql.ErrNoRows) {
				cErr = srv_errors.ShortUrlNotFound()
				return
			}

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("Short url have been successfully collected. ").
			Field("url", url).Write()
	}

	return
}

// Remove - удаление сокращенного url.
func (usecase *UseCase) Remove(ctx context.Context, id common_types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The process of deleting a short url has been started... ").
		Field("id", id).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of deleting the short url is completed. ").
			Field("id", id).Write()
	}()

	// Проверки
	{
		// Существования
		{
			if _, err := usecase.repositories.UrlsManagement.GetOne(ctx, id); err != nil {
				usecase.components.Logger.Error().
					Format("Could not get the shortened url by id: '%s'. ", err).Write()

				if errors.Is(err, sql.ErrNoRows) {
					cErr = srv_errors.ShortUrlNotFound()
					return
				}

				cErr = common_errors.InternalServerError()
				return
			}
		}
	}

	// Работа с Redis DB
	{
		var url *entities.ShortUrl

		// Получение
		{
			var err error

			if url, err = usecase.repositories.UrlsManagement.GetOne(ctx, id); err != nil {
				usecase.components.Logger.Error().
					Format("Could not get the shortened url by id: '%s'. ", err).Write()

				if errors.Is(err, sql.ErrNoRows) {
					cErr = srv_errors.ShortUrlNotFound()
					return
				}

				cErr = common_errors.InternalServerError()
				return
			}

			usecase.components.Logger.Info().
				Text("Short url have been successfully collected. ").
				Field("url", url).Write()
		}

		// Удаление из базы
		{
			if err := usecase.repositories.UrlsRedis.RemoveByReduction(ctx, url.Reduction); err != nil {
				usecase.components.Logger.Error().
					Format("The short url could not be deleted: '%s'. ", err).
					Field("reduction", url.Reduction).Write()

				cErr = common_errors.InternalServerError()
				return
			}
		}
	}

	// Удаление
	{
		var err error

		if err = usecase.repositories.UrlsManagement.Remove(ctx, id); err != nil {
			usecase.components.Logger.Error().
				Format("Couldn't delete short url by id: '%s'. ", err).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The short url has been successfully deleted. ").
			Field("id", id).Write()
	}

	return
}

// RemoveByReduction - удаление сокращенного url по сокращению.
func (usecase *UseCase) RemoveByReduction(ctx context.Context, reduction string) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, reduction)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The process of deleting a short url has been started... ").
		Field("reduction", reduction).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of deleting the short url is completed. ").
			Field("reduction", reduction).Write()
	}()

	// Проверки
	{
		// Существования
		{
			if _, err := usecase.repositories.UrlsManagement.GetOneByReduction(ctx, reduction); err != nil {
				usecase.components.Logger.Error().
					Format("Could not get the shortened url by reduction: '%s'. ", err).Write()

				if errors.Is(err, sql.ErrNoRows) {
					cErr = srv_errors.ShortUrlNotFound()
					return
				}

				cErr = common_errors.InternalServerError()
				return
			}
		}
	}

	// Работа с Redis DB
	{
		if err := usecase.repositories.UrlsRedis.RemoveByReduction(ctx, reduction); err != nil {
			usecase.components.Logger.Error().
				Format("The short url could not be deleted: '%s'. ", err).
				Field("reduction", reduction).Write()

			cErr = common_errors.InternalServerError()
			return
		}
	}

	// Удаление
	{
		var err error

		if err = usecase.repositories.UrlsManagement.RemoveByReduction(ctx, reduction); err != nil {
			usecase.components.Logger.Error().
				Format("Couldn't delete short url by reduction: '%s'. ", err).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The short url has been successfully deleted. ").
			Field("reduction", reduction).Write()
	}

	return
}

// Activate - активация сокращенного url.
func (usecase *UseCase) Activate(ctx context.Context, id common_types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The short url activation by id process has started... ").
		Field("id", id).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The short url activation by id process is completed. ").
			Field("id", id).Write()
	}()

	// Проверки
	{
		// Существования
		{
			if _, err := usecase.repositories.UrlsManagement.GetOne(ctx, id); err != nil {
				usecase.components.Logger.Error().
					Format("Could not get the shortened url by id: '%s'. ", err).Write()

				if errors.Is(err, sql.ErrNoRows) {
					cErr = srv_errors.ShortUrlNotFound()
					return
				}

				cErr = common_errors.InternalServerError()
				return
			}
		}
	}

	// Активация
	{
		var err error

		if err = usecase.repositories.UrlsManagement.UpdateActive(ctx, id, true); err != nil {
			usecase.components.Logger.Error().
				Format("Short url activation by id failed: '%s'. ", err).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The short url has been successfully activated by id. ").
			Field("id", id).Write()
	}

	// Работа с Redis DB
	{
		var url *entities.ShortUrl

		// Получение
		{
			var err error

			if url, err = usecase.repositories.UrlsManagement.GetOne(ctx, id); err != nil {
				usecase.components.Logger.Error().
					Format("Could not get the shortened url by id: '%s'. ", err).Write()

				if errors.Is(err, sql.ErrNoRows) {
					cErr = srv_errors.ShortUrlNotFound()
					return
				}

				cErr = common_errors.InternalServerError()
				return
			}

			usecase.components.Logger.Info().
				Text("Short url have been successfully collected. ").
				Field("url", url).Write()
		}

		// Запись в базу
		{
			if err := usecase.repositories.UrlsRedis.Set(ctx, url); err != nil {
				usecase.components.Logger.Error().
					Format("Failed to write short url in the redis database: '%s'. ", err).
					Field("url", url).Write()

				cErr = common_errors.InternalServerError()
				return
			}
		}
	}

	return
}

// ActivateByReduction - активация сокращенного url по сокращению.
func (usecase *UseCase) ActivateByReduction(ctx context.Context, reduction string) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, reduction)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The short url activation by reduction process has started... ").
		Field("reduction", reduction).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The short url activation by reduction process is completed. ").
			Field("reduction", reduction).Write()
	}()

	// Проверки
	{
		// Существования
		{
			if _, err := usecase.repositories.UrlsManagement.GetOneByReduction(ctx, reduction); err != nil {
				usecase.components.Logger.Error().
					Format("Could not get the shortened url by reduction: '%s'. ", err).Write()

				if errors.Is(err, sql.ErrNoRows) {
					cErr = srv_errors.ShortUrlNotFound()
					return
				}

				cErr = common_errors.InternalServerError()
				return
			}
		}
	}

	// Активация
	{
		var err error

		if err = usecase.repositories.UrlsManagement.UpdateActiveByReduction(ctx, reduction, true); err != nil {
			usecase.components.Logger.Error().
				Format("Short url activation by reduction failed: '%s'. ", err).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The short url has been successfully activated by reduction. ").
			Field("reduction", reduction).Write()
	}

	// Работа с Redis DB
	{
		var url *entities.ShortUrl

		// Получение
		{
			var err error

			if url, err = usecase.repositories.UrlsManagement.GetOneByReduction(ctx, reduction); err != nil {
				usecase.components.Logger.Error().
					Format("Could not get the shortened url by id: '%s'. ", err).Write()

				if errors.Is(err, sql.ErrNoRows) {
					cErr = srv_errors.ShortUrlNotFound()
					return
				}

				cErr = common_errors.InternalServerError()
				return
			}

			usecase.components.Logger.Info().
				Text("Short url have been successfully collected. ").
				Field("url", url).Write()
		}

		// Запись в базу
		{
			if err := usecase.repositories.UrlsRedis.Set(ctx, url); err != nil {
				usecase.components.Logger.Error().
					Format("Failed to write short url in the redis database: '%s'. ", err).
					Field("url", url).Write()

				cErr = common_errors.InternalServerError()
				return
			}
		}
	}

	return
}

// Deactivate - деактивация сокращенного url.
func (usecase *UseCase) Deactivate(ctx context.Context, id common_types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The short url deactivation by id process has started... ").
		Field("id", id).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The short url deactivation by id process is completed. ").
			Field("id", id).Write()
	}()

	// Проверки
	{
		// Существования
		{
			if _, err := usecase.repositories.UrlsManagement.GetOne(ctx, id); err != nil {
				usecase.components.Logger.Error().
					Format("Could not get the shortened url by id: '%s'. ", err).Write()

				if errors.Is(err, sql.ErrNoRows) {
					cErr = srv_errors.ShortUrlNotFound()
					return
				}

				cErr = common_errors.InternalServerError()
				return
			}
		}
	}

	// Деактивация
	{
		var err error

		if err = usecase.repositories.UrlsManagement.UpdateActive(ctx, id, false); err != nil {
			usecase.components.Logger.Error().
				Format("Short url deactivation by id failed: '%s'. ", err).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The short url has been successfully deactivated by id. ").
			Field("id", id).Write()
	}

	// Работа с Redis DB
	{
		var url *entities.ShortUrl

		// Получение
		{
			var err error

			if url, err = usecase.repositories.UrlsManagement.GetOne(ctx, id); err != nil {
				usecase.components.Logger.Error().
					Format("Could not get the shortened url by id: '%s'. ", err).Write()

				if errors.Is(err, sql.ErrNoRows) {
					cErr = srv_errors.ShortUrlNotFound()
					return
				}

				cErr = common_errors.InternalServerError()
				return
			}

			usecase.components.Logger.Info().
				Text("Short url have been successfully collected. ").
				Field("url", url).Write()
		}

		// Удаление из базы
		{
			if err := usecase.repositories.UrlsRedis.RemoveByReduction(ctx, url.Reduction); err != nil {
				usecase.components.Logger.Error().
					Format("The short url could not be deleted: '%s'. ", err).
					Field("reduction", url.Reduction).Write()

				cErr = common_errors.InternalServerError()
				return
			}
		}
	}

	return
}

// DeactivateByReduction - деактивация сокращенного url по сокращению.
func (usecase *UseCase) DeactivateByReduction(ctx context.Context, reduction string) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, reduction)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The short url deactivation by reduction process has started... ").
		Field("reduction", reduction).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The short url deactivation by reduction process is completed. ").
			Field("reduction", reduction).Write()
	}()

	// Проверки
	{
		// Существования
		{
			if _, err := usecase.repositories.UrlsManagement.GetOneByReduction(ctx, reduction); err != nil {
				usecase.components.Logger.Error().
					Format("Could not get the shortened url by reduction: '%s'. ", err).Write()

				if errors.Is(err, sql.ErrNoRows) {
					cErr = srv_errors.ShortUrlNotFound()
					return
				}

				cErr = common_errors.InternalServerError()
				return
			}
		}
	}

	// Деактивация
	{
		var err error

		if err = usecase.repositories.UrlsManagement.UpdateActiveByReduction(ctx, reduction, false); err != nil {
			usecase.components.Logger.Error().
				Format("Short url deactivation by reduction failed: '%s'. ", err).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The short url has been successfully deactivated by reduction. ").
			Field("reduction", reduction).Write()
	}

	// Работа с Redis DB
	{
		if err := usecase.repositories.UrlsRedis.RemoveByReduction(ctx, reduction); err != nil {
			usecase.components.Logger.Error().
				Format("The short url could not be deleted: '%s'. ", err).
				Field("reduction", reduction).Write()

			cErr = common_errors.InternalServerError()
			return
		}
	}

	return
}

// UpdateAccesses - обновления данных доступов к сокращенному url.
func (usecase *UseCase) UpdateAccesses(ctx context.Context, id common_types.ID, rolesID, permissionsID []common_types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, id, rolesID, permissionsID)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The process of updating the access data of the short url by id has been started... ").
		Field("id", id).
		Field("roles_id", rolesID).
		Field("permissions_id", permissionsID).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of updating the access data of the short url by id is completed. ").
			Field("id", id).
			Field("roles_id", rolesID).
			Field("permissions_id", permissionsID).Write()
	}()

	// Валидация
	{
		if rolesID == nil && permissionsID == nil {
			usecase.components.Logger.Error().
				Text("An invalid type value was passed. ").Write()

			cErr = common_errors.InvalidArguments()
			cErr.Details().Set("roles_id", "Is empty. ")
			cErr.Details().Set("permissions_id", "Is empty. ")

			return
		}
	}

	// Проверки
	{
		// Существования
		{
			if _, err := usecase.repositories.UrlsManagement.GetOne(ctx, id); err != nil {
				usecase.components.Logger.Error().
					Format("Could not get the shortened url by id: '%s'. ", err).Write()

				if errors.Is(err, sql.ErrNoRows) {
					cErr = srv_errors.ShortUrlNotFound()
					return
				}

				cErr = common_errors.InternalServerError()
				return
			}
		}
	}

	// Обновление
	{
		var err error

		if err = usecase.repositories.UrlsManagement.UpdateAccesses(ctx, id, rolesID, permissionsID); err != nil {
			usecase.components.Logger.Error().
				Format("Failed to update short url access data by id: '%s'. ", err).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The access data of the short url has been successfully updated by id. ").
			Field("id", id).Write()
	}

	return
}

// UpdateAccessesByReduction - обновления данных доступов к сокращенному url по сокращению.
func (usecase *UseCase) UpdateAccessesByReduction(ctx context.Context, reduction string, rolesID, permissionsID []common_types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelUseCase)

		trc.FunctionCall(ctx, reduction, rolesID, permissionsID)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	usecase.components.Logger.Info().
		Text("The process of updating the access data of the short url by ireductiond has been started... ").
		Field("reduction", reduction).
		Field("roles_id", rolesID).
		Field("permissions_id", permissionsID).Write()

	defer func() {
		usecase.components.Logger.Info().
			Text("The process of updating the access data of the short url by reduction is completed. ").
			Field("reduction", reduction).
			Field("roles_id", rolesID).
			Field("permissions_id", permissionsID).Write()
	}()

	// Валидация
	{
		if rolesID == nil && permissionsID == nil {
			usecase.components.Logger.Error().
				Text("An invalid type value was passed. ").Write()

			cErr = common_errors.InvalidArguments()
			cErr.Details().Set("roles_id", "Is empty. ")
			cErr.Details().Set("permissions_id", "Is empty. ")

			return
		}
	}

	// Проверки
	{
		// Существования
		{
			if _, err := usecase.repositories.UrlsManagement.GetOneByReduction(ctx, reduction); err != nil {
				usecase.components.Logger.Error().
					Format("Could not get the shortened url by reduction: '%s'. ", err).Write()

				if errors.Is(err, sql.ErrNoRows) {
					cErr = srv_errors.ShortUrlNotFound()
					return
				}

				cErr = common_errors.InternalServerError()
				return
			}
		}
	}

	// Обновление
	{
		var err error

		if err = usecase.repositories.UrlsManagement.UpdateAccessesByReduction(ctx, reduction, rolesID, permissionsID); err != nil {
			usecase.components.Logger.Error().
				Format("Failed to update short url access data by reduction: '%s'. ", err).Write()

			cErr = common_errors.InternalServerError()
			return
		}

		usecase.components.Logger.Info().
			Text("The access data of the short url has been successfully updated by reduction. ").
			Field("reduction", reduction).Write()
	}

	return
}
