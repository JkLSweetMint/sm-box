package models

import "sm-box/internal/common/types"

type (
	// RoleInfo - модель с информацией о роли пользователя в системе.
	RoleInfo struct {
		ID        types.ID `json:"id"         xml:"id,attr"`
		ProjectID types.ID `json:"project_id" xml:"project_id,attr"`

		Name         string               `json:"name"                   xml:"Name"`
		Inheritances RoleInfoInheritances `json:"inheritances,omitempty" xml:"Inheritances,omitempty>Inheritance"`
	}

	// RoleInfoInheritances - модель с информацией о наследованиях роли.
	RoleInfoInheritances []*RoleInfoInheritance

	// RoleInfoInheritance - модель с информацией о наследовании роли.
	RoleInfoInheritance struct {
		*RoleInfo
	}
)
