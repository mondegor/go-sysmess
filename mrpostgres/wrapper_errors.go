package mrpostgres

import (
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/mondegor/go-sysmess/errors"
)

// wrapError - обёртывает ошибки PostgreSQL в стандартные ошибки приложения:
// распознанные ошибки делегируются в wrapPgError, остальные оборачиваются в общую Internal-ошибку.
// Может возвращать следующие ошибки:
//   - ErrInternalStorageDuplicateKeyViolation - при нарушении уникальности ключа (код 23505);
//   - ErrSystemStorageUnexpectedEOF - при неожиданном конце соединения;
//   - ErrInternalStorageQueryFailed - при нераспознанной *pgconn.PgError;
//   - InternalError - для прочих ошибок.
func wrapError(err error) error {
	if ok, wrappedErr := wrapPgError(err); ok {
		if wrappedErr != nil {
			return wrappedErr
		}

		return errors.ErrInternalStorageQueryFailed.Wrap(err, "source", connectionName)
	}

	return errors.WrapInternalError(err, "failed", "source", connectionName)
}

// wrapPgError - преобразует специфичные ошибки PostgreSQL в стандартные ошибки приложения.
// Возвращает ok=true, если ошибка является *pgconn.PgError либо имеет текст "unexpected EOF".
// Может возвращать следующие ошибки:
//   - ErrInternalStorageDuplicateKeyViolation - при нарушении уникальности ключа (код 23505);
//   - ErrSystemStorageUnexpectedEOF - при тексте ошибки "unexpected EOF".
func wrapPgError(err error) (ok bool, wrappedErr error) {
	if e := (*pgconn.PgError)(nil); errors.As(err, &e) {
		// Code: 23505 duplicate key value violates unique constraint
		if e.Code == "23505" {
			return true, errors.ErrInternalStorageDuplicateKeyViolation.Wrap(err)
		}

		return true, nil
	}

	if err.Error() == "unexpected EOF" {
		return true, errors.ErrSystemStorageUnexpectedEOF.New("source", connectionName)
	}

	return false, nil
}

// wrapErrorFetchDataFailed - обёртывает ошибки получения данных из БД.
// Может возвращать следующие ошибки:
//   - ErrInternalStorageDuplicateKeyViolation или ErrSystemStorageUnexpectedEOF - если ошибка распознана wrapPgError;
//   - ErrInternalStorageFetchDataFailed - для прочих ошибок.
func wrapErrorFetchDataFailed(err error) error {
	if _, wrappedErr := wrapPgError(err); wrappedErr != nil {
		return wrappedErr
	}

	return errors.ErrInternalStorageFetchDataFailed.Wrap(err, "source", connectionName)
}

// wrapErrorCommandTag - обёртывает ошибки выполнения команд и проверяет количество затронутых строк.
// Возвращает ErrEventStorageRecordsNotAffected - если
// команда Insert/Update/Delete не затронула ни одной строки.
func wrapErrorCommandTag(commandTag pgconn.CommandTag, err error) error {
	if err != nil {
		return wrapError(err)
	}

	if commandTag.RowsAffected() < 1 {
		if commandTag.Insert() || commandTag.Update() || commandTag.Delete() {
			return errors.ErrEventStorageRecordsNotAffected
		}
	}

	return nil
}

// wrapErrorExecRow - проверяет, что командой была затронута ровно одна запись.
// Возвращает ErrEventStorageNoRecordFound - если не затронуто ни одной записи,
// или ErrInternalStorageQueryFailed (с details "many records affected") - если затронуто более одной записи.
func wrapErrorExecRow(commandTag pgconn.CommandTag, err error) error {
	if err != nil {
		return wrapError(err)
	}

	if commandTag.RowsAffected() == 1 {
		return nil
	}

	if commandTag.RowsAffected() < 1 {
		return errors.ErrEventStorageNoRecordFound
	}

	return errors.ErrInternalStorageQueryFailed.WithDetails(
		"many records affected",
		"source", connectionName,
	)
}

// wrapErrorExecAffected - возвращает число затронутых строк;
// 0 строк не считается ошибкой (в отличие от wrapErrorCommandTag).
func wrapErrorExecAffected(commandTag pgconn.CommandTag, err error) (int, error) {
	if err != nil {
		return 0, wrapError(err)
	}

	return int(commandTag.RowsAffected()), nil
}
