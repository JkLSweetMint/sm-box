package common_entities

import (
	"sm-box/internal/common/objects/models"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/tracer"
)

type (
	// User - пользователь системы.
	User struct {
		ID        types.ID
		ProjectID types.ID

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

// CheckHttpRouteAccesses - проверить доступ к http маршрутам.
func (entity *User) CheckHttpRouteAccesses(access HttpRouteAccesses) (ok bool) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall(access)
		defer func() { trc.FunctionCallFinished(ok) }()
	}

	for _, role := range entity.Accesses {
		for _, acc := range access {
			if role.ID == acc.ID {
				ok = true
				return
			}
		}
	}

	for _, role := range entity.Accesses {
		if ok = entity.checkUserRoleInheritancesForHttpRouteAccesses(role.Inheritances, access); ok {
			return
		}
	}

	return
}

// checkUserRoleInheritancesForHttpRouteAccesses - проверить доступы пользователя для http маршрутов.
func (entity *User) checkUserRoleInheritancesForHttpRouteAccesses(userAccess RoleInheritances, access HttpRouteAccesses) (ok bool) {
	for _, role := range userAccess {
		for _, acc := range access {
			if role.ID == acc.ID {
				ok = true
				return
			}
		}
	}

	for _, role := range userAccess {
		if ok = entity.checkUserRoleInheritancesForHttpRouteAccesses(role.Inheritances, access); ok {
			return
		}
	}

	return
}

// ToModel - получение модели.
func (entity *User) ToModel() (model *common_models.UserInfo) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &common_models.UserInfo{
		ID:        entity.ID,
		ProjectID: entity.ProjectID,
		Email:     entity.Email,
		Username:  entity.Username,
		Password:  entity.Password,
		Accesses:  make(common_models.UserInfoAccesses, 0),
	}

	for _, acc := range entity.Accesses {
		model.Accesses = append(model.Accesses, acc.ToModel())
	}

	return
}

// ToModel - получение модели.
func (entity *UserAccess) ToModel() (model *common_models.UserInfoAccess) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &common_models.UserInfoAccess{
		RoleInfo: entity.Role.ToModel(),
	}

	return
}
