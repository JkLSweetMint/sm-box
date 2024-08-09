package notification_notifier

import (
	authentication_entities "sm-box/internal/services/authentication/objects/entities"
)

const (
	NotificationTypeCreated = "created"
	NotificationTypeRemoved = "removed"
	NotificationTypeRead    = "read"
)

type (
	// Recipient - получатель уведомлений.
	Recipient struct {
		Keys     []string
		JwtToken *authentication_entities.JwtSessionToken
		channel  Channel
	}

	// NotificationType - тип уведомления для рассылки.
	NotificationType string

	// Notification - уведомления для рассылки.
	Notification struct {
		Type      NotificationType `json:"type"      xml:"type,attr"`
		Recipient string           `json:"recipient" xml:"recipient,attr"`
		Data      any              `json:"data"      xml:"Data"`
	}

	// Channel - канал для получения уведомлений.
	Channel chan *Notification
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
