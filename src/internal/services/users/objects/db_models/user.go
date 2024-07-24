package db_models

import common_types "sm-box/internal/common/types"

type (
	// User - модель базы данных пользователя.
	User struct {
		ID common_types.ID `db:"id"`

		Email    string `db:"email"`
		Username string `db:"username"`

		Password string `db:"password"`
	}

	// UserAccess - модель базы данных доступа пользователя.
	UserAccess struct {
		UserID       common_types.ID `db:"user_id"`
		RoleID       common_types.ID `db:"role_id"`
		PermissionID common_types.ID `db:"permission_id"`
	}
)
