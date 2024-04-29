package entities

import (
	"sm-box/internal/common/types"
	"time"
)

type (
	// HttpRoute - http маршрут системы.
	HttpRoute struct {
		ID     types.ID
		Active bool

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

// ActiveInt - получение числового представления булевого типа в поле Active.
func (entity *HttpRoute) ActiveInt() (v int) {
	if entity.Active {
		v = 1
	}

	return
}

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *HttpRoute) FillEmptyFields() *HttpRoute {
	if entity.Accesses == nil {
		entity.Accesses = make(HttpRouteAccesses, 0)
	}

	return entity
}
