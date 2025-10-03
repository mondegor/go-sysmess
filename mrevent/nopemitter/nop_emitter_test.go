package nopemitter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-sysmess/mrevent"
	"github.com/mondegor/go-sysmess/mrevent/nopemitter"
)

// Make sure the Emitter conforms with the mrevent.Emitter interface.
func TestEmitterImplementsEventEmitter(t *testing.T) {
	assert.Implements(t, (*mrevent.Emitter)(nil), &nopemitter.Emitter{})
}
