package models

import (
	"sm-box/internal/common/types"
)

type (
	// ProjectInfo - информация о проекте.
	ProjectInfo struct {
		ID types.ID `json:"id" xml:"id,attr"`

		Name        string `json:"name"        xml:"Name"`
		Description string `json:"description" xml:"Description"`
		Version     string `json:"version"     xml:"Version"`
	}

	// ProjectList - список проектов.
	ProjectList []*ProjectInfo
)
