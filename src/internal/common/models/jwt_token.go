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

		CreatedAt time.Time `json:"created_at" yaml:"CreatedAt" xml:"created_at,attr"`
		ExpiredAt time.Time `json:"expired_at" yaml:"ExpiredAt" xml:"expired_at,attr"`

		Params *JwtTokenInfoParams `json:"params" yaml:"Params" xml:"Params>Param"`
	}

	// JwtTokenInfoParams - информация о параметрах jwt токене системы доступа.
	JwtTokenInfoParams struct {
		RemoteAddr string `json:"remote_addr" yaml:"RemoteAddr" xml:"RemoteAddr"`
		UserAgent  string `json:"user_agent"  yaml:"UserAgent"  xml:"UserAgent"`
	}
)
