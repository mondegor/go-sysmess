package wire

import (
	"github.com/mondegor/go-sysmess/errors/kind"
	"github.com/mondegor/go-sysmess/errors/runtime"
	"github.com/mondegor/go-sysmess/errors/runtime/hint"
	"github.com/mondegor/go-sysmess/errors/runtime/hint/instance"
	"github.com/mondegor/go-sysmess/errors/runtime/hint/stacktrace"
)

// InitErrors - инициализирует работу с runtime ошибками.
// Для ошибок типа Internal при их создании формируется стек вызова и ID ошибки.
// Для системной ошибки только ID ошибки.
func InitErrors(opts ErrorConfig) {
	var (
		internalOptions []runtime.Option
		systemOptions   []runtime.Option
	)

	onCreateOption := runtime.WithOnCreate(
		func(_ kind.Enum, _ error) (bag any) {
			return hint.New(
				hint.WithErrorID(instance.GenerateID()),
			)
		},
	)

	systemOptions = append(systemOptions, onCreateOption)

	if opts.HasCaller {
		caller := stacktrace.NewCaller(
			stacktrace.WithDepth(opts.CallerDepth),
			stacktrace.WithStackTraceFilter(
				stacktrace.TrimUpperFilter(opts.CallerUpperBounds),
			),
		)

		onCreateOption = runtime.WithOnCreate(
			func(_ kind.Enum, _ error) (bag any) {
				return hint.New(
					hint.WithErrorID(instance.GenerateID()),
					hint.WithStackTrace(caller.Call()),
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
