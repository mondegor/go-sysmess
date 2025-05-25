package mrerrors_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-sysmess/mrerrors"
)

func TestNewProto(t *testing.T) {
	t.Parallel()

	got := mrerrors.NewProto("test-message", mrerrors.WithProtoKind(mrerrors.ErrorKindSystem), mrerrors.WithProtoCode("test-code"))
	require.Equal(t, "test-code", got.Code())
	require.Equal(t, mrerrors.ErrorKindSystem, got.Kind())
	assert.ErrorContains(t, got, "test-message")
}

func TestProtoInstantError_Error(t *testing.T) {
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
			want:    "my-message [INTERNAL]",
		},
		{
			name:    "test2",
			message: "my-message {Key1} - {Key2}",
			want:    "my-message {Key1} - {Key2} [INTERNAL]",
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := mrerrors.NewProto(tt.message)
			got := e.Error()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestProtoInstantError_Is(t *testing.T) {
	t.Parallel()

	errInternalTest1 := mrerrors.NewProto("test-message1")
	errInternalTest2 := mrerrors.NewProto("test-message2")

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
			target: errInternalTest1,
			want:   true,
		},
		{
			name:   "test3",
			target: errInternalTest2,
			want:   false,
		},
		{
			name:   "test4",
			target: errInternalTest1.New(),
			want:   true,
		},
		{
			name:   "test5",
			target: errInternalTest2.New(),
			want:   false,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := errInternalTest1.Is(tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}

var errTestProtoAppAs = mrerrors.NewProto("test-message1")

func TestProtoInstantError_As_PointerProtoByLink(t *testing.T) {
	t.Parallel()

	var target *mrerrors.ProtoError

	require.True(t, errTestProtoAppAs.As(&target))
	assert.Equal(t, errTestProtoAppAs, target)
}

func TestProtoInstantError_As_AnyPointerProtoByLink(t *testing.T) {
	t.Parallel()

	var target any = (*mrerrors.ProtoError)(nil)

	require.True(t, errTestProtoAppAs.As(&target))
	assert.Equal(t, errTestProtoAppAs, target)
}

func TestProtoInstantError_As_NilByLink(t *testing.T) {
	t.Parallel()

	var target any

	assert.False(t, errTestProtoAppAs.As(&target))
}

func TestProtoInstantError_As_NilByValue(t *testing.T) {
	t.Parallel()

	var target any

	require.Panics(t, func() { errTestProtoAppAs.As(target) })
	assert.Panics(t, func() { errTestProtoAppAs.As(nil) })
}

func TestProtoInstantError_As_PointerProtoByValue(t *testing.T) {
	t.Parallel()

	var target *mrerrors.ProtoError

	require.Panics(t, func() { errTestProtoAppAs.As(target) })
	assert.Panics(t, func() { errTestProtoAppAs.As((*mrerrors.ProtoError)(nil)) })
}

func TestProtoInstantError_As_2PointerProtoByValue(t *testing.T) {
	t.Parallel()

	var target **mrerrors.ProtoError

	require.Panics(t, func() { errTestProtoAppAs.As(target) })
	assert.Panics(t, func() { errTestProtoAppAs.As((**mrerrors.ProtoError)(nil)) })
}
