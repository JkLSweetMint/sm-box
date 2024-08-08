package db_models

import (
	"encoding/json"
	common_types "sm-box/internal/common/types"
)

type (
	// JwtAccessToken - модель база данных jwt токена доступа.
	JwtAccessToken struct {
		*JwtToken

		UserInfo *JwtAccessTokenUserInfo `json:"user_info"`
	}

	// JwtAccessTokenUserInfo - модель база данных с информацией о пользователя для jwt токена доступа.
	JwtAccessTokenUserInfo struct {
		Accesses *JwtAccessTokenUserInfoAccesses `json:"accesses"`
	}

	// JwtAccessTokenUserInfoAccesses - информация о доступах пользователя для jwt токена доступа.
	JwtAccessTokenUserInfoAccesses struct {
		Roles       []common_types.ID
		Permissions []common_types.ID
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
