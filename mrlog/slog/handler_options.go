package slog

import (
	"log/slog"
	"time"

	"github.com/mondegor/go-sysmess/mrlog"
)

func handlerOptions(opts options) *slog.HandlerOptions {
	return &slog.HandlerOptions{
		Level: slog.Level(opts.level),
		ReplaceAttr: func(_ []string, attr slog.Attr) slog.Attr {
			attr = opts.replaceAttr(attr)

			switch {
			case attr.Key == slog.TimeKey:
				if tm, ok := attr.Value.Any().(time.Time); ok {
					attr.Value = slog.StringValue(tm.Format(opts.timeFormat))
				}

			case attr.Key == slog.LevelKey:
				if lv, ok := attr.Value.Any().(slog.Level); ok {
					if lv < slog.LevelInfo || lv > slog.LevelError {
						attr.Value = slog.StringValue(mrlog.Level(lv).String())
					}
				}
			}

			return attr
		},
	}
}
