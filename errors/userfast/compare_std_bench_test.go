package userfast_test

import (
	"errors"
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
