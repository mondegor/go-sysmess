package slog_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-sysmess/mrlog"
	"github.com/mondegor/go-sysmess/mrlog/slog"
)

// Make sure the LoggerAdapter conforms with the Logger interface.
func TestLoggerAdapterImplementsLogger(t *testing.T) {
	assert.Implements(t, (*mrlog.Logger)(nil), &slog.LoggerAdapter{})
}
