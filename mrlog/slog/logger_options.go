package slog

import (
	"io"
	stdlog "log/slog"

	"github.com/mondegor/go-core/mrlog/level"
)

const (
	// AttrColorByDefaultKey - ключ используемый в WithColorizeAttr
	// для установки цветов ключей и значений по умолчанию.
	AttrColorByDefaultKey = "attrColorByDefault"
)

type (
	// Option - функция-опция для настройки LoggerAdapter.
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

// WithWriter - задаёт поток вывода для логирования (по умолчанию os.Stdout).
func WithWriter(value io.Writer) Option {
	return func(o *options) {
		o.stdout = value
	}
}

// WithLevel - задаёт минимальный уровень логирования (DEBUG, INFO, WARN, ERROR).
// Сообщения с более низким уровнем будут отфильтрованы.
func WithLevel(value string) Option {
	return func(o *options) {
		o.levelString = value
	}
}

// WithJsonFormat - включает вывод логов в формате JSON.
// При включении отключает цветной режим (colorMode).
func WithJsonFormat(value bool) Option {
	return func(o *options) {
		o.jsonFormat = value
	}
}

// WithTimeFormat - задаёт формат вывода времени в логах.
// Поддерживаемые именованные форматы: "RFC3339", "RFC3339Nano", "DateTime", "TimeOnly", "Kitchen".
func WithTimeFormat(value string) Option {
	return func(o *options) {
		o.timeFormat = value
	}
}

// WithMiddlewareHandler - добавляет один или несколько middleware-обработчиков.
// Middleware применяются в порядке добавления и позволяют модифицировать slog.Handler.
func WithMiddlewareHandler(value ...func(next stdlog.Handler) stdlog.Handler) Option {
	return func(o *options) {
		o.middlewareHandlers = append(o.middlewareHandlers, value...)
	}
}

// WithReplaceAttrs - задаёт функцию для кастомной обработки атрибутов перед выводом.
// Позволяет изменять ключи, значения и форматирование атрибутов.
func WithReplaceAttrs(value func(attr stdlog.Attr) (newAttr stdlog.Attr)) Option {
	return func(o *options) {
		o.replaceAttr = value
	}
}

// WithColorMode - включает/отключает цветной вывод в консоль.
// Игнорируется, если включен JsonFormat.
func WithColorMode(value bool) Option {
	return func(o *options) {
		o.colorMode = value
	}
}

// WithColorizeAttr - задаёт цвета для ключа/значения указанного атрибута.
// Параметры:
//   - attrKey - имя атрибута, которому назначаются цвета;
//   - colorKey, colorValue - ANSI-коды цветов из пакета color.
//
// При attrKey=AttrColorByDefaultKey цвета назначаются всем атрибутам,
// для которых явно не были определены цвета.
// Работает только при включенном ColorMode и отключенном JsonFormat.
func WithColorizeAttr(attrKey, colorKey, colorValue string) Option {
	return func(o *options) {
		color := attrColor{
			key:   colorKey,
			value: colorValue,
		}

		if attrKey == AttrColorByDefaultKey {
			o.attrColorByDefault = color

			return
		}

		o.attrKey2color[attrKey] = color
	}
}
