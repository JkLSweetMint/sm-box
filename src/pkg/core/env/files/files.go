package files

// Dev - хранилище файлов системы в режиме разработки.
type Dev struct {
	// SystemBin - скрипты и файлы администрирования.
	SystemBin *struct{}
}

// Prod - хранилище файлов системы в боевом режиме.
type Prod struct {
	// SystemBin - скрипты и файлы администрирования.
	SystemBin *struct {
		// Init - скрипт инициализации системы (первый запуск).
		Init string
	}
}

// Build - построение хранилища.
func (storage *Dev) Build() *Dev {
	return storage
}

// Build - построение хранилища.
func (storage *Prod) Build() *Prod {
	return storage
}
