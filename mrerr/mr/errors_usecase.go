package mr

import "github.com/mondegor/go-sysmess/mrerr"

var (
	// ErrUseCaseOperationFailed - operation failed (UseCaseErrorWrapper по умолчанию оборачивает в эту ошибку все нераспознанные ошибки).
	ErrUseCaseOperationFailed = mrerr.NewKindInternal("operation failed")

	// ErrUseCaseTemporarilyUnavailable - system is temporarily unavailable (UseCaseErrorWrapper оборачивает в эту ошибку все системные ошибки).
	// Системная ошибка, которая сообщает о сетевых проблемах, о работоспособности внешних ресурсов (БД, API, FileSystem).
	ErrUseCaseTemporarilyUnavailable = mrerr.NewKindSystem("system is temporarily unavailable", mrerr.WithCode(ErrorCodeTemporarilyUnavailable))

	// ErrUseCaseIncorrectInputData - input data is incorrect (желательно настраивать в слое выше валидатор этих данных).
	ErrUseCaseIncorrectInputData = mrerr.NewKindUser("IncorrectInputData", "input data is incorrect: '{Reason}'")

	// ErrUseCaseIncorrectInternalInputData - input data is incorrect
	// (когда нет возможности ответить клиенту, например, данные поступают из очереди,
	// или когда поступаемые данные не зависят от клиента, например, приходят из конфига,
	// в остальных случаях лучше использовать ErrUseCaseIncorrectInputData).
	ErrUseCaseIncorrectInternalInputData = mrerr.NewKindInternal("input data is incorrect")

	// ErrUseCaseAccessForbidden - access forbidden (данная ошибка заворачивается в 401 или 403).
	ErrUseCaseAccessForbidden = mrerr.NewKindUser("AccessForbidden", "access forbidden")

	// ErrUseCaseEntityNotFound - entity not found (данная ошибка заворачивается в 404).
	ErrUseCaseEntityNotFound = mrerr.NewKindUser("EntityNotFound", "entity not found")

	// ErrUseCaseEntityNotAvailable - entity is not available.
	ErrUseCaseEntityNotAvailable = mrerr.NewKindUser("EntityNotAvailable", "entity is not available")

	// ErrUseCaseEntityAlreadyExists - entity already exists.
	ErrUseCaseEntityAlreadyExists = mrerr.NewKindUser("EntityAlreadyExists", "entity already exists")

	// ErrUseCaseEntityVersionInvalid - entity version is invalid (данная ошибка заворачивается в 409).
	ErrUseCaseEntityVersionInvalid = mrerr.NewKindUser("EntityVersionInvalid", "entity version is invalid")

	// ErrUseCaseSwitchStatusRejected - switching from status to status is rejected.
	ErrUseCaseSwitchStatusRejected = mrerr.NewKindUser("SwitchStatusRejected", "switching from '{StatusFrom}' to '{StatusTo}' is rejected")

	// ErrUseCaseInvalidFile - file is invalid.
	ErrUseCaseInvalidFile = mrerr.NewKindUser("InvalidFile", "file is invalid")
)
