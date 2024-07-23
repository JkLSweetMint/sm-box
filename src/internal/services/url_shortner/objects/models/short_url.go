package models

import (
	"sm-box/internal/common/types"
	"time"
)

type (
	// ShortUrlInfo - внешняя модель короткой ссылки.
	ShortUrlInfo struct {
		ID        types.ID `json:"id"        xml:"id,attr"`
		Source    string   `json:"source"    xml:"source,attr"`
		Reduction string   `json:"reduction" xml:"reduction,attr"`

		Properties *ShortUrlInfoProperties `json:"properties" xml:"Properties"`
	}

	// ShortUrlInfoProperties - внешняя модель свойств короткой ссылке.
	ShortUrlInfoProperties struct {
		Type         string    `json:"type"           xml:"type,attr"`
		NumberOfUses uint16    `json:"number_of_uses" xml:"number_of_uses,attr"`
		StartActive  time.Time `json:"start_active"   xml:"start_active,attr"`
		EndActive    time.Time `json:"end_active"     xml:"end_active,attr"`
	}
)
