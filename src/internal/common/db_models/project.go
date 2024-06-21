package db_models

import "sm-box/internal/common/types"

type (
	// Project - модель проекта для базы данных.
	Project struct {
		ID      types.ID `db:"id"`
		UUID    string   `db:"uuid"`
		OwnerID types.ID `db:"owner_id"`

		Name        string `db:"name"`
		Description string `db:"description"`
		Version     string `db:"version"`
	}

	// ProjectOwner - модель владельца проекта для базы данных.
	ProjectOwner struct {
		ProjectID types.ID `db:"project_id"`
		OwnerID   types.ID `db:"owner_id"`
	}
)
