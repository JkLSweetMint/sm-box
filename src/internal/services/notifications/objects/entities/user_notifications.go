package entities

import (
	"fmt"
	"github.com/google/uuid"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/notifications/objects/models"
	"sm-box/internal/services/notifications/objects/types"
	"sm-box/pkg/core/components/tracer"
	"time"
)

type (
	// UserNotification - пользовательское уведомление
	UserNotification struct {
		ID   common_types.ID
		Type types.NotificationType

		SenderID    common_types.ID
		RecipientID common_types.ID

		Title     string
		TitleI18n uuid.UUID

		Text     string
		TextI18n uuid.UUID

		CreatedTimestamp time.Time
		ReadTimestamp    time.Time
		RemovedTimestamp time.Time
	}
)

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *UserNotification) FillEmptyFields() *UserNotification {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(entity) }()
	}

	return entity
}

// ToModel - получение внешней модели.
func (entity *UserNotification) ToModel() (model *models.UserNotificationInfo) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	entity.FillEmptyFields()

	model = &models.UserNotificationInfo{
		ID:   entity.ID,
		Type: entity.Type,

		SenderID:    entity.SenderID,
		RecipientID: entity.RecipientID,

		Title:     entity.Title,
		TitleI18n: entity.TitleI18n,

		Text:     entity.Text,
		TextI18n: entity.TextI18n,

		CreatedTimestamp: entity.CreatedTimestamp,
		ReadTimestamp:    entity.ReadTimestamp,
		RemovedTimestamp: entity.RemovedTimestamp,
	}

	return
}

// Recipient - получение получателя уведомления
func (entity *UserNotification) Recipient() (recipient string) {
	recipient = fmt.Sprintf("users:%d", entity.RecipientID)
	return
}
