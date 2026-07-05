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

	type testCase struct {
		name    string
		message string
		want    string
	}

	tests := []testCase{
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

	type testCase struct {
		name       string
		err        error
		want       bool
		mirrorWant bool
	}

	tests := []testCase{
		{
			name:       "test1",
			err:        errors.New("my-message"),
			want:       false,
			mirrorWant: false,
		},
		{
			name:       "test2",
			err:        errUserTest1,
			want:       true,
			mirrorWant: true,
		},
		{
			name:       "test3",
			err:        errUserTest2,
			want:       false,
			mirrorWant: false,
		},
		{
			name:       "test4",
			err:        errUserTest1.Wrap(errors.New("test-error3")),
			want:       true,
			mirrorWant: true,
		},
		{
			name:       "test5",
			err:        errUserTest1.Wrap(errUserTest1),
			want:       true,
			mirrorWant: true,
		},
		{
			name:       "test6",
			err:        errUserTest1.Wrap(errUserTest2),
			want:       true,
			mirrorWant: true,
		},
		{
			name:       "test7",
			err:        errUserTest2.Wrap(errors.New("test-error3")),
			want:       false,
			mirrorWant: false,
		},
		{
			name:       "test8",
			err:        errUserTest2.Wrap(errUserTest1),
			want:       true,
			mirrorWant: false,
		},
		{
			name:       "test9",
			err:        errUserTest2.Wrap(errUserTest2),
			want:       false,
			mirrorWant: false,
		},
		{
			name:       "test10",
			err:        errUserTest2.Wrap(errUserTest1.Wrap(errUserTest1)),
			want:       true,
			mirrorWant: false,
		},
		{
			name:       "test11",
			err:        errUserTest2.Wrap(errUserTest2.Wrap(errUserTest1)),
			want:       true,
			mirrorWant: false,
		},
		{
			name:       "test12",
			err:        errUserTest2.Wrap(errUserTest2.Wrap(errUserTest2)),
			want:       false,
			mirrorWant: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := errors.Is(tt.err, errUserTest1)
			assert.Equal(t, tt.want, got)

			got = errors.Is(errUserTest1, tt.err)
			assert.Equal(t, tt.mirrorWant, got)
		})
	}
}

func TestProto_Wrap_Nil(t *testing.T) {
	t.Parallel()

	proto := userfast.New("code", "message")
	err := proto.Wrap(nil)

	// При передаче nil Wrap возвращает сам прототип, а не обёртку
	assert.Same(t, proto, err)
}

func TestProto_Wrap_Error(t *testing.T) {
	t.Parallel()

	proto := userfast.New("code", "message")
	baseErr := errors.New("base error")
	wrapped := proto.Wrap(baseErr)

	// Wrap возвращает ошибку, содержащую сообщение прототипа
	require.ErrorContains(t, wrapped, "#code - message")
	require.ErrorContains(t, wrapped, "base error")

	// errors.Is находит прототип в обёртке
	require.ErrorIs(t, wrapped, proto)

	// errors.As извлекает ProtoError из обёртки
	var target userfast.ProtoError

	require.ErrorAs(t, wrapped, &target)
	assert.Equal(t, "code", target.Code())
}

func TestProto_As_FromWrapped(t *testing.T) {
	t.Parallel()

	proto := userfast.New("code", "message")
	wrapped := proto.Wrap(errors.New("inner error"))

	// errors.As извлекает ProtoError из обёртки.
	// Поскольку wrapError тоже реализует ProtoError,
	// errors.As возвращает саму обёртку, а не оригинальный protoError.
	var target userfast.ProtoError

	require.ErrorAs(t, wrapped, &target)
	assert.Equal(t, "code", target.Code())
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
