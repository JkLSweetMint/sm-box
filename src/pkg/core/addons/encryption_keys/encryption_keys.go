package encryption_keys

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"path"
	"sm-box/src/pkg/core/env"
)

const (
	prtFileName = "private.pem"
	pbFileName  = "public.pem"
)

// Init - инициализация ключей шифрования.
// По-умолчанию используется rsa 4096Bytes.
func Init() (err error) {
	var (
		dir = path.Join(env.Paths.SystemLocation, env.Paths.Src.Embed)
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
			exists = [2]bool{
				true,
				true,
			}
			files = [2]string{
				path.Join(dir, prtFileName),
				path.Join(dir, pbFileName),
			}
		)

		for i, p := range files {
			if _, err = os.Stat(p); errors.Is(err, os.ErrNotExist) {
				err = nil
				exists[i] = false
			} else if err != nil {
				return
			}
		}

		if exists[0] && exists[1] {
			if prt, pb, err = read(dir); err != nil {
				return
			}

			return
		} else if exists[0] != exists[1] {
			err = ErrInvalidEncryptionKeys
			return
		}
	}

	// Создание ключей
	{
		if prt, err = rsa.GenerateKey(rand.Reader, 4096); err != nil {
			return
		}

		pb = &prt.PublicKey
	}

	if err = write(dir, prt, pb); err != nil {
		return
	}

	return
}

// write - запись ключей шифрования.
func write(dir string, prt *rsa.PrivateKey, pb *rsa.PublicKey) (err error) {
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

	return err
}

// read - чтение ключей шифрования.
func read(dir string) (prt *rsa.PrivateKey, pb *rsa.PublicKey, err error) {
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

	return
}
