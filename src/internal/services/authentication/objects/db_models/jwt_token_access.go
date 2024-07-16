package db_models

import (
	"encoding/json"
	"sm-box/internal/common/types"
)

type (
	// JwtAccessToken - jwt токен доступа.
	JwtAccessToken struct {
		*JwtToken

		UserInfo *JwtAccessTokenUserInfo `json:"user_info"`
	}

	// JwtAccessTokenUserInfo - информация о пользователя для jwt токена доступа.
	JwtAccessTokenUserInfo struct {
		Accesses []types.ID `json:"accesses"`
	}
)

// MarshalBinary - упаковка структуры в бинарный формат.
func (entity *JwtAccessToken) MarshalBinary() ([]byte, error) {
	return json.Marshal(entity)
}

// UnmarshalBinary - распаковка структуры из бинарного формата.
func (entity *JwtAccessToken) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &entity)
}
