package notification_notifier

import (
	"context"
	"github.com/google/uuid"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
	"sync"
)

const (
	loggerInitiator = "system-[components]=notification_notifier"
)

var (
	once     = new(sync.Once)
	instance Notifier
)

type (
	// Notifier - описание компонента для рассылки уведомлений.
	Notifier interface {
		RegisterRecipient(recipient *Recipient)
		RemoveRecipient(sessionJwtTokenID uuid.UUID)
		Notify(notifications ...*Notification)
	}
)

// New - создание компонента для рассылки уведомлений.
func New(ctx context.Context) (component Notifier, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelComponent)

		trace.FunctionCall()

		defer func() { trace.FunctionCallFinished(component) }()
	}

	once.Do(func() {
		var notifier_ = &notifier{
			recipients: make([]*Recipient, 0, 10),
			rwMx:       new(sync.RWMutex),

			ctx: ctx,
		}

		// Компоненты
		{
			notifier_.components = new(components)

			// Logger
			{
				if notifier_.components.Logger, err = logger.New(loggerInitiator); err != nil {
					return
				}
			}
		}

		instance = notifier_
	})

	component = instance

	return
}
