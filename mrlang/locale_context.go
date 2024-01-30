package mrlang

import (
	"context"
)

type (
	ctxLocaleKey struct{}
)

func WithContext(ctx context.Context, locale *Locale) context.Context {
	return context.WithValue(ctx, ctxLocaleKey{}, locale)
}

func Ctx(ctx context.Context) *Locale {
	if value, ok := ctx.Value(ctxLocaleKey{}).(*Locale); ok {
		return value
	}

	return stubLocale
}
