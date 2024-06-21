package db_models

import (
	"sm-box/internal/common/types"
)

type (
	// JwtToken - модель jwt токена системы доступа для базы данных.
	JwtToken struct {
		ID     types.ID `db:"id"`
		UserID types.ID `db:"user_id"`

		Data string `db:"data"`

		ExpiresAt string `db:"expires_at"`
		NotBefore string `db:"not_before"`
		IssuedAt  string `db:"issued_at"`
	}

	// JwtTokenParams - модель с параметрами jwt токена системы доступа для базы данных.
	JwtTokenParams struct {
		TokenID types.ID `db:"token_id"`

		RemoteAddr string `db:"remote_addr"`
		UserAgent  string `db:"user_agent"`
	}
)