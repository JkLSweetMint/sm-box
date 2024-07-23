package models

import (
	"github.com/google/uuid"
	"sm-box/internal/common/types"
	"time"
)

type (
	// JwtTokenInfo - внешняя модель с информацией о jwt токене системы доступа.
	JwtTokenInfo struct {
		ID       uuid.UUID `json:"id"        xml:"id,attr"`
		ParentID uuid.UUID `json:"parent_id" xml:"parent_id,attr"`

		UserID    types.ID `json:"user_id"    xml:"user_id,attr"`
		ProjectID types.ID `json:"project_id" xml:"project_id,attr"`

		Type string `json:"type" xml:"type,attr"`
		Raw  string `json:"raw"  xml:"Raw"`

		ExpiresAt time.Time `json:"expires_at" xml:"expires_at,attr"`
		NotBefore time.Time `json:"not_before" xml:"not_before,attr"`
		IssuedAt  time.Time `json:"issued_at"  xml:"issued_at,attr"`
	}

	// JwtTokenInfoParams - внешняя модель с информацией о параметрах jwt токена системы доступа.
	JwtTokenInfoParams struct {
		Language   string `json:"language"    xml:"language,attr"`
		RemoteAddr string `json:"remote_addr" xml:"remote_addr,attr"`
		UserAgent  string `json:"user_agent"  xml:"UserAgent"`
	}
)
