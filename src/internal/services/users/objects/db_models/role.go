package db_models

import (
	"github.com/google/uuid"
	common_types "sm-box/internal/common/types"
)

type (
	// Role - модель базы данных роли пользователя.
	Role struct {
		ID        common_types.ID `db:"id"`
		ProjectID common_types.ID `db:"project_id"`

		Name     string    `db:"name"`
		NameI18n uuid.UUID `db:"name_i18n"`

		Description     string    `db:"description"`
		DescriptionI18n uuid.UUID `db:"description_i18n"`

		IsSystem bool `db:"is_system"`
	}

	// RoleInheritance - модель базы данных наследования роли.
	RoleInheritance struct {
		ParentID common_types.ID `db:"parent_id"`
		HeirID   common_types.ID `db:"heir_id"`
	}
)
