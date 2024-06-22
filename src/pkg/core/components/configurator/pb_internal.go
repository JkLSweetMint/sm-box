package configurator

import (
	"errors"
	"os"
	"path"
	"sm-box/pkg/core/components/configurator/encoders"
	"sm-box/pkg/core/env"
	"strings"
)

var (
	PbDir            = path.Join(env.Paths.SystemLocation, env.Paths.Etc)
	pbDefaultEncoder = encoders.YamlEncoder{}
)

// publicConfigurator - внутренняя реализация диспетчера публичных конфигураций.
type publicConfigurator[T any] struct {
	conf          Config[T]
	encoder       Encoder
	dir, filename string
}

// Encoder - задать кодировщик конфигурации.
//
// По умолчанию стоит encoders.YamlEncoder.
func (c *publicConfigurator[T]) Encoder(encoder Encoder) Public[T] {
	if encoder != nil {
		c.encoder = encoder
	}

	return c
}

// File - задать файл для взаимодействия с конфигурацией.
func (c *publicConfigurator[T]) File(dir, filename string) Public[T] {
	c.dir = strings.TrimSpace(dir)
	c.filename = strings.TrimSpace(filename)

	return c
}

// Profile - установить профиль конфигурации.
func (c *publicConfigurator[T]) Profile(profile PublicProfile) Public[T] {
	if profile.Dir = strings.TrimSpace(profile.Dir); profile.Dir != "" {
		c.dir = strings.TrimSpace(profile.Dir)
	}

	if profile.Filename = strings.TrimSpace(profile.Filename); profile.Filename != "" {
		c.filename = strings.TrimSpace(profile.Filename)
	}

	if profile.Encoder != nil {
		c.encoder = profile.Encoder
	}

	return c
}

// Init - инициализация конфигурации.
// Если файл конфигурации существует, читает, иначе создает.
func (c *publicConfigurator[T]) Init() (err error) {
	// Проверки
	{
		switch {
		case c.conf == nil:
			return ErrNilConfigurationInstanceIsSpecified
		case c.encoder == nil:
			return ErrNilConfigurationEncoderIsSpecified
		case c.filename == "":
			return ErrEmptyConfigurationFilenameIsSpecified
		}
	}

	if _, err = os.Stat(path.Join(PbDir, c.dir, c.filename)); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			c.conf.Default()
			return c.Write()
		}

		return
	}

	return c.Read()
}

// Write - запись конфигурации.
func (c *publicConfigurator[T]) Write() (err error) {
	// Проверки
	{
		switch {
		case c.conf == nil:
			return ErrNilConfigurationInstanceIsSpecified
		case c.encoder == nil:
			return ErrNilConfigurationEncoderIsSpecified
		case c.filename == "":
			return ErrEmptyConfigurationFilenameIsSpecified
		}
	}

	var data []byte

	if data, err = c.encoder.Encode(c.conf); err != nil {
		return
	}

	var dir = path.Join(PbDir, c.dir)

	if err = os.MkdirAll(dir, 0655); err != nil {
		return
	}

	if err = os.WriteFile(path.Join(dir, c.filename), data, 0655); err != nil {
		return
	}

	return
}

// Read - чтение конфигурации.
func (c *publicConfigurator[T]) Read() (err error) {
	// Проверки
	{
		switch {
		case c.conf == nil:
			return ErrNilConfigurationInstanceIsSpecified
		case c.encoder == nil:
			return ErrNilConfigurationEncoderIsSpecified
		case c.filename == "":
			return ErrEmptyConfigurationFilenameIsSpecified
		}
	}

	var data []byte

	if data, err = os.ReadFile(path.Join(PbDir, c.dir, c.filename)); err != nil {
		return
	}

	if err = c.encoder.Decode(data, c.conf); err != nil {
		return
	}

	return
}
