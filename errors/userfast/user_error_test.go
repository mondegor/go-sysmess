package userfast_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-sysmess/errors/kind"
	"github.com/mondegor/go-sysmess/errors/userfast"
)

func TestNew(t *testing.T) {
	t.Parallel()

	err := userfast.New("test-code", "test-message")
	got, ok := err.(interface {
		error
		Kind() kind.Enum
	})

	require.True(t, ok)
	assert.Equal(t, kind.User, got.Kind())
	assert.ErrorContains(t, got, "#test-code - test-message")
}

func TestCode(t *testing.T) {
	t.Parallel()

	err := userfast.New("test-code", "test-message")
	got, ok := err.(interface {
		error
		Code() string
	})

	require.True(t, ok)
	assert.Equal(t, "test-code", got.Code())
}

func TestMessage(t *testing.T) {
	t.Parallel()

	err := userfast.New("test-code", "test-message")
	got, ok := err.(interface {
		error
		Message() string
	})

	require.True(t, ok)
	assert.Equal(t, "test-message", got.Message())
}

func TestProto_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		message   string
		argsNames []string
		args      []any
		want      string
	}{
		{
			name:    "test1",
			message: "my-message",
			want:    "#test-code - my-message",
		},
		{
			name:    "test2",
			message: "my-message {Key1} - {Key2}",
			want:    "#test-code - my-message {Key1} - {Key2}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := userfast.New("test-code", tt.message)
			assert.Equal(t, tt.want, e.Error())
		})
	}
}

func TestProto_Is(t *testing.T) {
	t.Parallel()

	errUserTest1 := userfast.New("test-code1", "test-message1")
	errUserTest2 := userfast.New("test-code2", "test-message2")

	tests := []struct {
		name   string
		target error
		want   bool
	}{
		{
			name:   "test1",
			target: errors.New("my-message"),
			want:   false,
		},
		{
			name:   "test2",
			target: errUserTest1,
			want:   true,
		},
		{
			name:   "test3",
			target: errUserTest2,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := errors.Is(tt.target, errUserTest1)
			assert.Equal(t, tt.want, got)
		})
	}
}

var errTestProtoInstanceAs = userfast.New("test-code1", "test-message1")

func TestProtoInstance_As_PointerProtoByLink(t *testing.T) {
	t.Parallel()

	var target userfast.ProtoError

	require.ErrorAs(t, errTestProtoInstanceAs, &target)
	assert.Equal(t, errTestProtoInstanceAs, target)
}

func TestProtoInstance_As_AnyPointerProtoByLink(t *testing.T) {
	t.Parallel()

	var target any = (userfast.ProtoError)(nil)

	require.ErrorAs(t, errTestProtoInstanceAs, &target)
	assert.Equal(t, errTestProtoInstanceAs, target)
}

func TestProtoInstance_As_NilByLink(t *testing.T) {
	t.Parallel()

	var target any

	require.ErrorAs(t, errTestProtoInstanceAs, &target)
	assert.Equal(t, errTestProtoInstanceAs, target)
}
