package slog

import (
	"context"
	"fmt"
	stdlog "log/slog"
	"os"

	"github.com/mondegor/go-sysmess/mrlog/color"
	"github.com/mondegor/go-sysmess/mrlog/level"
	log "github.com/mondegor/go-sysmess/mrlog/logger"
)

type (
	// LoggerAdapter - логгер на крайний случай, например,
	// когда не был установлен логгер в контексте.
	LoggerAdapter struct {
		sl *stdlog.Logger
	}

	attrColor struct {
		keyColor   string
		valueColor string
	}
)

// NewLoggerAdapter - создаёт объект LoggerAdapter.
func NewLoggerAdapter(opts ...Option) (logger *LoggerAdapter, err error) {
	o := options{
		levelString:   level.Debug.String(),
		jsonFormat:    false,
		timeFormat:    "RFC3339",
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

	if o.stdout == nil {
		o.stdout = os.Stdout
	}

	o.level, err = level.Parse(o.levelString)
	if err != nil {
		return nil, fmt.Errorf("error parsing level: %w", err)
	}

	o.timeFormat, err = log.ParseDateTimeFormat(o.timeFormat)
	if err != nil {
		return nil, fmt.Errorf("error parsing time: %w", err)
	}

	if o.replaceAttr == nil {
		o.replaceAttr = func(attr stdlog.Attr) (newAttr stdlog.Attr) {
			return attr
		}
	}

	var handler stdlog.Handler

	switch {
	case o.jsonFormat:
		handler = stdlog.NewJSONHandler(o.stdout, handlerOptions(o))
	case o.colorMode:
		handler = newColoredHandler(o.stdout, o)
	default:
		handler = stdlog.NewTextHandler(o.stdout, handlerOptions(o))
	}

	for i := len(o.middlewareHandlers) - 1; i >= 0; i-- {
		handler = o.middlewareHandlers[i](handler)
	}

	return New(handler), nil
}

// New - создаёт объект LoggerAdapter.
func New(handler stdlog.Handler) *LoggerAdapter {
	return &LoggerAdapter{
		sl: stdlog.New(handler),
	}
}

// WithAttributes - возвращает новый логгер с прикреплёнными атрибутами.
// Один атрибут состоит из двух аргументов: ключ:string/значение:any.
func (l *LoggerAdapter) WithAttributes(attrs ...any) *LoggerAdapter {
	if len(attrs) == 0 {
		return l
	}

	c := *l
	c.sl = l.sl.With(attrs...)

	return &c
}

// WithAttrs - cм. WithAttributes.
func (l *LoggerAdapter) WithAttrs(attrs ...any) log.Logger {
	return l.WithAttributes(attrs...)
}

// Enabled - информирует включён ли указанный уровень логирования.
func (l *LoggerAdapter) Enabled(lvl level.Enum) bool {
	return l.sl.Enabled(context.Background(), stdlog.Level(lvl))
}

// Log - логирует сообщения на указанном уровне.
func (l *LoggerAdapter) Log(ctx context.Context, lvl level.Enum, msg string, args ...any) {
	l.sl.Log(ctx, stdlog.Level(lvl), msg, args...)
}

// Debug - логирует сообщения на уровне level.Debug.
func (l *LoggerAdapter) Debug(ctx context.Context, msg string, args ...any) {
	l.sl.DebugContext(ctx, msg, args...)
}

// DebugFunc - логирует сообщения на уровне level.Debug с их отложенным созданием.
// Применяется для того, чтобы исключить формирование в продуктовой среде больших отладочных
// сообщений с использованием многочисленных параметров.
func (l *LoggerAdapter) DebugFunc(ctx context.Context, createMsg func() string, args ...any) {
	if !l.Enabled(level.Debug) {
		return
	}

	l.sl.DebugContext(ctx, createMsg(), args...)
}

// Info - логирует сообщения на уровне level.Info.
func (l *LoggerAdapter) Info(ctx context.Context, msg string, args ...any) {
	l.sl.InfoContext(ctx, msg, args...)
}

// Warn - логирует сообщения на уровне level.Warn.
func (l *LoggerAdapter) Warn(ctx context.Context, msg string, args ...any) {
	l.sl.WarnContext(ctx, msg, args...)
}

// Error - логирует сообщения на уровне level.Error.
func (l *LoggerAdapter) Error(ctx context.Context, msg string, args ...any) {
	l.sl.ErrorContext(ctx, msg, args...)
}

// Trace - логирует сообщения на уровне level.Trace.
func (l *LoggerAdapter) Trace(ctx context.Context, args ...any) {
	l.sl.Log(ctx, stdlog.Level(level.Trace), "trace", args...)
}
