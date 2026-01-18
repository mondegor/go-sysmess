package kind_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-sysmess/errors/kind"
)

func TestEnum_String(t *testing.T) {
	t.Parallel()

	value := kind.Enum(0)
	require.Equal(t, "UNKNOWN", value.String())

	value = kind.Internal
	require.Equal(t, "INTERNAL", value.String())

	value = kind.System
	require.Equal(t, "SYSTEM", value.String())

	value = kind.User
	require.Equal(t, "USER", value.String())

	value = kind.Enum(4)
	require.Equal(t, "UNKNOWN", value.String())
}
