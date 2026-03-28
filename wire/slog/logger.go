package slog

import (
	"context"
	"fmt"
	stdlog "log/slog"
	"os"
	"strings"

	"github.com/mondegor/go-sysmess/errors/helper"
	"github.com/mondegor/go-sysmess/errors/kind"
	"github.com/mondegor/go-sysmess/errors/runtime/stacktrace"
	"github.com/mondegor/go-sysmess/mrapp"
	"github.com/mondegor/go-sysmess/mrlog/color"
	"github.com/mondegor/go-sysmess/mrlog/slog"
	"github.com/mondegor/go-sysmess/mrlog/slog/middleware"
	"github.com/mondegor/go-sysmess/mrtrace"
	"github.com/mondegor/go-sysmess/mrtrace/process"
	"github.com/mondegor/go-sysmess/wire"
)

type (
	runtimeError interface {
		error

		Kind() kind.Enum
		Hint() any
	}

	errorHint interface {
		ErrorID() string
		StackTraceIterator() func() (index int, name, file string, line int)
	}
)

// InitLogger - создаёт и инициализирует slog.LoggerAdapter.
func InitLogger(cfg wire.LoggerConfig) (logger *slog.LoggerAdapter, err error) {
	opts := initLoggerOptions(cfg)

	logger, err = slog.NewLoggerAdapter(opts...)
	if err != nil {
		return nil, fmt.Errorf("InitLogger: %w", err)
	}

	if !strings.HasPrefix(cfg.Environment, "local") {
		logger = logger.WithAttributes(mrapp.KeyAppEnvironment, cfg.Environment)

		if cfg.Version != "" {
			logger = logger.WithAttributes(mrapp.KeyAppVersion, cfg.Version)
		}
	}

	return logger, nil
}

func initLoggerOptions(cfg wire.LoggerConfig) []slog.Option {
	if cfg.Level == "" {
		cfg.Level = "info"
	}

	if cfg.TimeFormat == "" {
		cfg.TimeFormat = "RFC3339"
	}

	contextProcessIDs := process.ToKeyGetID(cfg.ContextProcessIDs)

	opts := []slog.Option{
		slog.WithWriter(os.Stdout),
		slog.WithLevel(strings.ToUpper(cfg.Level)),
		slog.WithJsonFormat(cfg.JsonFormat),
		slog.WithTimeFormat(cfg.TimeFormat),
		slog.WithMiddlewareHandler(
			middleware.BeforeHandle(
				func(ctx context.Context, rec stdlog.Record) stdlog.Record {
					var stack string

					rec.Attrs(
						func(attr stdlog.Attr) bool {
							if attr.Value.Kind() != stdlog.KindAny {
								return true
							}

							e, ok := attr.Value.Any().(runtimeError)
							if !ok {
								return true
							}

							if e.Kind() != kind.Internal && e.Kind() != kind.System {
								return true
							}

							rec.Add(
								helper.ExtractAttrs(
									e,
									func(key string) bool {
										return strings.HasPrefix(key, "log.") // TODO: filter debug
									},
								)...,
							)

							bag, ok := e.Hint().(errorHint)
							if !ok {
								return true
							}

							if id := bag.ErrorID(); id != "" {
								rec.Add(mrapp.KeyErrorID, id)
							}

							stack = strings.Join(stacktrace.ToStrings(bag.StackTraceIterator()), " | ")

							return true
						},
					)

					rec.Add(process.ExtractKeysValues(ctx, contextProcessIDs)...)

					// стек вызовов вставляется в самый конец
					if stack != "" {
						rec.Add(mrapp.KeyStackTrace, stack)
					}

					return rec
				},
			),
		),
		slog.WithReplaceAttrs(func(attr stdlog.Attr) (newAttr stdlog.Attr) {
			if key, ok := strings.CutPrefix(attr.Key, "log."); ok { // TODO: cut debug
				attr.Key = key
			}

			return attr
		}),
		slog.WithColorMode(cfg.ColorMode),
	}

	if cfg.ColorMode {
		opts = append(
			opts,
			slog.WithColorizeAttr(mrapp.KeyAppEnvironment, color.Yellow, color.LightGray),
			slog.WithColorizeAttr(mrapp.KeyAppVersion, color.Yellow, color.LightGray),
			slog.WithColorizeAttr(mrapp.KeyErrorID, color.Yellow, color.Red),

			slog.WithColorizeAttr(mrtrace.KeyProcessID, color.Yellow, color.LightGray),
			slog.WithColorizeAttr(mrtrace.KeyWorkerID, color.Yellow, color.LightGray),

			slog.WithColorizeAttr(mrtrace.KeyCorrelationID, color.LightYellow, color.LightGray),
			slog.WithColorizeAttr(mrtrace.KeyTaskID, color.LightYellow, color.LightGray),
			slog.WithColorizeAttr(mrtrace.KeyRequestID, color.LightYellow, color.LightGray),

			slog.WithColorizeAttr("sql", color.Cyan, color.Green),
		)
	}

	return opts
}
