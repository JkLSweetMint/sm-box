package db_models

import (
	"github.com/google/uuid"
	common_types "sm-box/internal/common/types"
)

type (
	// Permission - модель базы данных прав пользователя.
	Permission struct {
		ID        common_types.ID `db:"id"`
		ProjectID common_types.ID `db:"project_id"`

		Name     string    `db:"name"`
		NameI18n uuid.UUID `db:"name_i18n"`

		Description     string    `db:"description"`
		DescriptionI18n uuid.UUID `db:"description_i18n"`

		IsSystem bool `db:"is_system"`
	}
)
