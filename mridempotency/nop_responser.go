package mridempotency

import "net/http"

type (
	// responser - заглушка (no-op) реализующая интерфейс Responser идемпотентности.
	// Всегда возвращает пустое содержимое и HTTP-код 200 OK.
	responser struct{}
)

// NopResponser - создаёт no-op Responser.
func NopResponser() Responser {
	return responser{}
}

// StatusCode - всегда возвращает HTTP-код ответа 200 (OK).
func (r responser) StatusCode() int {
	return http.StatusOK
}

// Content - всегда возвращает nil (пустое тело ответа).
func (r responser) Content() []byte {
	return nil
}
