package db_models

import common_types "sm-box/internal/common/types"

type (
	// Project - модель проекта для базы данных.
	Project struct {
		ID common_types.ID `db:"id"`

		Name        string `db:"name"`
		Description string `db:"description"`
		Version     string `db:"version"`
	}
)
