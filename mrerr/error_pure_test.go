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

const (
	expectedCode    = "test-code"
	expectedMessage = "test-message"
)

func Test_pureError_IsTrue(t *testing.T) {
	t.Parallel()

	proto := mrerr.NewProto("test-code", mrerr.ErrorKindUser, "")
	err := proto.New()

	got := proto.Is(err)
	assert.True(t, got)
}

func Test_pureError_IsFalse(t *testing.T) {
	t.Parallel()

	proto := mrerr.NewProto(expectedCode, mrerr.ErrorKindUser, "")
	err := mrerr.NewProto("test-code2", mrerr.ErrorKindUser, "").New()

	got := proto.Is(err)
	assert.False(t, got)

	got = proto.Is(errors.New("external error"))
	assert.False(t, got)
}

func Test_pureError_TranslateKindUser(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	mockTranslator := mock_mrerr.NewMocktranslator(ctrl)

	proto := mrerr.NewProto(expectedCode, mrerr.ErrorKindUser, expectedMessage)
	expectedReason := mrmsg.ErrorMessage{Reason: "test"}

	mockTranslator.
		EXPECT().
		TranslateError(expectedCode, expectedMessage).
		Return(expectedReason)

	got := proto.Translate(mockTranslator)
	assert.Equal(t, expectedReason, got)
}

func Test_pureError_TranslateKindInternal(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	mockTranslator := mock_mrerr.NewMocktranslator(ctrl)

	proto := mrerr.NewProto(expectedCode, mrerr.ErrorKindInternal, "")
	expectedReason := mrmsg.ErrorMessage{Reason: "test"}

	mockTranslator.
		EXPECT().
		HasErrorCode(expectedCode).
		Return(false)

	mockTranslator.
		EXPECT().
		TranslateError(mrerr.ErrorCodeInternal, mrerr.ErrorCodeInternal).
		Return(expectedReason)

	got := proto.Translate(mockTranslator)
	assert.Equal(t, expectedReason, got)
}

func Test_pureError_TranslateKindInternalHasMessage(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	mockTranslator := mock_mrerr.NewMocktranslator(ctrl)

	proto := mrerr.NewProto(expectedCode, mrerr.ErrorKindInternal, expectedMessage)
	expectedReason := mrmsg.ErrorMessage{Reason: "test"}

	mockTranslator.
		EXPECT().
		HasErrorCode(expectedCode).
		Return(true)

	mockTranslator.
		EXPECT().
		TranslateError(expectedCode, expectedMessage).
		Return(expectedReason)

	got := proto.Translate(mockTranslator)
	assert.Equal(t, expectedReason, got)
}

func Test_pureError_TranslateKindSystem(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})

	mockTranslator := mock_mrerr.NewMocktranslator(ctrl)

	proto := mrerr.NewProto(expectedCode, mrerr.ErrorKindSystem, "")
	expectedReason := mrmsg.ErrorMessage{Reason: "test"}

	mockTranslator.
		EXPECT().
		HasErrorCode(expectedCode).
		Return(false)

	mockTranslator.
		EXPECT().
		TranslateError(mrerr.ErrorCodeSystem, mrerr.ErrorCodeSystem).
		Return(expectedReason)

	got := proto.Translate(mockTranslator)
	assert.Equal(t, expectedReason, got)
}
