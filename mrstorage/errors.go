package mrstorage

import (
	"github.com/mondegor/go-core/errors"
)

// ErrSystemFileProviderPingError - ошибка проверки работоспособности файлового провайдера.
var ErrSystemFileProviderPingError = errors.NewSystemProto("file provider ping error")
