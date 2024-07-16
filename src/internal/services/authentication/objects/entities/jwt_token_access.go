package entities

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"sm-box/internal/common/types"
	"sm-box/internal/services/authentication/objects/db_models"
	"sm-box/pkg/core/components/tracer"
	"sm-box/pkg/core/env"
	"time"
)

type (
	// JwtAccessToken - jwt токен доступа.
	JwtAccessToken struct {
		*JwtToken

		UserInfo *JwtAccessTokenUserInfo
	}

	// JwtAccessTokenUserInfo - информация о пользователя для jwt токена доступа.
	JwtAccessTokenUserInfo struct {
		Accesses []types.ID
	}

	// JwtAccessTokenClaims - данные для подписи jwt токена доступа.
	JwtAccessTokenClaims struct {
		*jwt.RegisteredClaims

		TokenID uuid.UUID
	}
)

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *JwtAccessToken) FillEmptyFields() *JwtAccessToken {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(entity) }()
	}

	if entity.JwtToken == nil {
		entity.JwtToken = new(JwtToken)
	}

	if entity.UserInfo == nil {
		entity.UserInfo = new(JwtAccessTokenUserInfo)
	}

	var emptyTime time.Time

	if entity.ExpiresAt == emptyTime {
		entity.ExpiresAt = time.Now().Add(3 * time.Minute)
	}

	if entity.NotBefore == emptyTime {
		entity.NotBefore = time.Now()
	}

	if entity.IssuedAt == emptyTime {
		entity.IssuedAt = time.Now()
	}

	entity.JwtToken.FillEmptyFields()
	entity.UserInfo.FillEmptyFields()

	return entity
}

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *JwtAccessTokenClaims) FillEmptyFields() *JwtAccessTokenClaims {
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

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *JwtAccessTokenUserInfo) FillEmptyFields() *JwtAccessTokenUserInfo {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(entity) }()
	}

	if entity.Accesses == nil {
		entity.Accesses = make([]types.ID, 0)
	}

	return entity
}

// Parse - парсинг данных токена доступа.
func (entity *JwtAccessToken) Parse(raw string) (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall(raw)
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	entity.FillEmptyFields()

	var (
		token  *jwt.Token
		claims = new(JwtAccessTokenClaims)
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

// Generate - генерация токена доступа.
func (entity *JwtAccessToken) Generate() (err error) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.Error(err).FunctionCallFinished() }()
	}

	entity.Type = JwtTokenTypeAccess

	entity.FillEmptyFields()

	var claims = &JwtAccessTokenClaims{
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
func (entity *JwtAccessToken) ToDbModel() (model *db_models.JwtAccessToken) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &db_models.JwtAccessToken{
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
		UserInfo: &db_models.JwtAccessTokenUserInfo{
			Accesses: entity.UserInfo.Accesses,
		},
	}

	return
}
