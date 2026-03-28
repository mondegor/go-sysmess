package errors

var (
	// ErrEventStorageNoRecordFound - no record found.
	// В хранилище данных нет указанной записи, в зависимости от логики это может быть не ошибкой.
	ErrEventStorageNoRecordFound = New("no record found")

	// ErrEventStorageRecordsNotAffected - records not affected.
	// В хранилище данных указанная запись не была обновлена, в зависимости от логики это может быть не ошибкой.
	ErrEventStorageRecordsNotAffected = New("records not affected")
)
