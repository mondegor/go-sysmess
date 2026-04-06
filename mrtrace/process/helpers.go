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

// // NewContextWithIDs - возвращает новый контекст содержащий
// // только все ID процессы, скопированные из указанного контекста.
// func NewContextWithIDs(originalCtx context.Context, values []GetIDWithID) context.Context {
// 	ctx := context.Background()
//
// 	if originalCtx == nil || originalCtx == ctx {
// 		return ctx
// 	}
//
// 	for _, v := range values {
// 		if value := v.GetID(ctx); value != "" {
// 			ctx = v.WithID(ctx, value)
// 		}
// 	}
//
// 	return ctx
// }

// ExtractKeysValues - извлекает ID процессов из контекста по указанному списку ключей.
// Возвращает плоский слайс пар ключ/значение: ["key1", "value1", "key2", "value2", ...].
// Пропускает ключи с пустыми значениями.
// Возвращает nil для nil-контекста, context.Background() или пустого списка ключей.
func ExtractKeysValues(ctx context.Context, values []KeyGetID) (keyValue []any) {
	if ctx == nil || ctx == context.Background() || len(values) == 0 {
		return nil
	}

	keyValue = make([]any, 0, len(values))

	for _, v := range values {
		if value := v.GetID(ctx); value != "" {
			keyValue = append(keyValue, v.Key, value)
		}
	}

	return keyValue[0:len(keyValue):len(keyValue)]
}

// ExtractCorrelationID - ищет первый непустой ID в контексте из списка ключей.
// Используется для определения ID корреляции, который может храниться под разными ключами.
// Возвращает пустую строку, если ни один ключ не найден.
func ExtractCorrelationID(ctx context.Context, values []KeyGetID) string {
	for _, v := range values {
		if id := v.GetID(ctx); id != "" {
			return id
		}
	}

	return ""
}

// ToKeyGetID - преобразует слайс KeyGetIDWithID в слайс KeyGetID,
// отбрасывая функции установки значений (WithID).
func ToKeyGetID(values []KeyGetIDWithID) []KeyGetID {
	if len(values) == 0 {
		return nil
	}

	list := make([]KeyGetID, 0, len(values))

	for _, v := range values {
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

// // ToGetIDWithID - возвращает преобразование []KeyGetIDWithID в []GetIDWithID.
// func ToGetIDWithID(values []KeyGetIDWithID) []GetIDWithID {
// 	if len(values) == 0 {
// 		return nil
// 	}
//
// 	list := make([]GetIDWithID, 0, len(values))
//
// 	for _, v := range values {
// 		list = append(
// 			list,
// 			GetIDWithID{
// 				GetID:  v.GetID,
// 				WithID: v.WithID,
// 			},
// 		)
// 	}
//
// 	return list
// }
