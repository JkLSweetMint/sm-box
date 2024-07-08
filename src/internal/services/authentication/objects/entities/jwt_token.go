package entities

import (
	"github.com/golang-jwt/jwt/v5"
	"sm-box/internal/common/types"
	"sm-box/internal/services/authentication/objects/db_models"
	"sm-box/internal/services/authentication/objects/models"
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
		Raw      string

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
func (entity *JwtToken) ToDbModel() (model *db_models.JwtToken) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &db_models.JwtToken{
		ID:        entity.ID,
		UserID:    entity.UserID,
		ProjectID: entity.ProjectID,

		Language: entity.Language,
		Raw:      entity.Raw,

		ExpiresAt: entity.ExpiresAt,
		NotBefore: entity.NotBefore,
		IssuedAt:  entity.IssuedAt,
	}

	return
}

// ToModel - получение модели.
func (entity *JwtToken) ToModel() (model *models.JwtTokenInfo) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &models.JwtTokenInfo{
		ID:        entity.ID,
		UserID:    entity.UserID,
		ProjectID: entity.ProjectID,

		Language: entity.Language,
		Raw:      entity.Raw,

		ExpiresAt: entity.ExpiresAt,
		NotBefore: entity.NotBefore,
		IssuedAt:  entity.IssuedAt,
	}

	return
}

// Parse - парсинг данных токена.
func (entity *JwtToken) Parse(data string) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall(data)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var t *jwt.Token

	if t, err = jwt.Parse(data, func(t *jwt.Token) (interface{}, error) {
		return env.Vars.EncryptionKeys.Public, nil
	}); err != nil {
		return
	}

	entity.Raw = t.Raw

	var (
		expirationTime, notBefore, issuedAt *jwt.NumericDate
	)

	if expirationTime, err = t.Claims.GetExpirationTime(); err != nil {
		return
	}

	if notBefore, err = t.Claims.GetNotBefore(); err != nil {
		return
	}

	if issuedAt, err = t.Claims.GetIssuedAt(); err != nil {
		return
	}

	entity.ExpiresAt = expirationTime.Time
	entity.NotBefore = notBefore.Time
	entity.IssuedAt = issuedAt.Time

	//fmt.Printf("\n\n\n%+v\n\n\n", t.Claims.(jwt.MapClaims))

	return
}

// ToDbModel - получение модели базы данных.
func (entity *JwtTokenParams) ToDbModel() (model *db_models.JwtTokenParams) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &db_models.JwtTokenParams{
		TokenID:    0,
		RemoteAddr: entity.RemoteAddr,
		UserAgent:  entity.UserAgent,
	}

	return
}

// ToModel - получение модели.
func (entity *JwtTokenParams) ToModel() (model *models.JwtTokenInfoParams) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &models.JwtTokenInfoParams{
		RemoteAddr: entity.RemoteAddr,
		UserAgent:  entity.UserAgent,
	}

	return
}
