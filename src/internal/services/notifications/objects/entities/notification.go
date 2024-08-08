package entities

type (
	Notification interface {
		Recipient() (recipient string)
	}
)
