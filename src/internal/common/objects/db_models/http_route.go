package common_db_models

import (
	"sm-box/internal/common/types"
	"time"
)

type (
	// HttpRoute - модель http маршрута система для базы данных.
	HttpRoute struct {
		ID types.ID `db:"id"`

		Name        string `db:"name"`
		Description string `db:"description"`

		Method string `db:"method"`
		Path   string `db:"path"`

		Active    bool `db:"active"`
		Authorize bool `db:"authorize"`

		RegisterTime time.Time `db:"register_time"`
	}

	// HttpRouteAccess - модель доступа http маршрута система для базы данных.
	HttpRouteAccess struct {
		RouteID types.ID `db:"route_id"`
		RoleID  types.ID `db:"role_id"`
	}
)
