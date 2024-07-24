package db_models

import (
	"encoding/json"
	users_models "sm-box/internal/services/users/objects/models"
)

type (
	// JwtAccessToken - модель база данных jwt токена доступа.
	JwtAccessToken struct {
		*JwtToken

		UserInfo *JwtAccessTokenUserInfo `json:"user_info"`
	}

	// JwtAccessTokenUserInfo - модель база данных с информацией о пользователя для jwt токена доступа.
	JwtAccessTokenUserInfo struct {
		Accesses *users_models.UserInfoAccesses `json:"accesses"`
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
