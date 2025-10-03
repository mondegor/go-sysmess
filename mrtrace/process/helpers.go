package process

import (
	"context"
)

type (
	// KeyGetIDWithID - тройка: текстовое название ID, функция получения ID из контекста и функция установки значения в контекст.
	KeyGetIDWithID struct {
		Key    string
		GetID  func(ctx context.Context) string
		WithID func(ctx context.Context, id string) context.Context
	}

	// KeyGetID - пара: текстовое название ID и функция получения ID из контекста.
	KeyGetID struct {
		Key   string
		GetID func(ctx context.Context) string
	}

	// GetIDWithID - пара: функция получения ID из контекста и функция установки значения в контекст.
	GetIDWithID struct {
		GetID  func(ctx context.Context) string
		WithID func(ctx context.Context, id string) context.Context
	}
)

// NewContextWithIDs - возвращает новый контекст содержащий
// только все ID процессы, скопированные из указанного контекста.
func NewContextWithIDs(originalCtx context.Context, list []GetIDWithID) context.Context {
	ctx := context.Background()

	if originalCtx == nil || originalCtx == ctx {
		return ctx
	}

	for _, v := range list {
		if value := v.GetID(ctx); value != "" {
			ctx = v.WithID(ctx, value)
		}
	}

	return ctx
}

// ExtractKeysValues - возвращает попарно (key/id-value) все имеющиеся
// ID процессов из указанного контекста.
func ExtractKeysValues(ctx context.Context, list []KeyGetID) (keyValue []any) {
	if ctx == nil || ctx == context.Background() {
		return nil
	}

	keyValue = make([]any, 0, len(list))

	for _, v := range list {
		if value := v.GetID(ctx); value != "" {
			keyValue = append(keyValue, v.Key, value)
		}
	}

	return keyValue[0:len(keyValue):len(keyValue)]
}

// ExtractCorrelationID - возвращает первый попавшийся ID из указанного контекста,
// который можно использовать в качестве CorrelationID.
func ExtractCorrelationID(ctx context.Context, list []KeyGetID) string {
	for _, v := range list {
		if value := v.GetID(ctx); value != "" {
			return value
		}
	}

	return ""
}

// ToKeyGetID - возвращает преобразование []KeyGetIDWithID в []KeyGetID.
func ToKeyGetID(value []KeyGetIDWithID) []KeyGetID {
	list := make([]KeyGetID, 0, len(value))

	for _, v := range value {
		list = append(
			list,
			KeyGetID{
				Key:   v.Key,
				GetID: v.GetID,
			},
		)
	}

	return list
}

// ToGetIDWithID - возвращает преобразование []KeyGetIDWithID в []GetIDWithID.
func ToGetIDWithID(value []KeyGetIDWithID) []GetIDWithID {
	list := make([]GetIDWithID, 0, len(value))

	for _, v := range value {
		list = append(
			list,
			GetIDWithID{
				GetID:  v.GetID,
				WithID: v.WithID,
			},
		)
	}

	return list
}
