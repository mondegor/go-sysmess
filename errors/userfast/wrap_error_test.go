package userfast_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-core/errors/kind"
	"github.com/mondegor/go-core/errors/userfast"
)

func TestWrapError_Is(t *testing.T) {
	t.Parallel()

	errUserTestWrapper := userfast.New("test-code1", "test-message1")
	errUserTest1 := errUserTestWrapper.Wrap(errors.New("test-message2"))
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
			err:        errUserTestWrapper,
			want:       true,
			mirrorWant: true,
		},
		{
			name:       "test3",
			err:        errUserTest1,
			want:       true,
			mirrorWant: true,
		},
		{
			name:       "test4",
			err:        errUserTest2,
			want:       false,
			mirrorWant: false,
		},
		{
			name:       "test5",
			err:        errUserTestWrapper.Wrap(errors.New("test-error3")),
			want:       true,
			mirrorWant: true,
		},
		{
			name:       "test6",
			err:        errUserTestWrapper.Wrap(errUserTest1),
			want:       true,
			mirrorWant: true,
		},
		{
			name:       "test7",
			err:        errUserTestWrapper.Wrap(errUserTest2),
			want:       true,
			mirrorWant: true,
		},
		{
			name:       "test8",
			err:        errUserTest2.Wrap(errors.New("test-error3")),
			want:       false,
			mirrorWant: false,
		},
		{
			name:       "test9",
			err:        errUserTest2.Wrap(errUserTest1),
			want:       true,
			mirrorWant: false,
		},
		{
			name:       "test10",
			err:        errUserTest2.Wrap(errUserTest2),
			want:       false,
			mirrorWant: false,
		},
		{
			name:       "test11",
			err:        errUserTest2.Wrap(errUserTestWrapper.Wrap(errUserTest1)),
			want:       true,
			mirrorWant: false,
		},
		{
			name:       "test12",
			err:        errUserTest2.Wrap(errUserTest2.Wrap(errUserTest1)),
			want:       true,
			mirrorWant: false,
		},
		{
			name:       "test13",
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

func TestWrapError_Kind(t *testing.T) {
	t.Parallel()

	proto := userfast.New("code", "message")
	wrapped := proto.Wrap(errors.New("inner error"))

	kindGetter, ok := wrapped.(interface{ Kind() kind.Enum })
	require.True(t, ok)
	assert.Equal(t, kind.User, kindGetter.Kind())
}

func TestWrapError_Code(t *testing.T) {
	t.Parallel()

	proto := userfast.New("my-code", "my message")
	wrapped := proto.Wrap(errors.New("inner error"))

	codeGetter, ok := wrapped.(interface{ Code() string })
	require.True(t, ok)
	assert.Equal(t, "my-code", codeGetter.Code())
}

func TestWrapError_Message(t *testing.T) {
	t.Parallel()

	proto := userfast.New("code", "original message")
	wrapped := proto.Wrap(errors.New("inner error"))

	msgGetter, ok := wrapped.(interface{ Message() string })
	require.True(t, ok)
	assert.Equal(t, "original message", msgGetter.Message())
}

func TestWrapError_Args(t *testing.T) {
	t.Parallel()

	proto := userfast.New("code", "message")
	wrapped := proto.Wrap(errors.New("inner error"))

	argsGetter, ok := wrapped.(interface{ Args() []any })
	require.True(t, ok)
	assert.Nil(t, argsGetter.Args())
}

func TestWrapError_Error(t *testing.T) {
	t.Parallel()

	proto := userfast.New("ERR_TIMEOUT", "request timed out")
	baseErr := errors.New("connection refused")
	wrapped := proto.Wrap(baseErr)

	// Формат: "#CODE - message: inner_error"
	assert.Equal(t, "#ERR_TIMEOUT - request timed out: connection refused", wrapped.Error())
}

func TestWrapError_Error_NestedProtos(t *testing.T) {
	t.Parallel()

	// Проверяем, что сообщение берётся от прототипа обёртки,
	// а не от вложенной ошибки
	innerProto := userfast.New("INNER", "inner message")
	outerProto := userfast.New("OUTER", "outer message")

	wrapped := outerProto.Wrap(innerProto)
	assert.Equal(t, "#OUTER - outer message: #INNER - inner message", wrapped.Error())
}

func TestWrapError_Unwrap(t *testing.T) {
	t.Parallel()

	proto := userfast.New("code", "message")
	baseErr := errors.New("inner error")
	wrapped := proto.Wrap(baseErr)

	unwrapper, ok := wrapped.(interface{ Unwrap() error })
	require.True(t, ok)
	assert.Equal(t, baseErr, unwrapper.Unwrap())
}

func TestWrapError_Unwrap_NestedWrap(t *testing.T) {
	t.Parallel()

	innerProto := userfast.New("INNER", "inner")
	outerProto := userfast.New("OUTER", "outer")
	innerWrapped := innerProto.Wrap(errors.New("base"))
	outerWrapped := outerProto.Wrap(innerWrapped)

	// Unwrap должен вернуть innerWrapped
	assert.Equal(t, innerWrapped, outerWrapped.(interface{ Unwrap() error }).Unwrap())
}

func TestWrapError_ErrorsUnwrap(t *testing.T) {
	t.Parallel()

	proto := userfast.New("code", "message")
	baseErr := errors.New("base error")
	wrapped := proto.Wrap(baseErr)

	// errors.Unwrap должен работать через интерфейс Unwraper
	assert.Equal(t, baseErr, errors.Unwrap(wrapped))
}

func TestWrapError_ErrorsAs_ProtoError(t *testing.T) {
	t.Parallel()

	proto := userfast.New("code", "message")
	wrapped := proto.Wrap(errors.New("inner error"))

	// errors.As возвращает первый ProtoError в цепочке.
	// Поскольку wrapError реализует ProtoError, errors.As
	// возвращает саму обёртку, а не оригинальный protoError.
	var target userfast.ProtoError

	require.ErrorAs(t, wrapped, &target)
	assert.Equal(t, "code", target.Code())
}

func TestWrapError_DeepChain_ErrorsIs(t *testing.T) {
	t.Parallel()

	// Цепочка: proto -> wrap(proto) -> wrap(wrap) -> wrap(wrap(wrap))
	proto := userfast.New("code", "message")
	err1 := errors.New("err1")
	wrap1 := proto.Wrap(err1)
	wrap2 := proto.Wrap(wrap1)
	wrap3 := proto.Wrap(wrap2)

	// Все уровни должны находить прототип
	require.ErrorIs(t, wrap1, proto)
	require.ErrorIs(t, wrap2, proto)
	require.ErrorIs(t, wrap3, proto)

	// И зеркальная проверка
	require.ErrorIs(t, proto, wrap1)
	require.ErrorIs(t, proto, wrap2)
	assert.ErrorIs(t, proto, wrap3)
}

func TestWrapError_DeepChain_ErrorsAs(t *testing.T) {
	t.Parallel()

	// Глубокая цепочка: proto2.wrap(proto1.wrap(proto1.wrap(err)))
	proto1 := userfast.New("code1", "message1")
	proto2 := userfast.New("code2", "message2")
	err := errors.New("base")
	wrap1 := proto1.Wrap(err)
	wrap2 := proto1.Wrap(wrap1)
	wrap3 := proto2.Wrap(wrap2)

	// errors.As находит первый ProtoError в цепочке - это wrap3 (proto2 обёртка).
	// wrapError теперь реализует ProtoError.
	var target userfast.ProtoError

	require.ErrorAs(t, wrap3, &target)
	assert.Equal(t, "code2", target.Code()) // code2, т.к. wrap3 - это обёртка proto2

	// errors.Is(wrap3, proto1) - wrap3 содержит proto1 внутри через цепочку
	require.ErrorIs(t, wrap3, proto1)

	// errors.Is(wrap3, proto2) - wrap3 сам является proto2 обёрткой
	assert.ErrorIs(t, wrap3, proto2)
}

func TestWrapError_MixedStdErrors(t *testing.T) {
	t.Parallel()

	proto := userfast.New("code", "message")
	stdErr1 := errors.New("std error 1")
	stdErr2 := errors.New("std error 2")

	wrap1 := proto.Wrap(stdErr1)
	wrap2 := proto.Wrap(wrap1)

	// errors.Is должен найти stdErr1 через цепочку
	require.ErrorIs(t, wrap2, stdErr1)
	assert.NotErrorIs(t, wrap2, stdErr2)
}

func TestWrapError_InterfaceContract(t *testing.T) {
	t.Parallel()

	proto := userfast.New("code", "message")
	wrapped := proto.Wrap(errors.New("inner"))

	// Проверяем, что wrapError реализует все ожидаемые интерфейсы
	type fullError interface {
		error
		Kind() kind.Enum
		Code() string
		Message() string
		Args() []any
	}

	var iface fullError

	require.ErrorAs(t, wrapped, &iface)
	assert.Equal(t, kind.User, iface.Kind())
	assert.Equal(t, "code", iface.Code())
	assert.Equal(t, "message", iface.Message())
	assert.Nil(t, iface.Args())
	assert.Contains(t, iface.Error(), "#code - message")
}
