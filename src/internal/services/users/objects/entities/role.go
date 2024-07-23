package entities

import (
	"sm-box/internal/common/types"
	"sm-box/internal/services/users/objects/models"
	"sm-box/pkg/core/components/tracer"
)

type (
	// Role - роль пользователя в системе.
	Role struct {
		ID        types.ID
		ProjectID types.ID

		Name     string
		IsSystem bool

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

// ToModel - получение внешней модели.
func (entity *Role) ToModel() (model *models.RoleInfo) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &models.RoleInfo{
		ID:           entity.ID,
		ProjectID:    entity.ProjectID,
		Name:         entity.Name,
		Inheritances: make(models.RoleInfoInheritances, 0),
	}

	for _, rl := range entity.Inheritances {
		model.Inheritances = append(model.Inheritances, rl.ToModel())
	}

	return
}

// ToModel - получение внешней модели.
func (entity *RoleInheritance) ToModel() (model *models.RoleInfoInheritance) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &models.RoleInfoInheritance{
		RoleInfo: entity.Role.ToModel(),
	}

	return
}
