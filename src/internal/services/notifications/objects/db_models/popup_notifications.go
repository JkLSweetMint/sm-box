package db_models

import (
	"github.com/google/uuid"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/notifications/objects/types"
	"time"
)

type (
	// PopupNotification - модель базы данных всплывающих уведомлений.
	PopupNotification struct {
		ID   common_types.ID        `db:"id"`
		Type types.NotificationType `db:"type"`

		SenderID    common_types.ID `db:"sender_id"`
		RecipientID string          `db:"recipient_id"`

		Title     string    `db:"title"`
		TitleI18n uuid.UUID `db:"title_i18n"`

		Text     string    `db:"text"`
		TextI18n uuid.UUID `db:"text_i18n"`

		CreatedTimestamp time.Time `db:"created_timestamp"`
	}
)
