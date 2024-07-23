package db_models

import (
	"github.com/google/uuid"
	"sm-box/internal/common/types"
	"time"
)

type (
	// JwtToken - модель базы данных jwt токена системы доступа.
	JwtToken struct {
		ID       uuid.UUID `json:"id"`
		ParentID uuid.UUID `json:"parent_id"`

		UserID    types.ID `json:"user_id"`
		ProjectID types.ID `json:"project_id"`

		Type string `json:"type"`

		ExpiresAt time.Time `json:"expires_at"`
		NotBefore time.Time `json:"not_before"`
		IssuedAt  time.Time `json:"issued_at"`
	}

	// JwtTokenParams - модель базы данных с параметрами jwt токена системы доступа.
	JwtTokenParams struct {
		RemoteAddr string `json:"remote_addr"`
		UserAgent  string `json:"user_agent"`
	}
)
