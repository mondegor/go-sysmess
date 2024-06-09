package mrerr

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock_mrerr "github.com/mondegor/go-sysmess/mrerr/mock"
	"github.com/mondegor/go-sysmess/mrmsg"
)

//go:generate mockgen -source=error_pure.go -destination=./mock/error_pure.go

func Test_pureError_IsTrue(t *testing.T) {
	t.Parallel()

	proto := &ProtoAppError{
		pureError: pureError{
			code: "test-code",
		},
	}

	err := &AppError{
		pureError: pureError{
			code: "test-code",
		},
	}

	got := proto.Is(err)
	assert.True(t, got)
}

func Test_pureError_IsFalse(t *testing.T) {
	t.Parallel()

	proto := &ProtoAppError{
		pureError: pureError{
			code: "test-code",
		},
	}

	err := &AppError{
		pureError: pureError{
			code: "test-code2",
		},
	}

	got := proto.Is(err)
	assert.False(t, got)

	got = proto.Is(errors.New("external error"))
	assert.False(t, got)
}

func Test_pureError_TranslateKindUser(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTranslator := mock_mrerr.NewMocktranslator(ctrl)

	proto := &ProtoAppError{
		pureError: pureError{
			code:    "test-code",
			kind:    ErrorKindUser,
			message: "test-message",
		},
	}

	expected := mrmsg.ErrorMessage{Reason: "test"}

	mockTranslator.
		EXPECT().
		TranslateError(proto.code, proto.message).
		Return(expected)

	got := proto.Translate(mockTranslator)
	assert.Equal(t, expected, got)
}

func Test_pureError_TranslateKindInternal(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTranslator := mock_mrerr.NewMocktranslator(ctrl)

	proto := &ProtoAppError{
		pureError: pureError{
			code: "test-code",
			kind: ErrorKindInternal,
		},
	}

	expected := mrmsg.ErrorMessage{Reason: "test"}

	mockTranslator.
		EXPECT().
		HasErrorCode(proto.code).
		Return(false)

	mockTranslator.
		EXPECT().
		TranslateError(ErrorCodeInternal, ErrorCodeInternal).
		Return(expected)

	got := proto.Translate(mockTranslator)
	assert.Equal(t, expected, got)
}

func Test_pureError_TranslateKindInternalHasMessage(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTranslator := mock_mrerr.NewMocktranslator(ctrl)

	proto := &ProtoAppError{
		pureError: pureError{
			code:    "test-code",
			kind:    ErrorKindInternal,
			message: "test-message",
		},
	}

	expected := mrmsg.ErrorMessage{Reason: "test"}

	mockTranslator.
		EXPECT().
		HasErrorCode(proto.code).
		Return(true)

	mockTranslator.
		EXPECT().
		TranslateError(proto.code, proto.message).
		Return(expected)

	got := proto.Translate(mockTranslator)
	assert.Equal(t, expected, got)
}

func Test_pureError_TranslateKindSystem(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTranslator := mock_mrerr.NewMocktranslator(ctrl)

	proto := &ProtoAppError{
		pureError: pureError{
			kind: ErrorKindSystem,
		},
	}

	expected := mrmsg.ErrorMessage{Reason: "test"}

	mockTranslator.
		EXPECT().
		HasErrorCode(proto.code).
		Return(false)

	mockTranslator.
		EXPECT().
		TranslateError(ErrorCodeSystem, ErrorCodeSystem).
		Return(expected)

	got := proto.Translate(mockTranslator)
	assert.Equal(t, expected, got)
}
