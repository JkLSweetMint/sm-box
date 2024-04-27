package entities

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

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *User) FillEmptyFields() *User {
	if entity.Accesses == nil {
		entity.Accesses = make(UserAccesses, 0)
	}

	return entity
}
