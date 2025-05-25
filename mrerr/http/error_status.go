package http

import (
	"errors"
	"net/http"

	"github.com/mondegor/go-sysmess/mrerr"
	"github.com/mondegor/go-sysmess/mrerr/mr"
)

type (
	// ErrorStatusGetter - только для: 4XX, 5XX.
	ErrorStatusGetter struct {
		unexpectedStatus int
	}
)

// NewErrorStatusGetter - создаёт объект ErrorStatusGetter.
func NewErrorStatusGetter(unexpectedStatus int) *ErrorStatusGetter {
	return &ErrorStatusGetter{
		unexpectedStatus: unexpectedStatus,
	}
}

// ErrorStatus - возвращает http код ответа на основе проанализированного типа ошибки и самой ошибки.
func (g *ErrorStatusGetter) ErrorStatus(analyzedKind mrerr.ErrorKind, err error) int {
	if analyzedKind == mrerr.ErrorKindInternal {
		return http.StatusInternalServerError
	}

	if analyzedKind == mrerr.ErrorKindSystem {
		return http.StatusServiceUnavailable
	}

	// если ошибка явно необработанна разработчиком (не обёрнута в InstantError),
	// то вместо 500 статуса отображается указанный g.unexpectedStatus
	if analyzedKind != mrerr.ErrorKindUnknown {
		return g.unexpectedStatus
	}

	// далее обрабатываются статусы пользовательских ошибок

	if mr.ErrHttpClientUnauthorized.Is(err) { // игнорируются вложенные ошибки
		return http.StatusUnauthorized
	}

	if mr.ErrHttpAccessForbidden.Is(err) || errors.Is(err, mr.ErrUseCaseAccessForbidden) {
		return http.StatusForbidden
	}

	if mr.ErrHttpResourceNotFound.Is(err) || errors.Is(err, mr.ErrUseCaseEntityNotFound) {
		return http.StatusNotFound
	}

	if errors.Is(err, mr.ErrHttpRequestParseData) {
		return http.StatusUnprocessableEntity
	}

	if errors.Is(err, mr.ErrUseCaseEntityVersionInvalid) {
		return http.StatusConflict
	}

	return http.StatusBadRequest
}
