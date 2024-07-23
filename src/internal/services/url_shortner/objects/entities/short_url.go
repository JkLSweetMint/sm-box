package entities

import (
	"sm-box/internal/common/types"
	"sm-box/internal/services/url_shortner/objects/db_models"
	"sm-box/internal/services/url_shortner/objects/models"
	"sm-box/pkg/core/components/tracer"
	"time"
)

type (
	// ShortUrl - короткая ссылка.
	ShortUrl struct {
		ID        types.ID
		Source    string
		Reduction string

		Properties *ShortUrlProperties
	}

	// ShortUrlProperties - свойства короткой ссылке.
	ShortUrlProperties struct {
		Type         string
		NumberOfUses uint16
		StartActive  time.Time
		EndActive    time.Time
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

	model = &models.ShortUrlInfo{
		ID:        entity.ID,
		Source:    entity.Source,
		Reduction: entity.Reduction,

		Properties: &models.ShortUrlInfoProperties{
			Type:         entity.Properties.Type,
			NumberOfUses: entity.Properties.NumberOfUses,
			StartActive:  entity.Properties.StartActive,
			EndActive:    entity.Properties.EndActive,
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
			Type:         entity.Properties.Type,
			NumberOfUses: entity.Properties.NumberOfUses,
			StartActive:  entity.Properties.StartActive,
			EndActive:    entity.Properties.EndActive,
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
