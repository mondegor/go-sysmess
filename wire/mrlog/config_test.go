package mrlog_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-core/mrtrace"
	"github.com/mondegor/go-core/wire/mrlog"
)

// TestDefaultProcessIDs_IsolatedBetweenCalls - проверяет, что каждый вызов
// возвращает независимый срез: изменение результата одним вызывающим
// не должно быть видно остальным.
func TestDefaultProcessIDs_IsolatedBetweenCalls(t *testing.T) {
	t.Parallel()

	first := mrlog.DefaultProcessIDs()
	require.NotEmpty(t, first)

	first[0].Key = "modified"

	second := mrlog.DefaultProcessIDs()
	require.NotEmpty(t, second)

	assert.Equal(t, mrtrace.KeyCorrelationID, second[0].Key)
}

// TestDefaultProcessIDs_Keys - фиксирует состав и порядок ключей по умолчанию.
func TestDefaultProcessIDs_Keys(t *testing.T) {
	t.Parallel()

	got := mrlog.DefaultProcessIDs()

	keys := make([]string, 0, len(got))

	for _, item := range got {
		require.NotNil(t, item.GetID, "GetID is not set for key %q", item.Key)
		require.NotNil(t, item.WithID, "WithID is not set for key %q", item.Key)

		keys = append(keys, item.Key)
	}

	assert.Equal(
		t,
		[]string{
			mrtrace.KeyCorrelationID,
			mrtrace.KeyRequestID,
			mrtrace.KeyProcessID,
			mrtrace.KeyWorkerID,
			mrtrace.KeyTaskID,
		},
		keys,
	)
}
