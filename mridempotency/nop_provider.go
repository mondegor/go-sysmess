package mridempotency

import (
	"context"
	"time"
)

const (
	nopProviderName = "IdempotencyNopProvider"
	defaultExpiry   = time.Second
)

type (
	// nopProvider - заглушка (no-op) реализующая интерфейс Provider идемпотентности.
	// Все методы всегда возвращают успешный результат, не выполняя реальной работы.
	nopProvider struct {
		tracer tracer
	}

	tracer interface {
		Trace(ctx context.Context, args ...any)
	}
)

// NopProvider - создаёт no-op провайдер идемпотентности для тестирования и отладки.
func NopProvider(tracer tracer) Provider {
	return &nopProvider{
		tracer: tracer,
	}
}

// Validate - всегда возвращает nil, эмулируя успешную валидацию любого ключа.
func (p *nopProvider) Validate(_ string) error {
	return nil
}

// Lock - эмулирует блокировку ключа идемпотентности без реальной синхронизации.
func (p *nopProvider) Lock(ctx context.Context, key string) (unlock func(), err error) {
	p.traceCmd(ctx, "Lock:"+defaultExpiry.String(), key)

	return func() {
		p.traceCmd(ctx, "Unlock", key)
	}, nil
}

// Get - всегда возвращает пустой ответ (NopResponser).
func (p *nopProvider) Get(ctx context.Context, key string) (Responser, error) {
	p.traceCmd(ctx, "Get:"+key, key)

	return NopResponser(), nil
}

// Save - эмулирует сохранение ответа без реальной записи в хранилище.
func (p *nopProvider) Save(ctx context.Context, key string, result Responser) error {
	p.traceCmd(ctx, "Save:"+key, key)

	p.tracer.Trace(
		ctx,
		"statusCode", result.StatusCode(),
		"body", result.Content(),
	)

	return nil
}

func (p *nopProvider) traceCmd(ctx context.Context, command, key string) {
	p.tracer.Trace(
		ctx,
		"source", nopProviderName,
		"cmd", command,
		"key", key,
	)
}
