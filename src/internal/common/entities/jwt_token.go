package entities

import (
	"sm-box/internal/common/types"
	"time"
)

type (
	// JwtToken - jwt токен системы доступа.
	JwtToken struct {
		ID     types.ID
		UserID types.ID

		Data string

		CreatedAt time.Time
		ExpiredAt time.Time

		Params *JwtTokenParams
	}

	// JwtTokenParams - параметры jwt токена системы доступа.
	JwtTokenParams struct {
		RemoteAddr string
		UserAgent  string
	}
)

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *JwtToken) FillEmptyFields() *JwtToken {
	if entity.Params == nil {
		entity.Params = new(JwtTokenParams)
	}

	return entity
}
