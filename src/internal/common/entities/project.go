package entities

import (
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/tracer"
)

type (
	// Project - проект.
	Project struct {
		ID types.ID

		Title       string
		Description string

		Owners ProjectOwners
	}

	// ProjectOwners - владельцы проекта.
	ProjectOwners []*ProjectOwner

	// ProjectOwner - владелец проекта.
	ProjectOwner struct {
		*User
	}
)

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *Project) FillEmptyFields() *Project {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(entity) }()
	}

	if entity.Owners == nil {
		entity.Owners = make(ProjectOwners, 0)
	}

	return entity
}
