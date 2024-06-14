package entities

import (
	"sm-box/internal/app/infrastructure/objects/db_models"
	"sm-box/internal/app/infrastructure/types"
	"sm-box/pkg/core/components/tracer"
	"strings"
	"time"
)

type (
	// HttpRoute - http маршрут системы.
	HttpRoute struct {
		ID        types.ID
		Active    bool
		Authorize bool

		Method string
		Path   string

		RegisterTime time.Time
		Accesses     HttpRouteAccesses
	}

	// HttpRouteAccesses - доступы к http маршруту системы.
	HttpRouteAccesses []*HttpRouteAccess

	// HttpRouteAccess - доступ к http маршруту системы.
	HttpRouteAccess struct {
		*Role
	}
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

// DbModel - получение модели базы данных.
func (entity *HttpRoute) DbModel() (model *db_models.HttpRoute) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &db_models.HttpRoute{
		ID: entity.ID,

		Active:    "off",
		Authorize: "off",

		Method: strings.ToUpper(entity.Method),
		Path:   entity.Path,

		RegisterTime: entity.RegisterTime.Format(time.RFC3339Nano),
	}

	if entity.Active {
		model.Active = "on"
	}

	if entity.Authorize {
		model.Authorize = "on"
	}

	return
}
