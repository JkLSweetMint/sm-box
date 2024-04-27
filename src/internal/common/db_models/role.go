package db_models

import "sm-box/internal/common/types"

type (
	// Role - модель роли пользователя для базы данных.
	Role struct {
		ID        types.ID `db:"id"`
		ProjectID types.ID `db:"project_id"`

		Title string `db:"title"`
	}

	// RoleInheritance - модель наследования роли для базы данных.
	RoleInheritance struct {
		Parent types.ID `db:"parent"`
		Heir   types.ID `db:"heir"`
	}
)
