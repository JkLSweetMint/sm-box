package entities

import (
	"sm-box/internal/services/i18n/infrastructure/objects/models"
	"sm-box/pkg/core/components/tracer"
)

type (
	// Language - язык.
	Language struct {
		Name   string
		Code   string
		Active bool
	}
)

// ToModel - получение модели.
func (entity *Language) ToModel() (model *models.Language) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = &models.Language{
		Code:   entity.Code,
		Name:   entity.Name,
		Active: entity.Active,
	}

	return
}
