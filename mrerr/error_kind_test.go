package mrerr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorKind_String(t *testing.T) {
	t.Parallel()

	kind := ErrorKindInternal
	assert.Equal(t, kind.String(), "Internal")

	kind = ErrorKindSystem
	assert.Equal(t, kind.String(), "System")

	kind = ErrorKindUser
	assert.Equal(t, kind.String(), "User")

	kind = ErrorKind(3)
	assert.Equal(t, kind.String(), "Unknown")
}
