package db_models

import "github.com/google/uuid"

type (
	// Text - модель базы данных языка локализации.
	Text struct {
		ID       uuid.UUID `db:"id"`
		Language string    `db:"language"`
		Section  uuid.UUID `db:"section"`
		Key      string    `db:"key"`
		Value    string    `db:"value"`
	}

	// Dictionary - модель базы данных словаря локализации.
	Dictionary []*Text
)
