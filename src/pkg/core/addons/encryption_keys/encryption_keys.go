package encryption_keys

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/fs"
	"os"
	"path"
	"slices"
	glob_embed "sm-box/embed"
	"sm-box/internal/system/init_cli/embed"
	"sm-box/pkg/core/env"
	env_mode "sm-box/pkg/core/env/mode"
)

const (
	prtFileName = "private.pem"
	pbFileName  = "public.pem"
)

// Init - инициализация ключей шифрования.
// По-умолчанию используется rsa 4096Bytes.
func Init() (err error) {
	var (
		prt *rsa.PrivateKey
		pb  *rsa.PublicKey
	)

	defer func() {
		if err == nil {
			env.Vars.EncryptionKeys.Private = prt
			env.Vars.EncryptionKeys.Public = pb
		}
	}()

	// Проверка существования
	{
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

		switch env.Mode {
		case env_mode.Dev:
			{
				var dir = path.Join(env.Paths.SystemLocation, env.Paths.Src.Embed)

				for i, p := range files {
					exists[i] = true

					if _, err = os.Stat(path.Join(dir, p)); errors.Is(err, os.ErrNotExist) {
						err = nil
						exists[i] = false
					} else if err != nil {
						return
					}
				}
			}
		case env_mode.Prod:
			{
				var dir []fs.DirEntry

				if dir, err = glob_embed.Dir.ReadDir("."); err != nil {
					return
				}

				for _, fl := range dir {
					if index := slices.Index(files, fl.Name()); index >= 0 {
						exists[index] = true
					}
				}
			}
		}

		if exists[0] && exists[1] {
			switch env.Mode {
			case env_mode.Dev:
				{
					var dir = path.Join(env.Paths.SystemLocation, env.Paths.Src.Embed)

					// Приватный ключ
					{
						var (
							data  []byte
							block *pem.Block
						)

						if data, err = os.ReadFile(path.Join(dir, prtFileName)); err != nil {
							return
						}

						block, _ = pem.Decode(data)

						if prt, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
							return
						}
					}

					// Публичный ключ
					{
						var (
							data  []byte
							block *pem.Block
						)

						if data, err = os.ReadFile(path.Join(dir, pbFileName)); err != nil {
							return
						}

						block, _ = pem.Decode(data)

						if pb, err = x509.ParsePKCS1PublicKey(block.Bytes); err != nil {
							return
						}
					}
				}
			case env_mode.Prod:
				{
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

						if prt, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
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

						if pb, err = x509.ParsePKCS1PublicKey(block.Bytes); err != nil {
							return
						}
					}
				}
			}

			return
		} else if exists[0] != exists[1] {
			err = ErrInvalidEncryptionKeys
			return
		}
	}

	if env.Mode != env_mode.Dev {
		err = ErrInvalidEncryptionKeys
		return
	}

	// Создание ключей
	{
		if prt, err = rsa.GenerateKey(rand.Reader, 4096); err != nil {
			return
		}

		pb = &prt.PublicKey
	}

	// Запись ключей
	{
		var dir = path.Join(env.Paths.SystemLocation, env.Paths.Src.Embed)

		// Приватный ключ
		{
			var block = &pem.Block{
				Type:  "RSA PRIVATE KEY",
				Bytes: x509.MarshalPKCS1PrivateKey(prt),
			}

			if err = os.WriteFile(path.Join(dir, prtFileName), pem.EncodeToMemory(block), 0644); err != nil {
				return
			}
		}

		// Публичный ключ
		{
			var block = &pem.Block{
				Type:  "PUBLIC KEY",
				Bytes: x509.MarshalPKCS1PublicKey(pb),
			}

			if err = os.WriteFile(path.Join(dir, pbFileName), pem.EncodeToMemory(block), 0644); err != nil {
				return
			}
		}
	}

	return
}
