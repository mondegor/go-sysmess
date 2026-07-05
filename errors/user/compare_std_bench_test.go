package user_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mondegor/go-core/errors/user"
)

// ============================================================================
// Сравнение с стандартной библиотекой errors
// ============================================================================

const (
	_testCode    = "TEST_CODE"
	_testMessage = "test message"
)

var errTestUserProto = user.New("TEST_CODE", "test message")

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
		err := errors.New("#" + testCode + " - " + testMessage)
		_ = err.Error()
	}
}

// BenchmarkUserProto_New - user.ProtoError.New() без аргументов.
func BenchmarkUserProto_New(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := errTestUserProto.New()
		_ = err.Error()
	}
}

// ----------------------------------------------------------------------------
// Создание ошибок с аргументами
// ----------------------------------------------------------------------------

// BenchmarkStd_FmtErrorf_WithArgs - стандартный fmt.Errorf с аргументами.
func BenchmarkStd_FmtErrorf_WithArgs(b *testing.B) {
	testCode := _testCode
	testMessage := _testMessage

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := fmt.Errorf("#%s - %s with %s and %d", testCode, testMessage, "stringx", 42)
		_ = err.Error()
	}
}

// BenchmarkUserProto_New_WithArgs - user.ProtoError.New() с аргументами.
func BenchmarkUserProto_New_WithArgs(b *testing.B) {
	testUserProtoWithArgs := user.New(_testCode, _testMessage+" with {A} and {B}")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := testUserProtoWithArgs.New("stringx", 42)
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

// BenchmarkUserProto_Wrap - user.ProtoError.Wrap().
func BenchmarkUserProto_Wrap(b *testing.B) {
	baseErr := errors.New("base error")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := errTestUserProto.Wrap(baseErr)
		_ = err.Error()
	}
}

// ----------------------------------------------------------------------------
// Проверка типа ошибки (errors.Is)
// ----------------------------------------------------------------------------

// BenchmarkStd_ErrorsIs - стандартный errors.Is.
func BenchmarkStd_ErrorsIs(b *testing.B) {
	err1 := errors.New("test error")
	err2 := errors.New("test error")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = errors.Is(err1, err2)
	}
}

// BenchmarkUserProto_Is - user.ProtoError.Is.
func BenchmarkUserProto_Is(b *testing.B) {
	err1 := errTestUserProto.New()
	err2 := errTestUserProto.New()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = errors.Is(err1, err2)
	}
}

// ----------------------------------------------------------------------------
// Получение метаданных
// ----------------------------------------------------------------------------

// BenchmarkUserProto_Code - получение кода ошибки.
func BenchmarkUserProto_Code(b *testing.B) {
	err := errTestUserProto.New()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = err.(interface{ Code() string }).Code()
	}
}

// BenchmarkUserProto_Args - получение аргументов.
func BenchmarkUserProto_Args(b *testing.B) {
	err := errTestUserProto.New("arg1", "arg2")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = err.(interface{ Args() []any }).Args()
	}
}

// BenchmarkUserProto_Message - получение сообщения.
func BenchmarkUserProto_Message(b *testing.B) {
	err := errTestUserProto.New()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = err.(interface{ Message() string }).Message()
	}
}
