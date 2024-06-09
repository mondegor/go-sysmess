package mrerr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorKind_String(t *testing.T) {
	t.Parallel()

	kind := ErrorKindInternal
	assert.Equal(t, "Internal", kind.String())

	kind = ErrorKindSystem
	assert.Equal(t, "System", kind.String())

	kind = ErrorKindUser
	assert.Equal(t, "User", kind.String())

	kind = ErrorKind(3)
	assert.Equal(t, "Unknown", kind.String())
}
