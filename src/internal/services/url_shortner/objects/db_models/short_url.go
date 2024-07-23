package db_models

import (
	"encoding/json"
	"sm-box/internal/common/types"
	"time"
)

type (
	// ShortUrl - модель базы данных короткой ссылке.
	ShortUrl struct {
		ID        types.ID `db:"id"`
		Source    string   `db:"source"`
		Reduction string   `db:"reduction"`
	}

	// ShortUrlProperties - модель базы данных свойств короткой ссылке.
	ShortUrlProperties struct {
		URL          types.ID  `db:"url"`
		Type         string    `db:"type"`
		NumberOfUses uint16    `db:"number_of_uses"`
		StartActive  time.Time `db:"start_active"`
		EndActive    time.Time `db:"end_active"`
	}

	// ShortUrlInfo - модель базы данных redis с информацией по короткой ссылке.
	ShortUrlInfo struct {
		ID        types.ID `json:"id"`
		Source    string   `json:"source"`
		Reduction string   `json:"reduction"`

		Properties *ShortUrlInfoProperties `json:"properties"`
	}

	// ShortUrlInfoProperties - модель базы данных redis с информацией по свойствам короткой ссылке.
	ShortUrlInfoProperties struct {
		Type         string    `json:"type"`
		NumberOfUses uint16    `json:"number_of_uses"`
		StartActive  time.Time `json:"start_active"`
		EndActive    time.Time `json:"end_active"`
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
