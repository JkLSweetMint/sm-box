package configurator

// Config - описание конфигурации для использования диспетчером.
type Config[T any] interface {
	FillEmptyFields() T
	Default() T
	Validate() (err error)
}

// Encoder - описание кодировщика для кодирования/декодирования данных диспетчером.
type Encoder interface {
	Encode(v any) ([]byte, error)
	Decode(data []byte, v any) error
}
