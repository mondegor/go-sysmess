package slog

import (
	"io"
	"log/slog"
	"time"

	"github.com/lmittmann/tint"

	"github.com/mondegor/go-sysmess/mrlog/color"
)

func newColoredHandler(w io.Writer, opts options) slog.Handler {
	return tint.NewHandler(
		w,
		&tint.Options{
			Level: slog.Level(opts.level),
			ReplaceAttr: func(_ []string, attr slog.Attr) slog.Attr {
				attr = opts.replaceAttr(attr)

				if attr.Key == slog.TimeKey {
					if tm, ok := attr.Value.Any().(time.Time); ok {
						attr.Value = slog.StringValue(color.ColorizeText(color.Gray, tm.Format(opts.timeFormat)))
					}

					return attr
				}

				if attr.Key == slog.LevelKey {
					if lv, ok := attr.Value.Any().(slog.Level); ok {
						switch {
						case lv > slog.LevelError+4:
							attr.Value = slog.StringValue(color.ColorizeText(color.Blue, "TRC"))
						case lv > slog.LevelError:
							attr.Value = slog.StringValue(color.ColorizeText(color.Red, "FAT"))
						case lv < slog.LevelInfo:
							attr.Value = slog.StringValue(color.ColorizeText(color.Yellow, "DBG"))
						}
					}

					return attr
				}

				if attr.Key == slog.MessageKey {
					return attr
				}

				if attr.Value.Kind() == slog.KindAny {
					if err, ok := attr.Value.Any().(error); ok {
						aErr := tint.Err(err)
						aErr.Key = attr.Key

						return aErr
					}
				}

				if clr, ok := opts.attrKey2color[attr.Key]; ok {
					return colorizeAttr(attr, clr)
				}

				return colorizeAttr(attr, opts.attrColorByDefault)
			},
		},
	)
}

func colorizeAttr(attr slog.Attr, clr attrColor) slog.Attr {
	if clr.keyColor != "" {
		attr.Key = color.ColorizeText(clr.keyColor, attr.Key)
	}

	if clr.valueColor == "" {
		return attr
	}

	return slog.String(attr.Key, color.ColorizeText(clr.valueColor, attr.Value.String()))
}
