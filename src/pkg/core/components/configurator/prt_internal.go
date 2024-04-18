package configurator

import (
	"os"
	"path"
	"sm-box/src/pkg/core/components/configurator/encoders"
	"sm-box/src/pkg/core/env"
	"strings"
)

var (
	prtDir            = path.Join(env.Paths.SystemLocation, env.Paths.System.Path)
	prtDefaultEncoder = encoders.XmlEncoder{}
)

// privateConfigurator - внутренняя реализация диспетчера приватных конфигураций.
// privateConfigurator - внутренняя реализация диспетчера системных конфигураций.
type privateConfigurator[T any] struct {
	conf          Config[T]
	encoder       Encoder
	dir, filename string
}

// Encoder - задать кодировщик конфигурации.
//
// По умолчанию стоит encoders.XmlEncoder.
func (c *privateConfigurator[T]) Encoder(encoder Encoder) Private[T] {
	if encoder != nil {
		c.encoder = encoder
	}

	return c
}

// File - задать файл для взаимодействия с конфигурацией.
func (c *privateConfigurator[T]) File(dir, filename string) Private[T] {
	c.dir = strings.TrimSpace(dir)
	c.filename = strings.TrimSpace(filename)

	return c
}

// Profile - установить профиль конфигурации.
func (c *privateConfigurator[T]) Profile(profile PrivateProfile) Private[T] {
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

// Read - чтение конфигурации.
func (c *privateConfigurator[T]) Read() (err error) {
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

	if data, err = os.ReadFile(path.Join(prtDir, c.dir, c.filename)); err != nil {
		return
	}

	if err = c.encoder.Decode(data, c.conf); err != nil {
		return
	}

	return
}
