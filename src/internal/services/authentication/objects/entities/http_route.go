package entities

import (
	"encoding/json"
	common_types "sm-box/internal/common/types"
	"sm-box/internal/services/authentication/objects/db_models"
	"sm-box/pkg/core/components/tracer"
	"strings"
)

type (
	// HttpRoute - http маршрут системы.
	HttpRoute struct {
		ID common_types.ID `json:"id"`

		SystemName  string `json:"system_name"`
		Name        string `json:"name"`
		Description string `json:"description"`

		Protocols  []string `json:"protocols"`
		Method     string   `json:"method"`
		Path       string   `json:"path"`
		RegexpPath string   `json:"regexp_path"`

		Active bool `json:"active"`

		Accesses *HttpRouteAccesses `json:"accesses"`
	}

	// HttpRouteAccesses - доступы к http маршруту системы.
	HttpRouteAccesses struct {
		Roles       []common_types.ID `json:"roles"`
		Permissions []common_types.ID `json:"permissions"`
	}
)

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *HttpRoute) FillEmptyFields() *HttpRoute {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(entity) }()
	}

	if entity.Accesses == nil {
		entity.Accesses = new(HttpRouteAccesses)
	}

	entity.Accesses.FillEmptyFields()

	return entity
}

// ToDbModel - получение модели базы данных.
func (entity *HttpRoute) ToDbModel() (model *db_models.HttpRoute) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &db_models.HttpRoute{
		ID: entity.ID,

		SystemName:  entity.SystemName,
		Name:        entity.Name,
		Description: entity.Description,

		Protocols:  entity.Protocols,
		Method:     strings.ToUpper(entity.Method),
		Path:       entity.Path,
		RegexpPath: entity.RegexpPath,

		Active: entity.Active,
	}

	return
}

// MarshalBinary - упаковка структуры в бинарный формат.
func (entity *HttpRoute) MarshalBinary() ([]byte, error) {
	return json.Marshal(entity)
}

// UnmarshalBinary - распаковка структуры из бинарного формата.
func (entity *HttpRoute) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &entity)
}

// FillEmptyFields - заполнение пустых полей сущности.
func (entity *HttpRouteAccesses) FillEmptyFields() *HttpRouteAccesses {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(entity) }()
	}

	if entity.Roles == nil {
		entity.Roles = make([]common_types.ID, 0)
	}

	if entity.Permissions == nil {
		entity.Permissions = make([]common_types.ID, 0)
	}

	return entity
}
