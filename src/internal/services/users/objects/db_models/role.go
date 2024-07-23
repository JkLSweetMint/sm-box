package db_models

import "sm-box/internal/common/types"

type (
	// Role - модель базы данных роли пользователя.
	Role struct {
		ID        types.ID `db:"id"`
		ProjectID types.ID `db:"project_id"`

		Name     string `db:"name"`
		IsSystem bool   `db:"is_system"`
	}

	// RoleInheritance - модель базы данных наследования роли.
	RoleInheritance struct {
		Parent types.ID `db:"parent"`
		Heir   types.ID `db:"heir"`
	}
)
