package mrerr_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/mondegor/go-sysmess/mrerr"
	mock_mrerr "github.com/mondegor/go-sysmess/mrerr/mock"
	"github.com/mondegor/go-sysmess/mrmsg"
)

func TestAppError_WithAttr(t *testing.T) {
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
			want:       "unnamed=<nil>",
		},
		{
			name:       "test2",
			paramName:  "test-name",
			paramValue: "test-value",
			want:       "test-name=test-value",
		},
		{
			name:       "test3",
			paramName:  "test-name",
			paramValue: 12345,
			want:       "test-name=12345",
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := &mrerr.AppError{}

			got := e.WithAttr(tt.paramName, tt.paramValue)
			assert.ErrorContains(t, got, tt.want)
		})
	}
}

func TestAppError_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		message    string
		args       []any
		instanceID string
		attrs      []mrmsg.NamedArg
		err        error
		want       string
	}{
		{
			name:    "test1",
			message: "my-message",
			want:    "my-message",
		},
		{
			name:       "test2",
			message:    "my-message {{ .key1 }} - {{ .key2 }}",
			args:       []any{"value1", "value2"},
			instanceID: "test-id",
			attrs: []mrmsg.NamedArg{
				{
					Name:  "attr1",
					Value: "value2",
				},
				{
					Name:  "attr1",
					Value: "value2",
				},
			},
			err:  errors.New("test external error"),
			want: "[test-id] my-message value1 - value2 (attr1=value2, attr1=value2): test external error",
		},
		{
			name:    "test3",
			message: "my-message {{ .key1 }} - {{ .key2 }}",
			args:    []any{"value1", "value2", "value3"},
			want:    "[WARNING!!! too many arguments in error message] my-message value1 - value2",
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var appErr *mrerr.AppError

			e := mrerr.NewProtoWithExtra(
				"",
				mrerr.ErrorKindInternal,
				tt.message,
				mrerr.ProtoExtra{
					Caller:    nil,
					OnCreated: func(err *mrerr.AppError) string { return tt.instanceID },
				},
			)

			if tt.err != nil {
				appErr = e.Wrap(tt.err, tt.args...)
			} else {
				appErr = e.New(tt.args...)
			}

			for _, attr := range tt.attrs {
				appErr = appErr.WithAttr(attr.Name, attr.Value)
			}

			got := appErr.Error()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAppError_NewWithInstanceID(t *testing.T) {
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
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := mrerr.NewProtoWithExtra(
				"",
				mrerr.ErrorKindInternal,
				"",
				mrerr.ProtoExtra{
					Caller:    nil,
					OnCreated: func(err *mrerr.AppError) string { return tt.instanceID },
				},
			).New()
			got := e.InstanceID()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAppError_WrapWithInstanceID(t *testing.T) {
	t.Parallel()

	e1 := errors.New("wrapped-error")
	e2 := mrerr.NewProtoWithExtra(
		"",
		mrerr.ErrorKindInternal,
		"",
		mrerr.ProtoExtra{
			Caller:    nil,
			OnCreated: func(err *mrerr.AppError) string { return "test-id" },
		},
	).Wrap(e1)

	got := e2.InstanceID()
	assert.Equal(t, "test-id", got)
}

func TestAppError_WrapWithWrappedInstanceID(t *testing.T) {
	t.Parallel()

	e1 := mrerr.NewProtoWithExtra(
		"",
		mrerr.ErrorKindInternal,
		"",
		mrerr.ProtoExtra{
			Caller:    nil,
			OnCreated: func(err *mrerr.AppError) string { return "test-id" },
		},
	).New()
	e2 := mrerr.NewProto("", mrerr.ErrorKindInternal, "").Wrap(e1)

	got := e2.InstanceID()
	assert.Equal(t, "test-id", got)
}

func TestAppError_WrapWithTripleInstanceID(t *testing.T) {
	t.Parallel()

	e1 := mrerr.NewProtoWithExtra(
		"",
		mrerr.ErrorKindInternal,
		"",
		mrerr.ProtoExtra{
			Caller:    nil,
			OnCreated: func(err *mrerr.AppError) string { return "test-id1" },
		},
	).New()

	e2 := mrerr.NewProtoWithExtra(
		"",
		mrerr.ErrorKindInternal,
		"",
		mrerr.ProtoExtra{
			Caller:    nil,
			OnCreated: func(err *mrerr.AppError) string { return "test-id2" },
		},
	).Wrap(e1)

	e3 := mrerr.NewProtoWithExtra(
		"",
		mrerr.ErrorKindInternal,
		"",
		mrerr.ProtoExtra{
			Caller:    nil,
			OnCreated: func(err *mrerr.AppError) string { return "test-id3" },
		},
	).Wrap(e2)

	got := e3.InstanceID()
	assert.Equal(t, "test-id1", got)
}

func TestAppError_NewWithStackTrace(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStackTracer := mock_mrerr.NewMockStackTracer(ctrl)

	proto := mrerr.NewProtoWithExtra(
		"",
		mrerr.ErrorKindInternal,
		"",
		mrerr.ProtoExtra{
			Caller:    func() mrerr.StackTracer { return mockStackTracer },
			OnCreated: nil,
		},
	)

	mockStackTracer.
		EXPECT().
		Count().
		Return(1)

	mockStackTracer.
		EXPECT().
		Item(0).
		Return("fun-name", "file-test", 15)

	got := proto.New()
	assert.ErrorContains(t, got, "in [fun-name] file-test:15")
}

func TestAppError_WrapWithStackTrace(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStackTracer := mock_mrerr.NewMockStackTracer(ctrl)

	proto := mrerr.NewProtoWithExtra(
		"",
		mrerr.ErrorKindInternal,
		"",
		mrerr.ProtoExtra{
			Caller:    func() mrerr.StackTracer { return mockStackTracer },
			OnCreated: nil,
		},
	)

	mockStackTracer.
		EXPECT().
		Count().
		Return(1)

	mockStackTracer.
		EXPECT().
		Item(0).
		Return("fun-name", "file-test", 15)

	got := proto.Wrap(errors.New("test-error"))
	assert.ErrorContains(t, got, "in [fun-name] file-test:15")
}

func TestAppError_WrapWithStackTraceTwoLines(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStackTracer := mock_mrerr.NewMockStackTracer(ctrl)

	proto := mrerr.NewProtoWithExtra(
		"",
		mrerr.ErrorKindInternal,
		"",
		mrerr.ProtoExtra{
			Caller:    func() mrerr.StackTracer { return mockStackTracer },
			OnCreated: nil,
		},
	)

	mockStackTracer.
		EXPECT().
		Count().
		Return(2)

	mockStackTracer.
		EXPECT().
		Item(0).
		Return("fun-name1", "file-test1", 15)

	mockStackTracer.
		EXPECT().
		Item(1).
		Return("fun-name2", "file-test2", 30)

	got := proto.Wrap(errors.New("test-error"))
	assert.ErrorContains(t, got, "in [fun-name1] file-test1:15 , [fun-name2] file-test2:30")
}

func TestAppError_WrapWithWrappedStackTrace(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStackTracer := mock_mrerr.NewMockStackTracer(ctrl)

	proto1 := mrerr.NewProtoWithExtra(
		"",
		mrerr.ErrorKindInternal,
		"",
		mrerr.ProtoExtra{
			Caller:    func() mrerr.StackTracer { return mockStackTracer },
			OnCreated: nil,
		},
	)

	proto2 := mrerr.NewProto("", mrerr.ErrorKindInternal, "")

	mockStackTracer.
		EXPECT().
		Count().
		Return(1)

	mockStackTracer.
		EXPECT().
		Item(0).
		Return("fun-name", "file-test", 15)

	got := proto2.Wrap(proto1.New())
	assert.ErrorContains(t, got, "in [fun-name] file-test:15")
}

func TestAppError_WrapWithDoubleStackTrace(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStackTracer1 := mock_mrerr.NewMockStackTracer(ctrl)
	mockStackTracer2 := mock_mrerr.NewMockStackTracer(ctrl)

	proto1 := mrerr.NewProtoWithExtra(
		"",
		mrerr.ErrorKindInternal,
		"",
		mrerr.ProtoExtra{
			Caller:    func() mrerr.StackTracer { return mockStackTracer1 },
			OnCreated: nil,
		},
	)

	proto2 := mrerr.NewProtoWithExtra(
		"",
		mrerr.ErrorKindInternal,
		"",
		mrerr.ProtoExtra{
			Caller:    func() mrerr.StackTracer { return mockStackTracer2 },
			OnCreated: nil,
		},
	)

	mockStackTracer1.
		EXPECT().
		Count().
		Return(1)

	mockStackTracer1.
		EXPECT().
		Item(0).
		Return("fun-name", "file-test", 15)

	got := proto2.Wrap(proto1.New())
	assert.ErrorContains(t, got, "in [fun-name] file-test:15")
}
