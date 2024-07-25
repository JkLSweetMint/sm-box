package models

import common_types "sm-box/internal/common/types"

type (
	// HttpRouteInfo - внешняя модель с информацией о http маршруте системы.
	HttpRouteInfo struct {
		ID common_types.ID `json:"id" xml:"id,attr"`

		SystemName  string `json:"system_name" xml:"SystemName"`
		Name        string `json:"name"        xml:"Name"`
		Description string `json:"description" xml:"Description"`

		Protocols  []string `json:"protocols"   xml:"protocols,attr"`
		Method     string   `json:"method"      xml:"method,attr"`
		Path       string   `json:"path"        xml:"path,attr"`
		RegexpPath string   `json:"regexp_path" xml:"regexp_path,attr"`

		Active bool `json:"active" xml:"active,attr"`

		Accesses *HttpRouteInfoAccesses `json:"accesses,omitempty" xml:"Accesses>Access,omitempty"`
	}

	// HttpRouteInfoAccesses - внешняя модель с информацией о доступах к маршруту.
	HttpRouteInfoAccesses struct {
		Roles       []common_types.ID `json:"roles"       xml:"Roles>Role"`
		Permissions []common_types.ID `json:"permissions" xml:"Permissions>Permission"`
	}
)
