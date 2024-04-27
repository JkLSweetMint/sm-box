package db_models

import "sm-box/internal/common/types"

type (
	// Project - модель проекта для базы данных.
	Project struct {
		ID types.ID `db:"id"`

		Title       string `db:"title"`
		Description string `db:"description"`
	}

	// ProjectOwner - модель владельца проекта для базы данных.
	ProjectOwner struct {
		ProjectID types.ID `db:"project_id"`
		OwnerID   types.ID `db:"owner_id"`
	}
)
