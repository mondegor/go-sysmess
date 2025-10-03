package process

import (
	"context"
	"fmt"
)

type (
	// ContextManager - отвечает за установку ID процессов
	// в контекст и за доступ к ним используемых в трейсинге.
	ContextManager struct {
		keyGetID  map[string]func(ctx context.Context) string
		keyWithID map[string]func(ctx context.Context, id string) context.Context

		newContextWithIDs    []GetIDWithID
		extractKeysValues    []KeyGetID
		extractCorrelationID []KeyGetID

		idGenerator idGenerator
		logger      logger
	}

	idGenerator interface {
		GenerateID() string
	}

	logger interface {
		Warn(ctx context.Context, msg string, args ...any)
	}
)

// NewContextManager - создаёт объект ContextManager.
func NewContextManager(
	keyGetIDWithID []KeyGetIDWithID,
	correlationKeys []string,
	idGenerator idGenerator,
	logger logger,
) (*ContextManager, error) {
	cm := ContextManager{
		logger:      logger,
		idGenerator: idGenerator,

		keyGetID:  make(map[string]func(ctx context.Context) string, len(keyGetIDWithID)),
		keyWithID: make(map[string]func(ctx context.Context, id string) context.Context, len(keyGetIDWithID)),

		newContextWithIDs:    make([]GetIDWithID, 0, len(keyGetIDWithID)),
		extractKeysValues:    make([]KeyGetID, 0, len(keyGetIDWithID)),
		extractCorrelationID: make([]KeyGetID, 0, len(correlationKeys)),
	}

	for _, v := range keyGetIDWithID {
		if _, ok := cm.keyGetID[v.Key]; ok {
			return nil, fmt.Errorf("duplicated key %q", v.Key)
		}

		cm.keyGetID[v.Key] = v.GetID
		cm.keyWithID[v.Key] = v.WithID

		cm.newContextWithIDs = append(
			cm.newContextWithIDs,
			GetIDWithID{
				GetID:  v.GetID,
				WithID: v.WithID,
			},
		)

		cm.extractKeysValues = append(
			cm.extractKeysValues,
			KeyGetID{
				Key:   v.Key,
				GetID: v.GetID,
			},
		)
	}

	for _, key := range correlationKeys {
		getID, ok := cm.keyGetID[key]
		if !ok {
			return nil, fmt.Errorf("correlation key %s not found", key)
		}

		cm.extractCorrelationID = append(
			cm.extractCorrelationID,
			KeyGetID{
				Key:   key,
				GetID: getID,
			},
		)
	}

	return &cm, nil
}

// ProcessID - возвращает из контекста ID процесса по его имени key.
func (e *ContextManager) ProcessID(ctx context.Context, key string) string {
	if fn, ok := e.keyGetID[key]; ok {
		return fn(ctx)
	}

	e.logger.Warn(ctx, "process key not found", "key", key)

	return ""
}

// WithProcessID - устанавливает ID процесса в контекст по его имени key и возвращает получившийся контекст.
func (e *ContextManager) WithProcessID(ctx context.Context, key, value string) context.Context {
	if fn, ok := e.keyWithID[key]; ok {
		return fn(ctx, value)
	}

	e.logger.Warn(ctx, "process key not found", "key", key)

	return ctx
}

// WithGeneratedProcessID - генерирует ID процесса,
// устанавливает его в контекст по его имени key и возвращает получившийся контекст.
func (e *ContextManager) WithGeneratedProcessID(ctx context.Context, key string) context.Context {
	if fn, ok := e.keyWithID[key]; ok {
		return fn(ctx, e.idGenerator.GenerateID())
	}

	e.logger.Warn(ctx, "process key not found", "key", key)

	return ctx
}

// NewContextWithIDs - возвращает результат вызова функции NewContextWithIDs.
func (e *ContextManager) NewContextWithIDs(originalCtx context.Context) context.Context {
	return NewContextWithIDs(originalCtx, e.newContextWithIDs)
}

// ExtractKeysValues - возвращает результат вызова функции ExtractKeysValues.
func (e *ContextManager) ExtractKeysValues(ctx context.Context) []any {
	return ExtractKeysValues(ctx, e.extractKeysValues)
}

// ExtractCorrelationID - возвращает результат вызова функции ExtractCorrelationID.
func (e *ContextManager) ExtractCorrelationID(ctx context.Context) string {
	return ExtractCorrelationID(ctx, e.extractCorrelationID)
}
