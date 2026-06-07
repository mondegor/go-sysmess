package mrrun

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"time"
)

const (
	defaultTimeout = 5 * time.Second
)

type (
	// HealthProbe - обёртка для проверки работоспособности процесса (сервиса).
	// Выполняет указанную функцию проверки с таймаутом и защитой от паник.
	//
	// Используется в healthcheck-эндпоинтах для мониторинга состояния
	// зависимостей приложения (базы данных, кэши, внешние сервисы и т.д.).
	HealthProbe struct {
		caption string
		check   func(ctx context.Context) error // функция проверки работоспособности процесса
		timeout time.Duration                   // таймаут, после которого функция check должна остановить своё выполнение
		logger  logger
	}
)

// ErrInternalProbeHasPanic - во время выполнения пробы произошла паника.
var ErrInternalProbeHasPanic = errors.New("probe has panic")

// NewHealthProbe - создаёт пробу для отслеживания работоспособности процесса.
// Если timeout равен 0, используется defaultTimeout (5 секунд).
func NewHealthProbe(logger logger, caption string, check func(ctx context.Context) error, timeout time.Duration) *HealthProbe {
	if timeout <= 0 {
		timeout = defaultTimeout
	}

	return &HealthProbe{
		caption: caption,
		check:   check,
		timeout: timeout,
		logger:  logger,
	}
}

// Caption - возвращает название пробы в свободной форме.
func (p *HealthProbe) Caption() string {
	return p.caption
}

// Check - выполняет проверку работоспособности процесса с защитой от паник.
//
// Особенности работы:
//  1. Создаёт контекст с таймаутом (p.timeout);
//  2. Выполняет функцию check в recover-обёртке;
//  3. При панике логирует ошибку со стеком и возвращает ошибку;
//  4. При истечении таймаута контекст автоматически отменяется;
//
// Возвращает ошибку от check-функции или ошибку при панике.
func (p *HealthProbe) Check(ctx context.Context) (err error) {
	ctx, cancel := context.WithTimeout(ctx, p.timeout)

	defer func() {
		cancel()

		if rvr := recover(); rvr != nil {
			p.logger.Error(
				ctx,
				"HealthProbe panic",
				"probe", p.caption,
				"recover", rvr,
				"stack_trace", string(debug.Stack()),
			)

			err = fmt.Errorf("%w [caption=%s]", ErrInternalProbeHasPanic, p.caption)
		}
	}()

	return p.check(ctx)
}
