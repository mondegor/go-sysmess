package mrerr

import (
	"context"
)

//go:generate mockgen -source=error_handler.go -destination=./mock/error_handler.go

type (
	// ErrorHandler - обработчик ошибок.
	ErrorHandler interface {
		Handle(ctx context.Context, err error)
		// TODO: попробовать analyzedKind перенести в err, чтобы extraHandler сделать не зависищим от типа
		HandleWith(ctx context.Context, err error, extraHandler func(analyzedKind ErrorKind, err error))
	}
)
