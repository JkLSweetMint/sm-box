package entities

import "sm-box/internal/common/types"

type (
	// Role - роль пользователя в системе.
	Role struct {
		ID        types.ID
		ProjectID types.ID

		Title        string
		Inheritances RoleInheritances
	}

	// RoleInheritances - наследования роли.
	RoleInheritances []*RoleInheritance

	// RoleInheritance - наследование роли.
	RoleInheritance struct {
		*Role
	}
)

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *Role) FillEmptyFields() *Role {
	if entity.Inheritances == nil {
		entity.Inheritances = make(RoleInheritances, 0)
	}

	return entity
}
