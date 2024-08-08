package notification_notifier

import (
	authentication_entities "sm-box/internal/services/authentication/objects/entities"
	"sm-box/internal/services/notifications/objects/entities"
)

type (
	// Channel - канал для получения уведомлений.
	Channel chan entities.Notification

	// Recipient - получатель уведомлений.
	Recipient struct {
		Keys     []string
		JwtToken *authentication_entities.JwtSessionToken
		channel  Channel
	}
)

// Channel - получение канала получателя.
func (recipient *Recipient) Channel() (channel Channel) {
	channel = recipient.channel
	return
}

// Close - закрытие канала получателя.
func (recipient *Recipient) Close() {
	close(recipient.channel)
	return
}
