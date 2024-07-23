package models

type (
	// Language - внешняя модель языка локализации.
	Language struct {
		Code   string `json:"code"   xml:"code,attr"`
		Name   string `json:"name"   xml:"name,attr"`
		Active bool   `json:"active" xml:"active,attr"`
	}
)
