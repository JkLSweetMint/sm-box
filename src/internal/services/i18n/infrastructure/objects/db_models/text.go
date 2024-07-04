package db_models

import "github.com/google/uuid"

type (
	// Text - текст.
	Text struct {
		ID       uuid.UUID `db:"id"`
		Language string    `db:"language"`
		Section  uuid.UUID `db:"section"`
		Key      string    `db:"key"`
		Value    string    `db:"value"`
	}

	// Dictionary - словарь локализации.
	Dictionary []*Text
)
