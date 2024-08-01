package constructors

import (
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/url_shortner/objects/types"
	"sm-box/pkg/core/components/tracer"
	"time"
)

type (
	// ShortUrl - конструктор короткой ссылки.
	ShortUrl struct {
		Source string

		Accesses   *ShortUrlAccesses
		Properties *ShortUrlProperties
	}

	// ShortUrlAccesses - конструктор информации по доступам к короткому url.
	ShortUrlAccesses struct {
		RolesID       []common_types.ID
		PermissionsID []common_types.ID
	}

	// ShortUrlProperties - конструктор свойств короткой ссылке.
	ShortUrlProperties struct {
		Type         types.ShortUrlType
		NumberOfUses int64
		StartActive  time.Time
		EndActive    time.Time
		Active       bool
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
