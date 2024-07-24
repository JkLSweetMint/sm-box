package entities

import (
	"github.com/google/uuid"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/users/objects/models"
	"sm-box/pkg/core/components/tracer"
)

type (
	// Role - роль пользователя в системе.
	Role struct {
		ID        common_types.ID
		ProjectID common_types.ID

		Name     string
		NameI18n uuid.UUID

		Description     string
		DescriptionI18n uuid.UUID

		IsSystem bool

		Permissions  []*Permission
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

	if entity.Permissions == nil {
		entity.Permissions = make([]*Permission, 0)
	}

	for _, permission := range entity.Permissions {
		if permission != nil {
			permission.FillEmptyFields()
		}
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

	entity.FillEmptyFields()

	model = &models.RoleInfo{
		ID:        entity.ID,
		ProjectID: entity.ProjectID,

		Name:            entity.Name,
		NameI18n:        entity.NameI18n,
		Description:     entity.Description,
		DescriptionI18n: entity.DescriptionI18n,

		IsSystem: entity.IsSystem,

		Permissions:  make([]*models.PermissionInfo, 0),
		Inheritances: make(models.RoleInfoInheritances, 0),
	}

	for _, rl := range entity.Inheritances {
		if rl != nil {
			model.Inheritances = append(model.Inheritances, rl.ToModel())
		}
	}

	for _, permission := range entity.Permissions {
		if permission != nil {
			model.Permissions = append(model.Permissions, permission.ToModel())
		}
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

	entity.FillEmptyFields()

	model = &models.RoleInfoInheritance{
		RoleInfo: entity.Role.ToModel(),
	}

	return
}
