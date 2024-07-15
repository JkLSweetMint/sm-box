package models

import (
	"sm-box/internal/common/types"
)

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

// ListIDs - получение списка ID доступов.
func (accesses UserInfoAccesses) ListIDs() (list []types.ID) {
	var writeInheritance func(rl *RoleInfo)

	writeInheritance = func(rl *RoleInfo) {
		list = append(list, rl.ID)

		for _, inheritRl := range rl.Inheritances {
			writeInheritance(inheritRl.RoleInfo)
		}
	}

	for _, rl := range accesses {
		writeInheritance(rl.RoleInfo)
	}

	return
}
