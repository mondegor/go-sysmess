package mrerrors_test

import (
	"context"
	"errors"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-sysmess/mrerrors"
	mock_mrerrors "github.com/mondegor/go-sysmess/mrerrors/mock"
	"github.com/mondegor/go-sysmess/mrmsg"
)

func TestInstantError_WithAttrs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		paramName  string
		paramValue any
		want       string
	}{
		{
			name:       "test1",
			paramName:  "",
			paramValue: nil,
			want:       `!EMPTYKEY=null`,
		},
		{
			name:       "test2",
			paramName:  "test-name",
			paramValue: "test-value",
			want:       `test-name="test-value"`,
		},
		{
			name:       "test3",
			paramName:  "test-name",
			paramValue: 12345,
			want:       "test-name=12345",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := mrerrors.NewProto("test internal error").New()

			got := e.WithAttr(tt.paramName, tt.paramValue)
			require.ErrorContains(t, got, tt.want)

			got2 := e.WithAttrs(tt.paramName, tt.paramValue)
			assert.ErrorContains(t, got2, tt.want)
		})
	}
}

func TestInstantError_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		message    string
		args       []any
		instanceID string
		attrs      []any
		err        error
		want       string
	}{
		{
			name:    "test1",
			message: "my-message",
			want:    "my-message [INTERNAL]",
		},
		{
			name:       "test2",
			message:    "my-message {Key1} - {Key2}",
			args:       []any{"value1", "value2"},
			instanceID: "test-id",
			attrs: []any{
				"attr1", "value1",
				"attr2", 222222,
			},
			err:  errors.New("test external error"),
			want: `my-message value1 - value2 [INTERNAL, id=test-id, attr1="value1", attr2=222222]: test external error`,
		},
		{
			name:    "test3",
			message: "my-message {Key1} - {Key2}",
			args:    []any{"value1", "value2", "value3"},
			want:    `my-message value1 - value2 [INTERNAL, value3="!MISSINGVALUE"]`,
		},
		{
			name:    "test4",
			message: "my-message {Key1} - {Key2}",
			args:    []any{"value1"},
			attrs: []any{
				"attr1", "attr-value1",
				"attr2",
			},
			want: `my-message value1 - !MISSINGARG [INTERNAL, attr1="attr-value1", attr2="!MISSINGVALUE"]`,
		},
		{
			name:    "test5",
			message: "my-message {Key1} - {Key2}",
			attrs: []any{
				"attr1", "attr-value1",
				"attr2",
			},
			want: `my-message !MISSINGARG - !MISSINGARG [INTERNAL, attr1="attr-value1", attr2="!MISSINGVALUE"]`,
		},
		{
			name:    "test6",
			message: "my-message {Key1} - {Key2}",
			args:    []any{"value1", "value2", "attr1", "attr-value1", "attr2"},
			want:    `my-message value1 - value2 [INTERNAL, attr1="attr-value1", attr2="!MISSINGVALUE"]`,
		},
		{
			name:    "test7",
			message: "my-message {Key1} - {Key2}",
			args:    []any{"value1", "value2", "attr1", "attr-value1", "attr2"},
			attrs: []any{
				222222,
				"attr3",
			},
			want: `my-message value1 - value2 [INTERNAL, attr1="attr-value1", attr2=222222, attr3="!MISSINGVALUE"]`,
		},
		{
			name:    "test8",
			message: "my-message {Key1} - {Key2}",
			args:    []any{"value1", "value2", "attr1", "attr-value1", "attr2"},
			attrs: []any{
				"attr-value2",
				"attr3", 333333,
			},
			want: `my-message value1 - value2 [INTERNAL, attr1="attr-value1", attr2="attr-value2", attr3=333333]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var instErr *mrerrors.InstantError

			e := mrerrors.NewProto(
				tt.message,
				mrerrors.WithProtoArgsReplacer(
					func(message string) mrerrors.MessageReplacer {
						return mrmsg.NewMessageReplacer("{", "}", message)
					},
				),
				mrerrors.WithProtoOnCreated(
					func(ctx context.Context, err error) string { return tt.instanceID },
				),
			)

			if tt.err != nil {
				instErr = e.Wrap(tt.err, tt.args...)
			} else {
				instErr = e.New(tt.args...)
			}

			instErr = instErr.WithAttrs(tt.attrs...)

			got := instErr.Error()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInstantError_NewWithInstanceID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		instanceID string
		want       string
	}{
		{
			name:       "test1",
			instanceID: "",
			want:       "",
		},
		{
			name:       "test2",
			instanceID: "test-id",
			want:       "test-id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := mrerrors.NewProto(
				"",
				mrerrors.WithProtoOnCreated(
					func(ctx context.Context, err error) string { return tt.instanceID },
				),
			).New()
			got := e.ID()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInstantError_WrapWithInstanceID(t *testing.T) {
	t.Parallel()

	e1 := errors.New("wrapped-error")
	e2 := mrerrors.NewProto(
		"",
		mrerrors.WithProtoOnCreated(
			func(ctx context.Context, err error) string { return "test-id" },
		),
	).Wrap(e1)

	got := e2.ID()
	assert.Equal(t, "test-id", got)
}

func TestInstantError_WrapWithWrappedInstanceID(t *testing.T) {
	t.Parallel()

	e1 := mrerrors.NewProto(
		"",
		mrerrors.WithProtoOnCreated(
			func(ctx context.Context, err error) string { return "test-id" },
		),
	).New()
	e2 := mrerrors.NewProto("").Wrap(e1)

	got := e2.ID()
	assert.Equal(t, "test-id", got)
}

func TestInstantError_DoubleWrapWithInstanceID(t *testing.T) {
	t.Parallel()

	e1 := mrerrors.NewProto(
		"",
		mrerrors.WithProtoOnCreated(
			func(ctx context.Context, err error) string { return "test-id1" },
		),
	).New()

	e2 := mrerrors.NewProto(
		"",
		mrerrors.WithProtoOnCreated(
			func(ctx context.Context, err error) string { return "test-id2" },
		),
	).Wrap(e1)

	got := e2.ID()
	assert.Equal(t, "test-id1", got)
}

func TestInstantError_Is(t *testing.T) {
	t.Parallel()

	errInternalTest1 := mrerrors.NewProto("test-message1")
	errInternalTest2 := mrerrors.NewProto("test-message2")

	tests := []struct {
		name   string
		err    *mrerrors.InstantError
		target error
		want   bool
	}{
		{
			name:   "test1",
			err:    errInternalTest1.New(),
			target: errors.New("my-message"),
			want:   false,
		},
		{
			name:   "test2",
			err:    errInternalTest1.New(),
			target: errInternalTest1,
			want:   true,
		},
		{
			name:   "test3",
			err:    errInternalTest1.New(),
			target: errInternalTest2,
			want:   false,
		},
		{
			name:   "test4",
			err:    errInternalTest1.New(),
			target: errInternalTest1.New(),
			want:   true,
		},
		{
			name:   "test5",
			err:    errInternalTest1.New(),
			target: errInternalTest2.New(),
			want:   false,
		},
		{
			name:   "test6",
			err:    errInternalTest2.Wrap(errInternalTest1),
			target: errInternalTest1,
			want:   true,
		},
		{
			name:   "test7",
			err:    errInternalTest2.Wrap(errInternalTest1.New()),
			target: errInternalTest1,
			want:   true,
		},
		{
			name:   "test8",
			err:    errInternalTest1.Wrap(errors.New("my-message")),
			target: errInternalTest1,
			want:   true,
		},
		{
			name:   "test9",
			err:    errInternalTest2.Wrap(errInternalTest1),
			target: errInternalTest1.New(),
			want:   true,
		},
		{
			name:   "test10",
			err:    errInternalTest2.Wrap(errInternalTest1.New()),
			target: errInternalTest1.New(),
			want:   true,
		},
		{
			name:   "test11",
			err:    errInternalTest1.Wrap(errors.New("my-message")),
			target: errInternalTest1.New(),
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := errors.Is(tt.err, tt.target)
			// got := tt.err.Is(tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInstantError_NewWithStackTrace(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	mockStackTracer := mock_mrerrors.NewMockStackTracer(ctrl)

	proto := mrerrors.NewProto(
		"",
		mrerrors.WithProtoCaller(
			func() mrerrors.StackTracer { return mockStackTracer },
		),
	)

	mockStackTracer.
		EXPECT().
		Count().
		Return(1)

	mockStackTracer.
		EXPECT().
		Source(0).
		Return("fun-name", "file-test", 15)

	got := proto.New()
	assert.ErrorContains(t, got, "in [fun-name] file-test:15")
}

func TestInstantError_WrapWithStackTrace(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	mockStackTracer := mock_mrerrors.NewMockStackTracer(ctrl)

	proto := mrerrors.NewProto(
		"",
		mrerrors.WithProtoCaller(
			func() mrerrors.StackTracer { return mockStackTracer },
		),
	)

	mockStackTracer.
		EXPECT().
		Count().
		Return(1)

	mockStackTracer.
		EXPECT().
		Source(0).
		Return("fun-name", "file-test", 15)

	got := proto.Wrap(errors.New("test-error"))
	assert.ErrorContains(t, got, "in [fun-name] file-test:15")
}

func TestInstantError_WrapWithStackTraceTwoLines(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	mockStackTracer := mock_mrerrors.NewMockStackTracer(ctrl)

	proto := mrerrors.NewProto(
		"",
		mrerrors.WithProtoCaller(
			func() mrerrors.StackTracer { return mockStackTracer },
		),
	)

	mockStackTracer.
		EXPECT().
		Count().
		Return(2)

	mockStackTracer.
		EXPECT().
		Source(0).
		Return("fun-name1", "file-test1", 15)

	mockStackTracer.
		EXPECT().
		Source(1).
		Return("fun-name2", "file-test2", 30)

	got := proto.Wrap(errors.New("test-error"))
	assert.ErrorContains(t, got, "in [fun-name1] file-test1:15, [fun-name2] file-test2:30")
}

func TestInstantError_WrapWithWrappedStackTrace(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	mockStackTracer := mock_mrerrors.NewMockStackTracer(ctrl)

	proto1 := mrerrors.NewProto(
		"",
		mrerrors.WithProtoCaller(
			func() mrerrors.StackTracer { return mockStackTracer },
		),
	)

	proto2 := mrerrors.NewProto("")

	mockStackTracer.
		EXPECT().
		Count().
		Return(1)

	mockStackTracer.
		EXPECT().
		Source(0).
		Return("fun-name", "file-test", 15)

	got := proto2.Wrap(proto1.New())
	assert.ErrorContains(t, got, "in [fun-name] file-test:15")
}

func TestInstantError_DoubleWrapWithStackTrace(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	mockStackTracer1 := mock_mrerrors.NewMockStackTracer(ctrl)
	mockStackTracer2 := mock_mrerrors.NewMockStackTracer(ctrl)

	proto1 := mrerrors.NewProto(
		"",
		mrerrors.WithProtoCaller(
			func() mrerrors.StackTracer { return mockStackTracer1 },
		),
	)

	proto2 := mrerrors.NewProto(
		"",
		mrerrors.WithProtoCaller(
			func() mrerrors.StackTracer { return mockStackTracer2 },
		),
	)

	mockStackTracer1.
		EXPECT().
		Source(0).
		Return("fun-name", "file-test", 15)

	gotFunction, gotFile, gotLine := proto2.Wrap(proto1.New()).StackTrace().Source(0)
	assert.Equal(t, "fun-name file-test:15", gotFunction+" "+gotFile+":"+strconv.Itoa(gotLine))
}
