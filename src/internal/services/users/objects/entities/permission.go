package entities

import (
	"github.com/google/uuid"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/users/objects/models"
	"sm-box/pkg/core/components/tracer"
)

type (
	// Permission - права пользователя.
	Permission struct {
		ID        common_types.ID
		ProjectID common_types.ID

		Name     string
		NameI18n uuid.UUID

		Description     string
		DescriptionI18n uuid.UUID

		IsSystem bool
	}
)

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *Permission) FillEmptyFields() *Permission {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(entity) }()
	}

	return entity
}

// ToModel - получение внешней модели.
func (entity *Permission) ToModel() (model *models.PermissionInfo) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	entity.FillEmptyFields()

	model = &models.PermissionInfo{
		ID:        entity.ID,
		ProjectID: entity.ProjectID,

		Name:            entity.Name,
		NameI18n:        entity.NameI18n,
		Description:     entity.Description,
		DescriptionI18n: entity.DescriptionI18n,

		IsSystem: entity.IsSystem,
	}

	return
}
