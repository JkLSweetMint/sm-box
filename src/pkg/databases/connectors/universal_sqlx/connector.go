package universal_sqlx

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"github.com/jmoiron/sqlx"
	"sm-box/pkg/core/components/logger"
	"sm-box/pkg/core/components/tracer"
)

// UniversalConnector - универсальный коннектор для работы с sql базами данных.
type UniversalConnector struct {
	DB *sqlx.DB

	Components  *Components
	Concurrency *Concurrency
}

// Concurrency - управление конкурентностью.
type Concurrency struct {
	Ctx     context.Context
	GlobCtx context.Context
	Cancel  context.CancelFunc
}

// Components - компоненты коннектора.
type Components struct {
	Logger logger.Logger
}

// Close - закрывает базу данных и предотвращает запуск новых запросов,
// затем ожидает завершения всех запросов, которые начали обрабатываться на сервере.
//
// Закрытие базы данных происходит редко, так как дескриптор базы данных должен быть
// долговечным и совместно использоваться многими подпрограммами.
func (connector *UniversalConnector) Close() (err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall()
		defer func() { trace.Error(err).FunctionCallFinished() }()
	}

	connector.Concurrency.Cancel()

	if err = connector.DB.Close(); err != nil {
		connector.Components.Logger.Error().
			Format("Error closing the database connection: '%s'. ", err).Write()
		return
	}

	return
}

// Ping проверяет что соединение с базой данных все еще работает, при необходимости устанавливает соединение.
//
// Использует глобальный context системы.
func (connector *UniversalConnector) Ping() (err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall()
		defer func() { trace.Error(err).FunctionCallFinished() }()
	}

	if err = connector.DB.PingContext(connector.Concurrency.Ctx); err != nil {
		connector.Components.Logger.Error().
			Format("Database query execution error: '%s'. ", err).Write()
		return
	}

	return
}

// PingContext - проверяет, что соединение с базой данных все еще работает, устанавливая соединение при необходимости.
func (connector *UniversalConnector) PingContext(ctx context.Context) (err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall()
		defer func() { trace.Error(err).FunctionCallFinished(ctx) }()
	}

	if err = connector.DB.PingContext(ctx); err != nil {
		connector.Components.Logger.Error().
			Format("Database query execution error: '%s'. ", err).Write()
		return
	}

	return
}

// Stats - возвращает статистику базы данных.
func (connector *UniversalConnector) Stats() (stats sql.DBStats) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall()
		defer func() { trace.FunctionCallFinished(stats) }()
	}

	stats = connector.DB.Stats()

	return
}

// Prepare - создает подготовленную инструкцию для последующих запросов или выполнений.
// Несколько запросов или выполнений могут выполняться одновременно из возвращаемой инструкции.
// Вызывающая сторона должна вызвать метод Close инструкции когда она больше не нужна.
//
// Используется для подготовки инструкции, а не для выполнения инструкции.
// Использует глобальный context системы.
func (connector *UniversalConnector) Prepare(query string) (stmt *sql.Stmt, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(query)
		defer func() { trace.Error(err).FunctionCallFinished(stmt) }()
	}

	if stmt, err = connector.DB.PrepareContext(connector.Concurrency.Ctx, query); err != nil {
		connector.Components.Logger.Error().
			Format("Database query execution error: '%s'. ", err).Write()
		return
	}

	return
}

// PrepareContext - создает подготовленную инструкцию для последующих запросов или выполнений.
// Несколько запросов или выполнений могут выполняться одновременно из возвращаемой инструкции.
// Вызывающая сторона должна вызвать метод Close инструкции когда она больше не нужна.
//
// Используется для подготовки инструкции, а не для ыполнения инструкции.
func (connector *UniversalConnector) PrepareContext(ctx context.Context, query string) (stmt *sql.Stmt, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(ctx, query)
		defer func() { trace.Error(err).FunctionCallFinished(stmt) }()
	}

	if stmt, err = connector.DB.PrepareContext(ctx, query); err != nil {
		connector.Components.Logger.Error().
			Format("Database query execution error: '%s'. ", err).Write()
		return
	}

	return
}

