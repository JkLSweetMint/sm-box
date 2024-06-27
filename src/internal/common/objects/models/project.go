package models

import "sm-box/internal/common/types"

type (
	// ProjectInfo - информация о проекте.
	ProjectInfo struct {
		ID types.ID `json:"id" xml:"id,attr"`

		Title       string `json:"title"       xml:"Title"`
		Description string `json:"description" xml:"Description"`

		Owners *ProjectInfoOwner `json:"owners" xml:"Owners"`
	}

	// ProjectInfoOwner - информация о владельце проекта.
	ProjectInfoOwner struct {
		*UserInfo
	}
)
