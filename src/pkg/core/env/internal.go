package env

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/fs"
	"os"
	"path"
	"reflect"
	"slices"
	"sm-box/embed"
	"strings"
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

// readEncryptionKeys - чтение ключей шифрования.
func readEncryptionKeys() (err error) {
	const (
		prtFileName = "private.pem"
		pbFileName  = "public.pem"
	)

	var (
		exists = []bool{
			false,
			false,
		}
		files = []string{
			prtFileName,
			pbFileName,
		}
	)

	// Проверка существования
	{
		var dir []fs.DirEntry

		if dir, err = embed.Dir.ReadDir("."); err != nil {
			return
		}

		for _, fl := range dir {
			if index := slices.Index(files, fl.Name()); index >= 0 {
				exists[index] = true
			}
		}
	}

	// Чтение
	{
		if exists[0] && exists[1] {
			// Приватный ключ
			{
				var (
					data  []byte
					block *pem.Block
				)

				if data, err = embed.Dir.ReadFile(prtFileName); err != nil {
					return
				}

				block, _ = pem.Decode(data)

				if Vars.EncryptionKeys.Private, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
					return
				}
			}

			// Публичный ключ
			{
				var (
					data  []byte
					block *pem.Block
				)

				if data, err = embed.Dir.ReadFile(pbFileName); err != nil {
					return
				}

				block, _ = pem.Decode(data)

				if Vars.EncryptionKeys.Public, err = x509.ParsePKCS1PublicKey(block.Bytes); err != nil {
					return
				}
			}

			return
		}
	}

	// Ключей нет
	{
		err = errors.New("Invalid encryption keys. ")
	}

	return
}
