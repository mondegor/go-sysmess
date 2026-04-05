package user_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-sysmess/errors/kind"
	"github.com/mondegor/go-sysmess/errors/user"
)

func TestNew(t *testing.T) {
	t.Parallel()

	err := user.New("test-code", "test-message")
	got, ok := err.(interface {
		error
		Kind() kind.Enum
	})

	require.True(t, ok)
	require.Equal(t, kind.User, got.Kind())
	assert.ErrorContains(t, got, "#test-code - test-message")
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

			e := user.New("test-code", tt.message)
			assert.Equal(t, tt.want, e.Error())
		})
	}
}

func TestProto_Is(t *testing.T) {
	t.Parallel()

	errUserTest1 := user.New("test-code1", "test-message1")
	errUserTest2 := user.New("test-code2", "test-message2")

	tests := []struct {
		name       string
		err        error
		want       bool
		mirrorWant bool
	}{
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
			err:        errUserTest1.New(),
			want:       true,
			mirrorWant: true,
		},
		{
			name:       "test5",
			err:        errUserTest1.New("arg1"),
			want:       true,
			mirrorWant: true,
		},
		{
			name:       "test6",
			err:        errUserTest2.New(),
			want:       false,
			mirrorWant: false,
		},
		{
			name:       "test7",
			err:        errUserTest1.Wrap(errors.New("my-message")),
			want:       true,
			mirrorWant: true,
		},
		{
			name:       "test8",
			err:        errUserTest2.Wrap(errUserTest1),
			want:       true,
			mirrorWant: false,
		},
		{
			name:       "test9",
			err:        errUserTest2.Wrap(errUserTest1.New()),
			want:       true,
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

	proto := user.New("code", "message")
	wrapped := proto.Wrap(nil)

	// В user.Wrap(nil) оборачивает ErrHasNilError вместо возврата прототипа
	require.Error(t, wrapped)
	assert.ErrorIs(t, wrapped, user.ErrHasNilError)

	// Ошибка содержит сообщение прототипа
	assert.ErrorContains(t, wrapped, "#code - message")
}

func TestProto_Wrap_Error(t *testing.T) {
	t.Parallel()

	proto := user.New("code", "message")
	baseErr := errors.New("base error")
	wrapped := proto.Wrap(baseErr)

	// Wrap возвращает ошибку, содержащую сообщение прототипа
	assert.ErrorContains(t, wrapped, "#code - message")
	assert.ErrorContains(t, wrapped, "base error")

	// errors.Is находит прототип в обёртке
	assert.True(t, errors.Is(wrapped, proto))

	// errors.As извлекает ProtoError из обёртки.
	// Возвращается *protoError с тем же id, но установленным err.
	var target user.ProtoError
	require.ErrorAs(t, wrapped, &target)
	// Сравнение через proto.Is — совпадает по id
	assert.True(t, errors.Is(wrapped, target))
	assert.True(t, errors.Is(target, proto))
	assert.Equal(t, "code", target.Code())
}

func TestProto_New_Args(t *testing.T) {
	t.Parallel()

	proto := user.New("VALIDATION", "field '{0}' is invalid")
	err := proto.New("email")

	// New подставляет аргументы в сообщение
	assert.ErrorContains(t, err, "#VALIDATION")
	assert.ErrorContains(t, err, "field")

	// Args возвращает аргументы для локализации
	argsGetter, ok := err.(interface{ Args() []any })
	require.True(t, ok)
	assert.Equal(t, []any{"email"}, argsGetter.Args())
}

func TestProto_New_NoArgs(t *testing.T) {
	t.Parallel()

	proto := user.New("SIMPLE", "simple message")
	err := proto.New()

	// New() без аргументов возвращает сам прототип (оптимизация)
	assert.Same(t, proto, err)
}

func TestProto_Unwrap(t *testing.T) {
	t.Parallel()

	proto := user.New("code", "message")
	baseErr := errors.New("base error")
	wrapped := proto.Wrap(baseErr)

	// errors.Unwrap должен возвращать вложенную ошибку
	assert.Equal(t, baseErr, errors.Unwrap(wrapped))
}

func TestProto_Unwrap_Chain(t *testing.T) {
	t.Parallel()

	innerProto := user.New("INNER", "inner")
	outerProto := user.New("OUTER", "outer")
	baseErr := errors.New("base")
	innerErr := innerProto.Wrap(baseErr)
	outerErr := outerProto.Wrap(innerErr)

	// Unwrap должен вернуть innerErr
	assert.Equal(t, innerErr, errors.Unwrap(outerErr))

	// errors.Is должен найти base error через цепочку (сравниваем тот же объект)
	assert.True(t, errors.Is(outerErr, baseErr))
	assert.True(t, errors.Is(outerErr, innerProto))
}

func TestProto_As_FromWrapped(t *testing.T) {
	t.Parallel()

	proto := user.New("code", "message")
	wrapped := proto.Wrap(errors.New("inner error"))

	// errors.As извлекает ProtoError из обёртки.
	// В user пакете Wrap создаёт новый *protoError,
	// который имеет тот же id и сравнивается через Is().
	var target user.ProtoError
	require.ErrorAs(t, wrapped, &target)
	assert.True(t, errors.Is(target, proto))
	assert.Equal(t, "code", target.Code())
}

func TestProto_As_FromWrappedWithArgs(t *testing.T) {
	t.Parallel()

	proto := user.New("VALIDATION", "field '{0}' is required")
	wrapped := proto.Wrap(errors.New("db error"), "email")

	var target user.ProtoError
	require.ErrorAs(t, wrapped, &target)
	assert.True(t, errors.Is(target, proto))
	assert.Equal(t, "VALIDATION", target.Code())

	argsGetter, ok := wrapped.(interface{ Args() []any })
	require.True(t, ok)
	assert.Equal(t, []any{"email"}, argsGetter.Args())
}

func TestProto_Kind(t *testing.T) {
	t.Parallel()

	proto := user.New("code", "message")
	kindGetter, ok := proto.(interface{ Kind() kind.Enum })
	require.True(t, ok)
	assert.Equal(t, kind.User, kindGetter.Kind())

	wrapped := proto.Wrap(errors.New("inner"))
	kindGetter, ok = wrapped.(interface{ Kind() kind.Enum })
	require.True(t, ok)
	assert.Equal(t, kind.User, kindGetter.Kind())
}

func TestProto_ErrorsIs_DeepChain(t *testing.T) {
	t.Parallel()

	// Цепочка: proto -> wrap(proto) -> wrap(wrap)
	proto := user.New("code", "message")
	err1 := errors.New("err1")
	wrap1 := proto.Wrap(err1)
	wrap2 := proto.Wrap(wrap1)

	assert.True(t, errors.Is(wrap1, proto))
	assert.True(t, errors.Is(wrap2, proto))
	assert.True(t, errors.Is(wrap1, err1))
	assert.True(t, errors.Is(wrap2, err1))
}

func TestProto_ErrorsIs_CrossProto(t *testing.T) {
	t.Parallel()

	proto1 := user.New("code1", "message1")
	proto2 := user.New("code2", "message2")
	wrapped := proto2.Wrap(proto1)

	// wrapped (proto2) содержит proto1 внутри через Unwrap
	assert.True(t, errors.Is(wrapped, proto1))
	assert.True(t, errors.Is(wrapped, proto2))

	// proto1 не содержит wrapped (разные id: proto1 != proto2)
	assert.False(t, errors.Is(proto1, wrapped))

	// proto2 содержит wrapped, т.к. wrapped имеет тот же id, что и proto2
	// (protoError.Is сравнивает по id, а не по указателю)
	assert.True(t, errors.Is(proto2, wrapped))
}

var errTestProtoInstanceAs = user.New("test-code1", "test-message1")

func TestProtoInstance_As_PointerProtoByLink(t *testing.T) {
	t.Parallel()

	var target user.ProtoError

	require.ErrorAs(t, errTestProtoInstanceAs, &target)
	assert.Equal(t, errTestProtoInstanceAs, target)
}

func TestProtoInstance_As_AnyPointerProtoByLink(t *testing.T) {
	t.Parallel()

	var target any = (user.ProtoError)(nil)

	require.ErrorAs(t, errTestProtoInstanceAs, &target)
	assert.Equal(t, errTestProtoInstanceAs, target)
}

func TestProtoInstance_As_NilByLink(t *testing.T) {
	t.Parallel()

	var target any

	require.ErrorAs(t, errTestProtoInstanceAs, &target)
	assert.Equal(t, errTestProtoInstanceAs, target)
}
