package wire

import (
	"github.com/mondegor/go-sysmess/errors/kind"
	"github.com/mondegor/go-sysmess/errors/runtime"
	"github.com/mondegor/go-sysmess/errors/runtime/hint"
	"github.com/mondegor/go-sysmess/errors/runtime/instance"
	"github.com/mondegor/go-sysmess/errors/runtime/stacktrace"
)

// InitErrors - инициализирует работу с runtime ошибками.
// Для ошибок типа Internal при их создании формируется стек вызова и ID ошибки.
// Для системной ошибки только ID ошибки.
func InitErrors(opts ErrorConfig) {
	var (
		internalOptions []runtime.Option
		systemOptions   []runtime.Option
	)

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

	systemOptions = append(systemOptions, onCreateOption)

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

	internalOptions = append(internalOptions, onCreateOption)

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
