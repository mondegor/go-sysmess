package slog_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-core/wire/mrlog"
	wireslog "github.com/mondegor/go-core/wire/mrlog/slog"
)

// TestInitLogger_EmptyConfigUsesAdapterDefaults - проверяет, что незаданные
// TimeFormat и TimeZone не приводят к ошибке, а отдаются на откуп умолчаниям
// slog.NewLoggerAdapter: wire-слой их не дублирует, поэтому пустой конфиг
// обязан оставаться рабочим.
func TestInitLogger_EmptyConfigUsesAdapterDefaults(t *testing.T) {
	t.Parallel()

	logger, err := wireslog.InitLogger(mrlog.LoggerConfig{})

	require.NoError(t, err)
	require.NotNil(t, logger)
}

// TestInitLogger_TimeZoneIsPassedThrough - проверяет, что заданный часовой пояс
// доходит до адаптера, то есть условное добавление опции ничего не потеряло.
func TestInitLogger_TimeZoneIsPassedThrough(t *testing.T) {
	t.Parallel()

	_, err := wireslog.InitLogger(mrlog.LoggerConfig{TimeZone: "Nowhere/Nowhere"})

	// несуществующий пояс обязан развалить создание логгера: если бы опция
	// не добавлялась, конструктор молча взял бы умолчание и вернул nil
	require.Error(t, err)
	assert.Contains(t, err.Error(), "Nowhere/Nowhere", "error must mention the rejected timezone")
}

// TestInitLogger_TimeFormatIsPassedThrough - то же для формата времени.
func TestInitLogger_TimeFormatIsPassedThrough(t *testing.T) {
	t.Parallel()

	_, err := wireslog.InitLogger(mrlog.LoggerConfig{TimeFormat: "NotAFormat"})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "NotAFormat", "error must mention the rejected time format")
}
