package models

import "sm-box/internal/common/types"

type (
	// HttpRoute - http маршрут системы.
	HttpRoute struct {
		ID     types.ID
		Active bool

		Method string
		Path   string

		Accesses HttpRouteAccesses
	}

	// HttpRouteAccesses - доступы к http маршруту системы.
	HttpRouteAccesses []*HttpRouteAccess

	// HttpRouteAccess - доступ к http маршруту системы.
	HttpRouteAccess struct {
		*Role
	}
)
