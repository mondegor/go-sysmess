package slog

import (
	"io"
	stdlog "log/slog"
	"time"

	"github.com/lmittmann/tint"

	"github.com/mondegor/go-core/mrlog/color"
)

func newColoredHandler(w io.Writer, opts options) stdlog.Handler {
	return tint.NewHandler(
		w,
		&tint.Options{
			Level: stdlog.Level(opts.level),
			ReplaceAttr: func(groups []string, attr stdlog.Attr) stdlog.Attr {
				attr = opts.replaceAttr(attr)

				// пустой атрибут - это запрос на его удаление из записи; раскрашивать
				// его нельзя, так как ANSI-коды попадут в ключ, атрибут перестанет
				// быть пустым и обработчик не сможет его отбросить
				if attr.Equal(stdlog.Attr{}) {
					return attr
				}

				// встроенные ключи (time, level) особые только на верхнем уровне записи,
				// вложенные атрибуты с такими же именами оформляются как обычные.
				// На msg это не распространяется: он рядовой ключ и красится везде
				// (см. colorizeBuiltinAttr)
				if len(groups) == 0 {
					if builtin, ok := colorizeBuiltinAttr(attr, opts.timeLocation, opts.timeFormat); ok {
						return builtin
					}
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

// colorizeBuiltinAttr - оформляет системный атрибут записи (time, level).
// Сообщает false, если атрибут системным не является.
//
// Сообщение записи (msg) системным здесь не считается: отличить его от
// пользовательского атрибута с тем же именем нельзя (оба - строки), поэтому
// msg остаётся рядовым ключом и красится всегда, включая вложенный;
// цвет задаётся как у любого другого атрибута (см. WithColorizeAttr).
//
// Системное поле записи всегда приходит со своим типом, поэтому несовпадение
// типа означает пользовательский атрибут, случайно названный именем поля.
// Обратный случай неустраним: атрибут верхнего уровня с именем time, несущий
// настоящий time.Time, от метки записи неотличим и будет пересчитан в часовой
// пояс логгера - контракт ReplaceAttr не даёт признака, чтобы их развести.
func colorizeBuiltinAttr(attr stdlog.Attr, loc *time.Location, format string) (stdlog.Attr, bool) {
	switch attr.Key {
	case stdlog.TimeKey:
		tm, ok := attr.Value.Any().(time.Time)
		if !ok {
			return attr, false
		}

		attr.Value = stdlog.StringValue(
			color.ColorizeText(color.Gray, formatTime(tm, loc, format)),
		)

	case stdlog.LevelKey:
		lv, ok := attr.Value.Any().(stdlog.Level)
		if !ok {
			return attr, false
		}

		switch {
		case lv > stdlog.LevelError+4:
			attr.Value = stdlog.StringValue(color.ColorizeText(color.Blue, "TRC"))
		case lv > stdlog.LevelError:
			attr.Value = stdlog.StringValue(color.ColorizeText(color.Red, "FAT"))
		case lv < stdlog.LevelInfo:
			attr.Value = stdlog.StringValue(color.ColorizeText(color.Yellow, "DBG"))
		}

	default:
		return attr, false
	}

	return attr, true
}

func colorizeAttr(attr stdlog.Attr, attrColor attrColor) stdlog.Attr {
	if attrColor.key != "" {
		attr.Key = color.ColorizeText(attrColor.key, attr.Key)
	}

	if attrColor.value == "" {
		return attr
	}

	return stdlog.String(attr.Key, color.ColorizeText(attrColor.value, attr.Value.String()))
}
