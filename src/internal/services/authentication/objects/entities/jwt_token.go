package entities

import (
	"github.com/golang-jwt/jwt/v5"
	app_models "sm-box/internal/app/objects/models"
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
		ID       types.ID
		ParentID types.ID

		UserID    types.ID
		ProjectID types.ID

		Raw string

		ExpiresAt time.Time
		NotBefore time.Time
		IssuedAt  time.Time

		Params *JwtTokenParams
	}

	// JwtTokenParams - параметры jwt токена системы доступа.
	JwtTokenParams struct {
		Language   string
		RemoteAddr string
		UserAgent  string
	}

	// JwtTokenClaims - claims для формирование jwt токена.
	JwtTokenClaims struct {
		jwt.RegisteredClaims

		Token *JwtToken
		User  *app_models.UserInfo
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

	var emptyTime time.Time

	if entity.ExpiresAt == emptyTime {
		entity.ExpiresAt = time.Now().Add(time.Hour)
	}

	if entity.NotBefore == emptyTime {
		entity.NotBefore = time.Now()
	}

	if entity.IssuedAt == emptyTime {
		entity.IssuedAt = time.Now()
	}

	if entity.Params == nil {
		entity.Params = new(JwtTokenParams)
	}

	entity.Params.FillEmptyFields()

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
		ID:       entity.ID,
		ParentID: entity.ParentID,

		UserID:    entity.UserID,
		ProjectID: entity.ProjectID,

		Raw: entity.Raw,

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

		Raw: entity.Raw,

		ExpiresAt: entity.ExpiresAt,
		NotBefore: entity.NotBefore,
		IssuedAt:  entity.IssuedAt,
	}

	return
}

// Parse - парсинг данных токена.
func (entity *JwtToken) Parse(raw string) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall(raw)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	var (
		t      *jwt.Token
		claims = new(JwtTokenClaims)
	)

	if t, err = jwt.ParseWithClaims(raw, claims, func(t *jwt.Token) (interface{}, error) {
		return env.Vars.EncryptionKeys.Public, nil
	}); err != nil {
		return
	}

	entity.ID = claims.Token.ID
	entity.ParentID = claims.Token.ParentID

	entity.UserID = claims.Token.UserID
	entity.ProjectID = claims.Token.ProjectID

	entity.Raw = t.Raw

	entity.ExpiresAt = claims.Token.ExpiresAt
	entity.NotBefore = claims.Token.NotBefore
	entity.IssuedAt = claims.Token.IssuedAt

	entity.Params = claims.Token.Params

	return
}

// Generate - генерация токена.
func (entity *JwtToken) Generate() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	entity.FillEmptyFields()

	var claims = &JwtTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    env.Vars.SystemName,
			ExpiresAt: &jwt.NumericDate{Time: entity.ExpiresAt},
			NotBefore: &jwt.NumericDate{Time: entity.NotBefore},
			IssuedAt:  &jwt.NumericDate{Time: entity.IssuedAt},
		},

		Token: &JwtToken{
			ID:       entity.ID,
			ParentID: entity.ParentID,

			UserID:    entity.UserID,
			ProjectID: entity.ProjectID,

			Raw: "",

			ExpiresAt: entity.ExpiresAt,
			NotBefore: entity.NotBefore,
			IssuedAt:  entity.IssuedAt,

			Params: &JwtTokenParams{
				Language:   entity.Params.Language,
				RemoteAddr: entity.Params.RemoteAddr,
				UserAgent:  entity.Params.UserAgent,
			},
		},
	}

	var tok = jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	if entity.Raw, err = tok.SignedString(env.Vars.EncryptionKeys.Private); err != nil {
		return
	}

	return
}

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *JwtTokenParams) FillEmptyFields() *JwtTokenParams {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(entity) }()
	}

	return entity
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
		Language:   entity.Language,
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
		Language:   entity.Language,
		RemoteAddr: entity.RemoteAddr,
		UserAgent:  entity.UserAgent,
	}

	return
}
