package entities

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"sm-box/internal/services/authentication/objects/db_models"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	"time"
)

type (
	// JwtRefreshToken - jwt токен обновления.
	JwtRefreshToken struct {
		*JwtToken
	}

	// JwtRefreshTokenClaims - данные для подписи jwt токена обновления.
	JwtRefreshTokenClaims struct {
		*jwt.RegisteredClaims

		TokenID uuid.UUID
	}
)

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *JwtRefreshToken) FillEmptyFields() *JwtRefreshToken {
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
func (entity *JwtRefreshTokenClaims) FillEmptyFields() *JwtRefreshTokenClaims {
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

// Parse - парсинг данных токена обновления.
func (entity *JwtRefreshToken) Parse(raw string) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall(raw)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	entity.FillEmptyFields()

	var (
		claims = new(JwtRefreshTokenClaims)
		token  *jwt.Token
	)

	if token, err = jwt.ParseWithClaims(raw, claims, func(t *jwt.Token) (interface{}, error) {
		return env.Vars.EncryptionKeys.Public, nil
	}); err != nil {
		return
	}

	entity.ID = claims.TokenID

	entity.Type = JwtTokenTypeSession
	entity.Raw = token.Raw

	return
}

// Generate - генерация токена обновления.
func (entity *JwtRefreshToken) Generate() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	entity.Type = JwtTokenTypeRefresh

	entity.FillEmptyFields()

	var claims = &JwtRefreshTokenClaims{
		TokenID: entity.JwtToken.ID,

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

// ToDbModel - получение модели базы данных.
func (entity *JwtRefreshToken) ToDbModel() (model *db_models.JwtRefreshToken) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &db_models.JwtRefreshToken{
		JwtToken: &db_models.JwtToken{
			ID:       entity.ID,
			ParentID: entity.ParentID,

			UserID:    entity.UserID,
			ProjectID: entity.ProjectID,

			Type: string(entity.Type),

			ExpiresAt: entity.ExpiresAt,
			NotBefore: entity.NotBefore,
			IssuedAt:  entity.IssuedAt,
		},
	}

	return
}
