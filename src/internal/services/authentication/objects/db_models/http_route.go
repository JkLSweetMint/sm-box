package db_models

import (
	"github.com/lib/pq"
	"sm-box/internal/common/types"
)

type (
	// HttpRoute - модель http маршрута система для базы данных.
	HttpRoute struct {
		ID types.ID `db:"id"`

		SystemName  string `db:"system_name"`
		Name        string `db:"name"`
		Description string `db:"description"`

		Protocols  pq.StringArray `db:"protocols"`
		Method     string         `db:"method"`
		Path       string         `db:"path"`
		RegexpPath string         `db:"regexp_path"`

		Active    bool `db:"active"`
		Authorize bool `db:"authorize"`
	}

	// HttpRouteAccess - модель доступа http маршрута система для базы данных.
	HttpRouteAccess struct {
		RouteID types.ID `db:"route_id"`
		RoleID  types.ID `db:"role_id"`
	}
)
