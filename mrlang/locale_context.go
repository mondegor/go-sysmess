package mrlang

import (
	"context"
)

type (
	ctxLocaleKey struct{}
)

// WithContext - возвращает указанный контекст с обогащённый объектом Locale.
func WithContext(ctx context.Context, locale *Locale) context.Context {
	return context.WithValue(ctx, ctxLocaleKey{}, locale)
}

// Ctx - возвращает объект Locale из указанного контекста.
func Ctx(ctx context.Context) *Locale {
	if value, ok := ctx.Value(ctxLocaleKey{}).(*Locale); ok {
		return value
	}

	return stubLocale
}
