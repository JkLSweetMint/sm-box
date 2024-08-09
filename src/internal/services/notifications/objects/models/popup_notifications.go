package models

import (
	"github.com/google/uuid"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/notifications/objects/types"
	"time"
)

type (
	// PopupNotificationInfo - внешняя модель с информацией по всплывающему уведомлению.
	PopupNotificationInfo struct {
		ID   common_types.ID        `json:"id"   xml:"id,attr"`
		Type types.NotificationType `json:"type" xml:"type,attr"`

		SenderID    common_types.ID `json:"sender_id"    xml:"sender_id,attr"`
		RecipientID string          `json:"recipient_id" xml:"recipient_id,attr"`

		Title     string    `json:"title"      xml:"Title"`
		TitleI18n uuid.UUID `json:"title_i18n" xml:"TitleI18N"`

		Text     string    `json:"text"      xml:"Text"`
		TextI18n uuid.UUID `json:"text_i18n" xml:"TextI18N"`

		CreatedTimestamp time.Time `json:"created_timestamp" xml:"created_timestamp,attr"`
	}
)
