package entities

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	"time"
)

type (
	// JwtSessionToken - jwt токен сессии.
	JwtSessionToken struct {
		*JwtToken

		Language string
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

	var (
		token  *jwt.Token
		claims = new(JwtSessionTokenClaims)
	)

	if token, err = jwt.ParseWithClaims(raw, claims, func(t *jwt.Token) (interface{}, error) {
		return env.Vars.EncryptionKeys.Public, nil
	}); err != nil {
		return
	}

	entity.ID = claims.Token.ID
	entity.ParentID = claims.Token.ParentID

	entity.UserID = claims.Token.UserID
	entity.ProjectID = claims.Token.ProjectID

	entity.Type = JwtTokenTypeSession
	entity.Raw = token.Raw

	entity.ExpiresAt = claims.ExpiresAt.Time
	entity.NotBefore = claims.NotBefore.Time
	entity.IssuedAt = claims.IssuedAt.Time

	entity.Language = claims.Language

	entity.Params = claims.Token.Params

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

	var claims = &JwtSessionTokenClaims{
		Token: &JwtTokenClaims{
			ID:       entity.ID,
			ParentID: entity.ParentID,

			UserID:    entity.UserID,
			ProjectID: entity.ProjectID,

			Params: entity.Params,
		},

		Language: entity.Language,

		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: entity.ExpiresAt,
			},
			NotBefore: &jwt.NumericDate{
				Time: entity.NotBefore,
			},
			IssuedAt: &jwt.NumericDate{
				Time: entity.IssuedAt,
			},
		},
	}

	var tok = jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	if entity.Raw, err = tok.SignedString(env.Vars.EncryptionKeys.Private); err != nil {
		return
	}

	return
}

// Value - метод для реализации интерфейса что бы хранить данные в postgresql.
func (entity JwtSessionToken) Value() (driver.Value, error) {
	return json.Marshal(entity)
}

// Scan - метод для реализации интерфейса что бы хранить данные в postgresql.
func (entity *JwtSessionToken) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &entity)
}
