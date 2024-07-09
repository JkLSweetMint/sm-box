package models

import "sm-box/internal/common/types"

type (
	// UserInfo - пользователь системы.
	UserInfo struct {
		ID        types.ID `json:"id"         xml:"id,attr"`
		ProjectID types.ID `json:"project_id" xml:"project_id,attr"`

		Email    string `json:"email"    xml:"Email"`
		Username string `json:"username" xml:"Username"`

		Password string `json:"password,omitempty" xml:"Password,omitempty"`

		Accesses UserInfoAccesses `json:"accesses" xml:"Accesses>Access"`
	}

	// UserInfoAccesses - информация о доступах пользователя.
	UserInfoAccesses []*UserInfoAccess

	// UserInfoAccess - информация о доступе пользователя.
	UserInfoAccess struct {
		*RoleInfo
	}
)
