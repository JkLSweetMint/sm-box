package configurator

// PublicProfile - публичный профиль конфигурации.
type PublicProfile struct {
	Encoder       Encoder
	Dir, Filename string
}

// PrivateProfile - приватный профиль конфигурации.
type PrivateProfile struct {
	Encoder       Encoder
	Dir, Filename string
}
