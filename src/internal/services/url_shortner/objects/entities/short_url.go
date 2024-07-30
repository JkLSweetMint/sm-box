package entities

import (
	common_types "sm-box/internal/common/types"
	authentication_entities "sm-box/internal/services/authentication/objects/entities"
	"sm-box/internal/services/url_shortner/objects/db_models"
	"sm-box/internal/services/url_shortner/objects/models"
	"sm-box/internal/services/url_shortner/objects/types"
	"sm-box/pkg/core/components/tracer"
	"time"
)

type (
	// ShortUrl - короткая ссылка.
	ShortUrl struct {
		ID        common_types.ID
		Source    string
		Reduction string

		Accesses   *ShortUrlAccesses
		Properties *ShortUrlProperties
	}

	// ShortUrlAccesses - информация по доступам к короткому url.
	ShortUrlAccesses struct {
		RolesID       []common_types.ID
		PermissionsID []common_types.ID
	}

	// ShortUrlProperties - свойства короткой ссылке.
	ShortUrlProperties struct {
		Type                 types.ShortUrlType
		NumberOfUses         int64
		RemainedNumberOfUses int64
		StartActive          time.Time
		EndActive            time.Time
		Active               bool
	}

	// ShortUrlUsageHistory - история использования короткой ссылке.
	ShortUrlUsageHistory struct {
		Status    string
		Timestamp time.Time
		TokenInfo *authentication_entities.JwtSessionToken
	}
)

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *ShortUrl) FillEmptyFields() *ShortUrl {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(entity) }()
	}

	if entity.Properties == nil {
		entity.Properties = new(ShortUrlProperties)
	}

	if entity.Accesses == nil {
		entity.Accesses = new(ShortUrlAccesses)
	}

	entity.Accesses.FillEmptyFields()
	entity.Properties.FillEmptyFields()

	return entity
}

// ToModel - получение внешней модели.
func (entity *ShortUrl) ToModel() (model *models.ShortUrlInfo) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	entity.FillEmptyFields()

	model = &models.ShortUrlInfo{
		ID:        entity.ID,
		Source:    entity.Source,
		Reduction: entity.Reduction,

		Accesses: &models.ShortUrlAccesses{
			RolesID:       entity.Accesses.RolesID,
			PermissionsID: entity.Accesses.PermissionsID,
		},
		Properties: &models.ShortUrlInfoProperties{
			Type:                 entity.Properties.Type,
			NumberOfUses:         entity.Properties.NumberOfUses,
			RemainedNumberOfUses: entity.Properties.RemainedNumberOfUses,
			StartActive:          entity.Properties.StartActive,
			EndActive:            entity.Properties.EndActive,
			Active:               entity.Properties.Active,
		},
	}

	return
}

// ToRedisDbModel - получение модели базы данных redis.
func (entity *ShortUrl) ToRedisDbModel() (model *db_models.ShortUrlInfo) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &db_models.ShortUrlInfo{
		ID:        entity.ID,
		Source:    entity.Source,
		Reduction: entity.Reduction,

		Properties: &db_models.ShortUrlInfoProperties{
			Type:                 entity.Properties.Type,
			NumberOfUses:         entity.Properties.NumberOfUses,
			RemainedNumberOfUses: entity.Properties.RemainedNumberOfUses,
			StartActive:          entity.Properties.StartActive,
			EndActive:            entity.Properties.EndActive,
		},
	}

	return
}

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *ShortUrlProperties) FillEmptyFields() *ShortUrlProperties {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(entity) }()
	}

	return entity
}

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *ShortUrlAccesses) FillEmptyFields() *ShortUrlAccesses {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(entity) }()
	}

	if entity.RolesID == nil {
		entity.RolesID = make([]common_types.ID, 0)
	}

	if entity.PermissionsID == nil {
		entity.PermissionsID = make([]common_types.ID, 0)
	}

	return entity
}

// ToModel - получение внешней модели.
func (entity *ShortUrlUsageHistory) ToModel() (model *models.ShortUrlUsageHistoryInfo) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &models.ShortUrlUsageHistoryInfo{
		Status:    entity.Status,
		Timestamp: entity.Timestamp,
		TokenInfo: entity.TokenInfo,
	}

	return
}
