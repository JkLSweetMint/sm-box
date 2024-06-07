package helpers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type (
	// Serialization - описание методов сериализации ошибок.
	Serialization interface {
		json.Marshaler
		xml.Marshaler

		json.Unmarshaler
	}

	// Stringer - описание методов для преобразование в строку.
	Stringer interface {
		fmt.Stringer
	}

	// Error - описание методов для связи с builtin ошибкой.
	Error interface {
		error
	}
)
