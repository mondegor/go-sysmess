package slog

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/mondegor/go-sysmess/mrlog"
	"github.com/mondegor/go-sysmess/mrlog/color"
	"github.com/mondegor/go-sysmess/mrlog/parser"
)

type (
	// LoggerAdapter - логгер на крайний случай, например,
	// когда не был установлен логгер в контексте.
	LoggerAdapter struct {
		sl *slog.Logger
	}

	options struct {
		stdout             io.Writer
		levelString        string
		level              mrlog.Level
		jsonFormat         bool
		timeFormat         string
		middlewareHandlers []func(next slog.Handler) slog.Handler
		replaceAttr        func(attr slog.Attr) (newAttr slog.Attr)
		colorMode          bool
		attrKey2color      map[string]attrColor
		attrColorByDefault attrColor
	}

	attrColor struct {
		keyColor   string
		valueColor string
	}
)

// NewLoggerAdapter - создаёт объект LoggerAdapter.
func NewLoggerAdapter(opts ...Option) (logger *LoggerAdapter, err error) {
	o := options{
		stdout:      os.Stderr,
		levelString: mrlog.LevelDebug.String(),
		jsonFormat:  false,
		timeFormat:  "RFC3339",
		replaceAttr: func(attr slog.Attr) (newAttr slog.Attr) {
			return attr
		},
		colorMode:     true,
		attrKey2color: make(map[string]attrColor),
		attrColorByDefault: attrColor{
			keyColor:   color.Cyan,
			valueColor: color.LightGray,
		},
	}

	for _, opt := range opts {
		opt(&o)
	}

	o.level, err = parser.ParseLevel(o.levelString)
	if err != nil {
		return nil, fmt.Errorf("error parsing level: %w", err)
	}

	o.timeFormat, err = parser.ParseDateTimeFormat(o.timeFormat)
	if err != nil {
		return nil, fmt.Errorf("error parsing time: %w", err)
	}

	var handler slog.Handler

	switch {
	case o.jsonFormat:
		handler = slog.NewJSONHandler(o.stdout, handlerOptions(o))
	case o.colorMode:
		handler = newColoredHandler(o.stdout, o)
	default:
		handler = slog.NewTextHandler(o.stdout, handlerOptions(o))
	}

	for i := len(o.middlewareHandlers) - 1; i >= 0; i-- {
		handler = o.middlewareHandlers[i](handler)
	}

	return New(handler), nil
}

// New - создаёт объект LoggerAdapter.
func New(handler slog.Handler) *LoggerAdapter {
	return &LoggerAdapter{
		sl: slog.New(handler),
	}
}

// WithAttrs - возвращает новый логгер с прикреплёнными атрибутами.
// Один атрибут состоит из двух аргументов: ключ:string/значение:any.
func (l *LoggerAdapter) WithAttrs(args ...any) mrlog.Logger {
	if len(args) == 0 {
		return l
	}

	c := *l
	c.sl = l.sl.With(args...)

	return &c
}

// Enabled - информирует включён ли указанный уровень логирования.
func (l *LoggerAdapter) Enabled(level mrlog.Level) bool {
	return l.sl.Enabled(context.Background(), slog.Level(level))
}

// Log - логирует сообщения на указанном уровне.
func (l *LoggerAdapter) Log(ctx context.Context, level mrlog.Level, msg string, args ...any) {
	l.sl.Log(ctx, slog.Level(level), msg, args...)
}

// Debug - логирует сообщения на уровне mrlog.LevelDebug.
func (l *LoggerAdapter) Debug(ctx context.Context, msg string, args ...any) {
	l.sl.DebugContext(ctx, msg, args...)
}

// DebugFunc - логирует сообщения на уровне mrlog.LevelDebug с их отложенным созданием.
// Применяется для того, чтобы исключить формирование в продуктовой среде больших отладочных
// сообщений с использованием многочисленных параметров.
func (l *LoggerAdapter) DebugFunc(ctx context.Context, createMsg func() string, args ...any) {
	if !l.Enabled(mrlog.LevelDebug) {
		return
	}

	l.sl.DebugContext(ctx, createMsg(), args...)
}

// Info - логирует сообщения на уровне mrlog.LevelInfo.
func (l *LoggerAdapter) Info(ctx context.Context, msg string, args ...any) {
	l.sl.InfoContext(ctx, msg, args...)
}

// Warn - логирует сообщения на уровне mrlog.LevelWarn.
func (l *LoggerAdapter) Warn(ctx context.Context, msg string, args ...any) {
	l.sl.WarnContext(ctx, msg, args...)
}

// Error - логирует сообщения на уровне mrlog.LevelError.
func (l *LoggerAdapter) Error(ctx context.Context, msg string, args ...any) {
	l.sl.ErrorContext(ctx, msg, args...)
}
