package onstartup_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-sysmess/mrprocess/onstartup"
)

type process interface {
	Caption() string
	ReadyTimeout() time.Duration
	Start(ctx context.Context, ready func()) error
	Shutdown(ctx context.Context) error
}

// Make sure the Process conforms with the process interface.
func TestProcessImplementsProcess(t *testing.T) {
	t.Parallel()

	assert.Implements(t, (*process)(nil), &onstartup.Process{})
}
