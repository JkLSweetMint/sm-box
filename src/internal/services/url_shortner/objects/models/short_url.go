package models

import (
	common_types "sm-box/internal/common/types"
	authentication_entities "sm-box/internal/services/authentication/objects/entities"
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
		RolesID       []common_types.ID `json:"roles_id"       xml:"RolesID>ID"`
		PermissionsID []common_types.ID `json:"permissions_id" xml:"PermissionsID>ID"`
	}

	// ShortUrlInfoProperties - внешняя модель свойств короткой ссылке.
	ShortUrlInfoProperties struct {
		Type   types.ShortUrlType `json:"type"   xml:"type,attr"`
		Active bool               `json:"active" xml:"active,attr"`

		NumberOfUses         int64 `json:"number_of_uses"          xml:"number_of_uses,attr"`
		RemainedNumberOfUses int64 `json:"remained_number_of_uses" xml:"remained_number_of_uses,attr"`

		StartActive time.Time `json:"start_active" xml:"start_active,attr"`
		EndActive   time.Time `json:"end_active"   xml:"end_active,attr"`
	}

	// ShortUrlUsageHistoryInfo - внешняя модель истории использования короткой ссылке.
	ShortUrlUsageHistoryInfo struct {
		Status    string                                   `json:"status"     xml:"status,attr"`
		Timestamp time.Time                                `json:"timestamp"  xml:"timestamp,attr"`
		TokenInfo *authentication_entities.JwtSessionToken `json:"token_info" xml:"TokenInfo"`
	}
)
