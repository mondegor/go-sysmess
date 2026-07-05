package errors

import (
	"github.com/mondegor/go-sysmess/errors/kind"
	"github.com/mondegor/go-sysmess/errors/runtime"
	"github.com/mondegor/go-sysmess/errors/runtime/hint"
	"github.com/mondegor/go-sysmess/errors/runtime/instance"
	"github.com/mondegor/go-sysmess/errors/runtime/stacktrace"
)

// InitErrors - инициализирует опции для разного типа runtime-ошибок.
// Для Internal-ошибок формируется стек вызовов (если HasCaller=true) и ID ошибки.
// Для System-ошибок формируется только ID ошибки.
func InitErrors(opts ErrorConfig) {
	// применяется для System и Internal ошибок по умолчанию
	onCreateOption := runtime.WithOnCreate(
		func(kind kind.Enum, wrappedErr error) (bag any) {
			if ht := hint.Extract(wrappedErr); ht.ErrorKind() != 0 {
				return ht
			}

			return hint.New(
				kind,
				instance.GenerateID(),
				nil,
			)
		},
	)

	systemOptions := []runtime.Option{onCreateOption}

	if opts.HasCaller {
		caller := stacktrace.NewCaller(
			stacktrace.WithDepth(int(opts.CallerDepth)),
			stacktrace.WithShowFunc(opts.CallerShowFunc),
			stacktrace.WithFindBottomBoundFunc(
				stacktrace.FindBottomBound(opts.CallerUpperBounds),
			),
		)

		// применяется только для Internal ошибок
		onCreateOption = runtime.WithOnCreate(
			func(kind kind.Enum, wrappedErr error) (bag any) {
				if ht := hint.Extract(wrappedErr); ht.ErrorKind() != 0 {
					return ht
				}

				return hint.New(
					kind,
					instance.GenerateID(),
					caller.Call(),
				)
			},
		)
	}

	internalOptions := []runtime.Option{onCreateOption}

	runtime.InitDelayedOptions(
		runtime.OptionsHandlerFunc(func(kindErr kind.Enum, _ string) []runtime.Option {
			switch kindErr {
			case kind.Internal:
				return internalOptions
			case kind.System:
				return systemOptions
			default:
				return nil
			}
		}),
	)
}
