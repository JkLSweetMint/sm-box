package entities

import (
	"github.com/google/uuid"
	"sm-box/internal/services/i18n/objects/models"
	"sm-box/pkg/core/components/tracer"
	"strings"
)

type (
	// Text - текст.
	Text struct {
		ID       uuid.UUID
		Language string
		Section  uuid.UUID
		Key      string
		Value    string
	}

	// Dictionary - словарь локализации.
	Dictionary []*Text
)

// ToModel - получение модели.
func (entity Dictionary) ToModel() (model models.Dictionary) {
	// tracer
	{
		var trc = tracer.New(tracer.LevelEntity)

		trc.FunctionCall()
		defer func() { trc.FunctionCallFinished(model) }()
	}

	model = make(models.Dictionary)

	for _, text := range entity {
		var (
			keys    = strings.Split(text.Key, ".")
			store   = model
			keysLen = len(keys)
		)

		for i, key := range keys {
			if i+1 == keysLen {
				store[key] = text.Value
				break
			}

			if v, ok := store[key]; ok {
				if store, ok = v.(map[string]any); ok {
					continue
				} else {
					store[key] = make(map[string]any)
					store = store[key].(map[string]any)
				}
			} else {
				store[key] = make(map[string]any)
				store = store[key].(map[string]any)
			}
		}
	}

	return
}
