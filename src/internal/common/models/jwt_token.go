package models

import (
	"sm-box/internal/common/types"
	"time"
)

type (
	// JwtTokenInfo - информация о jwt токене системы доступа.
	JwtTokenInfo struct {
		ID     types.ID `json:"id"      yaml:"ID"     xml:"id,attr"`
		UserID types.ID `json:"user_id" yaml:"UserID" xml:"user_id,attr"`

		Data string `json:"data" yaml:"Data" xml:"Data"`

		ExpiresAt time.Time `json:"expires_at" yaml:"ExpiresAt" xml:"expires_at,attr"`
		NotBefore time.Time `json:"not_before" yaml:"NotBefore" xml:"not_before,attr"`
		IssuedAt  time.Time `json:"issued_at"  yaml:"IssuedAt"  xml:"issued_at,attr"`

		Params *JwtTokenInfoParams `json:"params" yaml:"Params" xml:"Params>Param"`
	}

	// JwtTokenInfoParams - информация о параметрах jwt токене системы доступа.
	JwtTokenInfoParams struct {
		RemoteAddr string `json:"remote_addr" yaml:"RemoteAddr" xml:"RemoteAddr"`
		UserAgent  string `json:"user_agent"  yaml:"UserAgent"  xml:"UserAgent"`
	}
)