// Preparex - создает подготовленную инструкцию для последующих запросов или выполнений.
// Несколько запросов или выполнений могут выполняться одновременно из возвращаемой инструкции.
// Вызывающая сторона должна вызвать метод Close инструкции когда она больше не нужна.
//
// Используется для подготовки инструкции, а не для выполнения инструкции.
// Использует глобальный context системы.
// Возвращает sqlx.Stmt вместо sql.Stmt.
func (connector *UniversalConnector) Preparex(query string) (stmt *sqlx.Stmt, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(query)
		defer func() { trace.Error(err).FunctionCallFinished(stmt) }()
	}

	if stmt, err = connector.DB.PreparexContext(connector.Concurrency.Ctx, query); err != nil {
		connector.Components.Logger.Error().
			Format("Database query execution error: '%s'. ", err).Write()
		return
	}

	return
}

// PreparexContext - создает подготовленную инструкцию для последующих запросов или выполнений.
// Несколько запросов или выполнений могут выполняться одновременно из возвращаемой инструкции.
// Вызывающая сторона должна вызвать метод Close инструкции когда она больше не нужна.
//
// Используется для подготовки инструкции, а не для выполнения инструкции.
// Возвращает sqlx.Stmt вместо sql.Stmt.
func (connector *UniversalConnector) PreparexContext(ctx context.Context, query string) (stmt *sqlx.Stmt, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(ctx, query)
		defer func() { trace.Error(err).FunctionCallFinished(stmt) }()
	}

	if stmt, err = connector.DB.PreparexContext(ctx, query); err != nil {
		connector.Components.Logger.Error().
			Format("Database query execution error: '%s'. ", err).Write()
		return
	}

	return
}

// PrepareNamed - создает подготовленную именнованную инструкцию для последующих запросов или выполнений.
// Несколько запросов или выполнений могут выполняться одновременно из возвращаемой инструкции.
// Вызывающая сторона должна вызвать метод Close инструкции когда она больше не нужна.
//
// Используется для подготовки инструкции, а не для выполнения инструкции.
// Использует глобальный context системы.
func (connector *UniversalConnector) PrepareNamed(query string) (stmt *sqlx.NamedStmt, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(query)
		defer func() { trace.Error(err).FunctionCallFinished(stmt) }()
	}

	if stmt, err = connector.DB.PrepareNamedContext(connector.Concurrency.Ctx, query); err != nil {
		connector.Components.Logger.Error().
			Format("Database query execution error: '%s'. ", err).Write()
		return
	}

	return
}

// PrepareNamedContext - создает подготовленную именнованную инструкцию для последующих запросов или выполнений.
// Несколько запросов или выполнений могут выполняться одновременно из возвращаемой инструкции.
// Вызывающая сторона должна вызвать метод Close инструкции когда она больше не нужна.
//
// Используется для подготовки инструкции, а не для выполнения инструкции.
func (connector *UniversalConnector) PrepareNamedContext(ctx context.Context, query string) (stmt *sqlx.NamedStmt, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(ctx, query)
		defer func() { trace.Error(err).FunctionCallFinished(stmt) }()
	}

	if stmt, err = connector.DB.PrepareNamedContext(ctx, query); err != nil {
		connector.Components.Logger.Error().
			Format("Database query execution error: '%s'. ", err).Write()
		return
	}

	return
}

// Exec - выполняет запрос, не возвращая никаких строк.
// Аргументы предназначены для любых параметров-заполнителей в запросе.
//
// Использует глобальный context системы.
func (connector *UniversalConnector) Exec(query string, args ...any) (res sql.Result, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(query, args)
		defer func() { trace.Error(err).FunctionCallFinished(res) }()
	}

	if res, err = connector.DB.ExecContext(connector.Concurrency.Ctx, query, args...); err != nil {
		connector.Components.Logger.Error().
			Format("Database query execution error: '%s'. ", err).Write()
		return
	}

	return
}

// ExecContext - выполняет запрос, не возвращая никаких строк.
// Аргументы предназначены для любых параметров-заполнителей в запросе.
func (connector *UniversalConnector) ExecContext(ctx context.Context, query string, args ...any) (res sql.Result, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(ctx, query, args)
		defer func() { trace.Error(err).FunctionCallFinished(res) }()
	}

	if res, err = connector.DB.ExecContext(ctx, query, args...); err != nil {
		connector.Components.Logger.Error().
			Format("Database query execution error: '%s'. ", err).Write()
		return
	}

	return
}

