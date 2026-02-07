package main

import (
	"context"
	"errors"
	"fmt"
	stdlog "log/slog"
	"os"

	"github.com/mondegor/go-sysmess/errors/runtime/hint/instance"
	"github.com/mondegor/go-sysmess/mrlog"
	"github.com/mondegor/go-sysmess/mrlog/color"
	"github.com/mondegor/go-sysmess/mrlog/slog"
	"github.com/mondegor/go-sysmess/mrlog/slog/middleware"
)

// main - пример создания логгера и его использование с различными опциями.
func main() {
	logger, err := slog.NewLoggerAdapter(
		slog.WithWriter(os.Stdout),
		slog.WithLevel("INFO"),
		slog.WithJsonFormat(true),
		slog.WithTimeFormat("RFC3339Nano"),
	)
	if err != nil {
		mrlog.Fatal(err.Error())
	}

	ctx := context.WithValue(context.Background(), "env", "dev")

	printMsg(ctx, logger)

	logger, err = slog.NewLoggerAdapter(
		slog.WithWriter(os.Stdout),
		slog.WithLevel("DEBUG"),
		slog.WithJsonFormat(false),
		slog.WithTimeFormat("Kitchen"),
		slog.WithMiddlewareHandler(
			middleware.BeforeHandle(
				func(ctx context.Context, rec stdlog.Record) stdlog.Record {
					rec.Attrs(func(attr stdlog.Attr) bool {
						if attr.Value.Kind() == stdlog.KindAny {
							if e, ok := attr.Value.Any().(*baseError); ok {
								if id := e.ID(); id != "" {
									rec.Add("errorId", id)
								}

								rec.Add(e.Attrs()...)
							}
						}

						return true
					})

					// data from context
					rec.Add("env", ctx.Value("env"))

					return rec
				},
			),
		),
		slog.WithReplaceAttrs(func(attr stdlog.Attr) (newAttr stdlog.Attr) {
			if attr.Value.Kind() == stdlog.KindAny {
				if e, ok := attr.Value.Any().(*baseError); ok {
					attr.Value = stdlog.AnyValue((*lessVerboseError)(e))
				}
			}

			return attr
		}),
		slog.WithColorMode(true),
		slog.WithColorizeAttr("processId", color.Yellow, color.LightGray),
		slog.WithColorizeAttr("taskId", color.LightYellow, color.LightGray),
		slog.WithColorizeAttr("errorId", color.Yellow, color.Red),
		slog.WithColorizeAttr("env", color.LightCyan, color.LightGray),
		slog.WithColorizeAttr("version", color.LightCyan, color.LightGray),
		slog.WithColorizeAttr("sql", color.Cyan, color.Green),
	)
	if err != nil {
		mrlog.Fatal(err.Error())
	}

	loggerX2 := logger.WithAttrs("my-attr", 1)

	printMsg(ctx, loggerX2)

	mrlog.FatalError(logger, "Fatal error", "error", errors.New("my fatal error"))
}

func printMsg(ctx context.Context, logger mrlog.Logger) {
	logger.Info(ctx, "Logger info message - OK!")
	logger.Debug(ctx, "Logger DEBUG message", "version", "v1.0.0")
	logger.Error(ctx, "Error with error message", "error", errors.New("my error"))
	logger.Warn(ctx, "Warning with error message", "error", errors.New("my warning"))
	logger.Info(ctx, "Exec query", "sql", "SELECT COUNT(*) FROM table_name")

	err := error(&baseError{
		id:      instance.GenerateID(),
		message: "my error with attr-1",
		args:    []any{"err-attr-1", "err-value-1"},
	})
	logger.Error(ctx, "Error with error message and args", "error", err)

	logger = mrlog.WithAttrs(logger, "processId", "D8OR0E27-7WMZ-SC1A")
	logger.Info(ctx, "Start process", "service", "MainService")
	logger.Info(ctx, "Start task", "taskId", "D8OR2RFL-8751-N7V8")

	fmt.Println("-------------------------------------")
}

type (
	baseError struct {
		id      string
		message string
		args    []any
	}
)

func (e *baseError) Error() string {
	return fmt.Sprintf("[id: %s] %s: args%+v", e.id, e.message, e.args)
}

func (e *baseError) ID() string {
	return e.id
}

func (e *baseError) Attrs() []any {
	return e.args
}

type lessVerboseError baseError

func (e *lessVerboseError) Error() string {
	return e.message
}
