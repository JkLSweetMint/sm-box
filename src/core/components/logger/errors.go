package logger

import "errors"

var (
	ErrInstanceNoFound           = errors.New("The instance component was not found. ")
	ErrComponentCouldNotBeCopied = errors.New("The component could not be copied. ")
)
