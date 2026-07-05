package mrpostgres

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/mondegor/go-core/mrstorage"
	"github.com/mondegor/go-core/mrstorage/txisolevel"
)

type (
	transaction struct {
		tx       pgx.Tx
		isoLevel txisolevel.Enum
	}
)

// Query - отправляет SQL-запрос в рамках транзакции и возвращает результат в виде списка записей.
func (t *transaction) Query(ctx context.Context, sql string, args ...any) (mrstorage.DBQueryRows, error) {
	rows, err := t.tx.Query(ctx, sql, args...)
	if err != nil {
		return nil, wrapError(err)
	}

	return &queryRows{
		rows: rows,
	}, nil
}

// QueryRow - отправляет SQL-запрос в рамках транзакции и возвращает результат в виде одной записи.
func (t *transaction) QueryRow(ctx context.Context, sql string, args ...any) mrstorage.DBQueryRow {
	return &queryRow{
		row: t.tx.QueryRow(ctx, sql, args...),
	}
}

// Exec - отправляет SQL-запрос в рамках транзакции и исполняет его.
// Возвращает ErrEventStorageRecordsNotAffected,
// если команда Insert/Update/Delete не затронула ни одной строки.
func (t *transaction) Exec(ctx context.Context, sql string, args ...any) error {
	return wrapErrorCommandTag(t.tx.Exec(ctx, sql, args...))
}

// ExecRow - отправляет SQL-запрос в рамках транзакции и исполняет его, ожидая ровно одну затронутую запись.
// Возвращает ErrEventStorageNoRecordFound, если не затронуто ни одной записи,
// или ErrInternalStorageQueryFailed, если затронуто более одной.
func (t *transaction) ExecRow(ctx context.Context, sql string, args ...any) error {
	return wrapErrorExecRow(t.tx.Exec(ctx, sql, args...))
}

// ExecAffected - отправляет SQL-запрос в рамках транзакции, исполняет его и возвращает число затронутых строк.
func (t *transaction) ExecAffected(ctx context.Context, sql string, args ...any) (count int, err error) {
	return wrapErrorExecAffected(t.tx.Exec(ctx, sql, args...))
}

// mappingTxPgxOptions - преобразует внутренние настройки транзакции в настройки pgx.
func mappingTxPgxOptions(o mrstorage.TxOptions) (opts pgx.TxOptions) {
	switch o.IsoLevel {
	case txisolevel.Serializable:
		opts.IsoLevel = pgx.Serializable
	case txisolevel.RepeatableRead:
		opts.IsoLevel = pgx.RepeatableRead
	default:
	}

	return opts
}
