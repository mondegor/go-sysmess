package slog

import (
	"io"
	stdlog "log/slog"
	"time"

	"github.com/lmittmann/tint"

	"github.com/mondegor/go-sysmess/mrlog/color"
)

func newColoredHandler(w io.Writer, opts options) stdlog.Handler {
	return tint.NewHandler(
		w,
		&tint.Options{
			Level: stdlog.Level(opts.level),
			ReplaceAttr: func(_ []string, attr stdlog.Attr) stdlog.Attr {
				attr = opts.replaceAttr(attr)

				if attr.Key == stdlog.TimeKey {
					if tm, ok := attr.Value.Any().(time.Time); ok {
						attr.Value = stdlog.StringValue(color.ColorizeText(color.Gray, tm.Format(opts.timeFormat)))
					}

					return attr
				}

				if attr.Key == stdlog.LevelKey {
					if lv, ok := attr.Value.Any().(stdlog.Level); ok {
						switch {
						case lv > stdlog.LevelError+4:
							attr.Value = stdlog.StringValue(color.ColorizeText(color.Blue, "TRC"))
						case lv > stdlog.LevelError:
							attr.Value = stdlog.StringValue(color.ColorizeText(color.Red, "FAT"))
						case lv < stdlog.LevelInfo:
							attr.Value = stdlog.StringValue(color.ColorizeText(color.Yellow, "DBG"))
						}
					}

					return attr
				}

				if attr.Key == stdlog.MessageKey {
					return attr
				}

				if attr.Value.Kind() == stdlog.KindAny {
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

func colorizeAttr(attr stdlog.Attr, clr attrColor) stdlog.Attr {
	if clr.keyColor != "" {
		attr.Key = color.ColorizeText(clr.keyColor, attr.Key)
	}

	if clr.valueColor == "" {
		return attr
	}

	return stdlog.String(attr.Key, color.ColorizeText(clr.valueColor, attr.Value.String()))
}
