package types

type (
	// Details - детали ошибки.
	Details interface {
		PeekAll() (data map[string]string)
		Peek(k string) (v string)
		Set(k string, v string) Details
		Reset() Details

		PeekFields() (data []DetailsField)
		PeekFieldMessage(k string) (m DetailsFieldMessage)
		SetField(k DetailsFieldKey, m DetailsFieldMessage) Details
		SetFields(fields ...DetailsField) Details
		ResetFields() Details

		Clone() Details
	}
)

type (
	// DetailsFieldMessage - сообщение поля.
	DetailsFieldMessage Message

	// DetailsFieldKey - ключ поля.
	DetailsFieldKey interface {
		Add(path ...string) DetailsFieldKey
		AddArray(name string, index int) DetailsFieldKey
		AddMap(name string, key any) DetailsFieldKey
		String() (str string)

		Clone() DetailsFieldKey
	}

	// DetailsField - поле c сообщением об ошибке.
	DetailsField struct {
		Key     DetailsFieldKey
		Message DetailsFieldMessage
	}
)
