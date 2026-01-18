package errors

var (
	// ErrEventStorageNoRowFound - no row found.
	// В хранилище данных нет указанной записи, в зависимости от логики это может быть не ошибкой.
	ErrEventStorageNoRowFound = New("no row found")

	// ErrEventStorageRowsNotAffected - rows not affected.
	// В хранилище данных указанная запись не была обновлена, в зависимости от логики это может быть не ошибкой.
	ErrEventStorageRowsNotAffected = New("rows not affected")
)
