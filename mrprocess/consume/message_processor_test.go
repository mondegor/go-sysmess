package consume_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-core/mrprocess/consume"
)

type process interface {
	Caption() string
	ReadyTimeout() time.Duration
	Start(ctx context.Context, ready func()) error
	Shutdown(ctx context.Context) error
}

// Make sure the MessageProcessor conforms with the process interface.
func TestMessageProcessorImplementsProcess(t *testing.T) {
	t.Parallel()

	assert.Implements(t, (*process)(nil), &consume.MessageProcessor[any]{})
}
