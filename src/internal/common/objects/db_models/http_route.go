package db_models

import (
	"sm-box/internal/common/types"
	"time"
)

type (
	// HttpRoute - модель http маршрута система для базы данных.
	HttpRoute struct {
		ID        types.ID `db:"id"`
		Active    string   `db:"active"`
		Authorize string   `db:"authorize"`

		Method string `db:"method"`
		Path   string `db:"path"`

		RegisterTime time.Time `db:"register_time"`
	}

	// HttpRouteAccess - модель доступа http маршрута система для базы данных.
	HttpRouteAccess struct {
		RouteID types.ID `db:"route_id"`
		RoleID  types.ID `db:"role_id"`
	}
)
