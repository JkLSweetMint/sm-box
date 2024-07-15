package db_models

import "sm-box/internal/common/types"

type (
	// User - модель пользователя для базы данных.
	User struct {
		ID types.ID `db:"id"`

		Email    string `db:"email"`
		Username string `db:"username"`

		Password string `db:"password"`
	}

	// UserAccess - модель доступа пользователя для базы данных.
	UserAccess struct {
		UserID types.ID `db:"user_id"`
		RoleID types.ID `db:"role_id"`
	}
)
