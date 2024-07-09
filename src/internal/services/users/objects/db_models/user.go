package db_models

import "sm-box/internal/common/types"

type (
	// User - модель пользователя для базы данных.
	User struct {
		ID        types.ID `db:"id"`
		ProjectID types.ID `db:"project_id"`

		Email    string `db:"email"`
		Username string `db:"username"`
		Password []byte `db:"password"`
	}

	// UserAccess - модель доступа пользователя для базы данных.
	UserAccess struct {
		UserID types.ID `db:"user_id"`
		RoleID types.ID `db:"role_id"`
	}
)
