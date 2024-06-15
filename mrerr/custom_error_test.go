package mrerr_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-sysmess/mrerr"
)

func TestNewCustomError_UserAppError(t *testing.T) {
	t.Parallel()

	e := mrerr.NewProto("test-code", mrerr.ErrorKindUser, "").New()
	got := mrerr.NewCustomError("test-custom-code", e)
	assert.True(t, got.IsValid())
	assert.Equal(t, "test-custom-code", got.CustomCode())
	require.ErrorIs(t, got.Err(), e)
}

func TestNewCustomError_InternalAppError(t *testing.T) {
	t.Parallel()

	e := mrerr.NewProto("test-code", mrerr.ErrorKindInternal, "").New()
	got := mrerr.NewCustomError("test-custom-code", e)
	assert.False(t, got.IsValid())
	assert.Equal(t, "test-custom-code", got.CustomCode())
	require.ErrorIs(t, got.Err(), mrerr.ErrCustomErrorHasInternalError)
	require.ErrorIs(t, got.Err().Unwrap(), e)
}

func TestNewCustomError_UserProtoAppError(t *testing.T) {
	t.Parallel()

	proto := mrerr.NewProto("test-code", mrerr.ErrorKindUser, "")
	got := mrerr.NewCustomError("test-custom-code", proto)
	assert.True(t, got.IsValid())
	assert.Equal(t, "test-custom-code", got.CustomCode())
	assert.ErrorIs(t, got.Err(), proto)
}

func TestNewCustomError_InternalProtoAppError(t *testing.T) {
	t.Parallel()

	proto := mrerr.NewProto("test-code", mrerr.ErrorKindInternal, "")
	got := mrerr.NewCustomError("test-custom-code", proto)
	assert.False(t, got.IsValid())
	assert.Equal(t, "test-custom-code", got.CustomCode())
	require.ErrorIs(t, got.Err(), mrerr.ErrCustomErrorHasInternalError)
	require.ErrorIs(t, got.Err().Unwrap(), proto)
}

func TestNewCustomError_NilError(t *testing.T) {
	t.Parallel()

	got := mrerr.NewCustomError("test-custom-code", nil)
	assert.False(t, got.IsValid())
	assert.Equal(t, "test-custom-code", got.CustomCode())
	assert.ErrorIs(t, got.Err(), mrerr.ErrErrorIsNilPointer)
}

func TestNewCustomError_NoWrappedError(t *testing.T) {
	t.Parallel()

	got := mrerr.NewCustomError("test-custom-code", errors.New("no-wrapped-error"))
	assert.False(t, got.IsValid())
	assert.Equal(t, "test-custom-code", got.CustomCode())
	require.ErrorIs(t, got.Err(), mrerr.ErrCustomErrorHasNoWrappedError)
}
