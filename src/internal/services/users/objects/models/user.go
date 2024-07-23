package models

import (
	"sm-box/internal/common/types"
)

type (
	// UserInfo - внешняя модель пользователя системы.
	UserInfo struct {
		ID types.ID `json:"id" xml:"id,attr"`

		Email    string `json:"email"    xml:"Email"`
		Username string `json:"username" xml:"Username"`

		Accesses UserInfoAccesses `json:"accesses" xml:"Accesses>Access"`
	}

	// UserInfoAccesses - внешняя модель с информацией о доступах пользователя.
	UserInfoAccesses []*UserInfoAccess

	// UserInfoAccess - внешняя модель с информацией о доступе пользователя.
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
