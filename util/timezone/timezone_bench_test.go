package timezone_test

import (
	"testing"
	"time"

	"github.com/mondegor/go-core/util/timezone"
)

// ============================================================================
// Стоимость создания списка часовых поясов
// ============================================================================

// _benchNames - пояса и с переходом на летнее время, и без него,
// чтобы замер не приходился на один частный случай.
var _benchNames = []string{ //nolint:gochecknoglobals
	"Europe/Berlin",
	"Asia/Tokyo",
	"Australia/Sydney",
	"America/New_York",
}

// _benchNames16 - расширенный набор для проверки линейности стоимости.
var _benchNames16 = []string{ //nolint:gochecknoglobals
	"Europe/Berlin",
	"Asia/Tokyo",
	"Australia/Sydney",
	"America/New_York",
	"Europe/Moscow",
	"Europe/London",
	"Europe/Paris",
	"Europe/Lisbon",
	"America/Chicago",
	"America/Denver",
	"America/Los_Angeles",
	"America/Sao_Paulo",
	"Asia/Shanghai",
	"Asia/Kolkata",
	"Africa/Cairo",
	"Pacific/Auckland",
}

// ----------------------------------------------------------------------------
// Создание списка
// ----------------------------------------------------------------------------

// BenchmarkStd_LoadLocation - базовая линия: только загрузка поясов,
// без построения индекса подбора.
func BenchmarkStd_LoadLocation(b *testing.B) {
	names := _benchNames

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, name := range names {
			_, _ = time.LoadLocation(name)
		}
	}
}

// BenchmarkNewLocationList - загрузка поясов вместе с построением индекса.
// Разница с BenchmarkStd_LoadLocation - это цена индекса.
func BenchmarkNewLocationList(b *testing.B) {
	names := _benchNames

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = timezone.NewLocationList(names)
	}
}

// BenchmarkNewLocationList_ByCount - проверяет заявленную линейность
// стоимости создания по числу указанных поясов.
func BenchmarkNewLocationList_ByCount(b *testing.B) {
	benchmarks := []struct {
		name  string
		names []string
	}{
		{name: "1", names: _benchNames16[:1]},
		{name: "4", names: _benchNames16[:4]},
		{name: "16", names: _benchNames16[:16]},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			names := bm.names

			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_ = timezone.NewLocationList(names)
			}
		})
	}
}

// ----------------------------------------------------------------------------
// Обращения к готовому списку
// ----------------------------------------------------------------------------

// BenchmarkLocationList_NameByOffset - горячий путь: подбор по индексу
// должен стоить одно обращение к мапе.
func BenchmarkLocationList_NameByOffset(b *testing.B) {
	list := timezone.NewLocationList(_benchNames)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = list.NameByOffset(9*time.Hour, false)
	}
}

// BenchmarkLocationList_LocationByName - поиск пояса по имени.
func BenchmarkLocationList_LocationByName(b *testing.B) {
	list := timezone.NewLocationList(_benchNames)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = list.LocationByName("Asia/Tokyo")
	}
}
