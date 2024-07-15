package entities

import (
	"github.com/golang-jwt/jwt/v5"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	"time"
)

type (
	// JwtSessionToken - jwt токен сессии.
	JwtSessionToken struct {
		*JwtToken

		Claims *JwtSessionTokenClaims
	}

	// JwtSessionTokenClaims - данные для подписи jwt токена сессии.
	JwtSessionTokenClaims struct {
		*jwt.RegisteredClaims

		Token *JwtTokenClaims

		Language string
	}
)

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *JwtSessionToken) FillEmptyFields() *JwtSessionToken {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(entity) }()
	}

	if entity.JwtToken == nil {
		entity.JwtToken = new(JwtToken)
	}

	if entity.Claims == nil {
		entity.Claims = new(JwtSessionTokenClaims)
	}

	var emptyTime time.Time

	if entity.ExpiresAt == emptyTime {
		entity.ExpiresAt = time.Now().Add(7 * 24 * time.Hour)
	}

	if entity.NotBefore == emptyTime {
		entity.NotBefore = time.Now()
	}

	if entity.IssuedAt == emptyTime {
		entity.IssuedAt = time.Now()
	}

	entity.JwtToken.FillEmptyFields()
	entity.Claims.FillEmptyFields()

	return entity
}

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *JwtSessionTokenClaims) FillEmptyFields() *JwtSessionTokenClaims {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(entity) }()
	}

	if entity.RegisteredClaims == nil {
		entity.RegisteredClaims = new(jwt.RegisteredClaims)
	}

	return entity
}

// Parse - парсинг данных токена сессии.
func (entity *JwtSessionToken) Parse(raw string) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall(raw)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	entity.FillEmptyFields()

	var token *jwt.Token

	if token, err = jwt.ParseWithClaims(raw, entity.Claims, func(t *jwt.Token) (interface{}, error) {
		return env.Vars.EncryptionKeys.Public, nil
	}); err != nil {
		return
	}

	entity.ID = entity.Claims.Token.ID
	entity.ParentID = entity.Claims.Token.ParentID

	entity.UserID = entity.Claims.Token.UserID
	entity.ProjectID = entity.Claims.Token.ProjectID

	entity.Type = JwtTokenTypeSession
	entity.Raw = token.Raw

	entity.ExpiresAt = entity.Claims.ExpiresAt.Time
	entity.NotBefore = entity.Claims.NotBefore.Time
	entity.IssuedAt = entity.Claims.IssuedAt.Time

	entity.Params = entity.Claims.Token.Params

	return
}

// Generate - генерация токена сессии.
func (entity *JwtSessionToken) Generate() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	entity.Type = JwtTokenTypeSession

	entity.FillEmptyFields()

	entity.Claims.Token = &JwtTokenClaims{
		ID:       entity.ID,
		ParentID: entity.ParentID,

		UserID:    entity.UserID,
		ProjectID: entity.ProjectID,

		Params: entity.Params,
	}

	entity.Claims.ExpiresAt = &jwt.NumericDate{
		Time: entity.ExpiresAt,
	}
	entity.Claims.NotBefore = &jwt.NumericDate{
		Time: entity.NotBefore,
	}
	entity.Claims.IssuedAt = &jwt.NumericDate{
		Time: entity.IssuedAt,
	}

	var tok = jwt.NewWithClaims(jwt.SigningMethodRS256, entity.Claims)

	if entity.Raw, err = tok.SignedString(env.Vars.EncryptionKeys.Private); err != nil {
		return
	}

	return
}
