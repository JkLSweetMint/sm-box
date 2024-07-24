package entities

import (
	"github.com/google/uuid"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/authentication/objects/models"
	"sm-box/pkg/core/components/tracer"
	"time"
)

const (
	JwtTokenTypeSession JwtTokenType = "session"
	JwtTokenTypeAccess  JwtTokenType = "access"
	JwtTokenTypeRefresh JwtTokenType = "refresh"
)

type (
	// JwtToken - jwt токен системы доступа.
	JwtToken struct {
		ID       uuid.UUID
		ParentID uuid.UUID

		UserID    common_types.ID
		ProjectID common_types.ID

		Type JwtTokenType
		Raw  string

		ExpiresAt time.Time
		NotBefore time.Time
		IssuedAt  time.Time

		Params *JwtTokenParams
	}

	// JwtTokenType - тип токена.
	JwtTokenType string

	// JwtTokenParams - параметры jwt токена системы доступа.
	JwtTokenParams struct {
		RemoteAddr string
		UserAgent  string
	}

	// JwtTokenClaims - данные для подписи jwt токена.
	JwtTokenClaims struct {
		ID       uuid.UUID
		ParentID uuid.UUID

		UserID    common_types.ID
		ProjectID common_types.ID

		Params *JwtTokenParams
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

	if entity.ID.String() == new(uuid.UUID).String() {
		entity.ID = uuid.New()
	}

	if entity.Params == nil {
		entity.Params = new(JwtTokenParams)
	}

	entity.Params.FillEmptyFields()

	return entity
}

// ToModel - получение внешней модели.
func (entity *JwtToken) ToModel() (model *models.JwtTokenInfo) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	entity.FillEmptyFields()

	model = &models.JwtTokenInfo{
		ID:       entity.ID,
		ParentID: entity.ParentID,

		UserID:    entity.UserID,
		ProjectID: entity.ProjectID,

		Type: string(entity.Type),
		Raw:  entity.Raw,

		ExpiresAt: entity.ExpiresAt,
		NotBefore: entity.NotBefore,
		IssuedAt:  entity.IssuedAt,
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
