package models

import "sm-box/internal/common/types"

type (
	// UserInfo - пользователь системы.
	UserInfo struct {
		ID        types.ID `json:"id"         yaml:"ID"        xml:"id,attr"`
		ProjectID types.ID `json:"project_id" yaml:"ProjectID" xml:"project_id,attr"`

		Email    string `json:"email"    yaml:"Email"    xml:"Email"`
		Username string `json:"username" yaml:"Username" xml:"Username"`
		Password string `json:"password" yaml:"Password" xml:"Password"`

		Accesses UserInfoAccesses `json:"accesses" yaml:"Accesses" xml:"Accesses>Access"`
	}

	// UserInfoAccesses - информация о доступах пользователя.
	UserInfoAccesses []*UserInfoAccess

	// UserInfoAccess - информация о доступе пользователя.
	UserInfoAccess struct {
		*RoleInfo
	}
)
