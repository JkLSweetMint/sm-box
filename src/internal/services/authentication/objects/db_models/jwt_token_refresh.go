package db_models

import "encoding/json"

type (
	// JwtRefreshToken - модель база данных jwt токена обновления.
	JwtRefreshToken struct {
		*JwtToken
	}
)

// MarshalBinary - упаковка структуры в бинарный формат.
func (entity *JwtRefreshToken) MarshalBinary() ([]byte, error) {
	return json.Marshal(entity)
}

// UnmarshalBinary - распаковка структуры из бинарного формата.
func (entity *JwtRefreshToken) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &entity)
}
