package entities

import (
	"sm-box/internal/common/types"
	"sm-box/internal/services/authentication/objects/db_models"
	"sm-box/pkg/core/components/tracer"
	"strings"
)

type (
	// HttpRoute - http маршрут системы.
	HttpRoute struct {
		ID types.ID

		SystemName  string
		Name        string
		Description string

		Protocols  []string
		Method     string
		Path       string
		RegexpPath string

		Active    bool
		Authorize bool

		Accesses HttpRouteAccesses
	}

	// HttpRouteAccesses - доступы к http маршруту системы.
	HttpRouteAccesses []HttpRouteAccess

	// HttpRouteAccess - доступ к http маршруту системы.
	HttpRouteAccess types.ID
)

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *HttpRoute) FillEmptyFields() *HttpRoute {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(entity) }()
	}

	if entity.Accesses == nil {
		entity.Accesses = make(HttpRouteAccesses, 0)
	}

	return entity
}

// ToDbModel - получение модели базы данных.
func (entity *HttpRoute) ToDbModel() (model *db_models.HttpRoute) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &db_models.HttpRoute{
		ID: entity.ID,

		SystemName:  entity.SystemName,
		Name:        entity.Name,
		Description: entity.Description,

		Protocols:  entity.Protocols,
		Method:     strings.ToUpper(entity.Method),
		Path:       entity.Path,
		RegexpPath: entity.RegexpPath,

		Active:    entity.Active,
		Authorize: entity.Authorize,
	}

	return
}
