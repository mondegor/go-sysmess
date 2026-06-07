package mrlock

import (
	"errors"
)

var (
	// ErrLockKeyNotObtained - не удалось захватить ключ блокировки.
	// Возникает, когда ключ уже заблокирован другим процессом или истёк срок предыдущей блокировки.
	ErrLockKeyNotObtained = errors.New("lock key not obtained")

	// ErrLockKeyNotHeld - ключ блокировки не принадлежит текущему владельцу.
	// Возникает при попытке освободить ключ, который не был заблокирован текущим процессом.
	ErrLockKeyNotHeld = errors.New("lock key not held")
)
