package objects

import (
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/notifications/objects/types"
)

type (
	// UserNotificationSearch - атрибуты для поиска пользовательских уведомлений.
	UserNotificationSearch struct {
		Global string
	}

	// UserNotificationPagination - атрибуты для пагинации пользовательских уведомлений.
	UserNotificationPagination struct {
		Offset *int64
		Limit  *int64
	}

	// UserNotificationFilters - атрибуты для фильтрации пользовательских уведомлений.
	UserNotificationFilters struct {
		Type     *types.NotificationType
		NotRead  *bool
		SenderID *common_types.ID
	}
)
