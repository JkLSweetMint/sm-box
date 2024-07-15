package db_models

import "sm-box/internal/common/types"

type (
	// Project - модель проекта для базы данных.
	Project struct {
		ID types.ID `db:"id"`

		Name        string `db:"name"`
		Description string `db:"description"`
		Version     string `db:"version"`
	}
)