package noplocker_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-core/mrlock"
	"github.com/mondegor/go-core/mrlock/noplocker"
)

// Make sure the Locker conforms with the mrlock.Locker interface.
func TestLockerImplementsLocker(t *testing.T) {
	t.Parallel()

	assert.Implements(t, (*mrlock.Locker)(nil), &noplocker.Locker{})
}
