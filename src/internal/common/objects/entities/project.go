package common_entities

import (
	common_models "sm-box/internal/common/objects/models"
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

		Owner *ProjectOwner
	}

	// ProjectOwner - владелец проекта.
	ProjectOwner struct {
		*User
	}

	ProjectList []*ProjectListItem

	ProjectListItem struct {
		ID      types.ID
		Name    string
		Version string
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

// ToModel - получение модели.
func (entity ProjectList) ToModel() (list common_models.ProjectList) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(list) }()
	}

	list = make(common_models.ProjectList, len(entity))

	for i, item := range entity {
		list[i] = item.ToModel()
	}

	return
}

// ToModel - получение модели.
func (entity *ProjectListItem) ToModel() (model *common_models.ProjectListItem) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &common_models.ProjectListItem{
		ID:      entity.ID,
		Name:    entity.Name,
		Version: entity.Version,
	}

	return
}
