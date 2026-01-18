package custom_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-sysmess/errors/custom"
	"github.com/mondegor/go-sysmess/errors/kind"
)

type (
	testError struct {
		kind kind.Enum
		code string
	}
)

func (e *testError) Kind() kind.Enum {
	return e.kind
}

func (e *testError) Code() string {
	return e.code
}

func (e *testError) Error() string {
	return e.kind.String() + ": " + e.code
}

func TestNewCustom_UserError(t *testing.T) {
	t.Parallel()

	e := &testError{kind: kind.User, code: "MyCode"}
	got := custom.New(e, "test-custom-code")
	require.True(t, got.IsKindUser())
	require.Equal(t, "MyCode/test-custom-code", got.CustomCode())
	assert.ErrorIs(t, got.Unwrap(), e)
}

func TestNewCustom_InternalError(t *testing.T) {
	t.Parallel()

	e := &testError{kind: kind.Internal}
	got := custom.New(e, "test-custom-code")
	require.False(t, got.IsKindUser())
	require.Equal(t, "test-custom-code", got.CustomCode())
	require.ErrorIs(t, got.Unwrap(), e)
	assert.Contains(t, got.Error(), custom.ErrHasInternalError.Error())
}

func TestNewCustom_SystemError(t *testing.T) {
	t.Parallel()

	e := &testError{kind: kind.System}
	got := custom.New(e, "test-custom-code")
	require.False(t, got.IsKindUser())
	require.Equal(t, "test-custom-code", got.CustomCode())
	require.ErrorIs(t, got.Unwrap(), e)
	assert.Contains(t, got.Error(), custom.ErrHasSystemError.Error())
}

func TestNewCustom_NoWrappedError(t *testing.T) {
	t.Parallel()

	e := errors.New("no-wrapped-error")
	got := custom.New(e, "test-custom-code")
	require.False(t, got.IsKindUser())
	require.Equal(t, "test-custom-code", got.CustomCode())
	require.ErrorIs(t, got.Unwrap(), e)
	assert.Contains(t, got.Error(), custom.ErrHasUnexpectedError.Error())
}

func TestNewCustom_NilError(t *testing.T) {
	t.Parallel()

	got := custom.New(nil, "test-custom-code")
	require.False(t, got.IsKindUser())
	require.Equal(t, "test-custom-code", got.CustomCode())
	require.NoError(t, got.Unwrap())
	assert.Contains(t, got.Error(), custom.ErrHasNilError.Error())
}
