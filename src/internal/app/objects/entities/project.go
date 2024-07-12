package entities

import (
	"sm-box/internal/app/objects/models"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/tracer"
)

type (
	// Project - проект.
	Project struct {
		ID types.ID

		Name        string
		Description string
		Version     string
	}

	// ProjectList - список проектов.
	ProjectList []*Project
)

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *Project) FillEmptyFields() *Project {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(entity) }()
	}

	return entity
}

// ToModel - получение модели.
func (entity ProjectList) ToModel() (list models.ProjectList) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(list) }()
	}

	list = make(models.ProjectList, len(entity))

	for i, item := range entity {
		list[i] = item.ToModel()
	}

	return
}

// ToModel - получение модели.
func (entity *Project) ToModel() (model *models.ProjectInfo) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &models.ProjectInfo{
		ID: entity.ID,

		Name:        entity.Name,
		Description: entity.Description,
		Version:     entity.Version,
	}

	return
}
