package paths

import (
	"path"
	"testing"
)

// Dev - хранилища путей системы в режиме разработки.
type Dev struct {
	// SystemLocation - расположение системы.
	// Подгружается автоматически при запуске системы.
	SystemLocation string

	// Bin - путь к директории исполняемых файлов.
	// Стандартное расположение: /bin
	Bin string
	// Etc - путь к директории конфигурационных файлов.
	// Стандартное расположение: /etc
	Etc string

	// Src - директория исходных файлов.
	// Стандартное расположение: /src
	Src *struct {
		// Path - путь к директории.
		Path string
	}

	// System - директория системных файлов.
	// Стандартное расположение: /system
	System *struct {
		// Path - путь к директории.
		Path string
	}

	// Var - директория часто изменяемых файлов.
	// Стандартное расположение: /var
	Var *struct {
		// Path - путь к директории.
		Path string

		// Temp - путь к директории временных файлов.
		// Стандартное расположение: /temp
		Temp string
		// Logs - путь к директории файлов журнала.
		// Стандартное расположение: /logs
		Logs string
		// Data - путь к директории файлов системы.
		// Стандартное расположение: /data
		Data string
		// Cache - путь к директории для кеширования файлов.
		// Стандартное расположение: /cache
		Cache string
		// Lib - путь к директории для хранения системных библиотек и хранилищ системы.
		// Стандартное расположение: /lib
		Lib string
		// Run - путь к директории PID файлов.
		// Стандартное расположение: /run
		Run string

		// Test - директория тестовых файлов.
		// Стандартное расположение: /test
		Test *struct {
			// Path - путь к директории.
			Path string

			// Data - путь к директории тестовых файлов системы.
			// Стандартное расположение: /data
			Data string
			// Cache - путь к директории для кеширования тестовых файлов.
			// Стандартное расположение: /cache
			Cache string
		}
	}
}

// Prod - хранилища путей системы в боевом режиме.
type Prod struct {
	// SystemLocation - расположение системы.
	// Подгружается автоматически при запуске системы.
	SystemLocation string

	// Bin - путь к директории исполняемых файлов.
	// Стандартное расположение: /bin
	Bin string
	// Etc - путь к директории конфигурационных файлов.
	// Стандартное расположение: /etc
	Etc string

	// Src - директория исходных файлов.
	// Стандартное расположение: /src
	Src *struct {
		// Path - путь к директории.
		Path string
	}

	// System - директория системных файлов.
	// Стандартное расположение: /system
	System *struct {
		// Path - путь к директории.
		Path string
	}

	// Var - директория часто изменяемых файлов.
	// Стандартное расположение: /var
	Var *struct {
		// Path - путь к директории.
		Path string

		// Temp - путь к директории временных файлов.
		// Стандартное расположение: /temp
		Temp string
		// Logs - путь к директории файлов журнала.
		// Стандартное расположение: /logs
		Logs string
		// Data - путь к директории файлов системы.
		// Стандартное расположение: /data
		Data string
		// Cache - путь к директории для кеширования файлов.
		// Стандартное расположение: /cache
		Cache string
		// Lib - путь к директории для хранения системных библиотек и хранилищ системы.
		// Стандартное расположение: /lib
		Lib string
		// Run - путь к директории PID файлов.
		// Стандартное расположение: /run
		Run string
	}
}

// BuildOptions - опции построения хранилища.
type BuildOptions struct {
	ID string
}

// Build - построение хранилища.
func (storage *Dev) Build(options BuildOptions) *Dev {
	storage.SystemLocation = ""

	storage.Bin = "/bin"
	storage.Etc = "/etc"

	// Src
	{
		storage.Src = &struct {
			Path string
		}{
			Path: "/src",
		}
	}

	// System
	{
		storage.System = &struct {
			Path string
		}{
			Path: "/system",
		}
	}

	// Var
	{
		storage.Var = &struct {
			Path string

			Temp  string
			Logs  string
			Data  string
			Cache string
			Lib   string
			Run   string

			Test *struct {
				Path string

				Data  string
				Cache string
			}
		}{
			Path: "/var",
		}

		var s = storage.Var

		s.Temp = path.Join(s.Path, "/temp")
		s.Logs = path.Join(s.Path, "/logs")
		s.Data = path.Join(s.Path, "/data")
		s.Cache = path.Join(s.Path, "/cache")
		s.Lib = path.Join(s.Path, "/lib")
		s.Run = path.Join(s.Path, "/run")

		// Test
		{
			storage.Var.Test = &struct {
				Path string

				Data  string
				Cache string
			}{
				Path: path.Join(storage.Var.Path, "/test"),
			}

			var s = storage.Var.Test

			s.Data = path.Join(s.Path, "/data")
			s.Cache = path.Join(s.Path, "/cache")
		}
	}

	// Test
	{
		if testing.Testing() {
			var p = path.Join(storage.Var.Test.Cache, options.ID)

			storage.Etc = path.Join(p, storage.Etc)

			storage.Var.Path = path.Join(p, storage.Var.Path)
			storage.Var.Temp = path.Join(p, storage.Var.Temp)
			storage.Var.Logs = path.Join(p, storage.Var.Logs)
			storage.Var.Cache = path.Join(p, storage.Var.Cache)
			storage.Var.Run = path.Join(p, storage.Var.Run)
		}
	}

	return storage
}

// Build - построение хранилища.
func (storage *Prod) Build(options BuildOptions) *Prod {
	storage.SystemLocation = ""

	storage.Bin = "/bin"
	storage.Etc = "/etc"

	// Src
	{
		storage.Src = &struct {
			Path string
		}{
			Path: "/src",
		}
	}

	// System
	{
		storage.System = &struct {
			Path string
		}{
			Path: "/system",
		}
	}

	// Var
	{
		storage.Var = &struct {
			Path string

			Temp  string
			Logs  string
			Data  string
			Cache string
			Lib   string
			Run   string
		}{
			Path: "/var",
		}

		var s = storage.Var

		s.Temp = path.Join(s.Path, "/temp")
		s.Logs = path.Join(s.Path, "/logs")
		s.Data = path.Join(s.Path, "/data")
		s.Cache = path.Join(s.Path, "/cache")
		s.Lib = path.Join(s.Path, "/lib")
		s.Run = path.Join(s.Path, "/run")
	}

	return storage
}
