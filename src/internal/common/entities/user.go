package models

import "sm-box/internal/common/types"

type (
	// User - пользователь системы.
	User struct {
		ID        types.ID
		ProjectID types.ID

		Email    string
		Username string
		Password string

		Accesses UserAccesses
	}

	// UserAccesses - доступы пользователя.
	UserAccesses []*UserAccess

	// UserAccess - доступ пользователя.
	UserAccess struct {
		*Role
	}
)
