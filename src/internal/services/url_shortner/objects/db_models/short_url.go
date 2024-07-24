package db_models

import (
	"encoding/json"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/url_shortner/objects/types"
	"time"
)

type (
	// ShortUrl - модель базы данных короткой ссылке.
	ShortUrl struct {
		ID        common_types.ID `db:"id"`
		Source    string          `db:"source"`
		Reduction string          `db:"reduction"`
	}

	// ShortUrlProperties - модель базы данных свойств короткой ссылке.
	ShortUrlProperties struct {
		URL          common_types.ID    `db:"url"`
		Type         types.ShortUrlType `db:"type"`
		NumberOfUses int64              `db:"number_of_uses"`
		StartActive  time.Time          `db:"start_active"`
		EndActive    time.Time          `db:"end_active"`
	}

	// ShortUrlAccesses - информация по доступам к короткому url.
	ShortUrlAccesses struct {
		Roles       []common_types.ID `db:"roles"`
		Permissions []common_types.ID `db:"permissions"`
	}

	// ShortUrlInfo - модель базы данных redis с информацией по короткой ссылке.
	ShortUrlInfo struct {
		ID        common_types.ID `json:"id"`
		Source    string          `json:"source"`
		Reduction string          `json:"reduction"`

		Accesses   *ShortUrlInfoAccesses   `json:"accesses"`
		Properties *ShortUrlInfoProperties `json:"properties"`
	}

	// ShortUrlInfoProperties - модель базы данных redis с информацией по свойствам короткой ссылке.
	ShortUrlInfoProperties struct {
		Type         types.ShortUrlType `json:"type"`
		NumberOfUses int64              `json:"number_of_uses"`
		StartActive  time.Time          `json:"start_active"`
		EndActive    time.Time          `json:"end_active"`
	}

	// ShortUrlInfoAccesses - информация по доступам к короткому url.
	ShortUrlInfoAccesses struct {
		Roles       []common_types.ID `json:"roles"`
		Permissions []common_types.ID `json:"permissions"`
	}
)

// MarshalBinary - упаковка структуры в бинарный формат.
func (entity *ShortUrlInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(entity)
}

// UnmarshalBinary - распаковка структуры из бинарного формата.
func (entity *ShortUrlInfo) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &entity)
}