// NamedExec - выполняет именнованный запрос, не возвращая никаких строк.
// Аргументы предназначены для любых параметров-заполнителей в запросе.
//
// Использует глобальный context системы.
func (connector *UniversalConnector) NamedExec(query string, arg interface{}) (res sql.Result, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(query, arg)
		defer func() { trace.Error(err).FunctionCallFinished(res) }()
	}

	if res, err = connector.DB.NamedExecContext(connector.Concurrency.Ctx, query, arg); err != nil {
		connector.Components.Logger.Error().
			Format("Database query execution error: '%s'. ", err).Write()
		return
	}

	return
}

// NamedExecContext - выполняет именнованный запрос, не возвращая никаких строк.
// Аргументы предназначены для любых параметров-заполнителей в запросе.
func (connector *UniversalConnector) NamedExecContext(ctx context.Context, query string, arg interface{}) (res sql.Result, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(ctx, query, arg)
		defer func() { trace.Error(err).FunctionCallFinished(res) }()
	}

	if res, err = connector.DB.NamedExecContext(ctx, query, arg); err != nil {
		connector.Components.Logger.Error().
			Format("Database query execution error: '%s'. ", err).Write()
		return
	}

	return
}

// MustExec - выполняет запрос, не возвращая никаких строк.
// Аргументы предназначены для любых параметров-заполнителей в запросе.
//
// Использует глобальный context системы.
// Взовет panic в случае ошибки.
func (connector *UniversalConnector) MustExec(query string, args ...interface{}) (res sql.Result) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(query, args)
		defer func() { trace.FunctionCallFinished(res) }()
	}

	res = connector.DB.MustExecContext(connector.Concurrency.Ctx, query, args...)

	return
}

// MustExecContext - выполняет запрос, не возвращая никаких строк.
// Аргументы предназначены для любых параметров-заполнителей в запросе.
// Взовет panic в случае ошибки.
func (connector *UniversalConnector) MustExecContext(ctx context.Context, query string, args ...interface{}) (res sql.Result) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(ctx, query, args)
		defer func() { trace.FunctionCallFinished(res) }()
	}

	res = connector.DB.MustExecContext(ctx, query, args...)

	return
}

// Query - выполняет запрос, который возвращает строки, обычно SELECT.
// Аргументы предназначены для любых параметров-заполнителей в запросе.
//
// Использует глобальный context системы.
func (connector *UniversalConnector) Query(query string, args ...any) (rows *sql.Rows, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(query, args)
		defer func() { trace.Error(err).FunctionCallFinished(rows) }()
	}

	if rows, err = connector.DB.QueryContext(connector.Concurrency.Ctx, query, args...); err != nil {
		connector.Components.Logger.Error().
			Format("Database query execution error: '%s'. ", err).Write()
		return
	}

	return
}

// QueryContext - выполняет запрос, который возвращает строки, обычно SELECT.
// Аргументы предназначены для любых параметров-заполнителей в запросе.
func (connector *UniversalConnector) QueryContext(ctx context.Context, query string, args ...any) (rows *sql.Rows, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(ctx, query, args)
		defer func() { trace.Error(err).FunctionCallFinished(rows) }()
	}

	if rows, err = connector.DB.QueryContext(ctx, query, args...); err != nil {
		connector.Components.Logger.Error().
			Format("Database query execution error: '%s'. ", err).Write()
		return
	}

	return
}

// Queryx - выполняет запрос, который возвращает строки, обычно SELECT.
// Аргументы предназначены для любых параметров-заполнителей в запросе.
//
// Использует глобальный context системы.
// Возвращает sqlx.Rows вместо sql.Rows.
func (connector *UniversalConnector) Queryx(query string, args ...interface{}) (rows *sqlx.Rows, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(query, args)
		defer func() { trace.Error(err).FunctionCallFinished(rows) }()
	}

	if rows, err = connector.DB.QueryxContext(connector.Concurrency.Ctx, query, args...); err != nil {
		connector.Components.Logger.Error().
			Format("Database query execution error: '%s'. ", err).Write()
		return
	}

	return
}

