package models

import common_types "sm-box/internal/common/types"

type (
	// UserInfo - внешняя модель пользователя системы.
	UserInfo struct {
		ID common_types.ID `json:"id" xml:"id,attr"`

		Email    string `json:"email"    xml:"Email"`
		Username string `json:"username" xml:"Username"`

		Accesses *UserInfoAccesses `json:"accesses" xml:"Accesses>Access"`
	}

	// UserInfoAccesses - внешняя модель с информацией о доступах пользователя.
	UserInfoAccesses struct {
		Roles       []*RoleInfo       `json:"roles"       xml:"Roles"`
		Permissions []*PermissionInfo `json:"permissions" xml:"Permissions"`
	}
)

// RolesIDs - получение списка ID ролей.
func (accesses UserInfoAccesses) RolesIDs() (list []common_types.ID) {
	var writeInheritance func(rl *RoleInfo)

	writeInheritance = func(rl *RoleInfo) {
		list = append(list, rl.ID)

		for _, inheritRl := range rl.Inheritances {
			writeInheritance(inheritRl.RoleInfo)
		}
	}

	list = make([]common_types.ID, 0)

	for _, rl := range accesses.Roles {
		writeInheritance(rl)
	}

	return
}

// PermissionsIDs - получение списка ID прав.
func (accesses UserInfoAccesses) PermissionsIDs() (list []common_types.ID) {
	var writeInheritance func(rl *RoleInfo)

	writeInheritance = func(rl *RoleInfo) {
		for _, permission := range rl.Permissions {
			list = append(list, permission.ID)
		}

		for _, inheritRl := range rl.Inheritances {
			writeInheritance(inheritRl.RoleInfo)
		}
	}

	list = make([]common_types.ID, 0)

	for _, rl := range accesses.Roles {
		writeInheritance(rl)
	}

	return
}
