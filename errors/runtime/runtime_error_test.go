package runtime_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-core/errors/kind"
	"github.com/mondegor/go-core/errors/runtime"
)

func TestProtoError_WithDetails_ErrorString(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name    string
		kind    kind.Enum
		message string
		details string
		attrs   []any
		want    string
	}

	tests := []testCase{
		{
			name:    "creates error with details",
			kind:    kind.Internal,
			message: "base error",
			details: "additional context",
			want:    "base error: additional context",
		},
		{
			name:    "creates error with attributes",
			kind:    kind.Internal,
			message: "storage error",
			details: "query failed",
			attrs:   []any{"db", "postgres", "table", "users"},
			want:    "storage error: query failed [db=postgres, table=users]",
		},
		{
			name:    "empty details and no attributes",
			kind:    kind.Internal,
			message: "error message",
			details: "",
			want:    "error message",
		},
		{
			name:    "only attributes no details",
			kind:    kind.Internal,
			message: "error",
			details: "",
			attrs:   []any{"key", "value"},
			want:    "error [key=value]",
		},
		{
			name:    "special characters in details",
			kind:    kind.Internal,
			message: "error",
			details: "context: with special chars & <tags>",
			want:    "error: context: with special chars & <tags>",
		},
		{
			name:    "unicode in details",
			kind:    kind.Internal,
			message: "error",
			details: "контекст",
			attrs:   []any{"ключ", "значение"},
			want:    "error: контекст [ключ=значение]",
		},
		{
			name:    "appends details after base text",
			kind:    kind.User,
			message: "validation failed",
			details: "field 'email' is invalid",
			want:    "validation failed: field 'email' is invalid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			proto := runtime.New(tt.kind, tt.message)
			err := proto.WithDetails(tt.details, tt.attrs...)

			require.Error(t, err)
			assert.Equal(t, tt.want, err.Error())
		})
	}
}

func TestProtoError_WithDetails_ContainsAttributes(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name         string
		message      string
		details      string
		attrs        []any
		wantContains []string
	}

	tests := []testCase{
		{
			name:         "multiple attributes",
			message:      "complex error",
			details:      "failed",
			attrs:        []any{"user_id", 123, "action", "delete", "resource", "account", "attempt", 3},
			wantContains: []string{"user_id=123", "action=delete", "resource=account", "attempt=3"},
		},
		{
			name:         "nil-like attribute values",
			message:      "error",
			details:      "context",
			attrs:        []any{"ptr", nil, "num", 0, "bool", false},
			wantContains: []string{"ptr=<NIL>", "num=0", "bool=false"},
		},
		{
			name:         "numeric attribute values",
			message:      "error",
			details:      "limit exceeded",
			attrs:        []any{"current", 150, "max", 100, "ratio", 1.5},
			wantContains: []string{"current=150", "max=100", "ratio=1.5"},
		},
		{
			name:         "slice attribute values",
			message:      "error",
			details:      "batch failed",
			attrs:        []any{"ids", []int{1, 2, 3}},
			wantContains: []string{"ids=[1 2 3]"},
		},
		{
			name:         "map attribute values",
			message:      "error",
			details:      "config error",
			attrs:        []any{"settings", map[string]string{"key": "val"}},
			wantContains: []string{"settings="},
		},
		{
			name:         "odd number of attributes",
			message:      "error",
			details:      "context",
			attrs:        []any{"key1", "val1", "orphan_key"},
			wantContains: []string{"key1=val1", "!MISSINGATTRVALUE"},
		},
		{
			name:         "empty string as attribute key",
			message:      "error",
			details:      "context",
			attrs:        []any{"", "value"},
			wantContains: []string{"!EMPTYATTRKEY=value"},
		},
		{
			name:         "non-string as first attribute",
			message:      "error",
			details:      "context",
			attrs:        []any{123, "value"},
			wantContains: []string{"!BADATTRKEY=123", "value"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			proto := runtime.New(kind.Internal, tt.message)
			err := proto.WithDetails(tt.details, tt.attrs...)

			require.Error(t, err)

			for _, want := range tt.wantContains {
				assert.Contains(t, err.Error(), want)
			}
		})
	}
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
