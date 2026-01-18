package slog

import (
	"io"
	stdlog "log/slog"

	"github.com/mondegor/go-sysmess/mrlog/level"
)

const (
	// AttrColorByDefaultKey - ключ используемый в WithColorizeAttr
	// для установки цветов ключей и значений по умолчанию.
	AttrColorByDefaultKey = "attrColorByDefault"
)

type (
	// Option - настройка объекта LoggerAdapter.
	Option func(o *options)

	options struct {
		stdout             io.Writer
		levelString        string
		level              level.Enum
		jsonFormat         bool
		timeFormat         string
		middlewareHandlers []func(next stdlog.Handler) stdlog.Handler
		replaceAttr        func(attr stdlog.Attr) (newAttr stdlog.Attr)
		colorMode          bool
		attrKey2color      map[string]attrColor
		attrColorByDefault attrColor
	}
)

// WithWriter - устанавливает опцию stdout для LoggerAdapter.
func WithWriter(value io.Writer) Option {
	return func(o *options) {
		o.stdout = value
	}
}

// WithLevel - устанавливает уровень логирования для LoggerAdapter.
func WithLevel(value string) Option {
	return func(o *options) {
		o.levelString = value
	}
}

// WithJsonFormat - устанавливает признак логирования в json формате для LoggerAdapter.
func WithJsonFormat(value bool) Option {
	return func(o *options) {
		o.jsonFormat = value
	}
}

// WithTimeFormat - устанавливает формат времени при логировании для LoggerAdapter.
func WithTimeFormat(value string) Option {
	return func(o *options) {
		o.timeFormat = value
	}
}

// WithMiddlewareHandler - устанавливает middleware для LoggerAdapter.
func WithMiddlewareHandler(value ...func(next stdlog.Handler) stdlog.Handler) Option {
	return func(o *options) {
		o.middlewareHandlers = append(o.middlewareHandlers, value...)
	}
}

// WithReplaceAttrs - устанавливает функцию замены атрибутов для LoggerAdapter.
func WithReplaceAttrs(value func(attr stdlog.Attr) (newAttr stdlog.Attr)) Option {
	return func(o *options) {
		o.replaceAttr = value
	}
}

// WithColorMode - устанавливает режим логирования в цветном формате
// (только при выключенном JsonFormat) для LoggerAdapter.
func WithColorMode(value bool) Option {
	return func(o *options) {
		o.colorMode = value
	}
}

// WithColorizeAttr - добавляет определение цветов для указанного ключа и его значения
// (только при включенном ColorMode) для LoggerAdapter.
func WithColorizeAttr(key, keyColor, valueColor string) Option {
	return func(o *options) {
		color := attrColor{
			keyColor:   keyColor,
			valueColor: valueColor,
		}

		if key == AttrColorByDefaultKey {
			o.attrColorByDefault = color

			return
		}

		o.attrKey2color[key] = color
	}
}
