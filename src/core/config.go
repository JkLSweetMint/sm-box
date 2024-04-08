package core

// Config - конфигурация ядра системы.
type Config struct {
}

// BuildConfig - построения конфигурации системы.
func BuildConfig() (conf *Config, err error) {
	conf = new(Config)

	return
}
