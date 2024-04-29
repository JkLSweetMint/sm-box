package files

import "runtime"

// Dev - хранилище файлов системы в режиме разработки.
type Dev struct {
	// SystemBin - скрипты и файлы администрирования.
	SystemBin *struct {
		// Init - скрипт инициализации системы (первый запуск).
		Init string
	}
	// Var - директория часто изменяемых файлов.
	Var *struct {
		// Lib - системные библиотеки и хранилища системы.
		Lib *struct {
			// SystemDB - файл базы данных системы.
			SystemDB string
			// DashboardDB - файл базы данных панели управления.
			DashboardDB string
		}
	}
}

// Prod - хранилище файлов системы в боевом режиме.
type Prod struct {
	// SystemBin - скрипты и файлы администрирования.
	SystemBin *struct {
		// Init - скрипт инициализации системы (первый запуск).
		Init string
	}
	// Var - директория часто изменяемых файлов.
	Var *struct {
		// Lib - системные библиотеки и хранилища системы.
		Lib *struct {
			// SystemDB - файл базы данных системы.
			SystemDB string
			// DashboardDB - файл базы данных панели управления.
			DashboardDB string
		}
	}
}

// Build - построение хранилища.
func (storage *Dev) Build() *Dev {
	// SystemBin
	{
		switch runtime.GOOS {
		case "windows":
			{
				storage.SystemBin = &struct {
					Init string
				}{
					Init: "init.exe",
				}
			}
		case "linux":
			{
				storage.SystemBin = &struct {
					Init string
				}{
					Init: "init",
				}
			}
		}
	}

	// Var
	{
		storage.Var = new(struct {
			Lib *struct {
				SystemDB    string
				DashboardDB string
			}
		})

		// Lib
		{
			storage.Var.Lib = &struct {
				SystemDB    string
				DashboardDB string
			}{
				SystemDB:    "system.db",
				DashboardDB: "dashboard.db",
			}
		}
	}

	return storage
}

// Build - построение хранилища.
func (storage *Prod) Build() *Prod {
	// SystemBin
	{
		switch runtime.GOOS {
		case "windows":
			{
				storage.SystemBin = &struct {
					Init string
				}{
					Init: "init.exe",
				}
			}
		case "linux":
			{
				storage.SystemBin = &struct {
					Init string
				}{
					Init: "init",
				}
			}
		}
	}

	// Var
	{
		storage.Var = new(struct {
			Lib *struct {
				SystemDB    string
				DashboardDB string
			}
		})

		// Lib
		{
			storage.Var.Lib = &struct {
				SystemDB    string
				DashboardDB string
			}{
				SystemDB:    "system.db",
				DashboardDB: "dashboard.db",
			}
		}
	}

	return storage
}
