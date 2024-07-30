package urls_management_controller

import (
	"context"
	common_types "sm-box/internal/common/types"
	urls_management_usecase "sm-box/internal/services/url_shortner/infrastructure/usecases/urls_management"
	"sm-box/internal/services/url_shortner/objects"
	"sm-box/internal/services/url_shortner/objects/entities"
	"sm-box/internal/services/url_shortner/objects/models"
	"sm-box/internal/services/url_shortner/objects/types"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	c_errors "sm-box/pkg/errors"
	"time"
)

const (
	loggerInitiator = "infrastructure-[controllers]=urls_management"
)

// Controller - контроллер базовой аутентификации пользователей.
type Controller struct {
	components *components
	usecases   *usecases

	conf *Config
	ctx  context.Context
}

// usecases - логика контроллера.
type usecases struct {
	UrlsManagement interface {
		GetList(ctx context.Context,
			search *objects.ShortUrlsListSearch,
			sort *objects.ShortUrlsListSort,
			pagination *objects.ShortUrlsListPagination,
			filters *objects.ShortUrlsListFilters,
		) (list []*entities.ShortUrl, cErr c_errors.Error)
		GetOne(ctx context.Context, id common_types.ID) (url *entities.ShortUrl, cErr c_errors.Error)
		GetOneByReduction(ctx context.Context, reduction string) (url *entities.ShortUrl, cErr c_errors.Error)

		GetUsageHistory(ctx context.Context, id common_types.ID,
			sort *objects.ShortUrlsUsageHistoryListSort,
			pagination *objects.ShortUrlsUsageHistoryListPagination,
			filters *objects.ShortUrlsUsageHistoryListFilters,
		) (history []*entities.ShortUrlUsageHistory, cErr c_errors.Error)
		GetUsageHistoryByReduction(ctx context.Context, reduction string,
			sort *objects.ShortUrlsUsageHistoryListSort,
			pagination *objects.ShortUrlsUsageHistoryListPagination,
			filters *objects.ShortUrlsUsageHistoryListFilters,
		) (history []*entities.ShortUrlUsageHistory, cErr c_errors.Error)

		Create(ctx context.Context,
			source string,
			type_ types.ShortUrlType,
			numberOfUses int64,
			startActive, endActive time.Time) (url *entities.ShortUrl, cErr c_errors.Error)

		Remove(ctx context.Context, id common_types.ID) (cErr c_errors.Error)
		RemoveByReduction(ctx context.Context, reduction string) (cErr c_errors.Error)

		Activate(ctx context.Context, id common_types.ID) (cErr c_errors.Error)
		ActivateByReduction(ctx context.Context, reduction string) (cErr c_errors.Error)

		Deactivate(ctx context.Context, id common_types.ID) (cErr c_errors.Error)
		DeactivateByReduction(ctx context.Context, reduction string) (cErr c_errors.Error)
	}
}

// components - компоненты контроллера.
type components struct {
	Logger logger.Logger
}

// New - создание контроллера.
func New(ctx context.Context) (controller *Controller, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelMain, tracer.LevelController)

		trace.FunctionCall(ctx)

		defer func() { trace.Error(err).FunctionCallFinished(controller) }()
	}

	controller = new(Controller)
	controller.ctx = ctx

	// Конфигурация
	{
		controller.conf = new(Config).Default()

		if err = controller.conf.Read(); err != nil {
			return
		}
	}

	// Компоненты
	{
		controller.components = new(components)

		// Logger
		{
			if controller.components.Logger, err = logger.New(loggerInitiator); err != nil {
				return
			}
		}
	}

	// Логика
	{
		controller.usecases = new(usecases)

		// UrlsManagement
		{
			if controller.usecases.UrlsManagement, err = urls_management_usecase.New(ctx); err != nil {
				return
			}
		}
	}

	controller.components.Logger.Info().
		Format("A '%s' controller has been created. ", "urls_management").
		Field("config", controller.conf).Write()

	return
}

// GetList - получение списка сокращенных url.
func (controller *Controller) GetList(ctx context.Context,
	search *objects.ShortUrlsListSearch,
	sort *objects.ShortUrlsListSort,
	pagination *objects.ShortUrlsListPagination,
	filters *objects.ShortUrlsListFilters,
) (list []*models.ShortUrlInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, search, sort, pagination, filters)
		defer func() { trc.Error(cErr).FunctionCallFinished(list) }()
	}

	// Выполнения инструкций
	{
		var urls []*entities.ShortUrl

		if urls, cErr = controller.usecases.UrlsManagement.GetList(ctx, search, sort, pagination, filters); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}

		// Преобразование в модель
		{
			list = make([]*models.ShortUrlInfo, 0, len(urls))

			for _, url := range urls {
				list = append(list, url.ToModel())
			}
		}
	}

	return
}

// GetOne - получение сокращенного url.
func (controller *Controller) GetOne(ctx context.Context, id common_types.ID) (url *models.ShortUrlInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(cErr).FunctionCallFinished(url) }()
	}

	// Выполнения инструкций
	{
		var url_ *entities.ShortUrl

		if url_, cErr = controller.usecases.UrlsManagement.GetOne(ctx, id); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}

		// Преобразование в модель
		{
			if url_ != nil {
				url = url_.ToModel()
			}
		}
	}

	return
}

