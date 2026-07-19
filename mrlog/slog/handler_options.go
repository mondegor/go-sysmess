package slog

import (
	stdlog "log/slog"
	"time"

	"github.com/mondegor/go-core/mrlog/level"
)

func handlerOptions(opts options) *stdlog.HandlerOptions {
	return &stdlog.HandlerOptions{
		Level: stdlog.Level(opts.level),
		ReplaceAttr: func(groups []string, attr stdlog.Attr) stdlog.Attr {
			attr = opts.replaceAttr(attr)

			// встроенные ключи (time, level) особые только на верхнем уровне записи,
			// вложенные атрибуты с такими же именами трогать нельзя
			if len(groups) > 0 {
				return attr
			}

			switch attr.Key {
			case stdlog.TimeKey:
				if tm, ok := attr.Value.Any().(time.Time); ok {
					attr.Value = stdlog.StringValue(formatTime(tm, opts.timeLocation, opts.timeFormat))
				}

			case stdlog.LevelKey:
				if lv, ok := attr.Value.Any().(stdlog.Level); ok {
					if lv < stdlog.LevelInfo || lv > stdlog.LevelError {
						attr.Value = stdlog.StringValue(level.Enum(lv).String())
					}
				}
			}

			return attr
		},
	}
}
