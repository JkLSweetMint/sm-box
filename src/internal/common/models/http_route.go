package models

import (
	"sm-box/internal/common/types"
	"time"
)

type (
	// HttpRouteInfo - модель с информацией о http маршруте системы.
	HttpRouteInfo struct {
		ID     types.ID `json:"id"     yaml:"ID"     xml:"id,attr"`
		Active bool     `json:"active" yaml:"Active" xml:"active,attr"`

		Method string `json:"method" yaml:"Method" xml:"method,attr"`
		Path   string `json:"path"   yaml:"Path"   xml:"path,attr"`

		RegisterTime time.Time             `json:"register_time"      yaml:"RegisterTime"       xml:"register_time,attr"`
		Accesses     HttpRouteInfoAccesses `json:"accesses,omitempty" yaml:"Accesses,omitempty" xml:"Accesses,omitempty>Access"`
	}

	// HttpRouteInfoAccesses - модель с информацией о доступах к маршруту.
	HttpRouteInfoAccesses []*HttpRouteInfoAccess

	// HttpRouteInfoAccess - модель с информацией о доступе к маршруту.
	HttpRouteInfoAccess struct {
		*RoleInfo
	}
)
