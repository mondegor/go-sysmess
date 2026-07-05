package db_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-core/errors"
	"github.com/mondegor/go-core/mrpostgres/db"
)

func TestFieldUpdater_Update(t *testing.T) {
	t.Parallel()

	errExec := errors.New("exec failed")

	type testCase struct {
		name      string
		execErr   error
		wantErrIs error // nil - ожидается отсутствие ошибки
	}

	// Update - сквозная обёртка над ExecRow: проверяем проброс аргументов
	// и что ошибка ExecRow возвращается без изменений (маппинг счётчика - ответственность ExecRow).
	tests := []testCase{
		{
			name:      "exactly one updated",
			execErr:   nil,
			wantErrIs: nil,
		},
		{
			name:      "no record found is propagated",
			execErr:   errors.ErrEventStorageNoRecordFound,
			wantErrIs: errors.ErrEventStorageNoRecordFound,
		},
		{
			name:      "exec error is propagated",
			execErr:   errExec,
			wantErrIs: errExec,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			conn := &fakeConn{err: tt.execErr}
			updater := db.NewFieldUpdater[int, string](
				&fakeConnManager{conn: conn},
				"users", "user_id", "name", "deleted_at",
			)

			gotErr := updater.Update(context.Background(), 42, "bob")

			// id и value должны быть переданы в запрос.
			require.Equal(t, []any{42, "bob"}, conn.lastArgs)

			if tt.wantErrIs == nil {
				assert.NoError(t, gotErr)

				return
			}

			assert.ErrorIs(t, gotErr, tt.wantErrIs)
		})
	}
}
