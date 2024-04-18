package configurator

import "errors"

var (
	ErrNilConfigurationInstanceWasPassed     = errors.New("A nil configuration instance was passed. ")
	ErrNilConfigurationInstanceIsSpecified   = errors.New("A nil configuration instance is specified. ")
	ErrNilConfigurationEncoderIsSpecified    = errors.New("A nil configuration encoder is specified. ")
	ErrEmptyConfigurationFilenameIsSpecified = errors.New("A empty configuration filename is specified. ")
)
