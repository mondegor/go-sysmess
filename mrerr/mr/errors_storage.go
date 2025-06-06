package mr

import "github.com/mondegor/go-sysmess/mrerr"

var (
	// ErrStorageConnectionIsAlreadyCreated - connection is already created.
	ErrStorageConnectionIsAlreadyCreated = mrerr.NewKindInternal("connection is already created: '{Name}'")

	// ErrStorageConnectionIsNotOpened - connection is not opened.
	ErrStorageConnectionIsNotOpened = mrerr.NewKindInternal("connection is not opened: '{Name}'")

	// ErrStorageConnectionIsBusy - connection is busy.
	ErrStorageConnectionIsBusy = mrerr.NewKindSystem("connection is busy: '{Name}'")

	// ErrStorageConnectionFailed - connection is failed.
	ErrStorageConnectionFailed = mrerr.NewKindSystem("connection is failed: '{Name}'")

	// ErrStorageQueryFailed - query is failed (ErrStorageQueryFailed оборачивает в эту ошибку все нераспознанные ошибки).
	ErrStorageQueryFailed = mrerr.NewKindInternal("query is failed")

	// ErrStorageFetchDataFailed - fetching data is failed.
	ErrStorageFetchDataFailed = mrerr.NewKindInternal("fetching data is failed")

	// ErrStorageNoRowFound - no row found (в хранилище данных нет указанной записи, в зависимости от логики это может быть не ошибкой).
	// Это вспомогательная ошибка, для неё отключено формирование стека вызовов и отправление события о её создании.
	ErrStorageNoRowFound = mrerr.NewKindInternal("no row found", mrerr.WithDisabledCaller(), mrerr.WithDisabledOnCreated())

	// ErrStorageRowsNotAffected - rows not affected (в хранилище данных указанная запись не была обновлена, в зависимости от логики это может быть не ошибкой).
	// Это вспомогательная ошибка, для неё отключено формирование стека вызовов и отправление события о её создании.
	ErrStorageRowsNotAffected = mrerr.NewKindInternal("rows not affected", mrerr.WithDisabledCaller(), mrerr.WithDisabledOnCreated())

	// ErrStorageLockNotObtained - .
	// Это вспомогательная ошибка, для неё отключено формирование стека вызовов и отправление события о её создании.
	ErrStorageLockNotObtained = mrerr.NewKindInternal("lock not obtained", mrerr.WithDisabledCaller(), mrerr.WithDisabledOnCreated())

	// ErrStorageLockNotHeld - .
	// Это вспомогательная ошибка, для неё отключено формирование стека вызовов и отправление события о её создании.
	ErrStorageLockNotHeld = mrerr.NewKindInternal("lock not held", mrerr.WithDisabledCaller(), mrerr.WithDisabledOnCreated())
)
