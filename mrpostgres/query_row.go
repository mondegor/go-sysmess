package mrpostgres

import (
	"github.com/jackc/pgx/v5"

	"github.com/mondegor/go-core/errors"
)

type (
	queryRow struct {
		row pgx.Row
	}
)

// Scan - извлекает значения из одной записи результата запроса.
// Возвращает ErrEventStorageNoRecordFound - если запись не найдена,
// и ErrInternalStorageFetchDataFailed - если записей более одной.
func (qr *queryRow) Scan(dest ...any) error {
	if err := qr.row.Scan(dest...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errors.ErrEventStorageNoRecordFound
		}

		if errors.Is(err, pgx.ErrTooManyRows) {
			return errors.ErrInternalStorageFetchDataFailed.WithDetails(
				"too many rows",
				"source", connectionName,
			)
		}

		return wrapErrorFetchDataFailed(err)
	}

	return nil
}