// GetOneByReduction - получение сокращенного url по сокращению.
func (controller *Controller) GetOneByReduction(ctx context.Context, reduction string) (url *models.ShortUrlInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, reduction)
		defer func() { trc.Error(cErr).FunctionCallFinished(url) }()
	}

	// Выполнения инструкций
	{
		var url_ *entities.ShortUrl

		if url_, cErr = controller.usecases.UrlsManagement.GetOneByReduction(ctx, reduction); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}

		// Преобразование в модель
		{
			if url_ != nil {
				url = url_.ToModel()
			}
		}
	}

	return
}

// GetUsageHistory - получение истории использования сокращенного url.
func (controller *Controller) GetUsageHistory(ctx context.Context, id common_types.ID,
	sort *objects.ShortUrlsUsageHistoryListSort,
	pagination *objects.ShortUrlsUsageHistoryListPagination,
	filters *objects.ShortUrlsUsageHistoryListFilters,
) (history []*models.ShortUrlUsageHistoryInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, id, sort, pagination, filters)
		defer func() { trc.Error(cErr).FunctionCallFinished(history) }()
	}

	// Выполнения инструкций
	{
		var history_ []*entities.ShortUrlUsageHistory

		if history_, cErr = controller.usecases.UrlsManagement.GetUsageHistory(ctx, id, sort, pagination, filters); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}

		// Преобразование в модель
		{
			history = make([]*models.ShortUrlUsageHistoryInfo, 0, len(history_))

			for _, h := range history_ {
				history = append(history, h.ToModel())
			}
		}
	}

	return
}

// GetUsageHistoryByReduction - получение истории использования сокращенного url по сокращению.
func (controller *Controller) GetUsageHistoryByReduction(ctx context.Context, reduction string,
	sort *objects.ShortUrlsUsageHistoryListSort,
	pagination *objects.ShortUrlsUsageHistoryListPagination,
	filters *objects.ShortUrlsUsageHistoryListFilters,
) (history []*models.ShortUrlUsageHistoryInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, reduction, sort, pagination, filters)
		defer func() { trc.Error(cErr).FunctionCallFinished(history) }()
	}

	// Выполнения инструкций
	{
		var history_ []*entities.ShortUrlUsageHistory

		if history_, cErr = controller.usecases.UrlsManagement.GetUsageHistoryByReduction(ctx, reduction, sort, pagination, filters); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}

		// Преобразование в модель
		{
			history = make([]*models.ShortUrlUsageHistoryInfo, 0, len(history_))

			for _, h := range history_ {
				history = append(history, h.ToModel())
			}
		}
	}

	return
}

// Create - создание сокращенного url.
func (controller *Controller) Create(ctx context.Context,
	source string,
	type_ types.ShortUrlType,
	numberOfUses int64,
	startActive, endActive time.Time,
) (url *models.ShortUrlInfo, cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, source, type_, numberOfUses, startActive, endActive)
		defer func() { trc.Error(cErr).FunctionCallFinished(url) }()
	}

	// Выполнения инструкций
	{
		var url_ *entities.ShortUrl

		if url_, cErr = controller.usecases.UrlsManagement.Create(ctx, source, type_, numberOfUses, startActive, endActive); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}

		// Преобразование в модель
		{
			if url_ != nil {
				url = url_.ToModel()
			}
		}
	}

	return
}

// Remove - удаление сокращенного url.
func (controller *Controller) Remove(ctx context.Context, id common_types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	// Выполнения инструкций
	{
		if cErr = controller.usecases.UrlsManagement.Remove(ctx, id); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	return
}

// RemoveByReduction - удаление сокращенного url по сокращению.
func (controller *Controller) RemoveByReduction(ctx context.Context, reduction string) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, reduction)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	// Выполнения инструкций
	{
		if cErr = controller.usecases.UrlsManagement.RemoveByReduction(ctx, reduction); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	return
}

// Activate - активация сокращенного url.
func (controller *Controller) Activate(ctx context.Context, id common_types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	// Выполнения инструкций
	{
		if cErr = controller.usecases.UrlsManagement.Activate(ctx, id); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	return
}

// ActivateByReduction - активация сокращенного url по сокращению.
func (controller *Controller) ActivateByReduction(ctx context.Context, reduction string) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, reduction)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	// Выполнения инструкций
	{
		if cErr = controller.usecases.UrlsManagement.ActivateByReduction(ctx, reduction); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	return
}

// Deactivate - деактивация сокращенного url.
func (controller *Controller) Deactivate(ctx context.Context, id common_types.ID) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, id)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	// Выполнения инструкций
	{
		if cErr = controller.usecases.UrlsManagement.Deactivate(ctx, id); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	return
}

// DeactivateByReduction - деактивация сокращенного url по сокращению.
func (controller *Controller) DeactivateByReduction(ctx context.Context, reduction string) (cErr c_errors.Error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelController)

		trc.FunctionCall(ctx, reduction)
		defer func() { trc.Error(cErr).FunctionCallFinished() }()
	}

	// Выполнения инструкций
	{
		if cErr = controller.usecases.UrlsManagement.DeactivateByReduction(ctx, reduction); cErr != nil {
			controller.components.Logger.Error().
				Format("The controller instructions were executed with an error: '%s'. ", cErr).Write()

			return
		}
	}

	return
}
