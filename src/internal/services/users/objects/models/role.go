package models

import (
	"github.com/google/uuid"
	common_types "sm-box/internal/common/types"
)

type (
	// RoleInfo - внешняя модель с информацией о роли пользователя в системе.
	RoleInfo struct {
		ID        common_types.ID `json:"id"         xml:"id,attr"`
		ProjectID common_types.ID `json:"project_id" xml:"project_id,attr"`

		Name     string    `json:"name"      xml:"Name"`
		NameI18n uuid.UUID `json:"name_i18n" xml:"Name>i18n,attr"`

		Description     string    `json:"description"      xml:"Description"`
		DescriptionI18n uuid.UUID `json:"description_i18n" xml:"Description>i18n,attr"`

		IsSystem bool `json:"is_system" xml:"is_system,attr"`

		Permissions  []*PermissionInfo    `json:"permissions,omitempty"  xml:"Permissions>Permission,omitempty"`
		Inheritances RoleInfoInheritances `json:"inheritances,omitempty" xml:"Inheritances>Inheritance,omitempty"`
	}

	// RoleInfoInheritances - внешняя модель с информацией о наследованиях роли.
	RoleInfoInheritances []*RoleInfoInheritance

	// RoleInfoInheritance - внешняя модель с информацией о наследовании роли.
	RoleInfoInheritance struct {
		*RoleInfo
	}
)
