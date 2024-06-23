package entities

import (
	g_uuid "github.com/google/uuid"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/tracer"
)

type (
	// Project - проект.
	Project struct {
		ID   types.ID
		UUID g_uuid.UUID

		Name        string
		Description string
		Version     string

		Owner *ProjectOwner
	}

	// ProjectOwner - владелец проекта.
	ProjectOwner struct {
		*User
	}

	// ProjectEnvVar - переменная окружения проекта
	ProjectEnvVar struct {
		Key   string
		Value string
	}

	// ProjectEnv - переменные окружения проекта
	ProjectEnv []*ProjectEnvVar
)

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *Project) FillEmptyFields() *Project {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(entity) }()
	}

	if entity.Owner == nil {
		entity.Owner = &ProjectOwner{
			User: new(User),
		}

		entity.Owner.User.FillEmptyFields()
	}

	return entity
}
