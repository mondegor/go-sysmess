package mrerr_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-sysmess/mrerr"
)

func TestNewCustomError_UserInstantError(t *testing.T) {
	t.Parallel()

	e := mrerr.NewKindUser("test-code", "").New()
	got := mrerr.NewCustomError("test-custom-code", e)
	require.True(t, got.IsValid())
	require.Equal(t, "test-custom-code", got.CustomCode())
	assert.ErrorIs(t, got.Err(), e)
}

func TestNewCustomError_InternalInstantError(t *testing.T) {
	t.Parallel()

	e := mrerr.NewKindInternal("").New()
	got := mrerr.NewCustomError("test-custom-code", e)
	require.False(t, got.IsValid())
	require.Equal(t, "test-custom-code", got.CustomCode())
	require.ErrorIs(t, got.Err(), mrerr.ErrCustomErrorHasInternalError)
	assert.ErrorIs(t, got.Err().Unwrap(), e)
}

func TestNewCustomError_UserProtoInstantError(t *testing.T) {
	t.Parallel()

	proto := mrerr.NewKindUser("test-code", "")
	got := mrerr.NewCustomError("test-custom-code", proto)
	require.True(t, got.IsValid())
	require.Equal(t, "test-custom-code", got.CustomCode())
	assert.ErrorIs(t, got.Err(), proto)
}

func TestNewCustomError_InternalProtoInstantError(t *testing.T) {
	t.Parallel()

	proto := mrerr.NewKindInternal("")
	got := mrerr.NewCustomError("test-custom-code", proto)
	require.False(t, got.IsValid())
	require.Equal(t, "test-custom-code", got.CustomCode())
	require.ErrorIs(t, got.Err(), mrerr.ErrCustomErrorHasInternalError)
	assert.ErrorIs(t, got.Err().Unwrap(), proto)
}

func TestNewCustomError_NoWrappedError(t *testing.T) {
	t.Parallel()

	got := mrerr.NewCustomError("test-custom-code", errors.New("no-wrapped-error"))
	require.False(t, got.IsValid())
	require.Equal(t, "test-custom-code", got.CustomCode())
	assert.ErrorIs(t, got.Err(), mrerr.ErrCustomErrorHasNoWrappedError)
}

func TestNewCustomError_NilError(t *testing.T) {
	t.Parallel()

	got := mrerr.NewCustomError("test-custom-code", nil)
	require.False(t, got.IsValid())
	require.Equal(t, "test-custom-code", got.CustomCode())
	assert.ErrorIs(t, got.Err(), mrerr.ErrCustomErrorHasNilError)
}