// QueryxContext - выполняет запрос, который возвращает строки, обычно SELECT.
// Аргументы предназначены для любых параметров-заполнителей в запросе.
//
// Возвращает sqlx.Rows вместо sql.Rows.
func (connector *UniversalConnector) QueryxContext(ctx context.Context, query string, args ...interface{}) (rows *sqlx.Rows, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(ctx, query, args)
		defer func() { trace.Error(err).FunctionCallFinished(rows) }()
	}

	if rows, err = connector.DB.QueryxContext(ctx, query, args...); err != nil {
		connector.Components.Logger.Error().
			Format("Database query execution error: '%s'. ", err).Write()
		return
	}

	return
}

// NamedQuery - выполняет именнованный запрос, который возвращает строки, обычно SELECT.
// Аргументы предназначены для любых параметров-заполнителей в запросе.
//
// Использует глобальный context системы.
// Возвращает sqlx.Rows вместо sql.Rows.
func (connector *UniversalConnector) NamedQuery(query string, arg interface{}) (rows *sqlx.Rows, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(query, arg)
		defer func() { trace.Error(err).FunctionCallFinished(rows) }()
	}

	if rows, err = connector.DB.NamedQueryContext(connector.Concurrency.Ctx, query, arg); err != nil {
		connector.Components.Logger.Error().
			Format("Database query execution error: '%s'. ", err).Write()
		return
	}

	return
}

// NamedQueryContext - выполняет именнованный запрос, который возвращает строки, обычно SELECT.
// Аргументы предназначены для любых параметров-заполнителей в запросе.
//
// Возвращает sqlx.Rows вместо sql.Rows.
func (connector *UniversalConnector) NamedQueryContext(ctx context.Context, query string, arg interface{}) (rows *sqlx.Rows, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(ctx, query, arg)
		defer func() { trace.Error(err).FunctionCallFinished(rows) }()
	}

	if rows, err = connector.DB.NamedQueryContext(ctx, query, arg); err != nil {
		connector.Components.Logger.Error().
			Format("Database query execution error: '%s'. ", err).Write()
		return
	}

	return
}

// QueryRow - выполняет запрос, который возвращает строки, обычно SELECT.
// Аргументы предназначены для любых параметров-заполнителей в запросе.
//
// Использует глобальный context системы.
func (connector *UniversalConnector) QueryRow(query string, args ...any) (row *sql.Row) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(query, args)
		defer func() { trace.FunctionCallFinished(row) }()
	}

	row = connector.DB.QueryRowContext(connector.Concurrency.Ctx, query, args...)

	return
}

// QueryRowContext - выполняет запрос, который, как ожидается, вернет не более одной строки.
// Всегда возвращает значение, отличное от нуля. Ошибки откладываются до тех пор, пока не будет вызван метод сканирования Row.
// Если в запросе не выбрано ни одной строки, проверка *строки вернет значение ErrNoRows.
// В противном случае проверка *строки сканирует первую выбранную строку и отбрасывает остальные.
func (connector *UniversalConnector) QueryRowContext(ctx context.Context, query string, args ...any) (row *sql.Row) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(query, args)
		defer func() { trace.FunctionCallFinished(row) }()
	}

	row = connector.DB.QueryRowContext(ctx, query, args...)

	return
}

// QueryRowx - выполняет запрос, который возвращает строки, обычно SELECT.
// Аргументы предназначены для любых параметров-заполнителей в запросе.
//
// Использует глобальный context системы.
// Возвращает sqlx.Row вместо sql.Row.
func (connector *UniversalConnector) QueryRowx(query string, args ...interface{}) (row *sqlx.Row) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(query, args)
		defer func() { trace.FunctionCallFinished(row) }()
	}

	row = connector.DB.QueryRowxContext(connector.Concurrency.Ctx, query, args...)

	return
}

// QueryRowxContext - выполняет запрос, который, как ожидается, вернет не более одной строки.
// Всегда возвращает значение, отличное от нуля. Ошибки откладываются до тех пор, пока не будет вызван метод сканирования Row.
// Если в запросе не выбрано ни одной строки, проверка *строки вернет значение ErrNoRows.
// В противном случае проверка *строки сканирует первую выбранную строку и отбрасывает остальные.
//
// Возвращает sqlx.Row вместо sql.Row.
func (connector *UniversalConnector) QueryRowxContext(ctx context.Context, query string, args ...interface{}) (row *sqlx.Row) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(query, args)
		defer func() { trace.FunctionCallFinished(row) }()
	}

	row = connector.DB.QueryRowxContext(ctx, query, args...)

	return
}

