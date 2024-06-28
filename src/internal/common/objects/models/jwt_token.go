package models

import (
	"sm-box/internal/common/types"
	"time"
)

type (
	// JwtTokenInfo - информация о jwt токене системы доступа.
	JwtTokenInfo struct {
		ID        types.ID `json:"id"         xml:"id,attr"`
		UserID    types.ID `json:"user_id"    xml:"user_id,attr"`
		ProjectID types.ID `json:"project_id" xml:"project_id,attr"`

		Language string `json:"language" xml:"language,attr"`
		Data     string `json:"data"     xml:"Data"`

		ExpiresAt time.Time `json:"expires_at" xml:"expires_at,attr"`
		NotBefore time.Time `json:"not_before" xml:"not_before,attr"`
		IssuedAt  time.Time `json:"issued_at"  xml:"issued_at,attr"`
	}

	// JwtTokenInfoParams - информация о параметрах jwt токена системы доступа.
	JwtTokenInfoParams struct {
		RemoteAddr string `json:"remote_addr" xml:"RemoteAddr"`
		UserAgent  string `json:"user_agent"  xml:"UserAgent"`
	}
)
