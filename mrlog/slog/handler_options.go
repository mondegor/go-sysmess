package slog

import (
	stdlog "log/slog"
	"time"

	"github.com/mondegor/go-sysmess/mrlog/level"
)

func handlerOptions(opts options) *stdlog.HandlerOptions {
	return &stdlog.HandlerOptions{
		Level: stdlog.Level(opts.level),
		ReplaceAttr: func(_ []string, attr stdlog.Attr) stdlog.Attr {
			attr = opts.replaceAttr(attr)

			switch attr.Key {
			case stdlog.TimeKey:
				if tm, ok := attr.Value.Any().(time.Time); ok {
					attr.Value = stdlog.StringValue(tm.Format(opts.timeFormat))
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