// Begin - запускает транзакцию. Уровень изоляции по умолчанию зависит от драйвера.
//
// Использует глобальный context системы.
func (connector *UniversalConnector) Begin() (tx *sql.Tx, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall()
		defer func() { trace.Error(err).FunctionCallFinished(tx) }()
	}

	if tx, err = connector.DB.BeginTx(connector.Concurrency.Ctx, nil); err != nil {
		connector.Components.Logger.Error().
			Format("Transaction execution error: '%s'. ", err).Write()
		return
	}

	return
}

// BeginTx запускает транзакцию.
//
// Предоставленный контекст используется до тех пор, пока транзакция не будет зафиксирована или откатана.
// Если контекст отменен, пакет sql выполнит откат транзакции.
// транзакция. Tx.Commit вернет ошибку, если контекст, предоставленный для Begin Tx, отменен.
//
// Предоставленные TxOptions необязательны и могут быть равны нулю, если следует использовать значения по умолчанию.
// Если используется уровень изоляции, отличный от заданного по умолчанию, который драйвер не поддерживает,
// будет возвращена ошибка.
func (connector *UniversalConnector) BeginTx(ctx context.Context, opts *sql.TxOptions) (tx *sql.Tx, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(ctx, opts)
		defer func() { trace.Error(err).FunctionCallFinished(tx) }()
	}

	if tx, err = connector.DB.BeginTx(ctx, opts); err != nil {
		connector.Components.Logger.Error().
			Format("Transaction execution error: '%s'. ", err).Write()
		return
	}

	return
}

// Beginx - запускает транзакцию. Уровень изоляции по умолчанию зависит от драйвера.
//
// Использует глобальный context системы.
// Возвращает sqlx.Tx вместо sql.Tx.
func (connector *UniversalConnector) Beginx() (tx *sqlx.Tx, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall()
		defer func() { trace.Error(err).FunctionCallFinished(tx) }()
	}

	if tx, err = connector.DB.BeginTxx(connector.Concurrency.Ctx, nil); err != nil {
		connector.Components.Logger.Error().
			Format("Transaction execution error: '%s'. ", err).Write()
		return
	}

	return
}

// BeginTxx запускает транзакцию.
//
// Предоставленный контекст используется до тех пор, пока транзакция не будет зафиксирована или откатана.
// Если контекст отменен, пакет sql выполнит откат транзакции.
// транзакция. Tx.Commit вернет ошибку, если контекст, предоставленный для Begin Tx, отменен.
//
// Предоставленные TxOptions необязательны и могут быть равны нулю, если следует использовать значения по умолчанию.
// Если используется уровень изоляции, отличный от заданного по умолчанию, который драйвер не поддерживает,
// будет возвращена ошибка.
//
// Возвращает sqlx.Tx вместо sql.Tx.
func (connector *UniversalConnector) BeginTxx(ctx context.Context, opts *sql.TxOptions) (tx *sqlx.Tx, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(ctx, opts)
		defer func() { trace.Error(err).FunctionCallFinished(tx) }()
	}

	if tx, err = connector.DB.BeginTxx(ctx, opts); err != nil {
		connector.Components.Logger.Error().
			Format("Transaction execution error: '%s'. ", err).Write()
		return
	}

	return
}

// MustBegin - запускает транзакцию. Уровень изоляции по умолчанию зависит от драйвера.
//
// Использует глобальный context системы.
// Возвращает sqlx.Tx вместо sql.Tx.
// Взовет panic в случае ошибки.
func (connector *UniversalConnector) MustBegin() (tx *sqlx.Tx) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall()
		defer func() { trace.FunctionCallFinished(tx) }()
	}

	tx = connector.DB.MustBeginTx(connector.Concurrency.Ctx, nil)

	return
}

