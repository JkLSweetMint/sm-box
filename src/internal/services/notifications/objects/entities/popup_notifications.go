package entities

import (
	"github.com/google/uuid"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/notifications/objects/models"
	"sm-box/internal/services/notifications/objects/types"
	"sm-box/pkg/core/components/tracer"
	"time"
)

type (
	// PopupNotification - всплывающее уведомление.
	PopupNotification struct {
		ID   uuid.UUID
		Type types.NotificationType

		SenderID    common_types.ID
		RecipientID string

		Title     string
		TitleI18n uuid.UUID

		Text     string
		TextI18n uuid.UUID

		CreatedTimestamp time.Time
	}
)

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *PopupNotification) FillEmptyFields() *PopupNotification {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(entity) }()
	}

	if entity.ID.String() == "00000000-0000-0000-0000-000000000000" {
		entity.ID = uuid.New()
	}

	return entity
}

// ToModel - получение внешней модели.
func (entity *PopupNotification) ToModel() (model *models.PopupNotificationInfo) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	entity.FillEmptyFields()

	model = &models.PopupNotificationInfo{
		ID:   entity.ID,
		Type: entity.Type,

		SenderID:    entity.SenderID,
		RecipientID: entity.RecipientID,

		Title:     entity.Title,
		TitleI18n: entity.TitleI18n,

		Text:     entity.Text,
		TextI18n: entity.TextI18n,

		CreatedTimestamp: entity.CreatedTimestamp,
	}

	return
}

// Recipient - получение получателя уведомления
func (entity *PopupNotification) Recipient() (recipient string) {
	recipient = entity.RecipientID
	return
}
