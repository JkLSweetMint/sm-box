package models

import "sm-box/internal/common/types"

type (
	// UserInfo - пользователь системы.
	UserInfo struct {
		ID types.ID `json:"id"         xml:"id,attr"`

		Email    string `json:"email"    xml:"Email"`
		Username string `json:"username" xml:"Username"`

		Accesses UserInfoAccesses `json:"accesses" xml:"Accesses>Access"`
	}

	// UserInfoAccesses - информация о доступах пользователя.
	UserInfoAccesses []*UserInfoAccess

	// UserInfoAccess - информация о доступе пользователя.
	UserInfoAccess struct {
		*RoleInfo
	}
)
