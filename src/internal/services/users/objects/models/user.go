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
func (accesses UserInfoAccesses) RolesIDs() (roles []common_types.ID) {
	var (
		writeInheritance func(rl *RoleInfo)
		list             = make(map[common_types.ID]struct{})
	)

	roles = make([]common_types.ID, 0)

	writeInheritance = func(rl *RoleInfo) {
		list[rl.ID] = struct{}{}

		for _, inheritRl := range rl.Inheritances {
			writeInheritance(inheritRl.RoleInfo)
		}
	}

	for _, rl := range accesses.Roles {
		writeInheritance(rl)
	}

	for id, _ := range list {
		roles = append(roles, id)
	}

	return
}

// PermissionsIDs - получение списка ID прав.
func (accesses UserInfoAccesses) PermissionsIDs() (permissions []common_types.ID) {
	var (
		writeInheritance func(rl *RoleInfo)
		list             = make(map[common_types.ID]struct{})
	)

	permissions = make([]common_types.ID, 0)

	writeInheritance = func(rl *RoleInfo) {
		for _, permission := range rl.Permissions {
			list[permission.ID] = struct{}{}
		}

		for _, inheritRl := range rl.Inheritances {
			writeInheritance(inheritRl.RoleInfo)
		}
	}

	for _, rl := range accesses.Roles {
		writeInheritance(rl)
	}

	for id, _ := range list {
		permissions = append(permissions, id)
	}

	return
}
