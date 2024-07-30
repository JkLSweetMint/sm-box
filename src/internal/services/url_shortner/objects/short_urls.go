package objects

import (
	"sm-box/internal/services/url_shortner/objects/types"
	"time"
)

type (
	// ShortUrlsListSearch - атрибуты для поиска коротких маршрутов.
	ShortUrlsListSearch struct {
		Global string
	}

	// ShortUrlsListPagination - атрибуты для пагинации коротких маршрутов.
	ShortUrlsListPagination struct {
		Offset *int64
		Limit  *int64
	}

	// ShortUrlsListSort - атрибуты для сортировки коротких маршрутов.
	ShortUrlsListSort struct {
		Key  string
		Type string
	}

	// ShortUrlsListFilters - атрибуты для фильтрации коротких маршрутов.
	ShortUrlsListFilters struct {
		Active *bool
		Type   *types.ShortUrlType

		NumberOfUses     *int64
		NumberOfUsesType *string

		StartActive     *time.Time
		StartActiveType *string

		EndActive     *time.Time
		EndActiveType *string
	}

	// ShortUrlsUsageHistoryListPagination - атрибуты для пагинации истории использования коротких маршрутов.
	ShortUrlsUsageHistoryListPagination struct {
		Offset *int64
		Limit  *int64
	}

	// ShortUrlsUsageHistoryListSort - атрибуты для сортировки истории использования коротких маршрутов.
	ShortUrlsUsageHistoryListSort struct {
		Key  string
		Type string
	}

	// ShortUrlsUsageHistoryListFilters - атрибуты для фильтрации истории использования коротких маршрутов.
	ShortUrlsUsageHistoryListFilters struct {
		Status *types.ShortUrlUsageHistoryStatus

		Timestamp     *time.Time
		TimestampType *string
	}
)
