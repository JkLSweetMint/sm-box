package models

import "sm-box/internal/common/types"

type (
	// Project - проект.
	Project struct {
		ID types.ID

		Title       string
		Description string

		Owners ProjectOwners
	}

	// ProjectOwners - владельцы проекта.
	ProjectOwners []*ProjectOwner

	// ProjectOwner - владелец проекта.
	ProjectOwner struct {
		*User
	}
)
