package sqlite3

import (
	"database/sql"
	"github.com/mattn/go-sqlite3"
	"regexp"
)

func regex(re, s string) (bool, error) {
	return regexp.MatchString(re, s)
}

func init() {
	sql.Register(driverName,
		&sqlite3.SQLiteDriver{
			ConnectHook: func(conn *sqlite3.SQLiteConn) error {
				return conn.RegisterFunc("regexp", regex, true)
			},
		})
}
