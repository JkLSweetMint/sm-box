package models

import "sm-box/internal/app/infrastructure/types"

type (
	// ProjectInfo - информация о проекте.
	ProjectInfo struct {
		ID types.ID `json:"id" yaml:"ID" xml:"id,attr"`

		Title       string `json:"title"       yaml:"Title"       xml:"Title"`
		Description string `json:"description" yaml:"Description" xml:"Description"`

		Owners *ProjectInfoOwner `json:"owners" yaml:"Owners" xml:"Owners"`
	}

	// ProjectInfoOwner - информация о владельце проекта.
	ProjectInfoOwner struct {
		*UserInfo
	}
)
