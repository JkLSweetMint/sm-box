package entities

import (
	"sm-box/internal/app/infrastructure/types"
	"sm-box/pkg/core/components/tracer"
)

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
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(entity) }()
	}

	if entity.Inheritances == nil {
		entity.Inheritances = make(RoleInheritances, 0)
	}

	return entity
}
