package constructors

import (
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/url_shortner/objects/types"
	"sm-box/pkg/core/components/tracer"
	"time"
)

type (
	// ShortUrl - конструктор для создания короткого url.
	ShortUrl struct {
		Source string

		Accesses   *ShortUrlAccesses
		Properties *ShortUrlProperties
	}

	// ShortUrlAccesses - конструктор доступов к короткому url.
	ShortUrlAccesses struct {
		RolesID       []common_types.ID
		PermissionsID []common_types.ID
	}

	// ShortUrlProperties - конструктор свойств короткого url.
	ShortUrlProperties struct {
		Type         types.ShortUrlType
		NumberOfUses int64
		StartActive  time.Time
		EndActive    time.Time
		Active       bool
	}
)

// FillEmptyFields - заполнение пустых полей конструктора.
func (constructor *ShortUrl) FillEmptyFields() *ShortUrl {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConstructor)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(constructor) }()
	}

	if constructor.Properties == nil {
		constructor.Properties = new(ShortUrlProperties)
	}

	if constructor.Accesses == nil {
		constructor.Accesses = new(ShortUrlAccesses)
	}

	constructor.Accesses.FillEmptyFields()
	constructor.Properties.FillEmptyFields()

	return constructor
}

// FillEmptyFields - заполнение пустых полей конструктора.
func (constructor *ShortUrlProperties) FillEmptyFields() *ShortUrlProperties {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConstructor)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(constructor) }()
	}

	return constructor
}

// FillEmptyFields - заполнение пустых полей конструктора.
func (constructor *ShortUrlAccesses) FillEmptyFields() *ShortUrlAccesses {
	// tracer
	{
		var trc = tracer.New(tracer.LevelConstructor)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(constructor) }()
	}

	if constructor.RolesID == nil {
		constructor.RolesID = make([]common_types.ID, 0)
	}

	if constructor.PermissionsID == nil {
		constructor.PermissionsID = make([]common_types.ID, 0)
	}

	return constructor
}