// MustBeginTx запускает транзакцию.
//
// Предоставленный контекст используется до тех пор, пока транзакция не будет зафиксирована или откатана.
// Если контекст отменен, пакет sql выполнит откат транзакции.
// транзакция. Tx.Commit вернет ошибку, если контекст, предоставленный для Begin Tx, отменен.
//
// Предоставленные TxOptions необязательны и могут быть равны нулю, если следует использовать значения по умолчанию.
// Если используется уровень изоляции, отличный от заданного по умолчанию, который драйвер не поддерживает,
// будет возвращена ошибка.
//
// Взовет panic в случае ошибки.
func (connector *UniversalConnector) MustBeginTx(ctx context.Context, opts *sql.TxOptions) (tx *sqlx.Tx) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(ctx, opts)
		defer func() { trace.FunctionCallFinished(tx) }()
	}

	tx = connector.DB.MustBeginTx(ctx, opts)

	return
}

// Driver - возвращает драйвер базы данных.
func (connector *UniversalConnector) Driver() (dr driver.Driver) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall()
		defer func() { trace.FunctionCallFinished(dr) }()
	}

	dr = connector.DB.Driver()

	return
}

// DriverName - возвращает название драйвера базы данных.
func (connector *UniversalConnector) DriverName() (dr string) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall()
		defer func() { trace.FunctionCallFinished(dr) }()
	}

	dr = connector.DB.DriverName()

	return
}

// Conn - возвращает одно соединение, либо открывая новое соединение,
// либо возвращая существующее соединение из пула подключений. Conn будет
// блокироваться до тех пор, пока либо соединение не будет возвращено, либо ctx не будет отменен.
// Запросы, выполняемые в том же Conn, будут выполняться в том же сеансе базы данных.
//
// Каждый Conn должен быть возвращен в пул базы данных после использования путем вызова Conn.Close.
func (connector *UniversalConnector) Conn() (conn *sql.Conn, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall()
		defer func() { trace.Error(err).FunctionCallFinished(connector) }()
	}

	if conn, err = connector.DB.Conn(connector.Concurrency.Ctx); err != nil {
		connector.Components.Logger.Error().
			Format("Error getting a connection to the database: '%s'. ", err).Write()
		return
	}

	return
}

// ConnContext - возвращает одно соединение, либо открывая новое соединение,
// либо возвращая существующее соединение из пула подключений. Conn будет
// блокироваться до тех пор, пока либо соединение не будет возвращено, либо ctx не будет отменен.
// Запросы, выполняемые в том же Conn, будут выполняться в том же сеансе базы данных.
//
// Каждый Conn должен быть возвращен в пул базы данных после использования путем вызова Conn.Close.
func (connector *UniversalConnector) ConnContext(ctx context.Context) (conn *sql.Conn, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(ctx)
		defer func() { trace.Error(err).FunctionCallFinished(connector) }()
	}

	if conn, err = connector.DB.Conn(ctx); err != nil {
		connector.Components.Logger.Error().
			Format("Error getting a connection to the database: '%s'. ", err).Write()
		return
	}

	return
}

// Connx - возвращает одно соединение, либо открывая новое соединение,
// либо возвращая существующее соединение из пула подключений. Conn будет
// блокироваться до тех пор, пока либо соединение не будет возвращено, либо ctx не будет отменен.
// Запросы, выполняемые в том же Conn, будут выполняться в том же сеансе базы данных.
//
// Каждый Conn должен быть возвращен в пул базы данных после использования путем вызова Conn.Close.
//
// Использует глобальный context системы.
func (connector *UniversalConnector) Connx() (conn *sqlx.Conn, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall()
		defer func() { trace.Error(err).FunctionCallFinished(connector) }()
	}

	if conn, err = connector.DB.Connx(connector.Concurrency.Ctx); err != nil {
		connector.Components.Logger.Error().
			Format("Error getting a connection to the database: '%s'. ", err).Write()
		return
	}

	return
}

// ConnxContext - возвращает одно соединение, либо открывая новое соединение,
// либо возвращая существующее соединение из пула подключений. Conn будет
// блокироваться до тех пор, пока либо соединение не будет возвращено, либо ctx не будет отменен.
// Запросы, выполняемые в том же Conn, будут выполняться в том же сеансе базы данных.
//
// Каждый Conn должен быть возвращен в пул базы данных после использования путем вызова Conn.Close.
func (connector *UniversalConnector) ConnxContext(ctx context.Context) (conn *sqlx.Conn, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(ctx)
		defer func() { trace.Error(err).FunctionCallFinished(connector) }()
	}

	if conn, err = connector.DB.Connx(ctx); err != nil {
		connector.Components.Logger.Error().
			Format("Error getting a connection to the database: '%s'. ", err).Write()
		return
	}

	return
}

