package entities

import (
	"sm-box/internal/common/types"
	"sm-box/internal/services/users/objects/models"
	"sm-box/pkg/core/components/tracer"
)

type (
	// User - пользователь системы.
	User struct {
		ID types.ID

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
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(entity) }()
	}

	if entity.Accesses == nil {
		entity.Accesses = make(UserAccesses, 0)
	}

	return entity
}

// ToModel - получение внешней модели.
func (entity *User) ToModel() (model *models.UserInfo) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &models.UserInfo{
		ID: entity.ID,

		Email:    entity.Email,
		Username: entity.Username,

		Accesses: make(models.UserInfoAccesses, 0),
	}

	for _, acc := range entity.Accesses {
		model.Accesses = append(model.Accesses, acc.ToModel())
	}

	return
}

// ToModel - получение внешней модели.
func (entity *UserAccess) ToModel() (model *models.UserInfoAccess) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &models.UserInfoAccess{
		RoleInfo: entity.Role.ToModel(),
	}

	return
}
