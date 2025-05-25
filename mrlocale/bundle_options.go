package mrlocale

import (
	"golang.org/x/text/language"
)

type (
	// BundleOption - настройка объекта Bundle.
	BundleOption func(e *bundleOptions)
)

// WithMessageProvider - устанавливает провайдер для локализации сообщений.
func WithMessageProvider(createFunc func(langs []language.Tag) (MessageProvider, error)) BundleOption {
	return func(o *bundleOptions) {
		o.createProvider = createFunc
	}
}

// WithFormatMessage - устанавливает функцию для подстановки значений аргументов в сообщения.
func WithFormatMessage(fn func(msg string, args []any) (newMsg string, newArgs []any)) BundleOption {
	return func(o *bundleOptions) {
		o.formatMessage = fn
	}
}

// WithFormatError - устанавливает функцию для подстановки значений аргументов в сообщения об ошибках.
func WithFormatError(fn func(err error) (msg string, args []any)) BundleOption {
	return func(o *bundleOptions) {
		o.formatError = fn
	}
}

// WithMessagesDomain - устанавливает домен, который будет использоваться при локализации сообщений.
func WithMessagesDomain(value string) BundleOption {
	return func(o *bundleOptions) {
		if value != "" {
			o.messagesDomain = value
		}
	}
}

// WithErrorsDomain - устанавливает домен, который будет использоваться при локализации сообщений об ошибках.
func WithErrorsDomain(value string) BundleOption {
	return func(o *bundleOptions) {
		if value != "" {
			o.errorsDomain = value
		}
	}
}

// WithLanguages - добавляет языки для локализации сообщений.
func WithLanguages(values ...string) BundleOption {
	return func(o *bundleOptions) {
		o.languages = append(o.languages, values...)
	}
}

// WithDefaultLanguage - устанавливает язык по умолчанию.
func WithDefaultLanguage(value string) BundleOption {
	return func(o *bundleOptions) {
		if value != "" {
			o.defaultLanguage = value
		}
	}
}
