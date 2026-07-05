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

func TestNewCustom(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name           string
		inputErr       error
		wantKindUser   bool
		wantCustomCode string
		// wantUnwrap - ожидаемая ошибка под Unwrap(); если nil, ожидается сам inputErr
		wantUnwrap error
		// wantContains - sentinel-ошибка, чьё сообщение должно входить в got.Error(); nil - проверка пропускается
		wantContains error
	}

	tests := []testCase{
		{
			name:           "user error",
			inputErr:       &testError{kind: kind.User, code: "MyCode"},
			wantKindUser:   true,
			wantCustomCode: "MyCode/test-custom-code",
		},
		{
			name:           "internal error",
			inputErr:       &testError{kind: kind.Internal},
			wantKindUser:   false,
			wantCustomCode: "test-custom-code",
			wantContains:   custom.ErrHasInternalError,
		},
		{
			name:           "system error",
			inputErr:       &testError{kind: kind.System},
			wantKindUser:   false,
			wantCustomCode: "test-custom-code",
			wantContains:   custom.ErrHasSystemError,
		},
		{
			name:           "no wrapped error",
			inputErr:       errors.New("no-wrapped-error"),
			wantKindUser:   false,
			wantCustomCode: "test-custom-code",
			wantContains:   custom.ErrHasUnexpectedError,
		},
		{
			name:           "nil error",
			inputErr:       nil,
			wantKindUser:   false,
			wantCustomCode: "test-custom-code",
			wantUnwrap:     custom.ErrHasNilError,
			wantContains:   custom.ErrHasNilError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := custom.New(tt.inputErr, "test-custom-code")
			require.Equal(t, tt.wantKindUser, got.IsKindUser())
			require.Equal(t, tt.wantCustomCode, got.CustomCode())

			wantUnwrap := tt.inputErr
			if tt.wantUnwrap != nil {
				wantUnwrap = tt.wantUnwrap
			}

			require.ErrorIs(t, got.Unwrap(), wantUnwrap)

			if tt.wantContains != nil {
				assert.Contains(t, got.Error(), tt.wantContains.Error())
			}
		})
	}
}
