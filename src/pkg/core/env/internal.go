package env

import (
	"os"
	"path"
	"reflect"
	env_mode "sm-box/pkg/core/env/mode"
	"strings"
	"testing"
)

// initSystemDir - инициализация системных директорий.
func initSystemDir() (err error) {
	var getPaths func(value reflect.Value) (list []string)

	// getPaths
	{
		getPaths = func(value reflect.Value) (list []string) {
			list = make([]string, 0)

			for i := 0; i < value.NumField(); i++ {
				var f = value.Field(i)

				if value.Type().Field(i).Name == "SystemLocation" {
					continue
				}

				switch f.Kind().String() {
				case reflect.String.String():
					{
						list = append(list, f.Interface().(string))
					}
				case reflect.Ptr.String():
					{
						list = append(list, getPaths(f.Elem())...)
					}
				case reflect.Struct.String():
					{
						list = append(list, getPaths(f)...)
					}
				}
			}

			return
		}
	}

	var list = getPaths(reflect.ValueOf(Paths).Elem())

	for _, p := range list {
		p = path.Join(Paths.SystemLocation, p)

		if err = os.MkdirAll(p, 0755); err != nil {
			return
		}
	}

	return
}

// getSystemLocation - получение местоположения системы.
func getSystemLocation() (location string, err error) {
	defer func() {
		if location == "" && err == nil {
			err = ErrSystemLocationNotFound
		}
	}()

	// Test mode
	{
		if testing.Testing() || Mode == env_mode.Dev {
			location = devSystemLocation

			return
		}
	}

	// Остальное
	{
		if location, err = os.Getwd(); err != nil {
			return
		}

		location = strings.Replace(location, "\\", "/", -1)

		switch {
		case strings.HasSuffix(location, "/bin/"):
			{
				location = location[:len(location)-4]
			}
		case strings.HasSuffix(location, "/bin"):
			{
				location = location[:len(location)-3]
			}
		}
	}

	return
}
