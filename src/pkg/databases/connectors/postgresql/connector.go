package postgresql_connector

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Connector - описание коннектора для postgresql базы данных.
type Connector interface {
	Close() (err error)
	Ping() (err error)
	PingContext(ctx context.Context) (err error)
	Stats() (stats sql.DBStats)
	Prepare(query string) (stmt *sql.Stmt, err error)
	PrepareContext(ctx context.Context, query string) (stmt *sql.Stmt, err error)
	Preparex(query string) (stmt *sqlx.Stmt, err error)
	PreparexContext(ctx context.Context, query string) (stmt *sqlx.Stmt, err error)
	PrepareNamed(query string) (stmt *sqlx.NamedStmt, err error)
	PrepareNamedContext(ctx context.Context, query string) (stmt *sqlx.NamedStmt, err error)
	Exec(query string, args ...any) (res sql.Result, err error)
	ExecContext(ctx context.Context, query string, args ...any) (res sql.Result, err error)
	NamedExec(query string, arg interface{}) (res sql.Result, err error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (res sql.Result, err error)
	MustExec(query string, args ...interface{}) (res sql.Result)
	MustExecContext(ctx context.Context, query string, args ...interface{}) (res sql.Result)
	Query(query string, args ...any) (rows *sql.Rows, err error)
	QueryContext(ctx context.Context, query string, args ...any) (rows *sql.Rows, err error)
	Queryx(query string, args ...interface{}) (rows *sqlx.Rows, err error)
	QueryxContext(ctx context.Context, query string, args ...interface{}) (rows *sqlx.Rows, err error)
	NamedQuery(query string, arg interface{}) (rows *sqlx.Rows, err error)
	NamedQueryContext(ctx context.Context, query string, arg interface{}) (rows *sqlx.Rows, err error)
	QueryRow(query string, args ...any) (row *sql.Row)
	QueryRowContext(ctx context.Context, query string, args ...any) (row *sql.Row)
	QueryRowx(query string, args ...interface{}) (row *sqlx.Row)
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) (row *sqlx.Row)
	Begin() (tx *sql.Tx, err error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (tx *sql.Tx, err error)
	Beginx() (tx *sqlx.Tx, err error)
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (tx *sqlx.Tx, err error)
	MustBegin() (tx *sqlx.Tx)
	MustBeginTx(ctx context.Context, opts *sql.TxOptions) (tx *sqlx.Tx)
	Driver() (dr driver.Driver)
	DriverName() (dr string)
	Conn() (conn *sql.Conn, err error)
	ConnContext(ctx context.Context) (conn *sql.Conn, err error)
	Connx() (conn *sqlx.Conn, err error)
	ConnxContext(ctx context.Context) (conn *sqlx.Conn, err error)
	MapperFunc(mf func(string) string)
	Rebind(query string) (bound string)
	Unsafe() (db *sqlx.DB)
	BindNamed(query string, arg interface{}) (bound string, arglist []interface{}, err error)
	Select(dest interface{}, query string, args ...interface{}) (err error)
	Get(dest interface{}, query string, args ...interface{}) (err error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) (err error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) (err error)
}
