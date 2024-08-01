package service

import (
	"context"
	common_types "sm-box/internal/common/types"
	authentication_entities "sm-box/internal/services/authentication/objects/entities"
	urls_controller "sm-box/internal/services/url_shortner/infrastructure/controllers/urls"
	urls_management_controller "sm-box/internal/services/url_shortner/infrastructure/controllers/urls_management"
	"sm-box/internal/services/url_shortner/objects"
	"sm-box/internal/services/url_shortner/objects/constructors"
	"sm-box/internal/services/url_shortner/objects/models"
	"sm-box/internal/services/url_shortner/objects/types"
	c_errors "sm-box/pkg/errors"
)

// Controllers - описание контроллеров сервиса.
type Controllers interface {
	Urls() interface {
		RegisterToRedisDB(ctx context.Context) (cErr c_errors.Error)
		GetByReductionFromRedisDB(ctx context.Context, reduction string) (url *models.ShortUrlInfo, cErr c_errors.Error)
		UpdateInRedisDB(ctx context.Context, url *models.ShortUrlInfo) (cErr c_errors.Error)
		RemoveByReductionFromRedisDB(ctx context.Context, reduction string) (cErr c_errors.Error)

		Use(ctx context.Context, reduction string, token *authentication_entities.JwtSessionToken) (url *models.ShortUrlInfo, status types.ShortUrlUsageHistoryStatus, cErr c_errors.Error)
	}
	UrlsManagement() interface {
		GetList(ctx context.Context,
			search *objects.ShortUrlsListSearch,
			sort *objects.ShortUrlsListSort,
			pagination *objects.ShortUrlsListPagination,
			filters *objects.ShortUrlsListFilters,
		) (count int64, list []*models.ShortUrlInfo, cErr c_errors.Error)
		GetOne(ctx context.Context, id common_types.ID) (url *models.ShortUrlInfo, cErr c_errors.Error)
		GetOneByReduction(ctx context.Context, reduction string) (url *models.ShortUrlInfo, cErr c_errors.Error)

		GetUsageHistory(ctx context.Context, id common_types.ID,
			sort *objects.ShortUrlsUsageHistoryListSort,
			pagination *objects.ShortUrlsUsageHistoryListPagination,
			filters *objects.ShortUrlsUsageHistoryListFilters,
		) (count int64, history []*models.ShortUrlUsageHistoryInfo, cErr c_errors.Error)
		GetUsageHistoryByReduction(ctx context.Context, reduction string,
			sort *objects.ShortUrlsUsageHistoryListSort,
			pagination *objects.ShortUrlsUsageHistoryListPagination,
			filters *objects.ShortUrlsUsageHistoryListFilters,
		) (count int64, history []*models.ShortUrlUsageHistoryInfo, cErr c_errors.Error)

		Create(ctx context.Context, constructor *constructors.ShortUrl) (url *models.ShortUrlInfo, cErr c_errors.Error)

		Remove(ctx context.Context, id common_types.ID) (cErr c_errors.Error)
		RemoveByReduction(ctx context.Context, reduction string) (cErr c_errors.Error)

		Activate(ctx context.Context, id common_types.ID) (cErr c_errors.Error)
		ActivateByReduction(ctx context.Context, reduction string) (cErr c_errors.Error)

		Deactivate(ctx context.Context, id common_types.ID) (cErr c_errors.Error)
		DeactivateByReduction(ctx context.Context, reduction string) (cErr c_errors.Error)

		UpdateAccesses(ctx context.Context, id common_types.ID, rolesID, permissionsID []common_types.ID) (cErr c_errors.Error)
		UpdateAccessesByReduction(ctx context.Context, reduction string, rolesID, permissionsID []common_types.ID) (cErr c_errors.Error)
	}
}

// controllers - контроллеры сервиса.
type controllers struct {
	urls           *urls_controller.Controller
	urlsManagement *urls_management_controller.Controller
}

// Urls - получение контроллера сервиса.
func (controllers *controllers) Urls() interface {
	RegisterToRedisDB(ctx context.Context) (cErr c_errors.Error)
	GetByReductionFromRedisDB(ctx context.Context, reduction string) (url *models.ShortUrlInfo, cErr c_errors.Error)
	UpdateInRedisDB(ctx context.Context, url *models.ShortUrlInfo) (cErr c_errors.Error)
	RemoveByReductionFromRedisDB(ctx context.Context, reduction string) (cErr c_errors.Error)

	Use(ctx context.Context, reduction string, token *authentication_entities.JwtSessionToken) (url *models.ShortUrlInfo, status types.ShortUrlUsageHistoryStatus, cErr c_errors.Error)
} {
	return controllers.urls
}

// UrlsManagement - получение контроллера сервиса.
func (controllers *controllers) UrlsManagement() interface {
	GetList(ctx context.Context,
		search *objects.ShortUrlsListSearch,
		sort *objects.ShortUrlsListSort,
		pagination *objects.ShortUrlsListPagination,
		filters *objects.ShortUrlsListFilters,
	) (count int64, list []*models.ShortUrlInfo, cErr c_errors.Error)
	GetOne(ctx context.Context, id common_types.ID) (url *models.ShortUrlInfo, cErr c_errors.Error)
	GetOneByReduction(ctx context.Context, reduction string) (url *models.ShortUrlInfo, cErr c_errors.Error)

	GetUsageHistory(ctx context.Context, id common_types.ID,
		sort *objects.ShortUrlsUsageHistoryListSort,
		pagination *objects.ShortUrlsUsageHistoryListPagination,
		filters *objects.ShortUrlsUsageHistoryListFilters,
	) (count int64, history []*models.ShortUrlUsageHistoryInfo, cErr c_errors.Error)
	GetUsageHistoryByReduction(ctx context.Context, reduction string,
		sort *objects.ShortUrlsUsageHistoryListSort,
		pagination *objects.ShortUrlsUsageHistoryListPagination,
		filters *objects.ShortUrlsUsageHistoryListFilters,
	) (count int64, history []*models.ShortUrlUsageHistoryInfo, cErr c_errors.Error)

	Create(ctx context.Context, constructor *constructors.ShortUrl) (url *models.ShortUrlInfo, cErr c_errors.Error)

	Remove(ctx context.Context, id common_types.ID) (cErr c_errors.Error)
	RemoveByReduction(ctx context.Context, reduction string) (cErr c_errors.Error)

	Activate(ctx context.Context, id common_types.ID) (cErr c_errors.Error)
	ActivateByReduction(ctx context.Context, reduction string) (cErr c_errors.Error)

	Deactivate(ctx context.Context, id common_types.ID) (cErr c_errors.Error)
	DeactivateByReduction(ctx context.Context, reduction string) (cErr c_errors.Error)

	UpdateAccesses(ctx context.Context, id common_types.ID, rolesID, permissionsID []common_types.ID) (cErr c_errors.Error)
	UpdateAccessesByReduction(ctx context.Context, reduction string, rolesID, permissionsID []common_types.ID) (cErr c_errors.Error)
} {
	return controllers.urlsManagement
}
