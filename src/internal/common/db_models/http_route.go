package db_models

import "sm-box/internal/common/types"

type (
	// HttpRoute - модель http маршрута система для базы данных.
	HttpRoute struct {
		ID     types.ID `db:"id"`
		Active bool     `db:"active"`

		Method string `db:"method"`
		Path   string `db:"path"`
	}

	// HttpRouteAccess - модель доступа http маршрута система для базы данных.
	HttpRouteAccess struct {
		RouteID types.ID `db:"route_id"`
		RoleID  types.ID `db:"role_id"`
	}
)
