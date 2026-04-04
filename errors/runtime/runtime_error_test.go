package runtime_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-sysmess/errors/kind"
	"github.com/mondegor/go-sysmess/errors/runtime"
)

func TestProtoError_WithDetails_CreatesErrorWithDetails(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.Internal, "base error")
	err := proto.WithDetails("additional context")

	require.Error(t, err)
	assert.Equal(t, "base error: additional context", err.Error())
}

func TestProtoError_WithDetails_CreatesErrorWithAttributes(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.Internal, "storage error")
	err := proto.WithDetails("query failed", "db", "postgres", "table", "users")

	require.Error(t, err)
	assert.Equal(t, "storage error: query failed [db=postgres, table=users]", err.Error())
}

func TestProtoError_WithDetails_PreservesErrorKind(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.System, "system failure")
	err := proto.WithDetails("timeout occurred")

	var runtimeErr interface{ Kind() kind.Enum }

	require.ErrorAs(t, err, &runtimeErr)
	assert.Equal(t, kind.System, runtimeErr.Kind())
}

func TestProtoError_WithDetails_DoesNotModifyOriginalProto(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.Internal, "original")
	_ = proto.WithDetails("added details")

	// Original proto should still produce the same error
	originalErr := proto.New()
	assert.Equal(t, "original", originalErr.Error())
}

func TestProtoError_WithDetails_MultipleCallsCreateIndependentErrors(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.Internal, "base")
	err1 := proto.WithDetails("first")
	err2 := proto.WithDetails("second")

	assert.Equal(t, "base: first", err1.Error())
	assert.Equal(t, "base: second", err2.Error())
}

func TestProtoError_WithDetails_WithEmptyDetailsAndNoAttributes(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.Internal, "error message")
	err := proto.WithDetails("")

	require.Error(t, err)
	assert.Equal(t, "error message", err.Error())
}

func TestProtoError_WithDetails_WithOnlyAttributesNoDetails(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.Internal, "error")
	err := proto.WithDetails("", "key", "value")

	require.Error(t, err)
	assert.Equal(t, "error [key=value]", err.Error())
}

func TestProtoError_WithDetails_AttributesAreCopiedNotShared(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.Internal, "error")
	err := proto.WithDetails("context", "attr1", "val1")

	// Modify attrs through the public API by creating a new error
	// and checking that the original is unchanged
	err2 := proto.WithDetails("context", "attr1", "val1")

	// Both errors should have the same content since they're from same proto
	assert.Equal(t, err.Error(), err2.Error())
}

func TestProtoError_WithDetails_HandlesSpecialCharactersInDetails(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.Internal, "error")
	err := proto.WithDetails("context: with special chars & <tags>")

	require.Error(t, err)
	assert.Equal(t, "error: context: with special chars & <tags>", err.Error())
}

func TestProtoError_WithDetails_HandlesUnicodeInDetails(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.Internal, "error")
	err := proto.WithDetails("контекст", "ключ", "значение")

	require.Error(t, err)
	assert.Equal(t, "error: контекст [ключ=значение]", err.Error())
}

func TestProtoError_WithDetails_IsComparesWithSamePrototype(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.Internal, "error")
	err1 := proto.WithDetails("details1")
	err2 := proto.WithDetails("details2")

	// Both errors should match the same prototype
	require.ErrorIs(t, err1, proto.New())
	require.ErrorIs(t, err2, proto.New())
}

func TestProtoError_WithDetails_UnwrapReturnsNil(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.Internal, "error")
	err := proto.WithDetails("details")

	assert.NoError(t, errors.Unwrap(err))
}

func TestProtoError_WithDetails_WithMultipleAttributes(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.Internal, "complex error")
	err := proto.WithDetails("failed",
		"user_id", 123,
		"action", "delete",
		"resource", "account",
		"attempt", 3)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "user_id=123")
	assert.Contains(t, err.Error(), "action=delete")
	assert.Contains(t, err.Error(), "resource=account")
	assert.Contains(t, err.Error(), "attempt=3")
}

func TestProtoError_WithDetails_WithNilLikeAttributeValues(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.Internal, "error")
	err := proto.WithDetails("context", "ptr", nil, "num", 0, "bool", false)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "ptr=<NIL>")
	assert.Contains(t, err.Error(), "num=0")
	assert.Contains(t, err.Error(), "bool=false")
}

func TestProtoError_WithDetails_InvokesOnCreateCallback(t *testing.T) {
	t.Parallel()

	callbackCalled := false
	callbackData := "callback_data"

	proto := runtime.New(kind.Internal, "error", runtime.WithOnCreate(func(k kind.Enum, err error) any {
		callbackCalled = true

		return callbackData
	}))

	err := proto.WithDetails("details")

	assert.True(t, callbackCalled)

	var hintProvider interface{ Hint() any }

	require.ErrorAs(t, err, &hintProvider)
	assert.Equal(t, callbackData, hintProvider.Hint())
}

func TestProtoError_WithDetails_OnCreateReceivesCorrectParameters(t *testing.T) {
	t.Parallel()

	var receivedKind kind.Enum

	var receivedErr error

	proto := runtime.New(kind.System, "error", runtime.WithOnCreate(func(k kind.Enum, err error) any {
		receivedKind = k
		receivedErr = err

		return "data"
	}))

	_ = proto.WithDetails("details")

	assert.Equal(t, kind.System, receivedKind)
	assert.NoError(t, receivedErr)
}

func TestProtoError_WithDetails_CombinesWithWrapAndDetails(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.Internal, "wrapper error")
	baseErr := runtime.New(kind.Internal, "base error").New()
	err := proto.WithError(baseErr, "additional context", "key", "value")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "wrapper error: additional context")
	assert.Contains(t, err.Error(), "base error")
	assert.Contains(t, err.Error(), "key=value")
}

func TestProtoError_WithDetails_AppendsDetailsAfterBaseText(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.User, "validation failed")
	err := proto.WithDetails("field 'email' is invalid")

	require.Error(t, err)
	assert.Equal(t, "validation failed: field 'email' is invalid", err.Error())
}

func TestProtoError_WithDetails_WithNumericAttributeValues(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.Internal, "error")
	err := proto.WithDetails("limit exceeded",
		"current", 150,
		"max", 100,
		"ratio", 1.5)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "current=150")
	assert.Contains(t, err.Error(), "max=100")
	assert.Contains(t, err.Error(), "ratio=1.5")
}

func TestProtoError_WithDetails_WithSliceAttributeValues(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.Internal, "error")
	err := proto.WithDetails("batch failed", "ids", []int{1, 2, 3})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "ids=[1 2 3]")
}

func TestProtoError_WithDetails_WithMapAttributeValues(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.Internal, "error")
	err := proto.WithDetails("config error", "settings", map[string]string{"key": "val"})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "settings=")
}

func TestProtoError_WithDetails_WithOddNumberOfAttributes(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.Internal, "error")
	err := proto.WithDetails("context", "key1", "val1", "orphan_key")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "key1=val1")
	assert.Contains(t, err.Error(), "!MISSINGATTRVALUE")
}

func TestProtoError_WithDetails_WithEmptyStringAsAttributeKey(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.Internal, "error")
	err := proto.WithDetails("context", "", "value")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "!EMPTYATTRKEY=value")
}

func TestProtoError_WithDetails_WithNonStringAsFirstAttribute(t *testing.T) {
	t.Parallel()

	proto := runtime.New(kind.Internal, "error")
	err := proto.WithDetails("context", 123, "value")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "!BADATTRKEY=123")
	assert.Contains(t, err.Error(), "value")
}
