package db_models

type (
	// Language - язык.
	Language struct {
		Code   string `db:"code"`
		Name   string `db:"name"`
		Active bool   `db:"active"`
	}
)
