package configurator

// PublicProfile - публичный профиль конфигурации.
type PublicProfile struct {
	encoder       Encoder
	dir, filename string
}

// PrivateProfile - приватный профиль конфигурации.
type PrivateProfile struct {
	encoder       Encoder
	dir, filename string
}
