package common_models

import (
	"sm-box/internal/common/types"
)

type (
	// HttpRouteInfo - модель с информацией о http маршруте системы.
	HttpRouteInfo struct {
		ID types.ID `json:"id" xml:"id,attr"`

		Name        string `json:"name"        xml:"Name"`
		Description string `json:"description" xml:"Description"`

		Method string `json:"method" xml:"method,attr"`
		Path   string `json:"path"   xml:"path,attr"`

		Active    bool `json:"active"    xml:"active,attr"`
		Authorize bool `json:"authorize" xml:"authorize,attr"`

		Accesses HttpRouteInfoAccesses `json:"accesses,omitempty" xml:"Accesses,omitempty>Access"`
	}

	// HttpRouteInfoAccesses - модель с информацией о доступах к маршруту.
	HttpRouteInfoAccesses []*HttpRouteInfoAccess

	// HttpRouteInfoAccess - модель с информацией о доступе к маршруту.
	HttpRouteInfoAccess struct {
		RouteID types.ID `json:"route_id" xml:"route_id,attr"`
		RoleID  types.ID `json:"role_id"  xml:"role_id,attr"`
	}
)
