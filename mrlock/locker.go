package mrlock

import (
	"context"
	"time"
)

const (
	// DefaultExpiry - истечение срока блокировки по умолчанию.
	DefaultExpiry = time.Second
)

type (
	// Locker - интерфейс блокировщика указанного ключа.
	// Позволяет захватывать блокировку с временем жизни по умолчанию или с заданным временем истечения.
	// Возвращает функцию для освобождения блокировки.
	Locker interface {
		Lock(ctx context.Context, key string) (unlock func(), err error)
		LockWithExpiry(ctx context.Context, key string, expiry time.Duration) (unlock func(), err error)
	}
)
