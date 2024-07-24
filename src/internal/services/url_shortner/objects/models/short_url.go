package models

import (
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/url_shortner/objects/types"
	"time"
)

type (
	// ShortUrlInfo - внешняя модель короткой ссылки.
	ShortUrlInfo struct {
		ID        common_types.ID `json:"id"        xml:"id,attr"`
		Source    string          `json:"source"    xml:"source,attr"`
		Reduction string          `json:"reduction" xml:"reduction,attr"`

		Accesses   *ShortUrlAccesses       `json:"accesses"   xml:"Accesses"`
		Properties *ShortUrlInfoProperties `json:"properties" xml:"Properties"`
	}

	// ShortUrlAccesses - информация по доступам к короткому url.
	ShortUrlAccesses struct {
		Roles       []common_types.ID `json:"roles"       xml:"Roles"`
		Permissions []common_types.ID `json:"permissions" xml:"Permissions"`
	}

	// ShortUrlInfoProperties - внешняя модель свойств короткой ссылке.
	ShortUrlInfoProperties struct {
		Type         types.ShortUrlType `json:"type"           xml:"type,attr"`
		NumberOfUses int64              `json:"number_of_uses" xml:"number_of_uses,attr"`
		StartActive  time.Time          `json:"start_active"   xml:"start_active,attr"`
		EndActive    time.Time          `json:"end_active"     xml:"end_active,attr"`
	}
)
