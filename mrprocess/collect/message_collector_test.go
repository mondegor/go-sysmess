package collect_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-core/mrprocess/collect"
)

type process interface {
	Caption() string
	ReadyTimeout() time.Duration
	Start(ctx context.Context, ready func()) error
	Shutdown(ctx context.Context) error
}

// Make sure the MessageCollector conforms with the process interface.
func TestMessageCollectorImplementsProcess(t *testing.T) {
	t.Parallel()

	assert.Implements(t, (*process)(nil), &collect.MessageCollector[any]{})
}
