package slog

import (
	"io"
	"log/slog"
)

const (
	// AttrColorByDefaultKey - ключ используемый в WithColorizeAttr
	// для установки цветов ключей и значений по умолчанию.
	AttrColorByDefaultKey = "attrColorByDefault"
)

type (
	// Option - настройка объекта LoggerAdapter.
	Option func(o *options)
)

// WithWriter - устанавливает опцию stdout для LoggerAdapter.
func WithWriter(value io.Writer) Option {
	return func(o *options) {
		if o.stdout != nil {
			o.stdout = value
		}
	}
}

// WithLevel - устанавливает уровень логирования для LoggerAdapter.
func WithLevel(value string) Option {
	return func(o *options) {
		if value != "" {
			o.levelString = value
		}
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
		if value != "" {
			o.timeFormat = value
		}
	}
}

// WithReplaceArgs - устанавливает функцию замены аргумента сообщения для LoggerAdapter.
func WithReplaceArgs(value func(arg any) (newArg any)) Option {
	return func(o *options) {
		o.replaceArg = value
	}
}

// WithMiddlewareHandler - устанавливает middleware для LoggerAdapter.
func WithMiddlewareHandler(value ...func(next slog.Handler) slog.Handler) Option {
	return func(o *options) {
		o.middlewareHandlers = append(o.middlewareHandlers, value...)
	}
}

// WithReplaceAttrs - устанавливает функцию замены атрибутов для LoggerAdapter.
func WithReplaceAttrs(value func(attr slog.Attr) (newAttr slog.Attr)) Option {
	return func(o *options) {
		if value != nil {
			o.replaceAttr = value
		}
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
