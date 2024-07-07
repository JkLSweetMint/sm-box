package common_entities

import (
	"github.com/gofiber/fiber/v3"
)

type (
	// HttpRouteConstructor - конструктор Http запроса.
	HttpRouteConstructor struct {
		Name        string
		Description string

		Method string
		Path   string

		Authorize bool
	}
)

// Fill - заполнение данных конструктора с маршрута fiber.
func (entity *HttpRouteConstructor) Fill(route fiber.Route) {
	entity.Path = route.Path
	entity.Method = route.Method
}
