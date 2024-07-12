package helpers

import (
	"fmt"
)

type (
	// Stringer - описание методов для преобразование в строку.
	Stringer interface {
		fmt.Stringer
	}

	// Error - описание методов для связи с builtin ошибкой.
	Error interface {
		error
	}
)
