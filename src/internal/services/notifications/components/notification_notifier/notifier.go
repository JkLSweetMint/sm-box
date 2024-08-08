package notification_notifier

import (
	"context"
	"github.com/google/uuid"
	"slices"
	"sm-box/internal/services/notifications/objects/entities"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	"sync"
)

// notifier - компонент для рассылки уведомлений.
type notifier struct {
	recipients []*Recipient
	rwMx       *sync.RWMutex

	components *components
	ctx        context.Context
}

// components - компоненты.
type components struct {
	Logger logger.Logger
}

// RegisterRecipient - регистрация получателя уведомлений.
func (component *notifier) RegisterRecipient(recipient *Recipient) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelComponent)

		trace.FunctionCall(recipient)

		defer func() { trace.FunctionCallFinished() }()
	}

	component.components.Logger.Info().
		Text("The notification recipient registration process has started... ").
		Field("recipient", recipient).Write()

	defer func() {
		component.components.Logger.Info().
			Text("The notification recipient registration process is completed. ").
			Field("recipient", recipient).Write()
	}()

	// Проверки
	{
		if len(recipient.Keys) == 0 {
			component.components.Logger.Warn().
				Text("Zero number of notification recipient keys, the recipient has not been registered. ").
				Field("recipient", recipient).Write()
			return
		}
	}

	component.rwMx.Lock()
	defer component.rwMx.Unlock()

	recipient.channel = make(Channel)

	component.recipients = append(component.recipients, recipient)

	component.components.Logger.Info().
		Text("The notification recipient created. ").
		Field("recipient", recipient).Write()

	return
}

// RemoveRecipient - удаление получателя уведомлений.
func (component *notifier) RemoveRecipient(sessionJwtTokenID uuid.UUID) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelComponent)

		trace.FunctionCall(sessionJwtTokenID)

		defer func() { trace.FunctionCallFinished() }()
	}

	component.components.Logger.Info().
		Text("The notification recipient removing process has started... ").
		Field("session_jwt_token_id", sessionJwtTokenID).Write()

	defer func() {
		component.components.Logger.Info().
			Text("The notification recipient removing process is completed. ").
			Field("session_jwt_token_id", sessionJwtTokenID).Write()
	}()

	component.rwMx.Lock()
	defer component.rwMx.Unlock()

ForFindRecipient:
	for index, recipient := range component.recipients {
		if recipient == nil {
			continue
		}

		if recipient.JwtToken.ID == sessionJwtTokenID {
			component.components.Logger.Info().
				Text("The notification recipient has been found and removed. ").
				Field("session_jwt_token_id", sessionJwtTokenID).Write()

			component.recipients = append(component.recipients[:index], component.recipients[index+1:]...)
			recipient.Close()

			break ForFindRecipient
		}

		if index == len(component.recipients)-1 {
			component.components.Logger.Warn().
				Text("The notification recipient was not found. ").
				Field("session_jwt_token_id", sessionJwtTokenID).Write()
		}
	}

	return
}

// Notify - отправка уведомления получателю.
func (component *notifier) Notify(notifications ...entities.Notification) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelComponent)

		trace.FunctionCall(notifications)

		defer func() { trace.FunctionCallFinished() }()
	}

	component.components.Logger.Info().
		Text("The notification distribution process has been started... ").
		Field("notifications", notifications).Write()

	defer func() {
		component.components.Logger.Info().
			Text("The notification process has been completed. ").
			Field("notifications", notifications).Write()
	}()

	component.rwMx.RLock()
	defer component.rwMx.RUnlock()

	for _, notification := range notifications {
		if notification == nil {
			continue
		}

		for _, recipient := range component.recipients {
			if recipient == nil {
				continue
			}

			if slices.Contains(recipient.Keys, notification.Recipient()) && recipient.channel != nil {
				env.Synchronization.WaitGroup.Add(1)

				go component.notify(recipient.channel, notification)
				break
			}
		}
	}
}

// notify - отправка уведомления получателю.
func (component *notifier) notify(channel Channel, notification entities.Notification) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelComponentInternal)

		trace.FunctionCall(channel, notification)

		defer func() { trace.FunctionCallFinished() }()
	}

	defer env.Synchronization.WaitGroup.Done()

	component.components.Logger.Info().
		Text("The process of sending a notification to the recipient has been started... ").
		Field("notification", notification).Write()

	defer func() {
		component.components.Logger.Info().
			Text("The process of sending the notification to the recipient is completed. ").
			Field("notification", notification).Write()
	}()

	select {
	case channel <- notification:
		{
			component.components.Logger.Info().
				Text("The notification was successfully sent to the recipient. ").
				Field("notification", notification).Write()
		}
	case <-component.ctx.Done():
		{
			component.components.Logger.Warn().
				Text("The notification was not sent to the recipient. ").
				Field("notification", notification).Write()
			break
		}
	}
}
