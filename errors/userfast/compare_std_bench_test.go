package userfast_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mondegor/go-sysmess/errors/userfast"
)

// ============================================================================
// Сравнение с стандартной библиотекой errors
// ============================================================================

const (
	_testCode    = "TEST_CODE"
	_testMessage = "test message"
)

// ----------------------------------------------------------------------------
// Создание ошибок
// ----------------------------------------------------------------------------

// BenchmarkStd_ErrorsNew - стандартный errors.New.
func BenchmarkStd_ErrorsNew(b *testing.B) {
	testCode := _testCode
	testMessage := _testMessage

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = errors.New("#" + testCode + " - " + testMessage)
	}
}

// BenchmarkUserFast_New - userfast.ProtoError.
func BenchmarkUserFast_New(b *testing.B) {
	testCode := _testCode
	testMessage := _testMessage

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = userfast.New(testCode, testMessage)
	}
}

// BenchmarkStd_ErrorsNewAndError - стандартный errors.New.
func BenchmarkStd_ErrorsNewAndError(b *testing.B) {
	testCode := _testCode
	testMessage := _testMessage

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := errors.New("#" + testCode + " - " + testMessage)
		_ = err.Error()
	}
}

// BenchmarkUserFast_NewAndError - userfast.ProtoError.
func BenchmarkUserFast_NewAndError(b *testing.B) {
	testCode := _testCode
	testMessage := _testMessage

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := userfast.New(testCode, testMessage)
		_ = err.Error()
	}
}

// ----------------------------------------------------------------------------
// Обёртывание ошибок
// ----------------------------------------------------------------------------

// BenchmarkStd_FmtErrorf_Wrap - стандартный fmt.Errorf с обёртыванием.
func BenchmarkStd_FmtErrorf_Wrap(b *testing.B) {
	testCode := _testCode
	testMessage := _testMessage
	baseErr := errors.New("base error")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := fmt.Errorf("#%s - %s: %w", testCode, testMessage, baseErr)
		_ = err.Error()
	}
}

// BenchmarkUserFast_Wrap - userfast.ProtoError.Wrap().
func BenchmarkUserFast_Wrap(b *testing.B) {
	testCode := _testCode
	testMessage := _testMessage
	baseErr := errors.New("base error")
	testUserProto := userfast.New(testCode, testMessage)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := testUserProto.Wrap(baseErr)
		_ = err.Error()
	}
}

// ----------------------------------------------------------------------------
// Проверка типа ошибки (errors.Is)
// ----------------------------------------------------------------------------

// BenchmarkStd_ErrorsIs - стандартный errors.Is.
func BenchmarkStd_ErrorsIs(b *testing.B) {
	err1 := errors.New("#" + _testCode + " - " + _testMessage)
	err2 := errors.New("#" + _testCode + " - " + _testMessage)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = errors.Is(err1, err2)
	}
}

// BenchmarkUserFast_Is - userfast.ProtoError.Is.
func BenchmarkUserFast_Is(b *testing.B) {
	err1 := userfast.New(_testCode, _testMessage)
	err2 := userfast.New(_testCode, _testMessage)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = errors.Is(err1, err2)
	}
}

// BenchmarkStd_ErrorsIsDeep - стандартный errors.Is.
func BenchmarkStd_ErrorsIsDeep(b *testing.B) {
	err1 := errors.New("#" + _testCode + " - " + _testMessage)
	err2 := fmt.Errorf("#%s - %s: %w", _testCode, _testMessage, err1)
	err3 := fmt.Errorf("#%s - %s: %w", _testCode, _testMessage, err2)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = errors.Is(err1, err1)
		_ = errors.Is(err1, err2)
		_ = errors.Is(err1, err3)
		_ = errors.Is(err2, err1)
		_ = errors.Is(err3, err1)
	}
}

// BenchmarkUserFast_IsDeep - userfast.ProtoError.Is.
func BenchmarkUserFast_IsDeep(b *testing.B) {
	errWrapper := userfast.New(_testCode, _testMessage)
	err1 := userfast.New(_testCode, _testMessage)
	err2 := errWrapper.Wrap(err1)
	err3 := errWrapper.Wrap(err2)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = errors.Is(err1, err1)
		_ = errors.Is(err1, err2)
		_ = errors.Is(err1, err3)
		_ = errors.Is(err2, err1)
		_ = errors.Is(err3, err1)
	}
}
