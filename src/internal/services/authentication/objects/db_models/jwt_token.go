package db_models

import (
	"github.com/google/uuid"
	"sm-box/internal/common/types"
	"time"
)

type (
	// JwtToken - модель jwt токена системы доступа для базы данных.
	JwtToken struct {
		ID       uuid.UUID `db:"id"`
		ParentID uuid.UUID `db:"parent_id"`

		UserID    types.ID `db:"user_id"`
		ProjectID types.ID `db:"project_id"`

		Type string `db:"type"`
		Raw  string `db:"raw"`

		ExpiresAt time.Time `db:"expires_at"`
		NotBefore time.Time `db:"not_before"`
		IssuedAt  time.Time `db:"issued_at"`
	}

	// JwtTokenParams - модель с параметрами jwt токена системы доступа для базы данных.
	JwtTokenParams struct {
		TokenID uuid.UUID `db:"token_id"`

		RemoteAddr string `db:"remote_addr"`
		UserAgent  string `db:"user_agent"`
	}
)
