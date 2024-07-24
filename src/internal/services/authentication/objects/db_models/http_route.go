package db_models

import (
	"github.com/lib/pq"
	common_types "sm-box/internal/common/types"
)

type (
	// HttpRoute - модель базы данных http маршрута
	HttpRoute struct {
		ID common_types.ID `db:"id"`

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

	// HttpRouteAccess - модель базы данных доступа http маршрута.
	HttpRouteAccess struct {
		RouteID      common_types.ID `db:"route_id"`
		RoleID       common_types.ID `db:"role_id"`
		PermissionID common_types.ID `db:"permission_id"`
	}
)
