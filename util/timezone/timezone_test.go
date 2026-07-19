package timezone_test

// база часовых поясов встраивается в тестовый бинарник, чтобы тесты
// проходили в минимальных образах, где она отсутствует в системе.
import (
	"testing"
	"time"
	_ "time/tzdata"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mondegor/go-core/util/timezone"
)

func TestLocationList_LocationByName(t *testing.T) {
	t.Parallel()

	list := timezone.NewLocationList([]string{"Europe/Moscow"})

	tests := []struct {
		name    string
		value   string
		want    *time.Location
		wantErr bool
	}{
		{
			name:  "registered iana name",
			value: "Europe/Moscow",
			want:  mustLoadLocation(t, "Europe/Moscow"),
		},
		{
			name:  "utc is registered by default",
			value: "UTC",
			want:  time.UTC,
		},
		{
			name:  "local is registered by default",
			value: "Local",
			want:  time.Local,
		},
		{
			name:    "empty value returns utc and error",
			value:   "",
			want:    time.UTC,
			wantErr: true,
		},
		{
			name:    "known but not registered name returns utc and error",
			value:   "Asia/Tokyo",
			want:    time.UTC,
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := list.LocationByName(tc.value)

			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, tc.want, got)
		})
	}
}

// TestNewLocationList_SkipsInvalidNames - фиксирует, что негодные имена
// пропускаются, а не прерывают создание списка: имя, стоящее после негодного,
// должно быть зарегистрировано. Отвергать такой список - забота
// config.ValidateTimeZones, конструктор ошибку не возвращает.
func TestNewLocationList_SkipsInvalidNames(t *testing.T) {
	t.Parallel()

	list := timezone.NewLocationList([]string{"", "Nowhere/Nowhere", "Europe/Moscow"})

	t.Run("valid name after invalid ones is registered", func(t *testing.T) {
		t.Parallel()

		got, err := list.LocationByName("Europe/Moscow")

		require.NoError(t, err)
		assert.Equal(t, mustLoadLocation(t, "Europe/Moscow"), got)
	})

	t.Run("valid name after invalid ones is indexed by offset", func(t *testing.T) {
		t.Parallel()

		// Москва: +03:00 круглый год, перехода на летнее время нет
		got, ok := list.NameByOffset(3*time.Hour, false)

		require.True(t, ok)
		assert.Equal(t, "Europe/Moscow", got)
	})

	t.Run("unknown name is not registered", func(t *testing.T) {
		t.Parallel()

		got, err := list.LocationByName("Nowhere/Nowhere")

		require.Error(t, err)
		assert.Equal(t, time.UTC, got)
	})

	t.Run("always available names are registered", func(t *testing.T) {
		t.Parallel()

		utc, err := list.LocationByName("UTC")
		require.NoError(t, err)
		assert.Equal(t, time.UTC, utc)

		local, err := list.LocationByName("Local")
		require.NoError(t, err)
		assert.Equal(t, time.Local, local)
	})
}

// mustLoadLocation - загружает часовой пояс или прерывает тест.
func mustLoadLocation(t *testing.T, name string) *time.Location {
	t.Helper()

	loc, err := time.LoadLocation(name)
	require.NoError(t, err)

	return loc
}
