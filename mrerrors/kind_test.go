package mrerrors_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-sysmess/mrerrors"
)

func TestErrorKind_String(t *testing.T) {
	t.Parallel()

	kind := mrerrors.ErrorKind(0)
	require.Equal(t, "UNKNOWN", kind.String())

	kind = mrerrors.ErrorKindInternal
	require.Equal(t, "INTERNAL", kind.String())

	kind = mrerrors.ErrorKindSystem
	require.Equal(t, "SYSTEM", kind.String())

	kind = mrerrors.ErrorKindUser
	require.Equal(t, "USER", kind.String())

	kind = mrerrors.ErrorKind(4)
	require.Equal(t, "UNKNOWN", kind.String())
}
