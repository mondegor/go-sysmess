package mrerr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testGeneratedErrorID = "test-instance-id"

type mockStackTrace struct{}

func (m *mockStackTrace) Count() int                                 { return 0 }
func (m *mockStackTrace) FileLine(index int) (file string, line int) { return "", 0 }

func TestNewProto(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		code    string
		kind    ErrorKind
		message string
		want    *ProtoAppError
	}{
		{
			name:    "test1",
			code:    "test-code1",
			kind:    ErrorKindInternal,
			message: "test-message1",
			want: &ProtoAppError{
				pureError: pureError{
					code:    "test-code1",
					kind:    ErrorKindInternal,
					message: "test-message1",
				},
				generateID: nil,
				caller:     nil,
			},
		},
		{
			name:    "test1",
			code:    "test-code2",
			kind:    ErrorKindSystem,
			message: "test-message {{ .key1 }} and {{ .key2 }}",
			want: &ProtoAppError{
				pureError: pureError{
					code:      "test-code2",
					kind:      ErrorKindSystem,
					message:   "test-message {{ .key1 }} and {{ .key2 }}",
					argsNames: []string{"key1", "key2"},
					args:      []any{"missed-arg1", "missed-arg2"},
				},
				generateID: nil,
				caller:     nil,
			},
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := NewProto(tt.code, tt.kind, tt.message)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewWithExtra(t *testing.T) {
	t.Parallel()

	mockStack := &mockStackTrace{}

	expectedProto := ProtoAppError{
		pureError: pureError{
			code:    "errTestCode",
			kind:    ErrorKindUser,
			message: "test-message",
		},
		generateID: func() string { return testGeneratedErrorID },
		caller:     func() StackTracer { return mockStack },
	}

	got := NewProtoWithExtra(
		expectedProto.code,
		expectedProto.kind,
		expectedProto.message,
		expectedProto.generateID,
		expectedProto.caller,
	)

	assert.Equal(t, expectedProto.code, got.code)
	assert.Equal(t, expectedProto.kind, got.kind)
	assert.Equal(t, expectedProto.message, got.message)
	assert.NotNil(t, got.generateID)
	assert.NotNil(t, got.caller)
}

func TestProtoAppError_New(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		paramArgs  []any
		code       string
		kind       ErrorKind
		generateID func() string
		message    string
		argsNames  []string
		args       []any
		want       *AppError
	}{
		{
			name:      "test1",
			code:      "test-code",
			kind:      ErrorKindSystem,
			message:   "test-message {{ .key1 }} and {{ .key2 }}",
			argsNames: []string{"key1", "key2"},
			args:      []any{"test-arg1", "test-arg2"},
			want: &AppError{
				pureError: pureError{
					code:      "test-code",
					kind:      ErrorKindSystem,
					message:   "test-message {{ .key1 }} and {{ .key2 }}",
					argsNames: []string{"key1", "key2"},
					args:      []any{"test-arg1", "test-arg2"},
				},
			},
		},
		{
			name:      "test2",
			paramArgs: []any{"test-param-arg1"},
			argsNames: []string{"key1", "key2"},
			want: &AppError{
				pureError: pureError{
					argsNames: []string{"key1", "key2"},
					args:      []any{"test-param-arg1", "missed-arg2"},
				},
			},
		},
		{
			name:      "test3",
			paramArgs: []any{"test-param-arg1", "test-param-arg2"},
			argsNames: []string{"key1", "key2"},
			want: &AppError{
				pureError: pureError{
					argsNames: []string{"key1", "key2"},
					args:      []any{"test-param-arg1", "test-param-arg2"},
				},
			},
		},
		{
			name:      "test4",
			paramArgs: []any{"test-param-arg1", "test-param-arg2", "test-param-arg3"},
			argsNames: []string{"key1", "key2"},
			want: &AppError{
				pureError: pureError{
					argsNames: []string{"key1", "key2"},
					args:      []any{"test-param-arg1", "test-param-arg2", "test-param-arg3"},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := &ProtoAppError{
				pureError: pureError{
					code:      tt.code,
					kind:      tt.kind,
					message:   tt.message,
					argsNames: tt.argsNames,
					args:      tt.args,
				},
				generateID: tt.generateID,
			}
			got := e.New(tt.paramArgs...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestProtoAppError_NewWithGenerateID(t *testing.T) {
	t.Parallel()

	proto := &ProtoAppError{
		generateID: func() string { return testGeneratedErrorID },
	}

	got := proto.New()
	assert.Equal(t, testGeneratedErrorID, got.instanceID)
}

func TestProtoAppError_NewWithStackTrace(t *testing.T) {
	t.Parallel()

	mockStack := &mockStackTrace{}

	proto := &ProtoAppError{
		caller: func() StackTracer {
			return mockStack
		},
	}

	got := proto.New()
	assert.True(t, got.stackTrace.has)
	assert.Equal(t, mockStack, got.stackTrace.val)
}

func TestProtoAppError_Wrap(t *testing.T) {
	t.Parallel()

	testErr := errors.New("test-error")

	tests := []struct {
		name       string
		paramErr   error
		paramArgs  []any
		code       string
		kind       ErrorKind
		generateID func() string
		message    string
		argsNames  []string
		args       []any
		want       *AppError
	}{
		{
			name:      "test1",
			code:      "test-code",
			kind:      ErrorKindSystem,
			message:   "test-message {{ .key1 }} and {{ .key2 }}",
			argsNames: []string{"key1", "key2"},
			args:      []any{"test-arg1", "test-arg2"},
			want: &AppError{
				pureError: pureError{
					code:      "test-code",
					kind:      ErrorKindSystem,
					message:   "test-message {{ .key1 }} and {{ .key2 }}",
					argsNames: []string{"key1", "key2"},
					args:      []any{"test-arg1", "test-arg2"},
				},
				err: errSpecifiedErrorIsNil,
			},
		},
		{
			name:      "test2",
			paramErr:  testErr,
			paramArgs: []any{"test-param-arg1"},
			argsNames: []string{"key1", "key2"},
			want: &AppError{
				pureError: pureError{
					argsNames: []string{"key1", "key2"},
					args:      []any{"test-param-arg1", "missed-arg2"},
				},
				err: testErr,
			},
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := &ProtoAppError{
				pureError: pureError{
					code:      tt.code,
					kind:      tt.kind,
					message:   tt.message,
					argsNames: tt.argsNames,
					args:      tt.args,
				},
				generateID: tt.generateID,
			}
			got := e.Wrap(tt.paramErr, tt.paramArgs...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestProtoAppError_WrapWithGenerateID(t *testing.T) {
	t.Parallel()

	proto := &ProtoAppError{
		generateID: func() string { return testGeneratedErrorID },
	}

	wrappedErr := &AppError{}

	got := proto.Wrap(wrappedErr)
	assert.Equal(t, testGeneratedErrorID, got.instanceID)
}

func TestProtoAppError_WrapWrappedErrWithInstanceID(t *testing.T) {
	t.Parallel()

	proto := &ProtoAppError{
		generateID: func() string { return testGeneratedErrorID },
	}

	wrappedErr := &AppError{
		instanceID: testGeneratedErrorID,
	}

	got := proto.Wrap(wrappedErr)
	assert.Empty(t, got.instanceID)
	assert.Equal(t, wrappedErr.instanceID, *got.errInstanceID)
}

func TestProtoAppError_WrapWrappedErrWithPointerInstanceID(t *testing.T) {
	t.Parallel()

	proto := &ProtoAppError{
		generateID: func() string { return testGeneratedErrorID },
	}

	instanceID := testGeneratedErrorID
	wrappedErr := &AppError{
		errInstanceID: &instanceID,
	}

	got := proto.Wrap(wrappedErr)
	assert.Empty(t, got.instanceID)
	assert.Equal(t, wrappedErr.errInstanceID, got.errInstanceID)
}

func TestProtoAppError_WrapWithStackTrace(t *testing.T) {
	t.Parallel()

	mockStack := &mockStackTrace{}

	proto := &ProtoAppError{
		caller: func() StackTracer {
			return mockStack
		},
	}

	wrappedErr := &AppError{}

	got := proto.Wrap(wrappedErr)
	assert.True(t, got.stackTrace.has)
	assert.Equal(t, mockStack, got.stackTrace.val)
}

func TestProtoAppError_WrapWrappedErrWithHasStackTrace(t *testing.T) {
	t.Parallel()

	mockStack := &mockStackTrace{}

	proto := &ProtoAppError{
		caller: func() StackTracer {
			return mockStack
		},
	}

	wrappedErr := &AppError{
		stackTrace: stackTrace{
			has: true,
		},
	}

	got := proto.Wrap(wrappedErr)
	assert.True(t, got.stackTrace.has)
	assert.Nil(t, got.stackTrace.val)
}

func TestProtoAppError_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		message   string
		argsNames []string
		args      []any
		want      string
	}{
		{
			name:      "test1",
			message:   "my-message",
			argsNames: nil,
			args:      nil,
			want:      "my-message",
		},
		{
			name:      "test2",
			message:   "my-message {{ .key1 }} - {{ .key2 }}",
			argsNames: []string{"key1", "key2"},
			args:      []any{"value1", "value2"},
			want:      "my-message value1 - value2",
		},
		{
			name:      "test3",
			message:   "my-message {{ .key1 }} - {{ .key2 }}",
			argsNames: []string{"key1", "key2"},
			args:      []any{"value1", "value2", "value3"},
			want:      "my-message value1 - value2",
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := &ProtoAppError{
				pureError: pureError{
					message:   tt.message,
					argsNames: tt.argsNames,
					args:      tt.args,
				},
			}
			got := e.Error()
			assert.Equal(t, tt.want, got)
		})
	}
}
