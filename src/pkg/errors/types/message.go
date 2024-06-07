package types

type (
	// Message - описание сообщения ошибки.
	Message interface {
		String() (str string)
		Clone() Message
	}
)
