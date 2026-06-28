package errors

var (
	// ErrEventStorageNoRecordFound - запись не найдена.
	// В хранилище данных нет указанной записи, в зависимости от логики это может быть не ошибкой.
	ErrEventStorageNoRecordFound = New("no record found")

	// ErrEventStorageRecordsNotAffected - записи не затронуты.
	// В хранилище данных указанная запись не была обновлена, в зависимости от логики это может быть не ошибкой.
	ErrEventStorageRecordsNotAffected = New("records not affected")

	// ErrEventRecordAlreadyExists - такая запись уже существует.
	ErrEventRecordAlreadyExists = New("record already exists")
)
