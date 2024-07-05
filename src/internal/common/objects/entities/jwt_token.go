package common_entities

import (
	"github.com/golang-jwt/jwt/v5"
	common_db_models "sm-box/internal/common/objects/db_models"
	common_models "sm-box/internal/common/objects/models"
	"sm-box/internal/common/types"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	"time"
)

type (
	// JwtToken - jwt токен системы доступа.
	JwtToken struct {
		ID        types.ID
		UserID    types.ID
		ProjectID types.ID

		Language string
		Data     string

		ExpiresAt time.Time
		NotBefore time.Time
		IssuedAt  time.Time

		Params *JwtTokenParams
	}

	// JwtTokenParams - параметры jwt токена системы доступа.
	JwtTokenParams struct {
		RemoteAddr string
		UserAgent  string
	}
)

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *JwtToken) FillEmptyFields() *JwtToken {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(entity) }()
	}

	if entity.Params == nil {
		entity.Params = new(JwtTokenParams)
	}

	return entity
}

// ToDbModel - получение модели базы данных.
func (entity *JwtToken) ToDbModel() (model *common_db_models.JwtToken) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &common_db_models.JwtToken{
		ID:        entity.ID,
		UserID:    entity.UserID,
		ProjectID: entity.ProjectID,

		Language: entity.Language,
		Data:     entity.Data,

		ExpiresAt: entity.ExpiresAt,
		NotBefore: entity.NotBefore,
		IssuedAt:  entity.IssuedAt,
	}

	return
}

// ToModel - получение модели.
func (entity *JwtToken) ToModel() (model *common_models.JwtTokenInfo) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &common_models.JwtTokenInfo{
		ID:        entity.ID,
		UserID:    entity.UserID,
		ProjectID: entity.ProjectID,

		Language: entity.Language,
		Data:     entity.Data,

		ExpiresAt: entity.ExpiresAt,
		NotBefore: entity.NotBefore,
		IssuedAt:  entity.IssuedAt,
	}

	return
}

// Generate - генерация данных токена.
func (entity *JwtToken) Generate(claims *jwt.RegisteredClaims) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall(claims)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var tok = jwt.New(jwt.SigningMethodRS256)

	tok.Claims = claims

	if entity.Data, err = tok.SignedString(env.Vars.EncryptionKeys.Private); err != nil {
		return
	}

	return
}

// ToDbModel - получение модели базы данных.
func (entity *JwtTokenParams) ToDbModel() (model *common_db_models.JwtTokenParams) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &common_db_models.JwtTokenParams{
		TokenID:    0,
		RemoteAddr: entity.RemoteAddr,
		UserAgent:  entity.UserAgent,
	}

	return
}

// ToModel - получение модели.
func (entity *JwtTokenParams) ToModel() (model *common_models.JwtTokenInfoParams) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &common_models.JwtTokenInfoParams{
		RemoteAddr: entity.RemoteAddr,
		UserAgent:  entity.UserAgent,
	}

	return
}
