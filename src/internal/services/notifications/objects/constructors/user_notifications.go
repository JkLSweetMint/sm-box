package constructors

import (
	"github.com/google/uuid"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/notifications/objects/types"
	"sm-box/pkg/core/components/tracer"
)

type (
	// UserNotification - конструктор пользовательского уведомления.
	UserNotification struct {
		Type types.NotificationType

		SenderID    common_types.ID
		RecipientID common_types.ID

		Title     string
		TitleI18n uuid.UUID

		Text     string
		TextI18n uuid.UUID
	}
)

// FillEmptyFields - заполнение пустых полей конструктора.
func (constructor *UserNotification) FillEmptyFields() *UserNotification {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConstructor)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(constructor) }()
	}

	return constructor
}
