package db_models

type (
	// Language - модель базы данных языка локализации.
	Language struct {
		Code   string `db:"code"`
		Name   string `db:"name"`
		Active bool   `db:"active"`
	}
)
