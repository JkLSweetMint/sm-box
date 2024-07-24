package entities

import (
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/users/objects/models"
	"sm-box/pkg/core/components/tracer"
)

type (
	// User - пользователь системы.
	User struct {
		ID common_types.ID

		Email    string
		Username string

		Password string

		Accesses *UserAccesses
	}

	// UserAccesses - доступы пользователя.
	UserAccesses struct {
		Roles       []*Role
		Permissions []*Permission
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
		entity.Accesses = new(UserAccesses)
	}

	entity.Accesses.FillEmptyFields()

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

	entity.FillEmptyFields()

	model = &models.UserInfo{
		ID: entity.ID,

		Email:    entity.Email,
		Username: entity.Username,

		Accesses: entity.Accesses.ToModel(),
	}

	return
}

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *UserAccesses) FillEmptyFields() *UserAccesses {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(entity) }()
	}

	if entity.Roles == nil {
		entity.Roles = make([]*Role, 0)
	}

	if entity.Permissions == nil {
		entity.Permissions = make([]*Permission, 0)
	}

	return entity
}

// ToModel - получение внешней модели.
func (entity *UserAccesses) ToModel() (model *models.UserInfoAccesses) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	entity.FillEmptyFields()

	model = &models.UserInfoAccesses{
		Roles:       make([]*models.RoleInfo, 0),
		Permissions: make([]*models.PermissionInfo, 0),
	}

	for _, rl := range entity.Roles {
		if rl != nil {
			model.Roles = append(model.Roles, rl.ToModel())
		}
	}

	for _, permission := range entity.Permissions {
		if permission != nil {
			model.Permissions = append(model.Permissions, permission.ToModel())
		}
	}

	return
}
