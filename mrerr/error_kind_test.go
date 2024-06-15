package mrerr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-sysmess/mrerr"
)

func TestErrorKind_String(t *testing.T) {
	t.Parallel()

	kind := mrerr.ErrorKindInternal
	assert.Equal(t, "Internal", kind.String())

	kind = mrerr.ErrorKindSystem
	assert.Equal(t, "System", kind.String())

	kind = mrerr.ErrorKindUser
	assert.Equal(t, "User", kind.String())

	kind = mrerr.ErrorKind(3)
	assert.Equal(t, "Unknown", kind.String())
}
