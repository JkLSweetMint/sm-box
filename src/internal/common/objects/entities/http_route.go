package entities

import (
	"github.com/gofiber/fiber/v3"
	"sm-box/internal/common/objects/db_models"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/tracer"
	"strings"
	"time"
)

type (
	// HttpRoute - http маршрут системы.
	HttpRoute struct {
		ID types.ID

		Name        string
		Description string

		Method string
		Path   string

		Active    bool
		Authorize bool

		RegisterTime time.Time
		Accesses     HttpRouteAccesses
	}

	// HttpRouteConstructor - конструктор Http запроса.
	HttpRouteConstructor struct {
		Name        string
		Description string

		Method string
		Path   string

		Authorize bool
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

		Name:        entity.Name,
		Description: entity.Description,

		Method: strings.ToUpper(entity.Method),
		Path:   entity.Path,

		Active:    entity.Active,
		Authorize: entity.Authorize,

		RegisterTime: entity.RegisterTime,
	}

	return
}

// Fill - заполнение данных конструктора с маршрута fiber.
func (entity *HttpRouteConstructor) Fill(route fiber.Route) {
	entity.Path = route.Path
	entity.Method = route.Method
}

// Build - построение маршрута.
func (entity *HttpRouteConstructor) Build() (route *HttpRoute) {
	route = new(HttpRoute).FillEmptyFields()

	route.Name = entity.Name
	route.Description = entity.Description

	route.Path = entity.Path
	route.Method = entity.Method

	route.Active = true
	route.Authorize = entity.Authorize

	return
}
