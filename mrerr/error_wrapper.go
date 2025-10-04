package mrerr

import (
	"fmt"

	"github.com/mondegor/go-sysmess/mrargs"
)

const (
	errorSourceKey       = "errsource"
	errorSourceSeparator = "," // разделитель между источниками ошибок.
	defaultErrorSource   = "general.errors"
)

type (
	// ErrorWrapper - помощник для оборачивания ошибок.
	ErrorWrapper interface {
		WrapError(err error, attrs ...any) error
	}

	// UseCaseErrorWrapper - помощник для оборачивания ошибок
	// в зависимости от их вида. Используется в UseCase.
	UseCaseErrorWrapper interface {
		IsNotFoundOrNotAffectedError(err error) bool
		WrapErrorFailed(err error, attrs ...any) error
		WrapErrorNotFoundOrFailed(err error, attrs ...any) error
	}

	// UserErrorWrapper - помощник для оборачивания пользовательских ошибок.
	UserErrorWrapper interface {
		WrapError(err error, name string) error
	}
)

type (
	// errorWrapper - расширяет возможности ErrorWrapper добавляя к нему источник данных.
	errorWrapper struct {
		base        ErrorWrapper
		sourceValue string
	}
)

// NewErrorWrapper - создаёт объект errorWrapper.
func NewErrorWrapper(base ErrorWrapper, source string) ErrorWrapper {
	if source == "" {
		source = defaultErrorSource
	}

	return &errorWrapper{
		base:        base,
		sourceValue: source,
	}
}

// WrapError - возвращает ошибку с указанием источника.
func (u *errorWrapper) WrapError(err error, attrs ...any) error {
	return u.base.WrapError(err, addSourceToAttrs(u.sourceValue, attrs)...) //nolint:wrapcheck
}

type (
	// useCaseErrorWrapper - расширяет возможности UseCaseErrorWrapper добавляя к нему источник данных.
	useCaseErrorWrapper struct {
		base        UseCaseErrorWrapper
		sourceValue string
	}
)

// NewUseCaseErrorWrapper - создаёт объект useCaseErrorWrapper.
func NewUseCaseErrorWrapper(base UseCaseErrorWrapper, source string) UseCaseErrorWrapper {
	if source == "" {
		source = defaultErrorSource
	}

	return &useCaseErrorWrapper{
		base:        base,
		sourceValue: source,
	}
}

// IsNotFoundOrNotAffectedError - сообщает, связанна ли ошибка с отсутствием запрошенной записи,
// или она была найдена, но её изменение не потребовалось.
func (u *useCaseErrorWrapper) IsNotFoundOrNotAffectedError(err error) bool {
	return u.base.IsNotFoundOrNotAffectedError(err)
}

// WrapErrorFailed - возвращает обёрнутую ошибку с указанием источника.
func (u *useCaseErrorWrapper) WrapErrorFailed(err error, attrs ...any) error {
	return u.base.WrapErrorFailed(err, addSourceToAttrs(u.sourceValue, attrs)...) //nolint:wrapcheck
}

// WrapErrorNotFoundOrFailed - возвращает обёрнутую ошибку с указанием источника.
func (u *useCaseErrorWrapper) WrapErrorNotFoundOrFailed(err error, attrs ...any) error {
	return u.base.WrapErrorNotFoundOrFailed(err, addSourceToAttrs(u.sourceValue, attrs)...) //nolint:wrapcheck
}

func addSourceToAttrs(value string, attrs ...any) []any {
	return mrargs.AddKeyValue(
		errorSourceKey,
		func(index int, item any) (newitem any) {
			if index < 0 {
				return value
			}

			var itemStr string

			// если значение в виде строки, то оно дополняется,
			// иначе заменяется на новое
			switch v := item.(type) {
			case string:
				itemStr = v
			case fmt.Stringer:
				itemStr = v.String()
			}

			if itemStr != "" {
				value += errorSourceSeparator + itemStr
			}

			return value
		},
		attrs...,
	)
}
