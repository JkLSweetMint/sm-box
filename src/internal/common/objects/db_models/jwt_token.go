package common_db_models

import (
	"sm-box/internal/common/types"
	"time"
)

type (
	// JwtToken - модель jwt токена системы доступа для базы данных.
	JwtToken struct {
		ID        types.ID `db:"id"`
		UserID    types.ID `db:"user_id"`
		ProjectID types.ID `db:"project_id"`

		Language string `db:"language"`
		Data     string `db:"data"`

		ExpiresAt time.Time `db:"expires_at"`
		NotBefore time.Time `db:"not_before"`
		IssuedAt  time.Time `db:"issued_at"`
	}

	// JwtTokenParams - модель с параметрами jwt токена системы доступа для базы данных.
	JwtTokenParams struct {
		TokenID types.ID `db:"token_id"`

		RemoteAddr string `db:"remote_addr"`
		UserAgent  string `db:"user_agent"`
	}
)