// MapperFunc - устанавливает новый сопоставитель для этой базы данных, используя тег sql struct по умолчанию
// и предоставленную функцию сопоставления.
func (connector *UniversalConnector) MapperFunc(mf func(string) string) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(mf)
		defer func() { trace.FunctionCallFinished() }()
	}

	connector.DB.MapperFunc(mf)

	return
}

// Rebind - преобразует запрос из QUESTION в тип bind var драйвера базы данных.
func (connector *UniversalConnector) Rebind(query string) (bound string) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(query)
		defer func() { trace.FunctionCallFinished(bound) }()
	}

	bound = connector.DB.Rebind(query)

	return
}

// Unsafe - возвращает версию DB, которая автоматически завершит сканирование, когда
// столбцы в результате SQL не будут иметь полей в целевой структуре.
// sql.Stmt и sqlx.Tx, созданные из этой базы данных, унаследуют ее
// поведение безопасности.
func (connector *UniversalConnector) Unsafe() (db *sqlx.DB) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall()
		defer func() { trace.FunctionCallFinished(db) }()
	}

	db = connector.DB.Unsafe()

	return
}

// BindNamed - связывает запрос, используя тип bind var драйвера БД.
func (connector *UniversalConnector) BindNamed(query string, arg interface{}) (bound string, arglist []interface{}, err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(query, arg)
		defer func() { trace.Error(err).FunctionCallFinished(bound, arglist) }()
	}

	if bound, arglist, err = connector.DB.BindNamed(query, arg); err != nil {
		connector.Components.Logger.Error().
			Format("Error getting a connection to the database: '%s'. ", err).Write()
		return
	}

	return
}

// Select - выборка используя эту базу данных.
// Любые параметры-заполнители заменяются указанными аргументами.
//
// Использует глобальный context системы.
func (connector *UniversalConnector) Select(dest interface{}, query string, args ...interface{}) (err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(dest, query, args)
		defer func() { trace.Error(err).FunctionCallFinished() }()
	}

	if err = connector.DB.SelectContext(connector.Concurrency.Ctx, dest, query, args...); err != nil {
		connector.Components.Logger.Error().
			Format("Error getting a connection to the database: '%s'. ", err).Write()
		return
	}

	return
}

// Get - получение с помощью этой базы данных.
// Любые параметры-заполнители заменяются указанными аргументами.
// Возвращается ошибка, если результирующий набор пуст.
//
// Использует глобальный context системы.
func (connector *UniversalConnector) Get(dest interface{}, query string, args ...interface{}) (err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(dest, query, args)
		defer func() { trace.Error(err).FunctionCallFinished() }()
	}

	if err = connector.DB.GetContext(connector.Concurrency.Ctx, dest, query, args...); err != nil {
		connector.Components.Logger.Error().
			Format("Error getting a connection to the database: '%s'. ", err).Write()
		return
	}

	return
}

// SelectContext - выборка используя эту базу данных.
// Любые параметры-заполнители заменяются указанными аргументами.
func (connector *UniversalConnector) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) (err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(ctx, dest, query, args)
		defer func() { trace.Error(err).FunctionCallFinished() }()
	}

	if err = connector.DB.SelectContext(ctx, dest, query, args...); err != nil {
		connector.Components.Logger.Error().
			Format("Error getting a connection to the database: '%s'. ", err).Write()
		return
	}

	return
}

// GetContext - получение с помощью этой базы данных.
// Любые параметры-заполнители заменяются указанными аргументами.
// Возвращается ошибка, если результирующий набор пуст.
func (connector *UniversalConnector) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) (err error) {
	// tracer
	{
		var trace = tracer.New(tracer.LevelDatabaseConnector)

		trace.FunctionCall(ctx, dest, query, args)
		defer func() { trace.Error(err).FunctionCallFinished() }()
	}

	if err = connector.DB.GetContext(ctx, dest, query, args...); err != nil {
		connector.Components.Logger.Error().
			Format("Error getting a connection to the database: '%s'. ", err).Write()
		return
	}

	return
}
